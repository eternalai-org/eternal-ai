// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hybrid_model

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// HybridModelMetaData contains all meta data concerning the HybridModel contract.
var HybridModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ModelIdAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"IdentifierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"NameUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"WorkerHubUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"identifier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_rawFlag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_workerHub\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelCollection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_identifier\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_modelId\",\"type\":\"uint256\"}],\"name\":\"setModelId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_identifier\",\"type\":\"uint256\"}],\"name\":\"updateIdentifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"}],\"name\":\"updateMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"updateName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_workerHub\",\"type\":\"address\"}],\"name\":\"updateWorkerHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workerHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// HybridModelABI is the input ABI used to generate the binding from.
// Deprecated: Use HybridModelMetaData.ABI instead.
var HybridModelABI = HybridModelMetaData.ABI

// HybridModel is an auto generated Go binding around an Ethereum contract.
type HybridModel struct {
	HybridModelCaller     // Read-only binding to the contract
	HybridModelTransactor // Write-only binding to the contract
	HybridModelFilterer   // Log filterer for contract events
}

// HybridModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type HybridModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HybridModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HybridModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HybridModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HybridModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HybridModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HybridModelSession struct {
	Contract     *HybridModel      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HybridModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HybridModelCallerSession struct {
	Contract *HybridModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HybridModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HybridModelTransactorSession struct {
	Contract     *HybridModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HybridModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type HybridModelRaw struct {
	Contract *HybridModel // Generic contract binding to access the raw methods on
}

// HybridModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HybridModelCallerRaw struct {
	Contract *HybridModelCaller // Generic read-only contract binding to access the raw methods on
}

// HybridModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HybridModelTransactorRaw struct {
	Contract *HybridModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHybridModel creates a new instance of HybridModel, bound to a specific deployed contract.
func NewHybridModel(address common.Address, backend bind.ContractBackend) (*HybridModel, error) {
	contract, err := bindHybridModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HybridModel{HybridModelCaller: HybridModelCaller{contract: contract}, HybridModelTransactor: HybridModelTransactor{contract: contract}, HybridModelFilterer: HybridModelFilterer{contract: contract}}, nil
}

// NewHybridModelCaller creates a new read-only instance of HybridModel, bound to a specific deployed contract.
func NewHybridModelCaller(address common.Address, caller bind.ContractCaller) (*HybridModelCaller, error) {
	contract, err := bindHybridModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HybridModelCaller{contract: contract}, nil
}

// NewHybridModelTransactor creates a new write-only instance of HybridModel, bound to a specific deployed contract.
func NewHybridModelTransactor(address common.Address, transactor bind.ContractTransactor) (*HybridModelTransactor, error) {
	contract, err := bindHybridModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HybridModelTransactor{contract: contract}, nil
}

// NewHybridModelFilterer creates a new log filterer instance of HybridModel, bound to a specific deployed contract.
func NewHybridModelFilterer(address common.Address, filterer bind.ContractFilterer) (*HybridModelFilterer, error) {
	contract, err := bindHybridModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HybridModelFilterer{contract: contract}, nil
}

// bindHybridModel binds a generic wrapper to an already deployed contract.
func bindHybridModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HybridModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HybridModel *HybridModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HybridModel.Contract.HybridModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HybridModel *HybridModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridModel.Contract.HybridModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HybridModel *HybridModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HybridModel.Contract.HybridModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HybridModel *HybridModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HybridModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HybridModel *HybridModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HybridModel *HybridModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HybridModel.Contract.contract.Transact(opts, method, params...)
}

// Identifier is a free data retrieval call binding the contract method 0x7998a1c4.
//
// Solidity: function identifier() view returns(uint256)
func (_HybridModel *HybridModelCaller) Identifier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "identifier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Identifier is a free data retrieval call binding the contract method 0x7998a1c4.
//
// Solidity: function identifier() view returns(uint256)
func (_HybridModel *HybridModelSession) Identifier() (*big.Int, error) {
	return _HybridModel.Contract.Identifier(&_HybridModel.CallOpts)
}

