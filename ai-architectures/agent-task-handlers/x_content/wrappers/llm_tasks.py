import json
import os
import random
from openai import BaseModel
import requests
import logging
from typing import List, Optional

from . import redis_wrapper

from x_content.llm.local import SyncBasedEternalAI
from x_content.wrappers.magic import get_llm_tasks_first_interval, helpful_raise_for_status
from x_content.wrappers.magic import retry
from .log_decorators import log_function_call
from x_content import constants as const
from json_repair import repair_json

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

MAX_TEXT_LENGTH = 10000
MIN_TEXT_LENGTH_TO_SUMMARIZE = 100


def choose_suitable_language(chat_history: list) -> str:
    """
    Determines the most suitable language for the input chat history.

    Args:
        chat_history (list): A list of strings representing the chat conversation.

    Returns:
        str: The ISO 639-1 language code for the detected language. Defaults to 'en' if an error occurs.
    """
    valid_languages = [
        "aa",
        "ab",
        "af",
        "ak",
        "am",
        "ar",
        "an",
        "as",
        "av",
        "ae",
        "ay",
        "az",
        "ba",
        "bm",
        "be",
        "bn",
        "bi",
        "bo",
        "bs",
        "br",
        "bg",
        "ca",
        "cs",
        "ch",
        "ce",
        "cu",
        "cv",
        "kw",
        "co",
        "cr",
        "cy",
        "da",
        "de",
        "dv",
        "dz",
        "el",
        "en",
        "eo",
        "et",
        "eu",
        "ee",
        "fo",
        "fa",
        "fj",
        "fi",
        "fr",
        "fy",
        "ff",
        "gd",
        "ga",
        "gl",
        "gv",
        "gn",
        "gu",
        "ht",
        "ha",
        "sh",
        "he",
        "hz",
        "hi",
        "ho",
        "hr",
        "hu",
        "hy",
        "ig",
        "io",
        "ii",
        "iu",
        "ie",
        "ia",
        "id",
        "ik",
        "is",
        "it",
        "jv",
        "ja",
        "kl",
        "kn",
        "ks",
        "ka",
        "kr",
        "kk",
        "km",
        "ki",
        "rw",
        "ky",
        "kv",
        "kg",
        "ko",
        "kj",
        "ku",
        "lo",
        "la",
        "lv",
        "li",
        "ln",
        "lt",
        "lb",
        "lu",
        "lg",
        "mh",
        "ml",
        "mr",
        "mk",
        "mg",
        "mt",
        "mn",
        "mi",
        "ms",
        "my",
        "na",
        "nv",
        "nr",
        "nd",
        "ng",
        "ne",
        "nl",
        "nn",
        "nb",
        "no",
        "ny",
        "oc",
        "oj",
        "or",
        "om",
        "os",
        "pa",
        "pi",
        "pl",
        "pt",
        "ps",
        "qu",
        "rm",
        "ro",
        "rn",
        "ru",
        "sg",
        "sa",
        "si",
        "sk",
        "sl",
        "se",
        "sm",
        "sn",
        "sd",
        "so",
        "st",
        "es",
        "sq",
        "sc",
        "sr",
        "ss",
        "su",
        "sw",
        "sv",
        "ty",
        "ta",
        "tt",
        "te",
        "tg",
        "tl",
        "th",
        "ti",
        "to",
        "tn",
        "ts",
        "tk",
        "tr",
        "tw",
        "ug",
        "uk",
        "ur",
        "uz",
        "ve",
        "vi",
        "vo",
        "wa",
        "wo",
        "xh",
        "yi",
        "yo",
        "za",
        "zh",
        "zu",
    ]
    system_prompt = """
    You are a neutral and unbiased assistant tasked with determining the most suitable language for tweets.
    Analyze the input text objectively and specify the single most suitable language in which the conversation predominantly occurs.
    Treat all languages equally without favoring any specific language, including English.

    If the conversation contains mixed languages, choose the language that appears most frequently or is most contextually appropriate.
    Always return a single ISO 639-1 language code as the result.

    Provide the final response in a stringified JSON format with the key 'language' and the value as the corresponding ISO 639-1 language code.

    Examples:

    Input:
    @user1: 你好，你最近怎么样？
    @user2: 我很好，谢谢！你呢？

    Output:
    {
        "language": "zh"
    }

    Input:
    @user1: Hola, ¿cómo estás?
    @user2: Estoy bien, ¡gracias! ¿Y tú?

    Output:
    {
        "language": "es"
    }

    Input:
    @user1: Bonjour, ça va?
    @user2: I'm doing well, merci!

    Output:
    {
        "language": "fr"
    }

    Input:
    @user1: Hi, good morning!
    @user2: おはようございます！

    Output:
    {
        "language": "ja"
    }
    """

    chat_history = list(
        map(lambda x: f'@{x["user"]}: {x["message"]}', chat_history)
    )
    chat_history = "\n".join(chat_history)

    conversation = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": chat_history},
    ]

    url = const.SELF_HOSTED_LLAMA_405B_URL + "/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {const.SELF_HOSTED_LLAMA_API_KEY}",
    }

    data = {
        "model": const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
        "messages": conversation,
        "max_tokens": 100,
        "temperature": 0,
    }

    response = requests.post(url, headers=headers, json=data)
    helpful_raise_for_status(response)
    content = response.json()["choices"][0]["message"]["content"]
    language: str = repair_json(content, return_objects=True)["language"]

    if language in valid_languages:
        return language

    return "en"


