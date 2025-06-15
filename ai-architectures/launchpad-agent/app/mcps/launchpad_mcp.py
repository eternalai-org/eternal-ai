from fastmcp import FastMCP
from app.utils.launchpad_api_calls import search_launchpad, get_launchpad_detail
from typing import List, Dict, Any, Optional
import re
import logging

logger = logging.getLogger(__name__)

mcp = FastMCP(name="launchpad_enhanced")


async def _search_projects_by_keywords(keywords: List[str], max_results: int = 5, network_id: str = "11155111") -> Dict[str, Any]:
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
        unique_results = _deduplicate_and_rank_results(all_results, keywords)
        
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
async def search_projects_by_keywords(keywords: List[str], max_results: int = 5, network_id: str = "11155111") -> Dict[str, Any]:
    """Search launchpad projects using keywords extracted from tweet content with advanced matching"""
    return await _search_projects_by_keywords(keywords, max_results, network_id)

async def _search_launchpad_simple(query: str, network_id: str = "11155111") -> List[dict]:
    """Simple search wrapper that returns basic project information"""
    res = await search_launchpad(query, network_id)

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
    name="search_launchpad", 
    description="Search for a launchpad project by name",
    annotations={
        "query": "Query the name of the project"
    }
)
async def search_launchpad_simple(query: str, network_id: str = "11155111") -> List[dict]:
    """Simple search wrapper that returns basic project information"""
    return await _search_launchpad_simple(query, network_id)



async def _get_launchpad_detail_simple(id: str, network_id: str = "11155111") -> Optional[dict]:
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
async def get_launchpad_detail_simple(id: str, network_id: str = "11155111") -> Optional[dict]:
    """Simple detail wrapper that returns basic project information"""
    return await _get_launchpad_detail_simple(id, network_id)


async def _extract_project_keywords_from_tweet(tweet_content: str) -> Dict[str, Any]:
    """Extract keywords that might match launchpad projects"""
    
    try:
        # Clean the tweet content
        cleaned_content = _clean_tweet_content(tweet_content)
        
        # Extract different types of keywords
        token_symbols = _extract_token_symbols(cleaned_content)
        project_names = _extract_potential_project_names(cleaned_content)
        technology_keywords = _extract_technology_keywords(cleaned_content)
        company_names = _extract_company_names(cleaned_content)
        
        # Combine all keywords
        all_keywords = []
        all_keywords.extend(token_symbols)
        all_keywords.extend(project_names)
        all_keywords.extend(technology_keywords)
        all_keywords.extend(company_names)
        
        # Remove duplicates while preserving order
        unique_keywords = list(dict.fromkeys(all_keywords))
        
        return {
            "all_keywords": unique_keywords,
            "token_symbols": token_symbols,
            "project_names": project_names,
            "technology_keywords": technology_keywords,
            "company_names": company_names,
            "cleaned_content": cleaned_content
        }
        
    except Exception as e:
        logger.error(f"Error extracting keywords: {e}", exc_info=True)
        return {
            "all_keywords": [],
            "token_symbols": [],
            "project_names": [],
            "technology_keywords": [],
            "company_names": [],
            "cleaned_content": tweet_content,
            "error": str(e)
        }
        
@mcp.tool(
    name="extract_project_keywords_from_tweet",
    description="Extract potential project-related keywords from tweet content using NLP techniques",
    annotations={
        "tweet_content": "The tweet text to analyze for project-related keywords"
    }
)
async def extract_project_keywords_from_tweet(tweet_content: str) -> Dict[str, Any]:
    """Extract potential project-related keywords from tweet content using NLP techniques"""
    return await _extract_project_keywords_from_tweet(tweet_content)