// Identifier is a free data retrieval call binding the contract method 0x7998a1c4.
//
// Solidity: function identifier() view returns(uint256)
func (_HybridModel *HybridModelCallerSession) Identifier() (*big.Int, error) {
	return _HybridModel.Contract.Identifier(&_HybridModel.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(string)
func (_HybridModel *HybridModelCaller) Metadata(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "metadata")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(string)
func (_HybridModel *HybridModelSession) Metadata() (string, error) {
	return _HybridModel.Contract.Metadata(&_HybridModel.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(string)
func (_HybridModel *HybridModelCallerSession) Metadata() (string, error) {
	return _HybridModel.Contract.Metadata(&_HybridModel.CallOpts)
}

// ModelCollection is a free data retrieval call binding the contract method 0x4d9f5d67.
//
// Solidity: function modelCollection() view returns(address)
func (_HybridModel *HybridModelCaller) ModelCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "modelCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelCollection is a free data retrieval call binding the contract method 0x4d9f5d67.
//
// Solidity: function modelCollection() view returns(address)
func (_HybridModel *HybridModelSession) ModelCollection() (common.Address, error) {
	return _HybridModel.Contract.ModelCollection(&_HybridModel.CallOpts)
}

// ModelCollection is a free data retrieval call binding the contract method 0x4d9f5d67.
//
// Solidity: function modelCollection() view returns(address)
func (_HybridModel *HybridModelCallerSession) ModelCollection() (common.Address, error) {
	return _HybridModel.Contract.ModelCollection(&_HybridModel.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HybridModel *HybridModelCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HybridModel *HybridModelSession) Name() (string, error) {
	return _HybridModel.Contract.Name(&_HybridModel.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HybridModel *HybridModelCallerSession) Name() (string, error) {
	return _HybridModel.Contract.Name(&_HybridModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HybridModel *HybridModelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HybridModel *HybridModelSession) Owner() (common.Address, error) {
	return _HybridModel.Contract.Owner(&_HybridModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HybridModel *HybridModelCallerSession) Owner() (common.Address, error) {
	return _HybridModel.Contract.Owner(&_HybridModel.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HybridModel *HybridModelCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HybridModel *HybridModelSession) Paused() (bool, error) {
	return _HybridModel.Contract.Paused(&_HybridModel.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HybridModel *HybridModelCallerSession) Paused() (bool, error) {
	return _HybridModel.Contract.Paused(&_HybridModel.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_HybridModel *HybridModelCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_HybridModel *HybridModelSession) Version() (string, error) {
	return _HybridModel.Contract.Version(&_HybridModel.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_HybridModel *HybridModelCallerSession) Version() (string, error) {
	return _HybridModel.Contract.Version(&_HybridModel.CallOpts)
}

// WorkerHub is a free data retrieval call binding the contract method 0x860e9dc6.
//
// Solidity: function workerHub() view returns(address)
func (_HybridModel *HybridModelCaller) WorkerHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridModel.contract.Call(opts, &out, "workerHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkerHub is a free data retrieval call binding the contract method 0x860e9dc6.
//
// Solidity: function workerHub() view returns(address)
func (_HybridModel *HybridModelSession) WorkerHub() (common.Address, error) {
	return _HybridModel.Contract.WorkerHub(&_HybridModel.CallOpts)
}

// WorkerHub is a free data retrieval call binding the contract method 0x860e9dc6.
//
// Solidity: function workerHub() view returns(address)
func (_HybridModel *HybridModelCallerSession) WorkerHub() (common.Address, error) {
	return _HybridModel.Contract.WorkerHub(&_HybridModel.CallOpts)
}

// Infer is a paid mutator transaction binding the contract method 0x130c280b.
//
// Solidity: function infer(bytes _input, bool _rawFlag) payable returns(uint256)
func (_HybridModel *HybridModelTransactor) Infer(opts *bind.TransactOpts, _input []byte, _rawFlag bool) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "infer", _input, _rawFlag)
}

// Infer is a paid mutator transaction binding the contract method 0x130c280b.
//
// Solidity: function infer(bytes _input, bool _rawFlag) payable returns(uint256)
func (_HybridModel *HybridModelSession) Infer(_input []byte, _rawFlag bool) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer(&_HybridModel.TransactOpts, _input, _rawFlag)
}

// Infer is a paid mutator transaction binding the contract method 0x130c280b.
//
// Solidity: function infer(bytes _input, bool _rawFlag) payable returns(uint256)
func (_HybridModel *HybridModelTransactorSession) Infer(_input []byte, _rawFlag bool) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer(&_HybridModel.TransactOpts, _input, _rawFlag)
}

// Infer0 is a paid mutator transaction binding the contract method 0x67e950a8.
//
// Solidity: function infer(bytes _input) payable returns(uint256)
func (_HybridModel *HybridModelTransactor) Infer0(opts *bind.TransactOpts, _input []byte) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "infer0", _input)
}

// Infer0 is a paid mutator transaction binding the contract method 0x67e950a8.
//
// Solidity: function infer(bytes _input) payable returns(uint256)
func (_HybridModel *HybridModelSession) Infer0(_input []byte) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer0(&_HybridModel.TransactOpts, _input)
}

// Infer0 is a paid mutator transaction binding the contract method 0x67e950a8.
//
// Solidity: function infer(bytes _input) payable returns(uint256)
func (_HybridModel *HybridModelTransactorSession) Infer0(_input []byte) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer0(&_HybridModel.TransactOpts, _input)
}

// Infer1 is a paid mutator transaction binding the contract method 0x7c22c0e3.
//
// Solidity: function infer(bytes _input, address _creator, bool _flag) payable returns(uint256)
func (_HybridModel *HybridModelTransactor) Infer1(opts *bind.TransactOpts, _input []byte, _creator common.Address, _flag bool) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "infer1", _input, _creator, _flag)
}

// Infer1 is a paid mutator transaction binding the contract method 0x7c22c0e3.
//
// Solidity: function infer(bytes _input, address _creator, bool _flag) payable returns(uint256)
func (_HybridModel *HybridModelSession) Infer1(_input []byte, _creator common.Address, _flag bool) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer1(&_HybridModel.TransactOpts, _input, _creator, _flag)
}

// Infer1 is a paid mutator transaction binding the contract method 0x7c22c0e3.
//
// Solidity: function infer(bytes _input, address _creator, bool _flag) payable returns(uint256)
func (_HybridModel *HybridModelTransactorSession) Infer1(_input []byte, _creator common.Address, _flag bool) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer1(&_HybridModel.TransactOpts, _input, _creator, _flag)
}

// Infer2 is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_HybridModel *HybridModelTransactor) Infer2(opts *bind.TransactOpts, _input []byte, _creator common.Address) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "infer2", _input, _creator)
}

