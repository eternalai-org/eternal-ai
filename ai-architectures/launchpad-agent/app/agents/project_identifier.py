import json
from typing import Optional, Dict, List, Any
from app.agents.mcp_agent import mcp_agent_run
from app.schemas.evaluation import ProjectIdentification
from app.utils.lm import get_oai_async_client, get_model_id
import logging

logger = logging.getLogger(__name__)

async def identify_launchpad_project(tweet_content: str, tweet_id: str = None, network_id: str = "11155111") -> ProjectIdentification:
    """
    Stage 2: Identify which launchpad project the tweet is about
    
    Args:
        tweet_content: The text content of the tweet
        tweet_id: Optional tweet ID for tracking
        
    Returns:
        ProjectIdentification with project ID if found, None otherwise
    """
    
    try:
        # Step 1: Extract keywords from tweet
        keyword_extraction = await _extract_keywords_with_ai(tweet_content)
        
        if not keyword_extraction.get("keywords"):
            logger.info(f"No keywords extracted from tweet: {tweet_content}...")
            return ProjectIdentification(
                tweet_id=tweet_id or "unknown",
                launchpad_id=None,
                confidence=0.0,
                reasoning="No project-related keywords found in tweet",
                keywords_matched=[]
            )
        
        # Step 2: Search for projects using extracted keywords
        search_results = await _search_projects_with_mcp(keyword_extraction["keywords"], network_id)
        
        if not search_results.get("projects"):
            logger.info(f"No projects found for keywords: {keyword_extraction['keywords']}")
            return ProjectIdentification(
                tweet_id=tweet_id or "unknown",
                launchpad_id=None,
                confidence=0.0,
                reasoning=f"No projects found matching keywords: {keyword_extraction['keywords']}",
                keywords_matched=keyword_extraction["keywords"]
            )
        
        # Step 3: Use AI to determine the best match
        best_match = await _ai_select_best_project_match(
            tweet_content, 
            keyword_extraction["keywords"],
            search_results["projects"]
        )
        
        return _create_project_identification(
            tweet_id or "unknown",
            best_match,
            keyword_extraction["keywords"],
            search_results["projects"]
        )
        
    except Exception as e:
        logger.error(f"Error identifying project for tweet: {e}", exc_info=True)
        return ProjectIdentification(
            tweet_id=tweet_id or "unknown",
            launchpad_id=None,
            confidence=0.0,
            reasoning=f"Error during project identification: {str(e)}",
            keywords_matched=[]
        )

async def _extract_keywords_with_ai(tweet_content: str) -> Dict[str, Any]:
    """Use AI to extract project-related keywords from tweet"""
    
    system_prompt = """You are an expert at extracting cryptocurrency and blockchain project keywords from social media posts.

Your task is to identify keywords that could help match the tweet to launchpad projects. Return a JSON response with this structure:

{
    "keywords": ["keyword1", "keyword2", "keyword3"],
    "token_symbols": ["BTC", "ETH", "TOKEN"],
    "project_names": ["ProjectName", "ProtocolName"],
    "technology_categories": ["DeFi", "NFT", "GameFi"],
    "confidence": 0.85
}

Extract:
1. Token symbols (like $BTC, $ETH, or symbols mentioned without $)
2. Potential project names (capitalized words that could be projects)
3. Technology categories (DeFi, NFT, GameFi, AI, Web3, etc.)
4. Company/protocol names (ending in Labs, Protocol, Network, etc.)
5. Specific blockchain/crypto terminology

Focus on keywords that would help identify specific projects, not general crypto terms."""

    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": f"Extract project keywords from this tweet:\n\n{tweet_content}"}
    ]
    
    client = get_oai_async_client()
    
    try:
        response = await client.chat.completions.create(
            model=get_model_id(),
            messages=messages,
            temperature=0.1,
            max_tokens=500
        )
        
        response_text = response.choices[0].message.content.strip()
        
        # Extract JSON from response
        json_match = _extract_json_from_text(response_text)
        if json_match:
            result = json.loads(json_match)
            
            # Combine all keywords into a single list
            all_keywords = []
            all_keywords.extend(result.get("keywords", []))
            all_keywords.extend(result.get("token_symbols", []))
            all_keywords.extend(result.get("project_names", []))
            all_keywords.extend(result.get("technology_categories", []))
            
            # Remove duplicates and empty strings
            unique_keywords = list(set([k.strip() for k in all_keywords if k.strip()]))
            
            return {
                "keywords": unique_keywords,
                "details": result,
                "confidence": result.get("confidence", 0.5)
            }
        else:
            # Fallback: basic keyword extraction
            return _fallback_keyword_extraction(tweet_content)
            
    except Exception as e:
        logger.error(f"Error in AI keyword extraction: {e}")
        return _fallback_keyword_extraction(tweet_content)

async def _search_projects_with_mcp(keywords: List[str], network_id: str) -> Dict[str, Any]:
    """Search for projects using the enhanced launchpad MCP"""
    
    try:
        from app.mcps.launchpad_mcp import _search_projects_by_keywords

        # Call the enhanced search tool
        search_result = await _search_projects_by_keywords(keywords, max_results=5, network_id=network_id)

        if search_result and isinstance(search_result, dict):
            return search_result
        else:
            logger.warning(f"Unexpected search result format: {search_result}")
            return {"projects": [], "total_found": 0}
            
    except Exception as e:
        logger.error(f"Error searching projects with MCP: {e}")
        return {"projects": [], "total_found": 0}

