package ethapi

import (
	"context"
	"math/big"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/dagent721"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) Dagent721Mint(contractAddr string, prkHex string, to common.Address, uri string, data []byte, fee *big.Int, promptKey string, promptScheduler common.Address, modelId uint32) (string, error) {
	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	instanceABI, err := dagent721.Dagent721MetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"mint", to, uri, data, fee, promptKey, promptScheduler, modelId,
	)
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	var gasNumber uint64
	if chainID == 964 {
		_, err = client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  pbkHex,
			To:    &contractAddress,
			Data:  dataBytes,
			Value: common.Big0,
		})
		if err != nil {
			return "", err
		}
		gasNumber = 1000000
	} else {
		gasNumber, err = client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  pbkHex,
			To:    &contractAddress,
			Data:  dataBytes,
			Value: common.Big0,
		})
		if err != nil {
			return "", err
		}
	}
	nonceAt, err := c.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(
		&types.LegacyTx{
			Nonce:    nonceAt,
			GasPrice: gasPrice,
			Gas:      (gasNumber * 2),
			To:       &contractAddress,
			Value:    common.Big0,
			Data:     dataBytes,
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
