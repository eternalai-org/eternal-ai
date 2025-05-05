package usecase

import (
	"context"
	"math/big"
	"solo/config"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg/db/sqlite"
	"testing"

	interCommon "solo/chains/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-redis/redis"
)

func Test_miner_saveSuccessTask(t *testing.T) {
	type fields struct {
		cnf          *config.Config
		IsStaked     *bool
		currentBlock uint64
		tasksQueue   chan *model.Task
		rdb          *redis.Client
		chain        port.IChain
		staking      port.IStaking
		common       port.ICommon
		cluster      port.ICluster
	}
	type args struct {
		ctx        context.Context
		task       *model.Task
		resultData []byte
		tx         *types.Transaction
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				ctx: context.TODO(),
				task: &model.Task{
					TaskID:         "66",
					AssignmentID:   "66",
					ModelContract:  "700050",
					Params:         `{"messages":[{"role":"user","content":"hello"}],"model":"NousResearch/DeepHermes-3-Llama-3-8B-Preview","seed":0,"max_tokens":0,"temperature":0,"stream":false}`,
					AssignmentRole: "miner",
					ZKSync:         true,
					Requestor:      "0x4b5c74f43cb232264d546ec40658cc99f57f55d3",
					InferenceID:    "66",
				},
				resultData: []byte("eyJyZXN1bHRfdXJpIjoiaXBmczovL2xpZ2h0aG91c2UtZXJyb3IiLCJzdG9yYWdlIjoibGlnaHRob3VzZS1maWxlY29pbiIsImRhdGEiOiJleUpwWkNJNkltTm9ZWFJqYlhCc0xUTmxPVFZqTmpJNFptWmtNalJsTTJFNE1UTm1OVGcxTlRGaE56RTFZbU01SWl3aWIySnFaV04wSWpvaVkyaGhkQzVqYjIxd2JHVjBhVzl1TG1Ob2RXNXJJaXdpWTNKbFlYUmxaQ0k2TVRjME1EUTNOamd6TUN3aWJXOWtaV3dpT2lKT2IzVnpVbVZ6WldGeVkyZ3ZSR1ZsY0VobGNtMWxjeTB6TFV4c1lXMWhMVE10T0VJdFVISmxkbWxsZHlJc0ltTm9iMmxqWlhNaU9sdDdJbWx1WkdWNElqb3dMQ0p0WlhOellXZGxJanA3SW5KdmJHVWlPaUpoYzNOcGMzUmhiblFpTENKamIyNTBaVzUwSWpvaVNHVnNiRzhoSUVodmR5QmpZVzRnU1NCaGMzTnBjM1FnZVc5MUlIUnZaR0Y1UHlJc0luUnZiMnhmWTJGc2JITWlPbTUxYkd4OUxDSnNiMmR3Y205aWN5STZiblZzYkN3aVptbHVhWE5vWDNKbFlYTnZiaUk2SWlJc0luTjBiM0JmY21WaGMyOXVJanB1ZFd4c2ZWMHNJblZ6WVdkbElqcDdJbkJ5YjIxd2RGOTBiMnRsYm5NaU9qQXNJblJ2ZEdGc1gzUnZhMlZ1Y3lJNk1Dd2lZMjl0Y0d4bGRHbHZibDkwYjJ0bGJuTWlPakFzSW5CeWIyMXdkRjkwYjJ0bGJuTmZaR1YwWVdsc2N5STZiblZzYkgwc0luQnliMjF3ZEY5c2IyZHdjbTlpY3lJNmJuVnNiQ3dpYVhOZmMzUnZjQ0k2Wm1Gc2MyVXNJbTl1WTJoaGFXNWZaR0YwWVNJNmV5SnBibVpsY2w5cFpDSTZNQ3dpY0dKbWRGOWpiMjF0YVhSMFpXVWlPbTUxYkd3c0luQnliM0J2YzJWeUlqb2lJaXdpYVc1bVpYSmZkSGdpT2lJaUxDSndjbTl3YjNObFgzUjRJam9pSW4xOSJ9"),
				tx:         createFakeTransaction(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.cnf = &config.Config{
				NodeID:  "0553",
				Account: "bdcb9fef30a0e89291df236a8cbcd6e4c9cd03ca5812e41c4d662b076c382181",
				Rpc:     "https://base.llamarpc.com",
			}

			tr := &miner{
				cnf:          tt.fields.cnf,
				IsStaked:     tt.fields.IsStaked,
				currentBlock: tt.fields.currentBlock,
				tasksQueue:   tt.fields.tasksQueue,
				rdb:          tt.fields.rdb,
				chain:        tt.fields.chain,
				staking:      tt.fields.staking,
				common:       tt.fields.common,
				cluster:      tt.fields.cluster,
			}

			sqldb, err := sqlite.NewSqliteDB(tt.fields.cnf.NodeID)
			if err == nil {
				sqldb.Connect()
				tr.SetDB(sqldb)
			}

			//create tables
			defer sqldb.Close()

			db := model.SuccessTask{}
			sqldb.CreateTable(db.CreateDB())

			tr.common, err = interCommon.NewCommon(context.Background(), tt.fields.cnf)
			if err != nil {
				t.Errorf("miner.saveSuccessTask() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tr.saveSuccessTask(tt.args.ctx, tt.args.task, tt.args.resultData, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("miner.saveSuccessTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createFakeTransaction() *types.Transaction {
	// Create fake addresses
	// fromAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	toAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")

	// Create fake values
	value := big.NewInt(1000)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(1000000000) // 1 gwei
	nonce := uint64(0)
	data := []byte("fake transaction data")

	// Create the transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	return tx
}
