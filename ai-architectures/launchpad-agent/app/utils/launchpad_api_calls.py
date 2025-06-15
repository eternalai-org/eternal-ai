from app.schemas import launchpad, commons
from typing import List
import os
import httpx
from app.schemas import commons, launchpad
from app.config import settings

LAUNCHPAD_API_URL = settings.launchpad_api_url
LAUNCHPAD_API_KEY = settings.launchpad_api_key

TIMEOUT_CFG = httpx.Timeout(60.0, connect=10.0) 

async def search_launchpad(
    query: str,
    network_id: str,
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
            resp = await client.get(url, params={"search": query, "network_id": network_id}) 
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
    network_id: str,
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
            resp = await client.get(url, params={"network_id": network_id})
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

async def join_launchpad(
    launchpad_id: str,
    network_id: str,
    twitter_id: str,
    max_cap: str,
    tweet_id: str,
    tweet_content: str,
    launchpad_base_url: str = LAUNCHPAD_API_URL,
    launchpad_api_key: str = LAUNCHPAD_API_KEY,
) -> commons.ResponseMessage[launchpad.LaunchpadDepositInfo]:
    response_model = commons.ResponseMessage[launchpad.LaunchpadDepositInfo]
    url = f"{launchpad_base_url}/join"

    async with httpx.AsyncClient(
        base_url=launchpad_base_url,
        headers={
            "Authorization": f"Bearer {launchpad_api_key}"
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            resp = await client.post(url, json={
                "launchpad_id": launchpad_id,
                "twitter_id": twitter_id,
                "max_cap": max_cap,
                "tweet_id": tweet_id,
                "tweet_content": tweet_content,
                "network_id": network_id
            })
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
            return response_model(
                status=commons.APIStatus.OK,
                result=launchpad.LaunchpadDepositInfo.model_validate(resp_json)
            )
        except Exception as e:
            return response_model(
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )

async def get_investment_history(
    twitter_id: str,
    network_id: str,
    launchpad_base_url: str = LAUNCHPAD_API_URL,
    launchpad_api_key: str = LAUNCHPAD_API_KEY,
) -> commons.ResponseMessage[List[launchpad.Launchpad]]:
    response_model = commons.ResponseMessage[List[launchpad.Launchpad]]
    url = f"{launchpad_base_url}/investment-history"

    async with httpx.AsyncClient(
        base_url=launchpad_base_url,
        headers={
            "Authorization": f"Bearer {launchpad_api_key}",
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            resp = await client.get(
                url, 
                params={
                    "twitter_id": twitter_id, 
                    "network_id": network_id
                }
            )
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

async def reply(
    tweet_id: str,
    content: str,
    launchpad_base_url: str = LAUNCHPAD_API_URL,
    launchpad_api_key: str = LAUNCHPAD_API_KEY,
) -> commons.ResponseMessage[bool]:
    response_model = commons.ResponseMessage[bool]
    url = f"{launchpad_base_url}/reply-tweet"

    async with httpx.AsyncClient(
        base_url=launchpad_base_url,
        headers={
            "Authorization": f"Bearer {launchpad_api_key}"
        },
        timeout=TIMEOUT_CFG,
    ) as client:
        try:
            resp = await client.post(url, json={
                "reply_tweet_id": tweet_id,
                "content": content
            })
        except Exception as e:
            return response_model(
                result=False,
                status=commons.APIStatus.ERROR, 
                error=str(e)
            )
            
        if resp.status_code != 200:
            return response_model(
                result=False,
                status=commons.APIStatus.ERROR, 
                error=resp.text
            )

        return response_model(
            result=True,
            status=commons.APIStatus.OK,
        )