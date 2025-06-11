from fastmcp import FastMCP
from app.utils.launchpad_api_calls import (
    get_launchpad_detail as get_launchpad_detail_api, 
    search_launchpad as search_launchpad_api
)
from typing import List, Optional

mcp = FastMCP(
    name="launchpad_info",
    description="Explore launchpad projects",
)

@mcp.tool(
    name="search_launchpad", 
    description="Search for a launchpad project",
    annotations={
        "query": "query the name of the project"
    }
)
async def search_launchpad(query: str) -> List[dict]:
    # TODO: filter eneded launchpad project before returning results
    res = await search_launchpad_api(query)

    if res.result is not None:
        return [
            {
                "id": item.id,
                "name": item.name,
                "description": item.description
            }
            for item in res.result
        ]
        
    return []

@mcp.tool(
    name="get_launchpad_detail",
    description="Get the detail of a launchpad project",
    annotations={
        "id": "the id of the launchpad project"
    }
)
async def get_launchpad_detail(id: str) -> Optional[dict]:
    res = await get_launchpad_detail_api(id)

    if res.result is not None:
        return {
            "id": res.result.id,
            "name": res.result.name,
            "description": res.result.description
        }

    return None

@mcp.tool(
    name="get_investment_history",
    description="Get the investment history of a user",
    annotations={
        "user_id": "the id of the user"
    }
)
async def get_investment_history(user_id: str):
    # TODO: implement this
    return {
        "user_id": user_id,
        "investment_history": []
    }