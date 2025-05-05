import json
import logging
from typing import List

from x_content.constants.main import ModelName
from x_content.wrappers.api.twitter_v2.models.objects import (
    StructuredInformation,
    TweetObject,
)
from x_content.wrappers.llm_tasks import choose_suitable_language
from x_content.wrappers.magic import sync2async

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

ENHANCE_TWEET_PROMPT_TEMPLATE = """Analyze the provided base tweet and rephrase it using your distinct personality and writing style, ensuring that the main message remains unchanged.

Return the revised tweet as a JSON string with the key "tweet".

Example JSON Response:
{{
    "tweet": "Your rewritten tweet here"
}}

Base tweet: {base_tweet}
"""


def get_enhance_tweet_conversation(
    system_prompt: str, content: str, example_tweets: List[str] = []
):
    if example_tweets is not None and example_tweets != []:
        tweets_str = "\n".join([f"- {x}" for x in example_tweets])
        system_prompt = f"""{system_prompt}

Here are example tweets written in the defined style:
{tweets_str}
"""

    prompt_to_use = ENHANCE_TWEET_PROMPT_TEMPLATE.format(base_tweet=content)

    messages = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": prompt_to_use,
        },
    ]

    return messages


REPLY_TWEET_PROMPT_TEMPLATE = """Provide a single message to join the following conversation. Keep it concise (under 128 chars), No thread, No hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.

Provide the output in JSON format, e.g.:
{{
  "tweet": "<response tweet in the {language} language>"
}}

The conversation:
{chat_history}

Relevant information:
{structured_info}
"""


async def get_reply_tweet_conversation(
    system_prompt: str,
    task_prompt: str,
    tweets: List[TweetObject],
    structured_info: StructuredInformation,
):
    chat_history = [
        {
            "user": x.twitter_username,
            "message": x.full_text,
        }
        for x in tweets
    ]

    try:
        language = await sync2async(choose_suitable_language)(chat_history)
    except Exception as err:
        logger.info(
            f"[get_reply_tweet_conversation] Error detecting language: {err}"
        )
        language = "en"

    content_list = structured_info.news + structured_info.knowledge

    user_prompt = REPLY_TWEET_PROMPT_TEMPLATE.format(
        language=language,
        task_prompt=task_prompt,
        chat_history=json.dumps(chat_history),
        structured_info="\n".join([f"- {x}" for x in content_list]),
    )

    conversational_chat = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]

    return conversational_chat


REPLY_GAME_PROMPT_TEMPLATE = """Act as an expert in Twitter chat history analysis. Carefully review the provided Twitter chat history, a list of tweets in JSON format. Each tweet object contains a "user" field with the Twitter handle of the author, and a "message" field with the text content of the tweet. The tweets are ordered chronologically from oldest to most recent. The final tweet in the sequence contains a question that needs to be answered.

In addition to the chat history, you will also be provided with a bulleted list of relevant information that may be useful in formulating your response to the question. 

Provide your final answer to the question in valid JSON format. Here is the expected structure of your JSON response:
{{
  "thoughts": [
    "Your thought 1",
    "Your thought 2", 
    ...
  ],
  "answer": "Your final answer to the question"
}}

Chat History:
{chat_history}

Relevant Information:
{structured_info}
"""


async def get_reply_game_conversation(
    system_prompt: str,
    task_prompt: str,
    tweets: List[TweetObject],
    structured_info: StructuredInformation,
):
    chat_history = [
        {
            "user": x.twitter_username,
            "message": x.full_text,
        }
        for x in tweets
    ]

    try:
        language = await sync2async(choose_suitable_language)(chat_history)
    except Exception as err:
        logger.info(
            f"[get_reply_tweet_conversation] Error detecting language: {err}"
        )
        language = "en"

    content_list = structured_info.news + structured_info.knowledge

    user_prompt = REPLY_GAME_PROMPT_TEMPLATE.format(
        language=language,
        # task_prompt=task_prompt,
        chat_history=json.dumps(chat_history),
        structured_info="\n".join([f"- {x}" for x in content_list]),
    )

    conversational_chat = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]

    return conversational_chat

import re

def parse_deepseek_r1_result(content: str):
    result = {}

    pat = re.compile(
        r"[<think>]?(.*?)</think>(.*)", 
        re.DOTALL | re.MULTILINE | re.IGNORECASE
    )
    
    match = pat.match(content)

    if match is not None:
        result["think"] = match.group(1).strip()
        result["answer"] = match.group(2).strip()
    else:
        result["answer"] = content.strip()

    return result


def get_llm_result_by_model_name(content: str, model_name: str):
    if model_name == ModelName.DEEPSEEK_R1:
        return parse_deepseek_r1_result(content)["answer"]

    return content
