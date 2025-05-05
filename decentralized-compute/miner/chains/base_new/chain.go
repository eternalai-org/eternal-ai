package base_new

import (
	"context"
	"core/contracts/prompt_scheduler"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"solo/config"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg"
	"solo/pkg/eth"
	"solo/pkg/lighthouse"
	"solo/pkg/logger"
	"solo/pkg/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"go.uber.org/zap"
)

type chain struct {
	common          port.ICommon
	workerHub       *prompt_scheduler.PromptScheduler
	seizeMinerRoles map[string]bool
	task            *model.Task
}

func (b *chain) SetTask(task *model.Task) {
	b.task = task
}

func NewChain(ctx context.Context, c port.ICommon) (port.IChain, error) {
	wkHub, err := prompt_scheduler.NewPromptScheduler(
		c.GetWorkerHubAddress(),
		c.GetClient(),
	)
	if err != nil {
		return nil, err
	}

	return &chain{common: c, seizeMinerRoles: make(map[string]bool), workerHub: wkHub}, nil
}

func (b *chain) getAssigmentInfo(ctx context.Context, inferenceID *big.Int, out chan model.AssimentChan) {
	var err error
	assignment := new(model.Assiment)
	defer func() {
		out <- model.AssimentChan{
			AssismentID: inferenceID,
			Err:         err,
			Data:        assignment,
		}
	}()

	_assignment, err := b.workerHub.GetInferenceInfo(nil, inferenceID.Uint64())
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
	iter, err := c.workerHub.FilterNewInference(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks#error", zap.Error(err))
		return err
	}

	for iter.Next() {
		event := &model.EventPromptSchedulerNewInference{}
		err := utils.Copy(event, iter.Event)
		if err != nil {
			return err
		}

		currentBlock := event.Raw.BlockNumber

		task, err := c.getPendingTask(ctx, event, startBlock, endBlock)
		if err != nil {
			return err
		}

		if task == nil {
			continue
		}

		// push to channel
		logger.GetLoggerInstanceFromContext(ctx).Info("getPendingTask.Processing",
			zap.String("inference_id", task.TaskID),
			zap.Uint64("start_block", startBlock),
			zap.Uint64("end_block", endBlock),
			zap.Uint64("current_block", currentBlock),
		)
		out <- task
	}

	return nil
}

/*
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
*/
// Implement
func (b *chain) getPendingTask(ctx context.Context, event *model.EventPromptSchedulerNewInference, startBlock, endBlock uint64) (*model.Task, error) {
	var err error
	inferenceId := event.InferenceId
	//requestIdStr := requestId.String()
	//_ = requestIdStr
	if inferenceId == 0 {
		return nil, err
	}

	s, ok := b.seizeMinerRoles[fmt.Sprintf("%d", inferenceId)]
	if ok && s {
		/*logger.GetLoggerInstanceFromContext(ctx).Info("getPendingTask.Added",
			zap.Uint64("task.InferenceID", inferenceId),
			zap.Uint64("start_block", startBlock),
			zap.Uint64("end_block", endBlock),
		)*/
		return nil, nil
	}

	requestInfo, err := b.workerHub.GetInferenceInfo(nil, inferenceId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("getPendingTask",
			zap.Uint64("start_block", startBlock),
			zap.Uint64("end_block", endBlock),
			zap.Error(err))
		return nil, err
	}

	if !strings.EqualFold(b.common.GetWalletAddres().Hex(), requestInfo.ProcessedMiner.Hex()) {
		err := errors.New(fmt.Sprintf("task has been assigned to worker %s", requestInfo.ProcessedMiner.Hex()))
		return nil, err
	}

	_input, err := pkg.ExtractContent(string(requestInfo.Input))
	if err != nil {
		return nil, err
	}

	task := &model.Task{
		TaskID:         fmt.Sprintf("%d", inferenceId), // TaskID = AssignmentID = inferenceId
		AssignmentID:   fmt.Sprintf("%d", inferenceId),
		ModelContract:  fmt.Sprintf("%d", event.ModelId),
		Params:         _input, // here
		Requestor:      strings.ToLower(requestInfo.Creator.Hex()),
		ZKSync:         true,
		InferenceID:    fmt.Sprintf("%d", inferenceId),
		AssignmentRole: pkg.MODE_VALIDATOR,
	}

	task.AssignmentRole = pkg.MODE_MINER
	// zap.String("tx", transact.Hash().Hex()))

	var batchInfers []*model.BatchInferHistory
	var externalData *model.AgentInferExternalData

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
	b.seizeMinerRoles[task.TaskID] = true
	return task, nil
}

func (b *chain) SubmitTask(ctx context.Context, inferenceID *big.Int, result []byte) (*types.Transaction, error) {
	auth, err := eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), int64(b.common.GetGasLimit()))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#CreateBindTransactionOpts",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Any("inference_id", inferenceID.Uint64()),
			zap.Any("result", string(result)),
			zap.Error(err))
		return nil, err
	}

	auth.GasLimit = b.common.GetGasLimit()
	auth.GasPrice = auth.GasPrice.Mul(auth.GasPrice, big.NewInt(2))
	tx, err := b.workerHub.SubmitSolution(auth, inferenceID.Uint64(), result)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#SubmitSolution",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Any("inference_id", inferenceID.Uint64()),
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
			zap.Any("inference_id", inferenceID.Uint64()))
		return nil, err
	}

	receipt, err := b.common.GetClient().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#TransactionReceipt",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Error(err),
			zap.Any("inference_id", inferenceID.Uint64()))
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

	t, err := b.workerHub.GetInferenceByMiner(nil, b.common.GetWalletAddres())
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (b *chain) GetInferenceInfo(opt *bind.CallOpts, inferID uint64) (*model.InferInfo, error) {
	t, err := b.workerHub.GetInferenceInfo(opt, inferID)
	if err != nil {
		return nil, err
	}

	return &model.InferInfo{
		Value:   t.Value,
		Output:  t.Output,
		ModelId: t.ModelId,
		Input:   t.Input,
		Status:  t.Status,
	}, nil
}
