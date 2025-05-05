package ethapi

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/arbitrumnonfungiblepositionmanager"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// func (c *Client) ArbitrumNonfungiblePositionManagerMint(contractAddr string, privateHex string, weth9 common.Address, sqrtPriceX96 *big.Int, params []*basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams) (string, error) {
// 	addressHex, prk, err := c.parsePrkAuth(privateHex)
// 	if err != nil {
// 		return "", err
// 	}
// 	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
// 	if err != nil {
// 		return "", err
// 	}
// 	client, err := c.getClient()
// 	if err != nil {
// 		return "", err
// 	}
// 	value := big.NewInt(0)
// 	for _, param := range params {
// 		if !strings.EqualFold(weth9.Hex(), param.Token0.Hex()) {
// 			if param.Amount0Desired.Cmp(big.NewInt(0)) > 0 {
// 				allowance, err := c.Erc20Allowance(param.Token0.Hex(), addressHex.Hex(), contractAddr)
// 				if err != nil {
// 					return "", err
// 				}
// 				if allowance.Cmp(big.NewInt(0)) <= 0 {
// 					approveHash, err := c.Erc20ApproveMax(
// 						param.Token0.Hex(),
// 						privateHex,
// 						contractAddr,
// 					)
// 					if err != nil {
// 						return "", err
// 					}
// 					time.Sleep(5 * time.Second)
// 					err = c.WaitMined(approveHash)
// 					if err != nil {
// 						return "", err
// 					}
// 				}
// 			}
// 		} else {
// 			value = param.Amount0Desired
// 		}
// 		if !strings.EqualFold(weth9.Hex(), param.Token1.Hex()) {
// 			if param.Amount1Desired.Cmp(big.NewInt(0)) > 0 {
// 				allowance, err := c.Erc20Allowance(param.Token1.Hex(), addressHex.Hex(), contractAddr)
// 				if err != nil {
// 					return "", err
// 				}
// 				if allowance.Cmp(big.NewInt(0)) <= 0 {
// 					approveHash, err := c.Erc20ApproveMax(
// 						param.Token1.Hex(),
// 						privateHex,
// 						contractAddr,
// 					)
// 					if err != nil {
// 						return "", err
// 					}
// 					time.Sleep(5 * time.Second)
// 					err = c.WaitMined(approveHash)
// 					if err != nil {
// 						return "", err
// 					}
// 				}
// 			}
// 		} else {
// 			value = param.Amount1Desired
// 		}
// 	}
// 	contractAddress := helpers.HexToAddress(contractAddr)
// 	// EstimateGas
// 	instanceABI, err := abi.JSON(strings.NewReader(basenonfungiblepositionmanager.NonfungiblePositionManagerABI))
// 	if err != nil {
// 		return "", err
// 	}
// 	multicallBytes := [][]byte{}
// 	{
// 		for _, param := range params {
// 			multicallData, err := instanceABI.Pack(
// 				"createAndInitializePoolIfNecessary",
// 				param.Token0,
// 				param.Token1,
// 				param.Fee,
// 				sqrtPriceX96,
// 			)
// 			if err != nil {
// 				return "", err
// 			}
// 			multicallBytes = append(
// 				multicallBytes,
// 				multicallData,
// 			)
// 		}
// 	}
// 	{
// 		for _, param := range params {
// 			multicallData, err := instanceABI.Pack(
// 				"mint",
// 				basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
// 					Token0:         param.Token0,
// 					Token1:         param.Token1,
// 					Fee:            param.Fee,
// 					TickLower:      param.TickLower,
// 					TickUpper:      param.TickUpper,
// 					Amount0Desired: param.Amount0Desired,
// 					Amount1Desired: param.Amount1Desired,
// 					Amount0Min:     param.Amount0Min,
// 					Amount1Min:     param.Amount1Min,
// 					Recipient:      addressHex,
// 					Deadline:       param.Deadline,
// 				},
// 			)
// 			if err != nil {
// 				return "", err
// 			}
// 			multicallBytes = append(
// 				multicallBytes,
// 				multicallData,
// 			)
// 		}
// 	}
// 	for _, param := range params {
// 		if strings.EqualFold(weth9.Hex(), param.Token0.Hex()) || strings.EqualFold(weth9.Hex(), param.Token1.Hex()) {
// 			multicallData, err := instanceABI.Pack(
// 				"refundETH",
// 			)
// 			if err != nil {
// 				return "", err
// 			}
// 			multicallBytes = append(
// 				multicallBytes,
// 				multicallData,
// 			)
// 			break
// 		}
// 	}
// 	dataBytes, err := instanceABI.Pack(
// 		"multicall",
// 		multicallBytes,
// 	)
// 	if err != nil {
// 		return "", err
// 	}
// 	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 		From:  addressHex,
// 		To:    &contractAddress,
// 		Data:  dataBytes,
// 		Value: value,
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	chainID, err := c.GetChainID()
// 	if err != nil {
// 		return "", err
// 	}
// 	nonceAt, err := c.PendingNonceAt(context.Background(), addressHex)
// 	if err != nil {
// 		return "", err
// 	}
// 	rawTx := types.NewTx(&types.DynamicFeeTx{
// 		ChainID:   big.NewInt(int64(chainID)),
// 		Nonce:     nonceAt,
// 		GasFeeCap: gasPrice,
// 		GasTipCap: gasTipCap,
// 		Gas:       (gasNumber * 12 / 10),
// 		To:        &contractAddress,
// 		Value:     value,
// 		Data:      dataBytes,
// 	})
// 	signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(int64(chainID))), prk)
// 	if err != nil {
// 		return "", err
// 	}
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		return "", err
// 	}
// 	return signedTx.Hash().Hex(), nil
// }

