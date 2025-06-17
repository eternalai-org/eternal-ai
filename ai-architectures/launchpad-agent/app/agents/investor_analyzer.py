import json
from typing import Dict, List, Any, Optional
from app.schemas.evaluation import (
    InvestorProfile, InvestorGrade, ResearchInterest, 
    InvestmentBehavior, SocialMetrics
)
from app.utils.lm import get_oai_async_client, get_model_id
from app.utils.misc import float_clamp, retry
import logging
from json_repair import repair_json
from app.utils.twitter_api_calls import get_tweet_threads_by_twitter_id
from app.schemas.twitter import Tweet
from lite_logging import async_log
import asyncio
from app.config import settings

logger = logging.getLogger(__name__)

async def analyze_investor_profile(
    user_id: str, 
    tweet_content: str, 
    launchpad_id: str,
    network_id: str = "8453"
) -> InvestorProfile:
    """
    Stage 3: AI-Enhanced comprehensive analysis of the candidate investor
    """
    try:
        logger.info(f"Starting AI-enhanced investor analysis for user {user_id}")
        
        # Gather social data
        social_data = await gather_social_data(user_id)
        
        with open("social_data.json", "w") as f:
            import json
            json.dump(social_data, f, indent=4, default=lambda x: x.model_dump(mode="json") if isinstance(x, Tweet) else str(x))

        if social_data.get("profile") is None:
            return _create_error_profile(user_id, "Failed to get profile")

        # Get project details for context-aware analysis
        from app.utils.launchpad_api_calls import get_launchpad_detail
        req = await get_launchpad_detail(launchpad_id, network_id)

        if req.result is None:
            return _create_error_profile(user_id, "Failed to get project details")

        project_details = req.result.model_dump()
        
        # Use AI for enhanced analysis
        ai_analysis = await analyze_investor_with_ai(
            social_data, 
            project_details, 
            tweet_content,
            user_id
        )
        
        if ai_analysis:
            # AI analysis succeeded
            return ai_analysis

        # Fallback to basic analysis
        logger.warning(f"AI analysis failed for {user_id}, falling back to basic analysis")
        return await analyze_investor_basic(user_id, tweet_content, launchpad_id, social_data)
        
    except Exception as e:
        import traceback
        traceback_str = traceback.format_exc()

        asyncio.create_task(async_log(
            traceback_str, 
            channel=settings.lite_logging_channel,
            tags=["investor_analyzer", "error"],
            server_url=settings.lite_logging_base_url
        ))

        logger.error(f"Error analyzing investor {user_id}: {e}")
        return _create_error_profile(user_id, str(e))
    
from app.schemas.twitter import Tweet
async def summarize_thread(tweets: list[Tweet]) -> str:
    fmt_content = ""

    for tweet in tweets:
        fmt_content += f"{tweet.text}\n"

    system_prompt = """
You are a helpful assistant that summarizes a thread of tweets.
"""

    user_prompt = f"""
Here is the thread of tweets:
{fmt_content}

Your task now is to summarize the thread of tweets into a single tweet. The summarization should contain the following information:
- The main idea, topic of the thread
- The key points of the thread
- The main conclusion of the thread
- The main recommendation of the thread, if any
- The main action items of the thread, if any
- The main risks of the thread, if any

Write the summary in short and concise manner.
"""

    client = get_oai_async_client()
    model = get_model_id()

    response = await client.chat.completions.create(
        model=model,
        messages=[
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": user_prompt}
        ],
        temperature=0.3,
        max_tokens=1024
    )

    return response.choices[0].message.content

