// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agentupgradeable

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

// AgentUpgradeableMetaData contains all meta data concerning the AgentUpgradeable contract.
var AgentUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DigestAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthenticated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"retrieveAddress\",\"type\":\"address\"},{\"internalType\":\"enumIAgent.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIAgent.CodePointer\",\"name\":\"newPointer\",\"type\":\"tuple\"}],\"name\":\"CodePointerCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAddressByENS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"getAgentCode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAgentName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAgentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCodeLanguage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"getDepsAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"retrieveAddress\",\"type\":\"address\"},{\"internalType\":\"enumIAgent.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"internalType\":\"structIAgent.CodePointer[]\",\"name\":\"pointers\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"depsAgents\",\"type\":\"address[]\"}],\"name\":\"getHashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"agentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"agentVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"codeLanguage\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"retrieveAddress\",\"type\":\"address\"},{\"internalType\":\"enumIAgent.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"internalType\":\"structIAgent.CodePointer[]\",\"name\":\"pointers\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"depsAgents\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"agentOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nameService\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"retrieveAddress\",\"type\":\"address\"},{\"internalType\":\"enumIAgent.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"internalType\":\"structIAgent.CodePointer[]\",\"name\":\"pointers\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"depsAgents\",\"type\":\"address[]\"}],\"name\":\"publishAgentCode\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"retrieveAddress\",\"type\":\"address\"},{\"internalType\":\"enumIAgent.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"internalType\":\"structIAgent.CodePointer[]\",\"name\":\"pointers\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"depsAgents\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"publishAgentCodeWithSignature\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"internalType\":\"contractIBASERegistrarController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractIResolver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AgentUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentUpgradeableMetaData.ABI instead.
var AgentUpgradeableABI = AgentUpgradeableMetaData.ABI

// AgentUpgradeable is an auto generated Go binding around an Ethereum contract.
type AgentUpgradeable struct {
	AgentUpgradeableCaller     // Read-only binding to the contract
	AgentUpgradeableTransactor // Write-only binding to the contract
	AgentUpgradeableFilterer   // Log filterer for contract events
}

// AgentUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentUpgradeableSession struct {
	Contract     *AgentUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentUpgradeableCallerSession struct {
	Contract *AgentUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AgentUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentUpgradeableTransactorSession struct {
	Contract     *AgentUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AgentUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentUpgradeableRaw struct {
	Contract *AgentUpgradeable // Generic contract binding to access the raw methods on
}

// AgentUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentUpgradeableCallerRaw struct {
	Contract *AgentUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AgentUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentUpgradeableTransactorRaw struct {
	Contract *AgentUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentUpgradeable creates a new instance of AgentUpgradeable, bound to a specific deployed contract.
func NewAgentUpgradeable(address common.Address, backend bind.ContractBackend) (*AgentUpgradeable, error) {
	contract, err := bindAgentUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeable{AgentUpgradeableCaller: AgentUpgradeableCaller{contract: contract}, AgentUpgradeableTransactor: AgentUpgradeableTransactor{contract: contract}, AgentUpgradeableFilterer: AgentUpgradeableFilterer{contract: contract}}, nil
}

// NewAgentUpgradeableCaller creates a new read-only instance of AgentUpgradeable, bound to a specific deployed contract.
func NewAgentUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AgentUpgradeableCaller, error) {
	contract, err := bindAgentUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableCaller{contract: contract}, nil
}

// NewAgentUpgradeableTransactor creates a new write-only instance of AgentUpgradeable, bound to a specific deployed contract.
func NewAgentUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentUpgradeableTransactor, error) {
	contract, err := bindAgentUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableTransactor{contract: contract}, nil
}

// NewAgentUpgradeableFilterer creates a new log filterer instance of AgentUpgradeable, bound to a specific deployed contract.
func NewAgentUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentUpgradeableFilterer, error) {
	contract, err := bindAgentUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableFilterer{contract: contract}, nil
}

