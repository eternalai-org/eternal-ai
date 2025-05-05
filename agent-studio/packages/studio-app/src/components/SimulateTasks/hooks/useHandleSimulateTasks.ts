import { useState } from "react";
import useStudioAgentStore from "../../../stores/useStudioAgentStore";
import { CATEGORY_OPTION_KEYS } from "../../../constants/category-option-keys";
import AgentAPI from "../../../services/apis/agent";

export default function useHandleSimulateTasks() {
  const { simulatePrompt } = useStudioAgentStore();

  const [isLoading, setIsLoading] = useState(false);

  const [data, setData] = useState<any>({
    upperData: null,
    lowerData: null,
  });

  const getContent = (newsPrompt?: string[]) => {
    let content = simulatePrompt?.simulate_prompt || "";

    const news = newsPrompt?.join("\n") || "";

    if (
      simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post ||
      simulatePrompt?.simulate_type ===
        CATEGORY_OPTION_KEYS.missionOnFarcaster.mission_on_farcaster_post
    ) {
      content = `${simulatePrompt?.simulate_prompt}.
Tweet in 256 characters or less. Return the tweet-ready content only. No introduction. No gifs.`;
    }

    if (
      simulatePrompt?.simulate_type ===
      CATEGORY_OPTION_KEYS.missionOnX.mission_on_x_post_news
    ) {
      content = `${!!news ? `Content to use:` : ""}
${news || ""}
    
${simulatePrompt?.simulate_prompt}
`;
    }

    return content;
  };

  const handleSubmitSimulate = async (
    simulateContent?: string,
    newsPrompt?: string[]
  ) => {
    if (!simulatePrompt?.personality) return;

    try {
      setIsLoading(true);

      const historyMessages = [
        {
          role: "system",
          content: simulatePrompt?.personality || "",
        },
        {
          role: "user",
          content: getContent(newsPrompt),
        },
      ];

      const responseMessage = await AgentAPI.chatCompletions({
        messages: historyMessages,
      });

      const _content = !!newsPrompt
        ? `${!!newsPrompt ? `Content to use:` : ""}
${newsPrompt || ""}
          
${simulateContent}`
        : simulateContent;

      setData({
        upperData: {
          avatar: "",
          name: "You",
          content: _content,
        },
        lowerData: {
          avatar: "",
          name: simulatePrompt?.agent_name || "Agent",
          content: responseMessage || "Something went wrong!",
        },
      });
    } catch (error) {
    } finally {
      setIsLoading(false);
    }
  };

  return {
    isLoading,
    setIsLoading,
    handleSubmitSimulate,
    data,
    setData,
  };
}
