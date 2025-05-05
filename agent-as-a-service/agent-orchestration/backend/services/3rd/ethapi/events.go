package ethapi

import (
	"context"
	"math/big"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentfactory"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentshares"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentupgradeable"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/arbitrumfactory"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/arbitrumnonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/basenonfungiblepositionmanager"
	brigde "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/bridge"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc1155"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20"
	realworldagent "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20realworldagent"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc721"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/imagehub"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/iworkerhub"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/systempromptmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/uniswapv3factory"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/uniswapv3pool"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/vibetokenfactory"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/wbvm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type RealWorldAgentExecutionRequested struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Index           uint     `json:"log_index"`
	TxIndex         uint     `json:"tx_index"`
	BlockNumber     uint64   `json:"block_number"`
	ActId           *big.Int `json:"act_id"`
	Uuid            string   `json:"uuid"`
	Creator         string   `json:"creator"`
	Data            string   `json:"data"`
}

type BridgeMint struct {
	TxHash          string     `json:"tx_hash"`
	ContractAddress string     `json:"contract_address"`
	Timestamp       uint64     `json:"timestamp"`
	Index           uint       `json:"log_index"`
	TxIndex         uint       `json:"tx_index"`
	BlockNumber     uint64     `json:"block_number"`
	Tokens          []string   `json:"tokens"`
	Recipients      []string   `json:"recipients"`
	Amounts         []*big.Int `json:"amounts"`
}

type BridgeWithdraw struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Index           uint     `json:"log_index"`
	TxIndex         uint     `json:"tx_index"`
	BlockNumber     uint64   `json:"block_number"`
	Token           string   `json:"token"`
	Burner          string   `json:"burner"`
	Amount          *big.Int `json:"amount"`
	Extddr          string   `json:"extddr"`
	DestChainId     *big.Int `json:"dest_chain_id"`
}

type ImageHubImageTipTransferred struct {
	NetworkID       uint64 `json:"network_id"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Timestamp       uint64 `json:"timestamp"`
	Index           uint   `json:"log_index"`
	TxIndex         uint   `json:"tx_index"`
	BlockNumber     uint64 `json:"block_number"`
	User            string
	Creator         string
	ImageId         *big.Int
	Amount          *big.Int
	Fee             *big.Int
}

type ImageHubSubscriptionPriceUpdated struct {
	NetworkID       uint64 `json:"network_id"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Timestamp       uint64 `json:"timestamp"`
	Index           uint   `json:"log_index"`
	TxIndex         uint   `json:"tx_index"`
	BlockNumber     uint64 `json:"block_number"`
	Creator         string
	Duration        *big.Int
	Price           *big.Int
}

type ImageHubSubscriptionRegistered struct {
	NetworkID       uint64 `json:"network_id"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Timestamp       uint64 `json:"timestamp"`
	Index           uint   `json:"log_index"`
	TxIndex         uint   `json:"tx_index"`
	BlockNumber     uint64 `json:"block_number"`
	User            string
	Creator         string
	Duration        *big.Int
	Price           *big.Int
}

type ImageHubSubscriptionCharged struct {
	NetworkID       uint64 `json:"network_id"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Timestamp       uint64 `json:"timestamp"`
	Index           uint   `json:"log_index"`
	TxIndex         uint   `json:"tx_index"`
	BlockNumber     uint64 `json:"block_number"`
	User            string
	Creator         string
	Duration        *big.Int
	ExpiredAt       *big.Int
	Amount          *big.Int
	Fee             *big.Int
}

type WorkerHubNewInference struct {
	NetworkID       uint64   `json:"network_id"`
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Index           uint     `json:"log_index"`
	TxIndex         uint     `json:"tx_index"`
	BlockNumber     uint64   `json:"block_number"`
	InferenceId     *big.Int `json:"inference_id"`
	Model           string   `json:"model"`
	Creator         string   `json:"creator"`
	Value           *big.Int `json:"value"`
}

