import time
from datetime import datetime, timezone
from .schemas.services import EvaluationRequest
from .schemas.evaluation import EvaluationResult, TweetClassification
from .agents.tweet_classifier import classify_tweet
from .agents.project_identifier import identify_launchpad_project
from .agents.investor_analyzer import analyze_investor_profile
from .utils.mongodb import get_mongo_database
import logging 

logger = logging.getLogger(__name__)

async def pipeline():
    pass

async def evaluate_tweet(request: EvaluationRequest):
    """Complete tweet evaluation pipeline"""
    
    start_time = time.time()
    logger.info(f"Starting evaluation for tweet {request.tweet_id}")
    
    try:
        # Stage 1: Tweet Classification
        logger.info(f"Stage 1: Classifying tweet {request.tweet_id}")
        classification = await classify_tweet(request.tweet_content, request.tweet_id)
        
        if classification.classification != TweetClassification.CANDIDATE:
            logger.info(f"Tweet {request.tweet_id} classified as {classification.classification.value}, stopping evaluation")
            
            # Store result even for non-candidate tweets
            result = EvaluationResult(
                tweet_id=request.tweet_id,
                twitter_id=request.twitter_id,
                tweet_content=request.tweet_content,
                tweet_evaluation=classification,
                project_identification=None,
                investor_profile=None,
                processing_time_seconds=time.time() - start_time,
                status="completed_non_candidate"
            )
            await store_evaluation_result(result)
            return
        
        # Stage 2: Project Identification
        logger.info(f"Stage 2: Identifying project for tweet {request.tweet_id}")
        project_identification = await identify_launchpad_project(request.tweet_content, request.tweet_id, request.network_id)
        
        if not project_identification.launchpad_id:
            logger.info(f"No launchpad project identified for tweet {request.tweet_id}, stopping evaluation")
            
            # Store result for candidate tweets with no project match
            result = EvaluationResult(
                tweet_id=request.tweet_id,
                twitter_id=request.twitter_id,
                tweet_content=request.tweet_content,
                tweet_evaluation=classification,
                project_identification=project_identification,
                investor_profile=None,
                processing_time_seconds=time.time() - start_time,
                status="completed_no_project"
            )
            await store_evaluation_result(result)
            return
        
        # Stage 3: Investor Analysis
        logger.info(f"Stage 3: Analyzing investor {request.twitter_id} for project {project_identification.launchpad_id}")
        investor_profile = await analyze_investor_profile(
            request.twitter_id, 
            request.tweet_content, 
            project_identification.launchpad_id,
            request.network_id
        )
        
        # Complete evaluation
        result = EvaluationResult(
            tweet_id=request.tweet_id,
            twitter_id=request.twitter_id,
            tweet_content=request.tweet_content,
            tweet_evaluation=classification,
            project_identification=project_identification,
            investor_profile=investor_profile,
            processing_time_seconds=time.time() - start_time,
            status="completed_full"
        )
        
        await store_evaluation_result(result)
        
        logger.info(f"Evaluation complete: Tweet {request.tweet_id}, Project {project_identification.launchpad_id}, "
                   f"Grade {investor_profile.grade.value}, Score {investor_profile.score:.1f}, "
                   f"Time: {time.time() - start_time:.2f}s")
        
    except Exception as e:
        logger.error(f"Error evaluating tweet {request.tweet_id}: {e}", exc_info=True)
        
        # Store error result
        result = EvaluationResult(
            tweet_id=request.tweet_id,
            twitter_id=request.twitter_id,
            tweet_content=request.tweet_content,
            tweet_evaluation=None,
            project_identification=None,
            investor_profile=None,
            processing_time_seconds=time.time() - start_time,
            status=f"error: {str(e)}"
        )
        await store_evaluation_result(result)
        raise

async def store_evaluation_result(result: EvaluationResult):
    """Store evaluation results in MongoDB"""
    try:
        db = get_mongo_database("evaluations")
        collection = db.get_collection("tweet_evaluations")
        
        # Convert to dict for storage
        result_dict = result.model_dump(mode="json")
        
        # Ensure we have a unique identifier
        result_dict["_id"] = result.tweet_id
        
        # Store with upsert to handle duplicates
        collection.replace_one(
            {"_id": result.tweet_id},
            result_dict,
            upsert=True
        )
        
        logger.info(f"Stored evaluation result for tweet {result.tweet_id}")
        
    except Exception as e:
        logger.error(f"Error storing evaluation result for {result.tweet_id}: {e}", exc_info=True)