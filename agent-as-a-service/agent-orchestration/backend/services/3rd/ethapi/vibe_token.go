package ethapi

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *Client) GetSignatureForDeployToken(prk string, vibeTokenFactoryAddress common.Address, networkID uint64, nonce [32]byte, name string, symbol string, creator common.Address, tickLower *big.Int, tickUpper *big.Int, deadline *big.Int) (string, error) {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	int24Ty, _ := abi.NewType("int24", "int24", nil)
	stringTy, _ := abi.NewType("string", "string", nil)
	args := abi.Arguments{
		{Type: addressTy},
		{Type: uint256Ty},
		{Type: bytes32Ty},
		{Type: stringTy},
		{Type: stringTy},
		{Type: addressTy},
		{Type: int24Ty},
		{Type: int24Ty},
		{Type: uint256Ty},
	}
	data, err := args.Pack(
		vibeTokenFactoryAddress, big.NewInt(int64(networkID)), nonce, name, symbol, creator, tickLower, tickUpper, deadline,
	)
	if err != nil {
		return "", err
	}
	dataHash := crypto.Keccak256Hash(
		data,
	)
	signature, err := c.SignWithEthereum(prk, dataHash.Bytes())
	if err != nil {
		return "", err
	}
	return signature, nil
}
