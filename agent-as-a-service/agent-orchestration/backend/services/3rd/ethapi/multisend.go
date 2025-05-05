package ethapi

import (
	"context"
	"math/big"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/multisend"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) DeployMultisend(prkHex string) (string, string, error) {
	_, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(prk, big.NewInt(int64(chainID)))
	if err != nil {
		return "", "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", "", err
	}
	auth.GasPrice = gasPrice
	address, tx, _, err := multisend.DeployMultiSend(auth, client)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

func (c *Client) MultisendERC20Transfer(contractAddr string, prkHex string, token common.Address, addresses []common.Address, amounts []*big.Int) (string, error) {
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
	instanceABI, err := multisend.MultiSendMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"multiERC20Transfer",
		token,
		addresses,
		amounts,
	)
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	value := big.NewInt(0)
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  pbkHex,
		To:    &contractAddress,
		Data:  dataBytes,
		Value: value,
	})
	if err != nil {
		return "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	nonceAt, err := c.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonceAt,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Gas:       (gasNumber * 2),
		To:        &contractAddress,
		Value:     value,
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
	_, err = c.InscribeTxs([]string{signedTx.Hash().Hex()})
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
