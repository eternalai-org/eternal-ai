import json
from app.utils.launchpad_api_calls import get_launchpad_detail
from app.schemas.evaluation import ProjectIdentification
from app.utils.lm import get_oai_async_client, get_model_id
import logging
import re
from app.utils.misc import retry, float_clamp
from json_repair import repair_json
from lite_logging import async_log
from app.config import settings
import asyncio

logger = logging.getLogger(__name__)

SYSTEM_PROMPT = """
You are a Launchpad assistant. Launchpad is a platform where creators meet investors. Each project by a creator is identified by an ID (usually called launchpad_id)

Your task now is to read a tweet thread and identify which project is mentioned and targeted. You are provided tools to search or get details (of a project, by its launchpad_id)

Your output MUST be a single JSON object with the following structure:

{
  "launchpad_id": "string, the launchpad_id of the project, if not found, return null",
  "confidence": "confidence score, from 0 to 1, 1 means very confident",
  "reasoning": "string, explaining why you think this is the project"
}

In case the thread mentions more than one project, just focus on the last one and identify which project they are talking about. You should focus first on the launchpad id if any in the thread, then the hashtags, and finally the keywords in text, which is likely the unique name of something
"""

async def identify_launchpad_project(tweet_content: str, tweet_id: str = None, network_id: str = "8453") -> ProjectIdentification:
    """
    Stage 2: Identify which launchpad project the tweet is about
    
    Args:
        tweet_content: The text content of the tweet
        tweet_id: Optional tweet ID for tracking
        
    Returns:
        ProjectIdentification with project ID if found, None otherwise
    """
    
    try:    
        client = get_oai_async_client()

        async def wraps() -> ProjectIdentification:
            response = await client.chat.completions.create(
                model=get_model_id(),
                messages=[
                    {"role": "system", "content": SYSTEM_PROMPT},
                    {"role": "user", "content": tweet_content}
                ]
            )

            response_text = response.choices[0].message.content.strip()
            json_match = re.search(r'\{.*\}', response_text, re.DOTALL)
            json_generated_str = json_match.group() if json_match else None 

            if not json_generated_str:
                raise Exception("No JSON object found in the response")

            repaired_json = repair_json(json_generated_str)
            json_obj = json.loads(repaired_json)
            
            proj_details = await get_launchpad_detail(
                launchpad_id=json_obj.get("launchpad_id"),
                network_id=network_id
            )

            if proj_details.result is None:
                raise Exception("No project details found")

            return ProjectIdentification(
                tweet_id=tweet_id,
                launchpad_id=proj_details.result.id,
                project_name=proj_details.result.name,
                description=proj_details.result.description,
                confidence=float_clamp(json_obj.get("confidence", 0.5), 0, 1),
                reasoning=json_obj.get("reasoning"),
            )

        obj = await retry(wraps, max_retry=3, first_interval=10, interval_multiply=1)()
        return obj

    except Exception as e:
        import traceback
        traceback_str = traceback.format_exc()

        asyncio.create_task(async_log(
            traceback_str, 
            channel=settings.lite_logging_channel,
            tags=["project_identifier", "error"],
            server_url=settings.lite_logging_base_url
        ))

        return ProjectIdentification(
            tweet_id=tweet_id or "unknown",
            launchpad_id=None,
            confidence=0.0,
            reasoning=f"Error during project identification: {str(e)}",
        )

