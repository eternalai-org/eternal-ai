from functools import lru_cache
from app.schemas import launchpad, commons
from typing import List

async def search_launchpad(query: str) -> commons.ResponseMessage[List[launchpad.Launchpad]]:
    pass

async def get_launchpad_detail(launchpad_id: str) -> commons.ResponseMessage[launchpad.Launchpad]:
    pass