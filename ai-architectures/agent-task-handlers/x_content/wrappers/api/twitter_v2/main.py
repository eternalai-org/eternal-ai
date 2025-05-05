from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.magic import (
    get_response_content,
    helpful_raise_for_status,
    sync2async,
)
from x_content.wrappers.api.twitter_v2.models.response import (
    ExtendedTweetInfoDto,
    GenerateActionDto,
    GetRecentOwnTweetDto,
    InscribeTweetByIdDto,
    Response,
    ExtendedTweetInfosDto,
    SearchTweetDto,
    StructuredInformationDto,
    TweetInfosDto,
    TweetsDto,
    TwitterUserObjectDto,
    TwitterUsersDto,
    UsernamesDto,
)
from x_content.wrappers.api.twitter_v2.models.objects import (
    ExtendedTweetInfo,
    ExtendedTweetObject,
    MentionData,
    StructuredInformation,
    TweetInfo,
    TweetObject,
    TweetType,
    TwitterRequestAuthorization,
    TwitterUserObject,
)
from x_content.wrappers.log_decorators import log_function_call
from x_content.wrappers.vision_tasks import get_image_description

# BAD IMPORT HERE
from x_content.wrappers.llm_tasks import (
    generate_retrieval_query,
)

from typing import Any, List
import requests
import logging
from datetime import datetime

logger = logging.getLogger(__name__)
import re
import string
from x_content import constants as const
from x_content.cache.entity_cache import (
    ConversationRedisCache,
    FollowingListRedisCache,
    ShadowReplyRedisCache,
    TweetInscriptionRedisCache,
)
from x_content.wrappers.browsing import get_cleaned_text
from functools import lru_cache
import traceback
import json
import random

from x_content.wrappers.bing_search import search_from_bing
from x_content.wrappers.rag_search import get_random_from_collections, search_from_db

from datetime import timezone, timedelta


@lru_cache(maxsize=1)
def _get_api_headers():
    return {"api-key": const.TWITTER_API_KEY}


@lru_cache(maxsize=1)
def _get_conversation_redis_cache():
    return ConversationRedisCache()


@lru_cache(maxsize=1)
def _get_following_list_redis_cache():
    return FollowingListRedisCache()


@lru_cache(maxsize=1)
def _get_shadow_reply_redis_cache():
    return ShadowReplyRedisCache()


@lru_cache(maxsize=1)
def _get_tweet_inscription_redis_cache():
    return TweetInscriptionRedisCache()


def _preprocess_username(username: str):
    username = username.lstrip("@")
    if username.startswith("username="):
        username = username.split("=")[1].strip(' "')
    return username


def is_valid_tweet_id(tweet_id):
    if not isinstance(tweet_id, (str, int)):
        return False

    tweet_id_str = str(tweet_id)
    tweet_id_str = tweet_id_str.strip()

    if not tweet_id_str.isdigit():
        return False

    return 18 <= len(tweet_id_str) <= 19


@lru_cache(maxsize=128)
def _image_descriptions_from_tweet_id(tweet_id: str):
    url = f"{const.TWITTER_API_URL}/tweets"
    resp = requests.get(
        url, params={"ids": tweet_id}, headers=_get_api_headers()
    )

    if resp.status_code != 200:
        logger.error(f"Error occurred when calling api: {resp.text}")
        return []

    resp = resp.json()
    data = resp["result"][tweet_id]

    media = data["AttachmentMedia"]

    urls = [m["url"] for m in media if m["type"] == "photo" and m["url"]]

    image_descriptions = []

    for url in urls:
        try:
            description = get_image_description(url)
            image_descriptions.append(description)
        except Exception as err:
            logger.error(
                f"[_image_descriptions_from_tweet_id] Failed to get image description at {url}: {err}"
            )

    return image_descriptions


def _image_urls_from_tweet_id(tweet_id: str):
    try:
        url = f"{const.TWITTER_API_URL}/tweets"
        resp = requests.get(
            url, params={"ids": tweet_id}, headers=_get_api_headers()
        )
        helpful_raise_for_status(resp)

        resp = resp.json()
        data = resp["result"][tweet_id]

        media = data["AttachmentMedia"]

        urls = [m["url"] for m in media if m["type"] == "photo" and m["url"]]
        return urls
    except Exception as e:
        traceback.print_exc()
        logger.error(
            f"[image_urls_from_tweet_id] An unexpected error occured: {e}"
        )
        return []


def optimize_twitter_query(
    query: str,
    remove_punctuations=False,
    token_limit=-1,
    pat: re.Pattern = None,
    length_limit=30,
) -> str:
    and_token = re.compile(r"\bAND\b", flags=re.IGNORECASE)
    spacing = re.compile(r"\s+")

    query = and_token.sub(" ", query)
    query = spacing.sub(" ", query)

    tokenized_query = re.split(r"\bor\b", query, flags=re.IGNORECASE)
    filtered_tokenized_query = []

    if pat is not None:
        tokenized_query = [
            i.strip() for i in tokenized_query if pat.fullmatch(i.strip())
        ]

    # sort and remove duplicates
    tokenized_query = sorted(tokenized_query, key=len, reverse=True)

    for i in tokenized_query:
        i = i.strip(" '\"")

        if remove_punctuations:
            i = "".join([c for c in i if c not in string.punctuation])

        if len(filtered_tokenized_query) == 0:
            filtered_tokenized_query.append(i)
        else:
            if any([i.lower() in x.lower() for x in filtered_tokenized_query]):
                continue
            else:
                filtered_tokenized_query.append(i)

    random.shuffle(filtered_tokenized_query)

    if token_limit != -1:
        filtered_tokenized_query = filtered_tokenized_query[:token_limit]

    if len(filtered_tokenized_query) == 0:
        return ""

    res = ""
    for item in filtered_tokenized_query:
        if len(res) + len(item) > length_limit:
            break

        if len(res) > 0:
            res += " OR "

        res += item

    if len(res) == 0:
        e = tokenized_query[0].split()

        for ee in e:
            if len(res) + len(ee) > length_limit:
                break

            if len(res) > 0:
                res += " "

            res += ee

    return res


