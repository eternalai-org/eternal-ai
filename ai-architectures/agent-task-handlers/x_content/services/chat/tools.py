"""Define he agent's tools."""

import json
import uuid
from typing import Annotated, Optional

from langchain_core.runnables import RunnableConfig
from langchain_core.tools import InjectedToolArg
from langgraph.store.base import BaseStore
from langchain_core.tools import tool, BaseTool
from langchain_core.utils.function_calling import (
    convert_to_openai_tool,
)
from langgraph.prebuilt import InjectedStore

from x_content.services.chat.memory import get_memory_namespace

TOOL_TEMPLATE = """Use the function '{name}' to {description}

{tool_json}"""

TOOL_CALLING_TEMPLATE = """
You have tool calling capabilities. When you receive a tool call response, use the output to format an answer to the orginal use question. 
If you are using tools, respond in the format {{"name": function name, "parameters": dictionary of function arguments}}. Do not use variables.

You have access to the following functions:

{formatted_tools}
"""


def _lower_first_letter(text: str) -> str:
    return text[:1].lower() + text[1:]


def _format_tool(tool: dict) -> str:
    return TOOL_TEMPLATE.format(
        name=tool["function"]["name"],
        description=_lower_first_letter(tool["function"]["description"]),
        tool_json=json.dumps(tool, indent=2),
    )


def get_formatted_tools(tools: list[BaseTool]) -> str:
    formatted_tools = [convert_to_openai_tool(x) for x in tools]
    formatted_tools = "\n\n".join([_format_tool(x) for x in formatted_tools])
    if formatted_tools:
        formatted_tools = TOOL_CALLING_TEMPLATE.format(
            formatted_tools=formatted_tools
        )
    return formatted_tools


async def upsert_memory(
    content: str,
    context: str,
    *,
    memory_id: Optional[uuid.UUID] = None,
    # Hide these arguments from the model.
    config: Annotated[RunnableConfig, InjectedToolArg],
    store: Annotated[BaseStore, InjectedStore()],
):
    """Upsert a memory in the database.

    If a memory conflicts with an existing one, then just UPDATE the
    existing one by passing in memory_id - don't create two memories
    that are the same. If the user corrects a memory, UPDATE it.

    Args:
        content: The main content of the memory. For example:
            "User expressed interest in learning about French."
        context: Additional context for the memory. For example:
            "This was mentioned while discussing career options in Europe."
        memory_id: ONLY PROVIDE IF UPDATING AN EXISTING MEMORY.
        The memory to overwrite.
    """
    mem_id = memory_id or uuid.uuid4()
    user_id = config["configurable"]["user_id"]
    await store.aput(
        get_memory_namespace(user_id),
        key=str(mem_id),
        value={"content": content, "context": context},
    )
    return f"Stored memory {mem_id}"
