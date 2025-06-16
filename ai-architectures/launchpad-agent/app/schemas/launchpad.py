from pydantic import BaseModel
from typing import Literal
from typing import Optional

CURRENCY_UNIT = "EAI"

class Launchpad(BaseModel):
    id: int
    twitter_post_id: int
    tweet_id: str
    name: str
    description: str
    twitter_id: str
    twitter_username: str
    twitter_name: str
    address: str
    status: str
    start_at: str
    end_at: str
    finished_at: str
    fund_balance: str
    total_balance: str
    token_address: str
    token_name: str
    token_symbol: str
    token_image_url: str
    total_supply: str
    tge_balance: str
    airdrop_balance: Optional[str] = None
    liquidity_balance: str
    team_balance: str
    max_fund_balance: str
    refund_balance: str
    start_tweet_id: str
    end_tweet_id: str
    price_usd: str
    price_eai: str
    market_cap_usd: str


class LaunchpadDepositInfo(BaseModel):
    status: Literal["pending", "success", "failed"]
    eth_address: str