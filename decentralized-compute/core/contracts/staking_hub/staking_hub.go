// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_hub

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

// IStakingHubModel is an auto generated low-level Go binding around an user-defined struct.
type IStakingHubModel struct {
	MinimumFee *big.Int
	Tier       uint32
}

// IStakingHubUnstakeRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakingHubUnstakeRequest struct {
	Stake    *big.Int
	UnlockAt *big.Int
}

// StakingHubMetaData contains all meta data concerning the StakingHub contract.
var StakingHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWorkerHub\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinerInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameModelAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillBeingLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocks\",\"type\":\"uint256\"}],\"name\":\"BlocksPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldPercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPercent\",\"type\":\"uint16\"}],\"name\":\"FinePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"FraudulentMinerPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinFeeToUseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"}],\"name\":\"MinerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"MinerUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"ModelTierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"oldDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newDuration\",\"type\":\"uint40\"}],\"name\":\"PenaltyDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restake\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelayTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelayTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blocksPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"forceChangeModelForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMinerUnstakeRequests\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"unstakeAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"internalType\":\"structIStakingHub.UnstakeRequest[]\",\"name\":\"unstakeRequests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"getMinerAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_modelAddr\",\"type\":\"address\"}],\"name\":\"getModelInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"internalType\":\"structIStakingHub.Model\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNOMiner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wEAIAmt\",\"type\":\"uint256\"}],\"name\":\"increaseMinerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wEAI\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blocksPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_unstakeDelayTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_penaltyDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minFeeToUse\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"isMinerAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minerUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"restakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaskCompleted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMiner\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"}],\"name\":\"setFinePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minFeeToUse\",\"type\":\"uint256\"}],\"name\":\"setMinFeeToUse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinerMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardAmount\",\"type\":\"uint256\"}],\"name\":\"setNewRewardInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_penaltyDuration\",\"type\":\"uint40\"}],\"name\":\"setPenaltyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newUnstakeDelayTime\",\"type\":\"uint40\"}],\"name\":\"setUnstakDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wEAI\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_workerHub\",\"type\":\"address\"}],\"name\":\"setWorkerHubAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isFined\",\"type\":\"bool\"}],\"name\":\"slashMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeDelayTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"updateModelMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_tier\",\"type\":\"uint32\"}],\"name\":\"updateModelTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"validateModelOfMiner\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wEAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001657614734908161001c8239f35b600080fdfe60406080815260049081361015610020575b5050361561001e57600080fd5b005b600091823560e01c8063034438b0146131ce5780630716187f1461318a5780630738a9bb1461307d5780630dc7df531461302a578063191a54d814612fb75780631a8ef58414612d225780631fdadcb714612b8a5780632a1a8ca814612b4d5780633304f45614612b10578063339d0f781461292757806334875ec3146128b8578063351b2b33146127d857806336f4fb0214612723578063431a44571461266557806347253baa146125c35780634fb9bc1e146123d257806354917f83146123605780635aa1326c146123195780635c975abb146122d757806361d027b31461228457806361ea0a2514611f49578063648ec7b914611ea7578063652ff15914611e40578063656a1b2014611b1a578063674a63b914611ad6578063715018a614611a355780637362323c1461197c57806373df250d1461187c578063741727951461183857806376671808146117f1578063806b984f146117b457806384449a9d14611777578063885b050f146116b85780638da5cb5b146116655780639280f078146113f5578063969ceab4146112d75780639ae49cd314611217578063a8d6d3d114610fd0578063a9b3f8b714610f85578063af5e3be014610f0c578063afc1fce714610eaa578063b1d1a56b14610dc5578063b74cd19414610ced578063c325dc6314610af2578063cd23ea1414610a53578063d279c191146108f5578063d2d89be8146108b8578063d8f0166c14610812578063db2dab1d146105d7578063e32bd90c1461055e578063e4fefd651461051a578063e69d5b98146104d2578063e8d6f2f11461047f578063f06820541461043e578063f2fde38b146103515763f6a6ae1d146102955750610011565b3461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d5773ffffffffffffffffffffffffffffffffffffffff6102e1613277565b6102e9613331565b6102f1613774565b169182156103275750507fffffffffffffffffffffffff0000000000000000000000000000000000000000601354161760135580f35b517fec08c5b9000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b503461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57610389613277565b91610392613331565b73ffffffffffffffffffffffffffffffffffffffff8316156103bb57836103b8846133b0565b80f35b90602060849251917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602090600d549051908152f35b5080fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576104ce906104bb61401d565b90519182916020835260208301906132e7565b0390f35b83823461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5761050b613331565b610513613774565b35600b5580f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209064ffffffffff600c54169051908152f35b503461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d577f3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c4639035916105ba613331565b6105c2613774565b600f548151908152836020820152a1600f5580f35b5091903461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57610611613277565b9061061a613331565b610622613774565b73ffffffffffffffffffffffffffffffffffffffff809216918284526002602052600182852001805463ffffffff8116156107ea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000016905582845260066020528184205480156107bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9161070c6107076106c285600554016134e1565b9190549186860192856106d4856134e1565b92909360031b1c169082549060031b9173ffffffffffffffffffffffffffffffffffffffff809116831b921b1916179055565b6134e1565b90549060031b1c168552600660205282852055600554801561078f578492910161075a610738826134e1565b73ffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055565b60055582825260066020528120557f68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d78280a280f35b6024856031887f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b602486858551917f28a32c19000000000000000000000000000000000000000000000000000000008352820152fd5b8684517faba47339000000000000000000000000000000000000000000000000000000008152fd5b50903461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d5773ffffffffffffffffffffffffffffffffffffffff80610861613277565b1690818552600360205260028386200154168452826020526001828520019084526020528083205415610892578280f35b517fa7c1cb49000000000000000000000000000000000000000000000000000000008152fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576020906007549051908152f35b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5761092e613277565b906109376136e4565b61093f613596565b61094882613d8f565b9064ffffffffff600c5460701c169273ffffffffffffffffffffffffffffffffffffffff6109cb818316958688526003602052600285892001907fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff78ffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b8315158080610a4b575b15610a2a575091610a1b847f75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c299593602095888a5260108752898581205560125416614492565b51908152a25b60016101105580f35b929394505050610a3c575b5050610a21565b60106020528220553880610a35565b5060016109d5565b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57809173ffffffffffffffffffffffffffffffffffffffff610aa3613277565b8260208551610ab18161363c565b828152015216815260026020522060018251610acc8161363c565b602083549384835263ffffffff9384910154169101908152835192835251166020820152f35b50903461034d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57610b2a6132c2565b91610b3361329f565b90610b3c613596565b610b44613774565b73ffffffffffffffffffffffffffffffffffffffff8092168015610cc55780865260066020528386205415610cc55761ffff94858116958615908115610cb4575b50610c8c5733875260036020528487206002810193845460f01c610c655750600b54905582547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178255610c349291907fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905560125416600b54903090339061459e565b600b5490519081527f55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a160203392a380f35b86517f3a81d6fc000000000000000000000000000000000000000000000000000000008152fd5b8285517fe1423617000000000000000000000000000000000000000000000000000000008152fd5b9050600c5460601c16861138610b85565b5082517f13aa293e000000000000000000000000000000000000000000000000000000008152fd5b50903461034d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57610d25613277565b9073ffffffffffffffffffffffffffffffffffffffff60243592610d47613331565b610d4f613774565b169283855260026020528185209063ffffffff60018301541615610d9e575091602091817f923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7945551908152a280f35b82517f13aa293e000000000000000000000000000000000000000000000000000000008152fd5b50903461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57813591610e01613596565b610e09613774565b338452600360205281842090600282015460f01c15610e835750610e2e838254613600565b9055610e568273ffffffffffffffffffffffffffffffffffffffff601254163090339061459e565b519081527f3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c160203392a280f35b82517faba47339000000000000000000000000000000000000000000000000000000008152fd5b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b578060209273ffffffffffffffffffffffffffffffffffffffff610efc613277565b1681526002845220549051908152f35b503461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d577f37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c903591610f68613331565b610f70613774565b6014548151908152836020820152a160145580f35b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602090610fc9610fc4613277565b613eb6565b9051908152f35b503461034d5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57611008613277565b6024359261ffff84168094036112135760443591611024613331565b61102c613774565b73ffffffffffffffffffffffffffffffffffffffff8116938415610d9e5760145484106111ec5785156111c5578487526002602052828720600181019081549063ffffffff821661119d578690557fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000168717905584875260066020528287205461116e57600554906801000000000000000082101561114257509161112c7f7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c94926110ff856001602097016005556134e1565b90919082549060031b9173ffffffffffffffffffffffffffffffffffffffff809116831b921b1916179055565b600554858852600684528188205551908152a380f35b8760416024927f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b602490858451917f05e3de8f000000000000000000000000000000000000000000000000000000008352820152fd5b8386517f3a81d6fc000000000000000000000000000000000000000000000000000000008152fd5b82517fe1423617000000000000000000000000000000000000000000000000000000008152fd5b82517f732f9413000000000000000000000000000000000000000000000000000000008152fd5b8480fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5780519081809360055490818552602080950191600582527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0915b868282106112aa5785906104ce8861129b848903856136a3565b519282849384528301906132e7565b835473ffffffffffffffffffffffffffffffffffffffff1685528895509093019260019283019201611281565b5091903461047b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57611310613277565b6024359182151583036113f15773ffffffffffffffffffffffffffffffffffffffff8060ac5416331480156113e4575b156113875761134d613774565b82161561135f5750906103b8916138fe565b8490517fa7c1cb49000000000000000000000000000000000000000000000000000000008152fd5b60648660208451917f08c379a0000000000000000000000000000000000000000000000000000000008352820152601760248201527f4f6e6c79204f776e6572206f7220576f726b65724875620000000000000000006044820152fd5b5080601354163314611340565b8380fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5761142d61401d565b9180805b84518110156114ae5773ffffffffffffffffffffffffffffffffffffffff61145982876140cf565b5116825260206009815284832064ffffffffff600187519261147a8461363c565b805484520154169182910152611499575b611494906140a2565b611431565b916114a6611494916140a2565b92905061148b565b50906114b9816140e3565b916114c6845193846136a3565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe094856114f5846140e3565b0195602096368887013761152061150b856140e3565b94611518885196876136a3565b8086526140e3565b0186835b828110611644575050508192825b82518110156115e05773ffffffffffffffffffffffffffffffffffffffff8061155b83866140cf565b51168552600989528785209064ffffffffff60018a519361157b8561363c565b80548552015416808b84015261159c575b5050611597906140a2565b611532565b956115d89161159793976115b083886140cf565b51166115bc838b6140cf565b526115c782866140cf565b526115d281856140cf565b506140a2565b94903861158c565b86518781528083868a8c6115f68583018d6132e7565b9185830382870152818086519485815201950193905b8382106116195786860387f35b84518051875283015164ffffffffff168684015287965094850194938201936001919091019061160c565b819088516116518161363c565b868152868382015282828901015201611524565b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209073ffffffffffffffffffffffffffffffffffffffff60ac54169051908152f35b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b577fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff69ffffffffff000000000061171d6132d3565b611725613331565b61172d613774565b7ff7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e600c5494805164ffffffffff808860281c16825284166020820152a160281b16911617600c5580f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602090600f549051908152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602090600e549051908152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209064ffffffffff600c5460701c169051908152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209061ffff600c5460501c169051908152f35b50903461034d57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d576118b4613774565b338352600960205280832064ffffffffff600182015416431061195557805492831561192e5750839055611901823373ffffffffffffffffffffffffffffffffffffffff60125416614492565b519081527f1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b460203392a280f35b82517f15da940f000000000000000000000000000000000000000000000000000000008152fd5b50517fb3c383a1000000000000000000000000000000000000000000000000000000008152fd5b503461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d5773ffffffffffffffffffffffffffffffffffffffff6119c9613277565b6119d1613331565b6119d9613774565b16918215611a0f5750507fffffffffffffffffffffffff0000000000000000000000000000000000000000601254161760125580f35b517fe6c4247b000000000000000000000000000000000000000000000000000000008152fd5b8334611ad357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ad357611a6c613331565b8073ffffffffffffffffffffffffffffffffffffffff60ac547fffffffffffffffffffffffff0000000000000000000000000000000000000000811660ac55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602090610fc9611b15613277565b613d8f565b50903461034d57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57611b526136e4565b611b5a613596565b611b62613774565b3383526020916003835281842090600282019081548060f01c15611e185791600193917dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611c1d9416825587858454948281550155611bdf3373ffffffffffffffffffffffffffffffffffffffff166000526008602052604060002090565b54611cae575b507fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055338652600985528386205490613600565b64ffffffffff92600984611c3581600c541643613600565b1695825193611c438561363c565b845280840196875233885252852090518155019151167fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000825416179055337f8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e8280a260016101105580f35b611cb6613596565b611cbf33613d8f565b600c54338a5260038952878a2060020180547fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff1660309290921b78ffffffffff00000000000000000000000000000000000000001691909117905564ffffffffff908015158080611e11575b15611df85750338a52601089528988812055611d60813373ffffffffffffffffffffffffffffffffffffffff60125416614492565b87519081527f75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29893392a25b33895260118852868920907fffffffffffffffffffffffffffffffff000000000000ffffffffff000000000082549142169116179055611dca336141f3565b73ffffffffffffffffffffffffffffffffffffffff82541688528652611df233868920614306565b38611be5565b611e03575b50611d8b565b60108952878a205538611dfd565b508a611d2b565b5083517faba47339000000000000000000000000000000000000000000000000000000008152fd5b503461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d576080928291358152600a60205220908154916001810154916003600283015492015492815194855260208501528301526060820152f35b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57908160c09273ffffffffffffffffffffffffffffffffffffffff9283611efc613277565b1681526003602052209081549260026001840154930154928251948552602085015282169083015264ffffffffff808260a01c1660608401528160c81c16608083015260f01c60a0820152f35b50903461034d576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57611f83613277565b906084359164ffffffffff938484168094036122805760a43594851685036122805760c43561ffff8116810361227c5760799586549560ff8760081c16159485809661226f575b8015612258575b156121d55773ffffffffffffffffffffffffffffffffffffffff907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00988760018b8316178c556121a7575b5061203660ff8b5460081c166120318161341d565b61341d565b61203f336133b0565b61206f8a549960ff8b60081c16906120568261341d565b61205f8261341d565b60de541660de556120318161341d565b6001610110551693841561218057506c01000000000000000000000000929169ffffffffff00000000006bffff0000000000000000000092602435600b55604435600d55606435600f557fffffffffffffffffffffffffffffffffffff0000000000000000000000000000600c5460e43560145516179160281b16179160501b161717600c5543600e557fffffffffffffffffffffffff0000000000000000000000000000000000000000601254161760125561212a578380f35b7f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff6020931690555160018152a13880808380f35b86517fe6c4247b000000000000000000000000000000000000000000000000000000008152fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016610101178a553861201c565b60848560208951917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b158015611fd15750600160ff891614611fd1565b50600160ff891610611fca565b8680fd5b8580fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209073ffffffffffffffffffffffffffffffffffffffff600154169051908152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209060ff60de541690519015158152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209064ffffffffff600c5460281c169051908152f35b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57809173ffffffffffffffffffffffffffffffffffffffff6123b0613277565b16815260026020522063ffffffff600182549201541682519182526020820152f35b503461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d5761240a6132c2565b90612413613596565b61241b613774565b338452600960205282842091825492831561259b578590553385526003602052600284862061244b858254613600565b81550191825460f01c156124f8575b505073ffffffffffffffffffffffffffffffffffffffff9081815416156124ac575b541691519081527f5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d60203392a380f35b816124c46107076124bb6140fb565b600554906134a8565b90549060031b1c167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617815561247c565b61ffff808316801591821561258a575b5050612563575081547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178155388061245a565b84517fe1423617000000000000000000000000000000000000000000000000000000008152fd5b600c5460601c161090503880612508565b8285517f7c946ed7000000000000000000000000000000000000000000000000000000008152fd5b50903461034d57602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126113f15773ffffffffffffffffffffffffffffffffffffffff9283612615613277565b16855282528084209381518094848754928381520196835284832092905b82821061264c5786866104ce8761129b848d03856136a3565b8354811688529685019660019384019390910190612633565b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b577fffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffffff6bffff000000000000000000006126cc6132c2565b6126d4613331565b6126dc613774565b7fcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269600c5494805161ffff808860501c16825284166020820152a160501b16911617600c5580f35b50903461034d57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d5773ffffffffffffffffffffffffffffffffffffffff60135416330361277b57826103b8613774565b90602060649251917f08c379a0000000000000000000000000000000000000000000000000000000008352820152600e60248201527f4f6e6c7920576f726b65724875620000000000000000000000000000000000006044820152fd5b50903461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d576128116132d3565b9061281a613331565b612822613774565b64ffffffffff8092169283156128915750907fdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd4907fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000092600c549281519084168152856020820152a11617600c5580f35b90517faa7feadc000000000000000000000000000000000000000000000000000000008152fd5b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209061291d6128f7613277565b73ffffffffffffffffffffffffffffffffffffffff166000526008602052604060002090565b5415159051908152f35b5091903461047b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57612960613277565b9061296961329f565b91612972613331565b61297a613774565b73ffffffffffffffffffffffffffffffffffffffff8093168085526020916002835263ffffffff600185882001541615612ae8576129d88173ffffffffffffffffffffffffffffffffffffffff166000526008602052604060002090565b54156107ea5784811694858752600384526002858820015416828114612ac05791612a27600295949288996003956103b89a52808552612a1a82888d20614306565b838b528452858a2061412c565b8588528282528484892001817fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055875283815261ffff600184892001541694875252842001907dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffff00000000000000000000000000000000000000000000000000000000000083549260f01b169116179055565b8785517f77a9a35d000000000000000000000000000000000000000000000000000000008152fd5b8684517f13aa293e000000000000000000000000000000000000000000000000000000008152fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602090600b549051908152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576020906014549051908152f35b50903461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57612bc36132c2565b91612bcc613596565b612bd4613774565b61ffff92838116938415908115612d11575b50612ce95733855260036020528285206002810192835460f01c612cc25750600b54905581547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178155610c3490612c696107076124bb6140fb565b919073ffffffffffffffffffffffffffffffffffffffff928391549060031b1c167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905560125416600b54903090339061459e565b84517f3a81d6fc000000000000000000000000000000000000000000000000000000008152fd5b5090517fe1423617000000000000000000000000000000000000000000000000000000008152fd5b9050600c5460601c16841138612be6565b50903461034d57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d57612d5a613596565b612d62613774565b33835260036020528083209160028301908154938460f01c15612f8f5754600b5411612f685764ffffffffff93848160c81c164210612f405773ffffffffffffffffffffffffffffffffffffffff16855280602052612dc33384872061412c565b338552600860205282852054612f10576007549068010000000000000000821015612ee4575090612e2b612e00836001612e8b9501600755613547565b815473ffffffffffffffffffffffffffffffffffffffff60039290921b91821b19163390911b179055565b60075433865260086020528386205583600c5460701c167fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff78ffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b601160205282209042167fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000825416179055337fb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc54468280a280f35b8560416024927f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b6024908351907f05e3de8f0000000000000000000000000000000000000000000000000000000082523390820152fd5b5082517ff7bc3778000000000000000000000000000000000000000000000000000000008152fd5b82517f1cc3b37b000000000000000000000000000000000000000000000000000000008152fd5b5082517faba47339000000000000000000000000000000000000000000000000000000008152fd5b50503461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57809173ffffffffffffffffffffffffffffffffffffffff613007613277565b16815260096020522064ffffffffff600182549201541682519182526020820152f35b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209073ffffffffffffffffffffffffffffffffffffffff601254169051908152f35b503461034d57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d576130b4613277565b916024359163ffffffff93848416809403612280576130d1613331565b6130d9613774565b8315612ce95773ffffffffffffffffffffffffffffffffffffffff169384865260026020526001838720019182549182161561316357507fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001683179055519081527f64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a1590602090a280f35b83517f13aa293e000000000000000000000000000000000000000000000000000000008152fd5b50503461047b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760209061ffff600c5460601c169051908152f35b50903461034d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261034d5781359161320a613331565b613212613774565b821561325057507f3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d90600d548151908152836020820152a1600d5580f35b90517f1342eb4b000000000000000000000000000000000000000000000000000000008152fd5b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361329a57565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361329a57565b6004359061ffff8216820361329a57565b6004359064ffffffffff8216820361329a57565b90815180825260208080930193019160005b828110613307575050505090565b835173ffffffffffffffffffffffffffffffffffffffff16855293810193928101926001016132f9565b73ffffffffffffffffffffffffffffffffffffffff60ac5416330361335257565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60ac549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff000000000000000000000000000000000000000082161760ac55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b1561342457565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b81156134b2570690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6005548110156135185760056000527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6007548110156135185760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880190600090565b80548210156135185760005260206000200190600090565b60ff60de54166135a257565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b9190820180921161360d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6040810190811067ffffffffffffffff82111761365857604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761365857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761365857604052565b61011060028154146136f65760029055565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152fd5b9190820391821161360d57565b8181029291811591840414171561360d57565b600d54156138f857600e546137898143613754565b90600d549182156134b25782900491826137a3575b505050565b91906137b9906137b38385613761565b90613600565b600e558062f099c06137db600f54946137d6600796875490613761565b613761565b04915b1561379e57825464ffffffffff90600c9182549160709180600094841c168452600a60205260409160038386200155808554841c1684528660018093862001558454938185851c169182146138cb57509172ffffffffff0000000000000000000000000000917fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9301901b169116179055801561389d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01806137de565b602460007f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526011600452fd5b43600e55565b73ffffffffffffffffffffffffffffffffffffffff818116600081815260036020908152604080832060088352818420909794959294919392905415613d6657613946613596565b61394f87613d8f565b600c548785526003875285852060020180547fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff1660309290921b78ffffffffff00000000000000000000000000000000000000001691909117905564ffffffffff939087878215158080613d5f575b15613d43575081927f75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29928552601082528489812055613a02818d8860125416614492565b8851908152a25b868152601180875285822054851680613d3e5750636648cc0b5b8565ffffffffffff911681421603818111613d1257898452828952878420918083549216818360501c1601908111613ce657906fffffffffffff000000000000000000009392917fffffffffffffffffffffffffffffffff000000000000ffffffffffffffffffff94859160501b1691161790558883528088528683208642167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000082541617905560028b0194848654169a8b855260048a526001898620018b86528a5288852054613cc4575b50613b0187600c5460281c1642613600565b7fffff0000000000ffffffffffffffffffffffffffffffffffffffffffffffffff7dffffffffff0000000000000000000000000000000000000000000000000088549260c81b169116178655613b8757505050507f9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d949596505460c81c169051908152a3565b909160049796939550809450612710613bac600b5461ffff600c5460501c1690613761565b049a8b8154818110600014613cb2575050555b87845285528183209081541690558383601354168251968780927fe00246040000000000000000000000000000000000000000000000000000000082525afa918215613ca7578092613c48575b507f63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e439450613c3f88838560125416614492565b519687521694a4565b91509192938481813d8311613ca0575b613c6281836136a3565b8101031261047b5751908382168203611ad357507f63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e4393929138613c0c565b503d613c58565b9051903d90823e3d90fd5b613cbd929350613754565b9055613bbf565b613ce0908c865260048b52613cdb818b8820614306565b6141f3565b38613aef565b602485857f4e487b71000000000000000000000000000000000000000000000000000000008252600452fd5b602484847f4e487b71000000000000000000000000000000000000000000000000000000008252600452fd5b613a23565b915050613d51575b50613a09565b601087528582205538613d4b565b50846139be565b600484517fa7c1cb49000000000000000000000000000000000000000000000000000000008152fd5b613e1690613d9b613774565b600064ffffffffff9081600c5460701c1690613dd78473ffffffffffffffffffffffffffffffffffffffff166000526008602052604060002090565b54158015613e82575b15613e195750505073ffffffffffffffffffffffffffffffffffffffff6000915b16600052601060205260406000205490613600565b90565b613e6c612710926137d673ffffffffffffffffffffffffffffffffffffffff9560026040613e7b96898b168152600360205220015460a01c1662f099c0613e65600f54600d5490613761565b0492613754565b613e7585613eb6565b90613761565b0491613e01565b5073ffffffffffffffffffffffffffffffffffffffff84168152600360205282600260408320015460a01c16821115613de0565b613f3e62278d00916000613eea8273ffffffffffffffffffffffffffffffffffffffff166000526008602052604060002090565b54151580613fec575b15613f6f575073ffffffffffffffffffffffffffffffffffffffff636648cc0b915b166000526011602052613f3965ffffffffffff60406000205460501c164290613600565b613754565b04600c8110613f6a5750600c5b6101f4908082029182040361360d5761271090810180911161360d5790565b613f4b565b613f998273ffffffffffffffffffffffffffffffffffffffff166000526008602052604060002090565b54613fbb575073ffffffffffffffffffffffffffffffffffffffff4291613f15565b64ffffffffff604073ffffffffffffffffffffffffffffffffffffffff928385168152601160205220541691613f15565b5073ffffffffffffffffffffffffffffffffffffffff82168152601160205264ffffffffff60408220541615613ef3565b60405190600754808352826020918282019060076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688936000905b82821061407257505050614070925003836136a3565b565b855473ffffffffffffffffffffffffffffffffffffffff168452600195860195889550938101939091019061405a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461360d5760010190565b80518210156135185760209160051b010190565b67ffffffffffffffff81116136585760051b60200190565b60005460405160208101918252426040820152434060608201526060815261412281613687565b5190208060005590565b919091600181019073ffffffffffffffffffffffffffffffffffffffff8416916000918383528160205260408320546141c2578054680100000000000000008110156141955760409495966110ff82600161418a940185558461357e565b549382526020522055565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b602484604051907f05e3de8f0000000000000000000000000000000000000000000000000000000082526004820152fd5b73ffffffffffffffffffffffffffffffffffffffff8091166000918183526008602052604083205480156142d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9161426d6142686142568560075401613547565b9190549186860192856106d485613547565b613547565b90549060031b1c1684526008602052604084205560075480156142a8570161429761073882613547565b600755815260086020526040812055565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b602483604051907f28a32c190000000000000000000000000000000000000000000000000000000082526004820152fd5b600181019073ffffffffffffffffffffffffffffffffffffffff8093169060009382855283602052604085205480156143eb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9161438761438161436e858754018761357e565b9190549186860192856106d4858b61357e565b8561357e565b90549060031b1c168652846020526040862055815480156143be5701906143b1610738838361357e565b5582526020526040812055565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b602484604051907f28a32c190000000000000000000000000000000000000000000000000000000082526004820152fd5b3d15614475573d9067ffffffffffffffff8211613658576040519161446960207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601846136a3565b82523d6000602084013e565b606090565b9081602091031261329a5751801515810361329a5790565b600092918380937f7472616e7366657228616464726573732c75696e74323536290000000000000060206040516144c88161363c565b6019815201526040519073ffffffffffffffffffffffffffffffffffffffff60208301947fa9059cbb00000000000000000000000000000000000000000000000000000000865216602483015260448201526044815261452781613687565b51925af161453361441c565b901590811561456e575b5061454457565b60046040517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b8051801515925082614583575b50503861453d565b614596925060208091830101910161447a565b15388061457b565b60408051909467ffffffffffffffff9493916060810186811182821017613658577fffffffff000000000000000000000000000000000000000000000000000000009160259189528181527f74323536290000000000000000000000000000000000000000000000000000008960208301927f7472616e7366657246726f6d28616464726573732c616464726573732c75696e845201522016918651946020860193845273ffffffffffffffffffffffffffffffffffffffff809216602487015216604485015260648401526064835260a0830193838510908511176136585760008094938194875251925af161469361441c565b90159081156146ce575b506146a55750565b600490517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b80518015159250826146e3575b50503861469d565b6146f6925060208091830101910161447a565b1538806146db56fea2646970667358221220a4852957b4c8f95657a1f947f314d11b9678858460fd62781bf130c7c588b8c264736f6c63430008140033",
}

// StakingHubABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingHubMetaData.ABI instead.
var StakingHubABI = StakingHubMetaData.ABI

// StakingHubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingHubMetaData.Bin instead.
var StakingHubBin = StakingHubMetaData.Bin

// DeployStakingHub deploys a new Ethereum contract, binding an instance of StakingHub to it.
func DeployStakingHub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingHub, error) {
	parsed, err := StakingHubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingHubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingHub{StakingHubCaller: StakingHubCaller{contract: contract}, StakingHubTransactor: StakingHubTransactor{contract: contract}, StakingHubFilterer: StakingHubFilterer{contract: contract}}, nil
}

// StakingHub is an auto generated Go binding around an Ethereum contract.
type StakingHub struct {
	StakingHubCaller     // Read-only binding to the contract
	StakingHubTransactor // Write-only binding to the contract
	StakingHubFilterer   // Log filterer for contract events
}

// StakingHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingHubSession struct {
	Contract     *StakingHub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingHubCallerSession struct {
	Contract *StakingHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakingHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingHubTransactorSession struct {
	Contract     *StakingHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakingHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingHubRaw struct {
	Contract *StakingHub // Generic contract binding to access the raw methods on
}

// StakingHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingHubCallerRaw struct {
	Contract *StakingHubCaller // Generic read-only contract binding to access the raw methods on
}

// StakingHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingHubTransactorRaw struct {
	Contract *StakingHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingHub creates a new instance of StakingHub, bound to a specific deployed contract.
func NewStakingHub(address common.Address, backend bind.ContractBackend) (*StakingHub, error) {
	contract, err := bindStakingHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingHub{StakingHubCaller: StakingHubCaller{contract: contract}, StakingHubTransactor: StakingHubTransactor{contract: contract}, StakingHubFilterer: StakingHubFilterer{contract: contract}}, nil
}

// NewStakingHubCaller creates a new read-only instance of StakingHub, bound to a specific deployed contract.
func NewStakingHubCaller(address common.Address, caller bind.ContractCaller) (*StakingHubCaller, error) {
	contract, err := bindStakingHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingHubCaller{contract: contract}, nil
}

// NewStakingHubTransactor creates a new write-only instance of StakingHub, bound to a specific deployed contract.
func NewStakingHubTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingHubTransactor, error) {
	contract, err := bindStakingHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingHubTransactor{contract: contract}, nil
}

