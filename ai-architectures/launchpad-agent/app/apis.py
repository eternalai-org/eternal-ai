from fastapi import APIRouter, BackgroundTasks
import logging
from .schemas.services import EvaluationRequest
from .schemas.commons import ResponseMessage
from .handlers import evaluate_tweet

logger = logging.getLogger(__name__)
router = APIRouter(prefix="/api")

@router.post("/evaluate")
async def evaluate(request: EvaluationRequest, background_tasks: BackgroundTasks) -> ResponseMessage[str]:
    background_tasks.add_task(evaluate_tweet, request)
    return ResponseMessage[str](
        status="success",
        result=f"Evaluation request received (tweet-id: {request.tweet_id}, launchpad-id: {request.launchpad_id})",
    )
