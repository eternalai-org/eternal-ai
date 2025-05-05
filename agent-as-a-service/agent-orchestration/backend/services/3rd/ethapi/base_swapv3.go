package ethapi

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/basenonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/basequoterv2"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/baseswaprouter02"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// func (c *Client) BaseNonfungiblePositionManagerMint(contractAddr string, privateHex string, weth9 common.Address, sqrtPriceX96 *big.Int, params []*basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams) (string, error) {
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

func (c *Client) BaseNonfungiblePositionManagerMint(contractAddr string, privateHex string, weth9 common.Address, sqrtPriceX96 *big.Int, params *basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams) (string, error) {
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
	instanceABI, err := abi.JSON(strings.NewReader(basenonfungiblepositionmanager.NonfungiblePositionManagerABI))
	if err != nil {
		return "", err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"createAndInitializePoolIfNecessary",
			params.Token0,
			params.Token1,
			params.Fee,
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
			basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
				Token0:         params.Token0,
				Token1:         params.Token1,
				Fee:            params.Fee,
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
			"refundETH",
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

func (c *Client) BaseNonfungiblePositionManagerMintData(recipient string, sqrtPriceX96 *big.Int, params *basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(basenonfungiblepositionmanager.NonfungiblePositionManagerABI))
	if err != nil {
		return nil, err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"createAndInitializePoolIfNecessary",
			params.Token0,
			params.Token1,
			params.Fee,
			sqrtPriceX96,
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"mint",
			basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
				Token0:         params.Token0,
				Token1:         params.Token1,
				Fee:            params.Fee,
				TickLower:      params.TickLower,
				TickUpper:      params.TickUpper,
				Amount0Desired: params.Amount0Desired,
				Amount1Desired: params.Amount1Desired,
				Amount0Min:     params.Amount0Min,
				Amount1Min:     params.Amount1Min,
				Recipient:      helpers.HexToAddress(recipient),
				Deadline:       params.Deadline,
			},
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"refundETH",
		)
		if err != nil {
			return nil, err
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
		return nil, err
	}
	return dataBytes, nil
}

