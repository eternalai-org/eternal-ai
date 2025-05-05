package moralis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey string
}

func (c *Client) doWithAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) postJSON(apiURL string, headers map[string]string, jsonObject interface{}, result interface{}) error {
	bodyBytes, _ := json.Marshal(jsonObject)
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.doWithAuth(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
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

func (c *Client) getJSON(url string, headers map[string]string, result interface{}) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.doWithAuth(req)
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

type GetNFTsResp struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Cursor   string        `json:"cursor"`
	Result   []interface{} `json:"result"`
	Status   string        `json:"status"`
}

func (c *Client) GetNFTs(chain string, address string, cursor string, limit int) (*GetNFTsResp, error) {
	if limit <= 0 {
		limit = 500
	}
	rs := GetNFTsResp{}
	_, err := c.getJSON(
		fmt.Sprintf("https://deep-index.moralis.io/api/v2/%s/nft?chain=%s&format=decimal&cursor=%s&limit=%d", url.QueryEscape(address), url.QueryEscape(chain), url.QueryEscape(cursor), limit),
		map[string]string{
			"X-API-Key": c.APIKey,
		},
		&rs,
	)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

type TimeRange struct {
	OneHour        float64 `json:"1h"`
	FourHour       float64 `json:"4h"`
	TwelveHour     float64 `json:"12h"`
	TwentyFourHour float64 `json:"24h"`
}

type TrendingToken struct {
	ChainId            string    `json:"chainId"`
	TokenAddress       string    `json:"tokenAddress"`
	Name               string    `json:"name"`
	UniqueName         string    `json:"uniqueName"`
	Symbol             string    `json:"symbol"`
	Decimals           int       `json:"decimals"`
	Logo               string    `json:"logo"`
	UsdPrice           float64   `json:"usdPrice"`
	CreatedAt          int64     `json:"createdAt"`
	MarketCap          float64   `json:"marketCap"`
	LiquidityUsd       float64   `json:"liquidityUsd"`
	Holders            int       `json:"holders"`
	PricePercentChange TimeRange `json:"pricePercentChange"`
	TotalVolume        TimeRange `json:"totalVolume"`
	Transactions       TimeRange `json:"transactions"`
	BuyTransactions    TimeRange `json:"buyTransactions"`
	SellTransactions   TimeRange `json:"sellTransactions"`
	Buyers             TimeRange `json:"buyers"`
	Sellers            TimeRange `json:"sellers"`
}

func (c *Client) GetTrendingTokens(chain string) ([]TrendingToken, error) {
	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2.2/tokens/trending?chain=%s&limit=100", url.QueryEscape(chain))

	var tokens []TrendingToken
	_, err := c.getJSON(
		url,
		map[string]string{
			"X-API-Key": c.APIKey,
		},
		&tokens,
	)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