@log_function_call
def search_twitter_news(
    query: str,
    impression_count_limit=100,
    limit_api_results=50,
    use_raw=False,
    no_duplication=True,
) -> Response[TweetsDto]:
    try:
        if not use_raw:
            query = optimize_twitter_query(
                query, remove_punctuations=True, token_limit=5, length_limit=30
            )
            logger.info(f"[search_twitter_news] Optimized query: {query}")

        if query.strip() == "":
            logger.error("[search_twitter_news] Empty query")
            return Response(error="Empty query")

        url = f"{const.TWITTER_API_URL}/tweets/search/recent"

        params = {
            "query": f"{query} -is:retweet -is:reply -is:quote is:verified",
            "max_results": limit_api_results,
        }

        resp = requests.get(url, headers=_get_api_headers(), params=params)

        resp = resp.json()
        data = resp["result"]

        json.dumps(2)

        if resp["error"] is not None:
            logger.error(
                "[search_twitter_news] Error occurred when calling api: "
                + resp["error"]["message"]
            )
            return Response(
                error="Error occurred when calling api",
            )

        tweets: List[TweetObject] = []
        hashs = set([])

        for id, item in data["LookUps"].items():
            tweet = item["Tweet"]
            user = item["User"]

            if user is None:
                continue

            if (
                tweet["public_metrics"]["impression_count"]
                < impression_count_limit
            ):
                continue

            content_hash = hash(tweet["text"])

            if no_duplication and content_hash in hashs:
                continue

            hashs.add(content_hash)

            tweets.append(
                TweetObject(
                    tweet_id=tweet["id"],
                    twitter_username=(
                        user.get("username", "Anonymous")
                        if user is not None
                        else "Anonymous"
                    ),
                    twitter_id=tweet["author_id"],
                    like_count=tweet["public_metrics"]["like_count"],
                    retweet_count=tweet["public_metrics"]["retweet_count"],
                    reply_count=tweet["public_metrics"]["reply_count"],
                    impression_count=tweet["public_metrics"][
                        "impression_count"
                    ],
                    full_text=tweet["text"],
                    posted_at=tweet["created_at"],
                )
            )

        return Response(
            data=TweetsDto(
                tweets=tweets,
            )
        )
    except Exception as err:
        logger.error(
            f"[search_twitter_news] An unexpected error occured: {err}"
        )
        return Response(error="An unexpected error occured")


def search_for_token_news(tokens: list) -> Response[TweetsDto]:
    if isinstance(tokens, str):
        tokens = [tokens]

    query = " OR ".join([f"${x}" for x in tokens])
    return search_twitter_news(
        query,
        impression_count_limit=100,
        limit_api_results=50,
        use_raw=True,
        no_duplication=True,
    )


from ... import telegram


@lru_cache(maxsize=512)
def _get_username_by_id(user_id: str):
    try:
        user_url = f"{const.TWITTER_API_URL}/user/{user_id}"
        user_resp = requests.get(user_url, headers=_get_api_headers())
        user_resp_json = user_resp.json()

        if user_resp_json["result"] == None:
            raise Exception(
                f"[_get_username_by_id] User not found, url={user_url}"
            )

        username = user_resp_json["result"]["username"]
        return username

    except Exception as err:
        raise Exception(
            f"[_get_username_by_id] An unexpected error occurred: {err}"
        )


# TODO: Combine this and get_recent_mentioned_tweets_by_username
def get_recent_mentioned_tweets_by_username_v2(
    auth: TwitterRequestAuthorization,
    num_tweets=1,
    replied=0,
    max_num_tweets_in_conversation=3,
    preserve_img=False,
    get_all=False,
) -> Response[ExtendedTweetInfosDto]:
    try:
        conversation_redis_cache = _get_conversation_redis_cache()
        if get_all:
            url = f"{const.TWITTER_API_URL}/user/by/username/{auth.twitter_username}/mentions/all"
            params = {"max_results": 100}
        else:
            url = f"{const.TWITTER_API_URL}/user/by/username/{auth.twitter_username}/mentions"
            params = {"replied": replied}

        resp = requests.get(url, params=params, headers=_get_api_headers())

        if resp.status_code != 200:
            logger.error(
                f"[get_recent_mentioned_tweets_by_username_v2] Something went wrong (status code: {resp.status_code}, url: {resp.url})"
            )

            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nSomething went wrong (status code: {resp.status_code}, resp: {resp.json()}, url: {resp.url})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )

            return Response(
                error=f"Something went wrong (status code: {resp.status_code})",
            )

        resp_json = resp.json()

        if resp_json.get("error"):
            logger.error(
                f"[get_recent_mentioned_tweets_by_username_v2] Error occurred when calling API: {resp_json['error']['message']}, url: {resp.url}"
            )

            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nError occurred when calling API: {resp_json['error']['message']}, url: {resp.url}\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )
            return Response(
                error=f"Error occurred when calling API",
            )

        tweets = resp_json["result"]["data"]
        if not tweets:
            logger.info(
                f"[get_recent_mentioned_tweets_by_username_v2] No tweets found"
            )
            return Response(data=ExtendedTweetInfosDto(tweet_infos=[]))

        tweets_with_media = [
            tweet for tweet in tweets if tweet["attachments"]["media_keys"]
        ]

        tweets_with_media_ids = [tweet["id"] for tweet in tweets_with_media]

        res = []

        for idx, tweet in enumerate(tweets):
            root_conversation_id = tweet["conversation_id"]

            if conversation_redis_cache.is_threshold_exceeded(
                auth.twitter_username,
                root_conversation_id,
                max_num_tweets_in_conversation,
            ):
                continue

            reference_tweets = (
                []
                if tweet["referenced_tweets"] is None
                else tweet["referenced_tweets"]
            )

            parent_tweet_id = next(
                (
                    ref_tweet["id"]
                    for ref_tweet in reference_tweets
                    if ref_tweet["type"] == "replied_to"
                ),
                None,
            )

            author_id = tweet["author_id"]
            try:
                _username = _get_username_by_id(author_id)
            except:
                continue

            # Don't reply to itself
            if _username == auth.twitter_username:
                continue

            if tweet["id"] in tweets_with_media_ids:
                if preserve_img:
                    tweet["image_urls"] = _image_urls_from_tweet_id(
                        tweet["id"]
                    )
                else:
                    image_descriptions = _image_descriptions_from_tweet_id(
                        tweet["id"]
                    )
                    logger.info(
                        f"Image description of tweet {tweet['id']}: {image_descriptions}"
                    )
                    tweet["text"] += "\n\n".join(image_descriptions)

            full_text = get_cleaned_text(tweet.get("text", ""))
            mentions = tweet["entities"].get("mentions", []) or []
            tweet_object = ExtendedTweetObject(
                twitter_id=author_id,
                tweet_id=tweet["id"],
                twitter_username=_username,
                full_text=full_text,
                posted_at=tweet["created_at"],
                image_urls=tweet.get("image_urls", []) or [],
                mentions=[MentionData.from_dict(x) for x in mentions],
            )

            tweet_info = ExtendedTweetInfo(
                tweet_object=tweet_object,
                parent_tweet_id=parent_tweet_id,
                conversation_id=root_conversation_id,
            )

            res.append(tweet_info)

            if len(res) >= num_tweets:
                break

        return Response(
            data=ExtendedTweetInfosDto(
                tweet_infos=res,
            )
        )
    except Exception as err:
        logger.error(
            f"[get_recent_mentioned_tweets_by_username_v2] An unexpected error occured: {err}"
        )
        return Response(error=f"An unexpected error occured")


