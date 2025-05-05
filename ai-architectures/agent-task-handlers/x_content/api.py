import json
from fastapi import APIRouter, Depends, BackgroundTasks
from fastapi.responses import JSONResponse

from x_content.cache.chat_request_state_handler import ChatRequestStateHandler
from .verifications import (
    verify_opencall_x_token,
    verify_third_party_authorization_key,
    verify_x_token,
)
from .wrappers.api.twitter_v2.models.response import Response, SearchTweetDto
from x_content import constants as const
from .models import (
    ChatRequest,
    ReasoningLog,
    APIResponse,
    APIStatus,
    TwinTaskSubmitResponse,
    TwinTaskSubmitRequest,
    MissionChainState,
)

from x_content.wrappers.api import twitter_v2
from x_content.tasks.social_agent import post_v3
from x_content.wrappers import bing_search
from x_content.wrappers.magic import sync2async
from .service import MissionStateHandler, handle_chat_request, service_v2_handle_request
from .wrappers import redis_wrapper
from redis import asyncio as aredis
from .legacy_services.twin import twin_service
from .tasks import utils as task_utils
import time
from functools import lru_cache
import logging
from x_content.wrappers.log_decorators import compress_kwargs

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


@lru_cache(maxsize=1)
def get_state_handler():
    state_handler = MissionStateHandler()
    return state_handler


@lru_cache(maxsize=1)
def get_chat_request_state_handler():
    state_handler = ChatRequestStateHandler()
    return state_handler


router = APIRouter()


@router.get(
    "/api/twitter-news",
    dependencies=[
        Depends(verify_opencall_x_token),
        Depends(verify_third_party_authorization_key(const.BACKEND_API, {})),
    ],
)
async def get_twitter_news(query: str) -> JSONResponse:

    result: Response[SearchTweetDto] = await sync2async(
        twitter_v2.search_recent_tweets
    )(query, limit_observation=10)

    if result.is_error():
        return APIResponse(
            status=APIStatus.ERROR,
            error=result.error,
        )

    tweets = result.data.tweets or []

    return APIResponse(
        status=APIStatus.SUCCESS, data=[e.full_text for e in tweets]
    ).model_dump()


@router.get(
    "/api/bing-news/",
    dependencies=[
        Depends(verify_opencall_x_token),
        Depends(verify_third_party_authorization_key(const.BACKEND_API, {})),
    ],
)
async def get_bing_news(query: str) -> JSONResponse:

    result = await sync2async(bing_search.search_from_bing)(
        query, top_k=10, task_name="api:bing-news"
    )

    return APIResponse(
        status=APIStatus.SUCCESS,
        data=result,
    ).model_dump()


@router.get(
    "/api/post-v3-sample-content/",
    dependencies=[
        Depends(verify_opencall_x_token),
        Depends(verify_third_party_authorization_key(const.BACKEND_API, {})),
    ],
)
async def get_post_v3_sample_content(
    twitter_username, cutoff_hour: int = 2
) -> JSONResponse:

    return APIResponse(
        status=APIStatus.SUCCESS,
        data=await post_v3.PostV3.get_contents(
            twitter_username=twitter_username,
            time_cutoff_hours=cutoff_hour,
            top_k=5,
        ),
    ).model_dump()


@router.get("/debug/redis", dependencies=[Depends(verify_x_token)])
async def get_redis_by_key(key: str):

    async with aredis.Redis(
        connection_pool=redis_wrapper.get_aio_redis_connection_pool()
    ) as client:
        res = await client.get(key)

    if res is not None:
        return res
    else:
        return ""


@router.post(
    "/v1/twin/submit",
    response_model=TwinTaskSubmitResponse,
    dependencies=[Depends(verify_x_token)],
)
async def twin_task_submit(
    request: TwinTaskSubmitRequest, background_tasks: BackgroundTasks
) -> TwinTaskSubmitResponse:
    task_id = f"task_{int(time.time())}"

    background_tasks.add_task(
        logger.info,
        f"[twin_task_submit] Received request: task_id {task_id}, request={compress_kwargs(**request.model_dump())}",
    )

    background_tasks.add_task(
        twin_service.generate_twin, request.agent_id, request.twitter_ids
    )

    return TwinTaskSubmitResponse(status="success", task_id=task_id)


@router.post("/async/enqueue", dependencies=[Depends(verify_x_token)])
async def enqueue_api(
    request: ReasoningLog, background_tasks: BackgroundTasks
) -> ReasoningLog:

    background_tasks.add_task(
        logger.info,
        f"[enqueue_api] Received request: {compress_kwargs(**request.model_dump())}",
    )

    if request.state == MissionChainState.NEW:
        background_tasks.add_task(
            task_utils.notify_status_reasoning_log, request
        )

    background_tasks.add_task(service_v2_handle_request, request)

    await get_state_handler().acommit(request)
    return request


@router.get("/async/get", dependencies=[Depends(verify_x_token)])
async def get_result_api(id: str, thought_only: bool = False) -> JSONResponse:
    log = await get_state_handler().a_get(id, none_if_error=True)

    if log is None:
        return JSONResponse(
            {},
            status_code=404,
        )

    if thought_only:
        state = log.state

        thoughts = [
            e.get("final_answer", e.get("thought"))
            for e in log.scratchpad
            if "final_answer" in e or "thought" in e
        ]

        return JSONResponse(
            {
                "state": state,
                "tweets": thoughts,
                "agent_contract_id": log.meta_data.agent_contract_id,
                "chain_id": log.meta_data.chain_id,
                "ref_id": log.meta_data.ref_id,
            },
            status_code=200,
        )

    return JSONResponse(content=log.model_dump(), status_code=200)


@router.post("/async/chat/enqueue", dependencies=[Depends(verify_x_token)])
async def enqueue_chat(
    request: ChatRequest, background_tasks: BackgroundTasks
) -> ChatRequest:
    # logger.info(
    #     f"[enqueue_chat] Received request: {json.dumps(request.model_dump())}"
    # )
    if request.state == MissionChainState.NEW:
        await task_utils.notify_status_chat_request(request)

    await get_chat_request_state_handler().acommit(request)
    background_tasks.add_task(handle_chat_request, request)
    return request


@router.get("/async/chat/get", dependencies=[Depends(verify_x_token)])
async def get_chat_result_api(id: str) -> JSONResponse:
    request = await get_chat_request_state_handler().a_get(
        id, none_if_error=True
    )

    if request is None:
        return JSONResponse(
            {},
            status_code=404,
        )

    return JSONResponse(content=request.model_dump(), status_code=200)


@router.get(
    f"/async/internal-{const.API_SECRET_TOKEN}/get", include_in_schema=False
)
async def get_result_api_internal(
    id: str, thought_only: bool = False
) -> JSONResponse:
    return await get_result_api(id, thought_only)