def generate_retrieval_query(chat_history: list) -> Optional[str]:
    # Always use English because the model only supports English
    system_prompt = """
You are a helpful and intelligent assistant. Your task is to analyze the chat conversation and extract the most suitable query to retrieve relevant information from a database.

Ensure the query is clear, concise, and accurately represents the user's intent.

Your response **must** be in stringified JSON format, with the following structure:
{
  "query": "<extracted_query_string>"
}
"""

    formatted_chat_history = "\n".join(
        f'@{entry["user"]}: {entry["message"]}' for entry in chat_history
    )

    conversation = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": formatted_chat_history},
    ]

    llama_url = const.SELF_HOSTED_LLAMA_405B_URL
    llama_api_key = const.SELF_HOSTED_LLAMA_API_KEY
    llama_model_identity = const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY

    if not all([llama_url, llama_api_key, llama_model_identity]):
        logger.error(
            "[generate_retrieval_query] Missing environment variables"
        )
        return None

    url = f"{llama_url}/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {llama_api_key}",
    }

    data = {
        "model": llama_model_identity,
        "messages": conversation,
        "temperature": 0.01,
        "max_tokens": 1024,
    }

    response = requests.post(url, headers=headers, json=data)
    helpful_raise_for_status(response)

    content = (
        response.json()
        .get("choices", [{}])[0]
        .get("message", {})
        .get("content", "")
    )
    json_content = repair_json(content, return_objects=True)

    return json_content["query"]


@redis_wrapper.cache_for(3600 * 24 * 30)
@log_function_call
def is_analyzing_token_conversation(conversation: str) -> str:
    system_prompt = "You are a helpful assistant."

    prompt_to_use = """
Act as an expert in Natural Language Processing (NLP) and tweet analysis. Your task is to carefully examine a given Twitter tweet and determine whether it contains a request to analyze a specific token. To accomplish this, you will need to thoroughly analyze the language, tone, and content of the tweet, considering all possible contexts and nuances.

When evaluating the tweet, consider the following steps:

1. Identify any mentions of tokens, cryptocurrencies, or other relevant financial instruments.
2. Look for keywords or phrases that imply a request for analysis, such as "analyze," "evaluate," "assess," or "review."
3. Consider the tone and intent behind the tweet, taking into account the user's potential motivations and goals.
4. Weigh the possibility of the tweet being a genuine request for analysis against the likelihood of it being a spam or unrelated post.
5. If you determine that the tweet is a request to analyze a specific token, extract the name of the token from the tweet.

In your response, provide a JSON object containing three fields:
- "answer": a string indicating whether the tweet is a request to analyze a token ("yes" or "no")
- "token_name": the name of the token to be analyzed (if the answer is "yes"), or an empty string (if the answer is "no")
- "reason": a brief explanation for your answer, including any relevant context or insights gained from analyzing the tweet

Before providing your final answer, list out your thoughts and considerations, ensuring that you have thoroughly evaluated all aspects of the tweet. If complex reasoning is required, break down your thought process into step-by-step analysis, considering multiple perspectives and weighing the evidence.

Example of a potential tweet: "Can someone please analyze the current market trends for $BTC and predict its future value?"

When responding to the input tweet, utilize advanced prompt engineering techniques, such as Chain of Thought and Self Reflection, to ensure a comprehensive and accurate analysis. Consider the following questions:
- What are the key indicators that suggest the tweet is a request for analysis?
- How do the language and tone of the tweet support or contradict this interpretation?
- Are there any potential biases or assumptions that could influence my analysis, and how can I mitigate these factors?

By carefully considering these factors and providing a well-structured JSON response, you will be able to accurately detect whether a given tweet is a request to analyze a specific token.

Please respond with the JSON object containing your answer, token name, and reason.

Here is the input tweet: {}
""".format(
        conversation
    )

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

    def run_llm():
        model = SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=0.7,
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            api_key=const.SELF_HOSTED_LLAMA_API_KEY,
            model=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            seed=random.randint(1, int(1e9)),
        )

        result = model.generate(messages).generations[0].text
        log_data = {"conversation": conversation, "result": result}
        logger.info(
            f"[is_analyzing_token_conversation] {json.dumps(log_data)}"
        )
        result = repair_json(result, return_objects=True)
        result = result["answer"]

        if result.lower() not in ["yes", "no"]:
            raise Exception("Invalid answer")

        return result.lower()

    result = retry(
        run_llm,
        max_retry=3,
        first_interval=get_llm_tasks_first_interval(),
        interval_multiply=2,
    )()
    return result == "yes"


