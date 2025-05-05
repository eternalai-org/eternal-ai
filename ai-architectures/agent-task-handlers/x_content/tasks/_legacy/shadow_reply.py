import logging
from x_content.constants import MissionChainState
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.objects import StructuredInformation
from x_content.wrappers.api.twitter_v2.models.response import (
    ExtendedTweetInfosDto,
    Response,
    TweetInfosDto,
)
from x_content.wrappers.conversation import get_reply_tweet_conversation
from x_content.wrappers.postprocess import post_process_tweet

from x_content.tasks.base import MultiStepTaskBase
from x_content.tasks.utils import a_move_state, create_twitter_auth_from_reasoning_log
from x_content.llm.base import OnchainInferResult

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

from x_content.wrappers.magic import sync2async
import asyncio
from x_content.models import ReasoningLog


class ShadowReplyTask(MultiStepTaskBase):
    resumable = True

    async def process_task(self, log: ReasoningLog):

        if log.state == MissionChainState.NEW:
            resp: Response[TweetInfosDto] = await sync2async(
                twitter_v2.get_full_conversation_from_liked_tweets
            )(
                auth=create_twitter_auth_from_reasoning_log(log),
                num_tweets=5,
                ignore_replied_tweets=True,
            )

            if resp.is_error():
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Error retrieving shadow reply tweets: {resp.error}",
                )

            liked_tweets = [x.to_dict() for x in resp.data.tweet_infos]

            log.execute_info = {
                "tweets": liked_tweets,
                "task_result": [],
                "conversation": [],
            }

            for idx in range(len(liked_tweets)):
                tweet_info = liked_tweets[idx]

                context_resp: Response[ExtendedTweetInfosDto] = (
                    await sync2async(twitter_v2.get_full_context_by_tweet_id)(
                        tweet_info["tweet_object"]["tweet_id"]
                    )
                )

                if context_resp.is_error():
                    tweets_context = []
                else:
                    tweets_context = [
                        x.tweet_object for x in context_resp.data.tweet_infos
                    ]

                info_resp = await twitter_v2.get_relevent_information_v2(
                    self.kn_base,
                    tweets=tweets_context,
                    task_name=log.task,
                )

                knowledge_v2 = (
                    StructuredInformation(knowledge=[], news=[])
                    if info_resp.is_error()
                    else info_resp.data.structured_information
                )

                log.execute_info["tweets"][idx].update(
                    full_context=[x.to_dict() for x in tweets_context],
                    knowledge=knowledge_v2,
                    knowledge_v2=knowledge_v2,
                )

                log.execute_info["conversation"].append(
                    await get_reply_tweet_conversation(
                        system_prompt=log.agent_meta_data.persona,
                        task_prompt=log.prompt,
                        tweets=tweets_context,
                        structured_info=knowledge_v2,
                    )
                )

            return await a_move_state(
                log,
                MissionChainState.RUNNING,
                f"Start shadow reply task; found {len(liked_tweets)} liked tweets",
            )

        if log.state == MissionChainState.RUNNING:
            futures = [
                asyncio.ensure_future(
                    self.llm.agenerate(conversation_thread, temperature=0.7)
                )
                for conversation_thread in log.execute_info["conversation"]
            ]

            for idx, infer in enumerate(
                await asyncio.gather(*futures, return_exceptions=True)
            ):
                if isinstance(infer, Exception):
                    logger.info(
                        f"[{log.id}] Error while processing index {idx}: {infer} (inference fails)."
                    )
                    continue

                infer: OnchainInferResult
                result = post_process_tweet(
                    infer.generations[0].message.content
                )
                liked_tweet = log.execute_info["tweets"][idx]["tweet_object"]
                liked_tweet_id = liked_tweet["tweet_id"]

                await sync2async(twitter_v2.shadow_reply)(
                    create_twitter_auth_from_reasoning_log(log),
                    tweet_id=liked_tweet_id,
                    reply_content=result,
                    tx_hash=infer.tx_hash,
                )

                log.execute_info["task_result"].append(
                    {
                        "tweet_id": liked_tweet_id,
                        "reply_content": result,
                        "tx_hash": infer.tx_hash,
                    }
                )

            return await a_move_state(
                log,
                MissionChainState.DONE,
                f"Replied to all {len(futures)} liked tweets",
            )

        return log