// bindAgentUpgradeable binds a generic wrapper to an already deployed contract.
func bindAgentUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentUpgradeable *AgentUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentUpgradeable.Contract.AgentUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentUpgradeable *AgentUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.AgentUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentUpgradeable *AgentUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.AgentUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentUpgradeable *AgentUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentUpgradeable *AgentUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentUpgradeable *AgentUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentUpgradeable *AgentUpgradeableCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentUpgradeable *AgentUpgradeableSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentUpgradeable.Contract.Eip712Domain(&_AgentUpgradeable.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentUpgradeable.Contract.Eip712Domain(&_AgentUpgradeable.CallOpts)
}

// GetAddressByENS is a free data retrieval call binding the contract method 0x306d41b7.
//
// Solidity: function getAddressByENS(string name) view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetAddressByENS(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getAddressByENS", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByENS is a free data retrieval call binding the contract method 0x306d41b7.
//
// Solidity: function getAddressByENS(string name) view returns(address)
func (_AgentUpgradeable *AgentUpgradeableSession) GetAddressByENS(name string) (common.Address, error) {
	return _AgentUpgradeable.Contract.GetAddressByENS(&_AgentUpgradeable.CallOpts, name)
}

// GetAddressByENS is a free data retrieval call binding the contract method 0x306d41b7.
//
// Solidity: function getAddressByENS(string name) view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetAddressByENS(name string) (common.Address, error) {
	return _AgentUpgradeable.Contract.GetAddressByENS(&_AgentUpgradeable.CallOpts, name)
}

// GetAgentCode is a free data retrieval call binding the contract method 0xc1f3dd3c.
//
// Solidity: function getAgentCode(uint16 version) view returns(string code)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetAgentCode(opts *bind.CallOpts, version uint16) (string, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getAgentCode", version)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAgentCode is a free data retrieval call binding the contract method 0xc1f3dd3c.
//
// Solidity: function getAgentCode(uint16 version) view returns(string code)
func (_AgentUpgradeable *AgentUpgradeableSession) GetAgentCode(version uint16) (string, error) {
	return _AgentUpgradeable.Contract.GetAgentCode(&_AgentUpgradeable.CallOpts, version)
}

// GetAgentCode is a free data retrieval call binding the contract method 0xc1f3dd3c.
//
// Solidity: function getAgentCode(uint16 version) view returns(string code)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetAgentCode(version uint16) (string, error) {
	return _AgentUpgradeable.Contract.GetAgentCode(&_AgentUpgradeable.CallOpts, version)
}

// GetAgentName is a free data retrieval call binding the contract method 0x27c3d2bb.
//
// Solidity: function getAgentName() view returns(string)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetAgentName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getAgentName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAgentName is a free data retrieval call binding the contract method 0x27c3d2bb.
//
// Solidity: function getAgentName() view returns(string)
func (_AgentUpgradeable *AgentUpgradeableSession) GetAgentName() (string, error) {
	return _AgentUpgradeable.Contract.GetAgentName(&_AgentUpgradeable.CallOpts)
}

// GetAgentName is a free data retrieval call binding the contract method 0x27c3d2bb.
//
// Solidity: function getAgentName() view returns(string)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetAgentName() (string, error) {
	return _AgentUpgradeable.Contract.GetAgentName(&_AgentUpgradeable.CallOpts)
}

// GetAgentOwner is a free data retrieval call binding the contract method 0x3cce96ec.
//
// Solidity: function getAgentOwner() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetAgentOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getAgentOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgentOwner is a free data retrieval call binding the contract method 0x3cce96ec.
//
// Solidity: function getAgentOwner() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableSession) GetAgentOwner() (common.Address, error) {
	return _AgentUpgradeable.Contract.GetAgentOwner(&_AgentUpgradeable.CallOpts)
}

// GetAgentOwner is a free data retrieval call binding the contract method 0x3cce96ec.
//
// Solidity: function getAgentOwner() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetAgentOwner() (common.Address, error) {
	return _AgentUpgradeable.Contract.GetAgentOwner(&_AgentUpgradeable.CallOpts)
}

