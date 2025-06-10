import httpx
from typing import Union, Any, Dict, List
import os
from app.schemas import (
    twitter,
    commons
)
from pydantic import ValidationError
from app.utils.caching import mongo_cache, set_cache_value, get_cached_value, delete_cached_value

TWITTER_API_URL = os.getenv("TWITTER_API_URL")
TWITTER_API_KEY = os.getenv("TWITTER_API_KEY")
TWITTER_USERNAME_TO_ID = "twitter_username_to_id"
# key builders

def tweet_key_builder(tweet_id: Union[str, int], **kwargs) -> str:
    return f"tweet:{tweet_id}"

def twitter_profile_key_builder(**kwargs) -> str:
    if 'user_id' in kwargs and kwargs['user_id']:
        user_id = kwargs['user_id']
        return f"twitter-profile:{user_id}"

    elif 'username' in kwargs and kwargs['username']:
        username = kwargs['username']
        user_id = get_cached_value(TWITTER_USERNAME_TO_ID, username)
        return f"twitter-profile:{user_id or username}"

    else:
        raise ValueError("Either user_id or username must be provided")

@mongo_cache(
    collection_name="tweets",
    key_prefix="twitter",
    key_builder=tweet_key_builder,
    object_builder=lambda data: commons.ResponseMessage[twitter.Tweet].model_validate(data)
)
async def get_tweet_info(
    tweet_id: Union[str, int],
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[twitter.Tweet]:
    response_model = commons.ResponseMessage[twitter.Tweet]

    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
    ) as client: 
        res = await client.get(
            f"{twitter_api_base_url}/tweets", 
            params={"ids": tweet_id},
        )

        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        data = res.json()
        results: dict[str, Any] = data['result']

        if str(tweet_id) not in results:
            return response_model(
                status=commons.APIStatus.NOT_FOUND, 
                error=f"Tweet {tweet_id} not found"
            )

        tweet_data = results[str(tweet_id)]['Tweet']

        try:
            tweet = twitter.Tweet.model_validate(tweet_data)
            return response_model(result=tweet)
        except ValidationError as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

async def get_top_followers(
    user_id: Union[str, int],
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[List[twitter.ConnectionCard]]:
    response_model = commons.ResponseMessage[List[twitter.ConnectionCard]]
    url = f"{twitter_api_base_url}/user/follower"
    
    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
    ) as client:
        res = await client.get(url, params={"id": user_id})

        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        res_json: list[dict[str, Any]] = res.json()['result']
        
        try:
            obj = [
                twitter.ConnectionCard.model_validate(x) 
                for x in res_json
            ]

            return response_model(result=obj)

        except ValidationError as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )
            

async def get_top_following(
    user_id: Union[str, int],
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[List[twitter.ConnectionCard]]:
    response_model = commons.ResponseMessage[List[twitter.ConnectionCard]]
    url = f"{twitter_api_base_url}/user/{user_id}/following"
    
    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
    ) as client:
        res = await client.get(url)
        
        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        res_json: list[dict[str, Any]] = res.json()['result']

        try:
            obj = [
                twitter.ConnectionCard.model_validate(x) 
                for x in res_json
            ]

            return response_model(result=obj)

        except ValidationError as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

@mongo_cache(
    collection_name="twitter-users",
    key_prefix="twitter",
    key_builder=twitter_profile_key_builder,
    object_builder=lambda data: commons.ResponseMessage[twitter.TwitterUserInfo].model_validate(data)
)
async def get_twitter_user_info_by_id(
    user_id: Union[str, int], 
    twitter_api_base_url: str = TWITTER_API_URL, 
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[twitter.TwitterUserInfo]:
    response_model = commons.ResponseMessage[twitter.TwitterUserInfo]
    url = f"{twitter_api_base_url}/user/{user_id}"

    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
    ) as client:
        res = await client.get(url)
        
        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )
            
        res_json: dict[str, Any] = res.json()
        
        if not res_json.get('result'):
            return response_model(
                status=commons.APIStatus.NOT_FOUND, 
                error=f"User {user_id!r} not found"
            )
        
        try:
            obj = twitter.TwitterUserInfo.model_validate(res_json['result'])
            set_cache_value(TWITTER_USERNAME_TO_ID, obj.username, obj.id)

            followers_resp = await get_top_followers(obj.id)
            obj.followers = followers_resp.result or []
            
            following_resp = await get_top_following(obj.id)
            obj.following = following_resp.result or []

            return response_model(result=obj)
        except ValidationError as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

@mongo_cache(
    collection_name="twitter-users",
    key_prefix="twitter",
    key_builder=twitter_profile_key_builder,
    object_builder=lambda data: commons.ResponseMessage[twitter.TwitterUserInfo].model_validate(data)
)
async def get_twitter_user_info_by_username(
    username: str, 
    twitter_api_base_url: str = TWITTER_API_URL, 
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[twitter.TwitterUserInfo]:
    response_model = commons.ResponseMessage[twitter.TwitterUserInfo]
    url = f"{twitter_api_base_url}/user/by/username/{username}"
    
    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
    ) as client:
        res = await client.get(url)

        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        res_json: dict[str, Any] = res.json()

        if not res_json.get('result'):
            return response_model(
                status=commons.APIStatus.NOT_FOUND, 
                error=f"User {username!r} not found"
            )

        user_info = res_json['result']
        
        try:
            user = twitter.TwitterUserInfo.model_validate(user_info)
            set_cache_value(TWITTER_USERNAME_TO_ID, username, user.id)

            followers_resp = await get_top_followers(user.id)
            user.followers = followers_resp.result or []
            
            following_resp = await get_top_following(user.id)
            user.following = following_resp.result or []

            return response_model(result=user)
        except ValidationError as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )


@mongo_cache(
    collection_name="tweets",
    key_prefix="twitter",
    ttl_seconds=3600,
    key_builder=tweet_key_builder,
    object_builder=lambda data: commons.ResponseMessage[twitter.Tweet].model_validate(data)
)
async def list_tweets_of_user(
    user_id: Union[str, int],
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[List[twitter.Tweet]]:
    response_model = commons.ResponseMessage[List[twitter.Tweet]]
    