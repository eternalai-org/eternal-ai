import json
import logging
import random
from dotenv import load_dotenv
from typing import List

# TODO: bad import, change this
from x_content.legacy_services.utils import PROMPT_TEMPLATE
from x_content.legacy_services.utils import create_llm

from x_content.wrappers.api import twitter_v2
from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.log_decorators import log_function_call
from x_content.wrappers.magic import sync2async
from x_content.wrappers.rag_search import search_from_db, get_random_from_collections
from x_content.wrappers.bing_search import search_from_bing
from x_content import constants as const
from json_repair import repair_json


logger = logging.getLogger(__name__)

load_dotenv()
RETRY = 2


class BrainstormTweetService:

    def __init__(self):
        self.hermes = create_llm(
            base_url=const.SELF_HOSTED_HERMES_70B_URL + "/v1",
            model_id=const.SELF_HOSTED_HERMES_70B_MODEL_IDENTITY,
            api_key=const.SELF_HOSTED_HERMES_70B_API_KEY,
            temperature=1,
        )

        self.llama = create_llm(
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            model_id=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            temperature=1,
        )

    async def _generate_topic_from_question(
        self,
        infer_system_prompt: str,
        question: str,
        context_tweets: List[str] = [],
        retry: int = RETRY,
    ) -> str:
        tweets_str = "\n".join([f"- {x}" for x in context_tweets])
        for attempt in range(retry):
            try:
                system_prompt = f"""
You are a highly advanced language model capable of analyzing both the system prompt, a list of tweets and the user's query to generate a precise and relevant search topic for retrieving news and knowledge base information.

Your task is to produce ONE FOCUSED AND RELEVANT search topic based on the analysis of the system prompt, tweets and user query, specifically for news and knowledge searches.

The output must be a stringified JSON object with the key "topic," where the value is the search topic derived from your analysis of both the system and user prompts.

Be sure to keep the generated topic specific, concise, and directly related to the user's request. The result must not include any extraneous text, explanations, or details outside the JSON structure.

Example output:
{{ "topic": "Your topic" }}
"""

                user_prompt = f"""
System prompt:
{infer_system_prompt}

Tweets:
{tweets_str}

User query:
{question}
"""

                prompt = await PROMPT_TEMPLATE.ainvoke(
                    {"system_prompt": system_prompt, "question": user_prompt}
                )

                resp = await self.hermes.ainvoke(prompt)
                parsed_content = repair_json(resp.content, return_objects=True)
                return parsed_content["topic"]
            except Exception as err:
                if attempt + 1 == retry:
                    try:
                        resp = await self.llama.ainvoke(prompt)
                        parsed_content = repair_json(resp.content, True)
                        return parsed_content["topic"]
                    except Exception as llama_err:
                        return ""

    @log_function_call
    async def generate_content(
        self,
        infer_system_prompt: str,
        infer_user_prompt: str,
        kn_base: KnowledgeBase,
        top_k: int = 10,
        retry: int = RETRY,
        task_name: str = "N/A",
    ) -> str:
        context_tweets, err = await sync2async(get_random_from_collections)(
            kn_base.get_kn_ids(), n=10
        )

        debug_data = {"context_tweets": context_tweets}

        for attempt in range(retry):
            try:
                query = await self._generate_topic_from_question(
                    infer_system_prompt, infer_user_prompt, context_tweets
                )
                debug_data["query"] = query

                structured_information = {}
                if query != "":
                    knowledge = await search_from_db(
                        kn_base, query, top_k=5, threshold=0.85
                    )

                    news = await sync2async(search_from_bing)(
                        query, top_k=10, task_name=task_name
                    )

                    random_news = random.sample(news, min(1, len(news)))
                    structured_information["knowledge"] = knowledge
                    structured_information["news"] = random_news
                debug_data["structured_information"] = structured_information

                system_prompt = f"""{infer_system_prompt}

Your task is to craft a highly engaging tweet based on the provided information, ensuring it is relevant, timely, and captivating based on the user prompt and provided information.

The provided information includes:
    - **news**: Recent news retrieved from Bing search, which should be prioritized for timeliness and relevance.  
    - **knowledge**: Foundational information retrieved from the database for additional context.

The tweet must:
- Be concise, clear, and captivating, adhering to character limits (280 characters or fewer).
- Informatively and engagingly address the user query while prioritizing relevance and impact.
- Utilize compelling, precise language to grab attention, avoid redundancy, and maintain a professional yet approachable tone.
- Highlight timeliness and relevance from the news while seamlessly integrating relevant knowledge for context.

Output the response as a stringified JSON object with the key "tweet" containing the generated tweet.

Generate only one tweet per query, ensuring it is polished, impactful, and effective in delivering the intended message.
"""

                user_prompt = f"""
User prompt:
{infer_user_prompt}

Provided information:
{json.dumps(structured_information)}
"""
                debug_data["content_conversation"] = {
                    "system_prompt": system_prompt,
                    "question": user_prompt,
                }

                prompt = await PROMPT_TEMPLATE.ainvoke(
                    {"system_prompt": system_prompt, "question": user_prompt}
                )

                resp = await self.hermes.ainvoke(prompt)
                parsed_content = repair_json(resp.content, return_objects=True)
                debug_data["content"] = parsed_content["tweet"]
                return parsed_content["tweet"], debug_data
            except Exception as err:
                logger.error(
                    f"[tweets_rag] Attempt {attempt + 1} failed with hermes: {err}"
                )
                if attempt + 1 == retry:
                    try:
                        resp = await self.llama.ainvoke(prompt)
                        parsed_content = repair_json(
                            resp.content, return_objects=True
                        )
                        debug_data["content"] = parsed_content["tweet"]
                        return parsed_content["tweet"], debug_data
                    except Exception as llama_err:
                        logger.error(
                            f"[tweets_rag] Attempt {attempt + 1} failed with llama: {llama_err}"
                        )
                        return "", debug_data


brainstorm_post_service = BrainstormTweetService()