// GetCodeLanguage is a free data retrieval call binding the contract method 0x6681792d.
//
// Solidity: function getCodeLanguage() view returns(string)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetCodeLanguage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getCodeLanguage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCodeLanguage is a free data retrieval call binding the contract method 0x6681792d.
//
// Solidity: function getCodeLanguage() view returns(string)
func (_AgentUpgradeable *AgentUpgradeableSession) GetCodeLanguage() (string, error) {
	return _AgentUpgradeable.Contract.GetCodeLanguage(&_AgentUpgradeable.CallOpts)
}

// GetCodeLanguage is a free data retrieval call binding the contract method 0x6681792d.
//
// Solidity: function getCodeLanguage() view returns(string)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetCodeLanguage() (string, error) {
	return _AgentUpgradeable.Contract.GetCodeLanguage(&_AgentUpgradeable.CallOpts)
}

// GetCurrentVersion is a free data retrieval call binding the contract method 0xfabec44a.
//
// Solidity: function getCurrentVersion() view returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetCurrentVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getCurrentVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetCurrentVersion is a free data retrieval call binding the contract method 0xfabec44a.
//
// Solidity: function getCurrentVersion() view returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableSession) GetCurrentVersion() (uint16, error) {
	return _AgentUpgradeable.Contract.GetCurrentVersion(&_AgentUpgradeable.CallOpts)
}

// GetCurrentVersion is a free data retrieval call binding the contract method 0xfabec44a.
//
// Solidity: function getCurrentVersion() view returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetCurrentVersion() (uint16, error) {
	return _AgentUpgradeable.Contract.GetCurrentVersion(&_AgentUpgradeable.CallOpts)
}

