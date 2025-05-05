from x_content.wrappers.analytic.custom import send_log_to_custom_analytic
from x_content.wrappers.magic import helpful_raise_for_status
from .log_decorators import log_function_call
from typing import List
import logging
import requests
import json
from x_content import constants as const

logger = logging.getLogger(__name__)


@log_function_call
def search_from_bing(
    query: str, top_k: int = 1, task_name: str = "N/A"
) -> List[str]:
    """
    Retrieves and integrates up-to-date information on topics where recent events or developments (2024 onward)
    are essential to accurately address the user's query.

    This function should only be called when the knowledge cutoff (01 Sep 2022) leaves gaps that would
    prevent a comprehensive and accurate response. Retrieved insights should be directly relevant, carefully
    vetted, and seamlessly incorporated with existing knowledge to provide a clear, precise, and well-informed answer.

    Args:
        query (str): The search query to retrieve insights for.
        top_k (int): The number of top search results to retrieve.

    Returns:
        str: A consolidated summary of insights derived from the search results.
    """

    send_log_to_custom_analytic(
        "search_from_bing",
        params=[
            {"key": "query", "value": query},
            {"key": "top_k", "value": top_k},
            {"key": "task_name", "value": task_name},
        ],
        append_uuid=True,
    )

    url = "https://api.bing.microsoft.com/v7.0/news/search"

    headers = {"Ocp-Apim-Subscription-Key": const.BING_SEARCH_API_KEY}

    params = {
        "q": query,
        "count": top_k,
        "freshness": "Day",
        "safeSearch": "Moderate",
        "setLang": "en",
        "mkt": "en-US",
        "sortBy": "Relevance",
    }

    try:
        response = requests.get(url, headers=headers, params=params)
        helpful_raise_for_status(response)
        data = response.json().get("value", [])
        descriptions = [item.get("description", "") for item in data]
        final_content = descriptions if descriptions else []
        return final_content

    except requests.exceptions.RequestException as e:
        logger.error(f"[search_from_bing] Request failed: {e}")
        return []

    except json.JSONDecodeError as e:
        logger.error(f"[search_from_bing] JSON decode error: {e}")
        return []

    except Exception as e:
        logger.error(f"[search_from_bing] An unexpected error occurred: {e}")
        return []
