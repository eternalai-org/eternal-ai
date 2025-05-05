import logging

from json_repair import repair_json
from x_content.tasks.utils import create_twitter_auth_from_reasoning_log
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.objects import (
    ExtendedTweetInfo,
    StructuredInformation,
)
from x_content.wrappers.api.twitter_v2.models.response import (
    ExtendedTweetInfosDto,
    Response,
)
from x_content.wrappers.conversation import (
    get_llm_result_by_model_name,
    get_reply_game_conversation,
)
from x_content.tasks.reply_subtask_base import ReplySubtaskBase

from x_content.llm.base import OnchainInferResult

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

MINIMUM_REPLY_LENGTH = 32

from x_content.wrappers.game import GameAPIClient, GameState
from x_content.wrappers.magic import get_agent_llm_first_interval, retry, sync2async


def is_game_created(_tweet_id):
    try:
        game_info, err = GameAPIClient.get_game_info_by_tweet_id(_tweet_id)
        if err:
            raise err
        return game_info.status == GameState.RUNNING
    except Exception as err:
        logger.error(
            f"[is_game_created] Failed to check done status for tweet {_tweet_id}: {err}"
        )
        return False


class ReplyGameSubtask(ReplySubtaskBase):

    async def run(self) -> dict:
        self.tweet_info: ExtendedTweetInfo = await sync2async(
            twitter_v2.get_tweet_with_image_description_appended_to_text
        )(self.tweet_info)

        context_resp: Response[ExtendedTweetInfosDto] = await sync2async(
            twitter_v2.get_full_context_of_tweet
        )(self.tweet_info)

        if context_resp.is_error():
            tweets_context = []
        else:
            tweets_context = [
                x.tweet_object for x in context_resp.data.tweet_infos
            ]

        resp = await twitter_v2.get_relevent_information_v2(
            self.kn_base,
            tweets=tweets_context,
            task_name=self.log.task,
            use_bing_search=False,
        )
        knowledge_v2 = (
            StructuredInformation(knowledge=[], news=[])
            if resp.is_error()
            else resp.data.structured_information
        )

        # mentioned_tweets[idx].update(
        #     full_context=[x.to_dict() for x in tweets_context],
        #     knowledge=knowledge_v2,
        #     knowledge_v2=knowledge_v2
        # )

        base_reply_conversation = await get_reply_game_conversation(
            system_prompt=self.log.agent_meta_data.persona,
            task_prompt=self.log.prompt,
            tweets=tweets_context,
            structured_info=knowledge_v2,
        )

        async def get_base_reply():
            result: OnchainInferResult = await self.llm.agenerate(
                base_reply_conversation, temperature=0.7
            )
            response = result.generations[0].message.content
            response = get_llm_result_by_model_name(response, self.log.model)
            debug_data = {
                "tweets_context": [
                    {"user": x.twitter_username, "message": x.full_text}
                    for x in tweets_context
                ],
                "response": response,
            }
            logger.info(f"[ReplyGameSubtask.get_base_reply] {debug_data}")
            data = repair_json(response, return_objects=True)
            return data["answer"], result.tx_hash

        try:
            base_reply, base_reply_tx_hash = await retry(
                get_base_reply,
                max_retry=3,
                first_interval=get_agent_llm_first_interval(),
                interval_multiply=2,
            )()
        except Exception as err:
            raise Exception(f"Failed to generate base reply: {err}")

        if len(base_reply) >= MINIMUM_REPLY_LENGTH:
            await sync2async(twitter_v2.reply)(
                auth=create_twitter_auth_from_reasoning_log(self.log),
                tweet_id=self.tweet_info.tweet_object.tweet_id,
                reply_content=base_reply,
                tx_hash=base_reply_tx_hash,
            )

        return {
            "tweet_id": self.tweet_info.tweet_object.tweet_id,
            "conversation": [
                base_reply_conversation,
            ],
            "base_reply": base_reply,
            "base_reply_tx_hash": base_reply_tx_hash,
        }