// GetDepsAgents is a free data retrieval call binding the contract method 0xfa0d4179.
//
// Solidity: function getDepsAgents(uint16 version) view returns(address[])
func (_AgentUpgradeable *AgentUpgradeableCaller) GetDepsAgents(opts *bind.CallOpts, version uint16) ([]common.Address, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getDepsAgents", version)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDepsAgents is a free data retrieval call binding the contract method 0xfa0d4179.
//
// Solidity: function getDepsAgents(uint16 version) view returns(address[])
func (_AgentUpgradeable *AgentUpgradeableSession) GetDepsAgents(version uint16) ([]common.Address, error) {
	return _AgentUpgradeable.Contract.GetDepsAgents(&_AgentUpgradeable.CallOpts, version)
}

// GetDepsAgents is a free data retrieval call binding the contract method 0xfa0d4179.
//
// Solidity: function getDepsAgents(uint16 version) view returns(address[])
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetDepsAgents(version uint16) ([]common.Address, error) {
	return _AgentUpgradeable.Contract.GetDepsAgents(&_AgentUpgradeable.CallOpts, version)
}

// GetHashToSign is a free data retrieval call binding the contract method 0x69d07dc2.
//
// Solidity: function getHashToSign((address,uint8,string)[] pointers, address[] depsAgents) view returns(bytes32)
func (_AgentUpgradeable *AgentUpgradeableCaller) GetHashToSign(opts *bind.CallOpts, pointers []IAgentCodePointer, depsAgents []common.Address) ([32]byte, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "getHashToSign", pointers, depsAgents)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashToSign is a free data retrieval call binding the contract method 0x69d07dc2.
//
// Solidity: function getHashToSign((address,uint8,string)[] pointers, address[] depsAgents) view returns(bytes32)
func (_AgentUpgradeable *AgentUpgradeableSession) GetHashToSign(pointers []IAgentCodePointer, depsAgents []common.Address) ([32]byte, error) {
	return _AgentUpgradeable.Contract.GetHashToSign(&_AgentUpgradeable.CallOpts, pointers, depsAgents)
}

// GetHashToSign is a free data retrieval call binding the contract method 0x69d07dc2.
//
// Solidity: function getHashToSign((address,uint8,string)[] pointers, address[] depsAgents) view returns(bytes32)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) GetHashToSign(pointers []IAgentCodePointer, depsAgents []common.Address) ([32]byte, error) {
	return _AgentUpgradeable.Contract.GetHashToSign(&_AgentUpgradeable.CallOpts, pointers, depsAgents)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableSession) Owner() (common.Address, error) {
	return _AgentUpgradeable.Contract.Owner(&_AgentUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) Owner() (common.Address, error) {
	return _AgentUpgradeable.Contract.Owner(&_AgentUpgradeable.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableSession) Registrar() (common.Address, error) {
	return _AgentUpgradeable.Contract.Registrar(&_AgentUpgradeable.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) Registrar() (common.Address, error) {
	return _AgentUpgradeable.Contract.Registrar(&_AgentUpgradeable.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentUpgradeable.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableSession) Resolver() (common.Address, error) {
	return _AgentUpgradeable.Contract.Resolver(&_AgentUpgradeable.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_AgentUpgradeable *AgentUpgradeableCallerSession) Resolver() (common.Address, error) {
	return _AgentUpgradeable.Contract.Resolver(&_AgentUpgradeable.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd4a6cbf.
//
// Solidity: function initialize(string agentName, string agentVersion, string codeLanguage, (address,uint8,string)[] pointers, address[] depsAgents, address agentOwner, bytes nameService) payable returns()
func (_AgentUpgradeable *AgentUpgradeableTransactor) Initialize(opts *bind.TransactOpts, agentName string, agentVersion string, codeLanguage string, pointers []IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address, nameService []byte) (*types.Transaction, error) {
	return _AgentUpgradeable.contract.Transact(opts, "initialize", agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner, nameService)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd4a6cbf.
//
// Solidity: function initialize(string agentName, string agentVersion, string codeLanguage, (address,uint8,string)[] pointers, address[] depsAgents, address agentOwner, bytes nameService) payable returns()
func (_AgentUpgradeable *AgentUpgradeableSession) Initialize(agentName string, agentVersion string, codeLanguage string, pointers []IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address, nameService []byte) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.Initialize(&_AgentUpgradeable.TransactOpts, agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner, nameService)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd4a6cbf.
//
// Solidity: function initialize(string agentName, string agentVersion, string codeLanguage, (address,uint8,string)[] pointers, address[] depsAgents, address agentOwner, bytes nameService) payable returns()
func (_AgentUpgradeable *AgentUpgradeableTransactorSession) Initialize(agentName string, agentVersion string, codeLanguage string, pointers []IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address, nameService []byte) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.Initialize(&_AgentUpgradeable.TransactOpts, agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner, nameService)
}

// PublishAgentCode is a paid mutator transaction binding the contract method 0x89ffd55c.
//
// Solidity: function publishAgentCode((address,uint8,string)[] pointers, address[] depsAgents) returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableTransactor) PublishAgentCode(opts *bind.TransactOpts, pointers []IAgentCodePointer, depsAgents []common.Address) (*types.Transaction, error) {
	return _AgentUpgradeable.contract.Transact(opts, "publishAgentCode", pointers, depsAgents)
}

// PublishAgentCode is a paid mutator transaction binding the contract method 0x89ffd55c.
//
// Solidity: function publishAgentCode((address,uint8,string)[] pointers, address[] depsAgents) returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableSession) PublishAgentCode(pointers []IAgentCodePointer, depsAgents []common.Address) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.PublishAgentCode(&_AgentUpgradeable.TransactOpts, pointers, depsAgents)
}

// PublishAgentCode is a paid mutator transaction binding the contract method 0x89ffd55c.
//
// Solidity: function publishAgentCode((address,uint8,string)[] pointers, address[] depsAgents) returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableTransactorSession) PublishAgentCode(pointers []IAgentCodePointer, depsAgents []common.Address) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.PublishAgentCode(&_AgentUpgradeable.TransactOpts, pointers, depsAgents)
}

// PublishAgentCodeWithSignature is a paid mutator transaction binding the contract method 0x6c6ca629.
//
// Solidity: function publishAgentCodeWithSignature((address,uint8,string)[] pointers, address[] depsAgents, bytes signature) returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableTransactor) PublishAgentCodeWithSignature(opts *bind.TransactOpts, pointers []IAgentCodePointer, depsAgents []common.Address, signature []byte) (*types.Transaction, error) {
	return _AgentUpgradeable.contract.Transact(opts, "publishAgentCodeWithSignature", pointers, depsAgents, signature)
}

// PublishAgentCodeWithSignature is a paid mutator transaction binding the contract method 0x6c6ca629.
//
// Solidity: function publishAgentCodeWithSignature((address,uint8,string)[] pointers, address[] depsAgents, bytes signature) returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableSession) PublishAgentCodeWithSignature(pointers []IAgentCodePointer, depsAgents []common.Address, signature []byte) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.PublishAgentCodeWithSignature(&_AgentUpgradeable.TransactOpts, pointers, depsAgents, signature)
}

// PublishAgentCodeWithSignature is a paid mutator transaction binding the contract method 0x6c6ca629.
//
// Solidity: function publishAgentCodeWithSignature((address,uint8,string)[] pointers, address[] depsAgents, bytes signature) returns(uint16)
func (_AgentUpgradeable *AgentUpgradeableTransactorSession) PublishAgentCodeWithSignature(pointers []IAgentCodePointer, depsAgents []common.Address, signature []byte) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.PublishAgentCodeWithSignature(&_AgentUpgradeable.TransactOpts, pointers, depsAgents, signature)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_AgentUpgradeable *AgentUpgradeableTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _AgentUpgradeable.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_AgentUpgradeable *AgentUpgradeableSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.Renew(&_AgentUpgradeable.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_AgentUpgradeable *AgentUpgradeableTransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.Renew(&_AgentUpgradeable.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentUpgradeable *AgentUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentUpgradeable *AgentUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.RenounceOwnership(&_AgentUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentUpgradeable *AgentUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.RenounceOwnership(&_AgentUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentUpgradeable *AgentUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentUpgradeable *AgentUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.TransferOwnership(&_AgentUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentUpgradeable *AgentUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentUpgradeable.Contract.TransferOwnership(&_AgentUpgradeable.TransactOpts, newOwner)
}

// AgentUpgradeableCodePointerCreatedIterator is returned from FilterCodePointerCreated and is used to iterate over the raw logs and unpacked data for CodePointerCreated events raised by the AgentUpgradeable contract.
type AgentUpgradeableCodePointerCreatedIterator struct {
	Event *AgentUpgradeableCodePointerCreated // Event containing the contract specifics and raw log

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
func (it *AgentUpgradeableCodePointerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentUpgradeableCodePointerCreated)
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
		it.Event = new(AgentUpgradeableCodePointerCreated)
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
func (it *AgentUpgradeableCodePointerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentUpgradeableCodePointerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentUpgradeableCodePointerCreated represents a CodePointerCreated event raised by the AgentUpgradeable contract.
type AgentUpgradeableCodePointerCreated struct {
	Version    *big.Int
	PIndex     *big.Int
	NewPointer IAgentCodePointer
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCodePointerCreated is a free log retrieval operation binding the contract event 0x0c69de805f518c322126e1fa81f2e91d814412a2ad304638fcaed200bd7f43dc.
//
// Solidity: event CodePointerCreated(uint256 indexed version, uint256 indexed pIndex, (address,uint8,string) newPointer)
func (_AgentUpgradeable *AgentUpgradeableFilterer) FilterCodePointerCreated(opts *bind.FilterOpts, version []*big.Int, pIndex []*big.Int) (*AgentUpgradeableCodePointerCreatedIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var pIndexRule []interface{}
	for _, pIndexItem := range pIndex {
		pIndexRule = append(pIndexRule, pIndexItem)
	}

	logs, sub, err := _AgentUpgradeable.contract.FilterLogs(opts, "CodePointerCreated", versionRule, pIndexRule)
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableCodePointerCreatedIterator{contract: _AgentUpgradeable.contract, event: "CodePointerCreated", logs: logs, sub: sub}, nil
}

// WatchCodePointerCreated is a free log subscription operation binding the contract event 0x0c69de805f518c322126e1fa81f2e91d814412a2ad304638fcaed200bd7f43dc.
//
// Solidity: event CodePointerCreated(uint256 indexed version, uint256 indexed pIndex, (address,uint8,string) newPointer)
func (_AgentUpgradeable *AgentUpgradeableFilterer) WatchCodePointerCreated(opts *bind.WatchOpts, sink chan<- *AgentUpgradeableCodePointerCreated, version []*big.Int, pIndex []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var pIndexRule []interface{}
	for _, pIndexItem := range pIndex {
		pIndexRule = append(pIndexRule, pIndexItem)
	}

	logs, sub, err := _AgentUpgradeable.contract.WatchLogs(opts, "CodePointerCreated", versionRule, pIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentUpgradeableCodePointerCreated)
				if err := _AgentUpgradeable.contract.UnpackLog(event, "CodePointerCreated", log); err != nil {
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

// ParseCodePointerCreated is a log parse operation binding the contract event 0x0c69de805f518c322126e1fa81f2e91d814412a2ad304638fcaed200bd7f43dc.
//
// Solidity: event CodePointerCreated(uint256 indexed version, uint256 indexed pIndex, (address,uint8,string) newPointer)
func (_AgentUpgradeable *AgentUpgradeableFilterer) ParseCodePointerCreated(log types.Log) (*AgentUpgradeableCodePointerCreated, error) {
	event := new(AgentUpgradeableCodePointerCreated)
	if err := _AgentUpgradeable.contract.UnpackLog(event, "CodePointerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentUpgradeableEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AgentUpgradeable contract.
type AgentUpgradeableEIP712DomainChangedIterator struct {
	Event *AgentUpgradeableEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AgentUpgradeableEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentUpgradeableEIP712DomainChanged)
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
		it.Event = new(AgentUpgradeableEIP712DomainChanged)
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
func (it *AgentUpgradeableEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentUpgradeableEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentUpgradeableEIP712DomainChanged represents a EIP712DomainChanged event raised by the AgentUpgradeable contract.
type AgentUpgradeableEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentUpgradeable *AgentUpgradeableFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AgentUpgradeableEIP712DomainChangedIterator, error) {

	logs, sub, err := _AgentUpgradeable.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableEIP712DomainChangedIterator{contract: _AgentUpgradeable.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentUpgradeable *AgentUpgradeableFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AgentUpgradeableEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AgentUpgradeable.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentUpgradeableEIP712DomainChanged)
				if err := _AgentUpgradeable.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentUpgradeable *AgentUpgradeableFilterer) ParseEIP712DomainChanged(log types.Log) (*AgentUpgradeableEIP712DomainChanged, error) {
	event := new(AgentUpgradeableEIP712DomainChanged)
	if err := _AgentUpgradeable.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentUpgradeable contract.
type AgentUpgradeableInitializedIterator struct {
	Event *AgentUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *AgentUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentUpgradeableInitialized)
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
		it.Event = new(AgentUpgradeableInitialized)
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
func (it *AgentUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentUpgradeableInitialized represents a Initialized event raised by the AgentUpgradeable contract.
type AgentUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentUpgradeable *AgentUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentUpgradeableInitializedIterator, error) {

	logs, sub, err := _AgentUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableInitializedIterator{contract: _AgentUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentUpgradeable *AgentUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentUpgradeableInitialized)
				if err := _AgentUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AgentUpgradeable *AgentUpgradeableFilterer) ParseInitialized(log types.Log) (*AgentUpgradeableInitialized, error) {
	event := new(AgentUpgradeableInitialized)
	if err := _AgentUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentUpgradeable contract.
type AgentUpgradeableOwnershipTransferredIterator struct {
	Event *AgentUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AgentUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentUpgradeableOwnershipTransferred)
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
		it.Event = new(AgentUpgradeableOwnershipTransferred)
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
func (it *AgentUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the AgentUpgradeable contract.
type AgentUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentUpgradeable *AgentUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentUpgradeableOwnershipTransferredIterator{contract: _AgentUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentUpgradeable *AgentUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentUpgradeableOwnershipTransferred)
				if err := _AgentUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AgentUpgradeable *AgentUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*AgentUpgradeableOwnershipTransferred, error) {
	event := new(AgentUpgradeableOwnershipTransferred)
	if err := _AgentUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
