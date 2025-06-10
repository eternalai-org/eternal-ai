from pydantic import BaseModel, model_validator
from typing import TypeVar, Generic, Optional
from enum import Enum

_generic_type = TypeVar('_generic_type')

class APIStatus(str, Enum):
    OK = "ok"
    ERROR = "error"
    NOT_FOUND = "not_found"

class ResponseMessage(BaseModel, Generic[_generic_type]):
    result: Optional[_generic_type] = None
    error: Optional[str] = None
    status: APIStatus = APIStatus.OK
    
    @model_validator(mode="after")
    def refine_status(self):
        if self.error is not None:
            self.status = APIStatus.ERROR
            
        return self
    
