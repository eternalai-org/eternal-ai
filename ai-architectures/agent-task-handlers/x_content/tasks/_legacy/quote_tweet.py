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
from x_content.models import ReasoningLog

from x_content.tasks.base import MultiStepTaskBase
from x_content.tasks.utils import a_move_state, create_twitter_auth_from_reasoning_log
from x_content.llm.base import OnchainInferResult
import asyncio

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

from x_content.wrappers.magic import sync2async


class QuoteTweetTask(MultiStepTaskBase):
    resumable = True

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        quote_tweet_username = log.meta_data.params.get("quote_username")

        if quote_tweet_username is None or len(quote_tweet_username) == 0:
            return await a_move_state(
                log, MissionChainState.ERROR, "Invalid username to quote tweet"
            )

        if log.state == MissionChainState.NEW:
            resp: Response[TweetInfosDto] = await sync2async(
                twitter_v2.get_tweets_by_username_v2
            )(
                username=quote_tweet_username,
                num_tweets=5,
            )

            if resp.is_error():
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Error retrieving quote tweets: {resp.error}",
                )

            tweets = [x.to_dict() for x in resp.data.tweet_infos]

            log.execute_info = {
                "tweets": tweets,
                "task_result": [],
                "processing_idx": 0,
                "conversation": [],
            }

            for idx, tweet_info in enumerate(log.execute_info["tweets"]):
                context_resp: Response[
                    ExtendedTweetInfosDto
                ] = await sync2async(twitter_v2.get_full_context_by_tweet_id)(
                    tweet_info["tweet_object"]["tweet_id"],
                )

                if context_resp.is_error():
                    full_context = []
                else:
                    full_context = [
                        x.tweet_object for x in context_resp.data.tweet_infos
                    ]

                info_resp = await twitter_v2.get_relevent_information_v2(
                    self.kn_base,
                    tweets=full_context,
                    task_name=log.task,
                )

                knowledge_v2 = (
                    StructuredInformation(knowledge=[], news=[])
                    if info_resp.is_error()
                    else info_resp.data.structured_information
                )

                log.execute_info["tweets"][idx].update(
                    full_context=[x.to_dict() for x in full_context],
                    knowledge=knowledge_v2,
                    knowledge_v2=knowledge_v2,
                )

                log.execute_info["conversation"].append(
                    await get_reply_tweet_conversation(
                        system_prompt=log.agent_meta_data.persona,
                        task_prompt=log.prompt,
                        tweets=full_context,
                        structured_info=info_resp.data.structured_information,
                    )
                )

            return await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )

        if log.state == MissionChainState.RUNNING:
            totals = len(log.execute_info["conversation"])

            futures = [
                asyncio.ensure_future(
                    self.llm.agenerate(conversation, temperature=0.7)
                )
                for conversation in log.execute_info["conversation"]
            ]

            # TODO: rewrite this
            for i, infer in enumerate(
                await asyncio.gather(*futures, return_exceptions=True)
            ):
                if isinstance(infer, Exception):
                    logger.info(
                        f"[{log.id}] Error while processing index {i} (out of {totals}): {infer} (inference fails)."
                    )
                    continue

                infer: OnchainInferResult
                result = await sync2async(post_process_tweet)(
                    infer.generations[0].message.content.strip('" ')
                )
                liked_tweet = log.execute_info["tweets"][i]["tweet_object"]
                liked_tweet_id = liked_tweet["tweet_id"]

                await sync2async(twitter_v2.quote_tweet)(
                    auth=create_twitter_auth_from_reasoning_log(log),
                    tweet_id=liked_tweet_id,
                    comment=result,
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
                log, MissionChainState.DONE, "Quote tweet task completed"
            )

        return log
