// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package worker_hub

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

// IWorkerHubAssignment is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignment struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}

// IWorkerHubDAOTokenPercentage is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubDAOTokenPercentage struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}

// IWorkerHubDAOTokenReceiverInfor is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubDAOTokenReceiverInfor struct {
	Receiver common.Address
	Amount   *big.Int
	Role     uint8
}

// IWorkerHubInference is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubInference struct {
	Assignments    []*big.Int
	Input          []byte
	Value          *big.Int
	FeeL2          *big.Int
	FeeTreasury    *big.Int
	ModelAddress   common.Address
	SubmitTimeout  *big.Int
	CommitTimeout  *big.Int
	RevealTimeout  *big.Int
	Status         uint8
	Creator        common.Address
	ProcessedMiner common.Address
	Referrer       common.Address
}

// WorkerHubMetaData contains all meta data concerning the WorkerHub contract.
var WorkerHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRevealed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySeized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"Bytes32Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotFastForward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContext\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAssignedWorker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevealTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitTimeout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIWorkerHub.DAOTokenReceiverRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenReceiverInfor[]\",\"name\":\"receivers\",\"type\":\"tuple[]\"}],\"name\":\"DAOTokenMintedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"oldValue\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"DAOTokenPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerRoleSeized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"RawSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"nonce\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"name\":\"RevealSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"getAssignmentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIWorkerHub.Assignment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getAssignmentsByInference\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"assignments\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeL2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTreasury\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wEAI\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingHub\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_feeL2Percentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeTreasuryPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_minerRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"_submitDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_commitDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_revealDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"_feeRatioMinerValidor\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_daoTokenReward\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"_daoTokenPercentage\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_referrers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_referees\",\"type\":\"address[]\"}],\"name\":\"registerReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"resolveInference\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_nonce\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"seizeMinerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDAOTokenReward\",\"type\":\"uint256\"}],\"name\":\"setDAOTokenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assigmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isReferred\",\"type\":\"bool\"}],\"name\":\"validateDAOSupplyIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"notReachedLimit\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WorkerHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerHubMetaData.ABI instead.
var WorkerHubABI = WorkerHubMetaData.ABI

// WorkerHub is an auto generated Go binding around an Ethereum contract.
type WorkerHub struct {
	WorkerHubCaller     // Read-only binding to the contract
	WorkerHubTransactor // Write-only binding to the contract
	WorkerHubFilterer   // Log filterer for contract events
}

// WorkerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerHubSession struct {
	Contract     *WorkerHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerHubCallerSession struct {
	Contract *WorkerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WorkerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerHubTransactorSession struct {
	Contract     *WorkerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WorkerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerHubRaw struct {
	Contract *WorkerHub // Generic contract binding to access the raw methods on
}

// WorkerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerHubCallerRaw struct {
	Contract *WorkerHubCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerHubTransactorRaw struct {
	Contract *WorkerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerHub creates a new instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHub(address common.Address, backend bind.ContractBackend) (*WorkerHub, error) {
	contract, err := bindWorkerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerHub{WorkerHubCaller: WorkerHubCaller{contract: contract}, WorkerHubTransactor: WorkerHubTransactor{contract: contract}, WorkerHubFilterer: WorkerHubFilterer{contract: contract}}, nil
}

// NewWorkerHubCaller creates a new read-only instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubCaller(address common.Address, caller bind.ContractCaller) (*WorkerHubCaller, error) {
	contract, err := bindWorkerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubCaller{contract: contract}, nil
}

// NewWorkerHubTransactor creates a new write-only instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerHubTransactor, error) {
	contract, err := bindWorkerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTransactor{contract: contract}, nil
}

// NewWorkerHubFilterer creates a new log filterer instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerHubFilterer, error) {
	contract, err := bindWorkerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerHubFilterer{contract: contract}, nil
}