class IncludeExcludeItem(BaseModel):
    included: bool
    excluded: bool
    reason: str


class IncludeExcludeInfo(BaseModel):
    hashtags: IncludeExcludeItem
    urls: IncludeExcludeItem
    emojis: IncludeExcludeItem
    mentions: IncludeExcludeItem


@redis_wrapper.cache_for(3600 * 24 * 30)
@log_function_call
# @TODO: rewrite this
def detect_included_excluded_items(
    system_prompt, task_prompt
) -> IncludeExcludeInfo:
    req_system_prompt = """Persona: 
Act as an expert in natural language processing and prompt engineering with deep expertise in analyzing and synthesizing complex requirements for AI systems.

Task:
Your task is to analyze a given system prompt and user prompt intended for an AI that generates tweets for Twitter. Specifically, determine whether the system or user prompt explicitly requires the inclusion or exclusion of the following items from the generated tweets:
1. Hashtags
2. URLs
3. Emojis
4. Mentions
For each item, identify whether there is a clear and explicit requirement for inclusion or exclusion (e.g., "Include hashtags" or "Do not include URLs").

Context:
The AI generating tweets is designed to respond based on both system-level instructions (guiding its behavior generally) and user-level instructions (specific to individual requests). In cases where there is conflicting information between the system and user prompts, the user prompt takes precedence.

Response Format:
Provide your response in the following JSON format:

{
  "hashtags": {
    "included": <true or false>,
    "excluded": <true or false>,
    "reason": "Explain why hashtags are/are not required to be included or excluded."
  },
  "urls": {
    "included": <true or false>,
    "excluded": <true or false>,
    "reason": "Explain why URLs are/are not required to be included or excluded."
  },
  "emojis": {
    "included": <true or false>,
    "excluded": <true or false>,
    "reason": "Explain why emojis are/are not required to be included or excluded."
  },
  "mentions": {
    "included": <true or false>,
    "excluded": <true or false>,
    "reason": "Explain why mentions are/are not required to be included or excluded."
  }
}

Instructions for Reasoning:
- Carefully examine the system prompt for explicit instructions regarding the inclusion or exclusion of hashtags, URLs, emojis, and mentions.
- Analyze the user prompt for explicit instructions related to these items.
- If there are conflicting instructions, prioritize the user prompt over the system prompt.
- Focus solely on explicit requirements stated in the text.

Output the final analysis in the JSON format provided.
"""

    req_user_prompt = f"""Here are the system prompt and task prompt:

System prompt: {system_prompt}

Task prompt: {task_prompt}
"""

    messages = [
        {
            "role": "system",
            "content": req_system_prompt,
        },
        {
            "role": "user",
            "content": req_user_prompt,
        },
    ]

    def run_llm():
        model = SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=0.7,
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            api_key=const.SELF_HOSTED_LLAMA_API_KEY,
            model=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            seed=random.randint(1, int(1e9)),
        )
        result = model.generate(messages).generations[0].text
        result = repair_json(result, return_objects=True)

        print("result:", result)

        obj = IncludeExcludeInfo.model_validate(result)
        return obj

    obj = retry(
        run_llm,
        max_retry=3,
        first_interval=get_llm_tasks_first_interval(),
        interval_multiply=2,
    )()
    return obj


def clean_text(crawled_text: str):
    system_prompt = """You are a helpful assistant. Your task is to summarize the following crawled text in one paragraph, highlighting keywords, locations, dates, and other metadata. Provide the cleaned version of the text in a structured JSON format with the key "cleaned_text". Ensure no additional comments are added.

    Output format example:
    {
        "cleaned_text": "Your cleaned and summarized text here"
    }
    """
    url = os.path.join(const.SELF_HOSTED_LLAMA_405B_URL, "v1/chat/completions")

    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {const.SELF_HOSTED_LLAMA_API_KEY}",
    }

    data = {
        "model": const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": crawled_text[: const.MAX_TEXT_LENGTH]},
        ],
        "max_tokens": 1024,
        "temperature": 0.01,
    }

    resp = requests.post(url, headers=headers, data=json.dumps(data))

    if resp.status_code != 200:
        return crawled_text

    content = resp.json()["choices"][0]["message"]["content"]

    try:
        data = repair_json(content, return_objects=True)
        return data["cleaned_text"]
    except Exception as err:
        logger.error(f"[clean_text] An error occured: {err}")
        return crawled_text


