import httpx
from typing import Union, Any, Dict, List, Set
from app.schemas import (
    twitter,
    commons
)
from pydantic import ValidationError
from app.utils.caching import mongo_cache, set_cache_value, get_cached_value, delete_cached_value
import logging
from app.config import settings

logger = logging.getLogger(__name__)

TWITTER_API_URL = settings.twitter_api_url
TWITTER_API_KEY = settings.twitter_api_key
TWITTER_USERNAME_TO_ID = "twitter_username_to_id"
# key builders
TIMEOUT_CFG = httpx.Timeout(60.0, connect=10.0) 

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

def twitter_tweet_key_builder_w_page(**kwargs) -> str:
    pagination_token = kwargs.get('pagination_token', '') 

    if 'user_id' in kwargs and kwargs['user_id']:
        user_id = kwargs['user_id']
        return f"twitter-profile:{user_id}-{pagination_token}"

    elif 'username' in kwargs and kwargs['username']:
        username = kwargs['username']
        user_id = get_cached_value(TWITTER_USERNAME_TO_ID, username)
        return f"twitter-profile:{user_id or username}-{pagination_token}"

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
        timeout=TIMEOUT_CFG,
    ) as client: 
        try:
            res = await client.get(
                f"{twitter_api_base_url}/tweets", 
                params={"ids": tweet_id},
            )
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
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
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            res = await client.get(url, params={"id": user_id})
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        res_json: list[dict[str, Any]] = res.json().get('result', []) or []
        
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
    url = f"{twitter_api_base_url}/user/{user_id}/following_v1"

    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            res = await client.get(url)
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )
        
        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        res_json: list[dict[str, Any]] = res.json().get('result', []) or []

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
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            res = await client.get(url)
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )
        
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
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            res = await client.get(url)
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

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
    ttl_seconds=3600 * 6,
    key_builder=twitter_tweet_key_builder_w_page,
    object_builder=lambda data: commons.ResponseMessage[twitter.TweetPage].model_validate(data)
)
async def list_tweets_of_user(
    user_id: Union[str, int],
    pagination_token: str = "",
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY
) -> commons.ResponseMessage[twitter.TweetPage]:
    response_model = commons.ResponseMessage[twitter.TweetPage]
    url = f"{twitter_api_base_url}/tweets/{user_id}"
    
    async with httpx.AsyncClient(
        headers={
            "api-key": twitter_api_key,
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            res = await client.get(
                url, 
                params={
                   "pagination_token": pagination_token
                }
            )
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )
        
        if res.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=f"Error {res.status_code}: {res.text}"
            )

        res_json: dict[str, Any] = res.json().get('result')

        try:
            obj = twitter.TweetPage.model_validate(res_json)
            return response_model(result=obj)

        except ValidationError as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

is_valid_user = lambda user: (
    isinstance(user, twitter.TwitterUserInfo) 
    and user.id 
    and user.username
)

from queue import Queue

async def build_twitter_social_graph(
    user_id: str, 
    max_depth: int = 1,
    max_expansion: int = 10,
    twitter_api_base_url: str = TWITTER_API_URL, 
    twitter_api_key: str = TWITTER_API_KEY,
) -> Dict[str, List[str]]:
    user_req = await get_twitter_user_info_by_id(user_id, twitter_api_base_url, twitter_api_key)

    if not is_valid_user(user_req.result):
        return {}

    user = user_req.result

    graph: Dict[str, Set[str]] = {}

    users_map = {
        user.id: user
    }

    que = Queue() 
    que.put((user, 0))

    while not que.empty():
        user, depth = que.get()
        user: twitter.TwitterUserInfo

        if depth > max_depth:
            continue

        for follower in user.followers[:max_expansion]:
            if follower.rest_id not in users_map:
                info_req = await get_twitter_user_info_by_id(follower.rest_id, twitter_api_base_url, twitter_api_key)

                if is_valid_user(info_req.result):
                    users_map[follower.rest_id] = info_req.result
                    que.put((info_req.result, depth + 1))

        for following in user.following[:max_expansion]:
            if following.rest_id not in users_map:
                info_req = await get_twitter_user_info_by_id(following.rest_id, twitter_api_base_url, twitter_api_key)

                if is_valid_user(info_req.result):
                    users_map[following.rest_id] = info_req.result
                    que.put((info_req.result, depth + 1))

    logger.info(f"Built social graph for {user_id} with {len(users_map)} users")

    for id in users_map:
        graph[id] = set([])

    for id, user in users_map.items():
        for follower in user.followers: 
            if follower.rest_id in graph:
                graph[follower.rest_id].add(id)

            else:
                graph[follower.rest_id] = set([id])

        for following in user.following:
            if following.rest_id in graph:
                graph[following.rest_id].add(id)
            else:
                graph[following.rest_id] = set([id])

    return {
        k: list(v)
        for k, v in graph.items()
    }