# TODO: Combine this and get_tweets_by_username_v2
def get_tweets_by_username(
    username: str,
    top_k=5,
    max_num_tweets_in_conversation=3,
    filter_non_replied=False,
    owner_username=None,
) -> Response[TweetsDto]:
    try:
        username = _preprocess_username(username)

        if len(username) == 0:
            logger.error(
                "[get_tweets_by_username] get_tweets_by_username requires a valid username, an empty string is not"
            )
            return Response(
                error="get_tweets_by_username requires a valid username, an empty string is not",
            )

        url = f"{const.TWITTER_API_URL}/tweets/by/username/{username}"
        resp = requests.get(
            url,
            params={"max_results": max(5, top_k)},
            headers=_get_api_headers(),
        )

        if resp.status_code != 200:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nSomething went wrong (status code: {resp.status_code}; resp: {resp.json()}; url: {url})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )

            logger.error(
                f"[get_tweets_by_username] Something went wrong (status code: {resp.status_code}; resp: {resp.json()}; url: {url})"
            )

            return Response(
                error=f"Something went wrong (status code: {resp.status_code})",
            )

        resp = resp.json()

        if resp["error"] is not None:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )
            logger.error(
                "[get_tweets_by_username] Error occurred when calling api: "
                + resp["error"]["message"]
            )
            return Response(
                error="Error occurred when calling api",
            )

        tweets = resp["result"]["data"]
        if not tweets:
            logger.error("No tweets found")
            return Response(
                data=TweetsDto(
                    tweets=[],
                )
            )

        conversation_redis_cache = _get_conversation_redis_cache()

        if filter_non_replied:
            tweets = list(
                filter(
                    lambda x: not conversation_redis_cache.is_threshold_exceeded(
                        owner_username,
                        x["conversation_id"],
                        max_num_tweets_in_conversation=max_num_tweets_in_conversation,
                    ),
                    tweets,
                )
            )

        tweets = [
            TweetObject(
                tweet_id=x["id"],
                twitter_username=username,
                twitter_id=x["author_id"],
                like_count=x["public_metrics"]["like_count"],
                retweet_count=x["public_metrics"]["retweet_count"],
                reply_count=x["public_metrics"]["reply_count"],
                impression_count=x["public_metrics"]["impression_count"],
                full_text=x["text"],
                posted_at=x["created_at"],
                media=[
                    e["expanded_url"]
                    for e in x.get("entities", {}).get("urls", []) or []
                ],
                reference=x.get("referenced_tweets", []),
            )
            for x in tweets
        ]

        # reorder by posted time
        tweets = sorted(tweets, key=lambda x: x.posted_at, reverse=True)

        return Response(data=TweetsDto(tweets=tweets))
    except Exception as err:
        traceback.print_exc()
        return Response(
            error="An unexpected error occured",
        )


from typing import Any, List, Dict


def _get_following_by_username(
    username: str, minimum_followers=None
) -> Response[List[Dict[str, str]]]:
    username = _preprocess_username(username)

    key = f"{username}_followings"

    followings_list_redis_cache = _get_following_list_redis_cache()
    cached_list = followings_list_redis_cache.get(key)

    if cached_list:
        followings = json.loads(cached_list)

    else:
        url = f"{const.TWITTER_API_URL}/user/by/username/{username}/following"
        resp = requests.get(url, headers=_get_api_headers())

        if resp.status_code != 200:
            logger.error(
                f"Something went wrong (status code: {resp.status_code})"
            )
            return Response(
                error=f"Something went wrong (status code: {resp.status_code})",
            )

        resp = resp.json()

        if resp["error"] is not None:
            logger.error(
                f'Error occurred when calling api: {resp["error"]["message"]}'
            )
            return Response(
                error="Error occurred when calling api",
            )

        if resp["result"] is None:
            logger.error(f"Following user not found for username {username}")
            return Response(
                error="Not following any user",
            )

        followings = resp["result"]
        followings = [x for x in followings if x["rest_id"] != ""]
        if len(followings) >= 100:
            followings_list_redis_cache.commit(key, followings)

    return Response(data=followings)


def get_following_by_username(
    username: str, max_users=10, minimum_followers=None
) -> Response[UsernamesDto]:
    try:
        response = _get_following_by_username(username, minimum_followers)

        followings = response.data

        if len(followings) == 0:
            return Response(data=UsernamesDto(usernames=[]))

        if minimum_followers:
            followings = [
                user
                for user in followings
                if user["followers_count"] >= minimum_followers
            ]

        usernames = list(map(lambda x: x["screen_name"], followings))
        usernames = usernames[:max_users]
        return Response(data=UsernamesDto(usernames=usernames))
    except Exception as err:
        logger.error(
            f"[get_following_by_username] An unexpected error occured: {err}"
        )
        return Response(
            error="An unexpected error occured",
        )


async def get_relevent_information_v2(
    kn_base: KnowledgeBase,
    tweet_id: str = None,
    tweets: List[TweetObject] = None,
    task_name: str = "N/A",
    use_bing_search: bool = True,
) -> Response[StructuredInformationDto]:
    if tweets is None and tweet_id is None:
        return Response(error="Either tweet_id or tweets must be provided")

    if tweets is None:
        resp: Response[ExtendedTweetInfosDto] = await sync2async(
            get_full_context_by_tweet_id
        )(tweet_id)

        if resp.is_error():
            return Response(error="Retrieving full context failed")

        tweets = [x.tweet_object for x in resp.data.tweet_infos]

    chat_history = [
        {
            "user": x.twitter_username,
            "message": x.full_text,
        }
        for x in tweets
    ]

    logger.info(
        f"[get_relevent_information_v2] chat history: {json.dumps(chat_history)}"
    )

    try:
        retrieval_query = await sync2async(generate_retrieval_query)(
            chat_history
        )
        if retrieval_query == "":
            return Response(error="Generate retrieval query failed")
    except Exception as err:
        return Response(error=f"Generate retrieval query failed: {err}")

    knowledge = await search_from_db(
        kn_base, retrieval_query, top_k=5, threshold=0.85
    )
    bing_news = []
    if use_bing_search:
        bing_news = await sync2async(search_from_bing)(
            retrieval_query, top_k=10, task_name=task_name
        )
    twitter_resp: Response[TweetsDto] = await sync2async(search_twitter_news)(
        retrieval_query,
        limit_api_results=10,
        use_raw=True,
    )
    if not twitter_resp.is_error():
        twitter_news = [x.full_text for x in twitter_resp.data.tweets]
    else:
        twitter_news = []

    news = bing_news + twitter_news

    structured_information = StructuredInformation(
        knowledge=knowledge,
        news=news,
    )

    return Response(
        data=StructuredInformationDto(
            structured_information=structured_information
        )
    )


