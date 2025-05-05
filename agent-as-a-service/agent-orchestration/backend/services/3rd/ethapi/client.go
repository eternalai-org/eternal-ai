package ethapi

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/evmapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/trxapi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateETHAddress() (string, string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
	prk := hex.EncodeToString(key.D.Bytes())
	if len(prk) != 64 {
		return "", "", errors.New("bad private key len")
	}
	return addr, prk, nil
}

type BlockResp struct {
	time uint64
	hash string
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  *ResultResponse `json:"result"`
	ID      int             `json:"id"`
}

type ResultResponse struct {
	Hash          string  `json:"hash"`
	BlockHash     *string `json:"blockHash"`
	BlockNumber   *string `json:"blockNumber"`
	Number        *string `json:"number"`
	ChainID       string  `json:"chainId"`
	L1BlockNumber *string `json:"l1BlockNumber"`
}

func (b *BlockResp) Time() uint64 {
	return b.time
}

func (b *BlockResp) Hash() string {
	return b.hash
}

type Client struct {
	evmapi.BaseClient
	BaseURL     string
	BaseGasURL  string
	BlockMap    map[uint64]*BlockResp
	BlockMapMtx sync.Mutex
	//
	chainID           uint64
	client            *ethclient.Client
	gasClient         *ethclient.Client
	MinGasPrice       string
	BTCL1             bool
	BlockTimeDisabled bool
	//
	InscribeTxsLog func(txHash string, inscribeTxHash string, logErr string)
}

func (c *Client) getClient() (*ethclient.Client, error) {
	if c.client == nil {
		client, err := ethclient.Dial(c.BaseURL)
		if err != nil {
			return nil, err
		}
		c.client = client
	}
	return c.client, nil
}

func (c *Client) getGasClient() (*ethclient.Client, error) {
	if c.gasClient == nil {
		if c.BaseGasURL == "" {
			c.BaseGasURL = c.BaseURL
		}
		client, err := ethclient.Dial(c.BaseGasURL)
		if err != nil {
			return nil, err
		}
		c.gasClient = client
	}
	return c.gasClient, nil
}

func (c *Client) GetClient() (*ethclient.Client, error) {
	if c.client == nil {
		client, err := ethclient.Dial(c.BaseURL)
		if err != nil {
			return nil, err
		}
		c.client = client
	}
	return c.client, nil
}

func (c *Client) ChainID() (uint64, error) {
	_, err := c.GetChainID()
	if err != nil {
		return 0, err
	}
	return c.chainID, nil
}

func (c *Client) PendingNonceAt(ctx context.Context, addr common.Address) (uint64, error) {
	client, err := c.getClient()
	if err != nil {
		return 0, err
	}
	pNum, err := client.PendingNonceAt(ctx, addr)
	if err != nil {
		return 0, err
	}
	num, err := client.NonceAt(ctx, addr, nil)
	if err != nil {
		return 0, err
	}
	if pNum > num+3 {
		return 0, errors.New("pending nonce is too high")
	}
	return pNum, nil
}

func (c *Client) Address() (string, string, error) {
	return CreateETHAddress()
}

func (c *Client) BlockByNumber(blockNumber int64) (*types.Block, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *Client) GetLastBlockNumber() (int64, error) {
	client, err := c.getClient()
	if err != nil {
		return 0, err
	}
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return int64(blockNumber), nil
}

func (c *Client) GetBlockTime(n uint64) (uint64, error) {
	block, err := c.getBlock(n)
	if err != nil {
		return 0, err
	}
	return block.Time(), nil
}

func (c *Client) getBlock(n uint64) (*BlockResp, error) {
	if c.BlockMap == nil {
		c.BlockMap = map[uint64]*BlockResp{}
	}
	c.BlockMapMtx.Lock()
	blockResp, ok := c.BlockMap[n]
	c.BlockMapMtx.Unlock()
	if !ok {
		var blockInfoResp struct {
			Result *struct {
				Timestamp string `json:"timestamp"`
				Hash      string `json:"hash"`
			} `json:"result"`
		}
		err := c.postJSON(
			c.BaseURL,
			map[string]string{},
			map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      1,
				"method":  "eth_getBlockByNumber",
				"params": []interface{}{
					fmt.Sprintf("0x%s", big.NewInt(int64(n)).Text(16)),
					false,
				},
			},
			&blockInfoResp,
		)
		if err != nil {
			return nil, err
		}
		var timeBN *big.Int
		if strings.HasPrefix(blockInfoResp.Result.Timestamp, "0x") {
			timeBN, ok = big.NewInt(0).SetString(blockInfoResp.Result.Timestamp[2:], 16)
			if !ok {
				return nil, errors.New("wrong time")
			}
		} else {
			timeBN, ok = big.NewInt(0).SetString(blockInfoResp.Result.Timestamp, 10)
			if !ok {
				return nil, errors.New("wrong time")
			}
		}
		c.BlockMapMtx.Lock()
		//
		c.BlockMap[n] = &BlockResp{
			time: timeBN.Uint64(),
			hash: blockInfoResp.Result.Hash,
		}
		blockResp = c.BlockMap[n]
		//
		for i := n - 2000; i < n-1000; i++ {
			_, ok := c.BlockMap[i]
			if ok {
				delete(c.BlockMap, i)
			}
		}
		c.BlockMapMtx.Unlock()
	}
	return blockResp, nil
}

