// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agentfactory

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

// IAgentCodePointer is an auto generated low-level Go binding around an user-defined struct.
type IAgentCodePointer struct {
	RetrieveAddress common.Address
	FileType        uint8
	FileName        string
}

// AgentFactoryMetaData contains all meta data concerning the AgentFactory contract.
var AgentFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"agentId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"AgentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"agents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"agentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"agentVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"codeLanguage\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"retrieveAddress\",\"type\":\"address\"},{\"internalType\":\"enumIAgent.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"internalType\":\"structIAgent.CodePointer[]\",\"name\":\"pointers\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"depsAgents\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"agentOwner\",\"type\":\"address\"}],\"name\":\"createAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AgentFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentFactoryMetaData.ABI instead.
var AgentFactoryABI = AgentFactoryMetaData.ABI

// AgentFactory is an auto generated Go binding around an Ethereum contract.
type AgentFactory struct {
	AgentFactoryCaller     // Read-only binding to the contract
	AgentFactoryTransactor // Write-only binding to the contract
	AgentFactoryFilterer   // Log filterer for contract events
}

// AgentFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentFactorySession struct {
	Contract     *AgentFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentFactoryCallerSession struct {
	Contract *AgentFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AgentFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentFactoryTransactorSession struct {
	Contract     *AgentFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AgentFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentFactoryRaw struct {
	Contract *AgentFactory // Generic contract binding to access the raw methods on
}

// AgentFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentFactoryCallerRaw struct {
	Contract *AgentFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AgentFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentFactoryTransactorRaw struct {
	Contract *AgentFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentFactory creates a new instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactory(address common.Address, backend bind.ContractBackend) (*AgentFactory, error) {
	contract, err := bindAgentFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentFactory{AgentFactoryCaller: AgentFactoryCaller{contract: contract}, AgentFactoryTransactor: AgentFactoryTransactor{contract: contract}, AgentFactoryFilterer: AgentFactoryFilterer{contract: contract}}, nil
}

// NewAgentFactoryCaller creates a new read-only instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactoryCaller(address common.Address, caller bind.ContractCaller) (*AgentFactoryCaller, error) {
	contract, err := bindAgentFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryCaller{contract: contract}, nil
}

// NewAgentFactoryTransactor creates a new write-only instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentFactoryTransactor, error) {
	contract, err := bindAgentFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryTransactor{contract: contract}, nil
}

// NewAgentFactoryFilterer creates a new log filterer instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentFactoryFilterer, error) {
	contract, err := bindAgentFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryFilterer{contract: contract}, nil
}

// bindAgentFactory binds a generic wrapper to an already deployed contract.
func bindAgentFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentFactory *AgentFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentFactory.Contract.AgentFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentFactory *AgentFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentFactory.Contract.AgentFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentFactory *AgentFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentFactory.Contract.AgentFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentFactory *AgentFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentFactory *AgentFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentFactory *AgentFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentFactory.Contract.contract.Transact(opts, method, params...)
}

