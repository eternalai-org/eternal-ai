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
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRevealed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySeized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"Bytes32Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotFastForward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContext\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAssignedWorker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevealTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitTimeout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIWorkerHub.DAOTokenReceiverRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenReceiverInfor[]\",\"name\":\"receivers\",\"type\":\"tuple[]\"}],\"name\":\"DAOTokenMintedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"oldValue\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"DAOTokenPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerRoleSeized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"RawSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"nonce\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"name\":\"RevealSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"getAssignmentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIWorkerHub.Assignment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getAssignmentsByInference\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"assignments\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeL2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTreasury\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wEAI\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingHub\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_feeL2Percentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeTreasuryPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_minerRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"_submitDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_commitDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_revealDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"_feeRatioMinerValidor\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_daoTokenReward\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"_daoTokenPercentage\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_referrers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_referees\",\"type\":\"address[]\"}],\"name\":\"registerReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"resolveInference\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_nonce\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"seizeMinerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDAOTokenReward\",\"type\":\"uint256\"}],\"name\":\"setDAOTokenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wEAI\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assigmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isReferred\",\"type\":\"bool\"}],\"name\":\"validateDAOSupplyIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"notReachedLimit\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001757615c8790816200001d8239f35b600080fdfe608080604052600436101561001d575b50361561001b57600080fd5b005b60009081803560e01c91826308c0590314614aed5750508063121a301d14613b7b5780633f4ba83a14613adf5780634e50c75c14613a0557806354fd4d50146139895780635c975abb146139485780636029e786146130c75780636973d3f21461308b578063715018a614612fec5780637362323c14612f3a57806376e7ffae14612ef95780637c22c0e3146127305780638456cb59146126915780638da5cb5b1461263f5780639f004354146125b6578063a6ec47281461243f578063a96c79f414611d29578063afc1fce714611c56578063c41bf66514611b33578063d7acb1ea1461196f578063d98444581461100e578063e002460414610fbc578063e84dee6b1461082c578063f2f0387714610475578063f2fde38b14610389578063f80dca981461034d5763ffbc66610361000f573461034a576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610346576004359073ffffffffffffffffffffffffffffffffffffffff90838260125416803b15610346578180916004604051809481937f36f4fb020000000000000000000000000000000000000000000000000000000083525af1801561033b57610323575b508390526004815281600360408620015460281c1633036102f95782845260048152604084205491828552600282526007604086200154166102cf5760029083855260048152600360408620017902000000000000000000000000000000000000000000000000007fffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffff8254161790558285525260076040842001337fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905533917f3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c248480a480f35b60046040517ffbe08cbc000000000000000000000000000000000000000000000000000000008152fd5b60046040517f2ef424cc000000000000000000000000000000000000000000000000000000008152fd5b61032c90614ede565b6103375783386101e4565b8380fd5b6040513d84823e3d90fd5b5080fd5b80fd5b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576020600154604051908152f35b503461034a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576103c1615057565b6103c9615130565b73ffffffffffffffffffffffffffffffffffffffff8116156103f1576103ee906151af565b80f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b503461034a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576024356004356104b36152a7565b73ffffffffffffffffffffffffffffffffffffffff90838260125416803b15610346578180916004604051809481937f36f4fb020000000000000000000000000000000000000000000000000000000083525af1801561033b57610818575b505082156107ee5780845260209060048252604085209384549485875260028452604087209064ffffffffff80600584015460c81c16904316116107c457600682019560ff875460281c1660078110156107975760020361076d576003820154908160281c1633036107435760ff60019160c81c166105908161504d565b036107195760010180546106ef576105ab9183859255615533565b848652600583526040862080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060ff6105e6818416615593565b1691161790556040519081527f47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b833392a38284526005815260ff604085205416600982526040852054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116106c25714610663578380f35b81650300000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff7fbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d5094541617905560405160038152a23880808380f35b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b60046040517fbfec5558000000000000000000000000000000000000000000000000000000008152fd5b60046040517fd954416a000000000000000000000000000000000000000000000000000000008152fd5b60046040517f82b42900000000000000000000000000000000000000000000000000000000008152fd5b60046040517fef084b59000000000000000000000000000000000000000000000000000000008152fd5b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60046040517f88ac04e3000000000000000000000000000000000000000000000000000000008152fd5b60046040517fc06789fa000000000000000000000000000000000000000000000000000000008152fd5b61082190614ede565b610337578338610512565b503461034a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a5760043567ffffffffffffffff60243581811161033757610880903690600401614e40565b61088b9291926152a7565b73ffffffffffffffffffffffffffffffffffffffff90858260125416803b15610346578180916004604051809481937f36f4fb020000000000000000000000000000000000000000000000000000000083525af1801561033b57610fa4575b50508015610f7a57816012541692604051937f34875ec30000000000000000000000000000000000000000000000000000000085523360048601526020948581602481855afa908115610f6f578991610f42575b5015610f18578088913b156103465781602491604051928380927fd8f0166c0000000000000000000000000000000000000000000000000000000082523360048301525afa801561033b57610f00575b50869052600484526040872094604051926109a884614ef2565b86548452600193848801548782015260029586890154604083015260038901549164ffffffffff92838116606083015282610a1d6004608085019d8e848660281c16905260ff808660c81c169560a0880196610a038161504d565b875260d01c16610a128161504d565b60c087015201614f88565b9260e08101938452519b5116330361074357889051610a3b8161504d565b610a448161504d565b03610719575151610ed657888b5286885260408b20604051610a6581614ec1565b610a6e826154cc565b8152610a7b888301614f88565b8a820152888201546040820152600382015460608201526004820154608082015260058201549280841660a08301528460c0830194818160a01c16865260c81c1660e083015260068301549285841661010084015260ff8460281c16936007851015610ea857610120840185905260301c82166101408401526007810154821661016084015260080154166101809091015286900361076d578190511690431611610e7e578689528486526040892092888a5260048752600460408b2001908211610e51579189889284610b518c979654614e6e565b601f8111610df0575b5082601f8311600114610d04578a8484610bf994610c479a98967fbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d509f9e9d9c9a989560409691610cf9575b50828d1b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8460031b1c19161790555b80845196879485019889528585013782019083820152038a810184520182614f47565b5190208094838c526004885260408c208288820155015560068101650200000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff825416179055615533565b84875260068352610c6b816040892060019160005201602052604060002054151590565b15610ce0575b8652600782526040862080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060ff610cab818416615593565b169116179055604051908152a2337f9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd98380a380f35b84875260068352610cf48160408920615adc565b610c71565b905087013538610ba5565b8184528a8420969a99989796907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416855b818110610dc057509484610bf9946040947fbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d509f9997948f95610c479d9c9a10610d88575b50508b82811b019055610bd6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19908901351690553880610d7a565b949750929597509350988088819c999a9b868a013581550194019201928b95928d9795928f959c9b9a999c610d36565b9290939596508391945052878b20601f840160051c810191898510610e47575b928c92888c9693601f8f9a99970160051c01915b828110610e32575050610b5a565b600081558e99508d97508f95508a9101610e24565b9091508190610e10565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60046040517f487c58e5000000000000000000000000000000000000000000000000000000008152fd5b5060248f7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60046040517f9fbfc589000000000000000000000000000000000000000000000000000000008152fd5b610f0990614ede565b610f1457863861098e565b8680fd5b60046040517fa7c1cb49000000000000000000000000000000000000000000000000000000008152fd5b610f629150863d8811610f68575b610f5a8183614f47565b8101906154b4565b3861093e565b503d610f50565b6040513d8b823e3d90fd5b60046040517f5cb045db000000000000000000000000000000000000000000000000000000008152fd5b610fad90614ede565b610fb85785386108ea565b8580fd5b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57602073ffffffffffffffffffffffffffffffffffffffff600b5416604051908152f35b5060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a5760043567ffffffffffffffff811161034657611059903690600401614e40565b61106161507a565b9061106a6152a7565b6000604073ffffffffffffffffffffffffffffffffffffffff6012541660248251809481937fcd23ea140000000000000000000000000000000000000000000000000000000083523360048401525af19081156117b857600091611913575b506020015163ffffffff1615610743576110e460015461537c565b91826001558260005260026020526040600020600b5461ffff6111206127109182611114828660a01c16346153f9565b049360b01c16346153f9565b049067ffffffffffffffff85116118e45761113e6001840154614e6e565b601f811161189d575b50846000601f82116001146117cf57916111af916111b494936000916117c4575b508760011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8960031b1c19161760018601555b806003860155826004860155346153ec565b6153ec565b60028201556006810180547fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff0000000000008560301b16911617905573ffffffffffffffffffffffffffffffffffffffff8216600052600f602052600573ffffffffffffffffffffffffffffffffffffffff6040600020541691600881017fffffffffffffffffffffffff000000000000000000000000000000000000000093848254161790550190339082541617905564ffffffffff61129181600b5460d01c1643615484565b73ffffffffffffffffffffffffffffffffffffffff60056113466112bb85600c5416868616615491565b8860005260026020528260406000200180547dffffffffff000000000000000000000000000000000000000000000000008360c81b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff78ffffffffff00000000000000000000000000000000000000008960a01b1691161717905585600c5460281c1690615491565b87600052600260205260406000209065010000000000866006840192167fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000008354161717905501541691600073ffffffffffffffffffffffffffffffffffffffff60125416936024604051809681937f47253baa00000000000000000000000000000000000000000000000000000000835260048301525afa9283156117b8576000936116eb575b50825160ff600c5460501c16928382106116c15760005b8481106114bb57505050505050827f1619690726b58f924192551304a486d31e6b9753727252d31816b45f485533e873ffffffffffffffffffffffffffffffffffffffff6114a360209760405134815260008a82015283871690867f08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b60403393a46040519634885260008a890152608060408901526080880191615445565b936000606087015216938033940390a4604051908152f35b60ff6114fd60005460405160208101918252426040820152606043408183015281526114e681614f2b565b519020806000556114f784876153ec565b906154aa565b169061150982886153a9565b5161151482866153ec565b92837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101116116925761158e73ffffffffffffffffffffffffffffffffffffffff6115857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61168d97018c6153a9565b5116918a6153a9565b528a73ffffffffffffffffffffffffffffffffffffffff6115b060035461537c565b9283600355836000526004602052600360406000208481550179010000000000000000000000000000000000000000000000000081547fffffffffffff000000000000000000000000000000000000000000ffffffffff78ffffffffffffffffffffffffffffffffffffffff00000000008560281b169116171790551691826000526008602052611645816040600020615a7a565b81600052600960205261165c816040600020615a7a565b7f53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d560206040518a89168152a461537c565b611404565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60046040517f4069094a000000000000000000000000000000000000000000000000000000008152fd5b9290923d8083833e6116fd8183614f47565b8101906020818303126117b45780519267ffffffffffffffff841161034a5782601f85840101121561034a5783820151916117378361509d565b946117456040519687614f47565b838652602086019460208560051b8385010101116117b457602081830101945b60208560051b838501010186106117835750505050505091876113ed565b855173ffffffffffffffffffffffffffffffffffffffff811681036117b057815260209586019501611765565b8480fd5b8280fd5b6040513d6000823e3d90fd5b90508901358a611168565b600185018152602081209150805b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0881681106118855750916111af916111b49493887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081161061184d575b5050600187811b01600186015561119d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88a60031b161c19908b0135169055898061183b565b9091602060018192858d0135815501930191016117dd565b600184016000526020600020601f870160051c8101602088106118dd575b601f830160051c820181106118d1575050611147565b600081556001016118bb565b50806118bb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b906040823d604011611967575b8161192d60409383614f47565b8101031261034a5760206040519261194484614f0f565b8051845201519063ffffffff8216820361034a57506020808301919091526110c9565b3d9150611920565b503461034a57602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576004358015158103611b2e5715611a475773ffffffffffffffffffffffffffffffffffffffff600c5460581c1682600d546024604051809481937f082df8ca00000000000000000000000000000000000000000000000000000000835260048301525afa918215611a3b5791611a1e575b505b6040519015158152f35b611a359150823d8411610f6857610f5a8183614f47565b38611a12565b604051903d90823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff600c5460581c16600d54600e549061ffff91612710928082881c16840391848311611b0157611a9d9288959492611a979260301c16906153ec565b906153f9565b046024604051809481937f082df8ca00000000000000000000000000000000000000000000000000000000835260048301525afa918215611a3b5791611ae4575b50611a14565b611afb9150823d8411610f6857610f5a8183614f47565b38611ade565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b600080fd5b503461034a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a5767ffffffffffffffff6004358181116117b457611b849036906004016150b5565b906024359081116117b457611b9d9036906004016150b5565b611ba5615130565b8151815103610f7a57825b8251811015611c525773ffffffffffffffffffffffffffffffffffffffff9081611bda82866153a9565b511682611be783866153a9565b51169281158015611c4a575b610f7a57838752600f90602090828252604089205416610ed657611c459488525260408620907fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905561537c565b611bb0565b508315611bf3565b8380f35b503461034a57602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57611c90615057565b8273ffffffffffffffffffffffffffffffffffffffff602481601254169360405194859384927fafc1fce70000000000000000000000000000000000000000000000000000000084521660048301525afa918215611a3b578092611cf9575b5050604051908152f35b9091508282813d8311611d22575b611d118183614f47565b8101031261034a5750513880611cef565b503d611d07565b503461034a576102407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57611d62615057565b611d6a61507a565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203611b2e5773ffffffffffffffffffffffffffffffffffffffff6064351660643503611b2e576084359273ffffffffffffffffffffffffffffffffffffffff84168403611b2e5761ffff60a4351660a43503611b2e5761ffff60c4351660c43503611b2e5760ff60e4351660e435036117b05764ffffffffff928361010435166101043503611b2e5761012435908482168203611b2e578461014435166101443503611b2e5761ffff61016435166101643503611b2e5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe5c360112610f1457604051918260a081011067ffffffffffffffff60a0850111176118e45760a083016040526101a43561ffff81168103611b2e57835261ffff6101c435166101c43503611b2e576101c43560208401526101e43561ffff81168103611b2e57604084015261ffff61020435166102043503611b2e5761020435606084015261ffff61022435166102243503611b2e576102243560808401526077549560ff8760081c161596878098612432575b801561241b575b15612397578760017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831617607755612368575b50611f6960ff60775460081c16611f648161521c565b61521c565b611f72336151af565b60775497611fc360ff8a60081c16611f898161521c565b611f928161521c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060dc541660dc55611f648161521c565b600161010e5573ffffffffffffffffffffffffffffffffffffffff8716151580612349575b80612328575b80612309575b806122ea575b1561228c5773ffffffffffffffffffffffffffffffffffffffff8097167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55600b549387600c5492167fffffffffffffffffffffffff0000000000000000000000000000000000000000601254161760125561018435600d557fff000000000000000000000000000000000000000000000000000000000000009485897effffffffff00000000000000000000000000000000000000000000000000006101043560d01b16931691161775ffff000000000000000000000000000000000000000060a43560a01b161777ffff0000000000000000000000000000000000000000000060c43560b01b161779ffff0000000000000000000000000000000000000000000000006101643560c01b161717600b5569ffffffffff00000000006101443560281b16937effffffffffffffffffffffffffffffffffffffff000000000000000000000060643560581b169116176aff0000000000000000000060e43560501b161791161717600c5561ffff81511690600e549163ffff0000602083015160101b1665ffff00000000604084015160201b16917fffffffffffffffffffffffffffffffffffffffffffff0000000000000000000069ffff0000000000000000608067ffff000000000000606088015160301b1696015160401b1695161717171717600e55167fffffffffffffffffffffffff00000000000000000000000000000000000000006011541617601155612239575080f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166077557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a180f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f5a65726f206164647265737300000000000000000000000000000000000000006044820152fd5b5073ffffffffffffffffffffffffffffffffffffffff86161515611ffa565b5073ffffffffffffffffffffffffffffffffffffffff81161515611ff4565b5073ffffffffffffffffffffffffffffffffffffffff606435161515611fee565b5073ffffffffffffffffffffffffffffffffffffffff84161515611fe8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011760775538611f4e565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b158015611f1a5750600160ff821614611f1a565b50600160ff821610611f13565b503461034a57602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57604090606060e0835161248381614ef2565b8381528386820152838582015283838201528360808201528360a08201528360c08201520152600435815260048352206125b2604051916124c383614ef2565b80548352600181015490848401918252836002820154604082019081526003830154906060830164ffffffffff9081841681526080850192612551600460ff73ffffffffffffffffffffffffffffffffffffffff97888160281c16885260a0828260c81c169a01996125348161504d565b8a5260d01c169860c08d01996125498161504d565b8a5201614f88565b9760e08b019889526040519b8c9b818d5251908c01525160408b01525160608a015251166080880152511660a08601525161258b8161504d565b60c08501525161259a8161504d565b60e08401525161010080840152610120830190614da6565b0390f35b503461034a57602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57906004358252600981526125fe604083206154cc565b60405192828493840190808552835180925280604086019401925b82811061262857505050500390f35b835185528695509381019392810192600101612619565b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57602073ffffffffffffffffffffffffffffffffffffffff60aa5416604051908152f35b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576126c8615130565b6126d06152a7565b6126d86152a7565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060dc54161760dc557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b5060607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a5760043567ffffffffffffffff81116103465761277b903690600401614e40565b9061278461507a565b916044358015158103611b2e576127996152a7565b6000604073ffffffffffffffffffffffffffffffffffffffff6012541660248251809481937fcd23ea140000000000000000000000000000000000000000000000000000000083523360048401525af19081156117b857600091612e9d575b506020015163ffffffff16156107435761281360015461537c565b92836001558360005260026020526040600020600b5461ffff6128436127109182611114828660a01c16346153f9565b049067ffffffffffffffff86116118e4576128616001840154614e6e565b601f8111612e56575b50856000601f8211600114612d8857916111af916128d19493600091612d7d575b508860011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8a60031b1c1916176001860155806003860155826004860155346153ec565b60028201556006810180547fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff0000000000008960301b16911617905573ffffffffffffffffffffffffffffffffffffffff8616600052600f602052600573ffffffffffffffffffffffffffffffffffffffff6040600020541691600881017fffffffffffffffffffffffff000000000000000000000000000000000000000093848254161790550190339082541617905564ffffffffff6129ae81600b5460d01c1643615484565b73ffffffffffffffffffffffffffffffffffffffff6005612a636129d885600c5416868616615491565b8960005260026020528260406000200180547dffffffffff000000000000000000000000000000000000000000000000008360c81b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff78ffffffffff00000000000000000000000000000000000000008960a01b1691161717905585600c5460281c1690615491565b88600052600260205260406000209065010000000000866006840192167fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000008354161717905501541691600073ffffffffffffffffffffffffffffffffffffffff60125416936024604051809681937f47253baa00000000000000000000000000000000000000000000000000000000835260048301525afa9283156117b857600093612cb4575b50825160ff600c5460501c16928382106116c15760005b848110612bd9575050505050507f1619690726b58f924192551304a486d31e6b9753727252d31816b45f485533e873ffffffffffffffffffffffffffffffffffffffff602096612bc1879460405134815260008b82015284841690877f08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b60403393a46040519734895260008b8a0152608060408a01526080890191615445565b941515606087015216938033940390a4604051908152f35b60ff612c056000546040516020810191825242604082015243406060820152606081526114e681614f2b565b1690612c1182886153a9565b51612c1c82866153ec565b92837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81011161169257612c8d73ffffffffffffffffffffffffffffffffffffffff6115857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff612caf97018c6153a9565b528b73ffffffffffffffffffffffffffffffffffffffff6115b060035461537c565b612b21565b9290923d908183823e612cc78282614f47565b60208183810103126117b45780519267ffffffffffffffff841161034a57828201601f85840101121561034a578382015191612d028361509d565b94612d106040519687614f47565b8386526020860194820160208560051b8385010101116117b457602081830101945b60208560051b83850101018610612d50575050505050509188612b0a565b855173ffffffffffffffffffffffffffffffffffffffff811681036117b057815260209586019501612d32565b90508601358b61288b565b600185018152602081209150805b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe089168110612e3e5750916111af916128d19493897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0811610612e06575b5050600188811b01600186015561119d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88b60031b161c19908801351690558a80612df4565b9091602060018192858a013581550193019101612d96565b600184016000526020600020601f880160051c810160208910612e96575b601f830160051c82018110612e8a57505061286a565b60008155600101612e74565b5080612e74565b906040823d604011612ef1575b81612eb760409383614f47565b8101031261034a57602060405192612ece84614f0f565b8051845201519063ffffffff8216820361034a57506020808301919091526127f8565b3d9150612eaa565b503461034a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57612f31615130565b600435600d5580f35b503461034a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a5773ffffffffffffffffffffffffffffffffffffffff612f87615057565b612f8f615130565b168015612fc2577fffffffffffffffffffffffff0000000000000000000000000000000000000000601154161760115580f35b60046040517fe6c4247b000000000000000000000000000000000000000000000000000000008152fd5b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57613023615130565b600073ffffffffffffffffffffffffffffffffffffffff60aa547fffffffffffffffffffffffff0000000000000000000000000000000000000000811660aa55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576020600354604051908152f35b503461034a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576130ff6152a7565b600261010e54146138ea57600261010e558073ffffffffffffffffffffffffffffffffffffffff60125416803b156138e7578180916004604051809481937f36f4fb020000000000000000000000000000000000000000000000000000000083525af1801561033b576138d3575b505060043581526002602052604081206006810154600760ff8260281c16101561352257600160ff8260281c1614806138bc575b80613899575b61379a575b506006810154600760ff8260281c16101561352257600260ff8260281c161480613783575b61354f575b506006810154600760ff8260281c16101561352257600360ff8260281c161490816134e9575b5061324b575b600660ff91015460281c1661321a6040518092614e04565b7fbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50602060043592a2600161010e5580f35b6132566004356155d5565b61320257600435825260026020526040822060405161327481614ec1565b61327d826154cc565b815261328b60018301614f88565b6020820152600282015460408201526003820154606082015260048201546080820152600582015473ffffffffffffffffffffffffffffffffffffffff811660a083015264ffffffffff8181809360a01c1660c085015260c81c1660e08301526006830154908116610100830152600760ff8260281c1610156134bc5761337c929173ffffffffffffffffffffffffffffffffffffffff6008818460ff6133769660281c1661012086015260301c16948561014085015282600782015416610160850152015416610180820152608061336d6040830151606084015190615484565b91015190615484565b90615b8c565b60043582526009602052613392604083206154cc565b825b8151811015613483576133a781836153a9565b51845260046020526002604085200154156133cb575b6133c69061537c565b613394565b8373ffffffffffffffffffffffffffffffffffffffff601254166133ef83856153a9565b518252600460205273ffffffffffffffffffffffffffffffffffffffff600360408420015460281c16813b156117b45782916044839260405194859384927f969ceab400000000000000000000000000000000000000000000000000000000845260048401528160248401525af1801561033b5761346f575b50506133bd565b61347890614ede565b610337578338613468565b50506006810180547fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff1665040000000000179055613202565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b905064ffffffffff439116108015613502575b386131fc565b5060043582526005602052604082205460ff8082169160081c16146134fc565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60043583526005602052600160ff6040852054160160ff8111613756576004358452600960205260ff6135856040862054615a3c565b9116106135d45750600760ff600683015460281c1610156135a7575b386131d6565b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b650400000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff61364892161780600684015573ffffffffffffffffffffffffffffffffffffffff61363e6136336002860154600387015490615484565b600486015490615484565b9160301c16615b8c565b6004358252600960205261365e604083206154cc565b825b815181101561374f5761367381836153a9565b5184526004602052600160408520015415613697575b6136929061537c565b613660565b8373ffffffffffffffffffffffffffffffffffffffff601254166136bb83856153a9565b518252600460205273ffffffffffffffffffffffffffffffffffffffff600360408420015460281c16813b156117b45782916044839260405194859384927f969ceab400000000000000000000000000000000000000000000000000000000845260048401528160248401525af1801561033b5761373b575b5050613689565b61374490614ede565b610337578338613734565b50506135a1565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b5064ffffffffff600583015460c81c1643116131d1565b650500000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff6137f992161780600684015573ffffffffffffffffffffffffffffffffffffffff61363e6136336002860154600387015490615484565b8173ffffffffffffffffffffffffffffffffffffffff6012541673ffffffffffffffffffffffffffffffffffffffff600784015416813b156117b45782916044839260405194859384927f969ceab40000000000000000000000000000000000000000000000000000000084526004840152600160248401525af1801561033b57613885575b506131ac565b61388e90614ede565b61034657813861387f565b5073ffffffffffffffffffffffffffffffffffffffff60078301541615156131a7565b5064ffffffffff600583015460a01c1643116131a1565b6138dc90614ede565b61034a57803861316d565b50fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152fd5b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57602060ff60dc54166040519015158152f35b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a576125b26040516139c781614f0f565b600681527f76302e302e3200000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190614da6565b503461034a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57604090600435815260046020522080546125b2600183015492600281015490600381015473ffffffffffffffffffffffffffffffffffffffff60ff8260c81c1691613a89600460ff8360d01c169501614f88565b9460405198899889526020890152604088015264ffffffffff8116606088015260281c166080860152613abb8161504d565b60a0850152613ac98161504d565b60c08401526101008060e0850152830190614da6565b503461034a57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a57613b16615130565b613b1e615311565b613b26615311565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060dc541660dc557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b503461034a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034a5764ffffffffff602435166024350361034a5760443567ffffffffffffffff811161034657613bdd903690600401614e40565b90613be66152a7565b8273ffffffffffffffffffffffffffffffffffffffff60125416803b156103465781906004604051809481937f36f4fb020000000000000000000000000000000000000000000000000000000083525af18015614ae257614acf575b508115610f7a5764ffffffffff6024351615614aa55760043583526004602052604083209164ffffffffff600384015416614a7b57825492838552600260205260408520600681015464ffffffffff811664ffffffffff431611614a5157600760ff8260281c16101580614a2457600260ff8360281c16141580614a10575b61076d5764ffffffffff600584015460c81c1664ffffffffff43161080614996575b61496c5761493f57600260ff8260281c1614614908575b505060038101549073ffffffffffffffffffffffffffffffffffffffff8260281c16330361074357613d3160ff8360c81c1661504d565b600160ff8360c81c160361071957600181015480156148de5760405160208101907fffffffffff00000000000000000000000000000000000000000000000000000060243560d81b1682523360601b602582015285876039830137613da86039828881018c83820152036019810184520182614f47565b519020036148b457604051602081019086825284866040830137613dde6040828781018b83820152036020810184520182614f47565b519020917fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ffffffffff60243516911617600382015567ffffffffffffffff831161488757613e316004820154614e6e565b601f8111614842575b5081869184601f81116001146147785760029293899161476d575b508560011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8760031b1c19161760048201555b015583855260056020526040852080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff61ff00613ecd60ff8460081c16615593565b60081b1691161790558385526006602052613efb816040872060019160005201602052604060002054151590565b15614753575b845260076020526040842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060ff613f3c818416615593565b1691161790557ff7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc4170060405164ffffffffff6024351681526040602082015280613f8d6004359533956040840191615445565b0390a38082526005602052604082205460ff8082169160081c1614613fb0575080f35b613fb86152a7565b600261010e54146138ea57600261010e558173ffffffffffffffffffffffffffffffffffffffff60125416803b15610346578180916004604051809481937f36f4fb020000000000000000000000000000000000000000000000000000000083525af1801561033b5761473f575b50819052600260205260408220906006820154600760ff8260281c1610156143c657600160ff8260281c161480614728575b80614705575b614606575b506006820154600760ff8260281c1610156143c657600260ff8260281c1614806145ef575b6143f3575b50906006810154600760ff8260281c1610156143c657600360ff8260281c1614908161438f575b506140ff575b602060ff60067fbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d5093015460281c166140f56040518092614e04565ba2600161010e5580f35b614108826155d5565b6140ba579080835260026020526040832060405161412581614ec1565b61412e826154cc565b815261413c60018301614f88565b602082015260028201546040820152600382015460608201526004820154608082015264ffffffffff600583015473ffffffffffffffffffffffffffffffffffffffff811660a0840152818160a01c1660c084015260c81c1660e0820152600682015464ffffffffff8116610100830152600760ff8260281c16101561436257614221929173ffffffffffffffffffffffffffffffffffffffff6008818460ff6133769660281c1661012086015260301c16948561014085015282600782015416610160850152015416610180820152608061336d6040830151606084015190615484565b8083526009602052614235604084206154cc565b90835b82518110156143275761424b81846153a9565b518552600460205260026040862001541561426f575b61426a9061537c565b614238565b8473ffffffffffffffffffffffffffffffffffffffff6012541661429383866153a9565b518252600460205273ffffffffffffffffffffffffffffffffffffffff600360408420015460281c16813b156117b45782916044839260405194859384927f969ceab400000000000000000000000000000000000000000000000000000000845260048401528160248401525af1801561033b57614313575b5050614261565b61431c90614ede565b6117b057843861430c565b506006830180547fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff16650400000000001790559190506140ba565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b905064ffffffffff4391161080156143a8575b386140b4565b508183526005602052604083205460ff8082169160081c16146143a2565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b8184526005602052600160ff6040862054160160ff81116145c257828552600960205260ff6144256040872054615a3c565b911610614449575090600760ff600683015460281c161015613522575b903861408d565b650400000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff6144b392161780600685015573ffffffffffffffffffffffffffffffffffffffff61363e6144a86002870154600388015490615484565b600487015490615484565b80835260096020526144c7604084206154cc565b90835b82518110156145b9576144dd81846153a9565b5185526004602052600160408620015415614501575b6144fc9061537c565b6144ca565b8473ffffffffffffffffffffffffffffffffffffffff6012541661452583866153a9565b518252600460205273ffffffffffffffffffffffffffffffffffffffff600360408420015460281c16813b156117b45782916044839260405194859384927f969ceab400000000000000000000000000000000000000000000000000000000845260048401528160248401525af1801561033b576145a5575b50506144f3565b6145ae90614ede565b6117b057843861459e565b50919050614442565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b504364ffffffffff600585015460c81c1610614088565b650500000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff61466592161780600685015573ffffffffffffffffffffffffffffffffffffffff61363e6144a86002870154600388015490615484565b8273ffffffffffffffffffffffffffffffffffffffff6012541673ffffffffffffffffffffffffffffffffffffffff600785015416813b156117b45782916044839260405194859384927f969ceab40000000000000000000000000000000000000000000000000000000084526004840152600160248401525af1801561033b576146f1575b50614063565b6146fa90614ede565b6117b45782386146eb565b5073ffffffffffffffffffffffffffffffffffffffff600784015416151561405e565b504364ffffffffff600585015460a01c1610614058565b61474890614ede565b610346578138614026565b83855260066020526147688160408720615adc565b613f01565b905086013538613e55565b5060048101885260208820885b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe087168110614827575060029293867fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08116106147ef575b5050600185811b016004820155613e8a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c199088013516905538806147dd565b84880135825560209485019486945060019092019101614785565b60048201875260208720601f850160051c810160208610614880575b601f830160051c82018110614874575050613e3a565b6000815560010161485e565b508061485e565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60046040517f9ea6d127000000000000000000000000000000000000000000000000000000008152fd5b60046040517f81791cb4000000000000000000000000000000000000000000000000000000008152fd5b650300000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff60069216179101553880613cfa565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60046040517f8c736f8a000000000000000000000000000000000000000000000000000000008152fd5b50868852600560205260ff60408920541660096020526040892054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116149e3571415613ce3565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b50506000600360ff8360281c161415613cc1565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60046040517f514b55a4000000000000000000000000000000000000000000000000000000008152fd5b60046040517fa89ac151000000000000000000000000000000000000000000000000000000008152fd5b60046040517f756688fe000000000000000000000000000000000000000000000000000000008152fd5b614adb90939193614ede565b9138613c42565b6040513d86823e3d90fd5b8190346138e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126138e75780614b2b61018092614ec1565b60608152606060208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015282610140820152826101608201520152600435815260026020526040812060405190614b9382614ec1565b614b9c816154cc565b8252614baa60018201614f88565b602083015260028101546040830152600381015460608301526004810154608083015264ffffffffff600582015473ffffffffffffffffffffffffffffffffffffffff811660a0850152818160a01c1660c085015260c81c1660e0830152600681015464ffffffffff8116610100840152600760ff8260281c1610156143c65773ffffffffffffffffffffffffffffffffffffffff91828260ff60089460281c1661012087015260301c166101408501528260078201541661016085015201541661018082015260405190602082528051926101a060208401526101c08301845180915260206101e085019501915b818110614d905750505073ffffffffffffffffffffffffffffffffffffffff610180614cf4849560208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0878303016040880152614da6565b926040810151606086015260608101516080860152608081015160a08601528260a08201511660c086015264ffffffffff60c08201511660e086015264ffffffffff60e08201511661010086015264ffffffffff61010082015116610120860152614d69610120820151610140870190614e04565b82610140820151166101608601528261016082015116828601520151166101a08301520390f35b8251865260209586019590920191600101614c99565b919082519283825260005b848110614df05750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b602081830181015184830182015201614db1565b906007821015614e115752565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9181601f84011215611b2e5782359167ffffffffffffffff8311611b2e5760208381860195010111611b2e57565b90600182811c92168015614eb7575b6020831014614e8857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614e7d565b6101a0810190811067ffffffffffffffff8211176118e457604052565b67ffffffffffffffff81116118e457604052565b610100810190811067ffffffffffffffff8211176118e457604052565b6040810190811067ffffffffffffffff8211176118e457604052565b6080810190811067ffffffffffffffff8211176118e457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176118e457604052565b9060405191826000825492614f9c84614e6e565b90818452600194858116908160001461500b5750600114614fc8575b5050614fc692500383614f47565b565b9093915060005260209081600020936000915b818310614ff3575050614fc693508201013880614fb8565b85548884018501529485019487945091830191614fdb565b9050614fc69550602093507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013880614fb8565b60031115614e1157565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203611b2e57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203611b2e57565b67ffffffffffffffff81116118e45760051b60200190565b81601f82011215611b2e578035916150cc8361509d565b926150da6040519485614f47565b808452602092838086019260051b820101928311611b2e578301905b828210615104575050505090565b813573ffffffffffffffffffffffffffffffffffffffff81168103611b2e5781529083019083016150f6565b73ffffffffffffffffffffffffffffffffffffffff60aa5416330361515157565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60aa549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff000000000000000000000000000000000000000082161760aa55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b1561522357565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b60ff60dc54166152b357565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b60ff60dc54161561531e57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146116925760010190565b80518210156153bd5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190820391821161169257565b8181029291811591840414171561169257565b8115615416570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9190820180921161169257565b91909164ffffffffff8080941691160191821161169257565b8115615416570690565b90816020910312611b2e57518015158103611b2e5790565b9060405191828154918282526020928383019160005283600020936000905b82821061550157505050614fc692500383614f47565b8554845260019586019588955093810193909101906154eb565b80548210156153bd5760005260206000200190600090565b805490680100000000000000008210156118e4578161555a91600161558f9401815561551b565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b60ff1660ff81146116925760010190565b60ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019060ff821161169257565b90600091828181526020600681526040928383208451808285829454938481520190875285872092875b87828210615a265750505061561692500382614f47565b818452600683528484205490845b8281106159e25750505080600052600982526156438460002054615a3c565b60ff8085169182106159d65782600052600294858552615665876000206154cc565b9889519283156153bd578591878c01516000526004958689526156b3828b8d600020015414918b8d8c82600091829a8352522001549561ffff9182600e5416612710948591600d54906153f9565b0494156159b75750506156f961570e61571494846157076157006156ef8c9d986156e78d99600b5460c01c169e8f906153f9565b04809e6153ec565b876156f98a6155a4565b169061540c565b9a836153f9565b04906153ec565b926155a4565b505b60005b8581106157e0575050505050506006959650816000528383526003856000200154806157b8575b508160005283835284600020015480615790575b506000525260002001650600000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff825416179055600190565b6157b29073ffffffffffffffffffffffffffffffffffffffff600b5416615b8c565b38615754565b6157da9073ffffffffffffffffffffffffffffffffffffffff600a5416615b8c565b38615740565b6157ea818e6153a9565b516000908152878a528b90208a81015483146158e4576003017a0100000000000000000000000000000000000000000000000000007fffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffff8254161780915573ffffffffffffffffffffffffffffffffffffffff90816012541691823b15611b2e5760446000928f8c9585915196879586947f969ceab400000000000000000000000000000000000000000000000000000000865260281c1690840152600160248401525af180156158d957906158c592916158ca575b5061537c565b615719565b6158d390614ede565b386158bf565b8c513d6000823e3d90fd5b9060036158c592017a0200000000000000000000000000000000000000000000000000007fffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffff825416178082556001868260c81c166159418161504d565b0361597f57508580615955575b505061537c565b73ffffffffffffffffffffffffffffffffffffffff615978925460281c16615b8c565b388561594e565b9050868061598e57505061537c565b73ffffffffffffffffffffffffffffffffffffffff6159b09260281c16615b8c565b388661594e565b92509650506159ca816159d0939661540c565b9461540c565b50615716565b50600096505050505050565b6159ec81836153a9565b51806000526007865260ff808960002054169088168111615a18575b5050615a139061537c565b615624565b9099509550615a1338615a08565b85548452600195860195879550930192016155ff565b8060011b81810460021482151715611692576003615a6c920615600014615a6f57600360ff60005b169104615484565b90565b600360ff6001615a64565b60018101908260005281602052604060002054615aab57615a9b8382615533565b5491600052602052604060002055565b602483604051907f346c4a0e0000000000000000000000000000000000000000000000000000000082526004820152fd5b919060018301600090828252806020526040822054615b5b5784549468010000000000000000861015615b2e5783615b2161555a886001604098999a0185558461551b565b9055549382526020522055565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b602483604051907fbc436f5e0000000000000000000000000000000000000000000000000000000082526004820152fd5b60008080809481945af1903d15615c4b573d9067ffffffffffffffff8211615c1e5760405191615be460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160184614f47565b825260203d92013e5b15615bf457565b60046040517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b50615bed56fea26469706673582212202d771d73dd050e04df7fb85cd363d4c6bb7d92df31c3674c2550547a21a0c73264736f6c63430008140033",
}

// WorkerHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerHubMetaData.ABI instead.
var WorkerHubABI = WorkerHubMetaData.ABI

// WorkerHubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WorkerHubMetaData.Bin instead.
var WorkerHubBin = WorkerHubMetaData.Bin

// DeployWorkerHub deploys a new Ethereum contract, binding an instance of WorkerHub to it.
func DeployWorkerHub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WorkerHub, error) {
	parsed, err := WorkerHubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WorkerHubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WorkerHub{WorkerHubCaller: WorkerHubCaller{contract: contract}, WorkerHubTransactor: WorkerHubTransactor{contract: contract}, WorkerHubFilterer: WorkerHubFilterer{contract: contract}}, nil
}

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

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address _wEAI) returns()
func (_WorkerHub *WorkerHubTransactor) SetWEAIAddress(opts *bind.TransactOpts, _wEAI common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setWEAIAddress", _wEAI)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address _wEAI) returns()
func (_WorkerHub *WorkerHubSession) SetWEAIAddress(_wEAI common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetWEAIAddress(&_WorkerHub.TransactOpts, _wEAI)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address _wEAI) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetWEAIAddress(_wEAI common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetWEAIAddress(&_WorkerHub.TransactOpts, _wEAI)
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
