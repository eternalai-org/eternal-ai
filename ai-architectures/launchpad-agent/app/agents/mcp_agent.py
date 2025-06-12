from app.utils.lm import get_oai_async_client, get_model_id
import fastmcp
from typing import List
from ..mcps.utils import (
    convert_mcp_tools_to_openai_format, 
    execute_openai_compatible_toolcall, 
    refine_assistant_message,
    refine_mcp_response
)
from mcp.types import TextContent, EmbeddedResource
import json
import openai

async def mcp_agent_run(
    messages: list[dict[str, str]],
    mcp: fastmcp.FastMCP,
    max_calls: int = 25,
    **kwargs
) -> str:
    client = get_oai_async_client()

    tools = await mcp._mcp_list_tools()
    oai_tools = convert_mcp_tools_to_openai_format(tools)

    completion = await client.chat.completions.create(
        model=get_model_id(),
        messages=messages,
        tools=oai_tools,
        tool_choice="auto"
    )

    messages.append(await refine_assistant_message(completion.choices[0].message.model_dump()))
    
    n_calls = 0
    
    while completion.choices[0].message.tool_calls is not None \
        and len(completion.choices[0].message.tool_calls) > 0:
            
        n_calls += len(completion.choices[0].message.tool_calls)
        
        for call in completion.choices[0].message.tool_calls:
            _id, _name = call.id, call.function.name
            _args = json.loads(call.function.arguments)

            result = await execute_openai_compatible_toolcall(_name, _args, mcp)
            
            result = [
                r for r in result
                if isinstance(r, (TextContent, EmbeddedResource))
            ]

            messages.append(
                {
                    "role": "tool",
                    "tool_call_id": _id,
                    "content": await refine_mcp_response(result)
                }
            )
        
        completion = await client.chat.completions.create(
            model=get_model_id(),
            messages=messages,
            tools=oai_tools if n_calls < max_calls else openai._types.NOT_GIVEN,
            tool_choice="auto" if n_calls < max_calls else openai._types.NOT_GIVEN
        )

        messages.append(await refine_assistant_message(completion.choices[0].message.model_dump()))

    return completion.choices[0].message.content