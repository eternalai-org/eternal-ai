from pymongo import MongoClient
from functools import lru_cache
from app.config import settings

MONGO_URI = settings.mongo_uri

POOL_CONFIGURATIONS = {
    "maxPoolSize": 32,
    "minPoolSize": 4,
    "connectTimeoutMS": 2000,
    "socketTimeoutMS": 2000,
    "serverSelectionTimeoutMS": 5000
}

SERVICE_PREFIX = settings.service_prefix

@lru_cache(maxsize=1)
def get_mongodb_client():
    global POOL_CONFIGURATIONS, MONGO_URI

    return MongoClient(
        MONGO_URI, 
        **POOL_CONFIGURATIONS
    )
    
def get_mongo_database(name: str):
    global SERVICE_PREFIX
    db_fullname = f"{SERVICE_PREFIX}-{name}"
    client = get_mongodb_client()
    return client.get_database(db_fullname)