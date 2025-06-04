from fastapi import APIRouter
from fastapi import Request, Response
import logging
import time
from typing import Callable

logger = logging.getLogger(__name__)

router = APIRouter(prefix="/api")

@router.post("/evaluate")
async def evaluate(request: Request):
    data = await request.json()
    logger.info(f"Evaluating {data}")
    return {}
