from typing import List, Dict, Any, Union
import logging
from mcp.types import CallToolResult, TextContent, EmbeddedResource, Tool
from mcp import ClientSession
import datetime
import re
from pydantic import BaseModel

logger = logging.getLogger(__name__)

def convert_mcp_tools_to_openai_format(
    mcp_tools: List[Any]
) -> List[Dict[str, Any]]:
    """Convert MCP tool format to OpenAI tool format"""
    openai_tools = []
    
    logger.debug(f"Input mcp_tools type: {type(mcp_tools)}")
    logger.debug(f"Input mcp_tools: {mcp_tools}")
    
    # Extract tools from the response
    if hasattr(mcp_tools, 'tools'):
        tools_list = mcp_tools.tools
        logger.debug("Found ListToolsResult, extracting tools attribute")
    elif isinstance(mcp_tools, dict):
        tools_list = mcp_tools.get('tools', [])
        logger.debug("Found dict, extracting 'tools' key")
    else:
        tools_list = mcp_tools
        logger.debug("Using mcp_tools directly as list")
        
    logger.debug(f"Tools list type: {type(tools_list)}")
    logger.debug(f"Tools list: {tools_list}")
    
    # Process each tool in the list
    if isinstance(tools_list, list):
        logger.debug(f"Processing {len(tools_list)} tools")
        for tool in tools_list:
            logger.debug(f"Processing tool: {tool}, type: {type(tool)}")
            if hasattr(tool, 'name') and hasattr(tool, 'description'):
                openai_name = sanitize_tool_name(tool.name)
                logger.debug(f"Tool has required attributes. Name: {tool.name}")
                
                tool_schema = getattr(tool, 'inputSchema', {})
                (tool_schema.setdefault(k, v) for k, v in {
                    "type": "object",
                    "properties": {},
                    "required": []
                }.items()) 
                                
                openai_tool = {
                    "type": "function",
                    "function": {
                        "name": openai_name,
                        "description": tool.description,
                        "parameters": tool_schema
                    }
                }

                openai_tools.append(openai_tool)
                logger.debug(f"Converted tool {tool.name} to OpenAI format")
            else:
                logger.debug(
                    f"Tool missing required attributes: "
                    f"has name = {hasattr(tool, 'name')}, "
                    f"has description = {hasattr(tool, 'description')}"
                )
    else:
        logger.debug(f"Tools list is not a list, it's a {type(tools_list)}")
    
    return openai_tools

def sanitize_tool_name(name: str) -> str:
    """Sanitize tool name for OpenAI compatibility"""
    # Replace any characters that might cause issues
    return name.replace("-", "_").replace(" ", "_").lower()

def compare_toolname(openai_toolname: str, mcp_toolname: str) -> bool:
    return sanitize_tool_name(mcp_toolname) == openai_toolname


def strip_toolcall_noti(content: str) -> str:
    cleaned = re.sub(r"<details\b[^>]*>.*?</details>", "", content, flags=re.DOTALL | re.IGNORECASE)
    return cleaned.strip()

def strip_thinking_content(content: str) -> str:
    pat = re.compile(r"<think>.*?</think>", re.DOTALL | re.IGNORECASE)
    return pat.sub("", content).lstrip()


async def execute_openai_compatible_toolcall_via_opened_session(
    toolname: str, arguments: Dict[str, Any], mcp_tools: List[Tool], session: ClientSession
) -> CallToolResult:

    actual = [
        tool.name
        for tool in mcp_tools
        if compare_toolname(toolname, tool.name)
    ]

    if len(actual) > 1:
        logger.warning(
            "More than one tool has the same santizied"
            " name to the requested tool"
        )

    elif len(actual) == 0:
        return CallToolResult(
            content=[TextContent(text=f"Tool {toolname} not found")], 
            isError=True
        )

    toolname = actual[0]
    res = await session.call_tool(toolname, arguments)

    return res

import fastmcp

async def execute_openai_compatible_toolcall(
    toolname: str, arguments: Dict[str, Any], mcp: fastmcp.FastMCP
) -> list[Union[TextContent, EmbeddedResource]]:
    tools = await mcp._mcp_list_tools()
    candidate: List[Tool] = []

    for tool in tools:
        tool: Tool
        if compare_toolname(toolname, tool.name):
            candidate.append(tool)

    if len(candidate) > 1:
        logger.warning(
            "More than one tool has the same santizied"
            " name to the requested tool"
        )
        
    elif len(candidate) == 0:
        return CallToolResult(
            content=[TextContent(text=f"Tool {toolname} not found")], 
            isError=True
        )
        
    toolname = candidate[0].name
    res = await mcp._mcp_call_tool(toolname, arguments)

    return res



async def refine_chat_history(messages: list[dict[str, str]], system_prompt: str) -> list[dict[str, str]]:
    refined_messages = []
    
    current_time_utc_str = datetime.datetime.now(tz=datetime.timezone.utc).strftime("%Y-%m-%d %H:%M:%S")
    system_prompt += f'\nNote: Current time is {current_time_utc_str} UTC (only use this information when being asked or for searching purposes)'

    has_system_prompt = False
    for message in messages:
        message: dict[str, str]

        if isinstance(message, dict) and message.get('role', 'undefined') == 'system':
            message['content'] += f'\n{system_prompt}'
            has_system_prompt = True
            continue

        
        _message = {
            "role": message.get('role', 'assistant'),
            "content": strip_toolcall_noti(strip_thinking_content(message.get('content', '')))
        }

        refined_messages.append(_message)

    if not has_system_prompt and system_prompt != "":
        refined_messages.insert(0, {
            "role": "system",
            "content": system_prompt
        })

    if isinstance(refined_messages[-1], str):
        refined_messages[-1] = {
            "role": "user",
            "content": refined_messages[-1]
        }

    # current_time_utc_str = datetime.datetime.now(tz=datetime.timezone.utc).strftime("%Y-%m-%d %H:%M:%S")
    # refined_messages[-1]['content'] += f'\nCurrent time is {current_time_utc_str} UTC'

    return refined_messages


async def refine_assistant_message(
    assistant_message: dict[str, str]
) -> dict[str, str]:

    if 'content' in assistant_message:
        assistant_message['content'] = strip_thinking_content(assistant_message['content'] or "")

    return assistant_message


async def refine_mcp_response(something: Any) -> str:
    if isinstance(something, dict):
        return {
            k: await refine_mcp_response(v)
            for k, v in something.items()
        }

    elif isinstance(something, (list, tuple)):
        return [
            await refine_mcp_response(v)
            for v in something
        ]

    elif isinstance(something, BaseModel):
        return something.model_dump()

    return something