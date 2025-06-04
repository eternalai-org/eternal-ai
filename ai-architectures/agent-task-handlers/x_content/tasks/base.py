from abc import abstractmethod, ABC
from x_content.cache.base_state_handler import StatusHandlerBase
from x_content.cache.chat_request_state_handler import ChatRequestStateHandler
from x_content.constants import MissionChainState
from x_content.models import AgentKnowledgeBase, AutoAgentTask, ChatRequest, ReasoningLog
from x_content.llm import OpenAILLMBase
from x_content.cache import mission_state_handler
from x_content.wrappers import redis_wrapper
from x_content.toolcall.toolcall import IToolCall
from .utils import a_move_state
import traceback

from typing import TypeVar, Generic

T = TypeVar("T", bound=AutoAgentTask)


class BaseTask(ABC, Generic[T]):
    handler: StatusHandlerBase[T]
    resumable = False

    @abstractmethod
    async def process_task(self, task: T) -> T:
        raise NotImplementedError(
            "process_task method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    async def commit_log(self, task: T) -> T:
        return await self.handler.acommit(task)

    async def run(self, task: T) -> T:

        while True:
            try:
                task = await self.process_task(task) or task

            except Exception as err:
                task = await a_move_state(
                    task, MissionChainState.ERROR, f"Error: {err}"
                )
                traceback.print_exc()

            finally:
                await self.commit_log(task)

            if task.state in [MissionChainState.ERROR, MissionChainState.DONE]:
                break

        return task


class MultiStepTaskBase(BaseTask[ReasoningLog]):
    handler = mission_state_handler.MissionStateHandler()

    def __init__(
        self,
        llm: OpenAILLMBase,
        kn_base: AgentKnowledgeBase,
        toolcall: IToolCall = [],
        *args,
        **kwargs,
    ):
        super().__init__()

        self.llm: OpenAILLMBase = llm
        self.kn_base: AgentKnowledgeBase = kn_base
        self.toolcall: IToolCall = toolcall


class ChatTaskBase(BaseTask[ChatRequest]):
    handler = ChatRequestStateHandler()

    def __init__(
        self,
        llm: OpenAILLMBase,
        kn_base: AgentKnowledgeBase,
        *args,
        **kwargs,
    ):
        super().__init__()

        self.llm: OpenAILLMBase = llm
        self.kn_base: AgentKnowledgeBase = kn_base