func (c *Client) doWithAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) postJSON(apiURL string, headers map[string]string, jsonObject interface{}, result interface{}) error {
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
	resp, err := c.doWithAuth(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.Unmarshal(bodyBytes, result)
	}
	return nil
}

func (c *Client) TransactionConfirmed(hash string) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	if isPending {
		return errors.New("transaction is pending")
	}
	r, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return errors.New("transaction is not Successful")
	}
	return nil
}

func (c *Client) WaitMined(hash string) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	if isPending {
		return errors.New("transaction is pending")
	}
	r, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return errors.New("transaction is not Successful")
	}
	return nil
}

func (c *Client) WaitMinedTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	queryTicker := time.NewTicker(time.Second)
	maxRetry := 100
	defer queryTicker.Stop()
	retry := 0
	for {
		retry++
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil && receipt.BlockNumber != nil {
			return receipt, nil
		}
		_, _, err = c.CheckTxPendingByHash(txHash)
		if errors.Is(err, ethereum.NotFound) {
			return nil, err
		}

		if retry > maxRetry {
			return nil, err
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (c *Client) validateAddress(address string) {
	if !(strings.EqualFold(address, helpers.HexToAddress(address).Hex())) {
		panic("wrong address")
	}
}

func (c *Client) parsePrkAuth(prkHex string) (common.Address, *ecdsa.PrivateKey, error) {
	prkHex = strings.TrimPrefix(prkHex, "0x")
	prk, err := crypto.HexToECDSA(prkHex)
	if err != nil {
		return common.Address{}, nil, err
	}
	pbk, ok := prk.Public().(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, errors.New("error casting public key to ECDSA")
	}
	pbkHex := crypto.PubkeyToAddress(*pbk)
	return pbkHex, prk, nil
}

func (c *Client) GetChainID() (uint64, error) {
	if c.chainID > 0 {
		return c.chainID, nil
	}
	client, err := c.getClient()
	if err != nil {
		return 0, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	c.chainID = chainID.Uint64()
	return c.chainID, nil
}

func (c *Client) getGasPrice() (*big.Int, error) {
	client, err := c.getGasClient()
	if err != nil {
		return nil, err
	}
	gasPrice := common.Big0
	if c.MinGasPrice != "" {
		minGasPriceNum, err := strconv.ParseInt(c.MinGasPrice, 10, 0)
		if err == nil {
			gasPrice = big.NewInt(minGasPriceNum)
		}
	}
	if gasPrice == nil || gasPrice.Cmp(big.NewInt(0)) <= 0 {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
		gasPrice = common.Big0.Quo(
			common.Big0.Mul(gasPrice, big.NewInt(12)),
			big.NewInt(10),
		)
	}
	gasPriceRes, _ := common.Big0.SetString(gasPrice.Text(10), 10)
	return gasPriceRes, nil
}

func (c *Client) getGasTipCap() (*big.Int, error) {
	client, err := c.getGasClient()
	if err != nil {
		return nil, err
	}
	gasPrice := common.Big0
	if c.MinGasPrice != "" {
		minGasPriceNum, err := strconv.ParseInt(c.MinGasPrice, 10, 0)
		if err == nil {
			gasPrice = big.NewInt(minGasPriceNum)
		}
	}
	if gasPrice == nil || gasPrice.Cmp(big.NewInt(0)) <= 0 {
		gasPrice, err = client.SuggestGasTipCap(context.Background())
		if err != nil {
			return nil, err
		}
		if gasPrice.Cmp(big.NewInt(0)) <= 0 {
			gasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				return nil, err
			}
		}
		gasPrice = common.Big0.Quo(
			common.Big0.Mul(gasPrice, big.NewInt(12)),
			big.NewInt(10),
		)
	}
	gasPriceRes, _ := common.Big0.SetString(gasPrice.Text(10), 10)
	return gasPriceRes, nil
}

func (c *Client) GetCachedGasPriceAndTipCap() (*big.Int, *big.Int, error) {
	cachedGasPrice, err := c.getGasPrice()
	if err != nil {
		return nil, nil, err
	}
	cachedGasTipCap := cachedGasPrice
	return cachedGasPrice, cachedGasTipCap, nil
}

func (c *Client) Transfer(prkHex string, toAddr string, amount string, includeFee bool, wait bool) (string, error) {
	c.validateAddress(toAddr)
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
	}
	value, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		return "", errors.New("amount is invalid")
	}
	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	nonce, err := c.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	extraLimit := uint64(0)
	toAddress := helpers.HexToAddress(toAddr)
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: pbkHex,
		To:   &toAddress,
		Data: []byte{},
	})
	if err != nil {
		return "", err
	}
	if includeFee {
		gasFee := big.NewInt(0).Mul(gasPrice, big.NewInt(int64(gasLimit+extraLimit)))
		if value.Cmp(gasFee) <= 0 {
			return "", errors.New("amount is lower than gas fee")
		}
		value = big.NewInt(0).Sub(value, gasFee)
	}
	tx := types.NewTransaction(nonce, helpers.HexToAddress(toAddr), value, gasLimit, gasPrice, []byte{})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), prk)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	//
	if wait {
		c.WaitMined(signedTx.Hash().Hex())
	}
	return signedTx.Hash().Hex(), nil
}

