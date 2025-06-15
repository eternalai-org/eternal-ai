from pydantic import BaseModel, fields
from typing import Optional, List, Any
from datetime import datetime

class TweetAttachment(BaseModel):
    media_keys: Optional[List[str]] = None
    poll_ids: Optional[List[str]] = None
    
class TweetContentAnnotation(BaseModel):
    start: int
    end: int
    probability: float
    type: str
    normalized_text: str
    
class TweetEmbeddedURL(BaseModel):
    start: int
    end: int
    url: str
    expanded_url: str
    display_url: str
    status: int
    title: str
    description: str
    unwound_url: str
    
class TweetHashtag(BaseModel):
    pass

class TweetMention(BaseModel):
    pass

class TweetCashtag(BaseModel):
    pass 

class TwitterEntities(BaseModel):
    annotations: Optional[List[TweetContentAnnotation]] = None
    urls: Optional[List[TweetEmbeddedURL]] = None
    hashtags: Optional[List[TweetHashtag]] = None
    mentions: Optional[List[TweetMention]] = None
    cashtags: Optional[List[TweetCashtag]] = None

class TweetGeo(BaseModel):
    place_id: str
    coordinates: Optional[dict[str, Any]] = None

class TweetNonPublicMetrics(BaseModel):
    impression_count: int
    url_link_clicks: int
    user_profile_clicks: int
    like_count: int
    reply_count: int
    retweet_count: int
    quote_count: int
    
class OrganicMetrics(BaseModel):
    impression_count: int
    url_link_clicks: int
    user_profile_clicks: int
    like_count: int
    reply_count: int
    retweet_count: int
    quote_count: int
    
class PromotedMetrics(BaseModel):
    impression_count: int
    url_link_clicks: int
    user_profile_clicks: int
    like_count: int
    reply_count: int
    retweet_count: int
    quote_count: int
    
class PublicMetrics(BaseModel):
    impression_count: int
    url_link_clicks: int
    user_profile_clicks: int
    like_count: int
    reply_count: int
    retweet_count: int
    quote_count: int

class Tweet(BaseModel):
    id: str
    text: str
    attachments: Optional[TweetAttachment] = None
    author_id: str
    context_annotations: Optional[List[str]] = None
    conversation_id: str
    created_at: str
    entities: Optional[TwitterEntities] = None
    geo: Optional[TweetGeo] = None
    in_reply_to_user_id: Optional[str] = None
    lang: Optional[str] = None
    non_public_metrics: Optional[TweetNonPublicMetrics] = None
    organic_metrics: Optional[TweetNonPublicMetrics] = None
    possiby_sensitive: bool = False
    promoted_metrics: Optional[TweetNonPublicMetrics] = None
    public_metrics: Optional[TweetNonPublicMetrics] = None
    referenced_tweets: Optional[List[Any]] = None
    source: Optional[str] = None
    withheld: Optional[dict[str, Any]] = None
    note_tweet: Optional[dict[str, Any]] = None

    @property
    def created_timestamp(self) -> int:
        try:
            return int(datetime.strptime(self.created_at, "%Y-%m-%dT%H:%M:%S.%fZ").timestamp())
        except Exception as e:
            return 0

class Pagination(BaseModel):
    oldest_id: str
    newest_id: str
    result_count: int
    next_token: Optional[str] = None
    previous_token: Optional[str] = None

class TweetPage(BaseModel):
    data: List[Tweet]
    meta: Pagination
    
class ProfilePublicMetrics(BaseModel):
    followers_count: int
    following_count: int
    tweet_count: int
    listed_count: int


# sample data
class ConnectionCard(BaseModel):
    rest_id: str
    screen_name: str
    name: str
    profile_image_url_https: Optional[str] = None
    followers_count: int
    friends_count: int
    is_blue_verified: bool
    created_at: str


class TwitterUserInfo(BaseModel):
    id: str
    name: str
    username: str
    created_at: str
    description: str
    entities: Optional[dict[str, Any]] = None
    location: str
    pinned_tweet_id: str
    profile_image_url: Optional[str] = None
    protected: bool = False
    public_metrics: Optional[ProfilePublicMetrics] = None
    url: str
    verified: bool = False
    withheld: Optional[dict[str, Any]] = None

    followers: Optional[List[ConnectionCard]] = []
    following: Optional[List[ConnectionCard]] = []