async def gather_social_data(user_id: str) -> Dict[str, Any]:
    """Collect social media data"""
    try:
        from app.mcps.twitter_mcp import _get_twitter_user_info_by_id

        # Get profile
        profile_result = await _get_twitter_user_info_by_id(user_id)

        if profile_result is None:
            return {"profile": None, "tweets": []}

        req = await get_tweet_threads_by_twitter_id(user_id)
        threads = req.result
        tweets: list[Tweet] = []
        summary_tweet_content: Dict[str, str] = {}

        for thread_parent, thread_tweets in threads.items():
            tweets.extend(thread_tweets)

            if len(thread_tweets) > 1:
                summary_tweet_content[thread_parent] = await summarize_thread(thread_tweets)

            elif len(thread_tweets) == 1:
                summary_tweet_content[thread_parent] = thread_tweets[0].text

        return {
            "profile": profile_result,
            "tweets": tweets,
            "threads": threads,
            "summary_tweet_content": summary_tweet_content
        }
        
    except Exception as e:
        import traceback
        traceback_str = traceback.format_exc()

        asyncio.create_task(async_log(
            traceback_str, 
            channel=settings.lite_logging_channel,
            tags=["gather_social_data", "error"],
            server_url=settings.lite_logging_base_url
        ))

        logger.error(f"Error gathering social data: {e}")
        return {"profile": {}, "tweets": []}

async def analyze_research_interests(tweets: List[Tweet]) -> List[ResearchInterest]:
    """Basic research interest analysis"""
    if not tweets:
        return []
    
    # Simple keyword-based analysis for now
    categories = {
        "DeFi": ["defi", "dex", "yield", "farming", "liquidity", "swap"],
        "NFT": ["nft", "opensea", "mint", "collection", "art"],
        "GameFi": ["gamefi", "play", "earn", "gaming", "metaverse"],
        "AI": ["ai", "artificial", "intelligence", "machine", "learning"],
        "Data": ["DS", "DE", "Data", "Data Science", "Data Engineering", "Data Analysis", "Data Visualization"]
    }

    interests = []

    for category, keywords in categories.items():
        matches = 0

        for tweet in tweets:
            text = tweet.text.lower()

            if any(keyword in text for keyword in keywords):
                matches += 1

        if matches > 0:
            confidence = float_clamp(matches / len(tweets), 0, 1)
            interests.append(
                ResearchInterest(
                    category=category,
                    confidence=confidence,
                    evidence_tweets=[str(i) for i in range(min(matches, 3))],
                    technical_depth=float_clamp(confidence * 0.7, 0, 1),  # Estimate technical depth
                    keywords=keywords
                )
            )
    
    return interests

def analyze_investment_behavior_basic(tweets: List[Dict]) -> InvestmentBehavior:
    """Basic investment behavior analysis"""
    return InvestmentBehavior(
        risk_tolerance="Moderate",
        investment_size_preference="Medium",
        time_horizon="Medium-term",
        due_diligence_score=0.5,
        portfolio_diversity=0.5
    )

def calculate_social_metrics(profile: Dict[str, Any], social_posted_content: List[str]) -> SocialMetrics:
    """Calculate basic social metrics"""
    metrics = profile.get("metrics", {})
    
    return SocialMetrics(
        followers_count=metrics.get("followers_count", 0),
        following_count=metrics.get("following_count", 0),
        tweet_count=metrics.get("tweet_count", 0),
        account_age_days=365,  # Default
        engagement_rate=5.0,   # Default
        posting_frequency=1.0, # Default
        crypto_focus_ratio=_calculate_content_focus_ratio(social_posted_content)
    )

def _calculate_content_focus_ratio(social_posted_content: List[str]) -> float:
    """Calculate crypto focus ratio"""
    if not social_posted_content:
        return 0.0
    
    crypto_keywords = [
        'crypto', 'bitcoin', 'ethereum', 
        'blockchain', 'defi', 'nft',
        'web3', 'ai', 'artificial', 
        'intelligence', 'machine', 
        'learning', 'data', 
        'science', 'engineering', 
        'analysis', 'visualization'
    ]

    crypto_tweets = 0
    
    for content in social_posted_content:
        text = content.lower()
        if any(keyword in text for keyword in crypto_keywords):
            crypto_tweets += 1

    return crypto_tweets / len(social_posted_content)

