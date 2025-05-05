import traceback
import logging
from x_content.models import (
    ToolParamDtype,
    ToolDef,
    AdvanceToolDef,
    ToolParam,
)
from typing import List, Any, Union
import httpx
import json
from x_content import wrappers
from x_content.wrappers.magic import sync2async
from urllib.parse import urlparse

logger = logging.getLogger(__name__)
LIMIT_TOTAL_TOKENS_PER_OBSERVATION = 700 * 3


def is_float(xx: Any):
    try:
        float(xx)
        return True
    except ValueError:
        return False


def split_params(spec: dict, action_input):
    resp = {}

    for i, (k, v) in enumerate(spec.items()):
        if v["type"] == "number":
            resp[k] = int(action_input[i])
        else:  # v['type'] == "string":
            resp[k] = str(action_input[i])

    return resp


def make_response(content, success=True):
    return {
        "success": success,
        "content": content,
    }

def shorten_data(data: Union[dict, list, str], max_size_list: int = 10, max_size_str: int = 2048) -> dict | list:
    if isinstance(data, dict):
        data = {
            k: shorten_data(v, max_size_list, max_size_str)
            for k, v in data.items()
        }

    elif isinstance(data, list):
        data = data[:max_size_list]
        data = [
            shorten_data(item, max_size_list, max_size_str)
            for item in data
        ]
        
    elif isinstance(data, str):
        data = data[:max_size_str]

    return data

async def execute_http_toolcall(
    method: str, 
    endpoint: str,
    params: dict, 
    headers: dict=None
):
    parted_url = urlparse(endpoint)
    url = '{}://{}{}'.format(parted_url.scheme, parted_url.netloc, parted_url.path)

    payload = dict(
        method=method,
        url=url,
        params={}        
    )
    
    splited_query = [
        e.split('=') 
        for e in parted_url.query.split('&')
    ]

    splited_query = [
        e
        for e in splited_query 
        if len(e) == 2
    ]

    for k, v in splited_query:
        payload["params"][k] = v

    if headers is not None:
        payload["headers"] = headers

    if method == "GET":
        payload["params"].update(params)
    else:
        payload["json"] = params

    logger.info(f"Calling API with payload: {json.dumps(payload, indent=2)}")

    async with httpx.AsyncClient() as client:
        resp = await client.request(**payload)

    if resp.status_code != 200:
        pre_formated = (
            "Method: {}\n"
            "Endpoint: {}\n"
            "Params: \n<pre>{}</pre>\n"
            "Headers: \n<pre>{}</pre>\n"
            "Response: \n<pre>{}</pre>\n"
            "Code: {}\n"
        ).format(
            method,
            endpoint,
            json.dumps(params, indent=2),
            (
                json.dumps({k: "***" for k, _ in headers.items()}, indent=2)
                if isinstance(headers, dict)
                else "n/a"
            ),
            resp.text,
            resp.status_code,
        )

        task_id = params.get("request_id", None)

        await wrappers.telegram.a_send_message(
            "junk_nofitications",
            f"Request {task_id} (dynamic toolcall): Failed to call API\n\n{pre_formated}",
            room=wrappers.telegram.TELEGRAM_ALERT_ROOM,
            fmt="HTML",
        )

        return make_response(
            f"Error: {resp.text} (Code: {resp.status_code})", False
        )

    try:
        resp_to_agent = resp.json()
    except json.JSONDecodeError:
        resp_to_agent = resp.text

    return make_response(shorten_data(resp_to_agent), True)


def value_cast(dtype: ToolParamDtype, value: str):
    if dtype == ToolParamDtype.STRING:
        return value

    if dtype == ToolParamDtype.NUMBER:
        return float(value)

    if dtype == ToolParamDtype.BOOLEAN:
        return value.lower() in ["true", "1"]

    return value


def parse_toolcall_params(params: List[ToolParam], inputs):
    no_default_params = [p for p in params if p.default_value is None]
    res = {}

    for p in params:
        if p.default_value is not None:
            res[p.name] = p.default_value

    res = {
        **res,
        **{
            k.name: value_cast(k.dtype, v)
            for k, v in zip(no_default_params, inputs)
        },
    }

    return res


async def execute_advance_tool(
    tool: AdvanceToolDef, action_input: str, request_id: str = None
):
    inputs = [e.strip() for e in action_input.split("|")]
    no_default_params = [p for p in tool.params if p.default_value is None]

    if len(no_default_params) == 0:
        return [
            await execute_http_toolcall(
                tool.method,
                tool.executor,
                parse_toolcall_params(tool.params, []),
                tool.headers,
            )
        ]

    try:
        n_params = len(no_default_params)
        res = []

        for i in range(0, len(inputs), n_params):
            l, r = i, i + n_params

            if r > len(inputs):
                break

            res.append(
                await execute_http_toolcall(
                    tool.method,
                    tool.executor,
                    parse_toolcall_params(tool.params, inputs[l:r]),
                    tool.headers,
                )
            )

            if not tool.allow_multiple:
                break

        if len(res) == 0:
            if n_params > 0:
                return [
                    "Something went wrong while executing the tool and no results were returned."
                ]

            missing_params = [p.name for p in tool.params[-len(inputs) :]]
            return [
                f"{tool.name} requires {len(tool.params)} parameters, but only {len(inputs)} provided. Missing {len(missing_params)} value(s) for: {missing_params}"
            ]

        return res

    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"Error while executing {tool.name} with input {action_input}: {err}"
        )
        return [f"Error while executing {tool.name} with input {action_input}"]


async def execute_tool(
    tool: Union[ToolDef, AdvanceToolDef],
    action_input: str,
    request_id: str = None,
):
    if isinstance(tool, AdvanceToolDef):
        return await execute_advance_tool(
            tool, action_input, request_id
        )

    inputs = action_input.split("|")
    async_fn = sync2async(tool.executor)

    if len(tool.params) == 0:
            return [await async_fn()]

    try:
        n_params = len(tool.params)
        res = []

        for i in range(0, len(inputs), n_params):
            if i + n_params > len(inputs):
                break

            l, r = i, i + n_params
            result = await async_fn(*inputs[l:r])
            res.append(result)

            if not tool.allow_multiple:
                break

        if len(res) == 0:
            missing_params = [p.name for p in tool.params[-len(inputs) :]]
            return [
                f"{tool.name} requires {len(tool.params)} parameters, but only {len(inputs)} provided. Missing {len(missing_params)} value(s) for: {missing_params}"
            ]

        return res

    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"Error while executing {tool.name} with input {action_input}: {err}"
        )
        return [f"Error while executing {tool.name} with input {action_input}"]