func (c *Client) ArbitrumNonfungiblePositionManagerMint(contractAddr string, privateHex string, weth9 common.Address, sqrtPriceX96 *big.Int, params *arbitrumnonfungiblepositionmanager.INonfungiblePositionManagerMintParams) (string, error) {
	addressHex, prk, err := c.parsePrkAuth(privateHex)
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
	value := big.NewInt(0)
	if !strings.EqualFold(weth9.Hex(), params.Token0.Hex()) {
		if params.Amount0Desired.Cmp(big.NewInt(0)) > 0 {
			allowance, err := c.Erc20Allowance(params.Token0.Hex(), addressHex.Hex(), contractAddr)
			if err != nil {
				return "", err
			}
			if allowance.Cmp(big.NewInt(0)) <= 0 {
				approveHash, err := c.Erc20ApproveMax(
					params.Token0.Hex(),
					privateHex,
					contractAddr,
				)
				if err != nil {
					return "", err
				}
				time.Sleep(10 * time.Second)
				err = c.WaitMined(approveHash)
				if err != nil {
					return "", err
				}
			}
		}
	} else {
		value = params.Amount0Desired
	}
	if !strings.EqualFold(weth9.Hex(), params.Token1.Hex()) {
		if params.Amount1Desired.Cmp(big.NewInt(0)) > 0 {
			allowance, err := c.Erc20Allowance(params.Token1.Hex(), addressHex.Hex(), contractAddr)
			if err != nil {
				return "", err
			}
			if allowance.Cmp(big.NewInt(0)) <= 0 {
				approveHash, err := c.Erc20ApproveMax(
					params.Token1.Hex(),
					privateHex,
					contractAddr,
				)
				if err != nil {
					return "", err
				}
				time.Sleep(10 * time.Second)
				err = c.WaitMined(approveHash)
				if err != nil {
					return "", err
				}
			}
		}
	} else {
		value = params.Amount1Desired
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	// EstimateGas
	instanceABI, err := abi.JSON(strings.NewReader(arbitrumnonfungiblepositionmanager.NonfungiblePositionManagerABI))
	if err != nil {
		return "", err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"createAndInitializePoolIfNecessary",
			params.Token0,
			params.Token1,
			sqrtPriceX96,
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"mint",
			arbitrumnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
				Token0:         params.Token0,
				Token1:         params.Token1,
				TickLower:      params.TickLower,
				TickUpper:      params.TickUpper,
				Amount0Desired: params.Amount0Desired,
				Amount1Desired: params.Amount1Desired,
				Amount0Min:     params.Amount0Min,
				Amount1Min:     params.Amount1Min,
				Recipient:      addressHex,
				Deadline:       params.Deadline,
			},
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	if strings.EqualFold(weth9.Hex(), params.Token0.Hex()) || strings.EqualFold(weth9.Hex(), params.Token1.Hex()) {
		multicallData, err := instanceABI.Pack(
			"refundNativeToken",
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	dataBytes, err := instanceABI.Pack(
		"multicall",
		multicallBytes,
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  addressHex,
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
	nonceAt, err := c.PendingNonceAt(context.Background(), addressHex)
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
	return signedTx.Hash().Hex(), nil
}
