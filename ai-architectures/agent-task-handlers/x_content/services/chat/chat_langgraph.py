from datetime import datetime
import json
from typing import Annotated, TypedDict
from json_repair import repair_json
from langgraph.graph import StateGraph, START, END
from langgraph.graph.state import CompiledStateGraph
from langgraph.graph.message import add_messages
from langchain.schema import ChatMessage, AIMessage, SystemMessage, HumanMessage
from x_content.llm.base import OpenAILLMBase
from langgraph.store.memory import InMemoryStore, BaseStore
from langchain_core.runnables import RunnableConfig
from x_content.services.chat import tools
from x_content.services.chat.memory import get_formatted_memories
from langgraph.prebuilt import ToolNode, tools_condition
from langchain_core.messages import AnyMessage

SYS_PROMPT_TEMPLATE = """{persona}
{tools}
{user_info}
System Time: {time}
"""

in_memory_store = InMemoryStore()


class State(TypedDict):
    # Messages have the type "list". The `add_messages` function
    # in the annotation defines how this state key should be updated
    # (in this case, it appends messages to the list, rather than overwriting them)
    messages: Annotated[list[AnyMessage], add_messages]
    tx_hash: str = ""


def get_chat_langgraph(llm: OpenAILLMBase) -> CompiledStateGraph:
    tool_list = [tools.upsert_memory]

    async def chatbot(
        state: State, config: RunnableConfig, *, store: BaseStore
    ):
        persona = config["configurable"]["persona"]
        user_id = config["configurable"]["user_id"]

        messages = state["messages"]
        formatted_memories = await get_formatted_memories(
            store, user_id, messages
        )
        formatted_tools = tools.get_formatted_tools(tool_list)

        sys = SYS_PROMPT_TEMPLATE.format(
            persona=persona,
            tools=formatted_tools,
            user_info=formatted_memories,
            time=datetime.now().isoformat(),
        )

        messages = [SystemMessage(content=sys), *messages]

        infer_result = await llm.agenerate(messages)
        content = infer_result.generations[0].message.content
        messages.append(ChatMessage(role="assistant", content=content))

        tool_calls = []
        try:
            data = repair_json(content, return_objects=True)
            if "name" in data and "parameters" in data:
                tool_calls = [
                    {
                        "name": data["name"],
                        "args": data["parameters"],
                        "id": "tool_call_id",
                        "type": "tool_call",
                    }
                ]
        except Exception as err:
            pass

        messages.append(AIMessage(content=content, tool_calls=tool_calls))

        print(json.dumps([x.model_dump() for x in messages]))

        return {
            "messages": [AIMessage(content=content, tool_calls=tool_calls)],
            "tx_hash": infer_result.tx_hash,
        }

    graph_builder = StateGraph(State)
    graph_builder.add_node("chatbot", chatbot)

    tool_node = ToolNode(tools=tool_list)
    graph_builder.add_node("tools", tool_node)

    graph_builder.add_edge(START, "chatbot")
    graph_builder.add_conditional_edges(
        "chatbot",
        tools_condition,
    )
    graph_builder.add_edge("tools", "chatbot")

    graph = graph_builder.compile(store=in_memory_store)

    return graph