@lru_cache(maxsize=128)
def _get_tweet_info_from_tweet_id(
    tweet_id: str, preserve_img=False
) -> ExtendedTweetInfo:
    try:
        if not is_valid_tweet_id(tweet_id):
            raise Exception(f"'{tweet_id}' is not a valid tweet id")

        url = f"{const.TWITTER_API_URL}/tweets"
        resp = requests.get(
            url, params={"ids": tweet_id}, headers=_get_api_headers()
        )

        helpful_raise_for_status(resp)
        result = resp.json().get("result", {})

        if len(result) == 0:
            raise Exception(f"Tweet id {tweet_id} not found")

        key, value = list(result.items())[0]
        tweet = value["Tweet"]
        user = value["User"]
        media = value["AttachmentMedia"]

        image_descriptions = []
        if media is not None:
            urls = [
                m["url"] for m in media if m["type"] == "photo" and m["url"]
            ]

            if preserve_img:
                tweet["image_urls"] = urls
            else:
                for url in urls:
                    try:
                        description = get_image_description(url)
                        image_descriptions.append("\n\n" + description)
                    except Exception as err:
                        logger.error(
                            f"[_get_tweet_info_from_tweet_id] Failed to get image description at {url}: {err}"
                        )
                logger.info(
                    f"Image description of tweet {tweet['id']}: {image_descriptions}"
                )
                tweet["text"] += "".join(image_descriptions)

        user = user or {}

        mentions = tweet["entities"].get("mentions", []) or []
        tweet_object = ExtendedTweetObject(
            tweet_id=tweet["id"],
            twitter_id=user.get("id", "N/A"),
            twitter_username=user.get("username", "Anonymous"),
            full_text=tweet["text"],
            posted_at=tweet["created_at"],
            image_urls=tweet.get("image_urls", []) or [],
            mentions=[MentionData.from_dict(x) for x in mentions],
        )

        parent_tweet_id = None

        if key != tweet["conversation_id"]:
            reference_tweets = (
                []
                if tweet["referenced_tweets"] is None
                else tweet["referenced_tweets"]
            )
            parent_tweet_id = next(
                (
                    ref_tweet["id"]
                    for ref_tweet in reference_tweets
                    if ref_tweet["type"] == "replied_to"
                ),
                None,
            )

        tweet_info = ExtendedTweetInfo(
            tweet_object=tweet_object,
            parent_tweet_id=parent_tweet_id,
            conversation_id=tweet["conversation_id"],
        )

        return ExtendedTweetInfoDto(tweet_info=tweet_info)
    except Exception as e:
        traceback.print_exc()
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        logger.error(
            f"[get_tweet_info_from_tweet_id] (tweet_id={tweet_id}) An unexpected error occurred: {e}"
        )
        raise Exception("An unexpected error occurred")


@log_function_call
def get_tweet_info_from_tweet_id(
    tweet_id: str, preserve_img=False
) -> Response[ExtendedTweetInfoDto]:
    try:
        data = _get_tweet_info_from_tweet_id(tweet_id, preserve_img)
        return Response(data=data)
    except Exception as e:
        return Response(error=str(e))


template_msg = """
<strong>{agent_name} has made a {action_type}!</strong>
<i><strong>Ref-ID</strong> {ref_id};
<strong>Request-ID</strong> {nav};
<strong>Task</strong>: {task};
<strong>Toolset</strong>: {toolset};</i>
{line_str}
{action_input}
{line_str}
<strong>Success</strong>: {success}
{additional_info}
"""


def notify_agent_action(
    auth: TwitterRequestAuthorization,
    action_type: str,
    action_input: dict,
    success: bool,
    response: dict = None,
):
    additional_info = ""

    if not success:
        additional_info = "<strong>Response</strong>:\n<pre>{}</pre>".format(
            json.dumps(response, indent=2)
        )

    action_input = "<strong>Input</strong>:\n<pre>{}</pre>".format(
        json.dumps(action_input, indent=2)
    )
    line_str = "-" * 25

    msg = template_msg.format(
        agent_name=auth.twitter_username,
        action_type=action_type,
        task=auth.task,
        toolset=auth.toolset,
        ref_id=auth.ref_id,
        nav=auth.request_id,
        action_input=action_input,
        success=success,
        additional_info=additional_info,
        line_str=line_str,
    )

    telegram.send_message(
        "junk_nofitications", msg, room=telegram.TELEGRAM_ROOM, fmt="HTML"
    )


def generate_action(
    auth: TwitterRequestAuthorization,
    action_type: str,
    action_input: dict,
    tx_hash="",
) -> Response[GenerateActionDto]:
    try:
        url = f"{const.TWITTER_API_URL}/user/action"

        payload = {
            "agent_contract_id": str(auth.agent_contract_id),
            "chain_id": int(auth.chain_id),
            "action_type": action_type,
            "action_input": action_input,
            "ref_id": auth.ref_id,
            "inscribe_tx_hash": tx_hash,
        }

        response = requests.post(url, json=payload, headers=_get_api_headers())
        success = response.status_code == 200
        if success:
            logger.info(
                f"[generate_action] User {auth.twitter_username} performing action with payload {json.dumps(payload)}"
            )
        else:
            logger.error(
                f"[generate_action] User {auth.twitter_username} performing action failed, status_code={response.status_code}, payload={json.dumps(payload)}, response={get_response_content(response)}"
            )

        notify_agent_action(
            auth,
            action_type,
            action_input,
            success,
            get_response_content(response),
        )

        return Response(data=GenerateActionDto(success=success))
    except Exception as err:
        traceback.print_exc()
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        logger.error(f"[generate_action] An unexpected error occurred: {err}")
        return Response(error="An unexpected error occurred")


