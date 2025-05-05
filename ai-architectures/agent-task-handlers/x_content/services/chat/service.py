from x_content.llm.base import OpenAILLMBase
from x_content.models import AgentKnowledgeBase, ChatRequest
from x_content.services.chat.chat_langgraph import get_chat_langgraph
from langchain.schema import ChatMessage


class ChatService:

    def __init__(self):
        pass

    async def get_chat(
        self,
        chat_request: ChatRequest,
        llm: OpenAILLMBase,
        kn_base: AgentKnowledgeBase,
    ) -> str:
        graph = get_chat_langgraph(llm)

        messages = chat_request.messages.copy()
        if len(messages) == 0:
            raise ValueError("Input messages must not be empty")

        if messages[-1].role != "user":
            raise ValueError("Input messages must end with a user message")

        if messages[0].role == "system":
            messages = messages[1:]
        if len(messages) == 0:
            raise ValueError(
                "Input messages must have a least one user message"
            )
        # messages.insert(
        #     0,
        #     ChatMessage(
        #         role="system", content=chat_request.agent_meta_data.persona
        #     ),
        # )

        events = graph.astream(
            {"messages": messages},
            config={
                "configurable": {
                    "thread_id": chat_request.id,
                    "user_id": chat_request.user_address,
                    "persona": chat_request.agent_meta_data.persona,
                }
            },
        )

        events = [x async for x in events]

        if len(events) == 0:
            raise Exception("No result return from langgraph")

        result = ""
        for event in events:
            for value in event.values():
                result = value["messages"][-1].content

        return result


chat_service = ChatService()