type BlockChainEventResp struct {
	Transfer                            []*Erc20TokenTransferEventResp                            `json:"transfer"`
	NftTransfer                         []*NftTransferEventResp                                   `json:"nft_transfer"`
	ERC1155Transfer                     []*ERC1155ransferEventResp                                `json:"erc1155_transfer"`
	BridgeMint                          []*BridgeMint                                             `json:"bridge_mint"`
	BridgeWithdraw                      []*BridgeWithdraw                                         `json:"bridge_withdraw"`
	ImageHubImageTipTransferred         []*ImageHubImageTipTransferred                            `json:"image_hub_image_tip_transferred"`
	ImageHubSubscriptionPriceUpdated    []*ImageHubSubscriptionPriceUpdated                       `json:"image_hub_subscription_price_updated"`
	ImageHubSubscriptionRegistered      []*ImageHubSubscriptionRegistered                         `json:"image_hub_subscription_registered"`
	ImageHubSubscriptionCharged         []*ImageHubSubscriptionCharged                            `json:"image_hub_subscription_charged"`
	WorkerHubNewInference               []*WorkerHubNewInference                                  `json:"worker_hub_new_inference"`
	LastBlockNumber                     int64                                                     `json:"last_block_number"`
	BlockNumber                         uint64                                                    `json:"block_number"`
	MemePoolCreated                     []*UniswapPoolCreatedEventResp                            `json:"meme_pool_created"`
	MemeSwap                            []*UniswapSwapEventResp                                   `json:"meme_swap"`
	MemeIncreaseLiquidity               []*UniswapPositionLiquidity                               `json:"meme_increase_liquidity"`
	MemeDecreaseLiquidity               []*UniswapPositionLiquidity                               `json:"meme_decrease_liquidity"`
	AgentSharesTrades                   []*agentshares.AgentSharesTrade                           `json:"agent_shares_trades"`
	SystemPromptManagerNewTokens        []*systempromptmanager.SystemPromptManagerNewToken        `json:"system_prompt_manager_new_tokens"`
	SystemPromptManagerAgentDataUpdates []*systempromptmanager.SystemPromptManagerAgentDataUpdate `json:"system_prompt_manager_agent_data_updates"`
	SystemPromptManagerAgentURIUpdates  []*systempromptmanager.SystemPromptManagerAgentURIUpdate  `json:"system_prompt_manager_agent_uri_updates"`
	RealWorldAgentExecutionRequested    []*RealWorldAgentExecutionRequested                       `json:"real_world_agent_execution_requested"`
	CodePointerCreated                  []*agentupgradeable.AgentUpgradeableCodePointerCreated    `json:"code_pointer_created"`
	AgentCreated                        []*agentfactory.AgentFactoryAgentCreated                  `json:"agent_created"`
	VibeTokenFactoryTokenDeployed       []*vibetokenfactory.VibeTokenFactoryTokenDeployed         `json:"vibe_token_factory_token_deployed"`
}

func (c *Client) NewEventResp() *BlockChainEventResp {
	return &BlockChainEventResp{}
}

