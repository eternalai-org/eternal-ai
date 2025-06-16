from pydantic import BaseModel
from typing import Optional

class EvaluationRequest(BaseModel):
    twitter_id: str
    tweet_id: str
    tweet_content: str
    original_tweet: str
    launchpad_id: Optional[str] = None
    network_id: str = "8453"