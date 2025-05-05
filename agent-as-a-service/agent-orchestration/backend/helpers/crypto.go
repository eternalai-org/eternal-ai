package helpers

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func ValidateMessageSignature(msg string, signatureHex string, signer string) error {
	msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgBytes := []byte(msg)
	msgHash := crypto.Keccak256Hash(
		msgBytes,
	)
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}
	if signature[crypto.RecoveryIDOffset] > 1 {
		signature[crypto.RecoveryIDOffset] -= 27
	}
	sigPublicKey, err := crypto.SigToPub(msgHash.Bytes(), signature)
	if err != nil {
		return err
	}
	pbkHex := crypto.PubkeyToAddress(*sigPublicKey)
	if !strings.EqualFold(pbkHex.Hex(), signer) {
		return errors.New("not valid signer")
	}
	return nil
}

// get v, r, s variables from eip712 signature for permit function
func ERC20PermitSignature(pk *ecdsa.PrivateKey, domainSeparator, owner, spender common.Address, value, nonce, deadline *big.Int) (byte, [32]byte, [32]byte, error) {
	permitHash := common.HexToHash("6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9").Bytes()
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	args := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: addressTy,
		},
		{
			Type: addressTy,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}
	//abi.encode(PERMIT_TYPEHASH, owner, spender, value, nonce, deadline)
	bytes, err := args.Pack(
		permitHash,
		owner,
		spender,
		value,
		nonce,
		deadline,
	)
	if err != nil {
		return 0, [32]byte{}, [32]byte{}, err
	}
	//this is eq to keccak256(abi.encode(PERMIT_TYPEHASH, owner, spender, value, nonce, deadline))
	typedHash := crypto.Keccak256(bytes)
	//this is eq to keccak256(abi.encodePacked('x19x01', DOMAIN_SEPARATOR, hashed_args))
	hash := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator[:],
		typedHash,
	)
	sig, err := crypto.Sign(hash, pk)
	if err != nil {
		return 0, [32]byte{}, [32]byte{}, err
	}
	var v byte
	var r [32]byte
	var s [32]byte
	v = sig[64] + 27
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	return v, r, s, nil
}

func HexToBytes32(hex string) [32]byte {
	bytes := common.HexToHash(hex)
	if len(bytes) != 32 {
		panic("wrong bytes")
	}
	return bytes
}
