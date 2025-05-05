// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package w_eai

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

// WEaiMetaData contains all meta data concerning the WEai contract.
var WEaiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EAIUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EAIWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60803462000380576040906001600160401b0390808301828111828210176200036a578352600b81526020916a577261707065642045414960a81b83830152835192848401848110838211176200036a578552600493848152635745414960e01b8282015283519083821162000355576003928354926001968785811c951680156200034a575b8386101462000335578190601f95868111620002df575b50839086831160011462000278576000926200026c575b505060001982871b1c191690871b1784555b8151948511620002575786548681811c911680156200024c575b828210146200023757838111620001ec575b50809285116001146200017e575093839491849260009562000172575b50501b92600019911b1c19161790555b60058054336001600160a01b0319821681179092559151916001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a36113de9081620003868239f35b0151935038806200010f565b92919084601f1981168860005285600020956000905b89838310620001d15750505010620001b6575b50505050811b0190556200011f565b01519060f884600019921b161c1916905538808080620001a7565b85870151895590970196948501948893509081019062000194565b87600052816000208480880160051c8201928489106200022d575b0160051c019087905b82811062000220575050620000f2565b6000815501879062000210565b9250819262000207565b602288634e487b7160e01b6000525260246000fd5b90607f1690620000e0565b604187634e487b7160e01b6000525260246000fd5b015190503880620000b4565b90899350601f1983169188600052856000209260005b87828210620002c85750508411620002af575b505050811b018455620000c6565b015160001983891b60f8161c19169055388080620002a1565b8385015186558d979095019493840193016200028e565b90915086600052836000208680850160051c8201928686106200032b575b918b91869594930160051c01915b8281106200031b5750506200009d565b600081558594508b91016200030b565b92508192620002fd565b602289634e487b7160e01b6000525260246000fd5b94607f169462000086565b604186634e487b7160e01b6000525260246000fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe60806040818152600480361015610029575b505050361561001f57600080fd5b6100276112a7565b005b600092833560e01c90816306fdde0314610c0057508063095ea7b314610bb857806318160ddd14610b7b57806323b872dd14610a4d578063313ce56714610a13578063395093511461099957806340c10f191461094e57806370a08231146108ed578063715018a61461084f5780638da5cb5b146107fc57806395d89b41146106a3578063a457c2d7146105a1578063a9059cbb14610552578063d46eb1191461051a578063dd62ed3e146104a2578063de0e9a3e146102285763f2fde38b0361001157346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457610124610dad565b9061012d610df8565b73ffffffffffffffffffffffffffffffffffffffff8092169283156101a1575050600554827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600555167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b90602060849251917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b8280fd5b50903461022457602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049e57823592331561041d573385528483528185205484811061039b578490338752868552038286205583600254036002558482518581527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef853392a38480808087335af13d15610396573d67ffffffffffffffff811161036a57835190610307867fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160183610e77565b815286853d92013e5b156103435750907f5c2204eac89e5412535e43df69882d6530b6696edfbd2e060e871e14afd8fb7a91519283523392a280f35b90517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b6024876041857f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b610310565b508260849251917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152fd5b8260849251917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152fd5b8380fd5b50503461051657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261051657806020926104de610dad565b6104e6610dd5565b73ffffffffffffffffffffffffffffffffffffffff91821683526001865283832091168252845220549051908152f35b5080fd5b83807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261054f5761054c6112a7565b80f35b80fd5b50503461051657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126105165760209061059a610590610dad565b6024359033610f23565b5160018152f35b50823461054f57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261054f576105d9610dad565b918360243592338152600160205281812073ffffffffffffffffffffffffffffffffffffffff861682526020522054908282106106205760208561059a8585038733611132565b60849060208651917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152fd5b5091903461051657817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261051657805191809380549160019083821c928285169485156107f2575b60209586861081146107c657858952908115610784575060011461072c575b610728878761071e828c0383610e77565b5191829182610d47565b0390f35b81529295507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b82841061077157505050826107289461071e9282010194388061070d565b8054868501880152928601928101610753565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168887015250505050151560051b830101925061071e82610728388061070d565b6024846022857f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b93607f16936106ee565b50503461051657817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126105165760209073ffffffffffffffffffffffffffffffffffffffff600554169051908152f35b833461054f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261054f57610886610df8565b8073ffffffffffffffffffffffffffffffffffffffff6005547fffffffffffffffffffffffff00000000000000000000000000000000000000008116600555167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b5050346105165760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610516578060209273ffffffffffffffffffffffffffffffffffffffff61093f610dad565b16815280845220549051908152f35b505034610516577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261054f5761054c610988610dad565b610990610df8565b602435906112df565b50503461051657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126105165761059a602092610a0c6109da610dad565b913381526001865284812073ffffffffffffffffffffffffffffffffffffffff84168252865284602435912054610ee7565b9033611132565b50503461051657817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610516576020905160128152f35b508290346105165760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261051657610a87610dad565b610a8f610dd5565b91846044359473ffffffffffffffffffffffffffffffffffffffff8416815260016020528181203382526020522054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610af5575b60208661059a878787610f23565b848210610b1e5750918391610b136020969561059a95033383611132565b919394819350610ae7565b60649060208751917f08c379a0000000000000000000000000000000000000000000000000000000008352820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b50503461051657817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610516576020906002549051908152f35b50503461051657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126105165760209061059a610bf6610dad565b6024359033611132565b929190503461049e57837ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049e57600354600181811c9186908281168015610d3d575b6020958686108214610d115750848852908115610cd15750600114610c78575b610728868661071e828b0383610e77565b929550600383527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b828410610cbe57505050826107289461071e928201019438610c67565b8054868501880152928601928101610ca1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001687860152505050151560051b830101925061071e8261072838610c67565b8360226024927f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b93607f1693610c47565b60208082528251818301819052939260005b858110610d99575050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b818101830151848201604001528201610d59565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610dd057565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff82168203610dd057565b73ffffffffffffffffffffffffffffffffffffffff600554163303610e1957565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610eb857604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b91908201809211610ef457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8091169182156110ae571691821561102a57600082815280602052604081205491808310610fa657604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152fd5b73ffffffffffffffffffffffffffffffffffffffff80911691821561122457169182156111a05760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b6112b134336112df565b6040513481527f492434c7eb6fb2007b772f848d03d4e456494843f88ea88b25c029340eaf93b760203392a2565b73ffffffffffffffffffffffffffffffffffffffff1690811561134a577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60208261132e600094600254610ee7565b60025584845283825260408420818154019055604051908152a3565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fdfea2646970667358221220c7aa44fc9276614b6107435a202f2f9621e184add2a8b90d41c3fbe376ba288364736f6c63430008140033",
}

