import json
import re
from typing import Dict, List
from app.agents.mcp_agent import run_mcp_agent
from app.schemas.evaluation import TweetClassification, TweetEvaluation, SentimentScore
from app.utils.lm import get_oai_async_client, get_model_id
from app.utils.misc import float_clamp
import logging
from app.utils.twitter_api_calls import get_tweet_info
from lite_logging import async_log
from app.config import settings
import asyncio

logger = logging.getLogger(__name__)

INVESTMENT_KEYWORDS = [
    # Investment intent
    "invest", "investing", "investment", "buying", "purchase", "allocate", "portfolio",
    "position", "stake", "holding", "accumulating", "DCA", "dollar cost average",
    
    # Research intent  
    "research", "studying", "analyzing", "due diligence", "DYOR", "fundamentals",
    "whitepaper", "tokenomics", "roadmap", "team", "partnerships",
    
    # Positive engagement
    "bullish", "excited", "promising", "potential", "opportunity", "undervalued",
    "gem", "alpha", "early", "innovative", "revolutionary",
    
    # Learning intent
    "learn", "understand", "explain", "curious", "interested", "exploring",
    "deep dive", "breakdown", "analysis", "insights"
]

SPAM_INDICATORS = [
    "pump", "moon", "lambo", "🚀", "💎", "🙌", "HODL", "diamond hands",
    "to the moon", "100x", "1000x", "guaranteed", "risk free", "easy money",
    "get rich", "financial advice", "not financial advice", "NFA"
]

async def classify_tweet(threaded_content: str, target_tweet_id: str = None) -> TweetEvaluation:
    """
    Stage 1: Classify tweet as candidate, spam, irrelevant, or negative
    
    Args:
        tweet_content: The text content of the tweet
        tweet_id: Optional tweet ID for tracking
        
    Returns:
        TweetEvaluation with classification, sentiment, and reasoning
    """
    
    try:
        # Pre-process tweet content
        cleaned_content = _preprocess_tweet(threaded_content)
        logger.info(f"Cleaned content: {cleaned_content}")
        
        # Quick keyword analysis
        investment_keywords_found = _find_investment_keywords(cleaned_content)
        logger.info(f"Investment keywords found: {investment_keywords_found}")
        
        spam_score = _calculate_spam_score(cleaned_content)
        logger.info(f"Spam score: {spam_score}")

        # Use AI for detailed classification
        classification_result = await _ai_classify_tweet(cleaned_content, target_tweet_id)
        logger.info(f"Classification result: {classification_result}")

        # Calculate final classification
        final_classification = _determine_final_classification(
            classification_result, 
            investment_keywords_found, 
            spam_score
        )
        logger.info(f"Final classification: {final_classification}")
        
        # Extract sentiment scores
        sentiment = _extract_sentiment(classification_result)
        logger.info(f"Sentiment: {sentiment}")
        
        # Calculate investment intent score
        investment_intent_score = _calculate_investment_intent_score(
            classification_result, 
            investment_keywords_found
        )
        
        return TweetEvaluation(
            tweet_id=target_tweet_id,
            classification=final_classification,
            sentiment=sentiment,
            confidence=classification_result.get("confidence", 0.5),
            reasoning=classification_result.get("reasoning", "AI classification"),
            keywords_found=investment_keywords_found,
            investment_intent_score=investment_intent_score
        )
        
    except Exception as e:
        import traceback
        traceback_str = traceback.format_exc()

        asyncio.create_task(async_log(
            traceback_str, 
            channel=settings.lite_logging_channel,
            tags=["tweet_classifier", "error"],
            server_url=settings.lite_logging_base_url
        ))

        # Return safe default
        return TweetEvaluation(
            tweet_id=target_tweet_id or "unknown",
            classification=TweetClassification.IRRELEVANT,
            sentiment=SentimentScore(positive=0.0, negative=0.0, neutral=1.0, confidence=0.0),
            confidence=0.0,
            reasoning=f"Error during classification: {str(e)}",
            keywords_found=[],
            investment_intent_score=0.0
        )

async def _ai_classify_tweet(tweet_content: str, target_tweet_id: str = None) -> Dict:
    """Use AI to classify the tweet with detailed analysis"""
    
    system_prompt = """You are an expert tweet classifier for cryptocurrency/blockchain/ai products/ai agents investment analysis.

Your task is to classify a tweet (use context from the provided tweet thread). Analyze the investment relevance of the tweet. Return a JSON response with this exact structure:

{
    "reasoning": "Detailed explanation of classification",
    "classification": "candidate|spam|irrelevant|negative",
    "confidence": 0.85,
    "sentiment": {
        "positive": 0.7,
        "negative": 0.1,
        "neutral": 0.2,
        "confidence": 0.8
    },
    "investment_intent": 0.8,
    "keywords": ["keyword1", "keyword2"]
}

Classification criteria:

CANDIDATE tweets show:
- Genuine interest in learning about or investing in crypto/blockchain/ai products/ai agents projects
- Positive sentiment about specific technologies or projects
- Questions about project fundamentals, tokenomics, team, or roadmap
- Mentions of research, due diligence, or analysis
- Professional investment discussion

SPAM tweets contain:
- Pump and dump language ("moon", "100x", "guaranteed gains")
- Excessive emojis and hype language
- Promotional content without substance
- Bot-like repetitive patterns
- Scam indicators

IRRELEVANT tweets are:
- Not related to cryptocurrency, blockchain, or investing
- General tech discussion without investment context
- Personal life updates
- Non-crypto financial discussions

NEGATIVE tweets express:
- Complaints about projects, exchanges, or crypto in general
- FUD (Fear, Uncertainty, Doubt) spreading
- Criticism without constructive feedback
- Negative sentiment about the crypto space

Focus on genuine investment interest and learning intent, not just price speculation."""

    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": f"Classify this tweet thread:\n\n{tweet_content}\n\nPay more attention to the last one (Tweet: {target_tweet_id})."}
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
        json_match = re.search(r'\{.*\}', response_text, re.DOTALL)
        if json_match:
            return json.loads(json_match.group())
        else:
            # Fallback parsing
            return _parse_non_json_response(response_text)
            
    except Exception as e:
        import traceback
        traceback_str = traceback.format_exc()

        asyncio.create_task(async_log(
            traceback_str, 
            channel=settings.lite_logging_channel,
            tags=["tweet_classifier", "error"],
            server_url=settings.lite_logging_base_url
        ))

        return {
            "classification": "irrelevant",
            "confidence": 0.0,
            "reasoning": f"AI classification failed: {str(e)}",
            "sentiment": {"positive": 0.0, "negative": 0.0, "neutral": 1.0, "confidence": 0.0},
            "investment_intent": 0.0,
            "keywords": []
        }

