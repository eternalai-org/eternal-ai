import fastapi
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
import asyncio
import os
from app.apis import router as api_router
from fastapi import Request, Response
from typing import Callable
import time
import logging

logging_fmt = "%(asctime)s - %(name)s - %(levelname)s - %(message)s"
logging.basicConfig(level=logging.INFO, format=logging_fmt)
logger = logging.getLogger(__name__)

async def lifespan(app: fastapi.FastAPI):
    yield

def main():
    server_app = fastapi.FastAPI(
        lifespan=lifespan
    )

    server_app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    server_app.include_router(api_router)

    @server_app.get("/health")
    async def healthcheck():
        return {"status": "ok", "message": "Yo, I am alive"}
    
    @server_app.middleware("http")
    async def log_request_processing_time(request: Request, call_next: Callable) -> Response:
        start_time = time.time()
        response: Response = await call_next(request)

        if request.url.path.startswith(api_router.prefix):
            logger.info(f"{request.method} - {request.url.path} - {time.time() - start_time} seconds - {response.status_code}")

        return response

    event_loop = asyncio.new_event_loop()
    asyncio.set_event_loop(event_loop)

    config = uvicorn.Config(
        server_app,
        loop=event_loop,
        host=os.getenv("HOST", "0.0.0.0"),
        port=int(os.getenv("PORT", "80")),
        log_level="error",
        timeout_keep_alive=300,
    )

    server = uvicorn.Server(config)
    event_loop.run_until_complete(server.serve())

if __name__ == '__main__':
    main()