// Infer2 is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_HybridModel *HybridModelSession) Infer2(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer2(&_HybridModel.TransactOpts, _input, _creator)
}

// Infer2 is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_HybridModel *HybridModelTransactorSession) Infer2(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _HybridModel.Contract.Infer2(&_HybridModel.TransactOpts, _input, _creator)
}

// Initialize is a paid mutator transaction binding the contract method 0x9ab8367e.
//
// Solidity: function initialize(address _workerHub, address _modelCollection, uint256 _identifier, string _name, string _metadata) returns()
func (_HybridModel *HybridModelTransactor) Initialize(opts *bind.TransactOpts, _workerHub common.Address, _modelCollection common.Address, _identifier *big.Int, _name string, _metadata string) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "initialize", _workerHub, _modelCollection, _identifier, _name, _metadata)
}

// Initialize is a paid mutator transaction binding the contract method 0x9ab8367e.
//
// Solidity: function initialize(address _workerHub, address _modelCollection, uint256 _identifier, string _name, string _metadata) returns()
func (_HybridModel *HybridModelSession) Initialize(_workerHub common.Address, _modelCollection common.Address, _identifier *big.Int, _name string, _metadata string) (*types.Transaction, error) {
	return _HybridModel.Contract.Initialize(&_HybridModel.TransactOpts, _workerHub, _modelCollection, _identifier, _name, _metadata)
}

