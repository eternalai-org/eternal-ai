from .utils import get_and_warn

GAME_BASE_URL = get_and_warn("GAME_BASE_URL")

GAME_EMOJIS = ["🎮", "🎯", "🎪", "🎲", "🤔"]
GAME_KEYWORDS = ["game", "play", "guess", "riddle"]
FACT_EMOJIS = ["🎮", "🎯", "🎪", "🎲", "🤔"]
FACT_KEYWORDS = [
    "fact",
    "trivia",
    "did you know",
    "learn",
    "today i learned",
    "til",
]  # TODO: update this
# GAME_DURATION = 60 * 60 * 4 # 4 hours
GAME_DURATION = 60 * 10  # 10 minutes
GAME_BET_DURATION = 60 * 10  # should be the same as GAME_DURATION
# FACT_DURATION = 60 * 60 * 24  # 1 days
# FACT_BET_DURATION = 60 * 60  # 1 hours

FACT_DURATION = 60 * 30  # 30 minutes
FACT_BET_DURATION = 60 * 10  # 10 minutes

GAME_REDIS_CACHE = 24 * 60 * 60 * 2  # 2 day
CREATE_GAME_LOCK_EXPIRY = 1000 * 10
# 10 seconds
JUDGE_GAME_LOCK_EXPIRY = 1000 * 60 * 60 * 2
# 2 hours

# Reply tweet templates
GAME_CREATED_TWEET = "A challenge is created! Place your bet by deposit EAI on Base to the wallet of who you think will win:\n\nBet will close in {bet_hours:02d} hours {bet_minutes:02d} minutes\nGame will end in {hours:02d} hours {minutes:02d} minutes\n\n"
# Constants for tweet replies
WINNER_TWEET_TEMPLATE = "The winner is {}"
NO_WINNER_TWEET = "No winner"

CREATE_GAME_PREFIX_PATTERN = "[CREATE_GAME]"  # TODO: To be defined
