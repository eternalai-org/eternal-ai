from .react_agent import ReactAgent
import re
from .react_agent import dynamic_system_reminder
from x_content.models import ReasoningLog, ToolDef, MissionChainState
from typing import List
from x_content.tasks.utils import a_move_state
from x_content.toolcall.utils import execute_tool
from x_content.llm.base import OnchainInferResult
from x_content import constants as const
import logging

logger = logging.getLogger(__name__)


def format_prompt_v2(log: ReasoningLog, tools: List[ToolDef]):
    template_prompt = """
You have access to the following tools to get information or take actions:
{tools}

Your knowledge is outdated for years, it is recommended to use the provided tools to get the newest information before acting something.

Your answer must be formed in XML format with action and action input as follows:
- <action>must be one of {toolnames}</action>
- <action_input>the necessary parameters for the chosen action separated by |. For instance, if the action requires two parameters are 'a' and 'b', the action_input should be 'a|b'</action_input>

OR with exactly one final_answer:
- <final_answer>your final_answer</final_answer>

About you:
{base_system_prompt}

Again, shortly think about what to do and answer in the required format!
"""

    tool_names = ", ".join([tool.name for tool in tools])
    base_tool_str = "\n".join(e.prototype() for e in tools)

    return template_prompt.format(
        tools=base_tool_str,
        toolnames=tool_names,
        base_system_prompt=log._system_prompt_,
    )


def parse_conversational_react_response(response: str, verbose=True) -> dict:
    thought_pat = re.compile(
        r"[<think>]?(.*?)</think>", re.IGNORECASE | re.DOTALL
    )
    action_pat = re.compile(
        r"<action>(.*?)</action>", re.IGNORECASE | re.DOTALL
    )
    action_input_pat = re.compile(
        r"<action_input>(.*?)</action_input>", re.IGNORECASE | re.DOTALL
    )
    final_answer_pat = re.compile(
        r"<final_answer>(.*?)</final_answer>", re.IGNORECASE | re.DOTALL
    )

    action = action_pat.findall(response)
    action_input = action_input_pat.findall(response)
    thought = thought_pat.findall(response)
    final_answer = final_answer_pat.findall(response)
    segment_pad = {}

    if action is not None and len(action) > 0:
        segment_pad.update({"action": action[0]})

    if action_input is not None and len(action_input) > 0:
        segment_pad.update({"action_input": action_input[0]})

    if thought is not None and len(thought) > 0:
        segment_pad.update({"thought": thought[0]})

    if final_answer is not None and len(final_answer) > 0:
        segment_pad.update({"final_answer": final_answer[0]})

    return segment_pad


def render_conversation(log: ReasoningLog, tools: List[ToolDef]):
    conversation = [
        {"role": "system", "content": format_prompt_v2(log, tools)}
    ]

    for it, item in enumerate(log.scratchpad):
        user_message = {}
        for k in ["task", "observation", "hot_news"]:
            if k in item:
                user_message[k] = item[k]

        user_message["system_reminder"] = dynamic_system_reminder(
            log, it, tools
        )

        assistant_message = {}
        for k in ["thought", "action", "action_input", "final_answer"]:
            if k in item:
                assistant_message[k] = item[k]

        if len(assistant_message) > 0:
            assistant_resp = "<think>\n{}\n</think>".format(
                assistant_message.get("thought", "")
            )

            for item in ["action", "action_input", "final_answer"]:
                if item in assistant_message:
                    assistant_resp += "\n<{}>{}</{}>".format(
                        item, assistant_message[item], item
                    )

            conversation.append(
                {"role": "assistant", "content": assistant_resp}
            )

        response = ""

        for item in ["task", "observation", "hot_news", "system_reminder"]:
            if item in user_message:
                response += "\n<{}>{}</{}>".format(
                    item, user_message[item], item
                )

        conversation.append({"role": "user", "content": response})

    return conversation