def _preprocess_tweet(content: str) -> str:
    """Clean and preprocess tweet content"""
    # Remove URLs
    content = re.sub(r'http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\\(\\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+', '', content)

    # Remove extra whitespace
    content = re.sub(r'\s+', ' ', content).strip()

    return content

def _find_investment_keywords(content: str) -> List[str]:
    """Find investment-related keywords in the content"""
    content_lower = content.lower()
    found_keywords = []
    
    for keyword in INVESTMENT_KEYWORDS:
        if keyword.lower() in content_lower:
            found_keywords.append(keyword)
    
    return found_keywords

def _calculate_spam_score(content: str) -> float:
    """Calculate spam likelihood score (0-1)"""
    content_lower = content.lower()
    spam_indicators_found = 0
    WORDS_COUNT = len(content_lower.split())

    for indicator in SPAM_INDICATORS:
        if indicator.lower() in content_lower:
            spam_indicators_found += 1

    # Count excessive emojis
    emoji_count = len(re.findall(r'[🚀💎🙌📈📊💰🔥⚡️🌙]', content))
    if emoji_count > 3:
        spam_indicators_found += emoji_count

    # Count excessive capitalization
    caps_ratio = sum(1 for c in content if c.isupper()) / max(len(content), 1)
    if caps_ratio > 0.5:
        spam_indicators_found += 5

    return min(spam_indicators_found / WORDS_COUNT, 1.0)

def _determine_final_classification(
    ai_result: Dict, 
    investment_keywords: List[str], 
    spam_score: float
) -> TweetClassification:
    """Combine AI classification with rule-based checks"""
    
    ai_classification = str(ai_result.get("classification", "irrelevant")).lower()
    
    # Override with spam if high spam score
    if spam_score > 0.6:
        return TweetClassification.SPAM

    # Boost candidate classification if investment keywords present
    if len(investment_keywords) >= 2 and ai_classification in ["candidate", "irrelevant"]:
        return TweetClassification.CANDIDATE

    # Map AI classification to enum
    classification_map = {
        "candidate": TweetClassification.CANDIDATE,
        "spam": TweetClassification.SPAM,
        "irrelevant": TweetClassification.IRRELEVANT,
        "negative": TweetClassification.NEGATIVE
    }

    return classification_map.get(ai_classification, TweetClassification.IRRELEVANT)

def _extract_sentiment(ai_result: Dict) -> SentimentScore:
    """Extract sentiment scores from AI result"""
    sentiment_data: dict[str, float] = ai_result.get("sentiment", {})

    positive_score = float_clamp(sentiment_data.get("positive", 0.0), 0.0, 1.0)
    negative_score = float_clamp(sentiment_data.get("negative", 0.0), 0.0, 1.0)
    neutral_score = float_clamp(sentiment_data.get("neutral", 1.0), 0.0, 1.0)
    confidence_score = float_clamp(sentiment_data.get("confidence", 0.0), 0.0, 1.0)

    total_score = positive_score + negative_score + neutral_score

    if total_score > 0:
        positive_score = positive_score / total_score
        negative_score = negative_score / total_score
        neutral_score = neutral_score / total_score
    else:
        positive_score = 0.0
        negative_score = 0.0
        neutral_score = 1.0
        confidence_score = 0.3

    return SentimentScore(
        positive=positive_score,
        negative=negative_score,
        neutral=neutral_score,
        confidence=confidence_score
    )

def _calculate_investment_intent_score(ai_result: Dict, keywords: List[str]) -> float:
    """Calculate overall investment intent score"""

    ai_intent = ai_result.get("investment_intent", 0.0)
    ai_intent = float_clamp(ai_intent, 0.0, 1.0)
    keyword_boost = float_clamp(len(keywords) * 0.1, 0.0, 0.3)

    return float_clamp(ai_intent + keyword_boost, 0.0, 1.0)

def _parse_non_json_response(response_text: str) -> Dict:
    """Fallback parser for non-JSON responses"""
    classification = TweetClassification.IRRELEVANT.value

    # Simple keyword matching for fallback
    response_lower = response_text.lower()

    if any(word in response_lower for word in ["candidate", "investment", "positive"]):
        classification = TweetClassification.CANDIDATE.value
    elif any(word in response_lower for word in ["spam", "promotional", "pump"]):
        classification = TweetClassification.SPAM.value
    elif any(word in response_lower for word in ["negative", "fud", "complaint"]):
        classification = TweetClassification.NEGATIVE.value

    return {
        "classification": classification,
        "confidence": 0.3,
        "reasoning": "Fallback classification due to parsing error",
        "sentiment": {"positive": 0.0, "negative": 0.0, "neutral": 1.0, "confidence": 0.0},
        "investment_intent": 0.0,
        "keywords": []
    } 