from x_content.constants import MissionChainState, ModelName
from x_content.llm.base import OpenAILLMBase
from x_content.llm.eternal_ai import ASyncBasedEternalAI
from x_content.llm.local import SyncBasedEternalAI
from x_content.models import AgentKnowledgeBase, AutoAgentTask, ChatRequest, ReasoningLog, ToolDef
from x_content.toolcall import dynamic_toolcall, wrapped_external_apis
from x_content.utils import parse_knowledge_ids
from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.knowledge_base.eternals_kb import EternalKnowledgeBase
from x_content.wrappers import telegram
from x_content import constants as const

import random
import logging
import json
import requests
from typing import List
from x_content import constants as const

from x_content.wrappers.knowledge_base.local import KBStore
from x_content.wrappers.twin_agent import get_random_example_tweets

from x_content.wrappers.api import twitter_v2

logger = logging.getLogger(__name__)

def move_state(
    log: AutoAgentTask, state: MissionChainState, reason: str
):
    log.system_message = reason
    log.state = state
    return log

async def a_move_state(
    log: AutoAgentTask, state: MissionChainState, reason: str
):
    log.system_message = reason
    log.state = state
    return log


async def notify_status_reasoning_log(log: ReasoningLog):
    nav = f"<b>Request-ID</b>: {log.id};</i>"

    if log.state == MissionChainState.NEW and not log.is_resubmit:
        info = f"<i><b>Ref-ID</b>: {log.meta_data.ref_id};\n{nav}"
        msg = f"<strong>Received a new task {log.task} using toolset {log.toolset} for {log.meta_data.twitter_username}</strong>\nTraceback info:\n{info}"

    elif log.state in [MissionChainState.ERROR, MissionChainState.DONE]:
        info = f"<i><b>Ref-ID</b>: {log.meta_data.ref_id};\n{nav}"
        system_message = log.system_message

        # if log.state == MissionChainState.ERROR:
        #     await telegram.a_send_message(
        #         twitter_username="junk_notifications",
        #         room=telegram.TELEGRAM_ALERT_ROOM,
        #         schedule=True,
        #         message_to_send=f"<b>Request</b>: {log.id}\nModel: {log.model}; Task: {log.task}; Agent: {log.meta_data.twitter_username}\nError: {system_message}",
        #         fmt="HTML"
        #     )

        msg = f"<strong>Task {log.task} using toolset {log.toolset} for {log.meta_data.twitter_username} finished with state {log.state}</strong>\nTraceback info:\n{info}\nSystem message: {system_message}"

    else:
        return

    await telegram.a_send_message(
        twitter_username="junk_notifications",
        message_to_send=msg,
        schedule=True,
        room=telegram.TELEGRAM_TASK_IO_NOTIFICATION_ROOM,
        fmt="HTML",
    )


async def notify_status_chat_request(request: ChatRequest):
    nav = f"<b>Request-ID</b>: {request.id};</i>"

    task_name = "chat"
    if request.state == MissionChainState.NEW and not request.is_resubmit:
        info = f"<i><b>Ref-ID</b>: {request.meta_data.ref_id};\n{nav}"
        msg = f"<strong>Received a new task {task_name} for {request.meta_data.twitter_username}</strong>\nTraceback info:\n{info}"

    elif request.state in [MissionChainState.ERROR, MissionChainState.DONE]:
        info = f"<i><b>Ref-ID</b>: {request.meta_data.ref_id};\n{nav}"
        system_message = request.system_message
        msg = f"<strong>Task {task_name} for {request.meta_data.twitter_username} finished with state {request.state}</strong>\nTraceback info:\n{info}\nSystem message: {system_message}"

    else:
        return

    await telegram.a_send_message(
        twitter_username="junk_notifications",
        message_to_send=msg,
        schedule=True,
        room=telegram.TELEGRAM_TASK_IO_NOTIFICATION_ROOM,
        fmt="HTML",
    )


_alert_template = """
<strong>Alert</strong>
<i>Task {log.task} using toolset {log.toolset} for {log.meta_data.twitter_username} raised an alert</i>
{info}
<i>Reason: </i>
<pre>\n{reason}\n</pre>  
"""


async def send_alert(task: AutoAgentTask, reason: str):
    global _alert_template

    nav = f"<b>Request-ID</b>: {task.id};</i>"
    info = f"<i><b>Ref-ID</b>: {task.meta_data.ref_id};\n{nav}"

    msg = _alert_template.format(log=task, info=info, reason=reason)
    await telegram.a_send_message(
        twitter_username="junk_notifications",
        message_to_send=msg,
        schedule=True,
        room=telegram.TELEGRAM_ALERT_ROOM,
        fmt="HTML",
    )