EXTRACT_CONTENT_PROMPT = """Act as an expert in natural language processing (NLP) and information retrieval. You specialize in extracting the most relevant content from search results based on a given query. Your expertise ensures that only the most meaningful, contextually relevant information is included in the final output.

## Task:

Given a search query and a corresponding search result, extract a single, well-structured paragraph that is highly relevant to the search query. The extracted content should be:
- Concise: Remove unnecessary information that does not directly relate to the query.
- Contextually relevant: Preserve key details needed to understand the extracted content in relation to the query.
- Well-structured: Ensure the extracted content is readable as a single, coherent paragraph.

## Context:
- The search result may contain extraneous information such as ads, navigation links, or general introductions. Ignore these and focus only on content that directly addresses the query.
- If multiple relevant sections exist, merge them into a single, fluid paragraph without losing meaning.
- If no relevant information is found, return an empty string.
- If there is uncertainty about what to extract, list key reasoning steps before making a decision.

## Response Format (JSON):

Your output must be a valid JSON object with the following structure:

{{
  "content": "<single extracted paragraph>"
}}

where "content" is a single paragraph containing the extracted relevant content. If no relevant content is found, return an empty string ("").

## Instructions for Thought Process:

1. Analyze the search query to determine the core information being requested.
2. Review the search result and identify portions that directly answer the query.
3. Filter out irrelevant information such as ads, unrelated sections, or promotional content.
4. Merge relevant sections into a single, well-structured paragraph while maintaining clarity.
5. Format the output strictly as JSON, ensuring proper nesting and valid syntax.

If any ambiguity arises, list your reasoning steps before finalizing the extracted content.

## Provided Input:

Search Query: {query}

Search Results:
{result}
"""


def extract_content_relevant_to_query(query: str, contents: List[str]):
    system_prompt = "You are a helpful assistant."

    contents_str = "\n\n".join([f"- {x}" for x in contents])

    messages = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": EXTRACT_CONTENT_PROMPT.format(
                query=query, result=contents_str
            ),
        },
    ]

    def run_llm():
        model = SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=0.7,
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            api_key=const.SELF_HOSTED_LLAMA_API_KEY,
            model=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            seed=random.randint(1, int(1e9)),
        )
        result = model.generate(messages).generations[0].text
        result = repair_json(result, return_objects=True)

        return result["content"]

    obj = retry(
        run_llm,
        max_retry=3,
        first_interval=get_llm_tasks_first_interval(),
        interval_multiply=2,
    )()
    return obj


SUMMARIZE_JUDGE_COMMENTARY_PROMPT = """Act as an expert in extracting key insights and delivering sharp, impactful summaries. Your role is to distill a judge’s detailed commentary from a QA game into a concise yet compelling evaluation summary while preserving clarity and essential details.

## Task
Analyze the judge’s commentary, summarize their evaluation process, and explain the decision-making. Your summary must:
- Always include the winning user's name.
- Clearly state why they won.
- Support multiple reasons, each within 255 characters.
- Avoid mentioning specific judging criteria.

## Context
- Focus only on how the judge assessed the responses and reached a decision.
- Do not reference specific scoring metrics or criteria.
- Ensure each reason is independent and concise (≤255 characters per reason).
- Keep the language clear, neutral, and engaging.

## Response Format
Return a JSON object with multiple reasons in separate fields:

{{
  "summary": {{
    "1 ": "<First reason within 255 characters>",
    "2 ": "<Second reason within 255 characters>",
    "3 ": "<Third reason within 255 characters>",
    ...
  }}
}}

## Summarization Guidelines
1. **Emphasize the Process**: Explain how the judge reviewed, compared, and deliberated.
2. **Include the Winner**: Clearly mention who won.
3. **State Multiple Reasons**: Provide distinct, compelling reasons (≤255 characters each).
4. **Maximize Impact**: Use strong, direct language while staying neutral and precise.
5. **Ensure Readability**: Keep sentences structured, fluid, and engaging.

## Provided Input
Commentary: {commentary}
"""


def summarize_judge_commentary(commentary: str) -> str:
    system_prompt = "You are a helpful assistant."

    messages = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": SUMMARIZE_JUDGE_COMMENTARY_PROMPT.format(
                commentary=commentary
            ),
        },
    ]

    def run_llm():
        model = SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=0.7,
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            api_key=const.SELF_HOSTED_LLAMA_API_KEY,
            model=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            seed=random.randint(1, int(1e9)),
        )
        result = model.generate(messages).generations[0].text
        result = repair_json(result, return_objects=True)

        return result["summary"]

    obj = retry(
        run_llm,
        max_retry=3,
        first_interval=get_llm_tasks_first_interval(),
        interval_multiply=2,
    )()
    return obj
