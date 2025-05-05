import asyncio
import logging
import random
import traceback
import datetime
from typing import List

from json_repair import repair_json
from x_content.constants import MissionChainState
from x_content import constants as const
from x_content.llm.base import OnchainInferResult
from x_content.llm.local import SyncBasedEternalAI
from x_content.models import ReasoningLog
from x_content.tasks.base import MultiStepTaskBase
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.response import (
    ExtendedTweetInfoDto,
    GenerateActionDto,
    Response,
    TweetsDto,
)
from x_content.wrappers.game import (
    GameAPIClient,
    GameInfo,
    _get_game_redis_cache,
    GameStatus,
    GameState,
)
from x_content.wrappers.llm_tasks import summarize_judge_commentary
from x_content.tasks.game_agent.subtasks.reply.reply_create_game import _post_wallet_tweet
from x_content.services import fact_service
from x_content.services.fact.fact_check import SearchResponse

from x_content.tasks.utils import a_move_state, create_twitter_auth_from_reasoning_log

from x_content.wrappers.vision_tasks import get_image_description
from x_content.wrappers.magic import get_llm_tasks_first_interval, retry, sync2async

from x_content.tasks.game_agent.prompts import FACT_QUERY_PROMPT_TEMPLATE, JUDGE_GAME_PROMPT_TEMPLATE, JUDGE_FACT_PROMPT_TEMPLATE, JUDGE_GAME_WITH_FACTS_PROMPT_TEMPLATE
from x_content.wrappers.tweet_specialty import is_create_game_tweet, is_find_fact_tweet
from x_content.wrappers.postprocess import StringProcessor

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)


async def _get_ready_to_judge_games():
    # Get a list of games ready to judge. Return a list of ids.
    try:
        # Get all running games from redis cache
        logger.info(
            "[_get_ready_to_judge_games] Getting running games from Redis"
        )

        game_redis = _get_game_redis_cache()
        running_games = await sync2async(game_redis.get_running_games)()

        if not running_games:
            logger.info("[_get_ready_to_judge_games] No running games found")
            return []

        ready_tweet_ids = []

        # Filter games that have ended
        logger.info(
            f"[_get_ready_to_judge_games] Found {len(running_games)} running games, checking status..."
        )

        # Get game info for all running games in one API call
        games_info, err = await sync2async(
            GameAPIClient.get_game_info_by_tweet_ids
        )(running_games)

        if err:
            logger.error(
                f"[_get_ready_to_judge_games] Error getting game info: {err}"
            )
            return []

        if not games_info:
            logger.info("[_get_ready_to_judge_games] No game info returned")
            return []

        # Filter games that have ended status
        for game in games_info:
            if game.status == GameState.ENDED:
                logger.info(
                    f"[_get_ready_to_judge_games] Game {game.tweet_id} is ready to judge"
                )
                ready_tweet_ids.append(game.tweet_id)

        logger.info(
            f"[_get_ready_to_judge_games] Found {len(ready_tweet_ids)} games ready to judge"
        )
        return ready_tweet_ids

    except Exception as err:
        logger.error(
            f"[_get_ready_to_judge_games] Error getting ready games: {err}"
        )
        return []


def _is_judge_game_pending(tweet_id):
    game_redis = _get_game_redis_cache()
    status = game_redis.get_game_status(tweet_id)
    is_pending = (
        status == GameStatus.CREATE_DONE or status == GameStatus.JUDGE_PENDING
    )
    logger.info(
        f"[_is_judge_game_pending] Game {tweet_id} status: {status} {status == GameStatus.CREATE_DONE } pending status: {is_pending}"
    )
    return is_pending


def _mark_judge_game_pending(tweet_id):
    logger.info(
        f"[_mark_judge_game_pending] Marking game {tweet_id} as pending"
    )
    game_redis = _get_game_redis_cache()
    return game_redis.set_game_status(tweet_id, GameStatus.JUDGE_PENDING)