def calculate_basic_score(research_interests: List[ResearchInterest], social_metrics: SocialMetrics) -> float:
    """Calculate basic investor score"""
    score = 40.0  # Base score

    # Research interests bonus
    if research_interests:
        avg_confidence = sum(interest.confidence for interest in research_interests) / len(research_interests)
        score += avg_confidence * 20

    # Social metrics bonus
    if social_metrics.followers_count > 10000:
        score += 10
    
    elif social_metrics.followers_count > 1000:
        score += 5
    
    elif social_metrics.followers_count > 100:
        score += 2
    
    elif social_metrics.followers_count > 50:
        score += 1

    score += social_metrics.crypto_focus_ratio * 15
    return float_clamp(score, 0.0, 100.0)

def score_to_grade(score: float) -> InvestorGrade:
    """Convert score to grade"""
    if score >= 90:
        return InvestorGrade.A
    elif score >= 80:
        return InvestorGrade.B
    elif score >= 70:
        return InvestorGrade.C
    elif score >= 60:
        return InvestorGrade.D
    else:
        return InvestorGrade.E

def _create_error_profile(user_id: str, error_msg: str) -> InvestorProfile:
    """Create error profile"""

    return InvestorProfile(
        user_id=user_id,
        username="error",
        grade=InvestorGrade.E,
        score=0.0,
        research_interests=[],
        investment_behavior=InvestmentBehavior(
            risk_tolerance="Unknown",
            investment_size_preference="Unknown",
            time_horizon="Unknown",
            due_diligence_score=0.0,
            portfolio_diversity=0.0
        ),
        social_metrics=SocialMetrics(
            followers_count=0,
            following_count=0,
            tweet_count=0,
            account_age_days=0,
            engagement_rate=0.0,
            posting_frequency=0.0,
            crypto_focus_ratio=0.0
        ),
        risk_factors=[f"Analysis error: {error_msg}"],
        strengths=[],
        reasoning=f"Failed to analyze investor: {error_msg}"
    )