async def _get_project_details_for_matching(project_id: str, network_id: str = "11155111") -> Dict[str, Any]:
    """Get project details to help with matching analysis"""
    
    try:
        detail_result = await get_launchpad_detail(project_id, network_id)
        
        if detail_result.result:
            project = detail_result.result
            return {
                "id": project.id,
                "name": project.name,
                "description": project.description,
                "success": True
            }
        else:
            return {
                "id": project_id,
                "error": detail_result.error or "Project not found",
                "success": False
            }
            
    except Exception as e:
        logger.error(f"Error getting project details for {project_id}: {e}")
        return {
            "id": project_id,
            "error": str(e),
            "success": False
        }
        
@mcp.tool(
    name="get_project_details_for_matching",
    description="Get detailed information about a specific project for matching analysis",
    annotations={
        "project_id": "The launchpad project ID to get details for"
    }
)
async def get_project_details_for_matching(project_id: str, network_id: str = "11155111") -> Dict[str, Any]:
    """Get detailed information about a specific project for matching analysis"""
    return await _get_project_details_for_matching(project_id, network_id)

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

def _clean_tweet_content(content: str) -> str:
    """Clean tweet content for keyword extraction"""
    # Remove URLs
    content = re.sub(r'http[s]?://\S+', '', content)
    
    # Remove excessive whitespace
    content = re.sub(r'\s+', ' ', content)
    
    # Remove mentions and hashtags for cleaner analysis
    content = re.sub(r'@\w+', '', content)
    content = re.sub(r'#(\w+)', r'\1', content)  # Keep hashtag content without #
    
    return content.strip()

def _extract_token_symbols(content: str) -> List[str]:
    """Extract potential token symbols ($TOKEN format)"""
    # Match $SYMBOL pattern
    symbols = re.findall(r'\$([A-Z]{2,10})', content)
    
    # Also look for symbols in ALL CAPS that might be tokens
    words = content.split()
    for word in words:
        # Remove punctuation
        clean_word = re.sub(r'[^\w]', '', word)
        if (len(clean_word) >= 2 and len(clean_word) <= 6 and 
            clean_word.isupper() and clean_word.isalpha()):
            symbols.append(clean_word)
    
    return list(set(symbols))  # Remove duplicates

def _extract_potential_project_names(content: str) -> List[str]:
    """Extract words that could be project names"""
    # Look for capitalized words (potential project names)
    words = re.findall(r'\b[A-Z][a-z]+(?:\s+[A-Z][a-z]+)*\b', content)
    
    # Filter out common words
    common_words = {'The', 'And', 'Or', 'But', 'In', 'On', 'At', 'To', 'For', 'Of', 'With', 'By'}
    project_names = [word for word in words if word not in common_words and len(word) > 2]
    
    return project_names

def _extract_technology_keywords(content: str) -> List[str]:
    """Extract technology-related keywords"""
    tech_keywords = [
        'DeFi', 'NFT', 'GameFi', 'SocialFi', 'AI', 'ML', 'blockchain', 'crypto',
        'Web3', 'DAO', 'DEX', 'yield', 'farming', 'staking', 'lending',
        'metaverse', 'virtual', 'augmented', 'reality', 'AR', 'VR',
        'layer1', 'layer2', 'L1', 'L2', 'rollup', 'sidechain',
        'oracle', 'bridge', 'cross-chain', 'interoperability',
        'governance', 'token', 'coin', 'protocol', 'platform'
    ]
    
    content_lower = content.lower()
    found_keywords = []
    
    for keyword in tech_keywords:
        if keyword.lower() in content_lower:
            found_keywords.append(keyword)
    
    return found_keywords

def _extract_company_names(content: str) -> List[str]:
    """Extract potential company/organization names"""
    # Look for words ending in common company suffixes
    company_patterns = [
        r'\b\w+\s*(?:Inc|Corp|LLC|Ltd|Foundation|Labs|Protocol|Network|Finance|Capital|Ventures)\b',
        r'\b[A-Z]\w+\s+[A-Z]\w+\b'  # Two capitalized words (likely company names)
    ]
    
    companies = []
    for pattern in company_patterns:
        matches = re.findall(pattern, content, re.IGNORECASE)
        companies.extend(matches)
    
    return companies 
