package clanker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
	ApiKey  string
}

func NewClankerClient(apiKey, apiUrl string) *Client {
	return &Client{
		ApiKey:  apiKey,
		BaseURL: apiUrl,
	}
}

func (c *Client) buildUrl(resourcePath string) string {
	if resourcePath != "" {
		return c.BaseURL + "/" + resourcePath
	}
	return c.BaseURL
}

func (c *Client) doWithoutAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) methodJSON(method string, apiURL string, jsonObject interface{}, result interface{}) error {
	var buffer io.Reader
	if jsonObject != nil {
		bodyBytes, _ := json.Marshal(jsonObject)
		buffer = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(method, apiURL, buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", c.ApiKey)
	resp, err := c.doWithoutAuth(req)
	if err != nil {
		return err
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

type DeployTokenReq struct {
	Name                     string   `json:"name"`
	Symbol                   string   `json:"symbol"`
	Description              string   `json:"description"`
	Image                    string   `json:"image"`
	RequestKey               string   `json:"requestKey"`
	RequestorAddress         string   `json:"requestorAddress"`
	SocialMediaUrls          []string `json:"socialMediaUrls"`
	Platform                 string   `json:"platform"`
	CreatorRewardsPercentage float64  `json:"creatorRewardsPercentage"`
	CreatorRewardsAdmin      string   `json:"creatorRewardsAdmin"`
}

type DeployTokenResp struct {
	ID              int    `json:"id"`
	CreatedAt       string `json:"created_at"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	RequestorFID    int    `json:"requestor_fid"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	ImgURL          string `json:"img_url"`
	PoolAddress     string `json:"pool_address"`
	CastHash        string `json:"cast_hash"`
	Type            string `json:"type"`
	Pair            string `json:"pair"`
	PresaleID       int    `json:"presale_id"`
}

func (c *Client) DeployToken(req *DeployTokenReq) (*DeployTokenResp, error) {
	resp := &DeployTokenResp{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("tokens/deploy"),
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