// Agents is a free data retrieval call binding the contract method 0x2e20ee18.
//
// Solidity: function agents(bytes32 ) view returns(address)
func (_AgentFactory *AgentFactoryCaller) Agents(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AgentFactory.contract.Call(opts, &out, "agents", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Agents is a free data retrieval call binding the contract method 0x2e20ee18.
//
// Solidity: function agents(bytes32 ) view returns(address)
func (_AgentFactory *AgentFactorySession) Agents(arg0 [32]byte) (common.Address, error) {
	return _AgentFactory.Contract.Agents(&_AgentFactory.CallOpts, arg0)
}

// Agents is a free data retrieval call binding the contract method 0x2e20ee18.
//
// Solidity: function agents(bytes32 ) view returns(address)
func (_AgentFactory *AgentFactoryCallerSession) Agents(arg0 [32]byte) (common.Address, error) {
	return _AgentFactory.Contract.Agents(&_AgentFactory.CallOpts, arg0)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AgentFactory *AgentFactoryCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentFactory.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AgentFactory *AgentFactorySession) GetImplementation() (common.Address, error) {
	return _AgentFactory.Contract.GetImplementation(&_AgentFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AgentFactory *AgentFactoryCallerSession) GetImplementation() (common.Address, error) {
	return _AgentFactory.Contract.GetImplementation(&_AgentFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentFactory *AgentFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentFactory *AgentFactorySession) Owner() (common.Address, error) {
	return _AgentFactory.Contract.Owner(&_AgentFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentFactory *AgentFactoryCallerSession) Owner() (common.Address, error) {
	return _AgentFactory.Contract.Owner(&_AgentFactory.CallOpts)
}

// CreateAgent is a paid mutator transaction binding the contract method 0x0d1f52ff.
//
// Solidity: function createAgent(bytes32 agentId, string agentName, string agentVersion, string codeLanguage, (address,uint8,string)[] pointers, address[] depsAgents, address agentOwner) returns(address agent)
func (_AgentFactory *AgentFactoryTransactor) CreateAgent(opts *bind.TransactOpts, agentId [32]byte, agentName string, agentVersion string, codeLanguage string, pointers []IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "createAgent", agentId, agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner)
}

// CreateAgent is a paid mutator transaction binding the contract method 0x0d1f52ff.
//
// Solidity: function createAgent(bytes32 agentId, string agentName, string agentVersion, string codeLanguage, (address,uint8,string)[] pointers, address[] depsAgents, address agentOwner) returns(address agent)
func (_AgentFactory *AgentFactorySession) CreateAgent(agentId [32]byte, agentName string, agentVersion string, codeLanguage string, pointers []IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.CreateAgent(&_AgentFactory.TransactOpts, agentId, agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner)
}

// CreateAgent is a paid mutator transaction binding the contract method 0x0d1f52ff.
//
// Solidity: function createAgent(bytes32 agentId, string agentName, string agentVersion, string codeLanguage, (address,uint8,string)[] pointers, address[] depsAgents, address agentOwner) returns(address agent)
func (_AgentFactory *AgentFactoryTransactorSession) CreateAgent(agentId [32]byte, agentName string, agentVersion string, codeLanguage string, pointers []IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.CreateAgent(&_AgentFactory.TransactOpts, agentId, agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner, address implementation, address registrar, address resolver) returns()
func (_AgentFactory *AgentFactoryTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, implementation common.Address, registrar common.Address, resolver common.Address) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "initialize", owner, implementation, registrar, resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner, address implementation, address registrar, address resolver) returns()
func (_AgentFactory *AgentFactorySession) Initialize(owner common.Address, implementation common.Address, registrar common.Address, resolver common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.Initialize(&_AgentFactory.TransactOpts, owner, implementation, registrar, resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner, address implementation, address registrar, address resolver) returns()
func (_AgentFactory *AgentFactoryTransactorSession) Initialize(owner common.Address, implementation common.Address, registrar common.Address, resolver common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.Initialize(&_AgentFactory.TransactOpts, owner, implementation, registrar, resolver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentFactory *AgentFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentFactory *AgentFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentFactory.Contract.RenounceOwnership(&_AgentFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentFactory *AgentFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentFactory.Contract.RenounceOwnership(&_AgentFactory.TransactOpts)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation) returns()
func (_AgentFactory *AgentFactoryTransactor) SetImplementation(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "setImplementation", implementation)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation) returns()
func (_AgentFactory *AgentFactorySession) SetImplementation(implementation common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.SetImplementation(&_AgentFactory.TransactOpts, implementation)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation) returns()
func (_AgentFactory *AgentFactoryTransactorSession) SetImplementation(implementation common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.SetImplementation(&_AgentFactory.TransactOpts, implementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentFactory *AgentFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentFactory *AgentFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.TransferOwnership(&_AgentFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentFactory *AgentFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.TransferOwnership(&_AgentFactory.TransactOpts, newOwner)
}

// AgentFactoryAgentCreatedIterator is returned from FilterAgentCreated and is used to iterate over the raw logs and unpacked data for AgentCreated events raised by the AgentFactory contract.
type AgentFactoryAgentCreatedIterator struct {
	Event *AgentFactoryAgentCreated // Event containing the contract specifics and raw log

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
func (it *AgentFactoryAgentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentFactoryAgentCreated)
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
		it.Event = new(AgentFactoryAgentCreated)
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
func (it *AgentFactoryAgentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentFactoryAgentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentFactoryAgentCreated represents a AgentCreated event raised by the AgentFactory contract.
type AgentFactoryAgentCreated struct {
	AgentId [32]byte
	Agent   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentCreated is a free log retrieval operation binding the contract event 0x7c96960a1ebd8cc753b10836ea25bd7c9c4f8cd43590db1e8b3648cb0ec4cc89.
//
// Solidity: event AgentCreated(bytes32 indexed agentId, address indexed agent)
func (_AgentFactory *AgentFactoryFilterer) FilterAgentCreated(opts *bind.FilterOpts, agentId [][32]byte, agent []common.Address) (*AgentFactoryAgentCreatedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentFactory.contract.FilterLogs(opts, "AgentCreated", agentIdRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryAgentCreatedIterator{contract: _AgentFactory.contract, event: "AgentCreated", logs: logs, sub: sub}, nil
}

// WatchAgentCreated is a free log subscription operation binding the contract event 0x7c96960a1ebd8cc753b10836ea25bd7c9c4f8cd43590db1e8b3648cb0ec4cc89.
//
// Solidity: event AgentCreated(bytes32 indexed agentId, address indexed agent)
func (_AgentFactory *AgentFactoryFilterer) WatchAgentCreated(opts *bind.WatchOpts, sink chan<- *AgentFactoryAgentCreated, agentId [][32]byte, agent []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentFactory.contract.WatchLogs(opts, "AgentCreated", agentIdRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentFactoryAgentCreated)
				if err := _AgentFactory.contract.UnpackLog(event, "AgentCreated", log); err != nil {
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

// ParseAgentCreated is a log parse operation binding the contract event 0x7c96960a1ebd8cc753b10836ea25bd7c9c4f8cd43590db1e8b3648cb0ec4cc89.
//
// Solidity: event AgentCreated(bytes32 indexed agentId, address indexed agent)
func (_AgentFactory *AgentFactoryFilterer) ParseAgentCreated(log types.Log) (*AgentFactoryAgentCreated, error) {
	event := new(AgentFactoryAgentCreated)
	if err := _AgentFactory.contract.UnpackLog(event, "AgentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentFactoryImplementationSetIterator is returned from FilterImplementationSet and is used to iterate over the raw logs and unpacked data for ImplementationSet events raised by the AgentFactory contract.
type AgentFactoryImplementationSetIterator struct {
	Event *AgentFactoryImplementationSet // Event containing the contract specifics and raw log

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
func (it *AgentFactoryImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentFactoryImplementationSet)
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
		it.Event = new(AgentFactoryImplementationSet)
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
func (it *AgentFactoryImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentFactoryImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentFactoryImplementationSet represents a ImplementationSet event raised by the AgentFactory contract.
type AgentFactoryImplementationSet struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterImplementationSet is a free log retrieval operation binding the contract event 0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13.
//
// Solidity: event ImplementationSet(address indexed implementation)
func (_AgentFactory *AgentFactoryFilterer) FilterImplementationSet(opts *bind.FilterOpts, implementation []common.Address) (*AgentFactoryImplementationSetIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AgentFactory.contract.FilterLogs(opts, "ImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryImplementationSetIterator{contract: _AgentFactory.contract, event: "ImplementationSet", logs: logs, sub: sub}, nil
}

// WatchImplementationSet is a free log subscription operation binding the contract event 0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13.
//
// Solidity: event ImplementationSet(address indexed implementation)
func (_AgentFactory *AgentFactoryFilterer) WatchImplementationSet(opts *bind.WatchOpts, sink chan<- *AgentFactoryImplementationSet, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AgentFactory.contract.WatchLogs(opts, "ImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentFactoryImplementationSet)
				if err := _AgentFactory.contract.UnpackLog(event, "ImplementationSet", log); err != nil {
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

// ParseImplementationSet is a log parse operation binding the contract event 0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13.
//
// Solidity: event ImplementationSet(address indexed implementation)
func (_AgentFactory *AgentFactoryFilterer) ParseImplementationSet(log types.Log) (*AgentFactoryImplementationSet, error) {
	event := new(AgentFactoryImplementationSet)
	if err := _AgentFactory.contract.UnpackLog(event, "ImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentFactory contract.
type AgentFactoryInitializedIterator struct {
	Event *AgentFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *AgentFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentFactoryInitialized)
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
		it.Event = new(AgentFactoryInitialized)
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
func (it *AgentFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentFactoryInitialized represents a Initialized event raised by the AgentFactory contract.
type AgentFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentFactory *AgentFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentFactoryInitializedIterator, error) {

	logs, sub, err := _AgentFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentFactoryInitializedIterator{contract: _AgentFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentFactory *AgentFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentFactoryInitialized)
				if err := _AgentFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AgentFactory *AgentFactoryFilterer) ParseInitialized(log types.Log) (*AgentFactoryInitialized, error) {
	event := new(AgentFactoryInitialized)
	if err := _AgentFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentFactory contract.
type AgentFactoryOwnershipTransferredIterator struct {
	Event *AgentFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AgentFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentFactoryOwnershipTransferred)
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
		it.Event = new(AgentFactoryOwnershipTransferred)
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
func (it *AgentFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the AgentFactory contract.
type AgentFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentFactory *AgentFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryOwnershipTransferredIterator{contract: _AgentFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentFactory *AgentFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentFactoryOwnershipTransferred)
				if err := _AgentFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AgentFactory *AgentFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*AgentFactoryOwnershipTransferred, error) {
	event := new(AgentFactoryOwnershipTransferred)
	if err := _AgentFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