// WEaiABI is the input ABI used to generate the binding from.
// Deprecated: Use WEaiMetaData.ABI instead.
var WEaiABI = WEaiMetaData.ABI

// WEaiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WEaiMetaData.Bin instead.
var WEaiBin = WEaiMetaData.Bin

// DeployWEai deploys a new Ethereum contract, binding an instance of WEai to it.
func DeployWEai(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WEai, error) {
	parsed, err := WEaiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WEaiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WEai{WEaiCaller: WEaiCaller{contract: contract}, WEaiTransactor: WEaiTransactor{contract: contract}, WEaiFilterer: WEaiFilterer{contract: contract}}, nil
}

// WEai is an auto generated Go binding around an Ethereum contract.
type WEai struct {
	WEaiCaller     // Read-only binding to the contract
	WEaiTransactor // Write-only binding to the contract
	WEaiFilterer   // Log filterer for contract events
}

// WEaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type WEaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WEaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WEaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WEaiSession struct {
	Contract     *WEai             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WEaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WEaiCallerSession struct {
	Contract *WEaiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WEaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WEaiTransactorSession struct {
	Contract     *WEaiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WEaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type WEaiRaw struct {
	Contract *WEai // Generic contract binding to access the raw methods on
}

// WEaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WEaiCallerRaw struct {
	Contract *WEaiCaller // Generic read-only contract binding to access the raw methods on
}

// WEaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WEaiTransactorRaw struct {
	Contract *WEaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWEai creates a new instance of WEai, bound to a specific deployed contract.
func NewWEai(address common.Address, backend bind.ContractBackend) (*WEai, error) {
	contract, err := bindWEai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WEai{WEaiCaller: WEaiCaller{contract: contract}, WEaiTransactor: WEaiTransactor{contract: contract}, WEaiFilterer: WEaiFilterer{contract: contract}}, nil
}

// NewWEaiCaller creates a new read-only instance of WEai, bound to a specific deployed contract.
func NewWEaiCaller(address common.Address, caller bind.ContractCaller) (*WEaiCaller, error) {
	contract, err := bindWEai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WEaiCaller{contract: contract}, nil
}

// NewWEaiTransactor creates a new write-only instance of WEai, bound to a specific deployed contract.
func NewWEaiTransactor(address common.Address, transactor bind.ContractTransactor) (*WEaiTransactor, error) {
	contract, err := bindWEai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WEaiTransactor{contract: contract}, nil
}

// NewWEaiFilterer creates a new log filterer instance of WEai, bound to a specific deployed contract.
func NewWEaiFilterer(address common.Address, filterer bind.ContractFilterer) (*WEaiFilterer, error) {
	contract, err := bindWEai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WEaiFilterer{contract: contract}, nil
}

// bindWEai binds a generic wrapper to an already deployed contract.
func bindWEai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WEaiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WEai *WEaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WEai.Contract.WEaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WEai *WEaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEai.Contract.WEaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WEai *WEaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WEai.Contract.WEaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WEai *WEaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WEai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WEai *WEaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WEai *WEaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WEai.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WEai *WEaiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WEai *WEaiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WEai.Contract.Allowance(&_WEai.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WEai *WEaiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WEai.Contract.Allowance(&_WEai.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WEai *WEaiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WEai *WEaiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WEai.Contract.BalanceOf(&_WEai.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WEai *WEaiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WEai.Contract.BalanceOf(&_WEai.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WEai *WEaiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WEai *WEaiSession) Decimals() (uint8, error) {
	return _WEai.Contract.Decimals(&_WEai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WEai *WEaiCallerSession) Decimals() (uint8, error) {
	return _WEai.Contract.Decimals(&_WEai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WEai *WEaiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WEai *WEaiSession) Name() (string, error) {
	return _WEai.Contract.Name(&_WEai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WEai *WEaiCallerSession) Name() (string, error) {
	return _WEai.Contract.Name(&_WEai.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WEai *WEaiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WEai *WEaiSession) Owner() (common.Address, error) {
	return _WEai.Contract.Owner(&_WEai.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WEai *WEaiCallerSession) Owner() (common.Address, error) {
	return _WEai.Contract.Owner(&_WEai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WEai *WEaiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WEai *WEaiSession) Symbol() (string, error) {
	return _WEai.Contract.Symbol(&_WEai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WEai *WEaiCallerSession) Symbol() (string, error) {
	return _WEai.Contract.Symbol(&_WEai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WEai *WEaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEai.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WEai *WEaiSession) TotalSupply() (*big.Int, error) {
	return _WEai.Contract.TotalSupply(&_WEai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WEai *WEaiCallerSession) TotalSupply() (*big.Int, error) {
	return _WEai.Contract.TotalSupply(&_WEai.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WEai *WEaiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WEai *WEaiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Approve(&_WEai.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WEai *WEaiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Approve(&_WEai.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WEai *WEaiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WEai *WEaiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.DecreaseAllowance(&_WEai.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WEai *WEaiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.DecreaseAllowance(&_WEai.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WEai *WEaiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WEai *WEaiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.IncreaseAllowance(&_WEai.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WEai *WEaiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.IncreaseAllowance(&_WEai.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_WEai *WEaiTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_WEai *WEaiSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Mint(&_WEai.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_WEai *WEaiTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Mint(&_WEai.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WEai *WEaiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WEai *WEaiSession) RenounceOwnership() (*types.Transaction, error) {
	return _WEai.Contract.RenounceOwnership(&_WEai.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WEai *WEaiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WEai.Contract.RenounceOwnership(&_WEai.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WEai *WEaiTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WEai *WEaiSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Transfer(&_WEai.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WEai *WEaiTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Transfer(&_WEai.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WEai *WEaiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WEai *WEaiSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.TransferFrom(&_WEai.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WEai *WEaiTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.TransferFrom(&_WEai.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WEai *WEaiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WEai *WEaiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WEai.Contract.TransferOwnership(&_WEai.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WEai *WEaiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WEai.Contract.TransferOwnership(&_WEai.TransactOpts, newOwner)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 amount) returns()
func (_WEai *WEaiTransactor) Unwrap(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "unwrap", amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 amount) returns()
func (_WEai *WEaiSession) Unwrap(amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Unwrap(&_WEai.TransactOpts, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 amount) returns()
func (_WEai *WEaiTransactorSession) Unwrap(amount *big.Int) (*types.Transaction, error) {
	return _WEai.Contract.Unwrap(&_WEai.TransactOpts, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xd46eb119.
//
// Solidity: function wrap() payable returns()
func (_WEai *WEaiTransactor) Wrap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEai.contract.Transact(opts, "wrap")
}

// Wrap is a paid mutator transaction binding the contract method 0xd46eb119.
//
// Solidity: function wrap() payable returns()
func (_WEai *WEaiSession) Wrap() (*types.Transaction, error) {
	return _WEai.Contract.Wrap(&_WEai.TransactOpts)
}

// Wrap is a paid mutator transaction binding the contract method 0xd46eb119.
//
// Solidity: function wrap() payable returns()
func (_WEai *WEaiTransactorSession) Wrap() (*types.Transaction, error) {
	return _WEai.Contract.Wrap(&_WEai.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WEai *WEaiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEai.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WEai *WEaiSession) Receive() (*types.Transaction, error) {
	return _WEai.Contract.Receive(&_WEai.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WEai *WEaiTransactorSession) Receive() (*types.Transaction, error) {
	return _WEai.Contract.Receive(&_WEai.TransactOpts)
}

// WEaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WEai contract.
type WEaiApprovalIterator struct {
	Event *WEaiApproval // Event containing the contract specifics and raw log

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
func (it *WEaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEaiApproval)
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
		it.Event = new(WEaiApproval)
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
func (it *WEaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEaiApproval represents a Approval event raised by the WEai contract.
type WEaiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WEai *WEaiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WEaiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WEai.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WEaiApprovalIterator{contract: _WEai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WEai *WEaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WEaiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WEai.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEaiApproval)
				if err := _WEai.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WEai *WEaiFilterer) ParseApproval(log types.Log) (*WEaiApproval, error) {
	event := new(WEaiApproval)
	if err := _WEai.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEaiEAIUnwrappedIterator is returned from FilterEAIUnwrapped and is used to iterate over the raw logs and unpacked data for EAIUnwrapped events raised by the WEai contract.
type WEaiEAIUnwrappedIterator struct {
	Event *WEaiEAIUnwrapped // Event containing the contract specifics and raw log

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
func (it *WEaiEAIUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEaiEAIUnwrapped)
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
		it.Event = new(WEaiEAIUnwrapped)
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
func (it *WEaiEAIUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEaiEAIUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEaiEAIUnwrapped represents a EAIUnwrapped event raised by the WEai contract.
type WEaiEAIUnwrapped struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEAIUnwrapped is a free log retrieval operation binding the contract event 0x5c2204eac89e5412535e43df69882d6530b6696edfbd2e060e871e14afd8fb7a.
//
// Solidity: event EAIUnwrapped(address indexed user, uint256 amount)
func (_WEai *WEaiFilterer) FilterEAIUnwrapped(opts *bind.FilterOpts, user []common.Address) (*WEaiEAIUnwrappedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEai.contract.FilterLogs(opts, "EAIUnwrapped", userRule)
	if err != nil {
		return nil, err
	}
	return &WEaiEAIUnwrappedIterator{contract: _WEai.contract, event: "EAIUnwrapped", logs: logs, sub: sub}, nil
}

// WatchEAIUnwrapped is a free log subscription operation binding the contract event 0x5c2204eac89e5412535e43df69882d6530b6696edfbd2e060e871e14afd8fb7a.
//
// Solidity: event EAIUnwrapped(address indexed user, uint256 amount)
func (_WEai *WEaiFilterer) WatchEAIUnwrapped(opts *bind.WatchOpts, sink chan<- *WEaiEAIUnwrapped, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEai.contract.WatchLogs(opts, "EAIUnwrapped", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEaiEAIUnwrapped)
				if err := _WEai.contract.UnpackLog(event, "EAIUnwrapped", log); err != nil {
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

// ParseEAIUnwrapped is a log parse operation binding the contract event 0x5c2204eac89e5412535e43df69882d6530b6696edfbd2e060e871e14afd8fb7a.
//
// Solidity: event EAIUnwrapped(address indexed user, uint256 amount)
func (_WEai *WEaiFilterer) ParseEAIUnwrapped(log types.Log) (*WEaiEAIUnwrapped, error) {
	event := new(WEaiEAIUnwrapped)
	if err := _WEai.contract.UnpackLog(event, "EAIUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEaiEAIWrappedIterator is returned from FilterEAIWrapped and is used to iterate over the raw logs and unpacked data for EAIWrapped events raised by the WEai contract.
type WEaiEAIWrappedIterator struct {
	Event *WEaiEAIWrapped // Event containing the contract specifics and raw log

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
func (it *WEaiEAIWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEaiEAIWrapped)
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
		it.Event = new(WEaiEAIWrapped)
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
func (it *WEaiEAIWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEaiEAIWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEaiEAIWrapped represents a EAIWrapped event raised by the WEai contract.
type WEaiEAIWrapped struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEAIWrapped is a free log retrieval operation binding the contract event 0x492434c7eb6fb2007b772f848d03d4e456494843f88ea88b25c029340eaf93b7.
//
// Solidity: event EAIWrapped(address indexed user, uint256 amount)
func (_WEai *WEaiFilterer) FilterEAIWrapped(opts *bind.FilterOpts, user []common.Address) (*WEaiEAIWrappedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEai.contract.FilterLogs(opts, "EAIWrapped", userRule)
	if err != nil {
		return nil, err
	}
	return &WEaiEAIWrappedIterator{contract: _WEai.contract, event: "EAIWrapped", logs: logs, sub: sub}, nil
}

// WatchEAIWrapped is a free log subscription operation binding the contract event 0x492434c7eb6fb2007b772f848d03d4e456494843f88ea88b25c029340eaf93b7.
//
// Solidity: event EAIWrapped(address indexed user, uint256 amount)
func (_WEai *WEaiFilterer) WatchEAIWrapped(opts *bind.WatchOpts, sink chan<- *WEaiEAIWrapped, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEai.contract.WatchLogs(opts, "EAIWrapped", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEaiEAIWrapped)
				if err := _WEai.contract.UnpackLog(event, "EAIWrapped", log); err != nil {
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

// ParseEAIWrapped is a log parse operation binding the contract event 0x492434c7eb6fb2007b772f848d03d4e456494843f88ea88b25c029340eaf93b7.
//
// Solidity: event EAIWrapped(address indexed user, uint256 amount)
func (_WEai *WEaiFilterer) ParseEAIWrapped(log types.Log) (*WEaiEAIWrapped, error) {
	event := new(WEaiEAIWrapped)
	if err := _WEai.contract.UnpackLog(event, "EAIWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEaiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WEai contract.
type WEaiOwnershipTransferredIterator struct {
	Event *WEaiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WEaiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEaiOwnershipTransferred)
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
		it.Event = new(WEaiOwnershipTransferred)
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
func (it *WEaiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEaiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEaiOwnershipTransferred represents a OwnershipTransferred event raised by the WEai contract.
type WEaiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WEai *WEaiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WEaiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WEai.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WEaiOwnershipTransferredIterator{contract: _WEai.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WEai *WEaiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WEaiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WEai.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEaiOwnershipTransferred)
				if err := _WEai.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WEai *WEaiFilterer) ParseOwnershipTransferred(log types.Log) (*WEaiOwnershipTransferred, error) {
	event := new(WEaiOwnershipTransferred)
	if err := _WEai.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WEai contract.
type WEaiTransferIterator struct {
	Event *WEaiTransfer // Event containing the contract specifics and raw log

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
func (it *WEaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEaiTransfer)
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
		it.Event = new(WEaiTransfer)
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
func (it *WEaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEaiTransfer represents a Transfer event raised by the WEai contract.
type WEaiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WEai *WEaiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WEaiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WEai.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WEaiTransferIterator{contract: _WEai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WEai *WEaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WEaiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WEai.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEaiTransfer)
				if err := _WEai.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WEai *WEaiFilterer) ParseTransfer(log types.Log) (*WEaiTransfer, error) {
	event := new(WEaiTransfer)
	if err := _WEai.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