def follow(
    auth: TwitterRequestAuthorization, target_username: str
) -> Response[GenerateActionDto]:
    resp = generate_action(
        auth=auth,
        action_type="follow",
        action_input={
            "target_username": target_username,
        },
    )
    if resp.is_error():
        return Response(error="Generate follow action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )


def reply(
    auth: TwitterRequestAuthorization,
    tweet_id: str,
    reply_content: str,
    tx_hash="",
    max_num_tweets_in_conversation=3,
) -> Response[GenerateActionDto]:
    try:
        resp = get_tweet_info_from_tweet_id(tweet_id)
        if resp.is_error():
            return Response(error="Failed to get tweet info")

        tweet_info = resp.data.tweet_info.to_dict()
        conversation_id = tweet_info["conversation_id"]
        conversation_redis = _get_conversation_redis_cache()

        if conversation_redis.is_threshold_exceeded(
            auth.twitter_username,
            conversation_id,
            max_num_tweets_in_conversation=max_num_tweets_in_conversation,
        ):
            return Response(error="Conversation is already replied")

        conversation_redis.add_tweet_to_conversation(
            auth.twitter_username, conversation_id, tweet_info["tweet_object"]
        )

        resp = generate_action(
            auth=auth,
            action_type="reply",
            action_input={"tweet_id": tweet_id, "comment": reply_content},
            tx_hash=tx_hash,
        )
        if resp.is_error():
            return Response(error="Generate reply action failed")

        return Response(
            data=GenerateActionDto(
                success=resp.data.success,
            )
        )
    except Exception as err:
        traceback.print_exc()
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        logger.error(f"[reply] An unexpected error occurred: {err}")
        return Response(error="An unexpected error occurred")


def reply_multi(
    auth: TwitterRequestAuthorization,
    tweet_id: str,
    reply_content: str,
    tx_hash="",
):
    resp = generate_action(
        auth=auth,
        action_type="reply_multi",
        action_input={"tweet_id": tweet_id, "comment": reply_content},
        tx_hash=tx_hash,
    )

    if resp.is_error():
        return Response(error="Generate reply_multi action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )

    # if success:
    #     return f"Schedule to reply {tweet_id}", None

    # return f"Failed to reply {tweet_id}", Exception("Failed to reply")


def reply_multi_unlimited(
    auth: TwitterRequestAuthorization,
    tweet_id: str,
    reply_content: str,
    tx_hash="",
):
    resp = generate_action(
        auth=auth,
        action_type="reply_multi_unlimited",
        action_input={"tweet_id": tweet_id, "comment": reply_content},
        tx_hash=tx_hash,
    )

    if resp.is_error():
        return Response(error="Generate reply_multi_unlimited action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )


def shadow_reply(
    auth: TwitterRequestAuthorization,
    tweet_id: str,
    reply_content: str,
    tx_hash="",
):
    try:
        shadow_reply_redis = _get_shadow_reply_redis_cache()

        if shadow_reply_redis.is_threshold_exceeded(
            auth.twitter_username, tweet_id
        ):
            return f"Tweet is already replied"

        shadow_reply_redis.add_reply(auth.twitter_username, tweet_id)

    except Exception as err:
        logger.info(f"[shadow_reply] Error while checking reply {err}")
        return f"[shadow_reply] Error while checking reply {err}"

    resp = generate_action(
        auth=auth,
        action_type="reply",
        action_input={"tweet_id": tweet_id, "comment": reply_content},
        tx_hash=tx_hash,
    )

    if resp.is_error():
        return Response(error="Generate shadow_reply action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )

    # if success:
    #     return f"Schedule to reply {tweet_id}"

    # return f"Failed to reply {tweet_id}"


def quote_tweet(
    auth: TwitterRequestAuthorization, tweet_id: str, comment: str, tx_hash=""
):
    resp = generate_action(
        auth=auth,
        action_type="quote_tweet",
        action_input={"tweet_id": tweet_id, "comment": comment},
        tx_hash=tx_hash,
    )
    if resp.is_error():
        return Response(error="Generate quote tweet action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )


def tweet(
    auth: TwitterRequestAuthorization, content: str, tx_hash=""
) -> Response[GenerateActionDto]:
    resp = generate_action(
        auth=auth,
        action_type="tweet",
        action_input={"content": content},
        tx_hash=tx_hash,
    )

    if resp.is_error():
        return Response(error="Generate tweet action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )


def tweet_multi(auth: TwitterRequestAuthorization, content: List[str]):
    resp = generate_action(
        auth=auth,
        action_type="tweet_multi",
        action_input={"content": json.dumps(content)},
    )

    if resp.is_error():
        return Response(error="Generate tweet_multi action failed")

    return Response(
        data=GenerateActionDto(
            success=resp.data.success,
        )
    )


def inscribe_tweet_by_id(
    auth: TwitterRequestAuthorization,
    id: str,
    price: str,
    reason: str,
    tweet_type: TweetType,
) -> Response[InscribeTweetByIdDto]:
    try:
        if not is_float(price):
            return Response(error="Given price is not a valid float number")

        tweet_inscription_redis = _get_tweet_inscription_redis_cache()

        if tweet_inscription_redis.is_threshold_exceeded(
            auth.twitter_username, auth.request_id
        ):
            return Response(error="At most one inscription can be done")

        resp = get_tweet_info_from_tweet_id(id)

        if resp.is_error():
            return Response(error="Tweet id not found")

        tweet_info = resp.data.tweet_info.to_dict()

        content = tweet_info["tweet_object"]["full_text"]
        tweet_inscription_redis.add_inscription(
            auth.twitter_username, auth.request_id, id
        )

        action_input = {
            "tweet_id": id,
            "content": content,
            "price": price,
            "reason": reason,
        }

        action_type = (
            "inscribe_tweet"
            if tweet_type == TweetType.POST
            else "inscribe_reply"
        )
        resp = generate_action(
            auth=auth,
            action_type=action_type,
            action_input=action_input,
        )

        if resp.is_error():
            return Response(error="Generate inscribe tweet action failed")

        return Response(
            data=InscribeTweetByIdDto(
                success=resp.data.success,
                metadata={
                    "tweet_id": id,
                    "content": content,
                    "price": price,
                    "reason": reason,
                },
            )
        )
    except Exception as err:
        traceback.print_exc()
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        logger.error(f"[create_token] An unexpected error occurred: {err}")
        return Response(error="An unexpected error occurred")


# TODO: Move to utils.py
def is_float(xx: Any):
    try:
        float(xx)
        return True
    except ValueError:
        return False


def create_token(
    auth: TwitterRequestAuthorization,
    name: str,
    symbol: str,
    description: str,
    announcement_content: str,
) -> Response[GenerateActionDto]:
    try:
        symbol = symbol.upper()

        is_valid_symbol = lambda symbol: len(symbol) <= 8 and all(
            c.isalnum() or c.isalpha() for c in symbol
        )
        is_valid_name = lambda name: len(name) <= 20

        if not is_valid_symbol(symbol):
            return Response(
                error="Invalid symbol. Symbol must be alphanumeric and less or equal to 8 characters"
            )

        if not is_valid_name(name):
            return Response(
                error="Invalid name. Name must be less or equal to 20 characters"
            )

        create_token_action_input = {
            "name": name,
            "symbol": symbol,
            "description": description,
            "content": announcement_content,
        }

        resp = generate_action(
            auth=auth,
            action_type="create_token",
            action_input=create_token_action_input,
        )

        if resp.is_error():
            return Response(error="Generate create token action failed")

        return Response(
            data=GenerateActionDto(
                success=resp.data.success,
            )
        )
    except Exception as err:
        traceback.print_exc()
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        logger.error(f"[create_token] An unexpected error occurred: {err}")
        return Response(error="An unexpected error occurred")


def is_reply(tweet: dict[str, str]) -> bool:
    reference_tweets = (
        []
        if tweet["referenced_tweets"] is None
        else tweet["referenced_tweets"]
    )

    parent_tweet_id = next(
        (
            ref_tweet["id"]
            for ref_tweet in reference_tweets
            if ref_tweet["type"] == "replied_to"
        ),
        None,
    )

    return parent_tweet_id is not None


def is_post(tweet: dict[str, str]) -> bool:
    reference_tweets = (
        []
        if tweet["referenced_tweets"] is None
        else tweet["referenced_tweets"]
    )
    return len(reference_tweets) == 0


def is_correct_tweet_type(tweet: dict[str, str], type_whitelist):
    if type_whitelist == [] or type_whitelist is None:
        return True

    if TweetType.POST in type_whitelist and is_post(tweet):
        return True

    if TweetType.REPLY in type_whitelist and is_reply(tweet):
        return True

    return False


# TODO: Combine with get_tweets_by_username?
def get_posts_or_reply_by_username(
    username: str,
    top_k=10,
    length_limit=None,
    search_start: datetime = None,
    search_end: datetime = None,
    type_whitelist: List[TweetType] = [],
) -> Response[TweetsDto]:
    username = _preprocess_username(username)

    if len(username) == 0:
        return Response(
            error="get_replies_by_username requires a valid username, an empty string is not"
        )

    url = f"{const.TWITTER_API_URL}/tweets/by/username/{username}"
    resp = requests.get(url, headers=_get_api_headers())

    if resp.status_code != 200:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nSomething went wrong (status code: {resp.status_code}; response body: {resp.text}; url: {url})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )

        return Response(
            error=f"Something went wrong (status code: {resp.status_code})"
        )

    resp = resp.json()

    if resp["error"] is not None:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )

        return Response(
            error="Error occurred when calling api: "
            + resp["error"]["message"]
        )

    tweets = resp["result"]["data"]

    if len(tweets) == 0:
        return Response(error="No tweets found")

    posts = []
    # Get only posts, no reply, no retweet
    for tweet in tweets:
        if not is_correct_tweet_type(tweet, type_whitelist):
            continue

        if length_limit != None and len(tweet["text"]) > length_limit:
            continue

        if (
            search_start != None
            and tweet["created_at"] < search_start.isoformat()
        ):
            continue

        if search_end != None and tweet["created_at"] > search_end.isoformat():
            continue

        posts.append(tweet)

    posts = [
        TweetObject(
            tweet_id=x["id"],
            twitter_username=username,
            twitter_id=x["author_id"],
            like_count=x["public_metrics"]["like_count"],
            retweet_count=x["public_metrics"]["retweet_count"],
            reply_count=x["public_metrics"]["reply_count"],
            impression_count=x["public_metrics"]["impression_count"],
            full_text=x["text"],
            posted_at=x["created_at"],
        )
        for x in posts
    ]

    # reorder by posted time
    posts = sorted(posts, key=lambda x: x.posted_at, reverse=True)

    # Limit to 20 most recent posts
    posts = posts[:top_k]
    return Response(
        data=TweetsDto(tweets=posts),
    )


def get_own_recent_tweets(
    auth: TwitterRequestAuthorization, type_whitelist=[]
) -> Response[GetRecentOwnTweetDto]:
    try:
        search_end = datetime.now(tz=timezone.utc)
        search_start = search_end - timedelta(hours=24)
        resp = get_posts_or_reply_by_username(
            auth.twitter_username,
            top_k=50,
            length_limit=300,
            search_start=search_start,
            search_end=search_end,
            type_whitelist=type_whitelist,
        )

        if resp.is_error():
            logger.error(
                f"[get_own_recent_tweets] Error retrieving tweet by username: {resp.error}"
            )
            return Response(
                error="Error retrieving tweet by username",
            )

        tweets = resp.data.tweets

        tweet_inscription_redis = _get_tweet_inscription_redis_cache()
        inscribed_tweet_ids = tweet_inscription_redis.get_inscribed_tweets_ids(
            auth.twitter_username
        )

        tweets = list(
            filter(lambda x: x.tweet_id not in inscribed_tweet_ids, tweets)
        )

        return Response(
            data=GetRecentOwnTweetDto(
                tweets=tweets,
                search_start=search_start.isoformat(),
                search_end=search_end.isoformat(),
                tweet_count=len(tweets),
            )
        )
    except Exception as err:
        traceback.print_stack()
        logger.error(
            f"[get_own_recent_tweets] An unexpected error occurred: {err}"
        )
        return Response(
            error="An unexpected error occurred",
        )


# TODO: Combine with get_tweets_by_username
def get_tweets_by_username_v2(
    username: str, num_tweets=1, replied=0, filter_non_reply=False
) -> Response[TweetInfosDto]:
    try:
        username = _preprocess_username(username)

        if len(username) == 0:
            logger.error(
                "[get_tweets_by_username_v2] get_tweets_by_username_v2 requires a valid username, an empty string is not"
            )
            return Response(
                error="get_tweets_by_username_v2 requires a valid username, an empty string is not"
            )

        token = None

        res = []

        while len(res) < num_tweets:
            url = f"{const.TWITTER_API_URL}/tweets/by/username/{username}"
            params = {}

            if replied != None:
                params["replied"] = replied

            if token != None:
                params["pagination_token"] = token

            resp = requests.get(url, params=params, headers=_get_api_headers())

            if resp.status_code != 200:
                logger.error(
                    f"Something went wrong (status code: {resp.status_code}; url: {resp.url}; text: {resp.text})"
                )

                telegram.send_message(
                    "junk_nofitications",
                    f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {resp.url}; text: {resp.text})\n</pre>",
                    room=telegram.TELEGRAM_ALERT_ROOM,
                )

                return Response(
                    error=f"Something went wrong (status code: {resp.status_code})"
                )

            resp_json = resp.json()

            if resp_json.get("error"):
                logger.error(
                    f"Error occurred when calling API: {resp_json['error']['message']}, url: {resp.url}"
                )

                telegram.send_message(
                    "junk_nofitications",
                    f"<pre>\nError occurred when calling API: {resp_json['error']['message']}, url: {resp.url}\n</pre>",
                    room=telegram.TELEGRAM_ALERT_ROOM,
                )

                return Response(error=f"Error occurred when calling API")

            tweets = resp_json["result"]["data"]

            if tweets == None:
                return Response(error=f"User's tweet not found")

            for tweet in tweets:
                reference_tweets = (
                    []
                    if tweet["referenced_tweets"] is None
                    else tweet["referenced_tweets"]
                )

                parent_tweet_id = next(
                    (
                        ref_tweet["id"]
                        for ref_tweet in reference_tweets
                        if ref_tweet["type"] == "replied_to"
                    ),
                    None,
                )

                if filter_non_reply and len(reference_tweets) > 0:
                    continue

                author_id = tweet["author_id"]
                full_text = get_cleaned_text(tweet.get("text", ""))
                tweet_object = TweetObject(
                    twitter_id=author_id,
                    tweet_id=tweet["id"],
                    twitter_username=username,
                    full_text=full_text,
                    posted_at=tweet["created_at"],
                )

                tweet_info = TweetInfo(
                    tweet_object=tweet_object,
                    parent_tweet_id=parent_tweet_id,
                    conversation_id=tweet["conversation_id"],
                )

                res.append(tweet_info)

                if len(res) >= num_tweets:
                    return Response(data=TweetInfosDto(tweet_infos=res))

            break
        return Response(data=TweetInfosDto(tweet_infos=res))
    except Exception as e:
        traceback.print_exc()
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        logger.error(
            f"[get_tweets_by_username_v2] An unexpected error occurred: {e}"
        )
        return Response(error="An unexpected error occurred")


def get_full_context_of_tweet(
    tweet_info: ExtendedTweetInfo,
) -> Response[ExtendedTweetInfosDto]:
    try:
        tweets = [tweet_info]
        parent_tweet_id = tweet_info.parent_tweet_id
        while parent_tweet_id is not None and len(tweets) < 20:
            resp = get_tweet_info_from_tweet_id(parent_tweet_id)
            if resp.is_error():
                break

            tweet_info = resp.data.tweet_info

            tweet_info.tweet_object.full_text = get_cleaned_text(
                tweet_info.tweet_object.full_text
            )
            tweets.insert(0, tweet_info)
            parent_tweet_id = tweet_info.parent_tweet_id

        return Response(
            data=ExtendedTweetInfosDto(tweet_infos=tweets),
        )
    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"[get_full_context_of_tweet] An unexpected error occurred: {err}"
        )
        return Response(error="An unexpected error occured")


def get_full_context_by_tweet_id(
    tweet_id: str,
) -> Response[ExtendedTweetInfosDto]:
    try:
        resp = get_tweet_info_from_tweet_id(tweet_id)
        if resp.is_error():
            return Response(error=resp.error)
        return get_full_context_of_tweet(resp.data.tweet_info)
    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"[get_full_context_by_tweet_id] An unexpected error occurred: {err}"
        )
        return Response(error="An unexpected error occured")


def get_full_conversation_from_liked_tweets(
    auth: TwitterRequestAuthorization,
    num_tweets=1,
    replied=0,
    ignore_replied_tweets=False,
) -> Response[TweetInfosDto]:
    try:
        res = []
        shadow_reply_redis = _get_shadow_reply_redis_cache()

        url = f"{const.TWITTER_API_URL}/user/liked"
        params = {"replied": replied}
        resp = requests.get(url, params=params, headers=_get_api_headers())
        liked_tweet_infos = resp.json()["result"]

        if liked_tweet_infos == None:
            return Response(data=TweetInfosDto(tweet_infos=res))

        for idx, liked_tweet_info in enumerate(liked_tweet_infos):
            tweet_id = liked_tweet_info["tweet_id"]
            if ignore_replied_tweets:
                if shadow_reply_redis.is_threshold_exceeded(
                    auth.twitter_username, tweet_id
                ):
                    continue

            twitter_id = liked_tweet_info["twitter_id"]

            try:
                twitter_username = _get_username_by_id(twitter_id)
            except:
                continue

            in_reply_to_tweet_id = (
                None
                if liked_tweet_info["in_reply_to_tweet_id"] == ""
                else liked_tweet_info["in_reply_to_tweet_id"]
            )
            full_text = get_cleaned_text(liked_tweet_info.get("full_text", ""))
            tweet_object = TweetObject(
                twitter_id=twitter_id,
                tweet_id=tweet_id,
                twitter_username=twitter_username,
                full_text=full_text,
                posted_at=liked_tweet_info["posted_at"],
            )

            tweet_info = TweetInfo(
                tweet_object=tweet_object,
                parent_tweet_id=in_reply_to_tweet_id,
            )

            res.append(tweet_info)
            if len(res) >= num_tweets:
                return Response(data=TweetInfosDto(tweet_infos=res))

        return Response(data=TweetInfosDto(tweet_infos=res))
    except Exception as e:
        telegram.send_message(
            "junk_nofitications",
            f"""<pre>\n{traceback.format_exc()}\n</pre>""",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        traceback.print_exc()
        logger.error(
            f"[get_full_conversation_from_liked_tweets] An unexpected error occurred: {e}"
        )
        return Response(error=f"An unexpected error occurred")


def search_users(query: str) -> Response[TwitterUsersDto]:
    url = f"{const.TWITTER_API_URL}/user/search/"

    validate_pat = re.compile(r"[A-Za-z0-9_']{4,15}")
    optimized_query = optimize_twitter_query(
        query, remove_punctuations=True, pat=validate_pat
    )

    if len(optimized_query) == 0:
        return Response(
            error="Search users failed. Usernames must be alphanumeric with the length between 4-15 characters, no spaces, no punctuations except underscores"
        )

    logger.info(f"Optimized query: {optimized_query}")
    params = {
        "query": optimized_query,
    }

    resp = requests.get(url, params=params, headers=_get_api_headers())

    if resp.status_code != 200:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {resp.url}; text: {resp.text})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        return Response(
            error=f"Something went wrong (status code: {resp.status_code})"
        )

    resp = resp.json()
    if resp["error"] is not None:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        return Response(
            error="Error occurred when calling api: "
            + resp["error"]["message"]
        )

    if resp["result"] is None or len(resp["result"]) == 0:
        return Response(error="No user found with query: " + query)

    users = [
        TwitterUserObject(
            twitter_id=x["id"],
            username=x["username"],
            name=x["name"],
            followers_count=x["public_metrics"]["followers_count"],
            followings_count=x["public_metrics"]["following_count"],
            is_blue_verified=x["verified"],
        )
        for x in resp["result"]
    ]

    # Only get 20 top users
    users = users[:10]
    return Response(data=TwitterUsersDto(users=users))


# TODO: add to tool_call
def search_recent_retweeted_users(
    query: str, limit_observation=10
) -> Response[TwitterUsersDto]:
    if query.strip() == "":
        return Response(
            error="search_recent_tweets requires a valid query, an empty string is not",
        )

    max_results = limit_observation
    url = f"{const.TWITTER_API_URL}/tweets/search/recent"

    optimized_query = optimize_twitter_query(query)
    logger.info(f"Optimized query: {optimized_query}")
    if len(optimized_query) == 0:
        return Response(
            error="search_recent_tweets requires a valid query, an empty string is not",
        )

    params = {
        "query": optimized_query,
        "max_results": max_results,
    }
    resp = requests.get(url, params=params, headers=_get_api_headers())

    if resp.status_code != 200:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {resp.url}; text: {resp.text})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        return Response(
            error=f"Something went wrong (status code: {resp.status_code})"
        )

    resp = resp.json()
    data = resp["result"]

    if resp["error"] is not None:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        return Response(
            error="Error occurred when calling api: "
            + resp["error"]["message"]
        )

    if len(data["LookUps"]) == 0:
        return Response(error="No tweets found with query: " + query)

    ids = []
    for id, item in data["LookUps"].items():
        tweet = item["Tweet"]
        user = item["User"]

        if tweet["referenced_tweets"]:
            for t in tweet["referenced_tweets"]:
                if t["type"] == "retweeted":
                    ids.append(t["id"])

    url = f"{const.TWITTER_API_URL}/tweets"
    params = {
        "ids": ",".join(ids),
    }

    resp = requests.get(url, params=params, headers=_get_api_headers())

    if resp.status_code != 200:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {resp.url}; text: {resp.text})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        return Response(
            error=f"Something went wrong (status code: {resp.status_code})"
        )

    resp = resp.json()

    if resp["error"] is not None:
        telegram.send_message(
            "junk_nofitications",
            f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
            room=telegram.TELEGRAM_ALERT_ROOM,
        )
        return Response(error="No tweets found")

    data = resp["result"]
    res: List[TwitterUserObject] = []

    for id in ids:
        if id not in data:
            continue

        tweet = data[id]["Tweet"]
        user = data[id]["User"]
        res.append(
            TwitterUserObject(
                twitter_id=user["id"],
                username=user["username"],
                name=user["name"],
                followers_count=user["public_metrics"]["followers_count"],
                followings_count=user["public_metrics"]["following_count"],
                is_blue_verified=user["verified"],
            )
        )

    res = res[:10]
    return Response(data=TwitterUsersDto(users=res))