def format_prompt_v2(log: ReasoningLog, tools: List[ToolDef]):
    template_prompt = """
You have access to the following tools to get information or take actions:
{tools}

Your reply must be a single JSON object with exactly three keys described as follows.
thought: your own thought about the next step, reflecting your unique persona.
action: must be one of {toolnames}.
action_input: provide the necessary parameters for the chosen action, separating multiple parameters with the | character. For example, if there are two parameters "abc" and "123", the action_input field should be "abc|123".

OR with exactly two keys as follows.
thought: your final thought to conclude.
final_answer: your conclusion.

About you:
{base_system_prompt}

Again, only return a single JSON!
"""

    tool_names = ", ".join([tool.name for tool in tools])
    base_tool_str = "\n".join(e.prototype() for e in tools)

    return template_prompt.format(
        tools=base_tool_str,
        toolnames=tool_names,
        base_system_prompt=log._system_prompt_,
    )


def get_system_prompt(log: ReasoningLog):
    return log.agent_meta_data.persona


def get_system_prompt_with_random_example_tweets(log: ReasoningLog):
    system_prompt = get_system_prompt(log)
    example_tweets = get_random_example_tweets(log.meta_data.knowledge_base_id)

    if len(example_tweets) > 0:
        tweets_str = "\n".join([f"- {x}" for x in example_tweets])
        system_prompt += f"\n\nHere are example tweets written in the defined style:\n{tweets_str}"

    return system_prompt


def create_twitter_auth_from_reasoning_log(log: ReasoningLog):
    return twitter_v2.TwitterRequestAuthorization(
        agent_contract_id=log.meta_data.agent_contract_id,
        twitter_id=log.meta_data.twitter_id,
        twitter_username=log.meta_data.twitter_username,
        chain_id=int(log.meta_data.chain_id),
        knowledge_id=log.meta_data.knowledge_base_id,
        ref_id=log.meta_data.ref_id,
        request_id=log.id,
        task=log.task,
        toolset=log.toolset,
        kn_base=create_kn_base(log),
        prompt=log.prompt,
        model_name=log.model,
    )


def magic_toolset_from_reasoning_log(log: ReasoningLog, llm: OpenAILLMBase):
    if len(log.tool_list) > 0:
        toolcall = dynamic_toolcall.DynamicToolcall(log.tool_list)

    else:
        toolcall = wrapped_external_apis.LiveXDB(
            auth=create_twitter_auth_from_reasoning_log(log),
            agent_config=log.agent_meta_data,
            llm=llm,
        )

    return toolcall


def model_dependent_frequency_penalty(task: AutoAgentTask) -> float:
    frequency_penalty = 0

    if task.model == ModelName.INTELLECT_10B:
        frequency_penalty = 0.1

    return frequency_penalty


def create_llm(task: AutoAgentTask):
    if const.LLM_MODE == "0":
        return ASyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=const.DEFAULT_TEMPERATURE,
            base_url=const.BACKEND_API,
            api_key=const.BACKEND_AUTH_TOKEN,
            frequency_penalty=model_dependent_frequency_penalty(task),
            model=task.model,
            seed=random.randint(1, int(1e9)),
            chain_id=task.meta_data.chain_id,
            agent_contract_id=task.meta_data.agent_contract_id,
            metadata=task.meta_data.model_dump(),
        )

    else:
        models_info = const.SELF_HOSTED_MODELS
        models = list(filter(lambda x: x["model"] == task.model, models_info))

        if len(models) == 0:
            raise Exception(
                f"Model {task.model} not found in the list of SELF_HOSTED_MODELS"
            )

        model = models[0]

        return SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=const.DEFAULT_TEMPERATURE,
            base_url=model["url"].rstrip("/") + "/v1",
            api_key=model["api_key"],
            frequency_penalty=model_dependent_frequency_penalty(task),
            model=task.model,
            seed=random.randint(1, int(1e9)),
            metadata=task.meta_data.model_dump(),
        )


def create_kn_base(task: AutoAgentTask) -> KnowledgeBase:
    if len(task.agent_meta_data.kb_agents) == 0:
        kb_ids = parse_knowledge_ids(task.meta_data.knowledge_base_id)
        return KBStore(
            default_top_k=5,
            similarity_threshold=0.5,
            base_url=const.RAG_API,
            api_key=const.RAG_SECRET_TOKEN,
            kbs=[AgentKnowledgeBase(chain_id="", kb_id=id) for id in kb_ids],
        )
    if const.KN_BASE_MODE == "0":
        return EternalKnowledgeBase(
            default_top_k=5,
            similarity_threshold=0.5,
            base_url=const.BACKEND_API,
            api_key=const.BACKEND_AUTH_TOKEN,
            kbs=task.agent_meta_data.kb_agents,
        )
    else:
        return KBStore(
            default_top_k=5,
            similarity_threshold=0.5,
            base_url=const.RAG_API,
            api_key=const.RAG_SECRET_TOKEN,
            kbs=task.agent_meta_data.kb_agents,
        )
