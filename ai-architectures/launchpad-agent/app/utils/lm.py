import openai
from app.config import settings

def get_oai_client() -> openai.OpenAI:
    return openai.OpenAI(
        base_url=settings.llm_base_url,
        api_key=settings.llm_api_key,
    )

def get_oai_async_client() -> openai.AsyncOpenAI:
    return openai.AsyncOpenAI(
        base_url=settings.llm_base_url,
        api_key=settings.llm_api_key,
    )
    
def get_model_id() -> str:
    return settings.llm_model_id