func (c *Client) ValidateMessageSignature(msg string, signatureHex string, signer string) error {
	msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgBytes := []byte(msg)
	msgHash := crypto.Keccak256Hash(
		msgBytes,
	)
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}
	if signature[crypto.RecoveryIDOffset] > 1 {
		signature[crypto.RecoveryIDOffset] -= 27
	}
	sigPublicKey, err := crypto.SigToPub(msgHash.Bytes(), signature)
	if err != nil {
		return err
	}
	pbkHex := crypto.PubkeyToAddress(*sigPublicKey)
	pbStr := pbkHex.Hex()
	fmt.Println(pbStr)
	if !strings.EqualFold(pbStr, signer) {
		return errors.New("not valid signer")
	}
	return nil
}

func (c *Client) ReferralPaymentSignMessage(contractAddr, tokenAddr, accountAddr, prk string, balance *big.Int) (string, error) {
	datas := []byte{}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}

	datas = append(datas, common.HexToHash(contractAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(big.NewInt(int64(chainID)).Bytes()).Bytes()...)
	datas = append(datas, common.HexToHash(tokenAddr).Bytes()...)
	datas = append(datas, common.HexToHash(accountAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(balance.Bytes()).Bytes()...)

	dataByteHash := crypto.Keccak256Hash(
		datas,
	)

	signature, err := c.SignWithEthereum(prk, dataByteHash.Bytes())
	if err != nil {
		return "", err
	}

	return signature, nil
}

func (c *Client) SignWithEthereum(privateKey string, dataBytes []byte) (string, error) {
	signBytes := append([]byte("\x19Ethereum Signed Message:\n32"), dataBytes...)
	hash := crypto.Keccak256Hash(signBytes)
	prk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(hash.Bytes(), prk)
	if err != nil {
		return "", err
	}
	signature[crypto.RecoveryIDOffset] += 27
	sigHex := hexutil.Encode(signature)
	sigHex = sigHex[2:]
	return sigHex, nil
}

func (c *Client) Sign(privateKey string, hash [32]byte) ([]byte, error) {
	prk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	signature, err := crypto.Sign(hash[:], prk)
	if err != nil {
		return nil, err
	}
	signature[crypto.RecoveryIDOffset] += 27
	sigHex := hexutil.Encode(signature)
	sigHex = sigHex[2:]
	return signature, nil
}

func (c *Client) Erc20Transfer(erc20Addr string, prkHex string, toAddr string, amount string) (string, error) {
	if erc20Addr == "" ||
		!common.IsHexAddress(erc20Addr) {
		return "", errors.New("erc20Addr is invalid")
	}
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
	}
	value, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		return "", errors.New("amount is insufficient")
	}
	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	nonceAt, err := c.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(erc20Addr)
	// EstimateGas
	instanceABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"transfer",
		helpers.HexToAddress(toAddr),
		value,
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: pbkHex,
		To:   &contractAddress,
		Data: dataBytes,
	})
	if err != nil {
		return "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonceAt,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Gas:       gasNumber,
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      dataBytes,
	})
	signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(int64(chainID))), prk)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (c *Client) GetRewardSignature(contractAddr, tokenAddr, accountAddr, prk string, balance *big.Int) (string, error) {
	datas := []byte{}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}

	datas = append(datas, common.HexToHash(contractAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(big.NewInt(int64(chainID)).Bytes()).Bytes()...)
	datas = append(datas, common.HexToHash(tokenAddr).Bytes()...)
	datas = append(datas, common.HexToHash(accountAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(balance.Bytes()).Bytes()...)

	dataByteHash := crypto.Keccak256Hash(
		datas,
	)

	signature, err := c.SignWithEthereum(prk, dataByteHash.Bytes())
	if err != nil {
		return "", err
	}

	return signature, nil
}

func (c *Client) getErc20Instance(contractAddr common.Address) (*erc20.Erc20, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc20.NewErc20(contractAddr, client)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *Client) Erc20Allowance(erc20Addr string, addr string, spender string) (*big.Int, error) {
	if !common.IsHexAddress(erc20Addr) {
		return nil, errors.New("erc20Addr is invalid")
	}
	if !common.IsHexAddress(addr) {
		return nil, errors.New("addr is invalid")
	}
	if !common.IsHexAddress(spender) {
		return nil, errors.New("spender is invalid")
	}
	instance, err := c.getErc20Instance(helpers.HexToAddress(erc20Addr))
	if err != nil {
		return nil, err
	}
	balance, err := instance.Allowance(&bind.CallOpts{}, helpers.HexToAddress(addr), helpers.HexToAddress(spender))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) Erc20ApproveMaxData(toAddr string) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return nil, err
	}
	amount, _ := new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	dataBytes, err := instanceABI.Pack(
		"approve",
		helpers.HexToAddress(toAddr),
		amount,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}