async def _ai_select_best_project_match(
    tweet_content: str,
    keywords: List[str], 
    projects: List[Dict[str, Any]]
) -> Dict[str, Any]:
    """Use AI to select the best matching project from search results"""
    
    if not projects:
        return {"selected_project": None, "confidence": 0.0, "reasoning": "No projects to match"}
    
    # Format projects for AI analysis
    projects_text = "\n".join([
        f"ID: {p.get('id', 'N/A')}, Name: {p.get('name', 'N/A')}, Description: {p.get('description', 'N/A')}..."
        for p in projects
    ])
    
    system_prompt = """You are an expert at matching social media posts to cryptocurrency/blockchain projects.

Your task is to analyze a tweet and determine which of the provided launchpad projects (if any) it's most likely referring to.

Return a JSON response with this structure:

{
    "selected_project_id": "project_id_or_null",
    "confidence": 0.85,
    "reasoning": "Detailed explanation of why this project was selected",
    "keywords_matched": ["keyword1", "keyword2"],
    "alternative_matches": [{"id": "alt_id", "confidence": 0.3, "reason": "why"}]
}

Matching criteria:
1. Direct project name mentions
2. Token symbol matches
3. Technology category alignment
4. Description keyword overlap
5. Context and sentiment alignment

Only select a project if you're reasonably confident (>0.6) that the tweet is specifically about that project.
If no good match exists, return null for selected_project_id."""

    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": f"""Tweet: {tweet_content}

Extracted keywords: {keywords}

Available projects:
{projects_text}

Which project (if any) is this tweet most likely about?"""}
    ]
    
    client = get_oai_async_client()
    
    try:
        response = await client.chat.completions.create(
            model=get_model_id(),
            messages=messages,
            temperature=0.1,
            max_tokens=1000
        )
        
        response_text = response.choices[0].message.content.strip()
        
        # Extract JSON from response
        json_match = _extract_json_from_text(response_text)
        if json_match:
            return json.loads(json_match)
        else:
            # Fallback: select highest scoring project
            return _fallback_project_selection(projects)
            
    except Exception as e:
        logger.error(f"Error in AI project matching: {e}")
        return _fallback_project_selection(projects)

def _create_project_identification(
    tweet_id: str,
    match_result: Dict[str, Any],
    keywords: List[str],
    all_projects: List[Dict[str, Any]]
) -> ProjectIdentification:
    """Create ProjectIdentification from matching results"""
    
    selected_project_id = match_result.get("selected_project_id")
    confidence = match_result.get("confidence", 0.0)
    reasoning = match_result.get("reasoning", "AI project matching")
    keywords_matched = match_result.get("keywords_matched", keywords)
    
    # Find project name if ID is provided
    project_name = None
    if selected_project_id:
        for project in all_projects:
            if project.get("id") == selected_project_id:
                project_name = project.get("name")
                break
    
    # Prepare alternative projects
    alternatives = []
    for alt in match_result.get("alternative_matches", []):
        if alt.get("id") != selected_project_id:  # Don't include selected project as alternative
            alternatives.append(alt)
    
    return ProjectIdentification(
        tweet_id=tweet_id,
        launchpad_id=selected_project_id,
        project_name=project_name,
        confidence=confidence,
        keywords_matched=keywords_matched,
        reasoning=reasoning,
        alternative_projects=alternatives
    )

def _extract_json_from_text(text: str) -> Optional[str]:
    """Extract JSON object from text response"""
    import re
    
    # Look for JSON object in the text
    json_match = re.search(r'\{.*\}', text, re.DOTALL)
    if json_match:
        return json_match.group()
    
    return None

def _fallback_keyword_extraction(tweet_content: str) -> Dict[str, Any]:
    """Simple keyword extraction as fallback"""
    import re
    
    # Extract token symbols
    token_symbols = re.findall(r'\$([A-Z]{2,6})', tweet_content)
    
    # Extract capitalized words (potential project names)
    capitalized_words = re.findall(r'\b[A-Z][a-z]+\b', tweet_content)
    
    # Technology keywords
    tech_keywords = ['DeFi', 'NFT', 'GameFi', 'Web3', 'AI', 'blockchain', 'crypto']
    found_tech = [k for k in tech_keywords if k.lower() in tweet_content.lower()]
    
    all_keywords = list(set(token_symbols + capitalized_words + found_tech))
    
    return {
        "keywords": all_keywords,
        "details": {
            "token_symbols": token_symbols,
            "project_names": capitalized_words,
            "technology_categories": found_tech
        },
        "confidence": 0.3
    }

def _fallback_project_selection(projects: List[Dict[str, Any]]) -> Dict[str, Any]:
    """Simple project selection as fallback"""
    
    if not projects:
        return {"selected_project_id": None, "confidence": 0.0, "reasoning": "No projects available"}
    
    # Select the project with highest relevance score
    best_project = max(projects, key=lambda p: p.get("relevance_score", 0))
    
    confidence = min(best_project.get("relevance_score", 0) / 10.0, 0.8)  # Scale to 0-0.8
    
    return {
        "selected_project_id": best_project.get("id"),
        "confidence": confidence,
        "reasoning": f"Selected project with highest relevance score: {best_project.get('relevance_score', 0)}",
        "keywords_matched": [],
        "alternative_matches": []
    } 