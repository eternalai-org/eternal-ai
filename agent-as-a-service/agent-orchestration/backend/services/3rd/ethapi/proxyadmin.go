package ethapi

import (
	"context"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/proxyadmin"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) DeployProxyAdmin(prkHex string) (string, string, error) {
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
	address, tx, _, err := proxyadmin.DeployProxyAdmin(auth, client)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

func (c *Client) ProxyAdminUpgrade(contractAddr string, prkHex string, proxyAddress common.Address, logic common.Address) (string, error) {
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
	proAdmin, err := proxyadmin.NewProxyAdmin(helpers.HexToAddress(contractAddr), client)
	if err != nil {
		return "", err
	}
	nowLogic, err := proAdmin.GetProxyImplementation(&bind.CallOpts{}, proxyAddress)
	if err != nil {
		return "", err
	}
	if strings.EqualFold(nowLogic.Hex(), logic.Hex()) {
		return "", nil
	}
	instanceABI, err := proxyadmin.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"upgrade",
		proxyAddress,
		logic,
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