from app.utils.misc import dsu

async def get_tweet_threads_by_twitter_id(
    user_id: str,
    max_calls: int = 5,
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY,
) -> commons.ResponseMessage[dict[str, list[twitter.Tweet]]]:

    response_model = commons.ResponseMessage[dict[str, list[twitter.Tweet]]]
    tweets: list[twitter.Tweet] = []

    current_page = ""

    for i in range(max_calls):
        req = await list_tweets_of_user(
            user_id, 
            pagination_token=current_page,
            twitter_api_base_url=twitter_api_base_url,
            twitter_api_key=twitter_api_key
        ) 

        if req.result is None:
            logger.error(f"Error getting tweets for {user_id}: {req.error}")
            break

        tweets.extend(req.result.data)
        next_page = req.result.meta.next_token

        if current_page == next_page or next_page == "":
            break

        current_page = next_page

    map_idx = {
        val.id: i
        for i, val in enumerate(tweets)
    }

    relations = []

    for i, tweet in enumerate(tweets):
        for ref in (tweet.referenced_tweets or []):
            _type, _id = ref.get('type'), ref.get('id')

            if _type == "replied_to" and _id in map_idx:
                relations.append((map_idx[_id], i))

    parent = dsu(len(tweets), relations)
    unique_threads = set(parent)

    threads: dict[str, list[twitter.Tweet]] = {}

    for thread in unique_threads:
        threads[str(tweets[thread].id)] = [
            tweets[i]
            for i in range(len(tweets))
            if parent[i] == thread
        ]

    for thread in threads:
        threads[thread].sort(key=lambda x: x.created_timestamp)

    return response_model(result=threads)

async def get_tweet_threads_by_tweet_id(
    tweet_id: str,
    max_depth: int = 5,
    twitter_api_base_url: str = TWITTER_API_URL,
    twitter_api_key: str = TWITTER_API_KEY,
) -> commons.ResponseMessage[list[twitter.Tweet]]:
    response_model = commons.ResponseMessage[list[twitter.Tweet]]

    tweet_req = await get_tweet_info(tweet_id, twitter_api_base_url, twitter_api_key)
    current_tweet = tweet_req.result
    tweets: list[twitter.Tweet] = [current_tweet]

    for _ in range(max_depth):
        if current_tweet is None or not current_tweet.referenced_tweets:
            logger.info(f"No more referenced tweets for {tweet_id}; Exiting")
            break

        for ref in (current_tweet.referenced_tweets or []):
            ref: dict[str, Any]
            _type, _id = ref.get('type'), ref.get('id')

            if _type == "replied_to":
                tweet_req = await get_tweet_info(_id, twitter_api_base_url, twitter_api_key)

                if tweet_req.result is not None:
                    current_tweet = tweet_req.result
                    tweets.append(current_tweet)

                else:
                    logger.error(f"Error getting tweet {_id} for {tweet_id}; Message: {tweet_req.error}")

                break

    return response_model(result=tweets)
