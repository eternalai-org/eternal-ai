package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) JobAgentTwitterPostCreateTokenForImage2Video(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostCreateTokenForImage2Video",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)":                  {[]uint{s.conf.VideoAiAgentInfoId}},
						"status = ?":                            {models.AgentTwitterPostStatusInferSubmitted},
						"post_type = ? and type = ?":            {models.AgentSnapshotPostActionTypeGenerateVideo, models.AgentTwitterPostTypeImage2video},
						"token_name = '' and token_symbol = ''": {},
						"created_at > ?":                        {time.Now().Add(-1 * time.Hour)}, // only submitted before 1 hour
					},
					map[string][]interface{}{},
					[]string{
						"created_at asc",
					},
					0,
					20,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, agentTwitterPost := range twitterPosts {
					err = s.AgentTwitterPostGenerateTokenInfo(ctx, agentTwitterPost.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, agentTwitterPost.ID))
					} else {
						redisKey := s.GetRedisAgentTwitterPostGenerateTokenInfo(agentTwitterPost.ID)
						_ = s.rdb.Set(redisKey, "1", 24*time.Hour).Err()
					}
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}

type VisionCompletionRequest struct {
	Seed     int    `json:"seed"`
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content []struct {
			Type     string `json:"type"`
			Text     string `json:"text,omitempty"`
			ImageUrl struct {
				Url string `json:"url"`
			} `json:"image_url,omitempty"`
		} `json:"content"`
	} `json:"messages"`
	MaxTokens int `json:"max_tokens"`
}

func (s *Service) GetRedisAgentTwitterPostGenerateTokenInfo(id uint) string {
	redisKey := fmt.Sprintf("agent_twitter_post_generate_token_info_%v", id)
	return redisKey
}

