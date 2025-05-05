package ethapi

import (
	"context"
	"math/big"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/uniswapv3factory"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/uniswapv3pool"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type UniswapPoolCreatedEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Token0          string   `json:"token0"`
	Token1          string   `json:"token1"`
	Pool            string   `json:"pool"`
	Fee             *big.Int `json:"fee"`
	TickSpacing     *big.Int `json:"tick_spacing"`
	Index           uint     `json:"log_index"`
	NetworkID       uint64   `json:"network_id"`
}

type UniswapSwapEventResp struct {
	FromAddress     string   `json:"from_address"`
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Sender          string   `json:"sender"`
	Recipient       string   `json:"recipient"`
	Amount0         *big.Int `json:"amount0"`
	Amount1         *big.Int `json:"amount1"`
	SqrtPriceX96    *big.Int `json:"sqrt_price_x96"`
	Liquidity       *big.Int `json:"liquidity"`
	Tick            *big.Int `json:"tick"`
	Index           uint     `json:"log_index"`
	TxIndex         uint     `json:"tx_index"`
	BlockNumber     uint64   `json:"block_number"`
	NetworkID       uint64   `json:"network_id"`
}

type BlockChainExternalEventResp struct {
	UniV3PoolCreated []*UniswapPoolCreatedEventResp `json:"uni_v3_pool_created"`
	UniV3Swap        []*UniswapSwapEventResp        `json:"uni_v3_swap"`
	LastBlockNumber  int64                          `json:"last_block_number"`
	BlockNumber      uint64                         `json:"block_number"`
}

func (c *Client) NewExternalEventResp() *BlockChainExternalEventResp {
	return &BlockChainExternalEventResp{}
}

func (c *Client) ExternalEventEventsByTransaction(txHash string) (*BlockChainExternalEventResp, error) {
	resp := c.NewExternalEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	for _, log := range receipt.Logs {
		err = c.ParseExternalEventResp(resp, log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) ScanExternalEvents(contractAddrs []string, startBlock, endBlock int64) (*BlockChainExternalEventResp, error) {
	resp := c.NewExternalEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	var contractAddresses []common.Address
	if len(contractAddrs) > 0 {
		contractAddresses = []common.Address{}
		for _, contractAddr := range contractAddrs {
			if contractAddr != "" {
				contractAddresses = append(contractAddresses, helpers.HexToAddress(contractAddr))
			}
		}
	}
	ctx := context.Background()
	lastBlock, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	lastNumber := lastBlock.Number.Int64()
	if endBlock > lastNumber {
		endBlock = lastNumber
	}
	if startBlock > endBlock {
		return nil, nil
	}
	resp.LastBlockNumber = lastNumber
	resp.BlockNumber = lastBlock.Number.Uint64()
	logs, err := client.FilterLogs(
		ctx,
		ethereum.FilterQuery{
			FromBlock: big.NewInt(startBlock),
			ToBlock:   big.NewInt(endBlock),
			Addresses: contractAddresses,
			Topics: [][]common.Hash{
				{
					// univ3
					common.HexToHash("0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"),
					common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		err = c.ParseExternalEventResp(resp, &log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) ParseExternalEventResp(resp *BlockChainExternalEventResp, log *types.Log) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	blockTime := uint64(time.Now().Unix())
	uniswap, err := uniswapv3factory.NewUniswapv3factory(log.Address, client)
	if err != nil {
		return err
	}
	networkID, err := c.ChainID()
	if err != nil {
		return err
	}
	// ParsePoolCreated
	{
		logParsed, err := uniswap.ParsePoolCreated(*log)
		if err == nil {
			resp.UniV3PoolCreated = append(
				resp.UniV3PoolCreated,
				&UniswapPoolCreatedEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Token0:          logParsed.Token0.Hex(),
					Token1:          logParsed.Token1.Hex(),
					Fee:             logParsed.Fee,
					Pool:            logParsed.Pool.Hex(),
					TickSpacing:     logParsed.TickSpacing,
					Index:           log.Index,
					NetworkID:       networkID,
				},
			)
		}
	}

	uniswappair, err := uniswapv3pool.NewUniswapv3pool(log.Address, client)
	if err != nil {
		return err
	}
	// ParseSwap
	{
		logParsed, err := uniswappair.ParseSwap(*log)
		if err == nil {
			resp.UniV3Swap = append(
				resp.UniV3Swap,
				&UniswapSwapEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Sender:          logParsed.Sender.Hex(),
					Recipient:       logParsed.Recipient.Hex(),
					Amount0:         logParsed.Amount0,
					Amount1:         logParsed.Amount1,
					SqrtPriceX96:    logParsed.SqrtPriceX96,
					Liquidity:       logParsed.Liquidity,
					Tick:            logParsed.Tick,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					NetworkID:       networkID,
				},
			)
		}
	}
	return nil
}
