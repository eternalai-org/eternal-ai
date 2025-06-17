from enum import Enum
from pydantic import BaseModel, Field
from typing import Optional, List, Union
from datetime import datetime, timezone

class TweetClassification(str, Enum):
    CANDIDATE = "candidate"
    SPAM = "spam"
    IRRELEVANT = "irrelevant"
    NEGATIVE = "negative"

class SentimentScore(BaseModel):
    positive: float = Field(ge=0.0, le=1.0, description="Positive sentiment score")
    negative: float = Field(ge=0.0, le=1.0, description="Negative sentiment score")
    neutral: float = Field(ge=0.0, le=1.0, description="Neutral sentiment score")
    confidence: float = Field(ge=0.0, le=1.0, description="Confidence in sentiment analysis")

class TweetEvaluation(BaseModel):
    tweet_id: str
    classification: TweetClassification
    sentiment: SentimentScore
    confidence: float = Field(ge=0.0, le=1.0, description="Overall classification confidence")
    reasoning: str = Field(description="Explanation for the classification")
    keywords_found: List[str] = Field(default_factory=list, description="Investment-related keywords found")
    investment_intent_score: float = Field(ge=0.0, le=1.0, description="Score indicating investment intent")

class InvestorGrade(str, Enum):
    A = "A"  # Excellent investor profile (90-100)
    B = "B"  # Good investor profile (80-89)
    C = "C"  # Average investor profile (70-79)
    D = "D"  # Below average investor profile (60-69)
    E = "E"  # Poor investor profile (0-59)

class ResearchInterest(BaseModel):
    category: str = Field(description="Research category (DeFi, NFT, GameFi, etc.)")
    confidence: float = Field(ge=0.0, le=1.0, description="Confidence in this interest")
    evidence_tweets: List[str] = Field(description="Tweet IDs that support this interest")
    technical_depth: float = Field(ge=0.0, le=1.0, description="Level of technical understanding")
    keywords: List[str] = Field(default_factory=list, description="Keywords associated with this interest")

class InvestmentBehavior(BaseModel):
    risk_tolerance: str = Field(description="Conservative, Moderate, or Aggressive")
    investment_size_preference: str = Field(description="Small, Medium, or Large")
    time_horizon: str = Field(description="Short-term, Medium-term, or Long-term")
    due_diligence_score: float = Field(ge=0.0, le=1.0, description="Quality of research before investing")
    portfolio_diversity: float = Field(ge=0.0, le=1.0, description="How diversified their interests are")

class SocialMetrics(BaseModel):
    followers_count: int
    following_count: int
    tweet_count: int
    account_age_days: int
    engagement_rate: float = Field(ge=0.0, description="Average engagement per tweet")
    posting_frequency: float = Field(ge=0.0, description="Tweets per day")
    crypto_focus_ratio: float = Field(ge=0.0, le=1.0, description="Ratio of crypto-related tweets")

class InvestorProfile(BaseModel):
    user_id: str
    username: str
    name: Optional[str] = None
    grade: InvestorGrade
    score: float = Field(ge=0.0, le=100.0, description="Overall investor score")
    research_interests: List[ResearchInterest]
    investment_behavior: InvestmentBehavior
    social_metrics: SocialMetrics
    risk_factors: List[str] = Field(description="Potential concerns or red flags")
    strengths: List[str] = Field(description="Positive attributes")
    reasoning: str = Field(description="Detailed explanation of the grading")
    project_fit_score: Optional[float] = Field(default=None, ge=0.0, le=1.0, description="How well they fit the specific project")
    created_at: datetime = Field(default_factory=lambda: datetime.now(timezone.utc))

class ProjectIdentification(BaseModel):
    tweet_id: Union[int, str]
    launchpad_id: Optional[Union[int, str]] = "N/A"
    project_name: Optional[str] = "N/A"
    description: Optional[str] = "N/A"
    confidence: float = Field(ge=0.0, le=1.0, description="Confidence in project identification")
    reasoning: str = Field(description="Explanation for project identification")

class EvaluationResult(BaseModel):
    """Complete evaluation result for a tweet"""
    tweet_id: str
    twitter_id: str
    tweet_content: str
    
    # Stage 1: Classification
    tweet_evaluation: TweetEvaluation
    
    # Stage 2: Project Identification  
    project_identification: Optional[ProjectIdentification] = None
    
    # Stage 3: Investor Analysis (only if stages 1&2 succeed)
    investor_profile: Optional[InvestorProfile] = None
    
    # Metadata
    processing_time_seconds: float
    created_at: datetime = Field(default_factory=lambda: datetime.now(timezone.utc))
    status: str = Field(default="completed", description="Processing status") 