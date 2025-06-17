from fastapi import APIRouter, BackgroundTasks, HTTPException
from typing import List, Optional
import logging
from .schemas.services import EvaluationRequest
from .schemas.commons import ResponseMessage, APIStatus
from .schemas.evaluation import EvaluationResult, InvestorProfile, InvestorGrade
from .handlers import evaluate_tweet
from .utils.mongodb import get_mongo_database

logger = logging.getLogger(__name__)
router = APIRouter(prefix="/api")

@router.post("/evaluate")
async def evaluate(request: EvaluationRequest, background_tasks: BackgroundTasks) -> ResponseMessage[str]:
    """Submit a tweet for evaluation"""
    background_tasks.add_task(evaluate_tweet, request)
    return ResponseMessage[str](
        status=APIStatus.OK,
        result=f"Evaluation request received (tweet-id: {request.tweet_id})",
    )

@router.get("/health")
async def health_check():
    """Enhanced health check with system status"""
    try:
        # Test database connection
        db = get_mongo_database("evaluations")
        db.list_collection_names()
        
        return {
            "status": "healthy",
            "message": "Launchpad Agent is running",
            "services": {
                "database": "connected",
                "api": "active"
            }
        }
    except Exception as e:
        return {
            "status": "unhealthy", 
            "message": f"Health check failed: {str(e)}",
            "services": {
                "database": "error",
                "api": "active"
            }
        }