// NewStakingHubFilterer creates a new log filterer instance of StakingHub, bound to a specific deployed contract.
func NewStakingHubFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingHubFilterer, error) {
	contract, err := bindStakingHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingHubFilterer{contract: contract}, nil
}

// bindStakingHub binds a generic wrapper to an already deployed contract.
func bindStakingHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingHub *StakingHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingHub.Contract.StakingHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingHub *StakingHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.Contract.StakingHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingHub *StakingHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingHub.Contract.StakingHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingHub *StakingHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingHub *StakingHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingHub *StakingHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingHub.Contract.contract.Transact(opts, method, params...)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_StakingHub *StakingHubCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_StakingHub *StakingHubSession) BlocksPerEpoch() (*big.Int, error) {
	return _StakingHub.Contract.BlocksPerEpoch(&_StakingHub.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_StakingHub *StakingHubCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _StakingHub.Contract.BlocksPerEpoch(&_StakingHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_StakingHub *StakingHubCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_StakingHub *StakingHubSession) CurrentEpoch() (*big.Int, error) {
	return _StakingHub.Contract.CurrentEpoch(&_StakingHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_StakingHub *StakingHubCallerSession) CurrentEpoch() (*big.Int, error) {
	return _StakingHub.Contract.CurrentEpoch(&_StakingHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_StakingHub *StakingHubCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_StakingHub *StakingHubSession) FinePercentage() (uint16, error) {
	return _StakingHub.Contract.FinePercentage(&_StakingHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_StakingHub *StakingHubCallerSession) FinePercentage() (uint16, error) {
	return _StakingHub.Contract.FinePercentage(&_StakingHub.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_StakingHub *StakingHubCaller) GetAllMinerUnstakeRequests(opts *bind.CallOpts) (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IStakingHubUnstakeRequest
}, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getAllMinerUnstakeRequests")

	outstruct := new(struct {
		UnstakeAddresses []common.Address
		UnstakeRequests  []IStakingHubUnstakeRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnstakeAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UnstakeRequests = *abi.ConvertType(out[1], new([]IStakingHubUnstakeRequest)).(*[]IStakingHubUnstakeRequest)

	return *outstruct, err

}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_StakingHub *StakingHubSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IStakingHubUnstakeRequest
}, error) {
	return _StakingHub.Contract.GetAllMinerUnstakeRequests(&_StakingHub.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_StakingHub *StakingHubCallerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IStakingHubUnstakeRequest
}, error) {
	return _StakingHub.Contract.GetAllMinerUnstakeRequests(&_StakingHub.CallOpts)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_StakingHub *StakingHubCaller) GetMinFeeToUse(opts *bind.CallOpts, _modelAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getMinFeeToUse", _modelAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_StakingHub *StakingHubSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _StakingHub.Contract.GetMinFeeToUse(&_StakingHub.CallOpts, _modelAddress)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_StakingHub *StakingHubCallerSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _StakingHub.Contract.GetMinFeeToUse(&_StakingHub.CallOpts, _modelAddress)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_StakingHub *StakingHubCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_StakingHub *StakingHubSession) GetMinerAddresses() ([]common.Address, error) {
	return _StakingHub.Contract.GetMinerAddresses(&_StakingHub.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_StakingHub *StakingHubCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _StakingHub.Contract.GetMinerAddresses(&_StakingHub.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_StakingHub *StakingHubCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getMinerAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_StakingHub *StakingHubSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _StakingHub.Contract.GetMinerAddressesOfModel(&_StakingHub.CallOpts, _model)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_StakingHub *StakingHubCallerSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _StakingHub.Contract.GetMinerAddressesOfModel(&_StakingHub.CallOpts, _model)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_StakingHub *StakingHubCaller) GetModelAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getModelAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_StakingHub *StakingHubSession) GetModelAddresses() ([]common.Address, error) {
	return _StakingHub.Contract.GetModelAddresses(&_StakingHub.CallOpts)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_StakingHub *StakingHubCallerSession) GetModelAddresses() ([]common.Address, error) {
	return _StakingHub.Contract.GetModelAddresses(&_StakingHub.CallOpts)
}

// GetModelInfo is a free data retrieval call binding the contract method 0xcd23ea14.
//
// Solidity: function getModelInfo(address _modelAddr) view returns((uint256,uint32))
func (_StakingHub *StakingHubCaller) GetModelInfo(opts *bind.CallOpts, _modelAddr common.Address) (IStakingHubModel, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getModelInfo", _modelAddr)

	if err != nil {
		return *new(IStakingHubModel), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingHubModel)).(*IStakingHubModel)

	return out0, err

}

// GetModelInfo is a free data retrieval call binding the contract method 0xcd23ea14.
//
// Solidity: function getModelInfo(address _modelAddr) view returns((uint256,uint32))
func (_StakingHub *StakingHubSession) GetModelInfo(_modelAddr common.Address) (IStakingHubModel, error) {
	return _StakingHub.Contract.GetModelInfo(&_StakingHub.CallOpts, _modelAddr)
}

// GetModelInfo is a free data retrieval call binding the contract method 0xcd23ea14.
//
// Solidity: function getModelInfo(address _modelAddr) view returns((uint256,uint32))
func (_StakingHub *StakingHubCallerSession) GetModelInfo(_modelAddr common.Address) (IStakingHubModel, error) {
	return _StakingHub.Contract.GetModelInfo(&_StakingHub.CallOpts, _modelAddr)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_StakingHub *StakingHubCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_StakingHub *StakingHubSession) GetNOMiner() (*big.Int, error) {
	return _StakingHub.Contract.GetNOMiner(&_StakingHub.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_StakingHub *StakingHubCallerSession) GetNOMiner() (*big.Int, error) {
	return _StakingHub.Contract.GetNOMiner(&_StakingHub.CallOpts)
}

// IsMinerAddress is a free data retrieval call binding the contract method 0x34875ec3.
//
// Solidity: function isMinerAddress(address _miner) view returns(bool)
func (_StakingHub *StakingHubCaller) IsMinerAddress(opts *bind.CallOpts, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "isMinerAddress", _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinerAddress is a free data retrieval call binding the contract method 0x34875ec3.
//
// Solidity: function isMinerAddress(address _miner) view returns(bool)
func (_StakingHub *StakingHubSession) IsMinerAddress(_miner common.Address) (bool, error) {
	return _StakingHub.Contract.IsMinerAddress(&_StakingHub.CallOpts, _miner)
}

// IsMinerAddress is a free data retrieval call binding the contract method 0x34875ec3.
//
// Solidity: function isMinerAddress(address _miner) view returns(bool)
func (_StakingHub *StakingHubCallerSession) IsMinerAddress(_miner common.Address) (bool, error) {
	return _StakingHub.Contract.IsMinerAddress(&_StakingHub.CallOpts, _miner)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_StakingHub *StakingHubCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_StakingHub *StakingHubSession) LastBlock() (*big.Int, error) {
	return _StakingHub.Contract.LastBlock(&_StakingHub.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_StakingHub *StakingHubCallerSession) LastBlock() (*big.Int, error) {
	return _StakingHub.Contract.LastBlock(&_StakingHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_StakingHub *StakingHubCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_StakingHub *StakingHubSession) MaximumTier() (uint16, error) {
	return _StakingHub.Contract.MaximumTier(&_StakingHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_StakingHub *StakingHubCallerSession) MaximumTier() (uint16, error) {
	return _StakingHub.Contract.MaximumTier(&_StakingHub.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_StakingHub *StakingHubCaller) MinFeeToUse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "minFeeToUse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_StakingHub *StakingHubSession) MinFeeToUse() (*big.Int, error) {
	return _StakingHub.Contract.MinFeeToUse(&_StakingHub.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_StakingHub *StakingHubCallerSession) MinFeeToUse() (*big.Int, error) {
	return _StakingHub.Contract.MinFeeToUse(&_StakingHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_StakingHub *StakingHubCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_StakingHub *StakingHubSession) MinerMinimumStake() (*big.Int, error) {
	return _StakingHub.Contract.MinerMinimumStake(&_StakingHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_StakingHub *StakingHubCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _StakingHub.Contract.MinerMinimumStake(&_StakingHub.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_StakingHub *StakingHubCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "minerUnstakeRequests", arg0)

	outstruct := new(struct {
		Stake    *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_StakingHub *StakingHubSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _StakingHub.Contract.MinerUnstakeRequests(&_StakingHub.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_StakingHub *StakingHubCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _StakingHub.Contract.MinerUnstakeRequests(&_StakingHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_StakingHub *StakingHubCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "miners", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		Commitment       *big.Int
		ModelAddress     common.Address
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModelAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_StakingHub *StakingHubSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _StakingHub.Contract.Miners(&_StakingHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_StakingHub *StakingHubCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _StakingHub.Contract.Miners(&_StakingHub.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_StakingHub *StakingHubCaller) Models(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "models", arg0)

	outstruct := new(struct {
		MinimumFee *big.Int
		Tier       uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_StakingHub *StakingHubSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _StakingHub.Contract.Models(&_StakingHub.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_StakingHub *StakingHubCallerSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _StakingHub.Contract.Models(&_StakingHub.CallOpts, arg0)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_StakingHub *StakingHubCaller) Multiplier(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "multiplier", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_StakingHub *StakingHubSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _StakingHub.Contract.Multiplier(&_StakingHub.CallOpts, _miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_StakingHub *StakingHubCallerSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _StakingHub.Contract.Multiplier(&_StakingHub.CallOpts, _miner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingHub *StakingHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingHub *StakingHubSession) Owner() (common.Address, error) {
	return _StakingHub.Contract.Owner(&_StakingHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingHub *StakingHubCallerSession) Owner() (common.Address, error) {
	return _StakingHub.Contract.Owner(&_StakingHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingHub *StakingHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingHub *StakingHubSession) Paused() (bool, error) {
	return _StakingHub.Contract.Paused(&_StakingHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingHub *StakingHubCallerSession) Paused() (bool, error) {
	return _StakingHub.Contract.Paused(&_StakingHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_StakingHub *StakingHubCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_StakingHub *StakingHubSession) PenaltyDuration() (*big.Int, error) {
	return _StakingHub.Contract.PenaltyDuration(&_StakingHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_StakingHub *StakingHubCallerSession) PenaltyDuration() (*big.Int, error) {
	return _StakingHub.Contract.PenaltyDuration(&_StakingHub.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_StakingHub *StakingHubCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "rewardInEpoch", arg0)

	outstruct := new(struct {
		PerfReward         *big.Int
		EpochReward        *big.Int
		TotalTaskCompleted *big.Int
		TotalMiner         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerfReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalTaskCompleted = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalMiner = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_StakingHub *StakingHubSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _StakingHub.Contract.RewardInEpoch(&_StakingHub.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_StakingHub *StakingHubCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _StakingHub.Contract.RewardInEpoch(&_StakingHub.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_StakingHub *StakingHubCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_StakingHub *StakingHubSession) RewardPerEpoch() (*big.Int, error) {
	return _StakingHub.Contract.RewardPerEpoch(&_StakingHub.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_StakingHub *StakingHubCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _StakingHub.Contract.RewardPerEpoch(&_StakingHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StakingHub *StakingHubCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StakingHub *StakingHubSession) Treasury() (common.Address, error) {
	return _StakingHub.Contract.Treasury(&_StakingHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StakingHub *StakingHubCallerSession) Treasury() (common.Address, error) {
	return _StakingHub.Contract.Treasury(&_StakingHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_StakingHub *StakingHubCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_StakingHub *StakingHubSession) UnstakeDelayTime() (*big.Int, error) {
	return _StakingHub.Contract.UnstakeDelayTime(&_StakingHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_StakingHub *StakingHubCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _StakingHub.Contract.UnstakeDelayTime(&_StakingHub.CallOpts)
}

// ValidateModelOfMiner is a free data retrieval call binding the contract method 0xd8f0166c.
//
// Solidity: function validateModelOfMiner(address _miner) view returns()
func (_StakingHub *StakingHubCaller) ValidateModelOfMiner(opts *bind.CallOpts, _miner common.Address) error {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "validateModelOfMiner", _miner)

	if err != nil {
		return err
	}

	return err

}

// ValidateModelOfMiner is a free data retrieval call binding the contract method 0xd8f0166c.
//
// Solidity: function validateModelOfMiner(address _miner) view returns()
func (_StakingHub *StakingHubSession) ValidateModelOfMiner(_miner common.Address) error {
	return _StakingHub.Contract.ValidateModelOfMiner(&_StakingHub.CallOpts, _miner)
}

// ValidateModelOfMiner is a free data retrieval call binding the contract method 0xd8f0166c.
//
// Solidity: function validateModelOfMiner(address _miner) view returns()
func (_StakingHub *StakingHubCallerSession) ValidateModelOfMiner(_miner common.Address) error {
	return _StakingHub.Contract.ValidateModelOfMiner(&_StakingHub.CallOpts, _miner)
}

// WEAI is a free data retrieval call binding the contract method 0x0dc7df53.
//
// Solidity: function wEAI() view returns(address)
func (_StakingHub *StakingHubCaller) WEAI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingHub.contract.Call(opts, &out, "wEAI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAI is a free data retrieval call binding the contract method 0x0dc7df53.
//
// Solidity: function wEAI() view returns(address)
func (_StakingHub *StakingHubSession) WEAI() (common.Address, error) {
	return _StakingHub.Contract.WEAI(&_StakingHub.CallOpts)
}

// WEAI is a free data retrieval call binding the contract method 0x0dc7df53.
//
// Solidity: function wEAI() view returns(address)
func (_StakingHub *StakingHubCallerSession) WEAI() (common.Address, error) {
	return _StakingHub.Contract.WEAI(&_StakingHub.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_StakingHub *StakingHubTransactor) ClaimReward(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "claimReward", _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_StakingHub *StakingHubSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.ClaimReward(&_StakingHub.TransactOpts, _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_StakingHub *StakingHubTransactorSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.ClaimReward(&_StakingHub.TransactOpts, _miner)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_StakingHub *StakingHubTransactor) ForceChangeModelForMiner(opts *bind.TransactOpts, _miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "forceChangeModelForMiner", _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_StakingHub *StakingHubSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.ForceChangeModelForMiner(&_StakingHub.TransactOpts, _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_StakingHub *StakingHubTransactorSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.ForceChangeModelForMiner(&_StakingHub.TransactOpts, _miner, _modelAddress)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_StakingHub *StakingHubTransactor) IncreaseMinerStake(opts *bind.TransactOpts, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "increaseMinerStake", wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_StakingHub *StakingHubSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.IncreaseMinerStake(&_StakingHub.TransactOpts, wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_StakingHub *StakingHubTransactorSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.IncreaseMinerStake(&_StakingHub.TransactOpts, wEAIAmt)
}

// Initialize is a paid mutator transaction binding the contract method 0x61ea0a25.
//
// Solidity: function initialize(address _wEAI, uint256 _minerMinimumStake, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage, uint256 _minFeeToUse) returns()
func (_StakingHub *StakingHubTransactor) Initialize(opts *bind.TransactOpts, _wEAI common.Address, _minerMinimumStake *big.Int, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16, _minFeeToUse *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "initialize", _wEAI, _minerMinimumStake, _blocksPerEpoch, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage, _minFeeToUse)
}

// Initialize is a paid mutator transaction binding the contract method 0x61ea0a25.
//
// Solidity: function initialize(address _wEAI, uint256 _minerMinimumStake, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage, uint256 _minFeeToUse) returns()
func (_StakingHub *StakingHubSession) Initialize(_wEAI common.Address, _minerMinimumStake *big.Int, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16, _minFeeToUse *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.Initialize(&_StakingHub.TransactOpts, _wEAI, _minerMinimumStake, _blocksPerEpoch, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage, _minFeeToUse)
}

// Initialize is a paid mutator transaction binding the contract method 0x61ea0a25.
//
// Solidity: function initialize(address _wEAI, uint256 _minerMinimumStake, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage, uint256 _minFeeToUse) returns()
func (_StakingHub *StakingHubTransactorSession) Initialize(_wEAI common.Address, _minerMinimumStake *big.Int, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16, _minFeeToUse *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.Initialize(&_StakingHub.TransactOpts, _wEAI, _minerMinimumStake, _blocksPerEpoch, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage, _minFeeToUse)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_StakingHub *StakingHubTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_StakingHub *StakingHubSession) JoinForMinting() (*types.Transaction, error) {
	return _StakingHub.Contract.JoinForMinting(&_StakingHub.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_StakingHub *StakingHubTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _StakingHub.Contract.JoinForMinting(&_StakingHub.TransactOpts)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_StakingHub *StakingHubTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "registerMiner", tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_StakingHub *StakingHubSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _StakingHub.Contract.RegisterMiner(&_StakingHub.TransactOpts, tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_StakingHub *StakingHubTransactorSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _StakingHub.Contract.RegisterMiner(&_StakingHub.TransactOpts, tier)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0xc325dc63.
//
// Solidity: function registerMiner(uint16 tier, address model) returns()
func (_StakingHub *StakingHubTransactor) RegisterMiner0(opts *bind.TransactOpts, tier uint16, model common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "registerMiner0", tier, model)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0xc325dc63.
//
// Solidity: function registerMiner(uint16 tier, address model) returns()
func (_StakingHub *StakingHubSession) RegisterMiner0(tier uint16, model common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.RegisterMiner0(&_StakingHub.TransactOpts, tier, model)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0xc325dc63.
//
// Solidity: function registerMiner(uint16 tier, address model) returns()
func (_StakingHub *StakingHubTransactorSession) RegisterMiner0(tier uint16, model common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.RegisterMiner0(&_StakingHub.TransactOpts, tier, model)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_StakingHub *StakingHubTransactor) RegisterModel(opts *bind.TransactOpts, _model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "registerModel", _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_StakingHub *StakingHubSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.RegisterModel(&_StakingHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_StakingHub *StakingHubTransactorSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.RegisterModel(&_StakingHub.TransactOpts, _model, _tier, _minimumFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingHub *StakingHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingHub *StakingHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingHub.Contract.RenounceOwnership(&_StakingHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingHub *StakingHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingHub.Contract.RenounceOwnership(&_StakingHub.TransactOpts)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_StakingHub *StakingHubTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_StakingHub *StakingHubSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _StakingHub.Contract.RestakeForMiner(&_StakingHub.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_StakingHub *StakingHubTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _StakingHub.Contract.RestakeForMiner(&_StakingHub.TransactOpts, tier)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_StakingHub *StakingHubTransactor) RewardToClaim(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "rewardToClaim", _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_StakingHub *StakingHubSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.RewardToClaim(&_StakingHub.TransactOpts, _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_StakingHub *StakingHubTransactorSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.RewardToClaim(&_StakingHub.TransactOpts, _miner)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_StakingHub *StakingHubTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setBlocksPerEpoch", _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_StakingHub *StakingHubSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetBlocksPerEpoch(&_StakingHub.TransactOpts, _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_StakingHub *StakingHubTransactorSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetBlocksPerEpoch(&_StakingHub.TransactOpts, _blocks)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_StakingHub *StakingHubTransactor) SetFinePercentage(opts *bind.TransactOpts, _finePercentage uint16) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setFinePercentage", _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_StakingHub *StakingHubSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _StakingHub.Contract.SetFinePercentage(&_StakingHub.TransactOpts, _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_StakingHub *StakingHubTransactorSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _StakingHub.Contract.SetFinePercentage(&_StakingHub.TransactOpts, _finePercentage)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_StakingHub *StakingHubTransactor) SetMinFeeToUse(opts *bind.TransactOpts, _minFeeToUse *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setMinFeeToUse", _minFeeToUse)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_StakingHub *StakingHubSession) SetMinFeeToUse(_minFeeToUse *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetMinFeeToUse(&_StakingHub.TransactOpts, _minFeeToUse)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_StakingHub *StakingHubTransactorSession) SetMinFeeToUse(_minFeeToUse *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetMinFeeToUse(&_StakingHub.TransactOpts, _minFeeToUse)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_StakingHub *StakingHubTransactor) SetMinerMinimumStake(opts *bind.TransactOpts, _minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setMinerMinimumStake", _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_StakingHub *StakingHubSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetMinerMinimumStake(&_StakingHub.TransactOpts, _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_StakingHub *StakingHubTransactorSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetMinerMinimumStake(&_StakingHub.TransactOpts, _minerMinimumStake)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_StakingHub *StakingHubTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, _newRewardAmount *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setNewRewardInEpoch", _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_StakingHub *StakingHubSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetNewRewardInEpoch(&_StakingHub.TransactOpts, _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_StakingHub *StakingHubTransactorSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetNewRewardInEpoch(&_StakingHub.TransactOpts, _newRewardAmount)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_StakingHub *StakingHubTransactor) SetPenaltyDuration(opts *bind.TransactOpts, _penaltyDuration *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setPenaltyDuration", _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_StakingHub *StakingHubSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetPenaltyDuration(&_StakingHub.TransactOpts, _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_StakingHub *StakingHubTransactorSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetPenaltyDuration(&_StakingHub.TransactOpts, _penaltyDuration)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_StakingHub *StakingHubTransactor) SetUnstakDelayTime(opts *bind.TransactOpts, _newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setUnstakDelayTime", _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_StakingHub *StakingHubSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetUnstakDelayTime(&_StakingHub.TransactOpts, _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_StakingHub *StakingHubTransactorSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.SetUnstakDelayTime(&_StakingHub.TransactOpts, _newUnstakeDelayTime)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address _wEAI) returns()
func (_StakingHub *StakingHubTransactor) SetWEAIAddress(opts *bind.TransactOpts, _wEAI common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setWEAIAddress", _wEAI)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address _wEAI) returns()
func (_StakingHub *StakingHubSession) SetWEAIAddress(_wEAI common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.SetWEAIAddress(&_StakingHub.TransactOpts, _wEAI)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address _wEAI) returns()
func (_StakingHub *StakingHubTransactorSession) SetWEAIAddress(_wEAI common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.SetWEAIAddress(&_StakingHub.TransactOpts, _wEAI)
}

// SetWorkerHubAddress is a paid mutator transaction binding the contract method 0xf6a6ae1d.
//
// Solidity: function setWorkerHubAddress(address _workerHub) returns()
func (_StakingHub *StakingHubTransactor) SetWorkerHubAddress(opts *bind.TransactOpts, _workerHub common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "setWorkerHubAddress", _workerHub)
}

// SetWorkerHubAddress is a paid mutator transaction binding the contract method 0xf6a6ae1d.
//
// Solidity: function setWorkerHubAddress(address _workerHub) returns()
func (_StakingHub *StakingHubSession) SetWorkerHubAddress(_workerHub common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.SetWorkerHubAddress(&_StakingHub.TransactOpts, _workerHub)
}

// SetWorkerHubAddress is a paid mutator transaction binding the contract method 0xf6a6ae1d.
//
// Solidity: function setWorkerHubAddress(address _workerHub) returns()
func (_StakingHub *StakingHubTransactorSession) SetWorkerHubAddress(_workerHub common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.SetWorkerHubAddress(&_StakingHub.TransactOpts, _workerHub)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_StakingHub *StakingHubTransactor) SlashMiner(opts *bind.TransactOpts, _miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "slashMiner", _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_StakingHub *StakingHubSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _StakingHub.Contract.SlashMiner(&_StakingHub.TransactOpts, _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_StakingHub *StakingHubTransactorSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _StakingHub.Contract.SlashMiner(&_StakingHub.TransactOpts, _miner, _isFined)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingHub *StakingHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingHub *StakingHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.TransferOwnership(&_StakingHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingHub *StakingHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.TransferOwnership(&_StakingHub.TransactOpts, newOwner)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_StakingHub *StakingHubTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_StakingHub *StakingHubSession) UnregisterMiner() (*types.Transaction, error) {
	return _StakingHub.Contract.UnregisterMiner(&_StakingHub.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_StakingHub *StakingHubTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _StakingHub.Contract.UnregisterMiner(&_StakingHub.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_StakingHub *StakingHubTransactor) UnregisterModel(opts *bind.TransactOpts, _model common.Address) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "unregisterModel", _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_StakingHub *StakingHubSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.UnregisterModel(&_StakingHub.TransactOpts, _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_StakingHub *StakingHubTransactorSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _StakingHub.Contract.UnregisterModel(&_StakingHub.TransactOpts, _model)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_StakingHub *StakingHubTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_StakingHub *StakingHubSession) UnstakeForMiner() (*types.Transaction, error) {
	return _StakingHub.Contract.UnstakeForMiner(&_StakingHub.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_StakingHub *StakingHubTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _StakingHub.Contract.UnstakeForMiner(&_StakingHub.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_StakingHub *StakingHubTransactor) UpdateEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "updateEpoch")
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_StakingHub *StakingHubSession) UpdateEpoch() (*types.Transaction, error) {
	return _StakingHub.Contract.UpdateEpoch(&_StakingHub.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_StakingHub *StakingHubTransactorSession) UpdateEpoch() (*types.Transaction, error) {
	return _StakingHub.Contract.UpdateEpoch(&_StakingHub.TransactOpts)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_StakingHub *StakingHubTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, _model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "updateModelMinimumFee", _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_StakingHub *StakingHubSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.UpdateModelMinimumFee(&_StakingHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_StakingHub *StakingHubTransactorSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _StakingHub.Contract.UpdateModelMinimumFee(&_StakingHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_StakingHub *StakingHubTransactor) UpdateModelTier(opts *bind.TransactOpts, _model common.Address, _tier uint32) (*types.Transaction, error) {
	return _StakingHub.contract.Transact(opts, "updateModelTier", _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_StakingHub *StakingHubSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _StakingHub.Contract.UpdateModelTier(&_StakingHub.TransactOpts, _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_StakingHub *StakingHubTransactorSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _StakingHub.Contract.UpdateModelTier(&_StakingHub.TransactOpts, _model, _tier)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingHub *StakingHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingHub *StakingHubSession) Receive() (*types.Transaction, error) {
	return _StakingHub.Contract.Receive(&_StakingHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingHub *StakingHubTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingHub.Contract.Receive(&_StakingHub.TransactOpts)
}

// StakingHubBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the StakingHub contract.
type StakingHubBlocksPerEpochIterator struct {
	Event *StakingHubBlocksPerEpoch // Event containing the contract specifics and raw log

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
func (it *StakingHubBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubBlocksPerEpoch)
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
		it.Event = new(StakingHubBlocksPerEpoch)
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
func (it *StakingHubBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubBlocksPerEpoch represents a BlocksPerEpoch event raised by the StakingHub contract.
type StakingHubBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_StakingHub *StakingHubFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*StakingHubBlocksPerEpochIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &StakingHubBlocksPerEpochIterator{contract: _StakingHub.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_StakingHub *StakingHubFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *StakingHubBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubBlocksPerEpoch)
				if err := _StakingHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
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

// ParseBlocksPerEpoch is a log parse operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_StakingHub *StakingHubFilterer) ParseBlocksPerEpoch(log types.Log) (*StakingHubBlocksPerEpoch, error) {
	event := new(StakingHubBlocksPerEpoch)
	if err := _StakingHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the StakingHub contract.
type StakingHubFinePercentageUpdatedIterator struct {
	Event *StakingHubFinePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *StakingHubFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubFinePercentageUpdated)
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
		it.Event = new(StakingHubFinePercentageUpdated)
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
func (it *StakingHubFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubFinePercentageUpdated represents a FinePercentageUpdated event raised by the StakingHub contract.
type StakingHubFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_StakingHub *StakingHubFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*StakingHubFinePercentageUpdatedIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingHubFinePercentageUpdatedIterator{contract: _StakingHub.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_StakingHub *StakingHubFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *StakingHubFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubFinePercentageUpdated)
				if err := _StakingHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
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

// ParseFinePercentageUpdated is a log parse operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_StakingHub *StakingHubFilterer) ParseFinePercentageUpdated(log types.Log) (*StakingHubFinePercentageUpdated, error) {
	event := new(StakingHubFinePercentageUpdated)
	if err := _StakingHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the StakingHub contract.
type StakingHubFraudulentMinerPenalizedIterator struct {
	Event *StakingHubFraudulentMinerPenalized // Event containing the contract specifics and raw log

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
func (it *StakingHubFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubFraudulentMinerPenalized)
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
		it.Event = new(StakingHubFraudulentMinerPenalized)
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
func (it *StakingHubFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the StakingHub contract.
type StakingHubFraudulentMinerPenalized struct {
	Miner        common.Address
	ModelAddress common.Address
	Treasury     common.Address
	Fine         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_StakingHub *StakingHubFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (*StakingHubFraudulentMinerPenalizedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubFraudulentMinerPenalizedIterator{contract: _StakingHub.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_StakingHub *StakingHubFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *StakingHubFraudulentMinerPenalized, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubFraudulentMinerPenalized)
				if err := _StakingHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
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

// ParseFraudulentMinerPenalized is a log parse operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_StakingHub *StakingHubFilterer) ParseFraudulentMinerPenalized(log types.Log) (*StakingHubFraudulentMinerPenalized, error) {
	event := new(StakingHubFraudulentMinerPenalized)
	if err := _StakingHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingHub contract.
type StakingHubInitializedIterator struct {
	Event *StakingHubInitialized // Event containing the contract specifics and raw log

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
func (it *StakingHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubInitialized)
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
		it.Event = new(StakingHubInitialized)
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
func (it *StakingHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubInitialized represents a Initialized event raised by the StakingHub contract.
type StakingHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingHub *StakingHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingHubInitializedIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingHubInitializedIterator{contract: _StakingHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingHub *StakingHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingHubInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubInitialized)
				if err := _StakingHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingHub *StakingHubFilterer) ParseInitialized(log types.Log) (*StakingHubInitialized, error) {
	event := new(StakingHubInitialized)
	if err := _StakingHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinFeeToUseUpdatedIterator is returned from FilterMinFeeToUseUpdated and is used to iterate over the raw logs and unpacked data for MinFeeToUseUpdated events raised by the StakingHub contract.
type StakingHubMinFeeToUseUpdatedIterator struct {
	Event *StakingHubMinFeeToUseUpdated // Event containing the contract specifics and raw log

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
func (it *StakingHubMinFeeToUseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinFeeToUseUpdated)
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
		it.Event = new(StakingHubMinFeeToUseUpdated)
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
func (it *StakingHubMinFeeToUseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinFeeToUseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinFeeToUseUpdated represents a MinFeeToUseUpdated event raised by the StakingHub contract.
type StakingHubMinFeeToUseUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeToUseUpdated is a free log retrieval operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_StakingHub *StakingHubFilterer) FilterMinFeeToUseUpdated(opts *bind.FilterOpts) (*StakingHubMinFeeToUseUpdatedIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingHubMinFeeToUseUpdatedIterator{contract: _StakingHub.contract, event: "MinFeeToUseUpdated", logs: logs, sub: sub}, nil
}

// WatchMinFeeToUseUpdated is a free log subscription operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_StakingHub *StakingHubFilterer) WatchMinFeeToUseUpdated(opts *bind.WatchOpts, sink chan<- *StakingHubMinFeeToUseUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinFeeToUseUpdated)
				if err := _StakingHub.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
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

// ParseMinFeeToUseUpdated is a log parse operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_StakingHub *StakingHubFilterer) ParseMinFeeToUseUpdated(log types.Log) (*StakingHubMinFeeToUseUpdated, error) {
	event := new(StakingHubMinFeeToUseUpdated)
	if err := _StakingHub.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the StakingHub contract.
type StakingHubMinerDeactivatedIterator struct {
	Event *StakingHubMinerDeactivated // Event containing the contract specifics and raw log

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
func (it *StakingHubMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinerDeactivated)
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
		it.Event = new(StakingHubMinerDeactivated)
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
func (it *StakingHubMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinerDeactivated represents a MinerDeactivated event raised by the StakingHub contract.
type StakingHubMinerDeactivated struct {
	Miner        common.Address
	ModelAddress common.Address
	ActiveTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_StakingHub *StakingHubFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address) (*StakingHubMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubMinerDeactivatedIterator{contract: _StakingHub.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_StakingHub *StakingHubFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *StakingHubMinerDeactivated, miner []common.Address, modelAddress []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinerDeactivated)
				if err := _StakingHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
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

// ParseMinerDeactivated is a log parse operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_StakingHub *StakingHubFilterer) ParseMinerDeactivated(log types.Log) (*StakingHubMinerDeactivated, error) {
	event := new(StakingHubMinerDeactivated)
	if err := _StakingHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the StakingHub contract.
type StakingHubMinerExtraStakeIterator struct {
	Event *StakingHubMinerExtraStake // Event containing the contract specifics and raw log

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
func (it *StakingHubMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinerExtraStake)
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
		it.Event = new(StakingHubMinerExtraStake)
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
func (it *StakingHubMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinerExtraStake represents a MinerExtraStake event raised by the StakingHub contract.
type StakingHubMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_StakingHub *StakingHubFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*StakingHubMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubMinerExtraStakeIterator{contract: _StakingHub.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_StakingHub *StakingHubFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *StakingHubMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinerExtraStake)
				if err := _StakingHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
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

// ParseMinerExtraStake is a log parse operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_StakingHub *StakingHubFilterer) ParseMinerExtraStake(log types.Log) (*StakingHubMinerExtraStake, error) {
	event := new(StakingHubMinerExtraStake)
	if err := _StakingHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the StakingHub contract.
type StakingHubMinerJoinIterator struct {
	Event *StakingHubMinerJoin // Event containing the contract specifics and raw log

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
func (it *StakingHubMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinerJoin)
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
		it.Event = new(StakingHubMinerJoin)
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
func (it *StakingHubMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinerJoin represents a MinerJoin event raised by the StakingHub contract.
type StakingHubMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_StakingHub *StakingHubFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*StakingHubMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubMinerJoinIterator{contract: _StakingHub.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_StakingHub *StakingHubFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *StakingHubMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinerJoin)
				if err := _StakingHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
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

// ParseMinerJoin is a log parse operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_StakingHub *StakingHubFilterer) ParseMinerJoin(log types.Log) (*StakingHubMinerJoin, error) {
	event := new(StakingHubMinerJoin)
	if err := _StakingHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the StakingHub contract.
type StakingHubMinerRegistrationIterator struct {
	Event *StakingHubMinerRegistration // Event containing the contract specifics and raw log

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
func (it *StakingHubMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinerRegistration)
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
		it.Event = new(StakingHubMinerRegistration)
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
func (it *StakingHubMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinerRegistration represents a MinerRegistration event raised by the StakingHub contract.
type StakingHubMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_StakingHub *StakingHubFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*StakingHubMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubMinerRegistrationIterator{contract: _StakingHub.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_StakingHub *StakingHubFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *StakingHubMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinerRegistration)
				if err := _StakingHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
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

// ParseMinerRegistration is a log parse operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_StakingHub *StakingHubFilterer) ParseMinerRegistration(log types.Log) (*StakingHubMinerRegistration, error) {
	event := new(StakingHubMinerRegistration)
	if err := _StakingHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the StakingHub contract.
type StakingHubMinerUnregistrationIterator struct {
	Event *StakingHubMinerUnregistration // Event containing the contract specifics and raw log

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
func (it *StakingHubMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinerUnregistration)
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
		it.Event = new(StakingHubMinerUnregistration)
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
func (it *StakingHubMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinerUnregistration represents a MinerUnregistration event raised by the StakingHub contract.
type StakingHubMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_StakingHub *StakingHubFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*StakingHubMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubMinerUnregistrationIterator{contract: _StakingHub.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_StakingHub *StakingHubFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *StakingHubMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinerUnregistration)
				if err := _StakingHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
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

// ParseMinerUnregistration is a log parse operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_StakingHub *StakingHubFilterer) ParseMinerUnregistration(log types.Log) (*StakingHubMinerUnregistration, error) {
	event := new(StakingHubMinerUnregistration)
	if err := _StakingHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the StakingHub contract.
type StakingHubMinerUnstakeIterator struct {
	Event *StakingHubMinerUnstake // Event containing the contract specifics and raw log

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
func (it *StakingHubMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubMinerUnstake)
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
		it.Event = new(StakingHubMinerUnstake)
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
func (it *StakingHubMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubMinerUnstake represents a MinerUnstake event raised by the StakingHub contract.
type StakingHubMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_StakingHub *StakingHubFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*StakingHubMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubMinerUnstakeIterator{contract: _StakingHub.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_StakingHub *StakingHubFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *StakingHubMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubMinerUnstake)
				if err := _StakingHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
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

// ParseMinerUnstake is a log parse operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_StakingHub *StakingHubFilterer) ParseMinerUnstake(log types.Log) (*StakingHubMinerUnstake, error) {
	event := new(StakingHubMinerUnstake)
	if err := _StakingHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the StakingHub contract.
type StakingHubModelMinimumFeeUpdateIterator struct {
	Event *StakingHubModelMinimumFeeUpdate // Event containing the contract specifics and raw log

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
func (it *StakingHubModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubModelMinimumFeeUpdate)
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
		it.Event = new(StakingHubModelMinimumFeeUpdate)
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
func (it *StakingHubModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the StakingHub contract.
type StakingHubModelMinimumFeeUpdate struct {
	Model      common.Address
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_StakingHub *StakingHubFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, model []common.Address) (*StakingHubModelMinimumFeeUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubModelMinimumFeeUpdateIterator{contract: _StakingHub.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_StakingHub *StakingHubFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *StakingHubModelMinimumFeeUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubModelMinimumFeeUpdate)
				if err := _StakingHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
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

// ParseModelMinimumFeeUpdate is a log parse operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_StakingHub *StakingHubFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*StakingHubModelMinimumFeeUpdate, error) {
	event := new(StakingHubModelMinimumFeeUpdate)
	if err := _StakingHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the StakingHub contract.
type StakingHubModelRegistrationIterator struct {
	Event *StakingHubModelRegistration // Event containing the contract specifics and raw log

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
func (it *StakingHubModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubModelRegistration)
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
		it.Event = new(StakingHubModelRegistration)
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
func (it *StakingHubModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubModelRegistration represents a ModelRegistration event raised by the StakingHub contract.
type StakingHubModelRegistration struct {
	Model      common.Address
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_StakingHub *StakingHubFilterer) FilterModelRegistration(opts *bind.FilterOpts, model []common.Address, tier []uint16) (*StakingHubModelRegistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubModelRegistrationIterator{contract: _StakingHub.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_StakingHub *StakingHubFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *StakingHubModelRegistration, model []common.Address, tier []uint16) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubModelRegistration)
				if err := _StakingHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
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

// ParseModelRegistration is a log parse operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_StakingHub *StakingHubFilterer) ParseModelRegistration(log types.Log) (*StakingHubModelRegistration, error) {
	event := new(StakingHubModelRegistration)
	if err := _StakingHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the StakingHub contract.
type StakingHubModelTierUpdateIterator struct {
	Event *StakingHubModelTierUpdate // Event containing the contract specifics and raw log

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
func (it *StakingHubModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubModelTierUpdate)
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
		it.Event = new(StakingHubModelTierUpdate)
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
func (it *StakingHubModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubModelTierUpdate represents a ModelTierUpdate event raised by the StakingHub contract.
type StakingHubModelTierUpdate struct {
	Model common.Address
	Tier  uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_StakingHub *StakingHubFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, model []common.Address) (*StakingHubModelTierUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubModelTierUpdateIterator{contract: _StakingHub.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_StakingHub *StakingHubFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *StakingHubModelTierUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubModelTierUpdate)
				if err := _StakingHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
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

// ParseModelTierUpdate is a log parse operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_StakingHub *StakingHubFilterer) ParseModelTierUpdate(log types.Log) (*StakingHubModelTierUpdate, error) {
	event := new(StakingHubModelTierUpdate)
	if err := _StakingHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the StakingHub contract.
type StakingHubModelUnregistrationIterator struct {
	Event *StakingHubModelUnregistration // Event containing the contract specifics and raw log

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
func (it *StakingHubModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubModelUnregistration)
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
		it.Event = new(StakingHubModelUnregistration)
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
func (it *StakingHubModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubModelUnregistration represents a ModelUnregistration event raised by the StakingHub contract.
type StakingHubModelUnregistration struct {
	Model common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_StakingHub *StakingHubFilterer) FilterModelUnregistration(opts *bind.FilterOpts, model []common.Address) (*StakingHubModelUnregistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubModelUnregistrationIterator{contract: _StakingHub.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_StakingHub *StakingHubFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *StakingHubModelUnregistration, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubModelUnregistration)
				if err := _StakingHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
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

// ParseModelUnregistration is a log parse operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_StakingHub *StakingHubFilterer) ParseModelUnregistration(log types.Log) (*StakingHubModelUnregistration, error) {
	event := new(StakingHubModelUnregistration)
	if err := _StakingHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingHub contract.
type StakingHubOwnershipTransferredIterator struct {
	Event *StakingHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubOwnershipTransferred)
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
		it.Event = new(StakingHubOwnershipTransferred)
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
func (it *StakingHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubOwnershipTransferred represents a OwnershipTransferred event raised by the StakingHub contract.
type StakingHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingHub *StakingHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubOwnershipTransferredIterator{contract: _StakingHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingHub *StakingHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubOwnershipTransferred)
				if err := _StakingHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingHub *StakingHubFilterer) ParseOwnershipTransferred(log types.Log) (*StakingHubOwnershipTransferred, error) {
	event := new(StakingHubOwnershipTransferred)
	if err := _StakingHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingHub contract.
type StakingHubPausedIterator struct {
	Event *StakingHubPaused // Event containing the contract specifics and raw log

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
func (it *StakingHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubPaused)
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
		it.Event = new(StakingHubPaused)
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
func (it *StakingHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubPaused represents a Paused event raised by the StakingHub contract.
type StakingHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingHub *StakingHubFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingHubPausedIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingHubPausedIterator{contract: _StakingHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingHub *StakingHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingHubPaused) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubPaused)
				if err := _StakingHub.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingHub *StakingHubFilterer) ParsePaused(log types.Log) (*StakingHubPaused, error) {
	event := new(StakingHubPaused)
	if err := _StakingHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the StakingHub contract.
type StakingHubPenaltyDurationUpdatedIterator struct {
	Event *StakingHubPenaltyDurationUpdated // Event containing the contract specifics and raw log

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
func (it *StakingHubPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubPenaltyDurationUpdated)
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
		it.Event = new(StakingHubPenaltyDurationUpdated)
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
func (it *StakingHubPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the StakingHub contract.
type StakingHubPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_StakingHub *StakingHubFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*StakingHubPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingHubPenaltyDurationUpdatedIterator{contract: _StakingHub.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_StakingHub *StakingHubFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *StakingHubPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubPenaltyDurationUpdated)
				if err := _StakingHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
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

// ParsePenaltyDurationUpdated is a log parse operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_StakingHub *StakingHubFilterer) ParsePenaltyDurationUpdated(log types.Log) (*StakingHubPenaltyDurationUpdated, error) {
	event := new(StakingHubPenaltyDurationUpdated)
	if err := _StakingHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the StakingHub contract.
type StakingHubRestakeIterator struct {
	Event *StakingHubRestake // Event containing the contract specifics and raw log

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
func (it *StakingHubRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubRestake)
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
		it.Event = new(StakingHubRestake)
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
func (it *StakingHubRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubRestake represents a Restake event raised by the StakingHub contract.
type StakingHubRestake struct {
	Miner   common.Address
	Restake *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_StakingHub *StakingHubFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, model []common.Address) (*StakingHubRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubRestakeIterator{contract: _StakingHub.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_StakingHub *StakingHubFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *StakingHubRestake, miner []common.Address, model []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubRestake)
				if err := _StakingHub.contract.UnpackLog(event, "Restake", log); err != nil {
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

// ParseRestake is a log parse operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_StakingHub *StakingHubFilterer) ParseRestake(log types.Log) (*StakingHubRestake, error) {
	event := new(StakingHubRestake)
	if err := _StakingHub.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the StakingHub contract.
type StakingHubRewardClaimIterator struct {
	Event *StakingHubRewardClaim // Event containing the contract specifics and raw log

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
func (it *StakingHubRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubRewardClaim)
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
		it.Event = new(StakingHubRewardClaim)
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
func (it *StakingHubRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubRewardClaim represents a RewardClaim event raised by the StakingHub contract.
type StakingHubRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_StakingHub *StakingHubFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*StakingHubRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &StakingHubRewardClaimIterator{contract: _StakingHub.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_StakingHub *StakingHubFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *StakingHubRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubRewardClaim)
				if err := _StakingHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
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

// ParseRewardClaim is a log parse operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_StakingHub *StakingHubFilterer) ParseRewardClaim(log types.Log) (*StakingHubRewardClaim, error) {
	event := new(StakingHubRewardClaim)
	if err := _StakingHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the StakingHub contract.
type StakingHubRewardPerEpochIterator struct {
	Event *StakingHubRewardPerEpoch // Event containing the contract specifics and raw log

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
func (it *StakingHubRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubRewardPerEpoch)
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
		it.Event = new(StakingHubRewardPerEpoch)
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
func (it *StakingHubRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubRewardPerEpoch represents a RewardPerEpoch event raised by the StakingHub contract.
type StakingHubRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_StakingHub *StakingHubFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*StakingHubRewardPerEpochIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &StakingHubRewardPerEpochIterator{contract: _StakingHub.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_StakingHub *StakingHubFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *StakingHubRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubRewardPerEpoch)
				if err := _StakingHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
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

// ParseRewardPerEpoch is a log parse operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_StakingHub *StakingHubFilterer) ParseRewardPerEpoch(log types.Log) (*StakingHubRewardPerEpoch, error) {
	event := new(StakingHubRewardPerEpoch)
	if err := _StakingHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingHub contract.
type StakingHubUnpausedIterator struct {
	Event *StakingHubUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubUnpaused)
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
		it.Event = new(StakingHubUnpaused)
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
func (it *StakingHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubUnpaused represents a Unpaused event raised by the StakingHub contract.
type StakingHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingHub *StakingHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingHubUnpausedIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingHubUnpausedIterator{contract: _StakingHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingHub *StakingHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubUnpaused)
				if err := _StakingHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingHub *StakingHubFilterer) ParseUnpaused(log types.Log) (*StakingHubUnpaused, error) {
	event := new(StakingHubUnpaused)
	if err := _StakingHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingHubUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the StakingHub contract.
type StakingHubUnstakeDelayTimeIterator struct {
	Event *StakingHubUnstakeDelayTime // Event containing the contract specifics and raw log

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
func (it *StakingHubUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingHubUnstakeDelayTime)
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
		it.Event = new(StakingHubUnstakeDelayTime)
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
func (it *StakingHubUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingHubUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingHubUnstakeDelayTime represents a UnstakeDelayTime event raised by the StakingHub contract.
type StakingHubUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_StakingHub *StakingHubFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*StakingHubUnstakeDelayTimeIterator, error) {

	logs, sub, err := _StakingHub.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &StakingHubUnstakeDelayTimeIterator{contract: _StakingHub.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_StakingHub *StakingHubFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *StakingHubUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _StakingHub.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingHubUnstakeDelayTime)
				if err := _StakingHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
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

// ParseUnstakeDelayTime is a log parse operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_StakingHub *StakingHubFilterer) ParseUnstakeDelayTime(log types.Log) (*StakingHubUnstakeDelayTime, error) {
	event := new(StakingHubUnstakeDelayTime)
	if err := _StakingHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
