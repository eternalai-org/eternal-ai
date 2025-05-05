package eth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"solo/pkg"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"core/contracts/erc20"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	ethsecp "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
	ctx := context.Background()
	i := 0
	for {
		time.Sleep(2 * time.Second)
		if i > 100 {
			return errors.New("timeout")
		}
		i++
		_, isPending, err := client.TransactionByHash(context.Background(), tx)
		if err != nil {
			continue
		}

		receipt, err := client.TransactionReceipt(ctx, tx)
		if err != nil {
			if err == ethereum.NotFound {
				continue
			}
			continue
		}

		if !isPending && receipt.Status == 1 {
			break
		}

	}
	return nil
}

func WalletAddressFromCompressedPublicKey(publicKeyStr string) (string, error) {
	pubBytes, err := hex.DecodeString(publicKeyStr)
	if err != nil {
		return "", err
	}

	x, y := ethsecp.DecompressPubkey(pubBytes)

	pubkey := elliptic.Marshal(ethsecp.S256(), x, y)

	ecdsaPub, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		return "", err
	}
	ethAddress := crypto.PubkeyToAddress(*ecdsaPub).String()
	return ethAddress, nil
}

func GetAccountInfo(privKey string) (*ecdsa.PrivateKey, *common.Address, error) {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey := masterWallet.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	promptFeeAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return masterWallet, &promptFeeAddress, nil
}

func GenerateKeyFromSeedOld(seed string) (string, string, string, error) {
	seedHex := hex.EncodeToString([]byte(seed))
	master, ch := hd.ComputeMastersFromSeed([]byte(seedHex))
	path := "m/44'/1022'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", "", "", err
	}
	privateKey := secp256k1.GenPrivKeyFromSecret(priv)

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

func CreateBindTransactionOpts(ctx context.Context, client *ethclient.Client, privateKey string, gasLimit int64) (*bind.TransactOpts, error) {
	priv, address, err := GetAccountInfo(privateKey)
	if err != nil {
		return nil, err
	}

	nonce, err := client.NonceAt(ctx, *address, nil)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice = gasPrice
	_ = gasPrice

	return auth, nil
}

func ApproveERC20(ctx context.Context, client *ethclient.Client, privateKey string, contractAddress common.Address, erc20Address common.Address, gasLimit int64) error {
	auth, err := CreateBindTransactionOpts(ctx, client, privateKey, gasLimit)
	if err != nil {
		return err
	}

	maxBigInt := new(big.Int)
	maxBigInt.SetString("30000000000000000000000", 10)

	erc20Contract, err := erc20.NewErc20(erc20Address, client)
	if err != nil {
		return err
	}

	auth.GasLimit = pkg.LOCAL_CHAIN_GAS_LIMIT
	auth.GasPrice = big.NewInt(pkg.LOCAL_CHAIN_GAS_PRICE)
	tx, err := erc20Contract.Approve(auth, contractAddress, maxBigInt)
	if err != nil {
		return err
	}

	if err = WaitForTx(client, tx.Hash()); err != nil {
		return err
	}

	return err
}

func SignMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

func CheckTransactionReverted(ctx context.Context, client *ethclient.Client, txHash common.Hash) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		if err == ethereum.NotFound {
			return false, fmt.Errorf("transaction not found")
		}
		return false, fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	// Check the status of the transaction
	if receipt.Status == 0 {
		return true, nil // Transaction has reverted
	}

	return false, nil // Transaction was successful
}

func GetTokenIDFromTx(client *ethclient.Client, txHash common.Hash) (*big.Int, error) {
	transferEventSignature := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		if err == ethereum.NotFound {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	// Parse the logs for the Transfer event
	for _, log := range receipt.Logs {
		if log.Topics[0].Hex() == transferEventSignature.Hex() {
			// The tokenID is the third indexed parameter
			if len(log.Topics) < 4 {
				continue
			}

			tokenID := new(big.Int).SetBytes(log.Topics[3].Bytes())
			return tokenID, nil
		}
	}

	return nil, fmt.Errorf("no Transfer event found in transaction")
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

func SendToken(ctx context.Context, fromPrivKey string, toAddress string, amount *big.Int, client *ethclient.Client, customGaslimit int64) (*types.Transaction, error) {
	auth, err := CreateBindTransactionOpts(ctx, client, fromPrivKey, customGaslimit)
	if err != nil {
		return nil, err
	}

	cID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	_toAddress := common.HexToAddress(toAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   cID,
		Nonce:     auth.Nonce.Uint64(),
		Gas:       uint64(customGaslimit),
		GasFeeCap: big.NewInt(customGaslimit * 2),
		//GasTipCap: big.NewInt(customGaslimit),
		Value: amount,
		To:    &_toAddress})

	pKey, _, _ := GetAccountInfo(fromPrivKey)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(cID), pKey)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
