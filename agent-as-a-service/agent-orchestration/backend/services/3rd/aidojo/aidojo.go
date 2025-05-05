package aidojo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AiDojoBackend struct {
	BaseURL string
	ApiKey  string
}

type ChatResponse struct {
	Id      string `json:"id"`
	Choices []*struct {
		Message *struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	OnchainData struct {
		InferTx      string `json:"infer_tx"`
		SubmitTx     string `json:"submit_tx"`
		SeizeMinerTx string `json:"seize_miner_tx"`
	} `json:"onchain_data"`
}

func NewAiDojoBackend(baseUrl, apiKey string) *AiDojoBackend {
	return &AiDojoBackend{
		BaseURL: baseUrl,
		ApiKey:  apiKey,
	}
}

func (c *AiDojoBackend) postFormData(apiURL string, headers map[string]string, jsonObject map[string]interface{}, result interface{}) error {
	data := url.Values{}
	for key, val := range jsonObject {
		data.Set(key, fmt.Sprintf("%v", val))
	}
	encodedData := data.Encode()

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(encodedData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Authorization", fmt.Sprintf("TK1 %s", c.ApiKey))
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *AiDojoBackend) postJSON(apiURL string, headers map[string]string, jsonObject interface{}, result interface{}) error {
	bodyBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *AiDojoBackend) putJSON(apiURL string, headers map[string]string, jsonObject interface{}, result interface{}) error {
	bodyBytes, _ := json.Marshal(jsonObject)
	req, err := http.NewRequest(http.MethodPatch, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *AiDojoBackend) getJSON(url string, headers map[string]string, result interface{}) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, c.ApiKey))
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return resp.StatusCode, json.NewDecoder(resp.Body).Decode(result)
	}
	return resp.StatusCode, nil
}

func (c *AiDojoBackend) getBytes(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request: %v", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	return bodyBytes, nil
}

func (c *AiDojoBackend) LighthouseUpload(content string) (string, error) {
	var resp struct {
		Data string `json:"data"`
	}
	err := c.postJSON(
		fmt.Sprintf("%s/api/service/light-house/upload", c.BaseURL),
		map[string]string{
			"Authorization": fmt.Sprintf("TK1 %s", c.ApiKey),
		},
		map[string]string{
			"content": content,
		},
		&resp,
	)
	if err != nil {
		return "", err
	}
	return resp.Data, nil
}

type AgentType int

const (
	AgentTypeNormal         AgentType = 0
	AgentTypeReasoningAgent AgentType = 1
)

type AgentMetadataRequest struct {
	TokenInfo struct {
		Name    string `json:"name"`
		Symbol  string `json:"symbol"`
		Address string `json:"address"`
		Chain   string `json:"chain"`
	} `json:"token_info"`
}

func (c *AiDojoBackend) AgentBatchPromptItem(chainID, agentID string, outputMaxCharacter uint, agentType AgentType, toolset string, headSystemPrompt string, agentMetaData *AgentMetadataRequest, messages []*models.UserAgentInferDataItem, skipThough bool, toolList string) (string, error) {
	postData := map[string]interface{}{
		"chain_id":             chainID,
		"contract_agent_id":    agentID,
		"output_max_character": outputMaxCharacter,
		"agent_type":           agentType,
		"toolset":              toolset,
		"twitter_snapshot":     "",
		"head_system_prompt":   headSystemPrompt,
		"messages":             messages,
		"skip_though":          skipThough,
		"agent_meta_data":      agentMetaData,
		"tool_list":            toolList,
	}
	var resp struct {
		Data struct {
			Id string `json:"id"`
		} `json:"data"`
	}
	err := c.postJSON(
		fmt.Sprintf("%s/api/agent/batch-prompt-item", c.BaseURL),
		map[string]string{
			"Authorization": fmt.Sprintf("TK1 %s", c.ApiKey),
		},
		postData,
		&resp,
	)
	if err != nil {
		return "", err
	}
	if resp.Data.Id == "" {
		return "", errors.New("id not found")
	}
	return resp.Data.Id, nil
}

