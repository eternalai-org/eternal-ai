from pydantic import BaseModel, Field
from typing import Optional
from .utils import random_uuid
from .services import EvaluationRequest
from enum import Enum

class EvaluationStage(str, Enum):
    RECEIVED = "received"
    PROCESSING = "processing"
    COMPLETED = "completed"
    FAILED = "failed"


class EvaluationState(BaseModel):
    id: str = Field(default_factory=random_uuid)
    request: EvaluationRequest

    launchpad_id: str
    