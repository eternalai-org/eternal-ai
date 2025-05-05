import json
import logging
import random
import requests
from typing import List, Tuple

from x_content.wrappers.magic import helpful_raise_for_status
from x_content.wrappers.magic import retry

from ..models import TwinUpdateResponse
from .utils import PROMPT_TEMPLATE, names_list_to_string
from .utils import create_llm, pull_tweets

from x_content.wrappers.log_decorators import log_function_call
from x_content.wrappers.rag_search import insert_to_db
from x_content.wrappers.postprocess import (
    post_process_discord_message,
    post_process_knowledge_base_tweet,
)
from x_content import constants as const
from json_repair import repair_json


def try_load_json_from_str(s: str):
    try:
        return repair_json(s, return_objects=True)
    except json.JSONDecodeError:
        return None


logger = logging.getLogger(__name__)

MAX_RETRY = 3
DEFAULT_PREFIX_KNOWLEDGE_ID = "kn"
BATCH_SIZE = 100


class TwinService:

    def __init__(self):

        self.hermes = create_llm(
            base_url=const.SELF_HOSTED_HERMES_70B_URL + "/v1",
            model_id=const.SELF_HOSTED_HERMES_70B_MODEL_IDENTITY,
            api_key=const.SELF_HOSTED_HERMES_70B_API_KEY,
            temperature=0.01,
        )

        self.llama = create_llm(
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            model_id=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            temperature=0.01,
        )

    @log_function_call
    def _generate_style(self, context: str, current_style: dict) -> str:
        system_prompt = f"""
You are an advanced language model with expertise in analyzing and synthesizing writing styles of Twitter tweets. Your task is to analyze a provided list of tweets (up to 100 examples) and identify the key stylistic elements that define the user's writing style. Then, integrate these insights with a provided current style to create a unified and refined writing style.

### Instructions:
1. Carefully review the provided examples for patterns in tone, sentence structure, sentence flow and connection, vocabulary and any stylistic nuances.
2. Identify and describe the following key elements of the writing style of list of tweet examples:
- **Tone:** What emotional undertone does the writing convey? (e.g., humorous, formal, casual, motivational)
- **Sentence Structure:** Are the sentences short and punchy or long and elaborate? Are they mostly statements, questions, or exclamations?
- **Vocabulary:** Is the language simple, sophisticated, playful, technical, or colloquial?
- **Pacing and Rhythm:** Is the writing fast-paced or slow? Does it use repetition, pauses, or rhetorical devices?
- **Text Stylization:** Are there patterns in capitalization (e.g., all caps, all lowercase, sentence case, mixed)? Are there unique formatting choices like line breaks or repeated punctuation? 
- **Catch Phrases:** Are there any recurring phrases or expressions that are unique to the user? (e.g., "to the moon", "hodl", "buy the dip"). How frequently are they used?
- **Liked topics:** What are the topics that are liked by the user? (e.g., bitcoin, ethereum, dogecoin, etc.)
- **Disliked topics:** What are the topics that are disliked by the user? (e.g., politics, sports, etc.)
3. Provide a concise yet detailed summary of the writing style based on these elements.
4. Based on the analysis, synthesize the core style elements from the provided list of tweets and combine them with the **current style**. Your goal is combine the analyzed style of the tweet examples with the provided current style. Your final style must be retains the defining characteristics of both, ensuring it is comprehensive, adaptable, and accurate

#### Current Style:
{json.dumps(current_style, indent=2)}

Output your analysis in a stringified JSON format with the following keys:
- "tone"
- "sentence_structure"
- "vocabulary"
- "pacing_and_rhythm"
- "text_stylization"
- "catch_phrases"
- "liked_topics"
- "disliked_topics"

Only return the tweet's style based on your analysis above.
"""
        user_prompt = f"Here are the tweets:\n{context}"
        prompt = PROMPT_TEMPLATE.invoke(
            {"system_prompt": system_prompt, "question": user_prompt}
        )

        resp = self.hermes.invoke(prompt)
        parsed_content = try_load_json_from_str(resp.content)

        if parsed_content is not None:
            return parsed_content
        else:
            logger.error(f"Error in [_generate_style] with hermes")

            resp = self.llama.invoke(prompt)
            parsed_content = try_load_json_from_str(resp.content)

            if not parsed_content:
                logger.error(f"Error in [_generate_style] with llama")
                return {}

            return parsed_content

    def _generate_style_from_tweets(self, data_tweets: List[str]) -> str:
        shuffled_tweets = data_tweets.copy()
        random.shuffle(shuffled_tweets)
        styles = {}
        debug_data = []
        max_len = len(shuffled_tweets)

        for i in range(0, max_len, BATCH_SIZE):
            context = shuffled_tweets[i : i + BATCH_SIZE]
            context = "\n".join([f"- {x}" for x in context])

            styles = self._generate_style(
                context=context, current_style=styles
            )

            debug_data.append(
                {
                    "index": f"[{i}, {i+BATCH_SIZE})",
                    "shuffled_tweets": shuffled_tweets[i : i + BATCH_SIZE],
                    "styles": styles,
                }
            )

        return styles

    @log_function_call
    def _pull_tweets(self, ids: List[str]) -> Tuple[List[str], Exception]:
        return pull_tweets(ids)

    @log_function_call
    def _insert_to_db(self, agent_id: str, pull_tweets: List[str]) -> str:
        return insert_to_db(agent_id, pull_tweets)

    @log_function_call
    def generate_system_prompt(
        self, tweets: List[str], names: List[str] = []
    ) -> str:
        try:

            def try_get_system_prompt():
                style_attributes = self._generate_style_from_tweets(tweets)

                flavor_text = ""
                if len(names) > 0:
                    names_str = names_list_to_string(names)
                    flavor_text = f"trained on data inspired by {names_str} "

                system_prompt = f"""
You are an advanced AI agent {flavor_text}with a specific style in writing a tweet.

### Personalized tweet style:
- **Tone:** {style_attributes['tone']}
- **Sentence Structure:** {style_attributes['sentence_structure']}
- **Vocabulary:** {style_attributes['vocabulary']}
- **Pacing and Rhythm:** {style_attributes['pacing_and_rhythm']}
- **Text Stylization:** {style_attributes["text_stylization"]}
- **Catch Phrases:** {style_attributes["catch_phrases"]}
- **Liked topics:** {style_attributes["liked_topics"]}
- **Disliked topics:** {style_attributes["disliked_topics"]}
"""
                return system_prompt

            system_prompt = retry(
                try_get_system_prompt,
                max_retry=3,
                first_interval=60,
                interval_multiply=2,
            )()
            return system_prompt, None
        except Exception as e:
            logger.error(f"Error in [generate_system_prompt]: {e}")
            return None, e

    @log_function_call
    def _update_twin_status(self, twin_response: TwinUpdateResponse) -> bool:
        url = f"{const.BASE_TWIN_API_URL}/agent/update_twin_status"
        payload = twin_response.dict()
        try:
            resp = requests.post(url, json=payload, timeout=10)
            helpful_raise_for_status(resp)
            logger.info(
                f"[_update_twin_status] Response: {resp.status_code} - {resp.text}"
            )
            return True
        except requests.RequestException as e:
            logger.error(f"Error in [_update_twin_status]: {e}")
            return False

    @log_function_call
    def generate_twin(self, agent_id: str, twitter_ids: List[str]) -> dict:
        twin_status, knowledge_id, system_prompt = "running", "", ""

        def _update_error(error_msg, err):
            logger.error(error_msg + f": {err}")
            twin_status_err = "done_error"
            twin_resp = TwinUpdateResponse(
                agent_id=agent_id,
                twin_status=twin_status_err,
                knowledge_base_id="",
                system_prompt="",
                twin_training_progress=100.0,
                twin_training_message=error_msg,
            )
            self._update_twin_status(twin_resp)
            return None, err

        try:
            try:
                pulled_tweets, err = self._pull_tweets(twitter_ids)
                if err is not None:
                    raise Exception(f"Pulling tweets failed with error {err}")
                processed_tweets = [
                    post_process_knowledge_base_tweet(t) for t in pulled_tweets
                ]
                processed_tweets = list(
                    filter(lambda x: x != "", processed_tweets)
                )
            except Exception as e:
                return _update_error("Unable to pull tweets.", e)

            if len(processed_tweets) == 0:
                return _update_error(
                    "Failed. Try creating a clone by combining different DNA.",
                    None,
                )

            # Update progress
            twin_response = TwinUpdateResponse(
                agent_id=agent_id,
                twin_status=twin_status,
                knowledge_base_id="",
                system_prompt="",
                twin_training_progress=25.0,
            )
            self._update_twin_status(twin_response)

            knowledge_id = DEFAULT_PREFIX_KNOWLEDGE_ID + agent_id
            try:
                self._insert_to_db(knowledge_id, processed_tweets)
            except Exception as e:
                return _update_error("Error inserting tweets to DB", e)

            twin_response = TwinUpdateResponse(
                agent_id=agent_id,
                twin_status=twin_status,
                knowledge_base_id="",
                system_prompt="",
                twin_training_progress=65.0,
            )
            self._update_twin_status(twin_response)

            base_url = const.TWITTER_API_URL
            twitter_api_key = const.TWITTER_API_KEY
            session = requests.Session()
            session.params = {"api_key": twitter_api_key}

            names = []
            for twitter_id in twitter_ids:
                url = f"{base_url}/user/{twitter_id}"

                try:
                    resp = session.get(url)
                    helpful_raise_for_status(resp)
                    result = resp.json()["result"]
                    names.append(result["name"])
                except (requests.RequestException, KeyError) as err:
                    return _update_error(
                        "[generate_twin] Error retrieving twitter name", err
                    )

            for attempt in range(MAX_RETRY + 1):
                system_prompt, err = self.generate_system_prompt(
                    processed_tweets, names
                )

                if err is None:
                    twin_status = "done_success"
                    break

                logger.error(
                    f"[generate_twin] Error generating system prompt {repr(err)}"
                )

                if attempt == MAX_RETRY:
                    return _update_error(
                        "[generate_twin] Failed after all retries", err
                    )

                logger.info(
                    f"[generate_twin] Retrying {attempt + 1}/{MAX_RETRY}"
                )
        except Exception as e:
            return _update_error("Error in [generate_twin]", e)

        twin_response = TwinUpdateResponse(
            agent_id=agent_id,
            twin_status=twin_status,
            knowledge_base_id=knowledge_id,
            system_prompt=system_prompt,
            twin_training_progress=100.0,
        )

        self._update_twin_status(twin_response)
        return {
            "knowledge_id": knowledge_id,
            "system_prompt": system_prompt,
        }, None

    @log_function_call
    def generate_twin_from_discord_messages(
        self, agent_id: str, file_path: str, usernames: List[str]
    ) -> dict:
        twin_status, knowledge_id, system_prompt = "running", "", ""

        def _update_error(error_msg, err):
            logger.error(error_msg + f": {err}")
            return None, err

        try:
            with open(file_path, "r") as f:
                message_by_usernames = json.loads(f.read())

            messages = []
            for data in message_by_usernames:
                if data["name"] in usernames:
                    messages += data["messages"]
            messages = [post_process_discord_message(t) for t in messages]

            if len(messages) == 0:
                return _update_error(
                    "Failed. Try creating a clone by combining different DNA.",
                    None,
                )

            knowledge_id = DEFAULT_PREFIX_KNOWLEDGE_ID + agent_id
            try:
                self._insert_to_db(knowledge_id, messages)
            except Exception as e:
                return _update_error("Error inserting tweets to DB", e)

            for attempt in range(MAX_RETRY + 1):
                system_prompt, err = self.generate_system_prompt(
                    messages, usernames
                )

                if err is None:
                    twin_status = "done_success"
                    break

                logger.error(
                    f"[generate_twin_from_discord_messages] Error generating system prompt {repr(err)}"
                )

                if attempt == MAX_RETRY:
                    return _update_error(
                        "[generate_twin_from_discord_messages] Failed after all retries",
                        err,
                    )

                logger.info(
                    f"[generate_twin_from_discord_messages] Retrying {attempt + 1}/{MAX_RETRY}"
                )
        except Exception as e:
            return _update_error(
                "Error in [generate_twin_from_discord_messages]", e
            )

        return {
            "knowledge_id": knowledge_id,
            "system_prompt": system_prompt,
        }, None

    @log_function_call
    def twin_system_prompt(self, twitter_ids: List[str]) -> str:
        pulled_tweets, err = self._pull_tweets(twitter_ids)
        if err is not None:
            return "", []
        pulled_tweets = [
            post_process_knowledge_base_tweet(x) for x in pulled_tweets
        ]

        base_url = const.TWITTER_API_URL
        twitter_api_key = const.TWITTER_API_KEY
        session = requests.Session()
        session.params = {"api_key": twitter_api_key}

        names = []
        for twitter_id in twitter_ids:
            url = f"{base_url}/user/{twitter_id}"
            try:
                resp = session.get(url)
                result = resp.json()["result"]
                names.append(result["name"])
            except Exception as err:
                logger.error(
                    f"[generate_system_prompt] Error retrieving twitter name: {err}"
                )

        system_prompt, err = self.generate_system_prompt(pulled_tweets, names)
        return system_prompt, pulled_tweets


twin_service = TwinService()