// Initialize is a paid mutator transaction binding the contract method 0x9ab8367e.
//
// Solidity: function initialize(address _workerHub, address _modelCollection, uint256 _identifier, string _name, string _metadata) returns()
func (_HybridModel *HybridModelTransactorSession) Initialize(_workerHub common.Address, _modelCollection common.Address, _identifier *big.Int, _name string, _metadata string) (*types.Transaction, error) {
	return _HybridModel.Contract.Initialize(&_HybridModel.TransactOpts, _workerHub, _modelCollection, _identifier, _name, _metadata)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HybridModel *HybridModelTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HybridModel *HybridModelSession) Pause() (*types.Transaction, error) {
	return _HybridModel.Contract.Pause(&_HybridModel.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HybridModel *HybridModelTransactorSession) Pause() (*types.Transaction, error) {
	return _HybridModel.Contract.Pause(&_HybridModel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HybridModel *HybridModelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HybridModel *HybridModelSession) RenounceOwnership() (*types.Transaction, error) {
	return _HybridModel.Contract.RenounceOwnership(&_HybridModel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HybridModel *HybridModelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HybridModel.Contract.RenounceOwnership(&_HybridModel.TransactOpts)
}

// SetModelId is a paid mutator transaction binding the contract method 0x4d1c23c0.
//
// Solidity: function setModelId(uint256 _modelId) returns()
func (_HybridModel *HybridModelTransactor) SetModelId(opts *bind.TransactOpts, _modelId *big.Int) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "setModelId", _modelId)
}

// SetModelId is a paid mutator transaction binding the contract method 0x4d1c23c0.
//
// Solidity: function setModelId(uint256 _modelId) returns()
func (_HybridModel *HybridModelSession) SetModelId(_modelId *big.Int) (*types.Transaction, error) {
	return _HybridModel.Contract.SetModelId(&_HybridModel.TransactOpts, _modelId)
}

// SetModelId is a paid mutator transaction binding the contract method 0x4d1c23c0.
//
// Solidity: function setModelId(uint256 _modelId) returns()
func (_HybridModel *HybridModelTransactorSession) SetModelId(_modelId *big.Int) (*types.Transaction, error) {
	return _HybridModel.Contract.SetModelId(&_HybridModel.TransactOpts, _modelId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HybridModel *HybridModelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HybridModel *HybridModelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HybridModel.Contract.TransferOwnership(&_HybridModel.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HybridModel *HybridModelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HybridModel.Contract.TransferOwnership(&_HybridModel.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HybridModel *HybridModelTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HybridModel *HybridModelSession) Unpause() (*types.Transaction, error) {
	return _HybridModel.Contract.Unpause(&_HybridModel.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HybridModel *HybridModelTransactorSession) Unpause() (*types.Transaction, error) {
	return _HybridModel.Contract.Unpause(&_HybridModel.TransactOpts)
}

// UpdateIdentifier is a paid mutator transaction binding the contract method 0xe3284262.
//
// Solidity: function updateIdentifier(uint256 _identifier) returns()
func (_HybridModel *HybridModelTransactor) UpdateIdentifier(opts *bind.TransactOpts, _identifier *big.Int) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "updateIdentifier", _identifier)
}

// UpdateIdentifier is a paid mutator transaction binding the contract method 0xe3284262.
//
// Solidity: function updateIdentifier(uint256 _identifier) returns()
func (_HybridModel *HybridModelSession) UpdateIdentifier(_identifier *big.Int) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateIdentifier(&_HybridModel.TransactOpts, _identifier)
}

// UpdateIdentifier is a paid mutator transaction binding the contract method 0xe3284262.
//
// Solidity: function updateIdentifier(uint256 _identifier) returns()
func (_HybridModel *HybridModelTransactorSession) UpdateIdentifier(_identifier *big.Int) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateIdentifier(&_HybridModel.TransactOpts, _identifier)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x918b5be1.
//
// Solidity: function updateMetadata(string _metadata) returns()
func (_HybridModel *HybridModelTransactor) UpdateMetadata(opts *bind.TransactOpts, _metadata string) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "updateMetadata", _metadata)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x918b5be1.
//
// Solidity: function updateMetadata(string _metadata) returns()
func (_HybridModel *HybridModelSession) UpdateMetadata(_metadata string) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateMetadata(&_HybridModel.TransactOpts, _metadata)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x918b5be1.
//
// Solidity: function updateMetadata(string _metadata) returns()
func (_HybridModel *HybridModelTransactorSession) UpdateMetadata(_metadata string) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateMetadata(&_HybridModel.TransactOpts, _metadata)
}

