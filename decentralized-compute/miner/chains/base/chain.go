package base

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"core/contracts/worker_hub"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg"
	"solo/pkg/eth"
	"solo/pkg/lighthouse"
	"solo/pkg/utils"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"solo/config"
	"solo/pkg/logger"

	"go.uber.org/zap"
)

type chain struct {
	common          port.ICommon
	workerHub       *worker_hub.WorkerHub
	seizeMinerRoles map[string]bool
	task            *model.Task
}

func (b *chain) SetTask(task *model.Task) {
	b.task = task
}

func NewChain(ctx context.Context, c port.ICommon) (port.IChain, error) {
	wkHub, err := worker_hub.NewWorkerHub(
		c.GetWorkerHubAddress(),
		c.GetClient(),
	)
	if err != nil {
		return nil, err
	}

	return &chain{common: c, seizeMinerRoles: make(map[string]bool), workerHub: wkHub}, nil
}

func (b *chain) getAssigmentInfo(ctx context.Context, assignmentId *big.Int, out chan model.AssimentChan) {
	var err error
	assignment := new(model.Assiment)
	defer func() {
		out <- model.AssimentChan{
			AssismentID: assignmentId,
			Err:         err,
			Data:        assignment,
		}
	}()

	_assignment, err := b.workerHub.Assignments(nil, assignmentId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("WorkerHub.Assignments",
			zap.String("task.InferenceID", assignment.InferenceId.String()),
			zap.Error(err))
		return
	}

	_b, err := json.Marshal(_assignment)
	if err != nil {
		return
	}

	err = json.Unmarshal(_b, assignment)
}

func (c *chain) GetPendingTasks(ctx context.Context, startBlock, endBlock uint64, out chan *model.Task) error {
	iter, err := c.workerHub.FilterRawSubmitted(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks#error", zap.Error(err))
		return err
	}

	for iter.Next() {
		event := &model.Event{}
		err := utils.Copy(event, iter.Event)
		if err != nil {
			return err
		}

		task, err := c.getPendingTask(ctx, event, startBlock, endBlock)
		if err != nil {
			return err
		}

		// push to channel
		out <- task
	}

	return nil
}