def search_recent_tweet_by_tweetid(
    tweet_id: str, limit_observation=10
) -> Response[SearchTweetDto]:
    try:
        query = f"conversation_id:{tweet_id}"

        if query.strip() == "":
            return Response(
                error=f"Invalid empty query",
            )

        max_results = limit_observation
        url = f"{const.TWITTER_API_URL}/tweets/search/recent"
        optimized_query = query.strip()

        if len(optimized_query) == 0:
            logger.error(f"[search_recent_tweets] Invalid query: {query}")
            return Response(
                error=f"Invalid query",
            )

        params = {
            "query": f"{optimized_query}",
            "max_results": max_results,
        }
        resp = requests.get(url, params=params, headers=_get_api_headers())

        if resp.status_code != 200:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {resp.url}; resp: {resp.json()})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )
            return Response(
                error=f"Something went wrong (status code: {resp.status_code})",
            )

        resp = resp.json()
        data = resp["result"]

        if resp["error"] is not None:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )

            return Response(
                error="Error occurred when calling api",
            )

        if len(data["LookUps"]) == 0:
            logger.info(f"No tweets found with query: {optimized_query}")
            return Response(
                data=SearchTweetDto(
                    optimized_query=optimized_query,
                    tweets=[],
                )
            )

        tweets: List[TweetObject] = []
        for id, item in data["LookUps"].items():
            tweet = item["Tweet"]
            user = item["User"]

            reference_tweets = (
                []
                if tweet["referenced_tweets"] is None
                else [
                    ref
                    for ref in tweet["referenced_tweets"]
                    if ref["id"] == tweet_id
                ]
            )

            if not reference_tweets:
                continue

            tweets.append(
                TweetObject(
                    tweet_id=tweet["id"],
                    twitter_username=(
                        user.get("username", "Anonymous")
                        if user is not None
                        else "Anonymous"
                    ),
                    twitter_id=tweet["author_id"],
                    like_count=tweet["public_metrics"]["like_count"],
                    retweet_count=tweet["public_metrics"]["retweet_count"],
                    reply_count=tweet["public_metrics"]["reply_count"],
                    impression_count=tweet["public_metrics"][
                        "impression_count"
                    ],
                    full_text=tweet["text"],
                    posted_at=tweet["created_at"],
                )
            )

        tweets = [x for x in tweets if x.impression_count > 0]

        tweets = sorted(tweets, key=lambda x: x.posted_at, reverse=True)
        result = random.sample(tweets, min(len(tweets), limit_observation))

        return Response(
            data=SearchTweetDto(
                optimized_query=optimized_query,
                tweets=result,
            )
        )
    except Exception as err:
        logger.error(
            f"[search_recent_tweets] An unexpected error occured: {err}"
        )
        return Response(
            error="An unexpected error occured",
        )


