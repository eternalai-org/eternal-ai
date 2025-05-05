package base

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"core/contracts/staking_hub"
	"solo/internal/port"
	"solo/pkg/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type staking struct {
	stakingHub *staking_hub.StakingHub
	common     port.ICommon
}

func NewStaking(common port.ICommon, stakingHub *staking_hub.StakingHub) port.IStaking {
	return &staking{
		common:     common,
		stakingHub: stakingHub,
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
		return nil, err
	}

	err = eth.ApproveERC20(
		ctx, s.common.GetClient(), s.common.GetPrivateKey(), s.common.GetStakingHubAddress(), s.common.GetErc20contractAddress(), int64(s.common.GetGasLimit()),
	)
	if err != nil {
		return nil, err
	}

	minStake, err := s.stakingHub.MinerMinimumStake(nil)
	if err != nil {
		return nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, s.common.GetClient(), s.common.GetPrivateKey(), int64(s.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	auth.Value = minStake
	tx := new(types.Transaction)
	if s.common.GetModelAddress() != "" {
		tx, err = s.stakingHub.RegisterMiner0(auth, 1, common.HexToAddress(s.common.GetModelAddress()))
		if err != nil {
			return nil, err
		}
	} else {
		tx, err = s.stakingHub.RegisterMiner(auth, 1)
		if err != nil {
			return nil, err
		}
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

	tx, err := s.stakingHub.StakingHubTransactor.JoinForMinting(auth)
	if err != nil {
		return nil, err
	}

	_ = tx
	return tx, nil
}

func (s *staking) RewardToClaim(opts *bind.CallOpts) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (s *staking) ClaimReward() (*types.Transaction, error) {
	//TODO - implement me
	return nil, nil
}