// UpdateName is a paid mutator transaction binding the contract method 0x84da92a7.
//
// Solidity: function updateName(string _name) returns()
func (_HybridModel *HybridModelTransactor) UpdateName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "updateName", _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x84da92a7.
//
// Solidity: function updateName(string _name) returns()
func (_HybridModel *HybridModelSession) UpdateName(_name string) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateName(&_HybridModel.TransactOpts, _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x84da92a7.
//
// Solidity: function updateName(string _name) returns()
func (_HybridModel *HybridModelTransactorSession) UpdateName(_name string) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateName(&_HybridModel.TransactOpts, _name)
}

// UpdateWorkerHub is a paid mutator transaction binding the contract method 0x2b0a16f5.
//
// Solidity: function updateWorkerHub(address _workerHub) returns()
func (_HybridModel *HybridModelTransactor) UpdateWorkerHub(opts *bind.TransactOpts, _workerHub common.Address) (*types.Transaction, error) {
	return _HybridModel.contract.Transact(opts, "updateWorkerHub", _workerHub)
}

// UpdateWorkerHub is a paid mutator transaction binding the contract method 0x2b0a16f5.
//
// Solidity: function updateWorkerHub(address _workerHub) returns()
func (_HybridModel *HybridModelSession) UpdateWorkerHub(_workerHub common.Address) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateWorkerHub(&_HybridModel.TransactOpts, _workerHub)
}

