package ethapi

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc721"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) Erc721Transfer(contractAddr string, prkHex string, toAddr string, tokenId *big.Int) (string, error) {
	if contractAddr == "" ||
		!common.IsHexAddress(contractAddr) {
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
	contractAddress := helpers.HexToAddress(contractAddr)
	instanceABI, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"safeTransferFrom",
		pbkHex,
		helpers.HexToAddress(toAddr),
		tokenId,
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
