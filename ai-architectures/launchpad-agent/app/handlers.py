from .schemas.services import EvaluationRequest
import logging 
from .agents import (
    receptionist
)

logger = logging.getLogger(__name__)

async def pipeline():
    pass

async def evaluate_tweet(request: EvaluationRequest):
    logger.info(f"Evaluating tweet {request.tweet_id} for launchpad {request.launchpad_id}")