async def analyze_investor_with_ai(
    social_data: Dict[str, Any],
    project_details: Dict[str, Any], 
    tweet_content: str,
    user_id: str
) -> Optional[InvestorProfile]:
    """AI-powered investor analysis using LLM"""
    
    try:
        client = get_oai_async_client()
        model = get_model_id()
        
        # Prepare data for AI analysis
        
        summary_threads: dict[str, str] = social_data.get("summary_tweet_content", {})

        summary_tweet_content = ""
        social_posted_content = []
        
        for k, v in summary_threads.items():
            summary_tweet_content += f"Thread ID: {k}\nSummary: {v}\n\n"
            social_posted_content.append(v)

        profile: dict[str, Any] = social_data.get("profile", {})
        
        # Get project-specific insights
        project_insights = await analyze_project_specific_fit(project_details, social_data)
        
        # Build prompt for AI analysis
        analysis_prompt = f"""
You are an expert investment analyst evaluating whether a Twitter user would be a good fit as an investor for a specific blockchain/crypto project.

## PROJECT DETAILS:
- ID: {project_details.get('id', 'Unknown')}
- Name: {project_details.get('name', 'Unknown')}
- Description: {project_details.get('description', 'No description available')}
- Ticker: {project_details.get('token_symbol', 'Unknown')}
- Market Cap: {project_details.get('market_cap_usd', 'Unknown')}

## INVESTOR PROFILE:
- Username: {profile.get('username', 'Unknown')}
- Name: {profile.get('name', 'Unknown')}
- Followers: {profile.get('metrics', {}).get('followers_count', 0)}
- Following: {profile.get('metrics', {}).get('following_count', 0)}
- Tweet Count: {profile.get('metrics', {}).get('tweet_count', 0)}

## RECENT POSTED:
{summary_tweet_content}

## CANDIDATE TWEET:
"{tweet_content}"

## PROJECT-SPECIFIC INSIGHTS:
{project_insights}

Based on this information, analyze the investor and provide a comprehensive evaluation.

Please respond with a JSON object containing:
{{
    "reasoning": "<detailed explanation of the analysis and scoring>",
    "overall_score": <float 0-100>,
    "grade": "<A|B|C|D|E>",
    "project_fit_score": <float 0.0-1.0>,
    "research_interests": [
        {{
            "category": "<category>",
            "confidence": <float 0.0-1.0>,
            "technical_depth": <float 0.0-1.0>,
            "evidence": "<brief explanation>"
        }}
    ],
    "investment_behavior": {{
        "risk_tolerance": "<Conservative|Moderate|Aggressive>",
        "investment_size_preference": "<Small|Medium|Large>",
        "time_horizon": "<Short-term|Medium-term|Long-term>",
        "due_diligence_score": <float 0.0-1.0>,
        "portfolio_diversity": <float 0.0-1.0>
    }},
    "strengths": [<list of strength strings>],
    "risk_factors": [<list of risk factor strings>],
}}

Focus on:
1. Alignment between investor interests and project type
2. Quality of understanding demonstrated in tweets
3. Investment sophistication and experience
4. Social proof and credibility
5. Potential red flags or concerning patterns
6. How well this investor fits THIS SPECIFIC PROJECT
7. Project-specific insights and compatibility

Be thorough but concise in your analysis.
"""

        async def wraps():
            response = await client.chat.completions.create(
                model=model,
                messages=[
                    {
                        "role": "system", 
                        "content": "You are an expert blockchain investment analyst. Provide detailed, objective analysis."
                    },
                    {"role": "user", "content": analysis_prompt}
                ],
                temperature=0.3,
                max_tokens=2048
            )

            ai_result = response.choices[0].message.content

            try:
                repaired_json = repair_json(ai_result)
                return json.loads(repaired_json)
 
            except json.JSONDecodeError:
                import re
                json_match = re.search(r'```(?:json)?\s*(\{.*?\})\s*```', ai_result, re.DOTALL)

                if json_match:
                    repaired_json = repair_json(json_match.group(1))
                    return json.loads(repaired_json)

                raise Exception("Failed to parse AI response as JSON")
        
        ai_data: dict[str, Any] = await retry(wraps, max_retry=3, first_interval=10, interval_multiply=1)()

        # Build InvestorProfile from AI analysis
        research_interests = []
        for interest in ai_data.get("research_interests", []):

            if not isinstance(interest, dict):
                logger.warning(f"Invalid research interest: {interest}")
                continue

            research_interests.append(ResearchInterest(
                category=interest.get("category", "Unknown"),
                confidence=float_clamp(interest.get("confidence", 0.5), 0, 1),
                evidence_tweets=[],  # We could extract specific tweet IDs here
                technical_depth=float_clamp(interest.get("technical_depth", 0.5), 0, 1),
                keywords=[]
            ))

        investment_behavior = InvestmentBehavior(
            risk_tolerance=ai_data.get("investment_behavior", {}).get("risk_tolerance", "Moderate"),
            investment_size_preference=ai_data.get("investment_behavior", {}).get("investment_size_preference", "Medium"),
            time_horizon=ai_data.get("investment_behavior", {}).get("time_horizon", "Medium-term"),
            due_diligence_score=float_clamp(ai_data.get("investment_behavior", {}).get("due_diligence_score", 0.5), 0, 1),
            portfolio_diversity=float_clamp(ai_data.get("investment_behavior", {}).get("portfolio_diversity", 0.5), 0, 1)
        )
        
        # Calculate social metrics
        social_metrics = calculate_social_metrics(profile, list(summary_threads.values()))

        # Map grade string to enum
        grade_str = ai_data.get("grade", "D").upper()

        try:
            grade = InvestorGrade(grade_str)
        except ValueError:
            grade = InvestorGrade.D
        
        return InvestorProfile(
            user_id=user_id,
            username=profile.get("username", "unknown"),
            name=profile.get("name"),
            grade=grade,
            score=float_clamp(ai_data.get("overall_score", 50.0), 0, 100),
            research_interests=research_interests,
            investment_behavior=investment_behavior,
            social_metrics=social_metrics,
            risk_factors=ai_data.get("risk_factors", []),
            strengths=ai_data.get("strengths", []),
            reasoning=ai_data.get("reasoning", "AI-powered analysis completed"),
            project_fit_score=float_clamp(ai_data.get("project_fit_score", 0.5), 0, 1)
        )
        
    except Exception as e:
        import traceback
        traceback_str = traceback.format_exc()

        asyncio.create_task(async_log(
            traceback_str, 
            channel=settings.lite_logging_channel,
            tags=["analyze_investor_with_ai", "error"],
            server_url=settings.lite_logging_base_url
        ))

        return None