class ReactAgentUsingDeepSeekR1(ReactAgent):
    resumable = True

    async def is_react_complete(self, log: ReasoningLog) -> bool:
        return (
            len(log.scratchpad) > 0 and "final_answer" in log.scratchpad[-1]
        ) or len(log.scratchpad) > log.meta_data.params.get(
            "react_max_steps", const.DEFAULT_REACT_MAX_STEPS
        ) + 1

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        tools = self.toolcall.get_tools(log.toolset)

        if log.state == MissionChainState.NEW:
            log.scratchpad = [
                {
                    "task": log.prompt.replace("\n", " ").strip(),
                }
            ]

            log.execute_info = {
                "tool_call_metadata": [],
                "conversation": [],
            }

            return await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )

        if log.state == MissionChainState.RUNNING:
            while not await self.is_react_complete(log):
                conversation = render_conversation(log, tools)
                log.execute_info["conversation"].append(conversation)
                infer_result: OnchainInferResult = await self.llm.agenerate(
                    conversation, temperature=0.7
                )

                result = infer_result.generations[0].message.content
                pad: dict = parse_conversational_react_response(result)

                if len(pad) == 0:
                    return await a_move_state(
                        log,
                        MissionChainState.ERROR,
                        "No response (or wrong response format) from the agent message; Last: {}; receipt: {}; tx-hash: {}".format(
                            result, infer_result.receipt, infer_result.tx_hash
                        ),
                    )

                log = await self.update_react_scratchpad(log, pad)

                if log.state in [
                    MissionChainState.DONE,
                    MissionChainState.ERROR,
                ]:
                    break

                log.scratchpad[-1].update(
                    {
                        "tx_hash": infer_result.tx_hash,
                    }
                )

                log = await self.commit_log(log)

            else:
                if log.state == MissionChainState.RUNNING:
                    log = await a_move_state(
                        log, MissionChainState.DONE, "React task break"
                    )

        return log

    async def update_react_scratchpad(self, log: ReasoningLog, pad: dict):
        tools = self.toolcall.get_tools(log.toolset)

        if "thought" in pad:
            if "thought" in log.scratchpad[-1] and any(
                k not in log.scratchpad[-1]
                for k in ["action", "action_input", "observation"]
            ):
                for kk in ["action", "action_input", "observation"]:
                    if kk not in log.scratchpad[-1]:
                        log.scratchpad[-1][kk] = "Not found!"

                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    "Action/Action Input/Observation of the last step not found",
                )

            log.scratchpad.append({"thought": pad["thought"]})

            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🤔 Thought: {pad['thought']}"
            )

        if "action" in pad:
            if "action_input" not in pad:
                log.scratchpad[-1].update(
                    action=pad["action"], 
                    action_input="Not found!",
                    observation="Action input not found. Please provide a suitable action input to execute the selected action or leave it empty if no input is required."
                )

                return log

            if "task" in log.scratchpad[-1]:
                return await a_move_state(
                    log, MissionChainState.ERROR, "No thought found!"
                )

            action = pad["action"]
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🛠️ Action: {action}"
            )

            action_input = pad["action_input"]
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 📥 Action Input: {action_input}"
            )

            log.scratchpad[-1].update(action=action, action_input=action_input)

            observation = "Action not found!"

            # TODO: refactor this loop
            for tool in tools:
                if tool.name == action:

                    observation = await execute_tool(
                        tool, action_input, request_id=log.id
                    )
                    new_observation = []

                    for x in observation:
                        if isinstance(x, tuple):
                            x, metadata = x
                            log.execute_info["tool_call_metadata"].append(
                                metadata
                            )

                        new_observation.append(x)

                    observation = new_observation
                    break

            log.scratchpad[-1]["observation"] = observation
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🔍 Observation: {observation}"
            )

        if "final_answer" in pad:
            if any(
                k in log.scratchpad[-1]
                for k in ["action", "action_input", "observation"]
            ):
                log.scratchpad.append({})

            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🎯 Final Answer: {pad['final_answer']}"
            )
            log.scratchpad[-1].update({"final_answer": pad["final_answer"]})

            log = await a_move_state(
                log, MissionChainState.DONE, "Final answer found"
            )

        if len(log.scratchpad) > log.meta_data.params.get(
            "react_max_steps", const.DEFAULT_REACT_MAX_STEPS
        ):
            log = await a_move_state(
                log, MissionChainState.DONE, "Scratchpad length exceeded"
            )

        return log
