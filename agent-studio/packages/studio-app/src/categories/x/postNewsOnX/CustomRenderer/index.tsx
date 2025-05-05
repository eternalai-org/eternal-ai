import { StudioCategoryOptionRenderPayload } from "@agent-studio/studio-dnd";
import DetailsField from "./Fields/DetailsField";
import { PostNewsOnXFormData, PostNewsTopics } from "../types";
import FrequencyField from "./Fields/FrequencyField";
import { Button, Flex } from "@chakra-ui/react";
import { useMemo, useState } from "react";
import s from "../../styles.module.scss";
import SearchTopicField from "./Fields/SearchTopicField";
import AiModelField from "./Fields/AiModelField";
import useStudioAgentStore from "../../../../stores/useStudioAgentStore";
import { useDetectChainIdForMissionModel } from "../../../../hooks/useDetectChainIdForMissionModel";
import {
  findById,
  findParentById,
  findPersonalityObj,
} from "../../../../utils/process";
import { showError } from "../../../../utils/toast";
import CustomRendererBase from "../../../../components/CustomRendererBase";
import StudioFormWrapper from "../../../../components/form";

const CustomPostNewsOnXRenderer = ({
  id,
  formData,
  data,
  setFormFields,
}: StudioCategoryOptionRenderPayload<PostNewsOnXFormData>) => {
  const { frequency, details, topics, model } = formData;

  const [topic, setTopic] = useState("");

  const { setSimulatePrompt, simulatePrompt } = useStudioAgentStore();
  const node = useDetectChainIdForMissionModel(id);

  const getRandomValue = (arr: string[]): string => {
    const randomIndex = Math.floor(Math.random() * arr.length);
    return arr[randomIndex];
  };

  const getTopic = () => {
    const _topics = topics.values.split(",");
    const randomTopic = getRandomValue(_topics);

    return randomTopic.trim();
  };

  const handleSelectMissionPrompt = () => {
    const result = findById(data, id);
    const parentObj = findParentById(data, id);

    const personalityObj = findPersonalityObj(data, parentObj);

    let _topic = topic;

    if (!result.data.details) {
      showError({
        message: "Please fill the details field",
      });
      return;
    }

    if (!personalityObj) {
      showError({
        message:
          "Personality not found. Please connect the prompt to a personality",
      });
      return;
    }

    if (!_topic) {
      _topic = getTopic();
    }

    setSimulatePrompt({
      id: [personalityObj.id, result.id],
      personality: personalityObj?.data?.personality,
      simulate_prompt: result.data.details,
      simulate_type: result.idx,
      topics: simulatePrompt?.topics || { ...topics, values: _topic },
    });
  };

  const isActive = useMemo(() => {
    return simulatePrompt?.id.includes(id);
  }, [simulatePrompt]);

  return (
    <div data-lego-class={isActive && s.active}>
      <CustomRendererBase
        tag="Ability"
        title="Generate posts from specific topics"
      >
        <StudioFormWrapper>
          <AiModelField
            id={id}
            value={model as string}
            onChange={(id: string, name: string) => {
              setFormFields({ model: id, modelName: name });
            }}
            data={data}
          />

          <FrequencyField
            id={id}
            value={frequency as string}
            onChange={(v: string) => {
              setFormFields({ frequency: v });
            }}
          />
          <SearchTopicField
            id={id}
            data={data}
            topic={topic}
            setTopic={setTopic}
            value={
              topics || {
                values: "",
                bingSearch: false,
                twitterSearch: false,
              }
            }
            onChange={(v: PostNewsTopics) => {
              setFormFields({ topics: v });
            }}
          />

          <DetailsField
            id={id}
            value={details as string}
            onChange={(v: string) => {
              setFormFields({ details: v });
            }}
          />
          <Flex justifyContent={"flex-end"}>
            <Button
              borderRadius={"100px"}
              width={"fit-content"}
              mb="8px"
              onClick={handleSelectMissionPrompt}
            >
              Simulate
            </Button>
          </Flex>
        </StudioFormWrapper>
      </CustomRendererBase>
    </div>
  );
};

export default CustomPostNewsOnXRenderer;