async def analyze_project_specific_fit(project_details: Dict[str, Any], social_data: Dict[str, Any]) -> str:
    """Analyze project-specific compatibility factors"""
    
    project_name = project_details.get('name', '').lower()
    project_desc = project_details.get('description', '').lower()
    
    insights = []
    
    # Identify project type and relevant factors
    if any(keyword in f"{project_name} {project_desc}" for keyword in ['defi', 'dex', 'swap', 'yield', 'liquidity']):
        insights.append("PROJECT TYPE: DeFi - Look for understanding of DeFi mechanics, yield farming, AMMs")
    elif any(keyword in f"{project_name} {project_desc}" for keyword in ['nft', 'art', 'collectible', 'mint']):
        insights.append("PROJECT TYPE: NFT - Look for art appreciation, collecting behavior, creativity")
    elif any(keyword in f"{project_name} {project_desc}" for keyword in ['game', 'gaming', 'play', 'metaverse']):
        insights.append("PROJECT TYPE: GameFi - Look for gaming interest, virtual world engagement")
    elif any(keyword in f"{project_name} {project_desc}" for keyword in ['ai', 'artificial', 'intelligence', 'ml']):
        insights.append("PROJECT TYPE: AI - Look for tech sophistication, AI/ML understanding")
    else:
        insights.append("PROJECT TYPE: General - Evaluate overall crypto knowledge and investment experience")
    
    # Analyze tweet patterns for project relevance
    tweets = social_data.get("tweets", [])
    if tweets:
        relevant_tweets = 0
        total_analyzed = min(len(tweets), 20)
        
        for tweet in tweets[:total_analyzed]:
            text = tweet.text.lower()
            if any(keyword in text for keyword in project_name.split() if len(keyword) > 2):
                relevant_tweets += 1
        
        if relevant_tweets > 0:
            insights.append(f"RELEVANCE: Found {relevant_tweets}/{total_analyzed} tweets mentioning project-related terms")
        else:
            insights.append("RELEVANCE: No direct project mentions found in recent tweets")
    
    return "\n".join(insights)

async def analyze_investor_basic(
    user_id: str, 
    tweet_content: str, 
    launchpad_id: str,
    social_data: Dict[str, Any]
) -> InvestorProfile:
    """Fallback basic analysis when AI analysis fails"""
    
    # Analyze components using existing basic methods
    research_interests = await analyze_research_interests(social_data.get("tweets", []))
    investment_behavior = analyze_investment_behavior_basic(social_data.get("tweets", []))
    social_metrics = calculate_social_metrics(
        social_data.get("profile", {}), 
        list(social_data.get("summary_tweet_content", {}).values())
    )
    
    # Calculate score and grade
    final_score = calculate_basic_score(research_interests, social_metrics)
    grade = score_to_grade(final_score)
    
    return InvestorProfile(
        user_id=user_id,
        username=social_data.get("profile", {}).get("username", "unknown"),
        name=social_data.get("profile", {}).get("name"),
        grade=grade,
        score=final_score,
        research_interests=research_interests,
        investment_behavior=investment_behavior,
        social_metrics=social_metrics,
        risk_factors=[],
        strengths=[],
        reasoning=f"Basic analysis: Investor scored {final_score:.1f} points and received grade {grade.value}",
        project_fit_score=0.5  # Default neutral fit
    ) 