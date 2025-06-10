import json_repair
import json
from pydantic import BaseModel
from typing import Optional

class AgentPad(BaseModel):
    thought: str = ""
    final_answer: Optional[str] = None
    action: Optional[str] = None
    action_input: Optional[str] = None

def parse_conversational_react_response(response: str) -> AgentPad:

    try:
        json_response = json_repair.repair_json(response, return_objects=True)
    except json.JSONDecodeError:
        return AgentPad()

    segment_pad = AgentPad()

    if "thought" in json_response:
        segment_pad.thought = json_response["thought"]

    if "final_answer" in json_response:
        segment_pad.final_answer = json_response["final_answer"]
        return segment_pad

    if "action" in json_response:
        segment_pad.action = json_response["action"]

        if "action_input" not in json_response:
            json_response["action_input"] = ""

    if "action_input" in json_response:
        segment_pad.action_input = json_response["action_input"]

    return segment_pad

async def run_react_agent(task):
    pass