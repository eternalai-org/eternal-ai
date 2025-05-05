package eth

import (
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
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"

	"github.com/ethereum/go-ethereum"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const baseFeeWiggleMultiplier = 2

func NewEthClient(rpc string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewEthWsClient(ws string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(ws)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func WaitForTx(client *ethclient.Client, tx common.Hash) error {
	i := 0
	for {
		time.Sleep(2 * time.Second)
		if i > 30 {
			return errors.New("timeout")
		}
		i++
		_, isPending, err := client.TransactionByHash(context.Background(), tx)
		if err != nil {
			continue
		}
		if !isPending {
			time.Sleep(2 * time.Second)
			break
		}
	}
	return nil
}

func WaitForTxRetry(client *ethclient.Client, tx common.Hash, sleep int, retry int) error {
	i := 0
	for {
		time.Sleep(time.Duration(sleep) * time.Second)
		if i > retry {
			return errors.New("timeout")
		}
		i++
		_, isPending, err := client.TransactionByHash(context.Background(), tx)
		if err != nil {
			continue
		}
		if !isPending {
			time.Sleep(time.Duration(sleep) * time.Second)
			break
		}
	}
	return nil
}

func WaitForTxReceipt(client *ethclient.Client, tx common.Hash) (*types.Receipt, error) {
	i := 0
	for {
		time.Sleep(2 * time.Second)
		if i > 20 {
			return nil, errors.New("timeout")
		}
		i++
		txReceipt, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil {
			continue
		}
		if txReceipt != nil {
			time.Sleep(2 * time.Second)
			return txReceipt, nil
		}
	}
}

func WaitForTxReceiptRetry(client *ethclient.Client, tx common.Hash, sleep int, retry int) (*types.Receipt, error) {
	i := 0
	for {
		time.Sleep(time.Duration(sleep) * time.Second)
		if i > retry {
			return nil, errors.New("timeout")
		}
		i++
		txReceipt, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil {
			continue
		}
		if txReceipt != nil {
			time.Sleep(time.Duration(sleep) * time.Second)
			return txReceipt, nil
		}
	}
}

func GetAccountInfo(privKey string) (*ecdsa.PrivateKey, *common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	publicKeyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, &publicKeyAddress, nil
}

func GenerateKeyFromSeedOld(seed string) (string, string, string, error) {
	seedHex := hex.EncodeToString([]byte(seed))
	master, ch := hd.ComputeMastersFromSeed([]byte(seedHex))
	path := "m/44'/1022'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", "", "", err
	}
	var privateKey = secp256k1.GenPrivKeyFromSecret(priv)

	publicKey := privateKey.PubKey()

	privKey := hex.EncodeToString(priv)

	pubKey := hex.EncodeToString(publicKey.Bytes())

	address := "0x" + hex.EncodeToString(publicKey.Address().Bytes())

	return privKey, pubKey, address, nil
}

func GenerateKeyFromSeedNew(seed string) (string, string, string, error) {
	seedHex := hex.EncodeToString([]byte(seed))
	master, ch := hd.ComputeMastersFromSeed([]byte(seedHex))
	path := "m/44'/1022'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", "", "", err
	}

	privateKey, err := crypto.ToECDSA(priv)
	if err != nil {
		return "", "", "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	promptFeeAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	privKey := hex.EncodeToString(priv)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	pubKey := hex.EncodeToString(publicKeyBytes)

	return privKey, pubKey, strings.ToLower(promptFeeAddress.String()), nil
}

func GenerateAddress() (privKey, pubKey, address string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("failed to cast public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

func GenerateAddressFromPrivKey(privKey string) (pubKey, address string, err error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("failed to cast public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

type Client struct {
	eth *ethclient.Client
}

func NewClient(eth *ethclient.Client) *Client {
	return &Client{eth}
}

func (c *Client) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.eth.PendingNonceAt(ctx, address)
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.eth.SuggestGasPrice(ctx)
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.eth.NetworkID(ctx)
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.eth.SendTransaction(ctx, tx)
}

// transfer:
func (c *Client) Transfer(senderPrivKey, receiverAddress string, amount, gasPrice *big.Int, gasLimit, nonce uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(senderPrivKey)
	if err != nil {
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if nonce <= 0 {
		nonce, err = c.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return "", err
		}
	}

	if gasLimit == 0 {
		gasLimit = uint64(21000)
	}

	if gasPrice == nil {
		gasPrice, err = c.SuggestGasPrice(context.Background())
		if err != nil {
			return "", err
		}
	}

	fee := new(big.Int)
	fee.Mul(big.NewInt(int64(gasLimit)), gasPrice)

	toAddress := common.HexToAddress(receiverAddress)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

/*
Copy from the function `(c *BoundContract) transact` in the `eth` package.
*/
func CreateEthTransaction(client *ethclient.Client, from common.Address, to common.Address, value *big.Int, input []byte) (*types.Transaction, error) {

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("NetworkID :%v", err.Error())
	}
	//current block
	head, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("HeaderByNumber:%v , err:%v", chainID, err.Error())
	}
	gasLimit, err := client.EstimateGas(
		context.Background(),
		ethereum.CallMsg{
			From:  from,
			To:    &to,
			Value: value,
			Data:  input,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("err when estimate gas at chainID :%v ,from:%v ,to %v , value:%v, input:%v, err:%v", chainID, from, &to, value.String(), common.Bytes2Hex(input), err.Error())
	}
	// Estimate GasLimit
	if chainID.String() != configs.SubtensorEVMChainID {
		if gasLimit < head.GasLimit/2 {
			gasLimit = gasLimit * 2
		}
	} else {
		gasLimit = uint64(5000000)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, fmt.Errorf("err when PendingNonceAt:%v , err:%v", chainID, err.Error())
	}
	var tx *types.Transaction
	if head.BaseFee != nil {
		// Estimate TipCap
		gasTipCap := big.NewInt(0)
		gasFeeCap := big.NewInt(10000000000)
		if chainID.String() != configs.SubtensorEVMChainID {
			gasTipCap, err = client.SuggestGasTipCap(context.Background())
			if err != nil {
				return nil, fmt.Errorf("err when SuggestGasTipCap:%v , err:%v", chainID, err.Error())
			}
			gasFeeCap = new(big.Int).Add(
				gasTipCap,
				new(big.Int).Mul(head.BaseFee, big.NewInt(baseFeeWiggleMultiplier)),
			)
		} else {
			gasFeeCap, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				return nil, fmt.Errorf("err when SuggestGasPrice:%v , err:%v", chainID, err.Error())
			}
		}
		tx = types.NewTx(&types.DynamicFeeTx{
			To:        &to,
			Nonce:     uint64(nonce),
			GasFeeCap: gasFeeCap,
			GasTipCap: gasTipCap,
			Gas:       gasLimit,
			Value:     value,
			Data:      input,
		})
	} else {
		gasPrice := big.NewInt(10000000000)
		if chainID.String() != configs.SubtensorEVMChainID {
			gasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				return nil, fmt.Errorf("err when SuggestGasPrice:%v , err:%v", chainID, err.Error())
			}
		}
		tx = types.NewTx(&types.LegacyTx{
			To:       &to,
			Nonce:    uint64(nonce),
			GasPrice: gasPrice,
			Gas:      gasLimit,
			Value:    value,
			Data:     input,
		})
	}
	return tx, nil
}

func DecodeSignature(signature string) (r [32]byte, s [32]byte, v byte, err error) {
	if strings.HasPrefix(signature, "0x") {
		signature = signature[2:]
	}
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return r, s, v, fmt.Errorf("err when DecodeSignature:%v , err:%v", signature, err.Error())
	}
	if len(signatureBytes) != 65 {
		return r, s, v, fmt.Errorf("invalid signature length : %v", len(signatureBytes))
	}

	copy(r[:], signatureBytes[:32])
	copy(s[:], signatureBytes[32:64])
	v = signatureBytes[64]
	return r, s, v, nil
}

func CheckValidRpc(rpc string) (bool, error) {
	client, err := ethclient.Dial(rpc)
	defer client.Close()
	if err != nil {
		return false, nil
	}
	_, err = client.ChainID(context.Background())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetLastBlock(rpc string) (uint64, int64, error) {
	start := time.Now()
	response, err := http.DefaultClient.Post(rpc, "application/json", strings.NewReader("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"latest\",false],\"id\":1}"))
	if err != nil {
		return 0, 0, err
	}
	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, 0, err
	}
	if response.StatusCode != 200 {
		return 0, 0, fmt.Errorf("response %v !200 ,body :%v", response.StatusCode, string(res))
	}
	rpcResponse := configs.RPCResponse{}
	err = json.Unmarshal(res, &rpcResponse)
	if err != nil {
		return 0, 0, err
	}
	if rpcResponse.Result == nil {
		return 0, 0, fmt.Errorf("result not found in  :%v", string(res))
	}
	if rpcResponse.Result.Number == nil {
		return 0, 0, fmt.Errorf("block number not found in body :%v", string(res))
	}
	if strings.HasPrefix(*rpcResponse.Result.Number, "0x") {
		*rpcResponse.Result.Number = (*rpcResponse.Result.Number)[2:]
	}
	blockNumber, err := strconv.ParseUint(*rpcResponse.Result.Number, 16, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("block number can not parse in body :%v", string(res))
	}
	return blockNumber, time.Since(start).Milliseconds(), nil
}

func GetCurrentL1BlockNumber(rpc string) (uint64, error) {
	response, err := http.DefaultClient.Post(rpc, "application/json", strings.NewReader("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"latest\",false],\"id\":1}"))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	if response.StatusCode != 200 {
		return 0, fmt.Errorf("response %v !200 ,body :%v", response.StatusCode, string(res))
	}
	rpcResponse := configs.RPCResponse{}
	err = json.Unmarshal(res, &rpcResponse)
	if err != nil {
		return 0, err
	}
	if rpcResponse.Result == nil {
		return 0, fmt.Errorf("result not found in  :%v", string(res))
	}
	if rpcResponse.Result.L1BlockNumber == nil {
		return 0, fmt.Errorf("block number not found in body :%v", string(res))
	}
	if strings.HasPrefix(*rpcResponse.Result.L1BlockNumber, "0x") {
		*rpcResponse.Result.L1BlockNumber = (*rpcResponse.Result.L1BlockNumber)[2:]
	}
	blockNumber, err := strconv.ParseUint(*rpcResponse.Result.L1BlockNumber, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("block number can not parse in body :%v", string(res))
	}
	return blockNumber, nil
}
