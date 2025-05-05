package eth

import (
	"context"
	"core/contracts/proxy"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Deploy(client *ethclient.Client, privateKey string, contractABI string, contractByteCode string, args ...interface{}) (common.Address, *types.Transaction, error) {

	contractInst, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return common.Address{}, nil, err
	}

	bytecode := common.FromHex(contractByteCode)
	auth, err := CreateBindTransactionOpts(context.Background(), client, privateKey, 3000000)
	if err != nil {
		return common.Address{}, nil, err
	}

	// Send the transaction
	address, tx, _, err := bind.DeployContract(auth, contractInst, bytecode, client)
	if err != nil {
		return common.Address{}, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	_ = receipt
	return address, tx, err

}

func Initialize(client *ethclient.Client, privateKey string, contractABI string, contractByteCode string, args ...interface{}) (*common.Address, *types.Transaction, error) {

	contractInst, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, nil, err
	}

	bytecode := common.FromHex(contractByteCode)
	// Pack the constructor arguments
	data, err := contractInst.Pack("initialize", args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack constructor arguments: %v", err)
	}

	auth, err := CreateBindTransactionOpts(context.Background(), client, privateKey, 3000000)
	if err != nil {
		return nil, nil, err
	}

	bytecode = append(bytecode, data...)

	tx := types.NewTransaction(auth.Nonce.Uint64(), common.Address{}, big.NewInt(0), auth.GasLimit, auth.GasPrice, bytecode)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return nil, nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("signedTx===>", tx.Hash())
	fmt.Println("receipt===>", receipt)
	fmt.Println("ContractAddress===>", receipt.ContractAddress.Hex())

	return &receipt.ContractAddress, tx, nil

}

func DeployProxy(client *ethclient.Client, privateKey string, contractABI string, contractByteCode string, args ...interface{}) (common.Address, *types.Transaction, error) {

	implAddress, tx, err := Deploy(client, privateKey, contractABI, contractByteCode, args)
	if err != nil {
		fmt.Println("Deploy 1 ====> ", err)
		return common.Address{}, nil, err
	}

	fmt.Println("implAddress===> ", implAddress.Hex())
	address, tx, err := Deploy(client, privateKey, proxy.ProxyABI, proxy.ProxyBin, implAddress)
	if err != nil {
		fmt.Println("Deploy 2 ====> ", err)
		return common.Address{}, nil, err
	}

	_, tx, err = Initialize(client, privateKey, contractABI, contractByteCode, args...)
	if err != nil {
		fmt.Println("Deploy 3 ====> ", err)
		return common.Address{}, nil, err
	}

	fmt.Println("implAddress===> ", implAddress.Hex())
	fmt.Println("address===> ", address.Hex())
	return address, tx, nil
}
