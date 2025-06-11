from functools import lru_cache
from app.schemas import launchpad, commons
from typing import List
import os
import httpx
from app.schemas import commons, launchpad

LAUNCHPAD_API_URL = os.getenv("LAUNCHPAD_API_URL")
LAUNCHPAD_API_KEY = os.getenv("LAUNCHPAD_API_KEY", "super-secret")

TIMEOUT_CFG = httpx.Timeout(60.0, connect=10.0) 

async def search_launchpad(
    query: str,
    launchpad_base_url: str = LAUNCHPAD_API_URL,
    launchpad_api_key: str = LAUNCHPAD_API_KEY,
) -> commons.ResponseMessage[List[launchpad.Launchpad]]:
    response_model = commons.ResponseMessage[List[launchpad.Launchpad]]
    url = f"{launchpad_base_url}/list"

    async with httpx.AsyncClient(
        base_url=launchpad_base_url,
        headers={
            "Authorization": f"Bearer {launchpad_api_key}"
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            resp = await client.get(url, params={"search": query}) 
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

        if resp.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=resp.text
            )
            
        resp_json = resp.json()
        
        if 'result' not in resp_json:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error="No result found"
            )
            
        try:
            result = [
                launchpad.Launchpad.model_validate(item)
                for item in resp_json['result']
            ]

            return response_model(
                status=commons.APIStatus.OK,
                result=result
            )

        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )


async def get_launchpad_detail(
    launchpad_id: str,
    launchpad_base_url: str = LAUNCHPAD_API_URL,
    launchpad_api_key: str = LAUNCHPAD_API_KEY,
) -> commons.ResponseMessage[launchpad.Launchpad]:
    response_model = commons.ResponseMessage[launchpad.Launchpad]
    url = f"{launchpad_base_url}/detail/{launchpad_id}"

    async with httpx.AsyncClient(
        base_url=launchpad_base_url,
        headers={
            "Authorization": f"Bearer {launchpad_api_key}"
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            resp = await client.get(url)
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )
            
        if resp.status_code != 200:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=resp.text
            )

        resp_json = resp.json()

        if 'result' not in resp_json:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error="No result found"
            )

        try:
            result = launchpad.Launchpad.model_validate(resp_json['result'])
            
            return response_model(
                status=commons.APIStatus.OK,
                result=result
            )

        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )