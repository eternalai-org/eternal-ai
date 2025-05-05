import json
import logging
import uuid
import requests
import time
from x_content import constants as const
from x_content.wrappers.magic import helpful_raise_for_status

logger = logging.getLogger(__name__)


def send_log_to_custom_analytic(
    event_name: str, params: dict, append_uuid=False
):
    try:
        if append_uuid:
            event_name += ":" + str(uuid.uuid4())
        logger.info(
            f"[send_log_to_custom_analytic] Sending event {event_name} with params {json.dumps(params)}"
        )

        url = f"{const.CUSTOM_ANALYTIC_URL}/api/v1/event_tracking"
        headers = {"Authorization": const.CUSTOM_ANALYTIC_API_KEY}
        payload = {
            "event_name": event_name,
            "event_timestamp": time.time() * 1000,
            "data": {
                "platform": const.APP_NAME,
                "event_params": params,
            },
        }

        resp = requests.post(url=url, headers=headers, json=payload)
        helpful_raise_for_status(resp)
    except Exception as err:
        logger.error(
            f"[send_log_to_custom_analytic] An unexpected error occured: {err}"
        )
