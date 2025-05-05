import logging

from json_repair import repair_json

from x_content.constants import ToolSet
from x_content.llm.base import OpenAILLMBase, OnchainInferResult
from x_content.models import (
    ToolDef,
    ToolParam,
    ToolParamDtype,
    ToolLabel,
    AgentMetadata,
)
from x_content.wrappers.conversation import (
    get_enhance_tweet_conversation,
    get_llm_result_by_model_name,
    get_reply_tweet_conversation,
)
from x_content.wrappers import bing_search, trading
from typing import List

from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.objects import (
    TweetInfo,
    TweetObject,
    TweetType,
)
from x_content.wrappers.api.twitter_v2.models.response import (
    ExtendedTweetInfoDto,
    ExtendedTweetInfosDto,
    GenerateActionDto,
    Response,
    TweetsDto,
)
from x_content.wrappers.magic import get_agent_llm_first_interval, retry, sync2async
from x_content.wrappers.postprocess import postprocess_tweet_by_prompts
from x_content.wrappers.tweet_specialty import TweetSpecialty, detect_tweet_specialties
from x_content.wrappers.twin_agent import get_random_example_tweets
from .toolcall import IToolCall

logger = logging.getLogger(__name__)

MAX_TEXT_LENGTH = 10000
MIN_TEXT_LENGTH_TO_SUMMARIZE = 100

from x_content.utils import notify_trading_action

_map = {
    ToolSet.DEFAULT: [
        "get_user_info_by_username",
        "get_tweets_by_username",
        "get_following_by_username",
        "get_recent_mentioned_tweets_by_username",
        "search_users",
        "search_recent_tweets",
        "tweet",
        "reply",
        "follow",
    ],
    ToolSet.REPLY_NON_MENTIONS: [
        "get_popular_following_feed",
        "search_recent_tweets",
        "reply_by_tweet_id",
    ],
    ToolSet.REPLY_NON_MENTIONS_TRADITIONAL: [
        "get_popular_following_feed",
        "search_recent_tweets",
        "get_tweet_full_context",
        "reply",
    ],
    ToolSet.FOLLOW: [
        "get_following_by_username",
        "search_users",
        "follow",
    ],
    ToolSet.POST: [
        "get_popular_following_feed",
        "search_recent_tweets",
        # "research_about_topic",
        "tweet",
    ],
    ToolSet.QUOTE_TWEET: [
        "search_recent_tweets",
        "get_tweets_by_username",
        "get_tweet_full_context",
        "quote_tweet",
    ],
    ToolSet.TRADING: [
        "get_wallet_balance",
        "get_token_prices",
        "buy",
        "sell",
        "search_recent_tweets",
        "get_tweet_full_context",
    ],
    ToolSet.INSCRIBE_TWEET: [
        "get_recent_posts",
        "inscribe_post_by_id",
    ],
    ToolSet.ISSUE_TOKEN: [
        "get_your_own_tweets",
        "create_token",
    ],
    ToolSet.INSCRIBE_REPLY: [
        "get_recent_replies",
        "inscribe_reply_by_id",
    ],
}


def _tweet_obj_to_observation(tweet: TweetObject) -> dict:
    return {
        "tweet_id": tweet.tweet_id,
        "twitter_username": tweet.twitter_username,
        "full_text": tweet.full_text,
    }


