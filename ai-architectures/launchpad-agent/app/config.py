from pydantic_settings import BaseSettings
from pydantic import Field
from typing import Optional

class Settings(BaseSettings):
    # API Keys
    twitter_api_key: str = Field(alias="TWITTER_API_KEY", default="super-secret") 
    twitter_api_url: str = Field(alias="TWITTER_API_URL", default="https://imagine-backend.bvm.network/api/internal/twitter")
    launchpad_api_key: str = Field(alias="LAUNCHPAD_API_KEY", default="super-secret")
    launchpad_api_url: str = Field(alias="LAUNCHPAD_API_URL", default="https://api.launchpad.com/v1")
    llm_api_key: str = Field(alias="LLM_API_KEY", default="super-secret")
    llm_base_url: str = Field(alias="LLM_BASE_URL", default="https://api.openai.com/v1")
    llm_model_id: str = Field(alias="LLM_MODEL_ID", default="gpt-4o-mini")

    # Database
    mongo_uri: str = Field(alias="MONGO_URI", default="mongodb://localhost:27017/launchpad-agent")
    service_prefix: str = Field(alias="SERVICE_PREFIX", default="launchpad-agent")

    # Logging
    lite_logging_base_url: Optional[str] = Field(alias="LITE_LOGGING_BASE_URL", default=None)

    # app state
    app_env: str = Field(alias="APP_ENV", default="development")

    # Server
    host: str = Field(alias="HOST", default="0.0.0.0")
    port: int = Field(alias="PORT", default=80)

    class Config:
        env_file = ".env"
        case_sensitive = False

# Global settings instance
settings = Settings() 