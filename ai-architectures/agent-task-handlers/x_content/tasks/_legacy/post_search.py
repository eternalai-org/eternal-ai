import json
import logging
from x_content.constants import MissionChainState, ModelName
from x_content.tasks.utils import get_system_prompt_with_random_example_tweets
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.llm_tasks import detect_included_excluded_items
from x_content.wrappers.magic import sync2async
from x_content.wrappers.postprocess import post_process_tweet
from x_content.models import ReasoningLog

# TODO: change this
from x_content.legacy_services.post import brainstorm_post_service

from ..base import MultiStepTaskBase
from x_content.llm.base import OnchainInferResult
from ..utils import a_move_state, create_twitter_auth_from_reasoning_log

from json_repair import repair_json

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

NUM_OF_TWEETS_TO_POST = 1
MINIMUM_POST_LENGTH = 32

from ..utils import a_move_state


async def render_rewrite_tweet_conversation(log: ReasoningLog, content: str):
    if log.model == ModelName.INTELLECT_10B:
        prompt_to_use = "Write a tweet in 256 characters using your specific style. Content to follow:\n\n{}".format(
            content
        )
    else:
        prompt_to_use = """
Craft a tweet based on the provided content that aligns with your personality. 
- No introduction. 
- No gifs. 
- No hashtags. 
- No emojis.

Return the response as a stringified JSON with the key "tweet".

Example JSON Response:
{{
    "tweet": "Your tweet here"
}}

Content to follow: {}
""".format(
            content
        )

    conversational_chat = []
    conversational_chat.append(
        {
            "role": "system",
            "content": await sync2async(
                get_system_prompt_with_random_example_tweets
            )(log),
        }
    )

    conversational_chat.append(
        {
            "role": "user",
            "content": prompt_to_use,
        }
    )

    return conversational_chat


class PostSearchTask(MultiStepTaskBase):

    async def process_task(self, log: ReasoningLog):
        if log.state == MissionChainState.NEW:
            log.execute_info = {
                "news_and_knowledge": [],
                "past_tweets": [],
                "task_result": [],
                "debug_data": [],
                "processing_idx": 0,
            }

            tweet, debug_data = await self.get_tweets(log)
            tweet = await sync2async(post_process_tweet)(tweet)
            log.execute_info["debug_data"].append(debug_data)

            if len(tweet) == 0:
                return await a_move_state(
                    log, MissionChainState.ERROR, "No tweet generated"
                )

            logger.info(f"[{log.id}] Generated base tweet: {tweet}")
            conversation_thread = await self.generate_conversation_thread(
                log, tweet
            )

            log.execute_info["conversation"] = [conversation_thread]
            return await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )

        if log.state == MissionChainState.RUNNING:
            infer_result: OnchainInferResult = await self.llm.agenerate(
                log.execute_info["conversation"][-1], temperature=0.01
            )
            result = infer_result.generations[0].message.content

            if log.model != ModelName.INTELLECT_10B:
                try:
                    result = repair_json(result, return_objects=True)
                    result = result["tweet"]
                except Exception as e:
                    return await a_move_state(
                        log,
                        MissionChainState.ERROR,
                        f"Invalid LLM response '{result}'",
                    )

            system_prompt = log.agent_meta_data.persona
            user_prompt = await self.get_user_prompt(log)
            try:
                include_exclude_info = await sync2async(
                    detect_included_excluded_items
                )(system_prompt, user_prompt)
            except Exception as err:
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Detecting included/excluded items failed: '{err}'",
                )

            postprocessed_tweet = await sync2async(post_process_tweet)(
                result,
                keep_emojis=not include_exclude_info.emojis.excluded,
                keep_mentions=not include_exclude_info.mentions.excluded,
                keep_urls=include_exclude_info.urls.included,
                keep_hashtags=include_exclude_info.hashtags.included,
            )

            logger.info(
                f"[rewrite_tweet] Tweet postprocessed: old={result}, new={result}, include_exclude_data={json.dumps(include_exclude_info.model_dump())}"
            )

            if len(postprocessed_tweet) >= MINIMUM_POST_LENGTH:
                await sync2async(twitter_v2.tweet)(
                    auth=create_twitter_auth_from_reasoning_log(log),
                    content=postprocessed_tweet,
                    tx_hash=infer_result.tx_hash,
                )

            log.execute_info["task_result"].append(
                {
                    "content": postprocessed_tweet,
                    "tx_hash": infer_result.tx_hash,
                }
            )

            return await a_move_state(
                log, MissionChainState.DONE, "Final answer found"
            )

        return log

    async def get_user_prompt(self, log: ReasoningLog):
        return log.prompt

    async def get_tweets(self, log: ReasoningLog):
        system_prompt = log.agent_meta_data.persona
        user_prompt = await self.get_user_prompt(log)

        tweet, debug_data = await brainstorm_post_service.generate_content(
            system_prompt,
            user_prompt,
            self.kn_base,
            task_name=log.task,
        )

        return tweet, debug_data

    async def generate_conversation_thread(
        self, log: ReasoningLog, tweet: str
    ):
        conversation_thread = await render_rewrite_tweet_conversation(
            log, tweet
        )

        return conversation_thread
