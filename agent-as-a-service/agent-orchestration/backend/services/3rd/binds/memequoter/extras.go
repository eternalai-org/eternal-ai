package memequoter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (_QuoterV2 *QuoterV2Transactor) QuoteExactInputSingleCall(opts *bind.CallOpts, params IQuoterV2QuoteExactInputSingleParams) (*big.Int, *big.Int, uint32, uint32, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "quoteExactInputSingle", params)
	if err != nil {
		return new(big.Int), new(big.Int), 0, 0, err
	}
	amountOut := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	sqrtPriceX96After := *abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	initializedTicksCrossed := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	gasEstimate := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	return &amountOut, &sqrtPriceX96After, initializedTicksCrossed, gasEstimate, err
}

func (_QuoterV2 *QuoterV2Transactor) QuoteExactOutputSingleCall(opts *bind.CallOpts, params IQuoterV2QuoteExactOutputSingleParams) (*big.Int, *big.Int, uint32, uint32, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "quoteExactOutputSingle", params)
	if err != nil {
		return new(big.Int), new(big.Int), 0, 0, err
	}
	amountIn := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	sqrtPriceX96After := *abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	initializedTicksCrossed := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	gasEstimate := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	return &amountIn, &sqrtPriceX96After, initializedTicksCrossed, gasEstimate, err
}
