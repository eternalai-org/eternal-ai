from typing import Any, List, Optional, Dict
import httpx
from langchain.schema import BaseMessage, ChatGeneration, AIMessage
from langchain_community.adapters.openai import convert_message_to_dict
from asyncio import sleep as async_sleep
import time
import json
import logging
from x_content.llm.base import OnchainInferResult
from x_content.llm.time_estimation import ModelInferTimeEstimation
from .base import OpenAILLMBase
from typing import Dict
from functools import lru_cache
from .utils import check_and_get_infer_result, ServerInferenceResult
from x_content.wrappers.log_decorators import log_function_call, compress_kwargs
from x_content import constants as C


logger = logging.getLogger(__name__)


@lru_cache(maxsize=1)
def get_time_estimation():
    return ModelInferTimeEstimation()


# TODO: add retry logic
class ASyncBasedEternalAI(OpenAILLMBase):
    chain_id: str
    agent_contract_id: str
    metadata: dict
    timeout_seconds: int = 60 * 60 * 3  # expire in 3 hours
    failed_count_limt: int = 5

    def __init__(self, *args, **kwargs):
        super(ASyncBasedEternalAI, self).__init__(*args, **kwargs)

    async def submit_async_request(
        self,
        messages: List[BaseMessage],
        **kwargs,
    ) -> str:

        openai_messages = [
            convert_message_to_dict(m) if not isinstance(m, dict) else m
            for m in messages
        ]

        headers = {
            "Authorization": self.openai_api_key.get_secret_value(),
            "Content-Type": "application/json",
        }

        json_data = {
            "model": self.model_name,
            "messages": openai_messages,
            "temperature": self.temperature,
            "max_tokens": self.max_tokens,
            "top_p": 1.0,
            "presence_penalty": 0.1,
            "n": self.n,
            "logit_bias": None,
            "frequency_penalty": self.frequency_penalty,
            "seed": self.seed,
            "chain_id": self.chain_id,
            "contract_agent_id": self.agent_contract_id,
            "meta_data": self.metadata,
        }

        json_data.update(kwargs)

        url = f"{self.openai_api_base}/agent/async-batch-prompt"

        async with httpx.AsyncClient() as client:
            response = await client.post(
                url,
                headers=headers,
                json=json_data,
                timeout=httpx.Timeout(60.0),
            )

        if response.status_code != 200:
            raise ValueError(
                f"Failed to send request to '{url}'; code: {response.status_code}"
            )

        response_json = response.json()

        if (
            response_json["status"] < 0
            or response_json["data"].get("id") is None
        ):
            raise Exception("Inference request submission failed")

        receipt = response_json["data"]["id"]

        logger.info(
            f"Submitted async request; Receipt: {receipt}; Calldata: {compress_kwargs(**json_data)}"
        )

        return receipt

    async def wait(self, receipt: str, eta_seconds: float = 60):
        started_at = time.time()

        headers = {
            "Authorization": self.openai_api_key.get_secret_value(),
            "Content-Type": "application/json",
        }

        url = f"{self.openai_api_base}/agent/get-batch-item-output/{receipt}"

        while True:
            # step 3: sleep for a while
            await async_sleep(eta_seconds)
            eta_seconds = max(eta_seconds * 0.1, 5)

            # step 0: check if the task is timed out
            current_time = time.time()

            check_result: ServerInferenceResult = (
                await check_and_get_infer_result(url, headers)
            )

            if not check_result.skipped:
                resp = check_result.response

                # step 2: check and parse the response
                if resp["status"] < 0 or resp["data"]["status"] == "error":
                    raise Exception(
                        f"Inference request failed; Receipt: {receipt}; Raw Output: {resp}"
                    )

                if resp["data"]["status"] == "queue-handled":
                    tx_hash = resp["data"]["inscribe_tx_hash"]

                    try:
                        prompt_output = json.loads(
                            resp["data"]["prompt_output"]
                        )
                    except json.JSONDecodeError:
                        raise Exception(
                            f"Failed to decode prompt output; Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {resp}"
                        )

                    choices: List[Dict[str, Any]] = prompt_output.get(
                        "choices", []
                    )
                    if not choices:
                        raise ValueError(
                            f"No choices found in the OpenAI response; Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {prompt_output}"
                        )

                    if (
                        choices[0].get("message") is None
                        or choices[0].get("message", {}).get("content") is None
                    ):
                        raise ValueError(
                            f"Bad response from LLM-server. Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {prompt_output}"
                        )

                    logger.info(
                        f"[ASyncBasedEternalAI.wait] Inference request succeeded; Model: {self.model_name}; Receipt: {receipt}"
                    )

                    content = choices[0].get("message", {}).get("content", "")
                    token_usage = prompt_output.get("usage", {})

                    return {
                        "message": AIMessage(content=content),
                        "token_usage": token_usage,
                        "tx_hash": tx_hash,
                    }

            if current_time - started_at > self.timeout_seconds:
                logger.error(
                    f"[ASyncBasedEternalAI.wait] Inference request timed out; Model: {self.model_name}; Receipt: {receipt}"
                )
                raise Exception(
                    f"Inference request timed out; Model: {self.model_name}; Receipt: {receipt}"
                )

    @log_function_call
    async def agenerate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> OnchainInferResult:

        receipt: str = await self.submit_async_request(messages, **kwargs)
        submit_time = time.time()
        estimation = get_time_estimation()

        try:
            result: dict = await self.wait(
                receipt, estimation.estimate(self.chain_id)
            )

        finally:
            estimation.update(self.chain_id, time.time() - submit_time)

        if "message" not in result:
            raise ValueError(f"Unexpected response from OpenAI: {result}")

        # Return a ChatResult
        return OnchainInferResult(
            generations=[ChatGeneration(message=result["message"])],
            llm_output={"token_usage": result.get("token_usage", {})},
            tx_hash=result.get("tx_hash", ""),
            receipt=receipt,
        )

    def get_info(self):
        return {
            "model": self.model_name,
            "temperature": self.temperature,
            "max_tokens": self.max_tokens,
            "top_p": 1.0,
            "presence_penalty": 0.1,
            "n": self.n,
            "logit_bias": None,
            "frequency_penalty": self.frequency_penalty,
            "seed": self.seed,
            "chain_id": self.chain_id,
            "agent_contract_id": self.agent_contract_id,
            "metadata": self.metadata,
            "timeout_seconds": self.timeout_seconds,
            "failed_count_limt": self.failed_count_limt,
        }
