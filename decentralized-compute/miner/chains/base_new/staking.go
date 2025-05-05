package base_new

import (
	"context"
	"core/contracts/gpu_manager"
	"math/big"
	"solo/pkg/logger"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"

	"solo/internal/port"
	"solo/pkg/eth"

	"github.com/ethereum/go-ethereum/core/types"
)

type staking struct {
	stakingHub    *gpu_manager.GpuManager
	common        port.ICommon
	stakingHubAbi abi.ABI
}

func NewStaking(common port.ICommon, stakingHub *gpu_manager.GpuManager) port.IStaking {
	instanceABI, err := abi.JSON(strings.NewReader(gpu_manager.GpuManagerABI))
	if err != nil {
		return nil
	}

	return &staking{
		common:        common,
		stakingHub:    stakingHub,
		stakingHubAbi: instanceABI,
	}
}

func (s *staking) IsStaked() (bool, error) {
	workerInfo, err := s.stakingHub.Miners(nil, s.common.GetWalletAddres())
	if err != nil {
		return false, err
	}

	minStake, err := s.stakingHub.MinerMinimumStake(nil)
	if err != nil {
		return false, err
	}

	if workerInfo.Stake.Cmp(minStake) < 0 {
		return false, nil
	}

	pendingUnstake, err := s.stakingHub.MinerUnstakeRequests(nil, s.common.GetWalletAddres())
	if err != nil {
		return false, err
	}

	_ = pendingUnstake
	/*
		tskw.status.pendingUnstakeAmount = pendingUnstake.Stake
		if tskw.status.pendingUnstakeAmount.Cmp(new(big.Int).SetUint64(0)) > 0 {
			tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)
		}

		tskw.status.assignModel = workerInfo.ModelAddress.Hex()
		tskw.status.stakedAmount = workerInfo.Stake*/

	return true, nil
}

func (s *staking) StakeForWorker() (*types.Transaction, error) {
	ctx := context.Background()
	balance, err := s.common.GetErc20contract().BalanceOf(nil, s.common.GetWalletAddres())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return nil, err
	}

	err = eth.ApproveERC20(
		ctx, s.common.GetClient(), s.common.GetPrivateKey(), s.common.GetStakingHubAddress(),
		s.common.GetErc20contractAddress(), int64(s.common.GetGasLimit()),
	)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return nil, err
	}

	minStake, err := s.stakingHub.MinerMinimumStake(nil)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, s.common.GetClient(), s.common.GetPrivateKey(), int64(s.common.GetGasLimit()))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return nil, err
	}

	// auth.Value = minStake
	tx := new(types.Transaction)
	if s.common.GetConfig().ClusterID != "" {
		modelID, err1 := strconv.Atoi(s.common.GetConfig().ClusterID)
		if err1 == nil {
			tx, err = s.stakingHub.RegisterMiner0(auth, 1, uint32(modelID))
			if err != nil {

				logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
					zap.String("worker_address", s.common.GetWalletAddres().Hex()),
					zap.Int("model_id", modelID),
					zap.String("balance", balance.String()),
					zap.String("min_stake", minStake.String()),
					zap.Error(err))
				return nil, err
			}

			logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
				zap.String("worker_address", s.common.GetWalletAddres().Hex()),
				zap.Int("model_id", modelID),
				zap.String("balance", balance.String()),
				zap.String("min_stake", minStake.String()),
				zap.Any("tx", tx.Hash().Hex()),
			)
		}
	} else {
		tx, err = s.stakingHub.RegisterMiner(auth, 1)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
				zap.String("worker_address", s.common.GetWalletAddres().Hex()),
				zap.String("balance", balance.String()),
				zap.String("min_stake", minStake.String()),
				zap.Error(err))
			return nil, err
		}

		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.String("balance", balance.String()),
			zap.String("min_stake", minStake.String()),
			zap.Any("tx", tx.Hash().Hex()),
		)
	}

	// TODO - here
	_ = tx
	_ = balance
	_ = minStake
	return tx, nil
}

func (s *staking) JoinForMinting() (*types.Transaction, error) {
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, s.common.GetClient(), s.common.GetPrivateKey(), int64(s.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	tx, err := s.stakingHub.JoinForMinting(auth)
	if err != nil {
		return nil, err
	}

	_ = tx
	return tx, err
}

func (s *staking) RewardToClaim(opts *bind.CallOpts) (*big.Int, error) {
	//workerHub.JoinForMinting()
	dataBytes, err := s.stakingHubAbi.Pack(
		"rewardToClaim",
		s.common.GetWalletAddres(),
	)
	if err != nil {
		return nil, err
	}

	client := s.common.GetClient()
	cAddress := s.common.GetStakingHubAddress()

	msg := ethereum.CallMsg{
		To:   &cAddress,
		Data: dataBytes,
	}

	out, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}

	// Unpack the result
	var result *big.Int
	err = s.stakingHubAbi.UnpackIntoInterface(&result, "rewardToClaim", out)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}

	//fmt.Println(err)
	return result, nil
}

func (s *staking) ClaimReward() (*types.Transaction, error) {
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, s.common.GetClient(), s.common.GetPrivateKey(), int64(s.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	tx, err := s.stakingHub.ClaimReward(auth, s.common.GetWalletAddres())
	if err != nil {
		return nil, err
	}

	return tx, nil
}
