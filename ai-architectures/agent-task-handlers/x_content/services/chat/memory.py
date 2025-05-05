from typing import List
from langchain.schema import ChatMessage
from langgraph.store.memory import BaseStore

MEMORY_TEMPLATE = """<memories>
{memories}
</memories>"""


def get_memory_namespace(user_id: str) -> tuple:
    return (user_id, "memories")


async def get_formatted_memories(
    store: BaseStore, user_id: str, messages: List[ChatMessage]
):
    memories = await store.asearch(
        get_memory_namespace(user_id),
        query=str([m.content for m in messages[-3:]]),
        limit=10,
    )

    formatted_memories = "\n".join(
        f"[{mem.key}]: {mem.value} (similarity: {mem.score})"
        for mem in memories
    )
    if formatted_memories:
        formatted_memories = MEMORY_TEMPLATE.format(
            memories=formatted_memories
        )

    return formatted_memories
