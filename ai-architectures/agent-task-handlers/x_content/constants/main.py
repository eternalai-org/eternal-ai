from enum import Enum
from x_content.constants.utils import get_and_warn


class AgentUsername(str, Enum):
    CRYPTOCOMIC_AI = "YorelisDiaz1"


class AgentFullname(str, Enum):
    CRYPTOCOMIC_AI = "CryptoComic AI"


class ModelName(str, Enum):
    INTELLECT_10B = "PrimeIntellect/INTELLECT-1-Instruct"
    DEEPSEEK_R1 = "DeepSeek-R1"
    DEEPSEEK_V3 = "DeepSeek V3"
    DEEPSEEK_R1_DISTILL_LLAMA_70B = "DeepSeek-R1-Distill-Llama-70B"


class HTTPMethod(str, Enum):
    GET = "GET"
    POST = "POST"


class ToolSet(str, Enum):
    DEFAULT = "default"
    NOT_SPECIFIED = "not_specified"
    REPLY_NON_MENTIONS = "reply_non_mentions"
    REPLY_NON_MENTIONS_TRADITIONAL = "reply_non_mentions_traditional"
    FOLLOW = "follow"
    POST = "post"
    QUOTE_TWEET = "quote_tweet"
    TRADING = "trading"
    INSCRIBE_TWEET = "inscribe_tweet"
    ISSUE_TOKEN = "issue_token"
    INSCRIBE_REPLY = "inscribe_reply"


class AgentTask(str, Enum):
    REACT_AGENT = "react_agent"
    POST_V2 = "post_search_v2"
    POST_V3 = "post_search_v3"
    REPLY = "reply"
    QUOTE_TWEET = "quote_tweet"
    SHADOW_REPLY = "shadow_reply"
    TRADING = "trading"
    POST_TWEET = "post_tweet"
    POST_SEARCH = "post_search"
    CREATE_GAME = "create_game"
    JUDGE_GAME = "judge_game"
    DEFAULT = "default"


class MissionChainState(str, Enum):
    NEW = "new"
    RUNNING = "running"
    DONE = "done"
    ERROR = "error"


USE_RQ = get_and_warn("USE_RQ") or True

if isinstance(USE_RQ, str):
    USE_RQ = USE_RQ.lower() in ["true", "1"]


REDIS_LOG_BASE_KEY = get_and_warn(
    "REDIS_LOG_BASE_KEY", "ReactAgentReasoningLog-state-1"
)
REDIS_CHAT_REQUEST_BASE_KEY = get_and_warn(
    "REDIS_CHAT_REQUEST_BASE_KEY", "ChatRequest-state-1"
)
APP_ENV = get_and_warn("APP_ENV")
RUN_SERVICE_V2: bool = get_and_warn("RUN_SERVICE_V2") or False

if isinstance(RUN_SERVICE_V2, str):
    RUN_SERVICE_V2 = RUN_SERVICE_V2.lower() in ["true", "1"]

BACKEND_API = get_and_warn("BACKEND_API")
BACKEND_AUTH_TOKEN = get_and_warn("BACKEND_AUTH_TOKEN")

API_SECRET_TOKEN = get_and_warn("API_SECRET_TOKEN", "supersecret")
SERVER_HOST: str = get_and_warn("SERVER_HOST", "0.0.0.0")
SERVER_PORT: int = get_and_warn("SERVER_PORT") or 8000

if isinstance(SERVER_PORT, str):
    SERVER_PORT = int(SERVER_PORT)

MINIMUM_POST_LENGTH = 32
DEFAULT_REACT_MAX_STEPS = 10
NUM_OF_TWEETS_TO_POST = 1

REACT_MODELS_BLACKLIST = [
    ModelName.INTELLECT_10B,
]

REPLY_MODELS_BLACKLIST = [
    ModelName.INTELLECT_10B,
]

ALL_BLACKLIST = [
    "itsmechaseb",
]

GAME_TASKS_WHITELIST = [
    AgentUsername.CRYPTOCOMIC_AI,
]

GCP_MEME_IMAGES_FOLDER = get_and_warn("GCP_MEME_IMAGES_FOLDER")
GCP_BUCKET_NAME = get_and_warn("GCP_BUCKET_NAME")
GCS_TWIN_BUCKET = get_and_warn("GCS_TWIN_BUCKET")

TWITTER_API_URL = get_and_warn("TWITTER_API_URL", "").rstrip("/")
TWITTER_API_KEY = get_and_warn("TWITTER_API_KEY")

FETCH_NEWS_INTERVAL_SECONDS = (
    get_and_warn("FETCH_NEWS_INTERVAL_SECONDS") or 60 * 60 * 2
)

if isinstance(FETCH_NEWS_INTERVAL_SECONDS, str):
    FETCH_NEWS_INTERVAL_SECONDS = int(FETCH_NEWS_INTERVAL_SECONDS)

BING_SEARCH_API_KEY = get_and_warn("BING_SEARCH_API_KEY")
IS_MOCKING_ACTION: bool = get_and_warn("IS_MOCKING_ACTION") or True

if isinstance(IS_MOCKING_ACTION, str):
    IS_MOCKING_ACTION = IS_MOCKING_ACTION.lower() in ["true", "1"]

DISABLE_FETCH_NEWS: bool = get_and_warn("DISABLE_FETCH_NEWS") or False

if isinstance(DISABLE_FETCH_NEWS, str):
    DISABLE_FETCH_NEWS = DISABLE_FETCH_NEWS.lower() in ["true", "1"]


DISABLE_LOG_FUNCTION_CALL: bool = (
    get_and_warn("DISABLE_LOG_FUNCTION_CALL") or False
)

if isinstance(DISABLE_LOG_FUNCTION_CALL, str):
    DISABLE_LOG_FUNCTION_CALL = DISABLE_LOG_FUNCTION_CALL.lower() in [
        "true",
        "1",
    ]

MAX_TEXT_LENGTH = 10000
MIN_TEXT_LENGTH_TO_SUMMARIZE = 100

APP_NAME = get_and_warn("APP_NAME")
