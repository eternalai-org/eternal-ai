from fastapi import APIRouter, BackgroundTasks, HTTPException
from typing import List, Optional
import logging
from .schemas.services import EvaluationRequest
from .schemas.commons import ResponseMessage
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
        status="success",
        result=f"Evaluation request received (tweet-id: {request.tweet_id})",
    )

@router.get("/evaluations/{tweet_id}")
async def get_evaluation(tweet_id: str) -> ResponseMessage[EvaluationResult]:
    """Get evaluation results for a specific tweet"""
    try:
        db = get_mongo_database("evaluations")
        collection = db.get_collection("tweet_evaluations")
        
        result = collection.find_one({"_id": tweet_id})
        
        if not result:
            raise HTTPException(status_code=404, detail="Evaluation not found")
        
        # Remove MongoDB _id field and convert back to model
        result.pop("_id", None)
        evaluation = EvaluationResult.model_validate(result)
        
        return ResponseMessage[EvaluationResult](
            status="success",
            result=evaluation
        )
        
    except HTTPException:
        raise
    except Exception as e:
        logger.error(f"Error retrieving evaluation for {tweet_id}: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")

@router.get("/projects/{launchpad_id}/investors")
async def get_candidate_investors(
    launchpad_id: str, 
    min_grade: Optional[str] = None,
    limit: int = 100
) -> ResponseMessage[List[InvestorProfile]]:
    """Get candidate investors for a launchpad project"""
    try:
        db = get_mongo_database("evaluations")
        collection = db.get_collection("tweet_evaluations")
        
        # Build query
        query = {
            "project_identification.launchpad_id": launchpad_id,
            "investor_profile": {"$exists": True, "$ne": None},
            "status": "completed_full"
        }
        
        # Add grade filter if specified
        if min_grade:
            grade_order = {"A": 5, "B": 4, "C": 3, "D": 2, "E": 1}
            min_grade_value = grade_order.get(min_grade.upper())
            if min_grade_value:
                allowed_grades = [grade for grade, value in grade_order.items() if value >= min_grade_value]
                query["investor_profile.grade"] = {"$in": allowed_grades}
        
        # Execute query
        results = list(collection.find(query).limit(limit))
        
        # Extract investor profiles
        investors = []
        seen_users = set()
        
        for result in results:
            profile_data = result.get("investor_profile")
            if profile_data and profile_data.get("user_id") not in seen_users:
                investor = InvestorProfile.model_validate(profile_data)
                investors.append(investor)
                seen_users.add(investor.user_id)
        
        # Sort by score descending
        investors.sort(key=lambda x: x.score, reverse=True)
        
        return ResponseMessage[List[InvestorProfile]](
            status="success",
            result=investors
        )
        
    except Exception as e:
        logger.error(f"Error retrieving investors for project {launchpad_id}: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")

@router.get("/analytics/projects/{launchpad_id}")
async def get_project_analytics(launchpad_id: str) -> ResponseMessage[dict]:
    """Get analytics for investor interest in a project"""
    try:
        db = get_mongo_database("evaluations")
        collection = db.get_collection("tweet_evaluations")
        
        # Get all evaluations for this project
        evaluations = list(collection.find({
            "project_identification.launchpad_id": launchpad_id
        }))
        
        if not evaluations:
            return ResponseMessage[dict](
                status="success",
                result={
                    "project_id": launchpad_id,
                    "total_mentions": 0,
                    "candidate_tweets": 0,
                    "qualified_investors": 0,
                    "grade_distribution": {},
                    "average_score": 0.0,
                    "top_research_interests": []
                }
            )
        
        # Calculate analytics
        total_mentions = len(evaluations)
        candidate_tweets = len([e for e in evaluations if e.get("tweet_evaluation", {}).get("classification") == "candidate"])
        qualified_investors = len([e for e in evaluations if e.get("investor_profile") is not None])
        
        # Grade distribution
        grade_distribution = {}
        scores = []
        research_interests = {}
        
        for eval_result in evaluations:
            investor = eval_result.get("investor_profile")
            if investor:
                grade = investor.get("grade")
                if grade:
                    grade_distribution[grade] = grade_distribution.get(grade, 0) + 1
                
                score = investor.get("score", 0)
                if score > 0:
                    scores.append(score)
                
                # Collect research interests
                for interest in investor.get("research_interests", []):
                    category = interest.get("category")
                    if category:
                        if category not in research_interests:
                            research_interests[category] = []
                        research_interests[category].append(interest.get("confidence", 0))
        
        # Top research interests
        top_interests = []
        for category, confidences in research_interests.items():
            avg_confidence = sum(confidences) / len(confidences)
            top_interests.append({
                "category": category,
                "average_confidence": avg_confidence,
                "mention_count": len(confidences)
            })
        
        top_interests.sort(key=lambda x: x["average_confidence"], reverse=True)
        
        analytics = {
            "project_id": launchpad_id,
            "total_mentions": total_mentions,
            "candidate_tweets": candidate_tweets,
            "qualified_investors": qualified_investors,
            "grade_distribution": grade_distribution,
            "average_score": sum(scores) / len(scores) if scores else 0.0,
            "top_research_interests": top_interests[:10]
        }
        
        return ResponseMessage[dict](
            status="success",
            result=analytics
        )
        
    except Exception as e:
        logger.error(f"Error generating analytics for project {launchpad_id}: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")

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
