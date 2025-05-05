package privy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL   string
	AppID     string
	AppSecret string
}

func NewPrivyClient(appID, appSecret string) *Client {
	return &Client{
		AppID:     appID,
		AppSecret: appSecret,
		BaseURL:   "https://auth.privy.io/api/v1",
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
	req.Header.Add("privy-app-id", c.AppID)
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

type LinkedAccountReq struct {
	Type     string `json:"type"`
	Subject  string `json:"subject"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type LinkedAccountResp struct {
	Address  string `json:"address"`
	Type     string `json:"type"`
	Subject  string `json:"subject"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type CreateUserReq struct {
	CreateEthereumWallet bool               `json:"create_ethereum_wallet"`
	LinkedAccounts       []LinkedAccountReq `json:"linked_accounts"`
}

type CreateUserResp struct {
	ID              string              `json:"id"`
	LinkedAccounts  []LinkedAccountResp `json:"linked_accounts"`
	WalletAddresses []string            `json:"wallet_addresses"`
}

func (c *Client) CreateUser(req *CreateUserReq) (*CreateUserResp, error) {
	resp := &CreateUserResp{}

	// Create request with basic auth
	err := c.methodJSONWithAuth(
		http.MethodPost,
		c.buildUrl("users"),
		req,
		resp,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateUserEx(req *CreateUserReq, appID, appSecret string) (*CreateUserResp, error) {
	resp := &CreateUserResp{}

	// Create request with basic auth
	err := c.methodJSONWithAuthEx(
		http.MethodPost,
		c.buildUrl("users"),
		appID,
		appSecret,
		req,
		resp,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) methodJSONWithAuth(method string, apiURL string, jsonObject interface{}, result interface{}) error {
	var buffer io.Reader
	if jsonObject != nil {
		bodyBytes, _ := json.Marshal(jsonObject)
		buffer = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(method, apiURL, buffer)
	if err != nil {
		return err
	}

	// Add basic auth
	req.SetBasicAuth(c.AppID, c.AppSecret)

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("privy-app-id", c.AppID)

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

func (c *Client) methodJSONWithAuthEx(method, apiURL, appID, appSecret string, jsonObject interface{}, result interface{}) error {
	var buffer io.Reader
	if jsonObject != nil {
		bodyBytes, _ := json.Marshal(jsonObject)
		buffer = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(method, apiURL, buffer)
	if err != nil {
		return err
	}

	// Add basic auth
	req.SetBasicAuth(appID, appSecret)

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("privy-app-id", appID)

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