def search_recent_tweets(
    query: str, limit_observation=10
) -> Response[SearchTweetDto]:
    try:
        if query.strip() == "":
            return Response(
                error=f"Invalid empty query",
            )

        max_results = limit_observation
        url = f"{const.TWITTER_API_URL}/tweets/search/recent"

        optimized_query = optimize_twitter_query(query, token_limit=1)
        logger.info(f"Optimized query: {optimized_query}")
        if len(optimized_query) == 0:
            logger.error(f"[search_recent_tweets] Invalid query: {query}")
            return Response(
                error=f"Invalid query",
            )

        params = {
            "query": f"{optimized_query} -is:retweet -is:reply -is:quote is:verified",
            "max_results": max_results,
        }
        resp = requests.get(url, params=params, headers=_get_api_headers())

        if resp.status_code != 200:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {resp.url}; resp: {resp.json()})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )
            return Response(
                error=f"Something went wrong (status code: {resp.status_code})",
            )

        resp = resp.json()
        data = resp["result"]

        if resp["error"] is not None:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )

            return Response(
                error="Error occurred when calling api",
            )

        if len(data["LookUps"]) == 0:
            logger.info(f"No tweets found with query: {optimized_query}")
            return Response(
                data=SearchTweetDto(
                    optimized_query=optimized_query,
                    tweets=[],
                )
            )

        tweets: List[TweetObject] = []
        for id, item in data["LookUps"].items():
            tweet = item["Tweet"]
            user = item["User"]

            tweets.append(
                TweetObject(
                    tweet_id=tweet["id"],
                    twitter_username=(
                        user.get("username", "Anonymous")
                        if user is not None
                        else "Anonymous"
                    ),
                    twitter_id=tweet["author_id"],
                    like_count=tweet["public_metrics"]["like_count"],
                    retweet_count=tweet["public_metrics"]["retweet_count"],
                    reply_count=tweet["public_metrics"]["reply_count"],
                    impression_count=tweet["public_metrics"][
                        "impression_count"
                    ],
                    full_text=tweet["text"],
                    posted_at=tweet["created_at"],
                )
            )

        tweets = [x for x in tweets if x.impression_count > 0]

        tweets = sorted(tweets, key=lambda x: x.posted_at, reverse=True)
        result = random.sample(tweets, min(len(tweets), limit_observation))

        return Response(
            data=SearchTweetDto(
                optimized_query=optimized_query,
                tweets=result,
            )
        )
    except Exception as err:
        logger.error(
            f"[search_recent_tweets] An unexpected error occured: {err}"
        )
        return Response(
            error="An unexpected error occured",
        )


