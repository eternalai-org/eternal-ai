package blockchainutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type Client struct {
	BaseURL string
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

func (c *Client) SolanaAddress() (string, error) {
	resp := struct {
		Result struct {
			Address string `json:"address"`
		} `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/address"),
		nil,
		&resp,
	)
	if err != nil {
		return "", err
	}
	return resp.Result.Address, nil
}

func (c *Client) SolanaBalance(addr string) (uint64, error) {
	resp := struct {
		Result uint64 `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl(fmt.Sprintf("solana/balance/%s", addr)),
		nil,
		&resp,
	)
	if err != nil {
		return 0, err
	}
	return uint64(resp.Result), nil
}

func (c *Client) SolanaBlockheight() (int64, error) {
	resp := struct {
		Result int64 `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl("solana/blockheight"),
		nil,
		&resp,
	)
	if err != nil {
		return 0, err
	}
	return resp.Result, nil
}

func (c *Client) CleanHtml(htmlString string) (string, error) {
	resp := struct {
		Result string `json:"result"`
	}{}

	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("clean-html"),
		map[string]interface{}{
			"url":       "",
			"html_data": htmlString,
		},
		&resp,
	)
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}

type SolanaCreatePumpfunTokenReq struct {
	Address     string  `json:"address"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Description string  `json:"description"`
	Twitter     string  `json:"twitter"`
	Telegram    string  `json:"telegram"`
	Website     string  `json:"website"`
	Amount      float64 `json:"amount"`
	ImageBase64 string  `json:"image_base64"`
}

type SolanaCreatePumpfunTokenResp struct {
	Signature string `json:"signature"`
	Mint      string `json:"mint"`
}

func (c *Client) SolanaCreatePumpfunToken(req *SolanaCreatePumpfunTokenReq) (*SolanaCreatePumpfunTokenResp, error) {
	resp := struct {
		Result *SolanaCreatePumpfunTokenResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/create-pumfun"),
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaTradePumpfunTokenReq struct {
	Address string  `json:"address"`
	Action  string  `json:"action"`
	Mint    string  `json:"mint"`
	Amount  float64 `json:"amount"`
	Pool    string  `json:"pool"`
}

func (c *Client) SolanaTradePumpfunToken(req *SolanaTradePumpfunTokenReq) (string, error) {
	resp := struct {
		Result string `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/trade-pumfun"),
		req,
		&resp,
	)
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}

type SolanaTokenBalance struct {
	IsNative    bool   `json:"is_native"`
	Mint        string `json:"mint"`
	Owner       string `json:"owner"`
	State       string `json:"state"`
	TokenAmount struct {
		Amount         numeric.BigInt `json:"amount"`
		Decimals       int            `json:"decimals"`
		UIAmount       float64        `json:"ui_amount"`
		UIAmountString string         `json:"ui_amount_string"`
	} `json:"token_amount"`
}

func (c *Client) SolanaGetTokenBalances(address string) ([]*SolanaTokenBalance, error) {
	resp := struct {
		Result []*SolanaTokenBalance `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl("solana/balances/"+address),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaTokenInfoResp struct {
	Data *struct {
		Parsed *struct {
			Info *struct {
				Decimals int `json:"decimals"`
			} `json:"info"`
		} `json:"parsed"`
	} `json:"data"`
}

func (c *Client) SolanaTokenInfo(mint string) (*SolanaTokenInfoResp, error) {
	resp := struct {
		Result *SolanaTokenInfoResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl("solana/token-info/"+mint),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaTokenMetaDataResp struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Mint   string `json:"mint"`
}

func (c *Client) SolanaTokenMetaData(mint string) (*SolanaTokenMetaDataResp, error) {
	resp := struct {
		Result *SolanaTokenMetaDataResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl("solana/token-metadata/"+mint),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaTradeRaydiumTokenReq struct {
	Address    string  `json:"address"`
	InputMint  string  `json:"input_mint"`
	OutputMint string  `json:"output_mint"`
	Slippage   float64 `json:"slippage"`
	Amount     uint64  `json:"amount"`
}

type SolanaTradeRaydiumTokenResp struct {
	OutputAmount numeric.BigInt `json:"output_amount"`
	Signatures   []string       `json:"signatures"`
}

func (c *Client) SolanaTradeRaydiumToken(req *SolanaTradeRaydiumTokenReq) (*SolanaTradeRaydiumTokenResp, error) {
	resp := struct {
		Result *SolanaTradeRaydiumTokenResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/trade-raydium"),
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaComputeRaydiumTokenResp struct {
	SwapType             string         `json:"swap_type"`
	InputMint            string         `json:"input_mint"`
	InputAmount          numeric.BigInt `json:"input_amount"`
	OutputMint           string         `json:"output_mint"`
	OutputAmount         numeric.BigInt `json:"output_amount"`
	OtherAmountThreshold numeric.BigInt `json:"other_amount_threshold"`
	SlippageBps          float64        `json:"slippage_bps"`
	PriceImpactPct       float64        `json:"price_impact_pct"`
	ReferrerAmount       numeric.BigInt `json:"referrer_amount"`
	RoutePlan            []struct {
		PoolID     string         `json:"pool_id"`
		InputMint  string         `json:"input_mint"`
		OutputMint string         `json:"output_mint"`
		FeeMint    string         `json:"fee_mint"`
		FeeRate    int            `json:"fee_rate"`
		FeeAmount  numeric.BigInt `json:"fee_amount"`
	} `json:"route_plan"`
}

func (c *Client) SolanaComputeRaydiumToken(req *SolanaTradeRaydiumTokenReq) (*SolanaComputeRaydiumTokenResp, error) {
	resp := struct {
		Result *SolanaComputeRaydiumTokenResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/compute-raydium"),
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaTransferReq struct {
	ToAddress string  `json:"to_address"`
	Mint      string  `json:"mint"`
	Amount    float64 `json:"amount"`
}

func (c *Client) SolanaTransfer(address string, req *SolanaTransferReq) (string, error) {
	resp := struct {
		Result string `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/transfer/"+address),
		req,
		&resp,
	)
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}

func (c *Client) SolanaValidateAddress(mint string) (bool, error) {
	resp := struct {
		Result bool `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl("solana/validate/"+mint),
		nil,
		&resp,
	)
	if err != nil {
		return false, err
	}
	return resp.Result, nil
}

type SolanaBalanceByToken struct {
	Amount         string  `json:"amount"`
	Decimals       int     `json:"decimals"`
	UIAmount       float64 `json:"uiAmount"`
	UIAmountString string  `json:"uiAmountString"`
}

func (c *Client) SolanaBalanceByToken(address, mint string) (*SolanaBalanceByToken, error) {
	resp := struct {
		Result *SolanaBalanceByToken `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl(fmt.Sprintf(`solana/balance/%s/%s`, mint, address)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type ZoraCreateCoinReq struct {
	Name      string   `json:"name"`
	Symbol    string   `json:"symbol"`
	Uri       string   `json:"uri"`
	Recipient string   `json:"recipient"`
	Deployer  string   `json:"deployer"`
	Owners    []string `json:"owners"`
}

type ZoraCreateCoinResp struct {
	Hash       string `json:"hash"`
	Address    string `json:"address"`
	Deployment struct {
		PayoutRecipient string `json:"payoutRecipient"`
		Caller          string `json:"caller"`
		Uri             string `json:"uri"`
		Coin            string `json:"coin"`
		Pool            string `json:"pool"`
	} `json:"deployment"`
}

func (c *Client) ZoraCreateCoin(req *ZoraCreateCoinReq) (*ZoraCreateCoinResp, error) {
	resp := struct {
		Result *ZoraCreateCoinResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("/zora/create-coin"),
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type SolanaCreateTokenReq struct {
	ProjectID   string `json:"project_id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
	Amount      uint64 `json:"amount"`
}

type SolanaCreateTokenResp struct {
	Signature interface{} `json:"signature"`
	Mint      string      `json:"mint"`
}

func (c *Client) SolanaCreateToken(req *SolanaCreateTokenReq) (*SolanaCreateTokenResp, error) {
	resp := struct {
		Result *SolanaCreateTokenResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("solana/create-token"),
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