func (b *chain) seizeMinerRole(ctx context.Context, task *model.Task, assignment *model.Assiment) (*types.Transaction, error) {
	var err error
	assignmentId := task.AssignmentID
	auth, err := eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), 200_000)
	tx := new(types.Transaction)

	logs := new([]zap.Field)
	*logs = append(*logs, []zap.Field{
		zap.String("task.TaskID", task.TaskID),
		zap.String("task.InferenceID", task.InferenceID),
		zap.String("task.AssignmentRole", task.AssignmentRole),
		zap.String("task.AssignmentID", task.AssignmentID),
	}...)

	defer func() {
		if tx != nil {
			*logs = append(*logs, zap.String("tx", tx.Hash().Hex()))
		}

		if err != nil {
			*logs = append(*logs, zap.Error(err))
			logger.GetLoggerInstanceFromContext(ctx).Error("SeizeMinerRole", *logs...)
		} else {
			logger.GetLoggerInstanceFromContext(ctx).Info("SeizeMinerRole", *logs...)
		}
	}()

	if err != nil {
		return nil, err
	}

	auth.GasLimit = b.common.GetGasLimit()
	auth.GasPrice = auth.GasPrice.Mul(auth.GasPrice, big.NewInt(2))

	id, _ := big.NewInt(0).SetString(assignmentId, 10)
	tx, err = b.workerHub.SeizeMinerRole(auth, id)
	if err != nil {
		return nil, err
	}

	err = eth.WaitForTx(b.common.GetClient(), tx.Hash())
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Implement
func (b *chain) getPendingTask(ctx context.Context, event *model.Event, startBlock, endBlock uint64) (*model.Task, error) {
	var err error
	requestId := event.InferenceId
	requestIdStr := requestId.String()
	_ = requestIdStr
	requestInfo, err := b.workerHub.GetInferenceInfo(nil, requestId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetInferenceInfo", zap.Error(err))
		return nil, err
	}

	assignmentIds, err := b.workerHub.GetAssignmentsByInference(nil, requestId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetAssignmentsByInference", zap.Error(err))
		return nil, err
	}

	respChan := make(chan model.AssimentChan)
	for _, assignmentId := range assignmentIds {
		go b.getAssigmentInfo(ctx, assignmentId, respChan)
	}
	// here

	task := &model.Task{}
	for range assignmentIds {

		assignmentFChan := <-respChan
		if assignmentFChan.Err != nil {
			continue
		}

		assignment := assignmentFChan.Data
		assignmentId := assignmentFChan.AssismentID

		task = &model.Task{
			TaskID:         assignment.InferenceId.String(),
			AssignmentID:   assignmentId.String(),
			ModelContract:  strings.ToLower(event.Model.Hex()),
			Params:         string(requestInfo.Input), // here
			Requestor:      strings.ToLower(requestInfo.Creator.Hex()),
			ZKSync:         true,
			InferenceID:    event.InferenceId.String(),
			AssignmentRole: pkg.MODE_VALIDATOR,
		}

		s, ok := b.seizeMinerRoles[task.AssignmentID]
		if ok && s == true {
			logger.GetLoggerInstanceFromContext(ctx).Info("SeizeMinerRole.Added",
				zap.String("task.TaskID", task.TaskID),
				zap.String("task.InferenceID", task.InferenceID),
				zap.String("task.AssignmentRole", task.AssignmentRole),
				zap.String("task.AssignmentID", task.AssignmentID),
				zap.String("assignment.Worker", assignment.Worker.String()),
				zap.Uint64("start_block", startBlock),
				zap.Uint64("end_block", endBlock),
			)
			continue
		}

		if !strings.EqualFold(assignment.Worker.String(), b.common.GetWalletAddres().Hex()) {
			// fmt.Println("-----> ", assignment.Worker.String())
			continue
		}

		// register as a miner
		transact, err := b.seizeMinerRole(ctx, task, assignment)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("SeizeMinerRole.SeizeMinerRole",
				zap.String("task.TaskID", task.TaskID),
				zap.String("task.InferenceID", task.InferenceID),
				zap.String("task.AssignmentRole", task.AssignmentRole),
				zap.String("task.AssignmentID", task.AssignmentID),
				zap.String("assignment.Worker", assignment.Worker.String()),
				zap.Uint64("start_block", startBlock),
				zap.Uint64("end_block", endBlock),
				zap.Error(err))
			continue
		}

		isReverted, err := eth.CheckTransactionReverted(ctx, b.common.GetClient(), transact.Hash())
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("SeizeMinerRole.CheckTransactionReverted",
				zap.String("task.TaskID", task.TaskID),
				zap.String("task.InferenceID", task.InferenceID),
				zap.String("task.AssignmentRole", task.AssignmentRole),
				zap.String("task.AssignmentID", task.AssignmentID),
				zap.String("assignment.Worker", assignment.Worker.String()),
				zap.Uint64("start_block", startBlock),
				zap.Uint64("end_block", endBlock),
				zap.Error(err))

			continue
		}

		if isReverted {
			err := errors.New(fmt.Sprintf("tx: %s has been reverted", transact.Hash().Hex()))
			logger.GetLoggerInstanceFromContext(ctx).Error("SeizeMinerRole.JoinForMinting",
				zap.String("task.TaskID", task.TaskID),
				zap.String("task.InferenceID", task.InferenceID),
				zap.String("task.AssignmentRole", task.AssignmentRole),
				zap.String("task.AssignmentID", task.AssignmentID),
				zap.Uint64("start_block", startBlock),
				zap.Uint64("end_block", endBlock),
				zap.String("assignment.Worker", assignment.Worker.String()),
				zap.Error(err))
			continue
		}

		task.AssignmentRole = pkg.MODE_MINER

		logger.GetLoggerInstanceFromContext(ctx).Info("SeizeMinerRole.Done",
			zap.String("task.TaskID", task.TaskID),
			zap.String("task.InferenceID", task.InferenceID),
			zap.String("task.AssignmentRole", task.AssignmentRole),
			zap.String("task.AssignmentID", task.AssignmentID),
			zap.String("assignment.Worker", assignment.Worker.String()),
			zap.Uint64("start_block", startBlock),
			zap.Uint64("end_block", endBlock),
			zap.String("tx", transact.Hash().Hex()))

		var batchInfers []*model.BatchInferHistory
		var externalData *model.AgentInferExternalData

		/*
			TODO - chainConfig.AgentContractAddress ???
			if chainConfig.AgentContractAddress != "" {
				isAgentInfer, batchInfers, externalData, err = s.handleNewInferIsAgentInfer(ctx, modelInfo.ModelID.String(), chainConfig, ethClient, event.Raw.TxHash, aiZKClient)
				if err != nil {
					return err
				}
			}*/

		// Detect if  is batch
		isBatch := false
		if strings.HasPrefix(string(requestInfo.Input), config.IPFSPrefix) {
			// TODO - HERE
			inputBytes, _, err := lighthouse.DownloadDataSimpleWithRetry(string(requestInfo.Input))
			if err == nil {
				batchFullPrompts := []*model.BatchInferHistory{}
				err = json.Unmarshal(inputBytes, &batchFullPrompts)
				if err != nil {
					logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
				} else if len(batchFullPrompts) > 0 {
					batchInfers = batchFullPrompts
					isBatch = true
				}
			} else {
				logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
			}

		}

		task.IsBatch = isBatch
		task.BatchInfers = batchInfers // here
		task.ExternalData = externalData
		b.seizeMinerRoles[task.AssignmentID] = true
		return task, nil
	}

	return nil, nil
}

