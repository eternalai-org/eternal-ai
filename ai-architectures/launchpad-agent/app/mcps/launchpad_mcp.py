from fastmcp import FastMCP
from app.utils.launchpad_api_calls import search_launchpad, get_launchpad_detail
from typing import List, Dict, Any, Optional
import re
import logging

logger = logging.getLogger(__name__)

mcp = FastMCP(name="launchpad_enhanced")


async def _search_projects_by_keywords(keywords: List[str], max_results: int = 5, network_id: str = "8453") -> Dict[str, Any]:
    """Enhanced search that tries multiple keyword combinations and ranks results"""
    
    try:
        all_results = []
        search_terms_used = []
        
        # Search with individual keywords
        for keyword in keywords:
            if len(keyword.strip()) < 2:  # Skip very short keywords
                continue
                
            try:
                search_result = await search_launchpad(keyword.strip(), network_id)

                if search_result.result:
                    all_results.extend(search_result.result)
                    search_terms_used.append(keyword.strip())
                    logger.info(f"Found {len(search_result.result)} results for keyword: {keyword}")

                elif search_result.error:
                    logger.error(f"No results found for keyword: {keyword} (msg: {search_result.error})")

            except Exception as e:
                logger.warning(f"Search failed for keyword '{keyword}': {e}")
                continue
        
        # Search with combined keywords (for multi-word project names)
        if len(keywords) > 1:
            combined_query = " ".join(keywords[:5])  # Limit to first 3 keywords

            try:
                search_result = await search_launchpad(combined_query, network_id)
                if search_result.result:
                    all_results.extend(search_result.result)
                    search_terms_used.append(combined_query)
                    logger.info(f"Found {len(search_result.result)} results for combined query: {combined_query}")
            except Exception as e:
                logger.warning(f"Combined search failed for '{combined_query}': {e}")
        
        # Deduplicate and rank results
        print(all_results)
        unique_results = _deduplicate_and_rank_results(all_results, keywords)
        print(unique_results)
        
        return {
            "projects": unique_results[:max_results],
            "total_found": len(unique_results),
            "search_terms_used": search_terms_used,
            "original_keywords": keywords
        }
        
    except Exception as e:
        logger.error(f"Error in search_projects_by_keywords: {e}", exc_info=True)
        return {
            "projects": [],
            "total_found": 0,
            "search_terms_used": [],
            "original_keywords": keywords,
            "error": str(e)
        }

@mcp.tool(
    name="search_projects_by_keywords",
    description="Search launchpad projects using keywords extracted from tweet content with advanced matching",
    annotations={
        "keywords": "List of keywords to search for in project names and descriptions",
        "max_results": "Maximum number of results to return (default: 5)"
    }
)
async def search_projects_by_keywords(keywords: List[str], max_results: int = 5, network_id: str = "8453") -> Dict[str, Any]:
    """Search launchpad projects using keywords extracted from tweet content with advanced matching"""
    return await _search_projects_by_keywords(keywords, max_results, network_id)

async def _get_launchpad_detail_simple(id: str, network_id: str = "8453") -> Optional[dict]:
    """Simple detail wrapper that returns basic project information"""
    res = await get_launchpad_detail(id, network_id)

    if res.result is not None:
        return {
            "id": res.result.id,
            "name": res.result.name,
            "description": res.result.description
        }

    return None

@mcp.tool(
    name="get_launchpad_detail",
    description="Get the detail of a launchpad project",
    annotations={
        "id": "The id of the launchpad project"
    }
)
async def get_launchpad_detail_simple(id: str, network_id: str = "8453") -> Optional[dict]:
    """Simple detail wrapper that returns basic project information"""
    return await _get_launchpad_detail_simple(id, network_id)

def _deduplicate_and_rank_results(results: List[Any], keywords: List[str]) -> List[Dict[str, Any]]:
    """Deduplicate results and rank them by relevance to keywords"""
    
    if not results:
        return []
    
    # Deduplicate by ID
    seen_ids = set()
    unique_results = []
    
    for result in results:
        result_id = getattr(result, 'id', None)
        if result_id and result_id not in seen_ids:
            seen_ids.add(result_id)
            unique_results.append(result)
    
    # Rank by relevance
    ranked_results = []
    
    for result in unique_results:
        relevance_score = _calculate_relevance_score(result, keywords)
        
        ranked_results.append({
            "id": getattr(result, 'id', ''),
            "name": getattr(result, 'name', ''),
            "description": getattr(result, 'description', ''),
            "relevance_score": relevance_score
        })
    
    # Sort by relevance score (descending)
    ranked_results.sort(key=lambda x: x['relevance_score'], reverse=True)
    
    return ranked_results

def _calculate_relevance_score(result: Any, keywords: List[str]) -> float:
    """Calculate how relevant a project is to the given keywords"""
    
    name = getattr(result, 'name', '').lower()
    description = getattr(result, 'description', '').lower()
    
    score = 0.0
    
    for keyword in keywords:
        keyword_lower = keyword.lower()
        
        # Exact match in name (highest score)
        if keyword_lower == name:
            score += 10.0
        elif keyword_lower in name:
            score += 5.0
        
        # Match in description
        if keyword_lower in description:
            score += 2.0
        
        # Partial matches
        if any(keyword_lower in word for word in name.split()):
            score += 3.0
        
        if any(keyword_lower in word for word in description.split()):
            score += 1.0
    
    return score
