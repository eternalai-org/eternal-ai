from typing import Callable
import asyncio
import time
import traceback
import logging

logger = logging.getLogger(__name__)

def float_clamp(value: float | str, min_value: float = 0.0, max_value: float = 1.0, default_value: float = 0.0) -> float:

    try:
        value = float(value)
    except Exception:
        value = default_value

    return max(min(value, max_value), min_value)

def dsu(n: int, relations: list[tuple[int, int]]) -> list[int]:
    parent = list(range(n))
    rank = [0] * n

    def find(x: int) -> int:
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def unite(x: int, y: int) -> None:
        x_root = find(x)
        y_root = find(y)

        if x_root == y_root:
            return

        if rank[x] < rank[y]:
            parent[x_root] = y_root
        
        elif rank[x] > rank[y]:
            parent[y_root] = x_root
        
        else:
            parent[y_root] = x_root
            rank[x_root] += 1

    for x, y in relations:
        unite(x, y)

    return parent
    

def retry(func: Callable, max_retry=5, first_interval=10, interval_multiply=1):
    def sync_wrapper(*args, **kwargs):
        interval = first_interval
        for iter in range(max_retry + 1):
            try:
                result = func(*args, **kwargs)
                return result
            except Exception as err:
                traceback.print_exc()
                logger.error(
                    f"Function {func.__name__} failed with error '{err}'. Retry attempt {iter}/{max_retry}"
                )

                if iter == max_retry:
                    raise err

            time.sleep(interval)
            interval *= interval_multiply

    async def async_wrapper(*args, **kwargs):
        interval = first_interval
        for iter in range(max_retry + 1):
            try:
                result = await func(*args, **kwargs)
                return result
            except Exception as err:
                traceback.print_exc()
                logger.error(
                    f"Function {func.__name__} failed with error '{err}'. Retry attempt {iter}/{max_retry}"
                )
                
                if iter == max_retry:
                    raise err
                
            await asyncio.sleep(interval)
            interval *= interval_multiply

    return async_wrapper if asyncio.iscoroutinefunction(func) else sync_wrapper
