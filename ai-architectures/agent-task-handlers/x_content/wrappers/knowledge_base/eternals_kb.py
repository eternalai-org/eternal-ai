from .base import KnowledgeBase, RelatedInformation
import httpx
import time
import json
import logging
from asyncio import sleep as async_sleep

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)


class EternalKnowledgeBase(KnowledgeBase):
    DEFAULT_TIMEOUT_SECONDS: int = 1800

    async def submit_async_request(self, query: str, top_k: int = None):
        if len(self.kbs) == 0:
            return []

        url = f"{self.base_url}/agent/async-batch-knowledge"

        headers = {
            "Authorization": self.api_key,
        }

        body = {
            "query": query,
            "top_k": top_k or self.default_top_k,
            "kb": [self.kbs[0].kb_id],
            "chain_id": self.kbs[0].chain_id,
        }

        async with httpx.AsyncClient() as client:
            response = await client.post(
                url,
                headers=headers,
                json=body,
                timeout=httpx.Timeout(60.0),
            )

        response_json = response.json()

        if (
            response_json["status"] < 0
            or response_json["data"].get("id") is None
        ):
            raise Exception(
                f"Inference request submission failed: {json.dumps(response_json)}"
            )

        receipt = response_json["data"]["id"]
        return receipt

    async def wait(
        self, receipt: str, threshold: float = None, eta_seconds: float = 60
    ) -> list[RelatedInformation]:
        started_at = time.time()

        headers = {
            "Authorization": self.api_key,
            "Content-Type": "application/json",
        }

        url = f"{self.base_url}/agent/get-batch-item-output/{receipt}"

        while True:
            # step 3: sleep for a while
            await async_sleep(eta_seconds)
            eta_seconds = max(eta_seconds * 0.1, 5)

            # step 0: check if the task is timed out
            current_time = time.time()

            if current_time - started_at > self.DEFAULT_TIMEOUT_SECONDS:
                raise Exception(f"Inference request timed out; Receipt: {receipt}")

            async with httpx.AsyncClient() as client:
                response = await client.get(
                    url, 
                    headers=headers, 
                    timeout=httpx.Timeout(60.0)
                )

            if response.status_code != 200:
                continue

            resp = response.json()

            # step 2: check and parse the response
            if resp["status"] < 0 or resp["data"]["status"] == "error":
                raise Exception(
                    f"Retrieval request failed; Receipt: {receipt}; Raw Output: {resp}"
                )

            if resp["data"]["status"] == "queue-handled":
                tx_hash = resp["data"]["inscribe_tx_hash"]

                try:
                    prompt_output = json.loads(resp["data"]["prompt_output"])
                    assert "result" in prompt_output and isinstance(
                        prompt_output["result"], list
                    ), "Expect 'result' key in prompt output with list value"

                    validated_models = [
                        RelatedInformation.model_validate(m)
                        for m in prompt_output.get("result", [])
                    ]

                    validated_models = [
                        e
                        for e in validated_models
                        if e.score >= (threshold or self.similarity_threshold)
                    ]

                    return validated_models

                except json.JSONDecodeError:
                    raise Exception(
                        f"Failed to decode prompt output; Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {resp}"
                    )

                except Exception as e:
                    raise Exception(
                        f"Knowledge retrieval failed; Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {resp}"
                    )

    async def aretrieve(
        self,
        query: str,
        top_k: int = None,
        threshold: float = None,
        *args,
        **kwargs,
    ):
        receipt = await self.submit_async_request(query, top_k)
        assert receipt is not None, "Expect non-None receipt"
        logger.info(
            f"Submitted async request; Receipt: {receipt}; query: {json.dumps(query)[:100]}, kbs: {self.kbs}, top_k: {top_k}, threshold: {threshold}"
        )
        return await self.wait(receipt, threshold, *args, **kwargs)