def _mark_judge_game_running(tweet_id):
    logger.info(
        f"[_mark_judge_game_running] Marking game {tweet_id} as running"
    )
    game_redis = _get_game_redis_cache()
    return game_redis.set_game_status(tweet_id, GameStatus.JUDGE_RUNNING)


def _mark_judge_game_done(tweet_id):
    logger.info(f"[_mark_judge_game_done] Marking game {tweet_id} as done")
    try:
        game_redis = _get_game_redis_cache()
        game_redis.remove_running_game(tweet_id, GameStatus.JUDGE_DONE)
    except Exception as err:
        logger.error(
            f"[_mark_judge_game_done] Error marking game as done: {err}"
        )


def _try_acquire_judge_game_lock(_tweet_id):
    # Initialize lock variable
    lock = None

    try:
        game_redis = _get_game_redis_cache()
        # Attempt to acquire the lock with an expiry time
        lock = game_redis.acquire_judge_game_lock(
            _tweet_id, const.JUDGE_GAME_LOCK_EXPIRY  # e.g., 300 seconds
        )
        logger.info(
            f"[_try_acquire_judge_game_lock] Lock acquisition attempt for game {_tweet_id}: {'successful' if lock else 'failed'}"
        )
        return lock  # Returns True if acquired, False if already locked

    except Exception as e:
        # Log any errors during lock acquisition
        logger.error(f"Error acquiring lock for tweet {_tweet_id}: {str(e)}")
        return False

    finally:
        # Always execute this block, even if an exception occurs
        if not lock:  # If lock wasn't acquired
            try:
                # Attempt to release the lock
                game_redis.release_judge_game_lock(_tweet_id)
                logger.info(
                    f"[_try_acquire_judge_game_lock] Released lock for game {_tweet_id}"
                )
            except Exception as e:
                # Log any errors during lock release
                logger.error(
                    f"Error releasing lock for tweet {_tweet_id}: {str(e)}"
                )


def _get_facts_search_query_conversation(
    question: str, timestamp: str, content_images: List[str]
) -> List[dict]:
    system_message = "You are a helpful assistant."

    content_images_str = ""
    if len(content_images) > 0:
        content_images_str = "Image Descriptions:" + "\n\n".join(
            [f"- {x}" for x in content_images]
        )

    conversation_thread = [
        {
            "role": "system",
            "content": system_message,
        },
        {
            "role": "user",
            "content": FACT_QUERY_PROMPT_TEMPLATE.format(
                question=question,
                content_images=content_images_str,
                tweet_timestamp=timestamp,
                current_timestamp=datetime.datetime.now(
                    datetime.timezone.utc
                ).isoformat(),
            ),
        },
    ]

    logger.info(
        f"[_get_facts_search_query_conversation] Successfully built conversation: {conversation_thread}"
    )

    return conversation_thread


def _get_facts_search_query(
    question: str, timestamp: str, content_images: List[str]
) -> str:
    messages = _get_facts_search_query_conversation(
        question, timestamp, content_images
    )

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
        print("[_get_facts_search_query]", result)
        result = repair_json(result, return_objects=True)

        return result["query"]

    query = retry(
        run_llm,
        max_retry=3,
        first_interval=get_llm_tasks_first_interval(),
        interval_multiply=2,
    )()
    return query