func (b *chain) SubmitTask(ctx context.Context, assigmentID *big.Int, result []byte) (*types.Transaction, error) {
	auth, err := eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), int64(b.common.GetGasLimit()))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#CreateBindTransactionOpts",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Any("assigment_id", assigmentID),
			zap.Any("result", string(result)),
			zap.Error(err))
		return nil, err
	}

	auth.GasLimit = b.common.GetGasLimit()
	auth.GasPrice = auth.GasPrice.Mul(auth.GasPrice, big.NewInt(2))
	tx, err := b.workerHub.SubmitSolution(auth, assigmentID, result)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#SubmitSolution",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Any("assigment_id", assigmentID),
			zap.Any("result", string(result)),
			zap.Error(err),
		)
		return nil, err
	}

	err = eth.WaitForTx(b.common.GetClient(), tx.Hash())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#WaitForTx",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Error(err),
			zap.Any("assigment_id", assigmentID))
		return nil, err
	}

	receipt, err := b.common.GetClient().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#TransactionReceipt",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Error(err),
			zap.Any("assigment_id", assigmentID))
		return nil, err
	}

	//TODO - check this
	/*if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("tx failed")
	}

	for _, txLog := range receipt.Logs {
		feeLog, err := workerHub.WorkerHubFilterer.ParseTransferFee(*txLog)
		if err != nil {
			continue
		} else {
			if strings.EqualFold(feeLog.Miner.Hex(), tskw.address) {
				tskw.status.processedTasks++
				tskw.status.currentEarning.Add(tskw.status.currentEarning, feeLog.MingingFee)
			}
		}
	}*/
	_ = receipt

	return tx, nil
}

func (b *chain) GetInferenceByMiner() ([]*big.Int, error) {
	return []*big.Int{}, nil
}

func (b *chain) GetInferenceInfo(opt *bind.CallOpts, inferID uint64) (*model.InferInfo, error) {
	ID := big.NewInt(1).SetUint64(inferID)
	t, err := b.workerHub.GetInferenceInfo(opt, ID)
	if err != nil {
		return nil, err
	}

	_ = t
	return &model.InferInfo{}, nil
}