func (c *Client) EventsByTransaction(txHash string) (*BlockChainEventResp, error) {
	resp := c.NewEventResp()
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
		err = c.ParseEventResp(resp, log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) ScanEvents(contractAddrs []string, startBlock, endBlock int64) (*BlockChainEventResp, error) {
	resp := c.NewEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	var contractAddresses []common.Address
	if len(contractAddrs) > 0 {
		contractAddresses = []common.Address{}
		for _, contractAddr := range contractAddrs {
			if contractAddr != "" {
				contractAddresses = append(contractAddresses, helpers.HexToAddress(c.ConvertAddressForIn(contractAddr)))
			}
		}
	}
	ctx := context.Background()
	lastBlock, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	lastNumber := int64(lastBlock)
	if endBlock > lastNumber {
		endBlock = lastNumber
	}
	if startBlock > endBlock {
		return nil, nil
	}
	resp.LastBlockNumber = lastNumber
	resp.BlockNumber = lastBlock
	logs, err := client.FilterLogs(
		ctx,
		ethereum.FilterQuery{
			FromBlock: big.NewInt(startBlock),
			ToBlock:   big.NewInt(endBlock),
			Addresses: contractAddresses,
			Topics: [][]common.Hash{
				{
					// erc
					common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
					common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
					common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"),
					common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
					// imageHub
					common.HexToHash("0x9d239c3ad6158bc2f486c3310d25622c5250fa11f60eb415f1d5b848ac0d145e"),
					common.HexToHash("0xab7a1a256e1a3e6eefe1e11e862d92b2a9efd3ab45a0a0b3d330687f42021ec0"),
					common.HexToHash("0xe1b67d1786c6125e949791ee25291c8dc38b471ae947bf8c65433da2f31149bf"),
					common.HexToHash("0x0bc8ce05a473cd40acc9b6689277d1b7c59ff62cb5491e8ec926bddb9e365a36"),
					// workerhub
					common.HexToHash("0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b"),
					// Uniswap
					common.HexToHash("0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"),
					common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
					common.HexToHash("0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f"),
					common.HexToHash("0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4"),
					// AlgebraFactory
					common.HexToHash("0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db"),
					common.HexToHash("0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f"),
					// agentShares
					common.HexToHash("0x12d0646903287d48eb117ac55a8bcc90d4357c4180221d5b33e83e73860440ec"),
					//
					common.HexToHash("0x61beab98a81083e3c0239c33e149bef1316ca78f15b9f29125039f5521a06d06"),
					common.HexToHash("0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e"),
					common.HexToHash("0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362"),
					//RealWorldAgent
					common.HexToHash("0x9096741026bdd638bcc5cb995f0f00b4574b81f120a23c4a7086347116bf58a1"),
					// AgentUpgradeable
					common.HexToHash("0x0c69de805f518c322126e1fa81f2e91d814412a2ad304638fcaed200bd7f43dc"),
					// AgentFactory
					common.HexToHash("0x7c96960a1ebd8cc753b10836ea25bd7c9c4f8cd43590db1e8b3648cb0ec4cc89"),
					// VibeTokenFactory
					common.HexToHash("0x035044c956f0f85240e40898163f0636d588d72de5fd08d79a9168d626745173"),
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		err = c.ParseEventResp(resp, &log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) ParseEventResp(resp *BlockChainEventResp, log *types.Log) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	chainId, err := c.ChainID()
	if err != nil {
		return err
	}
	blockTime := uint64(time.Now().Unix())
	erc20, err := erc20.NewErc20(log.Address, client)
	if err != nil {
		return err
	}
	// ParseTransfer
	{
		logParsed, err := erc20.ParseTransfer(*log)
		if err == nil {
			resp.Transfer = append(
				resp.Transfer,
				&Erc20TokenTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: c.ConvertAddressForOut(logParsed.Raw.Address.Hex()),
					Timestamp:       blockTime,
					From:            c.ConvertAddressForOut(logParsed.From.Hex()),
					To:              c.ConvertAddressForOut(logParsed.To.Hex()),
					Value:           logParsed.Value,
					Index:           log.Index,
					BlockNumber:     log.BlockNumber,
					TxIndex:         log.TxIndex,
				},
			)
		}
	}
	// ParseTransfer
	{
		logParsed, err := erc20.ParseDeposit(*log)
		if err == nil {
			resp.Transfer = append(
				resp.Transfer,
				&Erc20TokenTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            helpers.HexToAddress("0x0000000000000000000000000000000000000000").Hex(),
					To:              logParsed.Dst.Hex(),
					Value:           logParsed.Wad,
					Index:           log.Index,
					BlockNumber:     log.BlockNumber,
					TxIndex:         log.TxIndex,
				},
			)
		}
	}
	{
		logParsed, err := erc20.ParseWithdrawal(*log)
		if err == nil {
			resp.Transfer = append(
				resp.Transfer,
				&Erc20TokenTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.Src.Hex(),
					To:              helpers.HexToAddress("0x0000000000000000000000000000000000000000").Hex(),
					Value:           logParsed.Wad,
					Index:           log.Index,
					BlockNumber:     log.BlockNumber,
					TxIndex:         log.TxIndex,
				},
			)
		}
	}
	//nft transfer
	erc721, err := erc721.NewErc721(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := erc721.ParseTransfer(*log)
		if err == nil {
			resp.NftTransfer = append(
				resp.NftTransfer,
				&NftTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.From.Hex(),
					To:              logParsed.To.Hex(),
					TokenId:         logParsed.TokenId,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
				},
			)
		}
	}
	//erc1155 transfer
	erc1155, err := erc1155.NewERC1155(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := erc1155.ParseTransferBatch(*log)
		if err == nil {
			for index := range logParsed.Ids {
				resp.ERC1155Transfer = append(
					resp.ERC1155Transfer,
					&ERC1155ransferEventResp{
						NetworkID:       chainId,
						TxHash:          log.TxHash.Hex(),
						ContractAddress: logParsed.Raw.Address.Hex(),
						Timestamp:       blockTime,
						From:            logParsed.From.Hex(),
						To:              logParsed.To.Hex(),
						Id:              logParsed.Ids[index],
						Value:           logParsed.Values[index],
					},
				)
			}

		}
	}

	{
		logParsed, err := erc1155.ParseTransferSingle(*log)
		if err == nil {
			resp.ERC1155Transfer = append(
				resp.ERC1155Transfer,
				&ERC1155ransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.From.Hex(),
					To:              logParsed.To.Hex(),
					Id:              logParsed.Id,
					Value:           logParsed.Value,
				},
			)
		}
	}
	//
	bridge, err := brigde.NewBrigde(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := bridge.ParseMint(*log)
		if err == nil {
			tmpObj := &BridgeMint{
				TxHash:          log.TxHash.Hex(),
				ContractAddress: logParsed.Raw.Address.Hex(),
				Timestamp:       blockTime,
				Index:           log.Index,
				TxIndex:         log.TxIndex,
				BlockNumber:     log.BlockNumber,
			}

			for i, _ := range logParsed.Tokens {
				tmpObj.Tokens = append(tmpObj.Tokens, logParsed.Tokens[i].Hex())
				tmpObj.Recipients = append(tmpObj.Recipients, logParsed.Recipients[i].Hex())
				tmpObj.Amounts = append(tmpObj.Amounts, logParsed.Amounts[i])
			}

			resp.BridgeMint = append(
				resp.BridgeMint, tmpObj,
			)
		}
	}
	{
		logParsed, err := bridge.ParseBridgeToken(*log)
		if err == nil {
			resp.BridgeWithdraw = append(
				resp.BridgeWithdraw,
				&BridgeWithdraw{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					Token:           logParsed.Token.Hex(),
					Burner:          logParsed.Burner.Hex(),
					Amount:          logParsed.Amount,
					Extddr:          logParsed.Extddr,
					DestChainId:     logParsed.DestChainId,
				},
			)
		}
	}
	imagehub, err := imagehub.NewImageHub(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := imagehub.ParseImageTipTransferred(*log)
		if err == nil {
			resp.ImageHubImageTipTransferred = append(
				resp.ImageHubImageTipTransferred,
				&ImageHubImageTipTransferred{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					User:            logParsed.User.Hex(),
					Creator:         logParsed.Creator.Hex(),
					ImageId:         logParsed.ImageId,
					Amount:          logParsed.Amount,
					Fee:             logParsed.Fee,
				},
			)
		}
	}
	{
		logParsed, err := imagehub.ParseSubscriptionPriceUpdated(*log)
		if err == nil {
			resp.ImageHubSubscriptionPriceUpdated = append(
				resp.ImageHubSubscriptionPriceUpdated,
				&ImageHubSubscriptionPriceUpdated{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					Creator:         logParsed.Creator.Hex(),
					Duration:        logParsed.Duration,
					Price:           logParsed.Price,
				},
			)
		}
	}
	{
		logParsed, err := imagehub.ParseSubscriptionRegistered(*log)
		if err == nil {
			resp.ImageHubSubscriptionRegistered = append(
				resp.ImageHubSubscriptionRegistered,
				&ImageHubSubscriptionRegistered{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					User:            logParsed.User.Hex(),
					Creator:         logParsed.Creator.Hex(),
					Duration:        logParsed.Duration,
					Price:           logParsed.Price,
				},
			)
		}
	}
	{
		logParsed, err := imagehub.ParseSubscriptionCharged(*log)
		if err == nil {
			resp.ImageHubSubscriptionCharged = append(
				resp.ImageHubSubscriptionCharged,
				&ImageHubSubscriptionCharged{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					User:            logParsed.User.Hex(),
					Creator:         logParsed.Creator.Hex(),
					Duration:        logParsed.Duration,
					ExpiredAt:       logParsed.ExpiredAt,
					Amount:          logParsed.Amount,
					Fee:             logParsed.Fee,
				},
			)
		}
	}
	//
	workerhub, err := iworkerhub.NewIWorkerHub(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := workerhub.ParseNewInference(*log)
		if err == nil {
			resp.WorkerHubNewInference = append(
				resp.WorkerHubNewInference,
				&WorkerHubNewInference{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					InferenceId:     logParsed.InferenceId,
					Model:           logParsed.Model.Hex(),
					Creator:         logParsed.Creator.Hex(),
					Value:           logParsed.Value,
				},
			)
		}
	}

	//

	uniswap, err := uniswapv3factory.NewUniswapv3factory(log.Address, client)
	if err != nil {
		return err
	}
	uniswapArb, err := arbitrumfactory.NewFactory(log.Address, client)
	if err != nil {
		return err
	}

	// ParsePoolCreated
	{
		logParsed, err := uniswap.ParsePoolCreated(*log)
		if err == nil {
			resp.MemePoolCreated = append(
				resp.MemePoolCreated,
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
				},
			)
		}
	}
	{
		logParsed, err := uniswapArb.ParsePool(*log)
		if err == nil {
			resp.MemePoolCreated = append(
				resp.MemePoolCreated,
				&UniswapPoolCreatedEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Token0:          logParsed.Token0.Hex(),
					Token1:          logParsed.Token1.Hex(),
					Pool:            logParsed.Pool.Hex(),
					Index:           log.Index,
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
			resp.MemeSwap = append(
				resp.MemeSwap,
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
					// FromAddress:     txFrom,
				},
			)
		}
	}

	position, err := basenonfungiblepositionmanager.NewNonfungiblePositionManager(log.Address, client)
	if err != nil {
		return err
	}
	// ParseIncreaseLiquidity
	{
		logParsed, err := position.ParseIncreaseLiquidity(*log)
		if err == nil {
			resp.MemeIncreaseLiquidity = append(
				resp.MemeIncreaseLiquidity,
				&UniswapPositionLiquidity{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Amount0:         logParsed.Amount0,
					Amount1:         logParsed.Amount1,
					TokenId:         logParsed.TokenId,
					Liquidity:       logParsed.Liquidity,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
				},
			)
		}
	}
	positionArb, err := arbitrumnonfungiblepositionmanager.NewNonfungiblePositionManager(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := positionArb.ParseIncreaseLiquidity(*log)
		if err == nil {
			resp.MemeIncreaseLiquidity = append(
				resp.MemeIncreaseLiquidity,
				&UniswapPositionLiquidity{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Amount0:         logParsed.Amount0,
					Amount1:         logParsed.Amount1,
					TokenId:         logParsed.TokenId,
					Liquidity:       logParsed.Liquidity,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
				},
			)
		}
	}

	// ParseDecreaseLiquidity
	{
		logParsed, err := position.ParseDecreaseLiquidity(*log)
		if err == nil {
			resp.MemeDecreaseLiquidity = append(
				resp.MemeDecreaseLiquidity,
				&UniswapPositionLiquidity{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Amount0:         logParsed.Amount0,
					Amount1:         logParsed.Amount1,
					TokenId:         logParsed.TokenId,
					Liquidity:       logParsed.Liquidity,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
				},
			)
		}
	}
	//
	{
		instance, err := agentshares.NewAgentShares(log.Address, client)
		if err != nil {
			return err
		}
		{
			logParsed, err := instance.ParseTrade(*log)
			if err == nil {
				resp.AgentSharesTrades = append(
					resp.AgentSharesTrades,
					logParsed,
				)
			}
		}
	}
	{
		instance, err := systempromptmanager.NewSystemPromptManager(log.Address, client)
		if err != nil {
			return err
		}
		{
			logParsed, err := instance.ParseNewToken(*log)
			if err == nil {
				resp.SystemPromptManagerNewTokens = append(
					resp.SystemPromptManagerNewTokens,
					logParsed,
				)
			}
		}
		{
			logParsed, err := instance.ParseAgentDataUpdate(*log)
			if err == nil {
				resp.SystemPromptManagerAgentDataUpdates = append(
					resp.SystemPromptManagerAgentDataUpdates,
					logParsed,
				)
			}
		}
		{
			logParsed, err := instance.ParseAgentURIUpdate(*log)
			if err == nil {
				resp.SystemPromptManagerAgentURIUpdates = append(
					resp.SystemPromptManagerAgentURIUpdates,
					logParsed,
				)
			}
		}
	}
	//realworld agent
	infra, err := realworldagent.NewERC20RealWorldAgent(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := infra.ParseExecutionRequested(*log)
		if err == nil {
			resp.RealWorldAgentExecutionRequested = append(
				resp.RealWorldAgentExecutionRequested,
				&RealWorldAgentExecutionRequested{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
					ActId:           logParsed.ActId,
					Uuid:            common.Bytes2Hex(logParsed.Uuid[:]),
					Creator:         logParsed.Creator.Hex(),
					Data:            string(logParsed.Data[:]),
				},
			)
		}
	}
	{
		instance, err := agentupgradeable.NewAgentUpgradeable(log.Address, client)
		if err != nil {
			return err
		}
		{
			logParsed, err := instance.ParseCodePointerCreated(*log)
			if err == nil {
				resp.CodePointerCreated = append(
					resp.CodePointerCreated,
					logParsed,
				)
			}
		}
	}
	// AgentFactory
	{
		instance, err := agentfactory.NewAgentFactory(log.Address, client)
		if err != nil {
			return err
		}
		{
			logParsed, err := instance.ParseAgentCreated(*log)
			if err == nil {
				resp.AgentCreated = append(resp.AgentCreated, logParsed)
			}
		}
	}
	{
		instance, err := vibetokenfactory.NewVibeTokenFactory(log.Address, client)
		if err != nil {
			return err
		}
		{
			logParsed, err := instance.ParseTokenDeployed(*log)
			if err == nil {
				resp.VibeTokenFactoryTokenDeployed = append(
					resp.VibeTokenFactoryTokenDeployed,
					logParsed,
				)
			}
		}
	}
	return nil
}