def get_popular_following_feed(
    auth: TwitterRequestAuthorization, top_k=10, minimum_followers=5000
) -> Response[TweetsDto]:
    try:
        resp = get_following_by_username(
            auth.twitter_username,
            max_users=20,
            minimum_followers=minimum_followers,
        )

        if resp.is_error():
            return Response(error=resp.error)

        # random pick for 5
        res: List[TweetObject] = []
        choices = random.sample(
            resp.data.usernames,
            k=min(top_k, len(resp.data.usernames)),
        )

        for choice in choices:
            resp = get_posts_or_reply_by_username(
                choice, top_k=5, type_whitelist=[TweetType.POST]
            )
            if resp.is_error():
                continue
            res.extend(resp.data.tweets)

        return Response(
            data=TweetsDto(
                tweets=res,
            )
        )
    except Exception as err:
        logger.error(
            f"[search_recent_tweets] An unexpected error occured: {err}"
        )
        return Response(
            error="An unexpected error occured",
        )


def get_user_info_by_username(username: str) -> Response[TwitterUserObjectDto]:
    try:
        username = _preprocess_username(username)

        url = f"{const.TWITTER_API_URL}/user/by/username/{username}"
        resp = requests.get(url, headers=_get_api_headers())

        if resp.status_code != 200:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nSomething went wrong (status code: {resp.status_code}; url: {url}; text: {resp.text})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )

            logger.error(
                f"[get_user_info_by_username] Something went wrong (status code: {resp.status_code}; url: {url}; text: {resp.text})"
            )

            return Response(
                error=f"Something went wrong (status code: {resp.status_code}"
            )

        resp = resp.json()
        if resp["error"] is not None:
            telegram.send_message(
                "junk_nofitications",
                f"<pre>\nError occurred when calling api (msg: {resp['error']['message']}; url: {url})\n</pre>",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )

            logger.error(
                f"[get_user_info_by_username] Error occurred when calling api (msg: {resp['error']['message']})"
            )

            return Response(
                error=f"Error occurred when calling api (msg: {resp['error']['message']})"
            )

        info = resp["result"]
        if info["id"] == "":
            return Response(error=f"Username not found")

        user = TwitterUserObject(
            twitter_id=info["id"],
            username=info["username"],
            name=info["name"],
            followers_count=info["public_metrics"]["followers_count"],
            followings_count=info["public_metrics"]["following_count"],
            is_blue_verified=info["verified"],
        )

        return Response(data=TwitterUserObjectDto(user=user))
    except Exception as err:
        logger.error(
            f"[get_user_info_by_username] An unexpected error occured: {err}"
        )
        return Response(
            error="An unexpected error occured",
        )


def get_tweet_with_image_description_appended_to_text(
    tweet_info: ExtendedTweetInfo,
) -> ExtendedTweetInfo:
    tweet_id = tweet_info.tweet_object.tweet_id
    image_descriptions = _image_descriptions_from_tweet_id(tweet_id)
    logger.info(f"Image description of tweet {tweet_id}: {image_descriptions}")
    tweet_info.tweet_object.full_text += "\n\n".join(image_descriptions)
    return tweet_info
