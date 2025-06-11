from fastmcp import FastMCP
from app.utils.twitter_api_calls import (
    get_twitter_user_info_by_id as get_twitter_user_info_by_id_api,
    list_tweets_of_user as list_tweets_of_user_api,
    get_tweet_info as get_tweet_info_api,
    get_twitter_user_info_by_username as get_twitter_user_info_by_username_api,
)
from typing import List, Optional

mcp = FastMCP(
    name="twitter_info",
    description="Explore twitter profiles",
)

@mcp.tool(
    name="get_twitter_user_info_by_id",
    description="Get the info of a twitter user by id",
    annotations={
        "user_id": "the id of the twitter user"
    }
)
async def get_twitter_user_info_by_id(user_id: str) -> Optional[dict]:
    req = await get_twitter_user_info_by_id_api(user_id)
    result = req.result

    if result is None:
        return f"Failed to get twitter user info for id {user_id!r}; Details: {req.error}"

    return {
        "id": result.id,
        "username": result.username,
        "name": result.name,
        "description": result.description,
        "metrics": result.public_metrics.model_dump(mode="json"),
        "verified": result.verified
    }
    

@mcp.tool(
    name="get_twitter_user_info_by_username",
    description="Get the info of a twitter user by username",
    annotations={
        "username": "the username of the twitter user"
    }
)
async def get_twitter_user_info_by_username(username: str) -> Optional[dict]:
    req = await get_twitter_user_info_by_username_api(username)
    result = req.result

    if result is None:
        return f"Failed to get twitter user info for username {username!r}; Details: {req.error}"

    return {
        "id": result.id,
        "username": result.username,
        "name": result.name,
        "description": result.description,
        "metrics": result.public_metrics.model_dump(mode="json"),
        "verified": result.verified
    }


@mcp.tool(
    name="list_tweets_of_user",
    description="List the tweets of a twitter user. Use the pagination_token to get the next page of tweets if it presents.",
    annotations={
        "user_id": "the id of the twitter user",
        "pagination_token": "the token to paginate the tweets, leave empty to get the first page (latest tweets)"
    }
)
async def list_tweets_of_user(user_id: str, pagination_token: str = "") -> Optional[dict]:
    req =  await list_tweets_of_user_api(user_id, pagination_token)
    result = req.result

    if result is None:
        return f"Failed to get tweets of profile id {user_id!r}; Details: {req.error}"

    tweets = [
        {
            "id": tweet.id,
            "text": tweet.text,
            "created_at": tweet.created_at,
            "metrics": tweet.public_metrics.model_dump(mode="json"),
            "in_reply_to_user_id": tweet.in_reply_to_user_id,
            "referenced_tweets": tweet.referenced_tweets,
        }
        for tweet in result.data
        if len(tweet.referenced_tweets or []) == 0
    ]

    pagination_token = result.meta.next_token

    return {
        "tweets": tweets,
        "pagination_token": pagination_token,
        "count": result.meta.result_count,
    }

@mcp.tool(
    name="get_tweet_info",
    description="Get the info of a tweet by id",
    annotations={
        "tweet_id": "the id of the tweet"
    }
)
async def get_tweet_info(tweet_id: str) -> Optional[dict]:
    req = await get_tweet_info_api(tweet_id)
    result = req.result

    if result is None:
        return f"Failed to get tweet info for id {tweet_id!r}; Details: {req.error}"

    return {
        "id": result.id,
        "text": result.text,
        "created_at": result.created_at,
        "metrics": result.public_metrics.model_dump(mode="json"),
        "in_reply_to_user_id": result.in_reply_to_user_id,
        "referenced_tweets": result.referenced_tweets,
        "author_id": result.author_id,
    }