// UpdateWorkerHub is a paid mutator transaction binding the contract method 0x2b0a16f5.
//
// Solidity: function updateWorkerHub(address _workerHub) returns()
func (_HybridModel *HybridModelTransactorSession) UpdateWorkerHub(_workerHub common.Address) (*types.Transaction, error) {
	return _HybridModel.Contract.UpdateWorkerHub(&_HybridModel.TransactOpts, _workerHub)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HybridModel *HybridModelTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridModel.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HybridModel *HybridModelSession) Receive() (*types.Transaction, error) {
	return _HybridModel.Contract.Receive(&_HybridModel.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HybridModel *HybridModelTransactorSession) Receive() (*types.Transaction, error) {
	return _HybridModel.Contract.Receive(&_HybridModel.TransactOpts)
}

// HybridModelIdentifierUpdateIterator is returned from FilterIdentifierUpdate and is used to iterate over the raw logs and unpacked data for IdentifierUpdate events raised by the HybridModel contract.
type HybridModelIdentifierUpdateIterator struct {
	Event *HybridModelIdentifierUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelIdentifierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelIdentifierUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelIdentifierUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelIdentifierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelIdentifierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelIdentifierUpdate represents a IdentifierUpdate event raised by the HybridModel contract.
type HybridModelIdentifierUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIdentifierUpdate is a free log retrieval operation binding the contract event 0x03e186105242a1073c60e85e4190610fc3cd7aa8685ef38b5fb00fb5ca834a46.
//
// Solidity: event IdentifierUpdate(uint256 newValue)
func (_HybridModel *HybridModelFilterer) FilterIdentifierUpdate(opts *bind.FilterOpts) (*HybridModelIdentifierUpdateIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "IdentifierUpdate")
	if err != nil {
		return nil, err
	}
	return &HybridModelIdentifierUpdateIterator{contract: _HybridModel.contract, event: "IdentifierUpdate", logs: logs, sub: sub}, nil
}

// WatchIdentifierUpdate is a free log subscription operation binding the contract event 0x03e186105242a1073c60e85e4190610fc3cd7aa8685ef38b5fb00fb5ca834a46.
//
// Solidity: event IdentifierUpdate(uint256 newValue)
func (_HybridModel *HybridModelFilterer) WatchIdentifierUpdate(opts *bind.WatchOpts, sink chan<- *HybridModelIdentifierUpdate) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "IdentifierUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelIdentifierUpdate)
				if err := _HybridModel.contract.UnpackLog(event, "IdentifierUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIdentifierUpdate is a log parse operation binding the contract event 0x03e186105242a1073c60e85e4190610fc3cd7aa8685ef38b5fb00fb5ca834a46.
//
// Solidity: event IdentifierUpdate(uint256 newValue)
func (_HybridModel *HybridModelFilterer) ParseIdentifierUpdate(log types.Log) (*HybridModelIdentifierUpdate, error) {
	event := new(HybridModelIdentifierUpdate)
	if err := _HybridModel.contract.UnpackLog(event, "IdentifierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HybridModel contract.
type HybridModelInitializedIterator struct {
	Event *HybridModelInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelInitialized represents a Initialized event raised by the HybridModel contract.
type HybridModelInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HybridModel *HybridModelFilterer) FilterInitialized(opts *bind.FilterOpts) (*HybridModelInitializedIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HybridModelInitializedIterator{contract: _HybridModel.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HybridModel *HybridModelFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HybridModelInitialized) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelInitialized)
				if err := _HybridModel.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HybridModel *HybridModelFilterer) ParseInitialized(log types.Log) (*HybridModelInitialized, error) {
	event := new(HybridModelInitialized)
	if err := _HybridModel.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the HybridModel contract.
type HybridModelMetadataUpdateIterator struct {
	Event *HybridModelMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelMetadataUpdate represents a MetadataUpdate event raised by the HybridModel contract.
type HybridModelMetadataUpdate struct {
	NewValue string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0x39b240244d76a12ce7613ed620720abc7a581774d12f26cda3b6b5d3ed535a31.
//
// Solidity: event MetadataUpdate(string newValue)
func (_HybridModel *HybridModelFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*HybridModelMetadataUpdateIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &HybridModelMetadataUpdateIterator{contract: _HybridModel.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0x39b240244d76a12ce7613ed620720abc7a581774d12f26cda3b6b5d3ed535a31.
//
// Solidity: event MetadataUpdate(string newValue)
func (_HybridModel *HybridModelFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *HybridModelMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelMetadataUpdate)
				if err := _HybridModel.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0x39b240244d76a12ce7613ed620720abc7a581774d12f26cda3b6b5d3ed535a31.
//
// Solidity: event MetadataUpdate(string newValue)
func (_HybridModel *HybridModelFilterer) ParseMetadataUpdate(log types.Log) (*HybridModelMetadataUpdate, error) {
	event := new(HybridModelMetadataUpdate)
	if err := _HybridModel.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelNameUpdateIterator is returned from FilterNameUpdate and is used to iterate over the raw logs and unpacked data for NameUpdate events raised by the HybridModel contract.
type HybridModelNameUpdateIterator struct {
	Event *HybridModelNameUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelNameUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelNameUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelNameUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelNameUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelNameUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelNameUpdate represents a NameUpdate event raised by the HybridModel contract.
type HybridModelNameUpdate struct {
	NewValue string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNameUpdate is a free log retrieval operation binding the contract event 0xba447370ea182903b32af1af1c0da2e966ba287288b29bf2b944259d335a9a8b.
//
// Solidity: event NameUpdate(string newValue)
func (_HybridModel *HybridModelFilterer) FilterNameUpdate(opts *bind.FilterOpts) (*HybridModelNameUpdateIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "NameUpdate")
	if err != nil {
		return nil, err
	}
	return &HybridModelNameUpdateIterator{contract: _HybridModel.contract, event: "NameUpdate", logs: logs, sub: sub}, nil
}

// WatchNameUpdate is a free log subscription operation binding the contract event 0xba447370ea182903b32af1af1c0da2e966ba287288b29bf2b944259d335a9a8b.
//
// Solidity: event NameUpdate(string newValue)
func (_HybridModel *HybridModelFilterer) WatchNameUpdate(opts *bind.WatchOpts, sink chan<- *HybridModelNameUpdate) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "NameUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelNameUpdate)
				if err := _HybridModel.contract.UnpackLog(event, "NameUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNameUpdate is a log parse operation binding the contract event 0xba447370ea182903b32af1af1c0da2e966ba287288b29bf2b944259d335a9a8b.
//
// Solidity: event NameUpdate(string newValue)
func (_HybridModel *HybridModelFilterer) ParseNameUpdate(log types.Log) (*HybridModelNameUpdate, error) {
	event := new(HybridModelNameUpdate)
	if err := _HybridModel.contract.UnpackLog(event, "NameUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HybridModel contract.
type HybridModelOwnershipTransferredIterator struct {
	Event *HybridModelOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelOwnershipTransferred represents a OwnershipTransferred event raised by the HybridModel contract.
type HybridModelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HybridModel *HybridModelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HybridModelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HybridModelOwnershipTransferredIterator{contract: _HybridModel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HybridModel *HybridModelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HybridModelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelOwnershipTransferred)
				if err := _HybridModel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HybridModel *HybridModelFilterer) ParseOwnershipTransferred(log types.Log) (*HybridModelOwnershipTransferred, error) {
	event := new(HybridModelOwnershipTransferred)
	if err := _HybridModel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HybridModel contract.
type HybridModelPausedIterator struct {
	Event *HybridModelPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelPaused represents a Paused event raised by the HybridModel contract.
type HybridModelPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HybridModel *HybridModelFilterer) FilterPaused(opts *bind.FilterOpts) (*HybridModelPausedIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HybridModelPausedIterator{contract: _HybridModel.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HybridModel *HybridModelFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HybridModelPaused) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelPaused)
				if err := _HybridModel.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HybridModel *HybridModelFilterer) ParsePaused(log types.Log) (*HybridModelPaused, error) {
	event := new(HybridModelPaused)
	if err := _HybridModel.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HybridModel contract.
type HybridModelUnpausedIterator struct {
	Event *HybridModelUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelUnpaused represents a Unpaused event raised by the HybridModel contract.
type HybridModelUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HybridModel *HybridModelFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HybridModelUnpausedIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HybridModelUnpausedIterator{contract: _HybridModel.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HybridModel *HybridModelFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HybridModelUnpaused) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelUnpaused)
				if err := _HybridModel.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HybridModel *HybridModelFilterer) ParseUnpaused(log types.Log) (*HybridModelUnpaused, error) {
	event := new(HybridModelUnpaused)
	if err := _HybridModel.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HybridModelWorkerHubUpdateIterator is returned from FilterWorkerHubUpdate and is used to iterate over the raw logs and unpacked data for WorkerHubUpdate events raised by the HybridModel contract.
type HybridModelWorkerHubUpdateIterator struct {
	Event *HybridModelWorkerHubUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HybridModelWorkerHubUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridModelWorkerHubUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HybridModelWorkerHubUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HybridModelWorkerHubUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HybridModelWorkerHubUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HybridModelWorkerHubUpdate represents a WorkerHubUpdate event raised by the HybridModel contract.
type HybridModelWorkerHubUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWorkerHubUpdate is a free log retrieval operation binding the contract event 0xc13b36de93e5efa00ccae6769264f17abfced5bee60a78d5735210c8c66b0181.
//
// Solidity: event WorkerHubUpdate(address newAddress)
func (_HybridModel *HybridModelFilterer) FilterWorkerHubUpdate(opts *bind.FilterOpts) (*HybridModelWorkerHubUpdateIterator, error) {

	logs, sub, err := _HybridModel.contract.FilterLogs(opts, "WorkerHubUpdate")
	if err != nil {
		return nil, err
	}
	return &HybridModelWorkerHubUpdateIterator{contract: _HybridModel.contract, event: "WorkerHubUpdate", logs: logs, sub: sub}, nil
}

// WatchWorkerHubUpdate is a free log subscription operation binding the contract event 0xc13b36de93e5efa00ccae6769264f17abfced5bee60a78d5735210c8c66b0181.
//
// Solidity: event WorkerHubUpdate(address newAddress)
func (_HybridModel *HybridModelFilterer) WatchWorkerHubUpdate(opts *bind.WatchOpts, sink chan<- *HybridModelWorkerHubUpdate) (event.Subscription, error) {

	logs, sub, err := _HybridModel.contract.WatchLogs(opts, "WorkerHubUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HybridModelWorkerHubUpdate)
				if err := _HybridModel.contract.UnpackLog(event, "WorkerHubUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWorkerHubUpdate is a log parse operation binding the contract event 0xc13b36de93e5efa00ccae6769264f17abfced5bee60a78d5735210c8c66b0181.
//
// Solidity: event WorkerHubUpdate(address newAddress)
func (_HybridModel *HybridModelFilterer) ParseWorkerHubUpdate(log types.Log) (*HybridModelWorkerHubUpdate, error) {
	event := new(HybridModelWorkerHubUpdate)
	if err := _HybridModel.contract.UnpackLog(event, "WorkerHubUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