func (c *Client) BaseSwapRouterExactInputSingle(contractAddr, privateHex string, weth9 common.Address, params *baseswaprouter02.IV3SwapRouterExactInputSingleParams) (string, error) {
	addressHex, prk, err := c.parsePrkAuth(privateHex)
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	if params.AmountIn.Cmp(big.NewInt(0)) <= 0 {
		return "", errors.New("amountIn is not enough for tx")
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	instanceABI, err := abi.JSON(strings.NewReader(baseswaprouter02.SwapRouter02ABI))
	if err != nil {
		return "", err
	}
	var dataBytes []byte
	value := big.NewInt(0)
	if !strings.EqualFold(params.TokenIn.Hex(), weth9.Hex()) {
		allowance, err := c.Erc20Allowance(params.TokenIn.Hex(), addressHex.Hex(), contractAddr)
		if err != nil {
			return "", err
		}
		if allowance.Cmp(params.AmountIn) < 0 {
			approveHash, err := c.Erc20ApproveMax(
				params.TokenIn.Hex(),
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
	if strings.EqualFold(params.TokenIn.Hex(), weth9.Hex()) {
		value = params.AmountIn
		exactInputSingleDataBytes, err := instanceABI.Pack(
			"exactInputSingle",
			baseswaprouter02.IV3SwapRouterExactInputSingleParams{
				TokenIn:           params.TokenIn,
				TokenOut:          params.TokenOut,
				Fee:               params.Fee,
				Recipient:         params.Recipient,
				AmountIn:          params.AmountIn,
				AmountOutMinimum:  params.AmountOutMinimum,
				SqrtPriceLimitX96: params.SqrtPriceLimitX96,
			},
		)
		if err != nil {
			return "", err
		}
		refundWETHDataBytes, err := instanceABI.Pack(
			"refundETH",
		)
		if err != nil {
			return "", err
		}
		dataBytes, err = instanceABI.Pack(
			"multicall0",
			big.NewInt(time.Now().Unix()+60),
			[][]byte{
				exactInputSingleDataBytes,
				refundWETHDataBytes,
			},
		)
		if err != nil {
			return "", err
		}
	} else if strings.EqualFold(params.TokenOut.Hex(), weth9.Hex()) {
		exactInputSingleDataBytes, err := instanceABI.Pack(
			"exactInputSingle",
			baseswaprouter02.IV3SwapRouterExactInputSingleParams{
				TokenIn:           params.TokenIn,
				TokenOut:          params.TokenOut,
				Fee:               params.Fee,
				Recipient:         contractAddress,
				AmountIn:          params.AmountIn,
				AmountOutMinimum:  params.AmountOutMinimum,
				SqrtPriceLimitX96: params.SqrtPriceLimitX96,
			},
		)
		if err != nil {
			return "", err
		}
		unwrapWETHDataBytes, err := instanceABI.Pack(
			"unwrapWETH9",
			big.NewInt(0),
			addressHex,
		)
		if err != nil {
			return "", err
		}
		sweepTokenDataBytes, err := instanceABI.Pack(
			"sweepToken",
			params.TokenOut,
			big.NewInt(0),
			addressHex,
		)
		if err != nil {
			return "", err
		}
		dataBytes, err = instanceABI.Pack(
			"multicall0",
			big.NewInt(time.Now().Unix()+60),
			[][]byte{
				exactInputSingleDataBytes,
				unwrapWETHDataBytes,
				sweepTokenDataBytes,
			},
		)
		if err != nil {
			return "", err
		}
	} else {
		exactInputSingleDataBytes, err := instanceABI.Pack(
			"exactInputSingle",
			baseswaprouter02.IV3SwapRouterExactInputSingleParams{
				TokenIn:           params.TokenIn,
				TokenOut:          params.TokenOut,
				Fee:               params.Fee,
				Recipient:         addressHex,
				AmountIn:          params.AmountIn,
				AmountOutMinimum:  params.AmountOutMinimum,
				SqrtPriceLimitX96: params.SqrtPriceLimitX96,
			},
		)
		if err != nil {
			return "", err
		}
		dataBytes, err = instanceABI.Pack(
			"multicall0",
			big.NewInt(time.Now().Unix()+60),
			[][]byte{
				exactInputSingleDataBytes,
			},
		)
		if err != nil {
			return "", err
		}
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

func (c *Client) BaseQuoterQuoteExactInputSingle(contractAddress string, params *basequoterv2.IQuoterV2QuoteExactInputSingleParams) (*big.Int, *big.Int, uint32, uint32, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, nil, 0, 0, err
	}
	instance, err := basequoterv2.NewQuoterV2(helpers.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, nil, 0, 0, err
	}
	amountOut, sqrtPriceX96After, initializedTicksCrossed, gasEstimate, err := instance.
		QuoteExactInputSingleCall(
			&bind.CallOpts{},
			*params,
		)
	if err != nil {
		return nil, nil, 0, 0, err
	}
	return amountOut, sqrtPriceX96After, initializedTicksCrossed, gasEstimate, nil
}

func (c *Client) BaseQuoterQuoteExactOutputSingle(contractAddress string, params *basequoterv2.IQuoterV2QuoteExactOutputSingleParams) (*big.Int, *big.Int, uint32, uint32, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, nil, 0, 0, err
	}
	instance, err := basequoterv2.NewQuoterV2(helpers.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, nil, 0, 0, err
	}
	amountIn, sqrtPriceX96After, initializedTicksCrossed, gasEstimate, err := instance.
		QuoteExactOutputSingleCall(
			&bind.CallOpts{},
			*params,
		)
	if err != nil {
		return nil, nil, 0, 0, err
	}
	return amountIn, sqrtPriceX96After, initializedTicksCrossed, gasEstimate, nil
}
