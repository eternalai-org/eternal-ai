from typing import List

from json_repair import repair_json

from x_content.wrappers.conversation import get_enhance_tweet_conversation, get_llm_result_by_model_name
from x_content.wrappers.postprocess import postprocess_tweet_by_prompts
from x_content.wrappers.magic import get_agent_llm_first_interval, retry, sync2async
from x_content.tasks.utils import (
    a_move_state,
    create_twitter_auth_from_reasoning_log,
    get_system_prompt_with_random_example_tweets,
)
from x_content.wrappers.api.twitter_v2.models.response import (
    GenerateActionDto,
    Response,
    TweetsDto,
)
from x_content.tasks.base import MultiStepTaskBase
from x_content.models import ReasoningLog, MissionChainState
from x_content.wrappers.bing_search import search_from_bing
from x_content.wrappers.twin_agent import get_random_example_tweets
from x_content.wrappers.api import twitter_v2
from x_content.llm.base import OnchainInferResult
import json
import random


def load_mission_configuration(log: ReasoningLog) -> dict:
    try:
        mission_configuration = json.loads(log.prompt)
    except:
        mission_configuration = {
            "prompt": log.prompt,
            "X": False,
            "bing": False,
            "topics": "",
        }
    return mission_configuration


class PostV2(MultiStepTaskBase):
    resumable = True

    POST_TOPICS_NEWS_PROMPT_TEMPLATE = """
{task_definition}

Content to use:
{news_to_use}
"""

    POST_TOPICS_PROMPT_TEMPLATE = """
{task_definition}

Topic list: {topics}
"""

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:

        if log.state == MissionChainState.NEW:
            mission_configuration = load_mission_configuration(log)

            missing_fields = [
                k
                for k in ["bing", "X", "topics", "prompt"]
                if k not in mission_configuration
            ]

            if len(missing_fields) > 0:
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Missing fields: {missing_fields} in prompt",
                )

            task_prompt = mission_configuration["prompt"]
            use_x_news = mission_configuration["X"]
            use_bing_news = mission_configuration["bing"]

            is_using_news = use_x_news or use_bing_news

            topics = [
                e.strip()
                for e in mission_configuration["topics"].split(",")
                if e.strip() != ""
            ]

            if is_using_news:
                if len(topics) == 0:
                    return await a_move_state(
                        log,
                        MissionChainState.ERROR,
                        "When news is used, topic list cannot be empty",
                    )
                log = await self._process_post_with_news(
                    log, task_prompt, topics, use_x_news, use_bing_news
                )
            else:
                if len(topics) > 0:
                    log = await self._process_post_with_topics(
                        log, task_prompt, topics
                    )
                else:
                    log = await self._process_post(log, task_prompt)

            return await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )

        if log.state == MissionChainState.RUNNING:
            mission_configuration = load_mission_configuration(log)
            task_prompt = mission_configuration["prompt"]

            async def get_base_tweet():
                base_tweet_conversation = log.execute_info["conversation"][-1]
                result: OnchainInferResult = await self.llm.agenerate(
                    base_tweet_conversation, temperature=0.7
                )
                content = result.generations[0].message.content
                base_tweet = get_llm_result_by_model_name(content, log.model)
                return base_tweet, result.tx_hash

            try:
                base_tweet, base_tweet_tx_hash = await retry(
                    get_base_tweet,
                    max_retry=3,
                    first_interval=get_agent_llm_first_interval(),
                    interval_multiply=2,
                )()
            except Exception as err:
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Failed to generate base tweet: {err}",
                )

            log.execute_info["task_result"] = [
                {
                    "base_tweet": base_tweet,
                    "base_tweet_tx_hash": base_tweet_tx_hash,
                }
            ]

            log.execute_info["conversation"].append(
                get_enhance_tweet_conversation(
                    system_prompt=log.agent_meta_data.persona,
                    content=base_tweet,
                    example_tweets=await sync2async(get_random_example_tweets)(
                        log.meta_data.knowledge_base_id
                    ),
                )
            )

            async def get_enhanced_tweet():
                enhance_tweet_conversation = log.execute_info["conversation"][
                    -1
                ]
                result: OnchainInferResult = await self.llm.agenerate(
                    enhance_tweet_conversation, temperature=0.01
                )
                content = result.generations[0].message.content
                content = get_llm_result_by_model_name(content, log.model)
                data = repair_json(content, return_objects=True)
                return data["tweet"], result.tx_hash

            try:
                enhanced_tweet, enhanced_tweet_tx_hash = await retry(
                    get_enhanced_tweet,
                    max_retry=3,
                    first_interval=get_agent_llm_first_interval(),
                    interval_multiply=2,
                )()
            except Exception as err:
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Failed to enhance the tweet: {err}",
                )

            log.execute_info["task_result"][-1].update(
                {
                    "enhanced_tweet": enhanced_tweet,
                    "enhanced_tweet_tx_hash": enhanced_tweet_tx_hash,
                }
            )

            try:
                postprocessed_tweet = await sync2async(
                    postprocess_tweet_by_prompts
                )(
                    system_prompt=log.agent_meta_data.persona,
                    task_prompt=task_prompt,
                    tweet=enhanced_tweet,
                )
            except Exception as err:
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Failed to postprocess the tweet: {err}",
                )

            resp: Response[GenerateActionDto] = await sync2async(
                twitter_v2.tweet
            )(
                auth=create_twitter_auth_from_reasoning_log(log),
                content=postprocessed_tweet,
                tx_hash=base_tweet_tx_hash,
            )

            if resp.is_error():
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Failed to perform tweet action: {resp.error}",
                )

            if not resp.data.success:
                return await a_move_state(
                    log, MissionChainState.ERROR, f"Failed to schedule tweet"
                )

            log.execute_info["task_result"][-1].update(
                {
                    "postprocessed_tweet": postprocessed_tweet,
                }
            )

            return await a_move_state(
                log, MissionChainState.DONE, "Task completed"
            )

        return log

    async def _process_post_with_news(
        self,
        log: ReasoningLog,
        prompt: str,
        topics: List[str],
        use_x_news: bool,
        use_bing_news: bool,
    ):
        news_source = []
        random.shuffle(topics)

        for today_topic in topics:
            x_news = []
            if use_x_news:
                search_resp: Response[TweetsDto] = await sync2async(
                    twitter_v2.search_twitter_news
                )(today_topic, limit_api_results=50, use_raw=True)

                tweets = (
                    search_resp.data.tweets
                    if not search_resp.is_error()
                    else []
                )
                x_news = [e.full_text for e in tweets]
                news_source.extend(x_news)

            bing_news = []
            if use_bing_news:
                bing_news = await sync2async(search_from_bing)(
                    today_topic,
                    top_k=5,
                    task_name=log.task,
                )
                news_source.extend(bing_news)

            if len(news_source) > 0:
                structured_news_source = {
                    "topic": today_topic,
                    "x_news": x_news,
                    "bing_news": bing_news,
                }
                break

        if len(news_source) == 0:
            return await a_move_state(
                log,
                MissionChainState.ERROR,
                f"No news found for any of provided topics {topics}",
            )

        messages = [
            {
                "role": "system",
                "content": await sync2async(
                    get_system_prompt_with_random_example_tweets
                )(log),
            },
            {
                "role": "user",
                "content": self.POST_TOPICS_NEWS_PROMPT_TEMPLATE.format(
                    task_definition=prompt,
                    news_to_use="\n".join("- " + e for e in news_source),
                ),
            },
        ]

        log.execute_info = {
            "news_source": structured_news_source,
            "conversation": [messages],
        }

        return log

    async def _process_post_with_topics(
        self, log: ReasoningLog, prompt: str, topics: List[str]
    ):
        messages = [
            {
                "role": "system",
                "content": await sync2async(
                    get_system_prompt_with_random_example_tweets
                )(log),
            },
            {
                "role": "user",
                "content": self.POST_TOPICS_PROMPT_TEMPLATE.format(
                    task_definition=prompt, topics=", ".join(topics)
                ),
            },
        ]

        log.execute_info = {
            "conversation": [messages],
        }

        return log

    async def _process_post(self, log: ReasoningLog, prompt: str):
        messages = [
            {
                "role": "system",
                "content": await sync2async(
                    get_system_prompt_with_random_example_tweets
                )(log),
            },
            {"role": "user", "content": prompt},
        ]

        log.execute_info = {
            "conversation": [messages],
        }

        return log