type AgentInscribeReq struct {
	Content     string    `json:"content"`
	MysqlID     uint      `json:"mysql_id"`
	TweetID     string    `json:"tweet_id"`
	AgentID     uint      `json:"agent_id"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	PostTweetAt time.Time `json:"post_tweet_at"`
}

func (c *AiDojoBackend) AgentBatchPromptItemOutput(inferId string) (string, string, error) {
	var resp struct {
		Data struct {
			Id             string `json:"id"`
			PromptOutput   string `json:"prompt_output"`
			InscribeTxHash string `json:"inscribe_tx_hash"`
		} `json:"data"`
	}
	_, err := c.getJSON(
		fmt.Sprintf("%s/api/agent/get-batch-item-output/%s", c.BaseURL, inferId),
		map[string]string{
			"Authorization": fmt.Sprintf("TK1 %s", c.ApiKey),
		},
		&resp,
	)
	if err != nil {
		return "", "", err
	}
	if resp.Data.Id == "" {
		return "", "", errors.New("id not found")
	}
	return resp.Data.PromptOutput, resp.Data.InscribeTxHash, nil
}

func (c *AiDojoBackend) OffchainAgentOutput(inferId string) (string, error) {
	data, err := c.getBytes(
		fmt.Sprintf("%s/api/agent/offchain-auto-agent-request/%s", c.BaseURL, inferId),
		map[string]string{
			"Authorization": fmt.Sprintf("TK1 %s", c.ApiKey),
		},
	)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *AiDojoBackend) OffchainAutoAgentOutput(endpoint string, funId string, xToken string) (string, error) {
	data, err := c.getBytes(
		fmt.Sprintf(endpoint+"/async/get?id=%s", funId),
		map[string]string{
			"x-token": xToken,
		},
	)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *AiDojoBackend) AgentMintNft(chainID uint64, agentID string) error {
	postData := map[string]interface{}{
		"agent_id": agentID,
		"chain_id": fmt.Sprintf("%d", chainID),
	}
	err := c.postJSON(
		fmt.Sprintf("%s/api/agent/mint?admin_key=eai2024", c.BaseURL),
		map[string]string{},
		postData,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *AiDojoBackend) GetAgentPoolBalance(chainID string, nftIds []string) (map[string]string, error) {
	postData := map[string]interface{}{
		"nft_ids":  nftIds,
		"chain_id": chainID,
	}
	var resp struct {
		Data map[string]string `json:"data"`
	}
	err := c.postJSON(
		fmt.Sprintf("%s/api/agent/pool_balance", c.BaseURL),
		map[string]string{
			"Authorization": fmt.Sprintf("TK1 %s", c.ApiKey),
		},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *AiDojoBackend) GetBitcoinTxHash(txHash string) (string, error) {
	var resp struct {
		Btc string `json:"btc"`
	}
	_, err := c.getJSON(
		fmt.Sprintf("https://webstat.shard-ai.l2aas.com/get-btctx?tc=%s", txHash),
		map[string]string{},
		&resp,
	)
	if err != nil {
		return "", err
	}
	return resp.Btc, nil
}

func (c *AiDojoBackend) AgentInscribe(postData *AgentInscribeReq) (string, error) {
	var resp struct {
		Data struct {
			Id string `json:"id"`
		} `json:"data"`
	}
	err := c.postJSON(
		fmt.Sprintf("%s/api/inscribe", c.BaseURL),
		map[string]string{
			"Authorization": fmt.Sprintf("TK1 %s", c.ApiKey),
		},
		postData,
		&resp,
	)
	if err != nil {
		return "", err
	}
	if resp.Data.Id == "" {
		return "", errors.New("id not found")
	}
	return resp.Data.Id, nil
}

type GenerateImageResponse struct {
	Output *struct {
		Result string `json:"result"`
	} `json:"output"`
}

func (c AiDojoBackend) GenerateImage(systemContent, baseUrl string) (string, error) {
	bodyReq := map[string]interface{}{
		"input": map[string]interface{}{
			"prompt": systemContent,
			"steps":  10,
			"seed":   time.Now().Unix(),
		},
	}

	chatResp := ""
	bodyBytes, _ := json.Marshal(bodyReq)
	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return chatResp, err
	}
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return chatResp, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return chatResp, err
	}

	m := GenerateImageResponse{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return chatResp, err
	}

	if m.Output != nil {
		chatResp = m.Output.Result
	}
	return chatResp, nil
}

type OffchainAutoAgentRequest struct {
	RequestInput    string             `bson:"request_input" json:"request_input"`
	Response        string             `bson:"response" json:"response"`
	ResponseId      string             `bson:"response_id" json:"response_id"`
	ResponseStatus  int                `bson:"response_status" json:"response_status"`
	Url             string             `bson:"url" json:"url"`
	Log             string             `bson:"log" json:"log"`
	AssistantID     primitive.ObjectID `bson:"assistant_id" json:"assistant_id"`
	MissionID       primitive.ObjectID `bson:"mission_id" json:"mission_id"`
	ContractAgentID string             `bson:"contract_agent_id" json:"contract_agent_id"`
	ChainID         string             `bson:"chain_id" json:"chain_id"`
	Output          string             `json:"output" bson:"output"`
	BatchItemInput  string             `json:"batch_item_input" bson:"batch_item_input"`
	Toolset         string             `json:"toolset" bson:"toolset"`
	Task            string             `json:"task" bson:"task"`
}

// ////
type ChainConfigsDetailResp struct {
	ChainId           string            `json:"chain_id"`
	SupportModelNames map[string]string `json:"support_model_names"`
}

type ChainConfigsResp struct {
	Data []*ChainConfigsDetailResp `json:"data"`
}

func (c *AiDojoBackend) GetChainConfigs() ([]*ChainConfigsDetailResp, error) {
	var resp ChainConfigsResp
	_, err := c.getJSON(
		fmt.Sprintf("%s/api/chain-config/list", c.BaseURL),
		map[string]string{},
		&resp,
	)

	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