async def _get_judge_game_with_facts_conversation(
    game_tweet_object: dict,
    answers: List[dict],
):
    """
    Return conversation to judge the winning agent from given game tweet and participating agents' answers
    :param game_tweet_object: tweet object with schema { "full_text": str, "img_url": str }
    :param answers: list of objects with schema { "username": str, "answer": str } representing the answer from participating agents
    :return conversation to submit LLM inference
    """
    logger.info(
        "[_get_judge_game_with_facts_conversation] Building conversation for judging"
    )

    system_message = "You are a helpful assistant."

    answers_content = ""
    for answer in answers:
        answers_content += f"- {answer['username']}: {answer['answer']}\n"

    content_images = []
    if game_tweet_object.get("image_urls"):
        for img_url in game_tweet_object.get("image_urls"):
            try:
                content_images.append(
                    await sync2async(get_image_description)(img_url)
                )
            except Exception as err:
                logger.error(
                    f"[_get_judge_game_with_facts_conversation] Failed to get image description at {img_url}: {err}"
                )

    # Perform search
    # Try to get cached response first
    game_redis = _get_game_redis_cache()
    cached_response = game_redis.get_fact_check(
        game_tweet_object.get("tweet_id")
    )

    given_facts = ""
    if cached_response:
        logger.info(
            f"[_get_judge_game_with_facts_conversation] Using cached fact check response for tweet {game_tweet_object.get('tweet_id')}"
        )
        given_facts = SearchResponse(**cached_response).results
    else:
        # No cache found, call fact service
        logger.info(
            f"[_get_judge_game_with_facts_conversation] No cache found, calling fact service for tweet {game_tweet_object.get('tweet_id')}"
        )

        query = await sync2async(_get_facts_search_query)(
            game_tweet_object.get("full_text"),
            game_tweet_object.get("posted_at"),
            content_images,
        )

        if query != "":
            logger.info(
                f"[_get_judge_game_with_facts_conversation] Received query '{query}' for fact searching for tweet {game_tweet_object.get('tweet_id')}"
            )

            response = await fact_service.search(
                query, time_range="d", search_depth="advanced"
            )

            # Cache the response
            game_redis.set_fact_check(
                game_tweet_object.get("id"), response.__dict__
            )

            given_facts = response.results
        else:
            logger.info(
                f"[_get_judge_game_with_facts_conversation] Fact searching is not required for tweet {game_tweet_object.get('tweet_id')}"
            )

    logger.info(
        "[_get_judge_game_with_facts_conversation] Building conversation for judging"
    )

    content_images_str = ""
    if len(content_images) > 0:
        content_images_str = "Content images in the tweet:" + "\n\n".join(
            [f"- {x}" for x in content_images]
        )
    if given_facts == "":
        given_facts = "Facts not found"

    user_prompt = JUDGE_GAME_WITH_FACTS_PROMPT_TEMPLATE.format(
        full_text=game_tweet_object.get("full_text"),
        content_images=content_images_str,
        tweet_timestamp=game_tweet_object.get("posted_at"),
        current_timestamp=datetime.datetime.now(
            datetime.timezone.utc
        ).isoformat(),
        given_facts=given_facts,
        answers_content=answers_content,
    )

    conversation_thread = [
        {
            "role": "system",
            "content": system_message,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]
    logger.info(
        f"[_get_judge_game_with_facts_conversation] Successfully built conversation: {conversation_thread}"
    )
    return conversation_thread


async def _get_judge_fact_conversation(tweet_object, answers):
    system_message = "You are a helpful assistant."

    answers_content = ""
    for answer in answers:
        answers_content += f"- {answer['username']}: {answer['answer']}\n"

    # Example claim to search
    query = tweet_object.get("full_text")
    query = (
        StringProcessor(query)
        .remove_tags()
        .remove_mentions()
        .remove_urls()
        .remove_emojis()
        .strip_head_and_tail_white_string()
        .get_text()
    )

    # Perform search
    # Try to get cached response first
    game_redis = _get_game_redis_cache()
    cached_response = game_redis.get_fact_check(tweet_object.get("id"))

    if cached_response:
        logger.info(
            f"[_get_judge_fact_conversation] Using cached fact check response for tweet {tweet_object.get('id')}"
        )
        response = SearchResponse(**cached_response)
    else:
        # No cache found, call fact service
        logger.info(
            f"[_get_judge_fact_conversation] No cache found, calling fact service for tweet {tweet_object.get('id')}"
        )
        response = await fact_service.search(query)

        # Cache the response
        game_redis.set_fact_check(tweet_object.get("id"), response.__dict__)

    user_prompt = JUDGE_FACT_PROMPT_TEMPLATE.format(
        query=tweet_object.get("full_text"),
        fact=response.results,
        answers=answers_content,
    )

    conversation_thread = [
        {
            "role": "system",
            "content": system_message,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]
    logger.info(
        f"[_get_judge_game_conversation] Successfully built conversation: {conversation_thread}"
    )
    return conversation_thread


async def _get_judge_game_conversation(game_tweet_object, answers):
    """
    Return conversation to judge the winning agent from given game tweet and participating agents' answers
    :param game_tweet_object: tweet object with schema { "full_text": str, "img_url": str }
    :param answers: list of objects with schema { "username": str, "answer": str } representing the answer from participating agents
    :return conversation to submit LLM inference
    """
    logger.info(
        "[_get_judge_game_conversation] Building conversation for judging"
    )

    system_message = "You are a helpful assistant."

    answers_content = ""
    for answer in answers:
        answers_content += f"- {answer['username']}: {answer['answer']}\n"

    content_images = ""
    if game_tweet_object.get("image_urls"):
        for img_url in game_tweet_object.get("image_urls"):
            try:
                content_images += await sync2async(get_image_description)(
                    img_url
                )
                content_images += "\n"
            except Exception as err:
                logger.error(
                    f"[_get_judge_game_conversation] Failed to get image description at {img_url}: {err}"
                )

    user_prompt = JUDGE_GAME_PROMPT_TEMPLATE.format(
        full_text=game_tweet_object.get("full_text"),
        content_images=content_images,
        answers_content=answers_content,
    )

    conversation_thread = [
        {
            "role": "system",
            "content": system_message,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]

    logger.info(
        f"[_get_judge_game_conversation] Successfully built conversation: {conversation_thread}"
    )

    return conversation_thread


async def _get_child_tweets(tweet_id):
    """
    Get all child tweets (replies) for a given tweet ID from the Twitter API
    :param tweet_id: str - ID of the parent tweet
    :return: list of child tweet objects with schema { "username": str, "full_text": str, "created_at": datetime }
    """
    try:
        logger.info(
            f"[_get_child_tweets] Fetching child tweets for tweet {tweet_id}"
        )

        resp: Response[TweetsDto] = await sync2async(
            twitter_v2.search_recent_tweet_by_tweetid
        )(tweet_id)

        if resp.is_error():
            raise Exception(resp.error)

        recent_replies = resp.data.tweets

        logger.info(
            f"[_get_child_tweets] Found {len(recent_replies)} child tweets"
        )

        # Call Twitter API to get replies
        # response = APICaller.GET_tweetReplies(tweet_id)

        # if not response:
        #     logger.error(f"[get_child_tweets_from_api] Failed to get replies for tweet {tweet_id}")
        #     return []
        # Extract relevant fields from response
        child_tweets = []
        for tweet in recent_replies:
            child_tweets.append(
                {
                    "username": tweet.twitter_username,
                    "full_text": tweet.full_text,
                    "posted_at": tweet.posted_at,
                }
            )

        return child_tweets

    except Exception as err:
        logger.error(
            f"[get_child_tweets_from_api] Error getting child tweets: {err}"
        )
        return []


async def _get_final_answers(log: ReasoningLog, _tweet_id):
    try:
        logger.info(
            f"[_get_final_answers] Getting final answers for game {_tweet_id}"
        )
        # Get child tweets and game data
        child_tweets, game_data, err = await _get_game_data_and_tweets(
            _tweet_id
        )

        if err is not None:
            return None, err

        # Get participant answers
        participants_answers, err = _get_participant_answers_from_tweets(
            game_data, child_tweets
        )

        if err is not None:
            return None, err

        answers = _format_answers_within_time_limit(
            game_data, participants_answers
        )

        logger.info(
            f"[_get_final_answers] Successfully got {len(answers)} final answers"
        )
        return answers, None

    except Exception as err:
        traceback.print_exc()
        logger.error(f"[_handle_judge_game_request] An error occured {err}")
        return None, err


async def _get_game_data_and_tweets(tweet_id):
    """Get child tweets and game data from BE"""
    logger.info(
        f"[_get_game_data_and_tweets] Getting child tweets for {tweet_id}"
    )
    child_tweets = await _get_child_tweets(tweet_id)
    if not child_tweets:
        logger.error(
            f"[_get_game_data_and_tweets] No child tweets found for tweet {tweet_id}"
        )
        return None, None, Exception("No child tweets found")
    logger.info(
        f"[_get_game_data_and_tweets] Found {len(child_tweets)} child tweets"
    )

    logger.info(
        f"[_get_game_data_and_tweets] Getting game data from BE for {tweet_id}"
    )
    game_data, err = await sync2async(GameAPIClient.get_game_info_by_tweet_id)(
        tweet_id
    )
    if err:
        return None, None, err

    return child_tweets, game_data, None


def _get_participant_answers_from_tweets(game_data: GameInfo, child_tweets):
    """Get first answer from each participant"""
    participants = [player.username for player in game_data.agent_wallets]
    logger.info(
        f"[_get_participant_answers_from_tweets] Found {len(participants)} participants"
    )

    # Sort by posted_at to get earliest tweets first
    sorted_child_tweets = sorted(child_tweets, key=lambda x: x["posted_at"])

    # Take first tweet from each participant
    seen_usernames = set()
    participants_answers = []

    for tweet in sorted_child_tweets:
        if (
            tweet["username"] in participants
            and tweet["username"] not in seen_usernames
        ):
            participants_answers.append(tweet)
            seen_usernames.add(tweet["username"])

    logger.info(
        f"[_get_participant_answers_from_tweets] Found {len(participants_answers)} participant answers"
    )

    return participants_answers, None


async def _handle_winner_without_llm(
    log: ReasoningLog, game_id, winner_agent: str
):
    """Handle case when there is no partipipant or only one participant, making them the winner by default"""
    logger.info(
        f"[_handle_winner_without_llm] Posting game result without calling LLM, winner if any: {winner_agent}"
    )
    _, err = await _post_game_result(log, game_id, winner_agent)
    return err


async def _handle_winner_with_llm(log: ReasoningLog, llm, game_id, answers):
    """Handle case when there are multiple participants, calling LLM to determine the winner"""
    # Create conversation for judging
    logger.info("[_handle_winner_with_llm] Creating conversation for judging")

    resp: Response[ExtendedTweetInfoDto] = await sync2async(
        twitter_v2.get_tweet_info_from_tweet_id
    )(game_id, True)

    if resp.is_error():
        return None, None, resp.error

    tweet_obj = resp.data.tweet_info.tweet_object.to_dict()
    # currently there are only 2 types: game & fact
    # so this works, consider alternative when the spec changes
    conversation_thread = await (
        _get_judge_game_conversation(tweet_obj, answers)
        if is_create_game_tweet(resp.data.tweet_info)
        else _get_judge_game_with_facts_conversation(tweet_obj, answers)
    )
    logger.info(
        f"[_handle_winner_with_llm] Created conversation thread for {game_id} judging: {conversation_thread}"
    )

    logger.info(
        f"[_handle_winner_with_llm] Calling LLM for {game_id} judgment"
    )
    infer_result: OnchainInferResult = await llm.agenerate(
        conversation_thread, temperature=0.7
    )
    logger.info(
        f"[_handle_winner_with_llm] Calling LLM for {game_id} llm.agenerate complete"
    )

    summarize = await sync2async(summarize_judge_commentary)(
        infer_result.generations[0].message.content
    )

    logger.info(
        f"[_handle_winner_with_llm] Calling LLM for {game_id} summarize_judge_commentary : {summarize}"
    )

    # err = await _post_wallet_tweet(log, game_id, summarize)
    for key, reason in summarize.items():
        err = await _post_wallet_tweet(log, game_id, reason)

    if err is not None:
        logger.error(
            f"[_handle_winner_with_llm] Error posting game {game_id} summarize: {err}"
        )
        return None, None, err

    winning_agent, err = _get_winner_from(
        infer_result.generations[0].message.content
    )
    if err is not None:
        logger.error(
            f"[_handle_winner_with_llm] Error getting winner for game {game_id} from LLM: {err}"
        )
        return None, None, err

    logger.info(
        f"[_handle_winner_with_llm] LLM for game {game_id} selected winner: {winning_agent}"
    )
    _, err = await _post_game_result(log, game_id, winning_agent)
    if err is not None:
        logger.error(
            f"[_handle_winner_with_llm] Error posting game {game_id} result: {err}"
        )
        return None, None, err

    return winning_agent, infer_result.tx_hash, None


def _format_answers_within_time_limit(
    game_data: GameInfo, participants_answers
):
    """Format answers that were submitted within time limit"""
    if not participants_answers:
        logger.info(
            "[_format_answers_within_time_limit] No participant answers found"
        )
        return []
    participants_answers_within_time_limit = [
        participant_answer
        for participant_answer in participants_answers
        if participant_answer["posted_at"] <= game_data.end_time
    ]

    logger.info(
        f"[_format_answers_within_time_limit] Found {len(participants_answers_within_time_limit)} answers within time limit"
    )

    return [
        {
            "username": participant_answer["username"],
            "answer": participant_answer["full_text"],
        }
        for participant_answer in participants_answers_within_time_limit
    ]


def _get_winner_from(result: str):
    """
    Parse returned result from LLM to get the name of winning agent
    :param result: returned result from LLM
    :return name of winning agent or None if no winner
    """
    try:
        logger.info(
            "[_get_winner_from] Parsing LLM result to determine winner"
        )
        result = repair_json(result, return_objects=True)
        result_found = result["result_found"]
        if not result_found:
            return None, Exception(f"Failed to find final result from llm")
        winner = result["winning_agent"]
        logger.info(f"[_get_winner_from] Successfully parsed winner: {winner}")
        return winner, None
    except Exception as err:
        traceback.print_exc()
        logger.error(f"[_get_winner_from] An error occured {err}")
        return None, err


async def _post_game_result(log: ReasoningLog, game_id, winning_agent):
    """Posts the game result to the API and tweets the winner.

    Args:
        _db: Database interface for posting tweets
        tweet_id: ID of the original game tweet
        winning_agent: Username of winning agent, or None if no winner

    Returns:
        Error if API call fails, None on success
    """
    logger.info(
        f"[_post_game_result] Posting game result for {game_id}, winner: {winning_agent if winning_agent else 'No winner'}"
    )

    # Call API to record game result
    _, err = await sync2async(GameAPIClient.submit_game_result)(
        game_id, winning_agent
    )

    if err is not None:
        logger.error(f"[_post_game_result] Error posting to API: {err}")
        return None, err

    # Post tweet announcing winner
    tweet_content = (
        const.WINNER_TWEET_TEMPLATE.format(winning_agent)
        if winning_agent
        else const.NO_WINNER_TWEET
    )
    logger.info(f"[_post_game_result] Posting result tweet: {tweet_content}")

    resp: Response[GenerateActionDto] = await sync2async(
        twitter_v2.reply_multi_unlimited
    )(
        auth=create_twitter_auth_from_reasoning_log(log),
        tweet_id=game_id,
        reply_content=tweet_content,
    )

    if resp.is_error() or not resp.data.success:
        logger.error(f"[_post_game_result] Error posting to Twitter: {err}")
        return None, err

    logger.info("[_post_game_result] Successfully posted game result")
    return None, None


class JudgeGameTask(MultiStepTaskBase):

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        if log.state == MissionChainState.NEW:
            logger.info("[JudgeGameTask.process_task] Starting new task")
            game_ids_need_to_judge = await _get_ready_to_judge_games()
            log = await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )
            log.execute_info = {
                "game_ids_need_to_judge": game_ids_need_to_judge,
                "task_result": [],
                "processing_idx": 0,
                "conversation": [],
            }
            log = await self.commit_log(log)
        else:
            logger.info("[JudgeGameTask.process_task] Resuming existing task")
            game_ids_need_to_judge = log.execute_info["game_ids_need_to_judge"]

        if len(game_ids_need_to_judge) == 0:
            logger.info("[JudgeGameTask.process_task] No games need judging")
            return await a_move_state(
                log, MissionChainState.DONE, "No game need to judge"
            )

        for idx in range(
            log.execute_info["processing_idx"], len(game_ids_need_to_judge)
        ):
            log.execute_info["processing_idx"] = idx
            game_id = game_ids_need_to_judge[idx]
            logger.info(
                f"[JudgeGameTask.process_task] Processing game {game_id} ({idx + 1}/{len(game_ids_need_to_judge)})"
            )

            if not _is_judge_game_pending(game_id):
                logger.info(
                    f"[JudgeGameTask.process_task] Game {game_id} is not pending, skipping"
                )
                continue

            if not _try_acquire_judge_game_lock(game_id):
                logger.info(
                    f"[JudgeGameTask.process_task] Game {game_id} is being processed by another instance, skipping..."
                )
                continue

            # Check if tweet exists and matches game ID
            resp: Response[ExtendedTweetInfoDto] = await sync2async(
                twitter_v2.get_tweet_info_from_tweet_id
            )(game_id)

            if (
                resp.is_error()
                or resp.data is None
                or resp.data.tweet_info is None
            ):
                logger.info(
                    f"[JudgeGameTask.process_task] Tweet {game_id} not found or error getting info, removing from cache"
                )
                # _mark_judge_game_done(game_id) # Remove from cache since tweet doesn't exist
                # maybe don't remove yet? it could be an error by server
                continue
            elif (
                resp.data.tweet_info.tweet_object.to_dict().get("tweet_id")
                != game_id
            ):
                logger.info(
                    f"[JudgeGameTask.process_task] Tweet {game_id} does not match game ID, removing from cache"
                )
                _mark_judge_game_done(
                    game_id
                )  # Remove from cache since tweet doesn't exist
                continue

            _mark_judge_game_running(game_id)

            answers, err = await _get_final_answers(log, game_id)

            if err is not None:
                logger.error(
                    f"[JudgeGameTask.process_task] Error getting final answers for game {game_id}: {err}"
                )
                _mark_judge_game_pending(game_id)
                return await a_move_state(
                    log, MissionChainState.ERROR, "Judge game failed"
                )

            # Handle case with no answers or single participant
            tx_hash = None
            winning_agent = None
            if len(answers) == 0:
                logger.info(
                    f"[JudgeGameTask.process_task] No answers for game {game_id}"
                )
                err = await _handle_winner_without_llm(log, game_id, None)
            elif len(answers) == 1:
                # Single participant case - no competition possible
                winning_agent = answers[0]["username"]
                logger.info(
                    f"[JudgeGameTask.process_task] Single participant for game {game_id}: {winning_agent}"
                )
                err = await _handle_winner_without_llm(
                    log, game_id, winning_agent
                )
            else:
                # Create conversation for judging
                logger.info(
                    f"[JudgeGameTask.process_task] Multiple participants for game {game_id}, using LLM"
                )
                winning_agent, tx_hash, err = await _handle_winner_with_llm(
                    log, self.llm, game_id, answers
                )

            if err is not None:
                logger.error(
                    f"[JudgeGameTask.process_task] Error handling winner for game {game_id}: {err}"
                )
                _mark_judge_game_pending(game_id)
                return await a_move_state(
                    log, MissionChainState.ERROR, "Judge game failed"
                )

            _mark_judge_game_done(game_id)
            log.execute_info["task_result"].append(
                {
                    "tweet_id": game_id,
                    "winning_agent": winning_agent,
                    "tx_hash": tx_hash,
                }
            )
            log = await self.commit_log(log)

        logger.info(
            "[JudgeGameTask.process_task] Successfully completed all games"
        )
        return await a_move_state(
            log, MissionChainState.DONE, "Judge game done"
        )