// ///////
type Erc20TokenTransferEventResp struct {
	NetworkID       uint64   `json:"network_id"`
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	From            string   `json:"from"`
	To              string   `json:"to"`
	Value           *big.Int `json:"value"`
	Index           uint     `json:"log_index"`
	BlockNumber     uint64   `json:"block_number"`
	TxIndex         uint     `json:"tx_index"`
}

type NftTransferEventResp struct {
	NetworkID       uint64   `json:"network_id"`
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	From            string   `json:"from"`
	To              string   `json:"to"`
	TokenId         *big.Int `json:"token_id"`
	Index           uint     `json:"log_index"`
	TxIndex         uint     `json:"tx_index"`
	BlockNumber     uint64   `json:"block_number"`
}

type ERC1155ransferEventResp struct {
	NetworkID       uint64 `json:"network_id"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Timestamp       uint64 `json:"timestamp"`
	Operator        string
	From            string
	To              string
	Id              *big.Int
	Value           *big.Int
}

type Erc20TokenEventResp struct {
	Transfer        []*Erc20TokenTransferEventResp `json:"transfer"`
	NftTransfer     []*NftTransferEventResp        `json:"nft_transfer"`
	ERC1155Transfer []*ERC1155ransferEventResp     `json:"erc1155_transfer"`
	LastBlockNumber int64                          `json:"last_block_number"`
}

func (c *Client) ScanTokenHolders(contracts []string, startBlock, endBlock int64) (*Erc20TokenEventResp, error) {
	resp := c.NewErc20TokenEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}

	var contractAddresses []common.Address
	if len(contracts) > 0 {
		contractAddresses = []common.Address{}
		for _, item := range contracts {
			contractAddresses = append(contractAddresses, helpers.HexToAddress(item))
		}
	}

	ctx := context.Background()
	lastBlock, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	if endBlock == 0 {
		endBlock = lastBlock.Number.Int64()
	}
	lastNumber := lastBlock.Number.Int64()
	if endBlock > lastNumber {
		endBlock = lastNumber
	}
	if startBlock > endBlock {
		return nil, nil
	}
	resp.LastBlockNumber = endBlock

	logs, err := client.FilterLogs(
		ctx,
		ethereum.FilterQuery{
			FromBlock: big.NewInt(startBlock),
			ToBlock:   big.NewInt(endBlock),
			Addresses: contractAddresses,
			Topics: [][]common.Hash{
				{
					common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
					common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
					common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"),
					common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		err = c.Erc20TokenEventResp(resp, &log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) NewErc20TokenEventResp() *Erc20TokenEventResp {
	return &Erc20TokenEventResp{}
}

func (c *Client) Erc20TokenEventResp(resp *Erc20TokenEventResp, log *types.Log) error {

	client, err := c.getClient()
	if err != nil {
		return err
	}
	chainId, err := c.ChainID()
	if err != nil {
		return err
	}
	blockTime := uint64(time.Now().Unix())
	erc20, err := erc20.NewErc20(log.Address, client)
	if err != nil {
		return err
	}
	// ParseTransfer
	{
		logParsed, err := erc20.ParseTransfer(*log)
		if err == nil {
			resp.Transfer = append(
				resp.Transfer,
				&Erc20TokenTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.From.Hex(),
					To:              logParsed.To.Hex(),
					Value:           logParsed.Value,
					Index:           log.Index,
					BlockNumber:     log.BlockNumber,
					TxIndex:         log.TxIndex,
				},
			)
		}
	}
	wbvm, err := wbvm.NewWBVM(log.Address, client)
	if err != nil {
		return err
	}
	// ParseTransfer
	{
		logParsed, err := wbvm.ParseDeposit(*log)
		if err == nil {
			resp.Transfer = append(
				resp.Transfer,
				&Erc20TokenTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            helpers.HexToAddress("0x0000000000000000000000000000000000000000").Hex(),
					To:              logParsed.Dst.Hex(),
					Value:           logParsed.Wad,
					Index:           log.Index,
					BlockNumber:     log.BlockNumber,
					TxIndex:         log.TxIndex,
				},
			)
		}
	}
	{
		logParsed, err := wbvm.ParseWithdrawal(*log)
		if err == nil {
			resp.Transfer = append(
				resp.Transfer,
				&Erc20TokenTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.Src.Hex(),
					To:              helpers.HexToAddress("0x0000000000000000000000000000000000000000").Hex(),
					Value:           logParsed.Wad,
					Index:           log.Index,
					BlockNumber:     log.BlockNumber,
					TxIndex:         log.TxIndex,
				},
			)
		}
	}
	//nft transfer
	erc721, err := erc721.NewErc721(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := erc721.ParseTransfer(*log)
		if err == nil {
			resp.NftTransfer = append(
				resp.NftTransfer,
				&NftTransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.From.Hex(),
					To:              logParsed.To.Hex(),
					TokenId:         logParsed.TokenId,
					Index:           log.Index,
					TxIndex:         log.TxIndex,
					BlockNumber:     log.BlockNumber,
				},
			)
		}
	}
	//erc1155 transfer
	erc1155, err := erc1155.NewERC1155(log.Address, client)
	if err != nil {
		return err
	}
	{
		logParsed, err := erc1155.ParseTransferBatch(*log)
		if err == nil {
			for index, _ := range logParsed.Ids {
				resp.ERC1155Transfer = append(
					resp.ERC1155Transfer,
					&ERC1155ransferEventResp{
						NetworkID:       chainId,
						TxHash:          log.TxHash.Hex(),
						ContractAddress: logParsed.Raw.Address.Hex(),
						Timestamp:       blockTime,
						From:            logParsed.From.Hex(),
						To:              logParsed.To.Hex(),
						Id:              logParsed.Ids[index],
						Value:           logParsed.Values[index],
					},
				)
			}
		}
	}
	{
		logParsed, err := erc1155.ParseTransferSingle(*log)
		if err == nil {
			resp.ERC1155Transfer = append(
				resp.ERC1155Transfer,
				&ERC1155ransferEventResp{
					NetworkID:       chainId,
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       blockTime,
					From:            logParsed.From.Hex(),
					To:              logParsed.To.Hex(),
					Id:              logParsed.Id,
					Value:           logParsed.Value,
				},
			)
		}
	}
	return nil
}

func (c *Client) Erc20EventsByTransaction(txHash string) (*Erc20TokenEventResp, error) {
	resp := c.NewErc20TokenEventResp()
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
		err = c.Erc20TokenEventResp(resp, log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