func (s *Service) AgentTwitterPostGenerateTokenInfo(ctx context.Context, agentTwitterPostID uint) error {
	redisKey := s.GetRedisAgentTwitterPostGenerateTokenInfo(agentTwitterPostID)
	redisCheckValue, err := s.rdb.Get(redisKey).Result()
	if err == nil && redisCheckValue != "" {
		return nil
	}

	tx := daos.GetDBMainCtx(ctx)
	agentTwitterPost, err := s.dao.FirstAgentTwitterPostByID(tx, agentTwitterPostID, map[string][]interface{}{}, true)
	if err != nil {
		return errs.NewError(err)
	}

	if agentTwitterPost.TokenName != "" || agentTwitterPost.TokenSymbol != "" {
		return nil
	}

	if agentTwitterPost.Type != models.AgentTwitterPostTypeImage2video {
		return nil
	}
	if agentTwitterPost.ExtractMediaContent == "" {
		return nil
	}

	// first get description of image
	visionRequest := &VisionCompletionRequest{
		Model: "Qwen/Qwen2.5-VL-7B-Instruct",
		Messages: []struct {
			Role    string `json:"role"`
			Content []struct {
				Type     string `json:"type"`
				Text     string `json:"text,omitempty"`
				ImageUrl struct {
					Url string `json:"url"`
				} `json:"image_url,omitempty"`
			} `json:"content"`
		}{
			{
				Role: "user",
				Content: []struct {
					Type     string `json:"type"`
					Text     string `json:"text,omitempty"`
					ImageUrl struct {
						Url string `json:"url"`
					} `json:"image_url,omitempty"`
				}{
					{
						Type: "text",
						Text: "What is in this image?",
					},
					{
						Type: "image_url",
						ImageUrl: struct {
							Url string `json:"url"`
						}{
							Url: agentTwitterPost.ExtractMediaContent,
						},
					},
				},
			},
		},
		MaxTokens: 512,
	}

	maxRetry := 5
	imageDescription := ""
	for i := 0; i < maxRetry; i++ {
		time.Sleep(time.Duration(i*10) * time.Second)
		respBytes, _, statusCode, err := helpers.HttpRequest(s.conf.KnowledgeBaseConfig.DirectServiceUrl, "POST",
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %v", s.conf.KnowledgeBaseConfig.OnchainAPIKey),
			}, visionRequest)
		if err != nil {
			continue
		}
		if statusCode != 200 {
			continue
		}
		visionResponse := &VisionResponse{}
		err = json.Unmarshal(respBytes, visionResponse)
		if err != nil {
			continue
		}

		if len(visionResponse.Choices) == 0 {
			continue
		}

		imageDescription = visionResponse.Choices[0].Message.Content
		break
	}

	if imageDescription == "" {
		s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[ERROR_IMAGE_DESC] Image description is empty, db_id=%v", agentTwitterPost.ID))
		return errs.NewError(fmt.Errorf("image description is empty"))
	}

	// using LLM to get token symbol and token name
	type TokenResponse struct {
		TokenSymbol string `json:"token_symbol"`
		TokenName   string `json:"token_name"`
	}

	for i := 0; i < maxRetry; i++ {
		userPrompt := fmt.Sprintf("User's prompt:%v\n\nImage description:%v", agentTwitterPost.ExtractMediaContent, imageDescription)
		request := &openai.ChatCompletionRequest{
			Model: "Llama3.3",
			Messages: []openai.ChatCompletionMessage{
				openai.ChatCompletionMessage{
					Role:    "system",
					Content: "You are an AI assistant tasked with generating a cryptocurrency token symbol and token name based on a given user's prompt and an image description. Prioritize the context of the user's prompt first, but also incorporate elements from the image description to ensure relevance.  \\r\\n\\r\\nThe response must be in strict JSON format with no additional text or explanation. Use the following format:  \\r\\n\\r\\n```json\\r\\n{\\\"token_symbol\\\": \\\"string\\\", \\\"token_name\\\": \\\"string\\\"}\\r\\n```  \\r\\n\\r\\n- **The token_symbol and token_name must NOT be empty.**  \\r\\n- **The token_symbol must be concise (3-5 uppercase letters) and relevant.**  \\r\\n- **The token_name should be creative yet clear.**  \\r\\n- **Avoid generating token_symbol or token_name similar to 'EAI', 'Eternal AI', 'CryptoEternalAI', or 'Crypto Eternal AI'.**  \\r\\n- Ensure that both maintain coherence with the provided inputs.",
				}, openai.ChatCompletionMessage{
					Role:    "user",
					Content: userPrompt,
				},
			},
		}

		respBytes, _, statusCode, err := helpers.HttpRequest(s.conf.KnowledgeBaseConfig.DirectServiceUrl, "POST",
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %v", s.conf.KnowledgeBaseConfig.OnchainAPIKey),
			}, request)
		if err != nil {
			continue
		}
		if statusCode != 200 {
			continue
		}

		chatCompletionResp := &openai.ChatCompletionResponse{}
		err = json.Unmarshal(respBytes, chatCompletionResp)
		if err != nil {
			continue
		}

		if len(chatCompletionResp.Choices) == 0 {
			continue
		}

		chatCompletionResp.Choices[0].Message.Content = strings.ReplaceAll(chatCompletionResp.Choices[0].Message.Content, "```json", "")
		chatCompletionResp.Choices[0].Message.Content = strings.ReplaceAll(chatCompletionResp.Choices[0].Message.Content, "```", "")

		tokenResponse := &TokenResponse{}
		err = json.Unmarshal([]byte(chatCompletionResp.Choices[0].Message.Content), tokenResponse)
		if err != nil {
			continue
		}

		if tokenResponse.TokenName == "" || tokenResponse.TokenSymbol == "" {
			continue
		}

		agentTwitterPost.TokenName = tokenResponse.TokenName
		agentTwitterPost.TokenSymbol = tokenResponse.TokenSymbol

		break
	}

	if agentTwitterPost.TokenName == "" || agentTwitterPost.TokenSymbol == "" {
		s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[ERROR_TOKEN_NAME_SYMBOL] Token name or symbol is empty, db_id=%v", agentTwitterPost.ID))
		return errs.NewError(fmt.Errorf("token name or symbol is empty"))
	}

	return s.dao.Save(tx, agentTwitterPost)
}

// =========
type VisionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role             string        `json:"role"`
			ReasoningContent interface{}   `json:"reasoning_content"`
			Content          string        `json:"content"`
			ToolCalls        []interface{} `json:"tool_calls"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
		StopReason   interface{} `json:"stop_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens        int         `json:"prompt_tokens"`
		TotalTokens         int         `json:"total_tokens"`
		CompletionTokens    int         `json:"completion_tokens"`
		PromptTokensDetails interface{} `json:"prompt_tokens_details"`
	} `json:"usage"`
	PromptLogprobs interface{} `json:"prompt_logprobs"`
}
