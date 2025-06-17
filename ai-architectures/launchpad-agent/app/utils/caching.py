from functools import wraps
from typing import Any, Callable, TypeVar, Optional
from datetime import datetime, timedelta, timezone
import inspect
from pydantic import BaseModel
from app.utils.mongodb import get_mongo_database
import logging
import pickle

logger = logging.getLogger(__name__) 

T = TypeVar('T')
F = TypeVar('F', bound=Callable[..., Any])

def mongo_cache(
    collection_name: str,
    ttl_seconds: int = 3600 * 24 * 30, # 30 days
    key_prefix: str = "",
    key_builder: Optional[Callable[..., str]] = None,
    object_builder: Optional[Callable[..., Any]] = None
):
    """
    Generic MongoDB caching decorator that works with any function type and return type.
    
    Args:
        collection_name (str): Name of the MongoDB collection to use for caching
        ttl_seconds (int): Time to live for cached items in seconds
        key_prefix (str): Prefix to add to cache keys
        key_builder (Callable): Optional custom function to build cache keys from function args
    """

    def decorator(func: F) -> F:
        ordered_args = inspect.getfullargspec(func).args

        @wraps(func)
        async def async_wrapper(*args: Any, **kwargs: Any) -> Any:
            call_args_dict = {}

            for i, e in enumerate(args):
                call_args_dict[ordered_args[i]] = e

            for k, v in kwargs.items():
                call_args_dict[k] = v

            return await _cache_wrapper_async(func, **call_args_dict)
            
        @wraps(func)
        def sync_wrapper(*args: Any, **kwargs: Any) -> Any:
            call_args_dict = {}

            for i, e in enumerate(args):
                call_args_dict[ordered_args[i]] = e

            for k, v in kwargs.items():
                call_args_dict[k] = v

            return _cache_wrapper_sync(func, **call_args_dict)
            
        def _cache_wrapper_sync(func: F, *args: Any, **kwargs: Any) -> Any:
            db = get_mongo_database("cache")
            collection = db.get_collection(collection_name)
            
            # Build cache key
            if key_builder:
                cache_key = f"{key_prefix}:{key_builder(*args, **kwargs)}"
            else:
                # Default key building from function name and arguments
                arg_str = ":".join(str(arg) for arg in args)
                kwarg_str = ":".join(f"{k}={v}" for k, v in sorted(kwargs.items()))
                cache_key = f"{key_prefix}:{func.__name__}:{arg_str}:{kwarg_str}"
            
            # Try to get from cache
            cached = collection.find_one({"_id": cache_key})
            
            if cached:
                # Ensure expires_at is timezone-aware
                expires_at = cached["expires_at"]
                if expires_at.tzinfo is None:
                    expires_at = expires_at.replace(tzinfo=timezone.utc)
                
                # Check if cache is still valid
                if datetime.now(timezone.utc) < expires_at:
                    try:
                        data = cached["data"]
                        logger.info(f"Cache hit for {cache_key!r} (collection: {collection_name})")

                        if object_builder:
                            return object_builder(data["model_data"])

                        elif isinstance(data, dict) and "model_type" in data:
                            model_class = globals()[data["model_type"]]
                            logger.info(f"Model class: {model_class}")

                            if issubclass(model_class, BaseModel):
                                return model_class.model_validate(data["model_data"])
                        else:
                            raise ValueError("No object builder or model type found in cache data")

                        return data
                    except Exception as err:
                        logger.error(f"Error validating cache data for {cache_key!r}: {err}", exc_info=True)
                        collection.delete_one({"_id": cache_key})
            
            # If not in cache or expired, call original function
            result = func(*args, **kwargs)
            
            # Cache the result
            if result is not None:
                cache_data = {
                    "data": result,
                    "expires_at": datetime.now(timezone.utc) + timedelta(seconds=ttl_seconds)
                }
                
                # Special handling for Pydantic models
                if isinstance(result, BaseModel):
                    cache_data["data"] = {
                        "model_type": result.__class__.__name__,
                        "model_data": result.model_dump()
                    }
                
                collection.update_one(
                    {"_id": cache_key},
                    {"$set": cache_data},
                    upsert=True
                )
            
            return result

        async def _cache_wrapper_async(func: F, *args: Any, **kwargs: Any) -> Any:
            db = get_mongo_database("cache")
            collection = db.get_collection(collection_name)

            # Build cache key
            if key_builder:
                cache_key = f"{key_prefix}:{key_builder(*args, **kwargs)}"
            else:
                # Default key building from function name and arguments
                arg_str = ":".join(str(arg) for arg in args)
                kwarg_str = ":".join(f"{k}={v}" for k, v in sorted(kwargs.items()))
                cache_key = f"{key_prefix}:{func.__name__}:{arg_str}:{kwarg_str}"
            
            # Try to get from cache
            cached = collection.find_one({"_id": cache_key})
            
            if cached:
                # Ensure expires_at is timezone-aware
                expires_at = cached["expires_at"]
                if expires_at.tzinfo is None:
                    expires_at = expires_at.replace(tzinfo=timezone.utc)
                
                # Check if cache is still valid
                if datetime.now(timezone.utc) < expires_at:
                    try:
                        data = cached["data"]
                        logger.info(f"Cache hit for {cache_key!r} (collection: {collection_name})")
                        # Handle Pydantic models
                        if object_builder:
                            return object_builder(data["model_data"])

                        elif isinstance(data, dict) and "model_type" in data:
                            model_class = globals()[data["model_type"]]
                            logger.info(f"Model class: {model_class}")

                            if issubclass(model_class, BaseModel):
                                return model_class.model_validate(data["model_data"])
                        else:
                            raise ValueError("No object builder or model type found in cache data")

                        return data
                    except Exception as err:
                        logger.error(f"Error validating cache data for {cache_key!r}: {err}", exc_info=True)
                        collection.delete_one({"_id": cache_key})

            # If not in cache or expired, call original function
            result = await func(*args, **kwargs)

            # Cache the result
            if result is not None:
                cache_data = {
                    "data": result,
                    "expires_at": datetime.now(timezone.utc) + timedelta(seconds=ttl_seconds)
                }

                # Special handling for Pydantic models
                if isinstance(result, BaseModel):
                    cache_data["data"] = {
                        "model_type": result.__class__.__name__,
                        "model_data": result.model_dump()   
                    }

                collection.update_one(
                    {"_id": cache_key},
                    {"$set": cache_data},
                    upsert=True
                )

            return result
        
        return async_wrapper if inspect.iscoroutinefunction(func) else sync_wrapper
    
    return decorator

def set_cache_value(
    collection_name: str, 
    key: str, 
    value: Any, 
    ttl_seconds: int = 3600 * 24 * 30,
    force_update: bool = False,
) -> None:
    db = get_mongo_database("cache")
    collection = db.get_collection(collection_name)
    
    if not force_update and collection.find_one({"_id": key}):
        return

    collection.update_one(
        {"_id": key},
        {"$set": {
            "data": pickle.dumps(value), 
            "expires_at": datetime.now(timezone.utc) + timedelta(seconds=ttl_seconds)}
        },
        upsert=True
    )

def get_cached_value(collection_name: str, key: str) -> Any:
    db = get_mongo_database("cache")
    collection = db.get_collection(collection_name)
    cached = collection.find_one({"_id": key})

    if cached:
        return pickle.loads(cached["data"])
    
    return None

def delete_cached_value(collection_name: str, key: str) -> None:
    db = get_mongo_database("cache")
    collection = db.get_collection(collection_name)
    collection.delete_one({"_id": key})