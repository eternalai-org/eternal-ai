import openai
import os 

base_url = os.getenv("LLM_BASE_URL", "http://localmodel:65534/v1")
api_key = os.getenv("LLM_API_KEY", "empty")
model_id = os.getenv("LLM_MODEL_ID", "empty")

def get_oai_client() -> openai.OpenAI:
    return openai.OpenAI(
        base_url=base_url,
        api_key=api_key,
    )

def get_oai_async_client() -> openai.AsyncOpenAI:
    return openai.AsyncOpenAI(
        base_url=base_url,
        api_key=api_key,
    )
    
def get_model_id() -> str:
    return model_id