// bindWorkerHub binds a generic wrapper to an already deployed contract.
func bindWorkerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkerHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHub *WorkerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHub.Contract.WorkerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHub *WorkerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.Contract.WorkerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHub *WorkerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHub.Contract.WorkerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHub *WorkerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHub *WorkerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHub *WorkerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHub.Contract.contract.Transact(opts, method, params...)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) AssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "assignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.AssignmentNumber(&_WorkerHub.CallOpts)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.AssignmentNumber(&_WorkerHub.CallOpts)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_WorkerHub *WorkerHubCaller) Assignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "assignments", arg0)

	outstruct := new(struct {
		InferenceId *big.Int
		Commitment  [32]byte
		Digest      [32]byte
		RevealNonce *big.Int
		Worker      common.Address
		Role        uint8
		Vote        uint8
		Output      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Digest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.RevealNonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Worker = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Role = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Vote = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Output = *abi.ConvertType(out[7], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_WorkerHub *WorkerHubSession) Assignments(arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	return _WorkerHub.Contract.Assignments(&_WorkerHub.CallOpts, arg0)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_WorkerHub *WorkerHubCallerSession) Assignments(arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	return _WorkerHub.Contract.Assignments(&_WorkerHub.CallOpts, arg0)
}

// GetAssignmentInfo is a free data retrieval call binding the contract method 0xa6ec4728.
//
// Solidity: function getAssignmentInfo(uint256 _assignmentId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes))
func (_WorkerHub *WorkerHubCaller) GetAssignmentInfo(opts *bind.CallOpts, _assignmentId *big.Int) (IWorkerHubAssignment, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAssignmentInfo", _assignmentId)

	if err != nil {
		return *new(IWorkerHubAssignment), err
	}

	out0 := *abi.ConvertType(out[0], new(IWorkerHubAssignment)).(*IWorkerHubAssignment)

	return out0, err

}

// GetAssignmentInfo is a free data retrieval call binding the contract method 0xa6ec4728.
//
// Solidity: function getAssignmentInfo(uint256 _assignmentId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes))
func (_WorkerHub *WorkerHubSession) GetAssignmentInfo(_assignmentId *big.Int) (IWorkerHubAssignment, error) {
	return _WorkerHub.Contract.GetAssignmentInfo(&_WorkerHub.CallOpts, _assignmentId)
}

// GetAssignmentInfo is a free data retrieval call binding the contract method 0xa6ec4728.
//
// Solidity: function getAssignmentInfo(uint256 _assignmentId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes))
func (_WorkerHub *WorkerHubCallerSession) GetAssignmentInfo(_assignmentId *big.Int) (IWorkerHubAssignment, error) {
	return _WorkerHub.Contract.GetAssignmentInfo(&_WorkerHub.CallOpts, _assignmentId)
}

// GetAssignmentsByInference is a free data retrieval call binding the contract method 0x9f004354.
//
// Solidity: function getAssignmentsByInference(uint256 _inferenceId) view returns(uint256[])
func (_WorkerHub *WorkerHubCaller) GetAssignmentsByInference(opts *bind.CallOpts, _inferenceId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAssignmentsByInference", _inferenceId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssignmentsByInference is a free data retrieval call binding the contract method 0x9f004354.
//
// Solidity: function getAssignmentsByInference(uint256 _inferenceId) view returns(uint256[])
func (_WorkerHub *WorkerHubSession) GetAssignmentsByInference(_inferenceId *big.Int) ([]*big.Int, error) {
	return _WorkerHub.Contract.GetAssignmentsByInference(&_WorkerHub.CallOpts, _inferenceId)
}

// GetAssignmentsByInference is a free data retrieval call binding the contract method 0x9f004354.
//
// Solidity: function getAssignmentsByInference(uint256 _inferenceId) view returns(uint256[])
func (_WorkerHub *WorkerHubCallerSession) GetAssignmentsByInference(_inferenceId *big.Int) ([]*big.Int, error) {
	return _WorkerHub.Contract.GetAssignmentsByInference(&_WorkerHub.CallOpts, _inferenceId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_WorkerHub *WorkerHubCaller) GetInferenceInfo(opts *bind.CallOpts, _inferenceId *big.Int) (IWorkerHubInference, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getInferenceInfo", _inferenceId)

	if err != nil {
		return *new(IWorkerHubInference), err
	}

	out0 := *abi.ConvertType(out[0], new(IWorkerHubInference)).(*IWorkerHubInference)

	return out0, err

}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_WorkerHub *WorkerHubSession) GetInferenceInfo(_inferenceId *big.Int) (IWorkerHubInference, error) {
	return _WorkerHub.Contract.GetInferenceInfo(&_WorkerHub.CallOpts, _inferenceId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_WorkerHub *WorkerHubCallerSession) GetInferenceInfo(_inferenceId *big.Int) (IWorkerHubInference, error) {
	return _WorkerHub.Contract.GetInferenceInfo(&_WorkerHub.CallOpts, _inferenceId)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_WorkerHub *WorkerHubCaller) GetMinFeeToUse(opts *bind.CallOpts, _modelAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinFeeToUse", _modelAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_WorkerHub *WorkerHubSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.GetMinFeeToUse(&_WorkerHub.CallOpts, _modelAddress)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.GetMinFeeToUse(&_WorkerHub.CallOpts, _modelAddress)
}

// GetTreasuryAddress is a free data retrieval call binding the contract method 0xe0024604.
//
// Solidity: function getTreasuryAddress() view returns(address)
func (_WorkerHub *WorkerHubCaller) GetTreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getTreasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasuryAddress is a free data retrieval call binding the contract method 0xe0024604.
//
// Solidity: function getTreasuryAddress() view returns(address)
func (_WorkerHub *WorkerHubSession) GetTreasuryAddress() (common.Address, error) {
	return _WorkerHub.Contract.GetTreasuryAddress(&_WorkerHub.CallOpts)
}

// GetTreasuryAddress is a free data retrieval call binding the contract method 0xe0024604.
//
// Solidity: function getTreasuryAddress() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) GetTreasuryAddress() (common.Address, error) {
	return _WorkerHub.Contract.GetTreasuryAddress(&_WorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) InferenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "inferenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHub.Contract.InferenceNumber(&_WorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHub.Contract.InferenceNumber(&_WorkerHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubSession) Owner() (common.Address, error) {
	return _WorkerHub.Contract.Owner(&_WorkerHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) Owner() (common.Address, error) {
	return _WorkerHub.Contract.Owner(&_WorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubSession) Paused() (bool, error) {
	return _WorkerHub.Contract.Paused(&_WorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) Paused() (bool, error) {
	return _WorkerHub.Contract.Paused(&_WorkerHub.CallOpts)
}

// ValidateDAOSupplyIncrease is a free data retrieval call binding the contract method 0xd7acb1ea.
//
// Solidity: function validateDAOSupplyIncrease(bool _isReferred) view returns(bool notReachedLimit)
func (_WorkerHub *WorkerHubCaller) ValidateDAOSupplyIncrease(opts *bind.CallOpts, _isReferred bool) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validateDAOSupplyIncrease", _isReferred)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateDAOSupplyIncrease is a free data retrieval call binding the contract method 0xd7acb1ea.
//
// Solidity: function validateDAOSupplyIncrease(bool _isReferred) view returns(bool notReachedLimit)
func (_WorkerHub *WorkerHubSession) ValidateDAOSupplyIncrease(_isReferred bool) (bool, error) {
	return _WorkerHub.Contract.ValidateDAOSupplyIncrease(&_WorkerHub.CallOpts, _isReferred)
}

// ValidateDAOSupplyIncrease is a free data retrieval call binding the contract method 0xd7acb1ea.
//
// Solidity: function validateDAOSupplyIncrease(bool _isReferred) view returns(bool notReachedLimit)
func (_WorkerHub *WorkerHubCallerSession) ValidateDAOSupplyIncrease(_isReferred bool) (bool, error) {
	return _WorkerHub.Contract.ValidateDAOSupplyIncrease(&_WorkerHub.CallOpts, _isReferred)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubSession) Version() (string, error) {
	return _WorkerHub.Contract.Version(&_WorkerHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubCallerSession) Version() (string, error) {
	return _WorkerHub.Contract.Version(&_WorkerHub.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_WorkerHub *WorkerHubTransactor) Commit(opts *bind.TransactOpts, _assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "commit", _assignId, _commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_WorkerHub *WorkerHubSession) Commit(_assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Commit(&_WorkerHub.TransactOpts, _assignId, _commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_WorkerHub *WorkerHubTransactorSession) Commit(_assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Commit(&_WorkerHub.TransactOpts, _assignId, _commitment)
}

// Infer is a paid mutator transaction binding the contract method 0x7c22c0e3.
//
// Solidity: function infer(bytes _input, address _creator, bool _flag) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactor) Infer(opts *bind.TransactOpts, _input []byte, _creator common.Address, _flag bool) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "infer", _input, _creator, _flag)
}

// Infer is a paid mutator transaction binding the contract method 0x7c22c0e3.
//
// Solidity: function infer(bytes _input, address _creator, bool _flag) payable returns(uint256)
func (_WorkerHub *WorkerHubSession) Infer(_input []byte, _creator common.Address, _flag bool) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer(&_WorkerHub.TransactOpts, _input, _creator, _flag)
}

// Infer is a paid mutator transaction binding the contract method 0x7c22c0e3.
//
// Solidity: function infer(bytes _input, address _creator, bool _flag) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) Infer(_input []byte, _creator common.Address, _flag bool) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer(&_WorkerHub.TransactOpts, _input, _creator, _flag)
}

// Infer0 is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactor) Infer0(opts *bind.TransactOpts, _input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "infer0", _input, _creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubSession) Infer0(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer0(&_WorkerHub.TransactOpts, _input, _creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) Infer0(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer0(&_WorkerHub.TransactOpts, _input, _creator)
}

// Initialize is a paid mutator transaction binding the contract method 0xa96c79f4.
//
// Solidity: function initialize(address _wEAI, address _l2Owner, address _treasury, address _daoToken, address _stakingHub, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint8 _minerRequirement, uint40 _submitDuration, uint40 _commitDuration, uint40 _revealDuration, uint16 _feeRatioMinerValidor, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubTransactor) Initialize(opts *bind.TransactOpts, _wEAI common.Address, _l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _stakingHub common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerRequirement uint8, _submitDuration *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _feeRatioMinerValidor uint16, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "initialize", _wEAI, _l2Owner, _treasury, _daoToken, _stakingHub, _feeL2Percentage, _feeTreasuryPercentage, _minerRequirement, _submitDuration, _commitDuration, _revealDuration, _feeRatioMinerValidor, _daoTokenReward, _daoTokenPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xa96c79f4.
//
// Solidity: function initialize(address _wEAI, address _l2Owner, address _treasury, address _daoToken, address _stakingHub, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint8 _minerRequirement, uint40 _submitDuration, uint40 _commitDuration, uint40 _revealDuration, uint16 _feeRatioMinerValidor, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubSession) Initialize(_wEAI common.Address, _l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _stakingHub common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerRequirement uint8, _submitDuration *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _feeRatioMinerValidor uint16, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _wEAI, _l2Owner, _treasury, _daoToken, _stakingHub, _feeL2Percentage, _feeTreasuryPercentage, _minerRequirement, _submitDuration, _commitDuration, _revealDuration, _feeRatioMinerValidor, _daoTokenReward, _daoTokenPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xa96c79f4.
//
// Solidity: function initialize(address _wEAI, address _l2Owner, address _treasury, address _daoToken, address _stakingHub, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint8 _minerRequirement, uint40 _submitDuration, uint40 _commitDuration, uint40 _revealDuration, uint16 _feeRatioMinerValidor, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubTransactorSession) Initialize(_wEAI common.Address, _l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _stakingHub common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerRequirement uint8, _submitDuration *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _feeRatioMinerValidor uint16, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _wEAI, _l2Owner, _treasury, _daoToken, _stakingHub, _feeL2Percentage, _feeTreasuryPercentage, _minerRequirement, _submitDuration, _commitDuration, _revealDuration, _feeRatioMinerValidor, _daoTokenReward, _daoTokenPercentage)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubSession) Pause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Pause(&_WorkerHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubTransactorSession) Pause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Pause(&_WorkerHub.TransactOpts)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_WorkerHub *WorkerHubTransactor) RegisterReferrer(opts *bind.TransactOpts, _referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerReferrer", _referrers, _referees)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_WorkerHub *WorkerHubSession) RegisterReferrer(_referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterReferrer(&_WorkerHub.TransactOpts, _referrers, _referees)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterReferrer(_referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterReferrer(&_WorkerHub.TransactOpts, _referrers, _referees)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHub.Contract.RenounceOwnership(&_WorkerHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHub.Contract.RenounceOwnership(&_WorkerHub.TransactOpts)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubTransactor) ResolveInference(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "resolveInference", _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResolveInference(&_WorkerHub.TransactOpts, _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubTransactorSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResolveInference(&_WorkerHub.TransactOpts, _inferenceId)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactor) Reveal(opts *bind.TransactOpts, _assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "reveal", _assignId, _nonce, _data)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_WorkerHub *WorkerHubSession) Reveal(_assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Reveal(&_WorkerHub.TransactOpts, _assignId, _nonce, _data)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactorSession) Reveal(_assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Reveal(&_WorkerHub.TransactOpts, _assignId, _nonce, _data)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubTransactor) SeizeMinerRole(opts *bind.TransactOpts, _assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "seizeMinerRole", _assignmentId)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubSession) SeizeMinerRole(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SeizeMinerRole(&_WorkerHub.TransactOpts, _assignmentId)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubTransactorSession) SeizeMinerRole(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SeizeMinerRole(&_WorkerHub.TransactOpts, _assignmentId)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_WorkerHub *WorkerHubTransactor) SetDAOTokenReward(opts *bind.TransactOpts, _newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setDAOTokenReward", _newDAOTokenReward)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_WorkerHub *WorkerHubSession) SetDAOTokenReward(_newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOTokenReward(&_WorkerHub.TransactOpts, _newDAOTokenReward)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetDAOTokenReward(_newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOTokenReward(&_WorkerHub.TransactOpts, _newDAOTokenReward)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactor) SubmitSolution(opts *bind.TransactOpts, _assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "submitSolution", _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitSolution(&_WorkerHub.TransactOpts, _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactorSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitSolution(&_WorkerHub.TransactOpts, _assigmentId, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.TransferOwnership(&_WorkerHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.TransferOwnership(&_WorkerHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubSession) Unpause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Unpause(&_WorkerHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Unpause(&_WorkerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubSession) Receive() (*types.Transaction, error) {
	return _WorkerHub.Contract.Receive(&_WorkerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) Receive() (*types.Transaction, error) {
	return _WorkerHub.Contract.Receive(&_WorkerHub.TransactOpts)
}

// WorkerHubCommitmentSubmissionIterator is returned from FilterCommitmentSubmission and is used to iterate over the raw logs and unpacked data for CommitmentSubmission events raised by the WorkerHub contract.
type WorkerHubCommitmentSubmissionIterator struct {
	Event *WorkerHubCommitmentSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubCommitmentSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubCommitmentSubmission)
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
		it.Event = new(WorkerHubCommitmentSubmission)
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
func (it *WorkerHubCommitmentSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubCommitmentSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubCommitmentSubmission represents a CommitmentSubmission event raised by the WorkerHub contract.
type WorkerHubCommitmentSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommitmentSubmission is a free log retrieval operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_WorkerHub *WorkerHubFilterer) FilterCommitmentSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubCommitmentSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "CommitmentSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubCommitmentSubmissionIterator{contract: _WorkerHub.contract, event: "CommitmentSubmission", logs: logs, sub: sub}, nil
}

// WatchCommitmentSubmission is a free log subscription operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_WorkerHub *WorkerHubFilterer) WatchCommitmentSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubCommitmentSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "CommitmentSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubCommitmentSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "CommitmentSubmission", log); err != nil {
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

// ParseCommitmentSubmission is a log parse operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_WorkerHub *WorkerHubFilterer) ParseCommitmentSubmission(log types.Log) (*WorkerHubCommitmentSubmission, error) {
	event := new(WorkerHubCommitmentSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "CommitmentSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubDAOTokenMintedV2Iterator is returned from FilterDAOTokenMintedV2 and is used to iterate over the raw logs and unpacked data for DAOTokenMintedV2 events raised by the WorkerHub contract.
type WorkerHubDAOTokenMintedV2Iterator struct {
	Event *WorkerHubDAOTokenMintedV2 // Event containing the contract specifics and raw log

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
func (it *WorkerHubDAOTokenMintedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubDAOTokenMintedV2)
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
		it.Event = new(WorkerHubDAOTokenMintedV2)
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
func (it *WorkerHubDAOTokenMintedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubDAOTokenMintedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubDAOTokenMintedV2 represents a DAOTokenMintedV2 event raised by the WorkerHub contract.
type WorkerHubDAOTokenMintedV2 struct {
	ChainId      *big.Int
	InferenceId  *big.Int
	ModelAddress common.Address
	Receivers    []IWorkerHubDAOTokenReceiverInfor
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenMintedV2 is a free log retrieval operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_WorkerHub *WorkerHubFilterer) FilterDAOTokenMintedV2(opts *bind.FilterOpts) (*WorkerHubDAOTokenMintedV2Iterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "DAOTokenMintedV2")
	if err != nil {
		return nil, err
	}
	return &WorkerHubDAOTokenMintedV2Iterator{contract: _WorkerHub.contract, event: "DAOTokenMintedV2", logs: logs, sub: sub}, nil
}

// WatchDAOTokenMintedV2 is a free log subscription operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_WorkerHub *WorkerHubFilterer) WatchDAOTokenMintedV2(opts *bind.WatchOpts, sink chan<- *WorkerHubDAOTokenMintedV2) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "DAOTokenMintedV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubDAOTokenMintedV2)
				if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenMintedV2", log); err != nil {
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

// ParseDAOTokenMintedV2 is a log parse operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_WorkerHub *WorkerHubFilterer) ParseDAOTokenMintedV2(log types.Log) (*WorkerHubDAOTokenMintedV2, error) {
	event := new(WorkerHubDAOTokenMintedV2)
	if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenMintedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubDAOTokenPercentageUpdatedIterator is returned from FilterDAOTokenPercentageUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenPercentageUpdated events raised by the WorkerHub contract.
type WorkerHubDAOTokenPercentageUpdatedIterator struct {
	Event *WorkerHubDAOTokenPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubDAOTokenPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubDAOTokenPercentageUpdated)
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
		it.Event = new(WorkerHubDAOTokenPercentageUpdated)
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
func (it *WorkerHubDAOTokenPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubDAOTokenPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubDAOTokenPercentageUpdated represents a DAOTokenPercentageUpdated event raised by the WorkerHub contract.
type WorkerHubDAOTokenPercentageUpdated struct {
	OldValue IWorkerHubDAOTokenPercentage
	NewValue IWorkerHubDAOTokenPercentage
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenPercentageUpdated is a free log retrieval operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_WorkerHub *WorkerHubFilterer) FilterDAOTokenPercentageUpdated(opts *bind.FilterOpts) (*WorkerHubDAOTokenPercentageUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "DAOTokenPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubDAOTokenPercentageUpdatedIterator{contract: _WorkerHub.contract, event: "DAOTokenPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenPercentageUpdated is a free log subscription operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_WorkerHub *WorkerHubFilterer) WatchDAOTokenPercentageUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubDAOTokenPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "DAOTokenPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubDAOTokenPercentageUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenPercentageUpdated", log); err != nil {
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

// ParseDAOTokenPercentageUpdated is a log parse operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_WorkerHub *WorkerHubFilterer) ParseDAOTokenPercentageUpdated(log types.Log) (*WorkerHubDAOTokenPercentageUpdated, error) {
	event := new(WorkerHubDAOTokenPercentageUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the WorkerHub contract.
type WorkerHubInferenceStatusUpdateIterator struct {
	Event *WorkerHubInferenceStatusUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInferenceStatusUpdate)
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
		it.Event = new(WorkerHubInferenceStatusUpdate)
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
func (it *WorkerHubInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the WorkerHub contract.
type WorkerHubInferenceStatusUpdate struct {
	InferenceId *big.Int
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []*big.Int) (*WorkerHubInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubInferenceStatusUpdateIterator{contract: _WorkerHub.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubInferenceStatusUpdate, inferenceId []*big.Int) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInferenceStatusUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
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

// ParseInferenceStatusUpdate is a log parse operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) ParseInferenceStatusUpdate(log types.Log) (*WorkerHubInferenceStatusUpdate, error) {
	event := new(WorkerHubInferenceStatusUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkerHub contract.
type WorkerHubInitializedIterator struct {
	Event *WorkerHubInitialized // Event containing the contract specifics and raw log

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
func (it *WorkerHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInitialized)
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
		it.Event = new(WorkerHubInitialized)
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
func (it *WorkerHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInitialized represents a Initialized event raised by the WorkerHub contract.
type WorkerHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkerHubInitializedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkerHubInitializedIterator{contract: _WorkerHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkerHubInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInitialized)
				if err := _WorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseInitialized(log types.Log) (*WorkerHubInitialized, error) {
	event := new(WorkerHubInitialized)
	if err := _WorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerRoleSeizedIterator is returned from FilterMinerRoleSeized and is used to iterate over the raw logs and unpacked data for MinerRoleSeized events raised by the WorkerHub contract.
type WorkerHubMinerRoleSeizedIterator struct {
	Event *WorkerHubMinerRoleSeized // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerRoleSeizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerRoleSeized)
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
		it.Event = new(WorkerHubMinerRoleSeized)
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
func (it *WorkerHubMinerRoleSeizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerRoleSeizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerRoleSeized represents a MinerRoleSeized event raised by the WorkerHub contract.
type WorkerHubMinerRoleSeized struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerRoleSeized is a free log retrieval operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_WorkerHub *WorkerHubFilterer) FilterMinerRoleSeized(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*WorkerHubMinerRoleSeizedIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerRoleSeized", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerRoleSeizedIterator{contract: _WorkerHub.contract, event: "MinerRoleSeized", logs: logs, sub: sub}, nil
}

// WatchMinerRoleSeized is a free log subscription operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_WorkerHub *WorkerHubFilterer) WatchMinerRoleSeized(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerRoleSeized, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerRoleSeized", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerRoleSeized)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerRoleSeized", log); err != nil {
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

// ParseMinerRoleSeized is a log parse operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_WorkerHub *WorkerHubFilterer) ParseMinerRoleSeized(log types.Log) (*WorkerHubMinerRoleSeized, error) {
	event := new(WorkerHubMinerRoleSeized)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerRoleSeized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the WorkerHub contract.
type WorkerHubNewAssignmentIterator struct {
	Event *WorkerHubNewAssignment // Event containing the contract specifics and raw log

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
func (it *WorkerHubNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewAssignment)
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
		it.Event = new(WorkerHubNewAssignment)
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
func (it *WorkerHubNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewAssignment represents a NewAssignment event raised by the WorkerHub contract.
type WorkerHubNewAssignment struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	ExpiredAt    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) FilterNewAssignment(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*WorkerHubNewAssignmentIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewAssignmentIterator{contract: _WorkerHub.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *WorkerHubNewAssignment, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewAssignment)
				if err := _WorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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

// ParseNewAssignment is a log parse operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) ParseNewAssignment(log types.Log) (*WorkerHubNewAssignment, error) {
	event := new(WorkerHubNewAssignment)
	if err := _WorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the WorkerHub contract.
type WorkerHubNewInferenceIterator struct {
	Event *WorkerHubNewInference // Event containing the contract specifics and raw log

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
func (it *WorkerHubNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewInference)
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
		it.Event = new(WorkerHubNewInference)
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
func (it *WorkerHubNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewInference represents a NewInference event raised by the WorkerHub contract.
type WorkerHubNewInference struct {
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*WorkerHubNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewInferenceIterator{contract: _WorkerHub.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *WorkerHubNewInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewInference)
				if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
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

// ParseNewInference is a log parse operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) ParseNewInference(log types.Log) (*WorkerHubNewInference, error) {
	event := new(WorkerHubNewInference)
	if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkerHub contract.
type WorkerHubOwnershipTransferredIterator struct {
	Event *WorkerHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorkerHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubOwnershipTransferred)
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
		it.Event = new(WorkerHubOwnershipTransferred)
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
func (it *WorkerHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubOwnershipTransferred represents a OwnershipTransferred event raised by the WorkerHub contract.
type WorkerHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkerHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubOwnershipTransferredIterator{contract: _WorkerHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkerHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubOwnershipTransferred)
				if err := _WorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseOwnershipTransferred(log types.Log) (*WorkerHubOwnershipTransferred, error) {
	event := new(WorkerHubOwnershipTransferred)
	if err := _WorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WorkerHub contract.
type WorkerHubPausedIterator struct {
	Event *WorkerHubPaused // Event containing the contract specifics and raw log

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
func (it *WorkerHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubPaused)
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
		it.Event = new(WorkerHubPaused)
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
func (it *WorkerHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubPaused represents a Paused event raised by the WorkerHub contract.
type WorkerHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) FilterPaused(opts *bind.FilterOpts) (*WorkerHubPausedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubPausedIterator{contract: _WorkerHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WorkerHubPaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubPaused)
				if err := _WorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParsePaused(log types.Log) (*WorkerHubPaused, error) {
	event := new(WorkerHubPaused)
	if err := _WorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRawSubmittedIterator is returned from FilterRawSubmitted and is used to iterate over the raw logs and unpacked data for RawSubmitted events raised by the WorkerHub contract.
type WorkerHubRawSubmittedIterator struct {
	Event *WorkerHubRawSubmitted // Event containing the contract specifics and raw log

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
func (it *WorkerHubRawSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRawSubmitted)
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
		it.Event = new(WorkerHubRawSubmitted)
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
func (it *WorkerHubRawSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRawSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRawSubmitted represents a RawSubmitted event raised by the WorkerHub contract.
type WorkerHubRawSubmitted struct {
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Input             []byte
	Flag              bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRawSubmitted is a free log retrieval operation binding the contract event 0x1619690726b58f924192551304a486d31e6b9753727252d31816b45f485533e8.
//
// Solidity: event RawSubmitted(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId, bytes input, bool flag)
func (_WorkerHub *WorkerHubFilterer) FilterRawSubmitted(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*WorkerHubRawSubmittedIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RawSubmitted", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRawSubmittedIterator{contract: _WorkerHub.contract, event: "RawSubmitted", logs: logs, sub: sub}, nil
}

// WatchRawSubmitted is a free log subscription operation binding the contract event 0x1619690726b58f924192551304a486d31e6b9753727252d31816b45f485533e8.
//
// Solidity: event RawSubmitted(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId, bytes input, bool flag)
func (_WorkerHub *WorkerHubFilterer) WatchRawSubmitted(opts *bind.WatchOpts, sink chan<- *WorkerHubRawSubmitted, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RawSubmitted", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRawSubmitted)
				if err := _WorkerHub.contract.UnpackLog(event, "RawSubmitted", log); err != nil {
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

// ParseRawSubmitted is a log parse operation binding the contract event 0x1619690726b58f924192551304a486d31e6b9753727252d31816b45f485533e8.
//
// Solidity: event RawSubmitted(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId, bytes input, bool flag)
func (_WorkerHub *WorkerHubFilterer) ParseRawSubmitted(log types.Log) (*WorkerHubRawSubmitted, error) {
	event := new(WorkerHubRawSubmitted)
	if err := _WorkerHub.contract.UnpackLog(event, "RawSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRevealSubmissionIterator is returned from FilterRevealSubmission and is used to iterate over the raw logs and unpacked data for RevealSubmission events raised by the WorkerHub contract.
type WorkerHubRevealSubmissionIterator struct {
	Event *WorkerHubRevealSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubRevealSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRevealSubmission)
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
		it.Event = new(WorkerHubRevealSubmission)
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
func (it *WorkerHubRevealSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRevealSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRevealSubmission represents a RevealSubmission event raised by the WorkerHub contract.
type WorkerHubRevealSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Nonce       *big.Int
	Output      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealSubmission is a free log retrieval operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_WorkerHub *WorkerHubFilterer) FilterRevealSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubRevealSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RevealSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRevealSubmissionIterator{contract: _WorkerHub.contract, event: "RevealSubmission", logs: logs, sub: sub}, nil
}

// WatchRevealSubmission is a free log subscription operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_WorkerHub *WorkerHubFilterer) WatchRevealSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubRevealSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RevealSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRevealSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "RevealSubmission", log); err != nil {
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

// ParseRevealSubmission is a log parse operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_WorkerHub *WorkerHubFilterer) ParseRevealSubmission(log types.Log) (*WorkerHubRevealSubmission, error) {
	event := new(WorkerHubRevealSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "RevealSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the WorkerHub contract.
type WorkerHubSolutionSubmissionIterator struct {
	Event *WorkerHubSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubSolutionSubmission)
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
		it.Event = new(WorkerHubSolutionSubmission)
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
func (it *WorkerHubSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubSolutionSubmission represents a SolutionSubmission event raised by the WorkerHub contract.
type WorkerHubSolutionSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubSolutionSubmissionIterator{contract: _WorkerHub.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubSolutionSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubSolutionSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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

// ParseSolutionSubmission is a log parse operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) ParseSolutionSubmission(log types.Log) (*WorkerHubSolutionSubmission, error) {
	event := new(WorkerHubSolutionSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubStreamedDataIterator is returned from FilterStreamedData and is used to iterate over the raw logs and unpacked data for StreamedData events raised by the WorkerHub contract.
type WorkerHubStreamedDataIterator struct {
	Event *WorkerHubStreamedData // Event containing the contract specifics and raw log

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
func (it *WorkerHubStreamedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubStreamedData)
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
		it.Event = new(WorkerHubStreamedData)
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
func (it *WorkerHubStreamedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubStreamedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubStreamedData represents a StreamedData event raised by the WorkerHub contract.
type WorkerHubStreamedData struct {
	AssignmentId *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStreamedData is a free log retrieval operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerHub *WorkerHubFilterer) FilterStreamedData(opts *bind.FilterOpts, assignmentId []*big.Int) (*WorkerHubStreamedDataIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubStreamedDataIterator{contract: _WorkerHub.contract, event: "StreamedData", logs: logs, sub: sub}, nil
}

// WatchStreamedData is a free log subscription operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerHub *WorkerHubFilterer) WatchStreamedData(opts *bind.WatchOpts, sink chan<- *WorkerHubStreamedData, assignmentId []*big.Int) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubStreamedData)
				if err := _WorkerHub.contract.UnpackLog(event, "StreamedData", log); err != nil {
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

// ParseStreamedData is a log parse operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerHub *WorkerHubFilterer) ParseStreamedData(log types.Log) (*WorkerHubStreamedData, error) {
	event := new(WorkerHubStreamedData)
	if err := _WorkerHub.contract.UnpackLog(event, "StreamedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WorkerHub contract.
type WorkerHubUnpausedIterator struct {
	Event *WorkerHubUnpaused // Event containing the contract specifics and raw log

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
func (it *WorkerHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubUnpaused)
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
		it.Event = new(WorkerHubUnpaused)
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
func (it *WorkerHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubUnpaused represents a Unpaused event raised by the WorkerHub contract.
type WorkerHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WorkerHubUnpausedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubUnpausedIterator{contract: _WorkerHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WorkerHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubUnpaused)
				if err := _WorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseUnpaused(log types.Log) (*WorkerHubUnpaused, error) {
	event := new(WorkerHubUnpaused)
	if err := _WorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