func (c *Client) Erc20ApproveMax(erc20Addr string, prkHex string, toAddr string) (string, error) {
	if erc20Addr == "" ||
		!common.IsHexAddress(erc20Addr) {
		return "", errors.New("erc20Addr is invalid")
	}
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
	}
	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	nonceAt, err := c.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(erc20Addr)
	// EstimateGas
	instanceABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return "", err
	}
	amount, _ := new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	dataBytes, err := instanceABI.Pack(
		"approve",
		helpers.HexToAddress(toAddr),
		amount,
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: pbkHex,
		To:   &contractAddress,
		Data: dataBytes,
	})
	if err != nil {
		return "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonceAt,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Gas:       (gasNumber * 12 / 10),
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      dataBytes,
	})
	signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(int64(chainID))), prk)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (c *Client) GetFrom(tx *types.Transaction) (string, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	return from.String(), err
}

func (c *Client) GetFromFromHash(hash string) (string, error) {
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return "", err
	}
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	return from.String(), err
}

type TxResp struct {
	TxHash      string
	FromAddress string
	ToAddrress  string
	Amount      *big.Int
}

func (c *Client) GetReceiveAddrsByHeight(height int64) ([]string, []*TxResp, error) {
	var blockInfoResp struct {
		Result *struct {
			Timestamp    string `json:"timestamp"`
			Hash         string `json:"hash"`
			Transactions []*struct {
				Hash  string  `json:"hash"`
				From  string  `json:"from"`
				To    *string `json:"to"`
				Value *string `json:"value"`
			} `json:"transactions"`
		} `json:"result"`
	}
	err := c.postJSON(
		c.BaseURL,
		map[string]string{},
		map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      time.Now().UnixNano(),
			"method":  "eth_getBlockByNumber",
			"params": []interface{}{
				fmt.Sprintf("0x%s", big.NewInt(height).Text(16)),
				true,
			},
		},
		&blockInfoResp,
	)
	if err != nil {
		return nil, nil, err
	}
	if blockInfoResp.Result == nil {
		return nil, nil, errors.New("not found")
	}
	addrMap := map[string]bool{}
	txResps := []*TxResp{}
	for _, tx := range blockInfoResp.Result.Transactions {
		value := big.NewInt(0)
		to := ""
		if tx.To != nil {
			to = strings.ToLower(*tx.To)
		}
		if tx.Value != nil {
			value, _ = new(big.Int).SetString((*tx.Value)[2:], 16)
		}
		if to != "" &&
			value.Cmp(big.NewInt(0)) > 0 {
			addrMap[to] = true
		}
		txResps = append(txResps, &TxResp{
			TxHash:      tx.Hash,
			FromAddress: tx.From,
			ToAddrress:  to,
			Amount:      value,
		})
	}
	addrs := []string{}
	for addr := range addrMap {
		addrs = append(addrs, strings.ToLower(addr))
	}
	return addrs, txResps, nil
}