class LiveXDB(IToolCall):

    def __init__(
        self,
        auth: twitter_v2.TwitterRequestAuthorization,
        agent_config: AgentMetadata,
        llm: OpenAILLMBase,
    ):
        self.auth = auth
        self.agent_config = agent_config
        self.llm = llm

    def quote_tweet(self, tweet_id: str, comment: str):
        resp = twitter_v2.quote_tweet(
            auth=self.auth,
            tweet_id=tweet_id,
            comment=comment,
        )

        if resp.is_error():
            return resp.error

        if not resp.data.success:
            return f"Failed to quote tweet {tweet_id}"

        return f"Schedule to quote tweet {tweet_id}"

    def tweet(self, content: str):
        resp = twitter_v2.tweet(self.auth, content)

        if resp.is_error():
            return resp.error

        if not resp.data.success:
            return "Failed to schedule the tweet"

        return "The tweet is scheduled to be posted"

    def tweet_multi(self, content: List[str]):
        resp = twitter_v2.tweet_multi(self.auth, content)

        if resp.is_error():
            return resp.error

        if not resp.data.success:
            return "Failed to schedule the tweet thread"

        return "The tweet thread is scheduled to be posted"

    async def tweet_with_enhancement(self, content: str):
        enhance_tweet_conversation = get_enhance_tweet_conversation(
            system_prompt=self.agent_config.persona,
            content=content,
            example_tweets=await sync2async(get_random_example_tweets)(
                self.auth.knowledge_id
            ),
        )

        async def get_enhanced_tweet():
            result: OnchainInferResult = await self.llm.agenerate(
                enhance_tweet_conversation, temperature=0.01
            )
            content = result.generations[0].message.content
            content = get_llm_result_by_model_name(
                content, self.auth.model_name
            )
            data = repair_json(content, return_objects=True)
            return data["tweet"]

        try:
            enhanced_tweet = await retry(
                get_enhanced_tweet,
                max_retry=3,
                first_interval=get_agent_llm_first_interval(),
                interval_multiply=2,
            )()
        except Exception as err:
            logger.error(
                f"[tweet_with_enhancement] Failed to enhance the tweet {err}"
            )
            return "Failed to enhance the tweet"

        try:
            postprocessed_tweet = await sync2async(
                postprocess_tweet_by_prompts
            )(
                system_prompt=self.agent_config.persona,
                task_prompt=self.auth.prompt,
                tweet=enhanced_tweet,
            )
        except Exception as err:
            logger.error(
                f"[tweet_with_enhancement] Failed to postprocess the tweet {err}"
            )
            return "Failed to postprocess the tweet"

        resp: Response[GenerateActionDto] = await sync2async(twitter_v2.tweet)(
            auth=self.auth,
            content=postprocessed_tweet,
        )

        if resp.is_error():
            return resp.error

        metadata = {
            "tool_call": "tweet_with_enhancement",
            "original_tweet": content,
            "enhanced_tweet": enhanced_tweet,
            "postprocessed_tweet": postprocessed_tweet,
        }

        if not resp.data.success:
            return f"Failed to schedule the tweet", metadata

        return "The tweet is scheduled to be posted", metadata

    def reply(self, tweet_id: str, reply_content: str):
        resp = twitter_v2.reply(self.auth, tweet_id, reply_content)

        if resp.is_error():
            return resp.error

        if not resp.data.success:
            return f"Failed to reply {tweet_id}"

        return f"Schedule to reply {tweet_id}"

    async def reply_by_tweet_id(self, tweet_id: str):
        resp: Response[ExtendedTweetInfoDto] = await sync2async(
            twitter_v2.get_tweet_info_from_tweet_id
        )(tweet_id)

        if resp.is_error():
            return "Error getting tweet by id"

        tweet_info = resp.data.tweet_info

        specialties: List[TweetSpecialty] = await sync2async(
            detect_tweet_specialties
        )(tweet_info)

        if len(specialties) > 0:
            return "Cannot reply this tweet. Please select another tweet to reply."

        context_resp: Response[ExtendedTweetInfosDto] = await sync2async(
            twitter_v2.get_full_context_of_tweet
        )(tweet_info)

        if context_resp.is_error() or len(context_resp.data.tweet_infos) == 0:
            return "Error getting tweet full context"

        tweets = [x.tweet_object for x in context_resp.data.tweet_infos]

        info_resp = await twitter_v2.get_relevent_information_v2(
            self.auth.kn_base,
            tweets=tweets,
            task_name=f"{self.auth.task}:{self.auth.toolset}",
            use_bing_search=False,
        )

        conversational_chat = await get_reply_tweet_conversation(
            system_prompt=self.agent_config.persona,
            task_prompt=self.auth.prompt,
            tweets=tweets,
            structured_info=info_resp.data.structured_information,
        )

        async def get_reply_tweet():
            result: OnchainInferResult = await self.llm.agenerate(
                conversational_chat, temperature=0.7
            )
            content = result.generations[0].message.content
            content = get_llm_result_by_model_name(
                content, self.auth.model_name
            )
            data = repair_json(content, return_objects=True)
            return data["tweet"]

        try:
            reply_tweet = await retry(
                get_reply_tweet,
                max_retry=3,
                first_interval=get_agent_llm_first_interval(),
                interval_multiply=2,
            )()
        except Exception as err:
            logger.error(
                f"[reply_by_tweet_id] Failed to get reply tweet {err}"
            )
            return "Failed to get reply tweet"

        try:
            postprocessed_reply = await sync2async(
                postprocess_tweet_by_prompts
            )(
                system_prompt=self.agent_config.persona,
                task_prompt=self.auth.prompt,
                tweet=reply_tweet,
            )
        except Exception as err:
            logger.error(
                f"[reply_by_tweet_id] Failed to postprocess the tweet {err}"
            )
            return "Failed to postprocess the tweet"

        metadata = {
            "tool_name": "reply_by_tweet_id",
            "params": {
                "tweet_id": tweet_id,
            },
            "metadata": {
                "base_reply": reply_tweet,
                "postprocessed_reply": postprocessed_reply,
            },
        }

        action_resp: Response[GenerateActionDto] = await sync2async(
            twitter_v2.reply
        )(
            auth=self.auth,
            tweet_id=tweet_id,
            reply_content=reply_tweet,
        )

        if action_resp.is_error():
            return action_resp.error, metadata

        if not action_resp.data.success:
            return f"Failed to reply to tweet {tweet_id}", metadata

        return f"Schedule to reply {tweet_id}", metadata

    def follow(self, target_username: str):
        resp = twitter_v2.follow(self.auth, target_username=target_username)
        if resp.is_error():
            return resp.error
        if not resp.data.success:
            return f"Failed to follow {target_username}"
        return f"Decided to follow {target_username}"

    def create_token(
        self,
        name: str,
        symbol: str,
        description: str,
        announcement_content: str,
    ):
        resp = twitter_v2.create_token(
            auth=self.auth,
            name=name,
            symbol=symbol,
            description=description,
            announcement_content=announcement_content,
        )

        if resp.is_error():
            return resp.error

        if not resp.data.success:
            return "Failed to schedule token creation"

        return "Token creation is scheduled"

    def inscribe_reply_by_id(self, tweet_id: str, price: str, reason: str):
        resp = twitter_v2.inscribe_tweet_by_id(
            auth=self.auth,
            id=tweet_id,
            price=price,
            reason=reason,
            tweet_type=TweetType.REPLY,
        )

        if resp.is_error():
            return resp.error

        metadata = {
            "tool_name": "inscribe_reply_by_id",
            "params": {
                "id": tweet_id,
                "price": price,
                "reason": reason,
            },
            "metadata": resp.data.metadata,
        }

        if not resp.data.success:
            return "Failed to schedule the inscription", metadata

        return "The reply is scheduled to be inscribed", metadata

    def inscribe_post_by_id(self, tweet_id: str, price: str, reason: str):
        resp = twitter_v2.inscribe_tweet_by_id(
            auth=self.auth,
            id=tweet_id,
            price=price,
            reason=reason,
            tweet_type=TweetType.POST,
        )

        metadata = {
            "tool_name": "inscribe_post_by_id",
            "params": {
                "id": tweet_id,
                "price": price,
                "reason": reason,
            },
            "metadata": resp.data.metadata,
        }

        if resp.is_error():
            return resp.error

        if not resp.data.success:
            return "Failed to schedule the inscription", metadata

        return "The post is scheduled to be inscribed", metadata

    def get_tweets_by_username(self, username: str):
        resp = twitter_v2.get_tweets_by_username(
            username=username,
            top_k=10,
            filter_non_replied=True,
            owner_username=self.auth.twitter_username,
        )

        if resp.is_error():
            return resp.error

        if resp.data.tweets == []:
            return "No tweets found"

        return [_tweet_obj_to_observation(x) for x in resp.data.tweets]

    def get_user_info_by_username(self, username: str):
        resp = twitter_v2.get_user_info_by_username(username=username)

        if resp.is_error():
            return resp.error

        return resp.data.user.to_dict()

    def get_tweets_by_username_selfblock(self, username: str):
        if username == self.auth.twitter_username:
            return "Search your own tweets is not allowed"

        resp = twitter_v2.get_tweets_by_username(
            username=username,
            top_k=10,
            filter_non_replied=True,
            owner_username=self.auth.twitter_username,
        )

        if resp.is_error():
            return resp.error

        if resp.data.tweets == []:
            return "No tweets found"

        return [_tweet_obj_to_observation(x) for x in resp.data.tweets]

    def get_your_own_tweets(self):
        resp = twitter_v2.get_tweets_by_username_v2(
            username=self.auth.twitter_username, num_tweets=30
        )

        if resp.is_error():
            return resp.error

        if len(resp.data.tweet_infos) == 0:
            return "No tweets found"

        return [
            _tweet_obj_to_observation(x.tweet_object)
            for x in resp.data.tweet_infos
        ]

    def get_recent_posts(self):
        resp = twitter_v2.get_own_recent_tweets(
            self.auth, type_whitelist=[TweetType.POST]
        )

        if resp.is_error():
            return resp.error

        metadata = {
            "tool_name": "get_recent_posts",
            "params": {},
            "metadata": {
                "search_start": resp.data.search_start.isoformat(),
                "search_end": resp.data.search_end.isoformat(),
                "tweet_count": len(resp.data.tweets),
            },
        }

        return [
            _tweet_obj_to_observation(x) for x in resp.data.tweets
        ], metadata

    def get_recent_mentioned_tweets_by_username(self, username: str):
        resp = twitter_v2.get_recent_mentioned_tweets_by_username_v2(
            auth=self.auth
        )

        if resp.is_error():
            return resp.error

        if resp.data.tweet_infos == []:
            return "No tweet found"

        return [
            _tweet_obj_to_observation(x.tweet_object)
            for x in resp.data.tweet_infos
        ]

    def react_get_tweet_full_context(self, tweet_id: str):
        resp = twitter_v2.get_full_context_by_tweet_id(tweet_id=tweet_id)

        if resp.is_error():
            return resp.error

        tweets = [x.tweet_object for x in resp.data.tweet_infos]
        if tweets == []:
            return "Tweet context not found"

        return [_tweet_obj_to_observation(x) for x in tweets]

    def get_following_by_username(self, username: str):
        resp = twitter_v2.get_following_by_username(username=username)

        if resp.is_error():
            return resp.error

        if resp.data.usernames == []:
            return "You are not following any user"

        return resp.data.usernames

    def search_users(self, query: str):
        resp = twitter_v2.search_users(query=query)

        if resp.is_error():
            return resp.error

        if resp.data.users == []:
            return "No user found"

        return [x.to_dict() for x in resp.data.users]

    def get_popular_following_feed(self):
        resp = twitter_v2.get_popular_following_feed(auth=self.auth)

        if resp.is_error():
            return resp.error

        if resp.data.tweets == []:
            return "You are not following any user. Please use another toolcall to retrieve tweet."

        return [_tweet_obj_to_observation(x) for x in resp.data.tweets]

    def search_recent_tweets(self, query: str):
        resp = twitter_v2.search_recent_tweets(
            query=query, limit_observation=10
        )

        if resp.is_error():
            return resp.error

        metadata = {
            "tool_call": "search_recent_tweets",
            "search_query": resp.data.optimized_query,
        }

        if resp.data.tweets == []:
            return "No recent tweets found", metadata

        return [
            _tweet_obj_to_observation(x) for x in resp.data.tweets
        ], metadata

    def get_recent_replies(self):
        resp = twitter_v2.get_own_recent_tweets(
            self.auth, type_whitelist=[TweetType.REPLY]
        )

        if resp.is_error():
            return resp.error

        metadata = {
            "tool_name": "get_recent_replies",
            "params": {},
            "metadata": {
                "search_start": resp.data.search_start.isoformat(),
                "search_end": resp.data.search_end.isoformat(),
                "tweet_count": len(resp.data.tweets),
            },
        }

        return [
            _tweet_obj_to_observation(x) for x in resp.data.tweets
        ], metadata

    def buy(self, symbol: str, amount: float):
        res = trading.buy(
            self.auth.chain_id,
            self.auth.agent_contract_id,
            symbol,
            amount,
            self.auth.ref_id,
        )

        notify_trading_action(
            "buy",
            {
                "symbol": symbol,
                "amount": amount,
                "result": res,
            },
            self.auth.twitter_username,
            self.auth.ref_id,
            self.auth.request_id,
        )

        return res

    def sell(self, symbol: str, amount: float):
        res = trading.sell(
            self.auth.chain_id,
            self.auth.agent_contract_id,
            symbol,
            amount,
            self.auth.ref_id,
        )

        notify_trading_action(
            "sell",
            {
                "symbol": symbol,
                "amount": amount,
                "result": res,
            },
            self.auth.twitter_username,
            self.auth.ref_id,
            self.auth.request_id,
        )

        return res

    def get_wallet_balance(self):
        return trading.get_wallet_balance(
            self.auth.chain_id, self.auth.agent_contract_id
        )

    def get_token_prices(self):
        return trading.get_token_price()

    def research_about_topic(self, topic):
        return bing_search.search_from_bing(
            topic, top_k=10, task_name=f"{self.auth.task}:{self.auth.toolset}"
        )

    def tool_list(self) -> List[ToolDef]:
        # Twitter API get tools
        resp = [
            ToolDef(
                name="get_user_info_by_username",
                description="Get info of a single user by their username, returns user info of a user",
                params=[
                    ToolParam(name="username", dtype=ToolParamDtype.STRING)
                ],
                executor=self.get_user_info_by_username,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_tweets_by_username",
                description="returns a list of most recent tweets by a specified username",
                params=[
                    ToolParam(name="username", dtype=ToolParamDtype.STRING)
                ],
                executor=self.get_tweets_by_username,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_following_by_username",
                description="Get the list of twitter user that a user follows, returns a list of username",
                params=[
                    ToolParam(name="username", dtype=ToolParamDtype.STRING)
                ],
                executor=self.get_following_by_username,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_recent_mentioned_tweets_by_username",
                description="returns a list of most recent tweets mentioning a specified username",
                params=[
                    ToolParam(name="username", dtype=ToolParamDtype.STRING)
                ],
                executor=self.get_recent_mentioned_tweets_by_username,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="search_users",
                description="search users with one topic keyword, return a list of users",
                params=[ToolParam(name="query", dtype=ToolParamDtype.STRING)],
                executor=self.search_users,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="search_recent_tweets",
                description="search recent tweets by 14-15 topic keywords seperated by OR, separated by spaces",
                params=[ToolParam(name="query", dtype=ToolParamDtype.STRING)],
                executor=self.search_recent_tweets,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_popular_following_feed",
                description="search recent tweets from the most popular users that you are following",
                params=[],
                executor=self.get_popular_following_feed,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_tweet_full_context",
                description="Get full context for a tweet by tweet_id, returning a list of all the ancestor tweets of the given tweet.",
                params=[
                    ToolParam(name="tweet_id", dtype=ToolParamDtype.STRING)
                ],
                executor=self.react_get_tweet_full_context,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_your_own_tweets",
                description="returns a list of your own tweets",
                params=[],
                executor=self.get_your_own_tweets,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_recent_posts",
                description="returns a list of your recent posts that is not already inscribed",
                params=[],
                executor=self.get_recent_posts,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_recent_replies",
                description="returns a list of your recent replies that is not already inscribed",
                params=[],
                executor=self.get_recent_replies,
                label=ToolLabel.QUERY,
            ),
            ToolDef(
                name="get_tweets_by_username_selfblock",
                description="returns a list of most recent tweets by a specified username. Don't call it with your own username.",
                params=[
                    ToolParam(name="username", dtype=ToolParamDtype.STRING)
                ],
                executor=self.get_tweets_by_username_selfblock,
                label=ToolLabel.QUERY,
            ),
        ]

        # Twitter API action tools
        resp += [
            ToolDef(
                name="tweet",
                description="Post a tweet",
                params=[
                    ToolParam(name="content", dtype=ToolParamDtype.STRING)
                ],
                executor=self.tweet_with_enhancement,
                label=ToolLabel.ACTION,
            ),
            ToolDef(
                name="reply",
                description="Post a reply to a tweet, specified by the tweet id",
                params=[
                    ToolParam(name="tweet_id", dtype=ToolParamDtype.STRING),
                    ToolParam(
                        name="reply_content", dtype=ToolParamDtype.STRING
                    ),
                ],
                executor=self.reply,
                label=ToolLabel.ACTION,
            ),
            ToolDef(
                name="reply_by_tweet_id",
                description="Auto reply a tweet. Only specify the tweet id when using this tool (as the reply will be automatically generated).",
                params=[
                    ToolParam(name="tweet_id", dtype=ToolParamDtype.STRING)
                ],
                executor=self.reply_by_tweet_id,
                label=ToolLabel.ACTION,
            ),
            ToolDef(
                name="follow",
                description="Follow a Twitter user that I am not following",
                params=[
                    ToolParam(
                        name="target_username", dtype=ToolParamDtype.STRING
                    )
                ],
                executor=self.follow,
                label=ToolLabel.ACTION,
            ),
            ToolDef(
                name="quote_tweet",
                description="Quote a tweet, specified by the tweet id",
                params=[
                    ToolParam(name="tweet_id", dtype=ToolParamDtype.STRING),
                    ToolParam(name="comment", dtype=ToolParamDtype.STRING),
                ],
                executor=self.quote_tweet,
                label=ToolLabel.ACTION,
            ),
            ToolDef(
                name="create_token",
                description="Create a new token on pumpfun and announce it on Twitter",
                params=[
                    ToolParam(name="name", dtype=ToolParamDtype.STRING),
                    ToolParam(name="symbol", dtype=ToolParamDtype.STRING),
                    ToolParam(name="description", dtype=ToolParamDtype.STRING),
                    ToolParam(
                        name="announcement_content",
                        dtype=ToolParamDtype.STRING,
                    ),
                ],
                executor=self.create_token,
                label=ToolLabel.ACTION,
            ),
        ]

        # Trading tools
        _tradable_symbols = trading.get_tradable_symbols()
        _tradable_symbols_str = ", ".join(["$" + e for e in _tradable_symbols])

        if len(_tradable_symbols) > 0:
            resp += [
                ToolDef(
                    name="buy",
                    description="Use SOL to buy token.",
                    params=[
                        ToolParam(name="symbol", dtype=ToolParamDtype.STRING),
                        ToolParam(
                            name="sol_amount", dtype=ToolParamDtype.NUMBER
                        ),
                    ],
                    executor=self.buy,
                    # allow_multiple=True,
                    label=ToolLabel.ACTION,
                ),
                ToolDef(
                    name="sell",
                    description="Sell an amount of token and get SOL back.",
                    params=[
                        ToolParam(name="symbol", dtype=ToolParamDtype.STRING),
                        ToolParam(
                            name="token_amount", dtype=ToolParamDtype.NUMBER
                        ),
                    ],
                    executor=self.sell,
                    # allow_multiple=True,
                    label=ToolLabel.ACTION,
                ),
                ToolDef(
                    name="get_wallet_balance",
                    description="Get the wallet balance",
                    params=[],
                    executor=self.get_wallet_balance,
                    label=ToolLabel.QUERY,
                ),
                ToolDef(
                    name="get_token_prices",
                    description=f"Get the current price of tradable tokens ({_tradable_symbols_str})",
                    params=[],
                    executor=self.get_token_prices,
                    label=ToolLabel.QUERY,
                ),
            ]

        # Bitcoin Ordinal inscribing tools
        resp += [
            ToolDef(
                name="inscribe_post_by_id",
                description="Inscribe a post to Bitcoin Ordinal by its id and set a selling price for it. Include your reasoning for the selected reply and selling price.",
                params=[
                    ToolParam(name="tweet_id", dtype=ToolParamDtype.STRING),
                    ToolParam(name="price", dtype=ToolParamDtype.STRING),
                    ToolParam(name="reason", dtype=ToolParamDtype.STRING),
                ],
                executor=self.inscribe_post_by_id,
                label=ToolLabel.ACTION,
            ),
            ToolDef(
                name="inscribe_reply_by_id",
                description="Inscribe a reply to Bitcoin Ordinal by its id and set a selling price for it. Include your reasoning for the selected reply and selling price.",
                params=[
                    ToolParam(name="tweet_id", dtype=ToolParamDtype.STRING),
                    ToolParam(name="price", dtype=ToolParamDtype.STRING),
                    ToolParam(name="reason", dtype=ToolParamDtype.STRING),
                ],
                executor=self.inscribe_reply_by_id,
                label=ToolLabel.ACTION,
            ),
        ]

        # Research tool
        resp += [
            ToolDef(
                name="research_about_topic",
                description=f"Research more information about a topic. Always call this toolcall before posting a tweet.",
                params=[
                    ToolParam(name="topic", dtype=ToolParamDtype.STRING),
                ],
                executor=self.research_about_topic,
                label=ToolLabel.QUERY,
            ),
        ]

        for i in range(len(resp)):
            resp[i].executor = sync2async(resp[i].executor)

        return resp

    def get_tools_by_toolset(self, toolset: ToolSet) -> List[ToolDef]:

        tools = self.tool_list()
        targets = set(_map.get(toolset, []))
        res = []

        for tool in tools:
            if tool.name in targets:
                res.append(tool)

        return res

    def get_tools(self, toolset=None):
        return self.get_tools_by_toolset(toolset or ToolSet.DEFAULT)