func (c *Client) InscribeTxs(txHashs []string) (string, error) {
	if c.BTCL1 {
		txHash, err := func() (string, error) {
			var feeResp struct {
				FastestFee int `json:"fastestFee"`
			}
			err := helpers.CurlURL(
				"https://mempool.space/api/v1/fees/recommended",
				http.MethodGet,
				map[string]string{},
				nil,
				&feeResp,
			)
			if err != nil {
				return "", err
			}
			if feeResp.FastestFee == 0 {
				return "", errs.NewError(fmt.Errorf("fastest_fee is zero"))
			}
			var resp struct {
				Result string `json:"result"`
			}
			err = c.postJSON(
				c.BaseURL,
				map[string]string{},
				map[string]interface{}{
					"id":      1,
					"jsonrpc": "1.0",
					"method":  "eth_inscribeBatchTxsWithTargetFeeRate",
					"params": []interface{}{
						txHashs,
						feeResp.FastestFee,
					},
				},
				&resp,
			)
			if err != nil {
				return "", err
			}
			return resp.Result, nil
		}()
		if err != nil {
			if c.InscribeTxsLog != nil {
				for _, txHash := range txHashs {
					c.InscribeTxsLog(txHash, "", err.Error())
				}
			}
			return "", err
		}
		return txHash, nil
	}
	return "", nil
}

func (c *Client) Transact(contractAddr string, prkHex string, dataBytes []byte, value *big.Int) (string, error) {
	if value == nil {
		value = common.Big0
	}
	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	nonceAt, err := c.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	var contractAddress *common.Address
	//contractAddr = "" , tx create new contract ,tx.To = nil
	if len(contractAddr) > 0 {
		contractAddress = new(common.Address)
		*contractAddress = helpers.HexToAddress(contractAddr)
	}
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  pbkHex,
		To:    contractAddress,
		Data:  dataBytes,
		Value: value,
	})
	if err != nil {
		return "", err
	}
	if contractAddress == nil {
		gasNumber = gasNumber * 2
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonceAt,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Gas:       (gasNumber * 12 / 10),
		To:        contractAddress,
		Data:      dataBytes,
		Value:     value,
	})
	signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(int64(chainID))), prk)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	_, err = c.InscribeTxs([]string{signedTx.Hash().Hex()})
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (c *Client) IsContract(address string) (bool, error) {
	client, err := c.getClient()
	if err != nil {
		return false, err
	}
	bytecode, err := client.CodeAt(context.Background(), helpers.HexToAddress(address), nil) // nil is latest block
	if err != nil {
		return false, err
	}
	isContract := len(bytecode) > 0
	return isContract, nil
}

func (c *Client) CheckTxPendingByHash(txHash common.Hash) (*RPCResponse, bool, error) {
	jsonData := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\",\"params\":[\"" + txHash.Hex() + "\"],\"id\":1}"
	resp, err := http.Post(c.BaseURL, "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("status code :%v = 200", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	var rpcResponse RPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return nil, false, err
	}
	if rpcResponse.Result == nil {
		return nil, false, ethereum.NotFound
	}
	return &rpcResponse, rpcResponse.Result.BlockHash == nil, nil
}

func (c *Client) ConvertAddressForIn(addr string) string {
	chainID, err := c.GetChainID()
	if err != nil {
		panic(err)
	}
	if chainID == 728126428 {
		return trxapi.AddrTronToEvm(addr)
	}
	return addr
}

func (c *Client) ConvertAddressForOut(addr string) string {
	chainID, err := c.GetChainID()
	if err != nil {
		panic(err)
	}
	if chainID == 728126428 {
		return trxapi.AddrEvmToTron(addr)
	}
	return addr
}

func (c *Client) GetSignatureTimestamp(prk string, timestamp int64) (string, error) {
	datas := []byte{}
	datas = append(datas, common.BytesToHash(big.NewInt(int64(timestamp)).Bytes()).Bytes()...)

	dataByteHash := crypto.Keccak256Hash(
		datas,
	)

	signature, err := c.SignWithEthereum(prk, dataByteHash.Bytes())
	if err != nil {
		return "", err
	}

	return signature, nil
}
