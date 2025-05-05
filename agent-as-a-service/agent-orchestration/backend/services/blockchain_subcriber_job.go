package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/bridgeapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
)

func (s *Service) MemeEventsByTransaction(ctx context.Context, networkID uint64, txHash string) error {
	if !strings.EqualFold(txHash, common.HexToHash(txHash).Hex()) {
		return errs.NewError(errs.ErrBadRequest)
	}
	if networkID != models.TRON_CHAIN_ID {
		evmClient := s.GetEthereumClient(ctx, networkID)
		err := evmClient.WaitMined(txHash)
		if err != nil {
			return errs.NewError(err)
		}
	}
	ethClient := s.GetEthereumClient(ctx, networkID)
	eventResp, err := ethClient.EventsByTransaction(
		txHash,
	)
	if err != nil {
		return errs.NewError(err)
	}
	for i := 0; i < 5; i++ {
		err = s.MemeEventsByTransactionEventResp(
			ctx,
			networkID,
			eventResp,
			false,
		)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateScanBlockError(ctx context.Context, chainID uint, lastBlockError error) error {
	mapError := map[string]any{}
	if lastBlockError != nil {
		mapError["last_block_error"] = lastBlockError.Error()
		mapError["updated_at"] = time.Now()
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.BlockScanInfo{}).
		Where("id = ?", chainID).
		Updates(mapError).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateScanBlockErrorForNetwork(ctx context.Context, networkID uint64, lastBlockError error) error {
	mapError := map[string]any{}
	if lastBlockError != nil {
		mapError["last_block_error"] = lastBlockError.Error()
		mapError["updated_at"] = time.Now()
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.BlockScanInfo{}).
		Where("network_id = ?", networkID).
		Updates(mapError).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateScanBlockNumber(ctx context.Context, chainID uint, lastBlockEvent int64) error {
	mapError := map[string]any{}
	if lastBlockEvent > 0 {
		mapError["last_block_number"] = lastBlockEvent
		mapError["last_block_error"] = "OK"
		mapError["updated_at"] = time.Now()
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.BlockScanInfo{}).
		Where("id = ?", chainID).
		Updates(mapError).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) MemeEventsByTransactionEventResp(ctx context.Context, networkID uint64, eventResp *ethapi.BlockChainEventResp, forScan bool) error {
	var retErr error
	// handle transfer events
	{
		{
			poolMap := map[string]bool{}
			for _, event := range eventResp.Transfer {
				poolMap[strings.ToLower(event.ContractAddress)] = true
			}
			poolArr := []string{}
			for pool := range poolMap {
				poolArr = append(poolArr, pool)
			}
			poolMap = map[string]bool{}
			if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
				poolMap[strings.ToLower(strings.ToLower(s.conf.GetConfigKeyString(networkID, "eai_contract_address")))] = true
			}
			if len(poolArr) > 0 {
				memes, err := s.dao.FindMeme(
					daos.GetDBMainCtx(ctx),
					map[string][]any{
						"token_address in (?)": {poolArr},
					},
					map[string][]any{},
					[]string{},
					0,
					999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, meme := range memes {
					poolMap[strings.ToLower(meme.TokenAddress)] = true
				}
			}
			eventTransfers := []*ethapi.Erc20TokenTransferEventResp{}
			for _, event := range eventResp.Transfer {
				if poolMap[strings.ToLower(event.ContractAddress)] {
					eventTransfers = append(eventTransfers, event)
				}
			}
			eventResp.Transfer = eventTransfers
			eventResp.NftTransfer = []*ethapi.NftTransferEventResp{}
			eventResp.ERC1155Transfer = []*ethapi.ERC1155ransferEventResp{}
			for range 3 {
				err := s.TokenTransferEventsByTransactionV2(
					ctx,
					networkID,
					eventResp,
				)
				if err == nil {
					break
				}
			}
		}
		{
			poolMap := map[string]bool{}
			for _, event := range eventResp.Transfer {
				poolMap[strings.ToLower(event.To)] = true
				poolMap[strings.ToLower(event.From)] = true
			}
			poolArr := []string{}
			for pool := range poolMap {
				poolArr = append(poolArr, pool)
			}
			poolMap = map[string]bool{}
			if len(poolArr) > 0 {
				{
					agents, err := s.dao.FindAgentInfo(
						daos.GetDBMainCtx(ctx),
						map[string][]any{
							"eth_address in (?)": {poolArr},
						},
						map[string][]any{},
						[]string{},
						0,
						999999,
					)
					if err != nil {
						return errs.NewError(err)
					}
					for _, agent := range agents {
						poolMap[strings.ToLower(agent.ETHAddress)] = true
					}
				}
				{
					users, err := s.dao.FindUser(
						daos.GetDBMainCtx(ctx),
						map[string][]any{
							"eth_address in (?)": {poolArr},
						},
						map[string][]any{},
						[]string{},
						0,
						999999,
					)
					if err != nil {
						return errs.NewError(err)
					}
					for _, user := range users {
						poolMap[strings.ToLower(user.EthAddress)] = true
					}
				}
			}
			for _, event := range eventResp.Transfer {
				if poolMap[strings.ToLower(event.To)] || poolMap[strings.ToLower(event.From)] {
					err := s.CreateErc20TokenTransferEvent(ctx, networkID, event)
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
		}
		// {
		// 	poolMap := map[string]bool{}
		// 	for _, event := range eventResp.Transfer {
		// 		poolMap[strings.ToLower(event.To)] = true
		// 	}
		// 	poolArr := []string{}
		// 	for pool := range poolMap {
		// 		poolArr = append(poolArr, pool)
		// 	}
		// 	if len(poolArr) > 0 {
		// 		lps, err := s.dao.FindLaunchpad(
		// 			daos.GetDBMainCtx(ctx),
		// 			map[string][]any{
		// 				"address in (?)": {poolArr},
		// 			},
		// 			map[string][]any{},
		// 			[]string{},
		// 			0,
		// 			999999,
		// 		)
		// 		if err != nil {
		// 			return errs.NewError(err)
		// 		}
		// 		poolMap = map[string]bool{}
		// 		for _, lp := range lps {
		// 			poolMap[strings.ToLower(lp.Address)] = true
		// 		}
		// 	}
		// 	for _, event := range eventResp.Transfer {
		// 		if poolMap[strings.ToLower(event.To)] {
		// 			err := s.CreateErc20TokenTransferEventLaunchpad(ctx, networkID, event)
		// 			if err != nil {
		// 				return errs.NewError(err)
		// 			}
		// 		}
		// 	}
		// }
	}
	// handle vibe token factory token deployed events
	{
		for _, event := range eventResp.VibeTokenFactoryTokenDeployed {
			err := s.VibeTokenFactoryTokenDeployedEvent(ctx, networkID, event)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	// handle meme pool created events
	{
		var baseTokenETH, baseTokenEAI string
		if s.conf.ExistsedConfigKey(networkID, "weth9_contract_address") {
			baseTokenETH = strings.ToLower(s.conf.GetConfigKeyString(networkID, "weth9_contract_address"))
		}
		if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
			baseTokenEAI = strings.ToLower(s.conf.GetConfigKeyString(networkID, "eai_contract_address"))
		}
		poolMap := map[string]bool{}
		for _, event := range eventResp.MemePoolCreated {
			if !strings.EqualFold(event.Token0, baseTokenETH) && !strings.EqualFold(event.Token0, baseTokenEAI) {
				poolMap[strings.ToLower(event.Token0)] = true
			}
			if !strings.EqualFold(event.Token1, baseTokenETH) && !strings.EqualFold(event.Token1, baseTokenEAI) {
				poolMap[strings.ToLower(event.Token1)] = true
			}
		}
		poolArr := []string{}
		for pool := range poolMap {
			poolArr = append(poolArr, pool)
		}
		if len(poolArr) > 0 {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"token_address in (?)": {poolArr},
				},
				map[string][]any{},
				[]string{},
				0,
				999999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if len(memes) > 0 {
				poolMap := map[string]bool{}
				for _, meme := range memes {
					poolMap[strings.ToLower(meme.TokenAddress)] = true
				}
				for _, event := range eventResp.MemePoolCreated {
					if poolMap[strings.ToLower(event.Token0)] || poolMap[strings.ToLower(event.Token1)] {
						err = s.CreateMemePool(
							ctx, networkID, event, false,
						)
						if err != nil {
							retErr = errs.MergeError(retErr, err)
						}
					}
				}
			}
		}
	}
	// handle meme swap events
	{
		poolMap := map[string]bool{}
		for _, event := range eventResp.MemeSwap {
			poolMap[strings.ToLower(event.ContractAddress)] = true
		}
		poolArr := []string{}
		for pool := range poolMap {
			poolArr = append(poolArr, pool)
		}
		if len(poolArr) > 0 {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"pool in (?) or uniswap_pool in (?)": {poolArr, poolArr},
				},
				map[string][]any{},
				[]string{},
				0,
				999999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if len(memes) > 0 {
				poolMap := map[string]bool{}
				for _, meme := range memes {
					poolMap[strings.ToLower(meme.Pool)] = true
					poolMap[strings.ToLower(meme.UniswapPool)] = true
				}
				for _, event := range eventResp.MemeSwap {
					if poolMap[strings.ToLower(event.ContractAddress)] {
						err = s.CreateMemeTradeHistory(
							ctx, event,
						)
						if err != nil {
							retErr = errs.MergeError(retErr, err)
						}
					}
				}
			}
		}
	}
	if !forScan {
		// handle meme increase liquidity events
		for _, event := range eventResp.MemeIncreaseLiquidity {
			err := s.UpdateMemeLiquidityPosition(
				ctx, networkID, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	// handle system prompt manager new tokens events
	{
		for _, event := range eventResp.SystemPromptManagerNewTokens {
			err := s.SystemPromptManagerNewTokenEvent(
				ctx, networkID, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		for _, event := range eventResp.SystemPromptManagerAgentDataUpdates {
			err := s.SystemPromptManagerAgentDataUpdateEvent(
				ctx, networkID, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		for _, event := range eventResp.SystemPromptManagerAgentURIUpdates {
			err := s.SystemPromptManagerAgentURIUpdateEvent(
				ctx, networkID, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	// handle agent created events
	{
		for _, event := range eventResp.AgentCreated {
			err := s.AgentFactoryAgentCreatedEvent(ctx, networkID, event)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	// handle agent upgradeable code pointer created events
	{
		for _, event := range eventResp.CodePointerCreated {
			err := s.HandleAgentUpgradeableCodePointerCreated(ctx, event)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	return retErr
}

func (s *Service) TokenTransferEventsByTransactionV2(ctx context.Context, networkID uint64, eventResp *ethapi.BlockChainEventResp) error {
	var retErr, err error
	if eventResp != nil {
		{
			changeBalance := map[string]bool{}
			for _, v := range eventResp.Transfer {
				changeBalance[fmt.Sprintf("%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.From))] = true
				changeBalance[fmt.Sprintf("%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.To))] = true
			}
			err = s.CreateErc20TransferEvent(ctx, networkID, changeBalance)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		for _, event := range eventResp.NftTransfer {
			err = s.CreateErc721TransferEvent(
				ctx, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		{
			changeBalance1155 := map[string]bool{}
			for _, v := range eventResp.ERC1155Transfer {
				changeBalance1155[fmt.Sprintf("%s|%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.From), v.Id.Text(10))] = true
				changeBalance1155[fmt.Sprintf("%s|%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.To), v.Id.Text(10))] = true
			}
			err = s.CreateErc1155TransferEvent(ctx, networkID, changeBalance1155)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}

	return retErr
}

func (s *Service) ScanTokenTransferTxHash(ctx context.Context, networkID uint64, hash string) error {
	eventResp, err := s.GetEthereumClient(ctx, networkID).Erc20EventsByTransaction(hash)
	if err != nil {
		return errs.NewError(err)
	}
	err = s.TokenTransferEventsByTransaction(ctx, networkID, eventResp)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) TokenTransferEventsByTransaction(ctx context.Context, networkID uint64, eventResp *ethapi.Erc20TokenEventResp) error {
	var retErr, err error
	if eventResp != nil {
		{
			changeBalance := map[string]bool{}
			for _, v := range eventResp.Transfer {
				changeBalance[fmt.Sprintf("%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.From))] = true
				changeBalance[fmt.Sprintf("%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.To))] = true
				err = s.CreateErc20TokenTransferEvent(ctx, networkID, v)
				if err != nil {
					retErr = errs.MergeError(retErr, err)
				}
			}
			err = s.CreateErc20TransferEvent(ctx, networkID, changeBalance)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		for _, event := range eventResp.NftTransfer {
			err = s.CreateErc721TransferEvent(
				ctx, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		{
			changeBalance1155 := map[string]bool{}
			for _, v := range eventResp.ERC1155Transfer {
				changeBalance1155[fmt.Sprintf("%s|%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.From), v.Id.Text(10))] = true
				changeBalance1155[fmt.Sprintf("%s|%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.To), v.Id.Text(10))] = true
			}
			err = s.CreateErc1155TransferEvent(ctx, networkID, changeBalance1155)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	return retErr
}

func (s *Service) ErcTransferEventsByTransactionV2(ctx context.Context, networkID uint64, eventResp *ethapi.BlockChainEventResp) error {
	var retErr, err error
	if eventResp != nil {
		{
			changeBalance := map[string]bool{}
			for _, v := range eventResp.Transfer {
				changeBalance[fmt.Sprintf("%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.From))] = true
				changeBalance[fmt.Sprintf("%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.To))] = true
				err = s.CreateErc20TokenTransferEvent(ctx, networkID, v)
				if err != nil {
					retErr = errs.MergeError(retErr, err)
				}
			}
			err = s.CreateErc20TransferEvent(ctx, networkID, changeBalance)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		for _, event := range eventResp.NftTransfer {
			err = s.CreateErc721TransferEvent(
				ctx, event,
			)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
		{
			changeBalance1155 := map[string]bool{}
			for _, v := range eventResp.ERC1155Transfer {
				changeBalance1155[fmt.Sprintf("%s|%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.From), v.Id.Text(10))] = true
				changeBalance1155[fmt.Sprintf("%s|%s|%s", strings.ToLower(v.ContractAddress), strings.ToLower(v.To), v.Id.Text(10))] = true
			}
			err = s.CreateErc1155TransferEvent(ctx, networkID, changeBalance1155)
			if err != nil {
				retErr = errs.MergeError(retErr, err)
			}
		}
	}
	return retErr
}

func (s *Service) CreateErc20TransferEvent(ctx context.Context, networkID uint64, changeBalance map[string]bool) error {
	ethClient := s.GetEthereumClient(ctx, networkID)
	updateStm := "INSERT INTO erc20_holders (network_id, contract_address, address, balance, created_at) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE balance = ?, updated_at = ?"
	now := time.Now()
	for k := range changeBalance {
		conAddress := strings.Split(k, "|")[0]
		userAddress := strings.Split(k, "|")[1]
		if userAddress != "0x0000000000000000000000000000000000000000" {
			var err error
			var balance *big.Int
			if strings.EqualFold(conAddress, "0x000000000000000000000000000000000000800A") {
				balance, err = ethClient.Balance(userAddress)
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				chainId, err := ethClient.ChainID()
				if err != nil {
					return err
				}
				if chainId == models.TRON_CHAIN_ID {
					balance, err = s.trxApi.Trc20Balance(conAddress, userAddress)
					if err != nil {
						fmt.Println(err.Error())
					}
				} else {
					balance, err = ethClient.Erc20Balance(conAddress, userAddress)
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}
			if balance == nil {
				balance = big.NewInt(0)
			}
			fBalance := models.ConvertWeiToBigFloat(balance, 18)
			valueArgs := []any{}
			valueArgs = append(valueArgs, networkID)
			valueArgs = append(valueArgs, conAddress)
			valueArgs = append(valueArgs, userAddress)
			valueArgs = append(valueArgs, numeric.NewBigFloatFromFloat(fBalance))
			valueArgs = append(valueArgs, now)
			valueArgs = append(valueArgs, numeric.NewBigFloatFromFloat(fBalance))
			valueArgs = append(valueArgs, now)
			err = daos.GetDBMainCtx(ctx).Exec(updateStm, valueArgs...).Error
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) CreateErc721TransferEvent(ctx context.Context, event *ethapi.NftTransferEventResp) error {
	ethClient := s.GetEthereumClient(ctx, event.NetworkID)
	updateStm := "INSERT INTO erc721_holders (network_id, contract_address, token_id, owner_address, created_at) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE owner_address=?, updated_at = ?"
	now := time.Now()
	owner, err := ethClient.NftOwnerOf(event.ContractAddress, fmt.Sprintf("%d", uint(event.TokenId.Uint64())))
	if err != nil {
		if strings.Contains(err.Error(), "execution reverted") {
			owner = "0x0000000000000000000000000000000000000000"
		} else {
			return errs.NewError(err)
		}
	}
	valueArgs := []any{}
	valueArgs = append(valueArgs, event.NetworkID)
	valueArgs = append(valueArgs, strings.ToLower(event.ContractAddress))
	valueArgs = append(valueArgs, uint(event.TokenId.Uint64()))
	valueArgs = append(valueArgs, strings.ToLower(owner))
	valueArgs = append(valueArgs, now)
	valueArgs = append(valueArgs, strings.ToLower(owner))
	valueArgs = append(valueArgs, now)
	err = daos.GetDBMainCtx(ctx).Exec(updateStm, valueArgs...).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateErc1155TransferEvent(ctx context.Context, networkID uint64, changeBalance map[string]bool) error {
	ethClient := s.GetEthereumClient(ctx, networkID)
	updateStm := "INSERT INTO erc1155_holders (network_id, contract_address, address, token_id, balance, created_at) VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE balance = ?, updated_at = ?"
	now := time.Now()
	for k := range changeBalance {
		err := func() error {
			conAddress := strings.Split(k, "|")[0]
			userAddress := strings.Split(k, "|")[1]
			tokenId := strings.Split(k, "|")[2]
			if userAddress != "0x0000000000000000000000000000000000000000" {
				balance, err := ethClient.Erc1155Balance(conAddress, tokenId, userAddress)
				if err != nil {
					return errs.NewError(err)
				}
				fBalance := models.ConvertWeiToBigFloat(balance, 18)
				valueArgs := []any{}
				valueArgs = append(valueArgs, networkID)
				valueArgs = append(valueArgs, conAddress)
				valueArgs = append(valueArgs, userAddress)
				valueArgs = append(valueArgs, tokenId)
				valueArgs = append(valueArgs, numeric.NewBigFloatFromFloat(fBalance))
				valueArgs = append(valueArgs, now)
				valueArgs = append(valueArgs, numeric.NewBigFloatFromFloat(fBalance))
				valueArgs = append(valueArgs, now)
				err = daos.GetDBMainCtx(ctx).Exec(updateStm, valueArgs...).Error
				if err != nil {
					return errs.NewError(err)
				}
			}
			return nil
		}()
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func (s *Service) ImageHubSubscriptionPriceUpdatedEvent(ctx context.Context, networkID uint64, event *ethapi.ImageHubSubscriptionPriceUpdated) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			imageHubAddress := s.conf.GetConfigKeyString(networkID, "image_hub_contract_address")
			if !strings.EqualFold(event.ContractAddress, imageHubAddress) {
				return nil
			}
			user, err := s.GetUser(
				tx,
				networkID,
				event.Creator,
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			priceUpdated := helpers.GetTimeIndex(uint(event.BlockNumber), event.TxIndex, event.Index)
			switch event.Duration.Uint64() {
			case models.DURATION_30D:
				{
					if user.Price30dUpdated < priceUpdated {
						err = tx.Model(user).
							Where("price30d_updated < ?", priceUpdated).
							Updates(
								map[string]any{
									"price30d":         numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Price, 18)),
									"price30d_updated": priceUpdated,
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				}
			case models.DURATION_90D:
				{
					if user.Price90dUpdated < priceUpdated {
						err = tx.Model(user).
							Where("price90d_updated < ?", priceUpdated).
							Updates(
								map[string]any{
									"price90d":         numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Price, 18)),
									"price90d_updated": priceUpdated,
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateErc20TokenTransferEvent(ctx context.Context, networkID uint64, event *ethapi.Erc20TokenTransferEventResp) error {
	if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
		contractAddress := strings.ToLower(event.ContractAddress)
		eaiAddress := s.conf.GetConfigKeyString(networkID, "eai_contract_address")
		toAddress := strings.ToLower(event.To)
		fromAddress := strings.ToLower(event.From)
		var istelePost bool
		if !strings.EqualFold(toAddress, models.ETH_ZERO_ADDRESS) && strings.EqualFold(contractAddress, eaiAddress) {
			var agent *models.AgentInfo
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					var err error
					{
						agent, err = s.dao.FirstAgentInfo(
							tx,
							map[string][]any{
								"eth_address = ?": {toAddress},
							},
							map[string][]any{},
							[]string{},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agent != nil {
							eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxHash, event.Index)
							m, err := s.dao.FirstAgentEaiTopup(
								tx,
								map[string][]any{
									"event_id = ?": {eventId},
								},
								map[string][]any{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							if m == nil {
								m = &models.AgentEaiTopup{
									NetworkID:      networkID,
									EventId:        eventId,
									AgentInfoID:    agent.ID,
									Type:           models.AgentEaiTopupTypeDeposit,
									DepositAddress: strings.ToLower(event.From),
									DepositTxHash:  event.TxHash,
									Amount:         numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Value, 18)),
									Status:         models.AgentEaiTopupStatusDone,
									ToAddress:      agent.ETHAddress,
								}
								if (m.NetworkID == models.ABSTRACT_TESTNET_CHAIN_ID ||
									m.NetworkID == models.MONAD_TESTNET_CHAIN_ID ||
									m.NetworkID == models.MEGAETH_TESTNET_CHAIN_ID ||
									m.NetworkID == models.BASE_SEPOLIA_CHAIN_ID) && m.NetworkID != agent.NetworkID {
									m.Status = models.AgentEaiTopupStatusCancelled
								}
								err = s.dao.Create(
									tx,
									m,
								)
								if err != nil {
									return errs.NewError(err)
								}
								if m.Status == models.AgentEaiTopupStatusDone {
									err = tx.Model(agent).
										UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", m.Amount)).
										UpdateColumn("eai_wallet_balance", gorm.Expr("eai_wallet_balance + ?", m.Amount)).
										Error
									if err != nil {
										return errs.NewError(err)
									}
								}
								istelePost = true
							}
						}
					}
					{
						agent, err = s.dao.FirstAgentInfo(
							tx,
							map[string][]any{
								"eth_address = ?": {fromAddress},
							},
							map[string][]any{},
							[]string{},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agent != nil {
							eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxHash, event.Index)
							m, err := s.dao.FirstAgentEaiTopup(
								tx,
								map[string][]any{
									"event_id = ?": {eventId},
								},
								map[string][]any{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							if m == nil {
								m = &models.AgentEaiTopup{
									NetworkID:      networkID,
									EventId:        eventId,
									AgentInfoID:    agent.ID,
									Type:           models.AgentEaiTopupTypeTransfer,
									DepositAddress: strings.ToLower(event.From),
									DepositTxHash:  event.TxHash,
									Amount:         numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Value, 18)),
									Status:         models.AgentEaiTopupStatusDone,
									ToAddress:      agent.ETHAddress,
								}
								err = s.dao.Create(
									tx,
									m,
								)
								if err != nil {
									return errs.NewError(err)
								}
							}
						}
					}
					{
						switch networkID {
						case models.BASE_CHAIN_ID:
							{
								err = s.UtilityTwitterHandleDeposit(tx, networkID, event)
								if err != nil {
									return errs.NewError(err)
								}
							}
						}
					}
					// {
					// 	err = s.LaunchpadErc20TokenTransferEvent(tx, networkID, event)
					// 	if err != nil {
					// 		return errs.NewError(err)
					// 	}
					// }
					// {
					// 	switch networkID {
					// 	case models.ETHEREUM_CHAIN_ID,
					// 		models.BASE_CHAIN_ID:
					// 		{
					// 			eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxHash, event.Index)
					// 			s.ProcessDeposit(ctx, networkID, eventId, event.TxHash, toAddress, event.Value)
					// 			err = s.LaunchpadErc20TokenTransferEvent(tx, networkID, event)
					// 			if err != nil {
					// 				return errs.NewError(err)
					// 			}
					// 		}
					// 	}
					// }
					{
						user, err := s.dao.FirstUser(
							tx,
							map[string][]any{
								"eth_address = ?": {toAddress},
							},
							map[string][]any{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if user != nil {
							eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxHash, event.Index)
							m, err := s.dao.FirstUserTransaction(
								tx,
								map[string][]any{
									"event_id = ?": {eventId},
								},
								map[string][]any{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							if m == nil {
								m = &models.UserTransaction{
									NetworkID:   networkID,
									EventId:     eventId,
									UserID:      user.ID,
									Type:        models.UserTransactionTypeDeposit,
									FromAddress: strings.ToLower(event.From),
									ToAddress:   strings.ToLower(event.To),
									TxHash:      event.TxHash,
									Amount:      numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Value, 18)),
									Status:      models.UserTransactionStatusDone,
								}
								if m.NetworkID == models.ABSTRACT_TESTNET_CHAIN_ID ||
									m.NetworkID == models.MONAD_TESTNET_CHAIN_ID ||
									m.NetworkID == models.MEGAETH_TESTNET_CHAIN_ID ||
									m.NetworkID == models.BASE_SEPOLIA_CHAIN_ID {
									m.Status = models.UserTransactionStatusCancelled
								}
								err = s.dao.Create(
									tx,
									m,
								)
								if err != nil {
									return errs.NewError(err)
								}
								if m.Status == models.UserTransactionStatusDone {
									err = tx.Model(user).
										UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", m.Amount)).
										Error
									if err != nil {
										return errs.NewError(err)
									}
								}
							}
						}
					}
					return nil
				},
			)
			if err != nil {
				return errs.NewError(err)
			}
			if istelePost && agent != nil {
				if event.Value.Cmp(common.Big0) > 0 {
					go s.AgentTeleAlertByID(ctx, agent.ID, event.TxHash, models.ConvertWeiToBigFloat(event.Value, 18), networkID)
				}
			}
		}
	}
	return nil
}

func (s *Service) CreateErc20TokenTransferEventLaunchpad(ctx context.Context, networkID uint64, event *ethapi.Erc20TokenTransferEventResp) error {
	if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
		contractAddress := strings.ToLower(event.ContractAddress)
		eaiAddress := s.conf.GetConfigKeyString(networkID, "eai_contract_address")
		toAddress := strings.ToLower(event.To)
		if !strings.EqualFold(toAddress, models.ETH_ZERO_ADDRESS) && strings.EqualFold(contractAddress, eaiAddress) {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					err := s.LaunchpadErc20TokenTransferEvent(tx, networkID, event)
					if err != nil {
						return errs.NewError(err)
					}
					return nil
				},
			)
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) CreateSolanaTokenTransferEvent(ctx context.Context, networkID uint64, event *bridgeapi.SolanaEAITxResp) error {
	switch networkID {
	case models.SOLANA_CHAIN_ID:
		{
			toAddress := event.DepositNativeAddress
			contractAddress := event.Token
			if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
				eaiAddress := s.conf.GetConfigKeyString(networkID, "eai_contract_address")
				if strings.EqualFold(contractAddress, eaiAddress) {
					var agent *models.AgentInfo
					err := daos.WithTransaction(
						daos.GetDBMainCtx(ctx),
						func(tx *gorm.DB) error {
							var err error
							{
								agent, err = s.dao.FirstAgentInfo(
									tx,
									map[string][]any{
										"sol_address = ?": {toAddress},
									},
									map[string][]any{},
									[]string{},
								)
								if err != nil {
									return errs.NewError(err)
								}
								if agent != nil {
									eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxReceivedDeposit, 0)
									m, err := s.dao.FirstAgentEaiTopup(
										tx,
										map[string][]any{
											"event_id = ?": {eventId},
										},
										map[string][]any{},
										[]string{},
									)
									if err != nil {
										return errs.NewError(err)
									}
									if m == nil {
										m = &models.AgentEaiTopup{
											NetworkID:      networkID,
											EventId:        eventId,
											AgentInfoID:    agent.ID,
											Type:           models.AgentEaiTopupTypeDeposit,
											DepositAddress: event.DepositNativeAddress,
											DepositTxHash:  event.TxReceivedDeposit,
											Amount:         numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(&event.Amount.Int, 6)),
											Status:         models.AgentEaiTopupStatusDone,
											ToAddress:      agent.ETHAddress,
										}
										err = s.dao.Create(
											tx,
											m,
										)
										if err != nil {
											return errs.NewError(err)
										}
										err = tx.Model(agent).
											UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", m.Amount)).
											UpdateColumn("eai_wallet_balance", gorm.Expr("eai_wallet_balance + ?", m.Amount)).
											Error
										if err != nil {
											return errs.NewError(err)
										}
									}
								}
							}
							{
								user, err := s.dao.FirstUser(
									tx,
									map[string][]any{
										"sol_address = ?": {toAddress},
									},
									map[string][]any{},
									false,
								)
								if err != nil {
									return errs.NewError(err)
								}
								if user != nil {
									eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxReceivedDeposit, 0)
									m, err := s.dao.FirstUserTransaction(
										tx,
										map[string][]any{
											"event_id = ?": {eventId},
										},
										map[string][]any{},
										[]string{},
									)
									if err != nil {
										return errs.NewError(err)
									}
									if m == nil {
										m = &models.UserTransaction{
											NetworkID:   networkID,
											EventId:     eventId,
											UserID:      user.ID,
											Type:        models.UserTransactionTypeDeposit,
											FromAddress: event.DepositNativeAddress,
											ToAddress:   toAddress,
											TxHash:      event.TxReceivedDeposit,
											Amount:      numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(&event.Amount.Int, 6)),
											Status:      models.UserTransactionStatusDone,
										}
										err = s.dao.Create(
											tx,
											m,
										)
										if err != nil {
											return errs.NewError(err)
										}
										if m.Status == models.UserTransactionStatusDone {
											err = tx.Model(user).
												UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", m.Amount)).
												Error
											if err != nil {
												return errs.NewError(err)
											}
										}
									}
								}
							}
							return nil
						},
					)
					if err != nil {
						return errs.NewError(err)
					}
					if agent != nil {
						if event.Amount.Int.Cmp(common.Big0) > 0 {
							go s.AgentTeleAlertByID(ctx, agent.ID, event.TxReceivedDeposit, models.ConvertWeiToBigFloat(&event.Amount.Int, 6), networkID)
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) DeleteFilterAddrs(ctx context.Context, networkID uint64) error {
	for range 5 {
		err := s.DeleteRedisCachedWithKey(fmt.Sprintf("GetFilterAddrs_%d", networkID))
		if err == nil {
			break
		}
	}
	return nil
}

func (s *Service) GetFilterAddrs(ctx context.Context, networkID uint64) ([]string, error) {
	addrs := []string{}
	err := s.RedisCached(
		fmt.Sprintf("GetFilterAddrsV1_%d", networkID),
		true,
		5*time.Minute,
		&addrs,
		func() (any, error) {
			addrs := []string{}
			memeAddresses, err := s.dao.FindAllMemeTokenAddress(
				daos.GetDBMainCtx(ctx),
				networkID,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}
			addrs = append(addrs, memeAddresses...)
			if s.conf.ExistsedConfigKey(networkID, "meme_position_mamanger_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "meme_position_mamanger_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "meme_factory_contract_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "meme_factory_contract_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "uniswap_position_mamanger_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "uniswap_position_mamanger_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "uniswap_factory_contract_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "uniswap_factory_contract_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "eai_contract_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "agent_contract_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "agent_contract_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "order_payment_contract_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "order_payment_contract_address"))
			}
			if s.conf.ExistsedConfigKey(networkID, "agent_factory_address") {
				address := s.conf.GetConfigKeyString(networkID, "agent_factory_address")
				if address != "" {
					addrs = append(addrs, address)
				}
			}
			if s.conf.ExistsedConfigKey(networkID, "vibe_token_factory_address") {
				addrs = append(addrs, s.conf.GetConfigKeyString(networkID, "vibe_token_factory_address"))
			}
			agentAddresses, err := s.dao.FindAllAgentAddress(
				daos.GetDBMainCtx(ctx),
				networkID,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}
			addrs = append(addrs, agentAddresses...)
			return addrs, nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return addrs, nil
}

func (s *Service) ScanEventsByChain(ctx context.Context, networkID uint64) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("ScanEventsByChain_%d", networkID),
		func() error {
			if networkID > 0 {
				switch networkID {
				case models.SOLANA_CHAIN_ID:
					{
						err := func(networkID uint64) error {
							chain, err := s.dao.FirstBlockScanInfo(
								daos.GetDBMainCtx(ctx),
								map[string][]any{
									"type = ?":              {"solana"},
									"network_id = ?":        {networkID},
									"last_block_number > 0": {},
								},
								map[string][]any{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							if chain == nil {
								chain = &models.BlockScanInfo{
									Type:      "solana",
									NetworkID: networkID,
									NumBlocks: 100,
								}
								lastBlockNumber, err := s.blockchainUtils.SolanaBlockheight()
								if err != nil {
									return errs.NewError(err)
								}
								if lastBlockNumber <= 0 {
									lastBlockNumber = 1
								}
								chain.LastBlockNumber = lastBlockNumber
								err = s.dao.Create(
									daos.GetDBMainCtx(ctx),
									chain,
								)
								if err != nil {
									return errs.NewError(err)
								}
							}
							if !chain.Enabled || chain.LastBlockNumber == 0 {
								return nil
							}
							lastBlockNumber := chain.LastBlockNumber
							err = func() error {
								txs, err := s.bridgeAPI.GetSolanaEAITxs(uint64(chain.LastBlockNumber))
								if err != nil {
									return errs.NewError(err)
								}
								for _, tx := range txs {
									err = s.CreateSolanaTokenTransferEvent(ctx, chain.NetworkID, tx)
									if err != nil {
										return errs.NewError(err)
									}
									if tx.Block > int(lastBlockNumber) {
										lastBlockNumber = int64(tx.Block)
									}
								}
								return nil
							}()
							if err != nil {
								return errs.NewError(err)
							}
							if lastBlockNumber > 0 {
								err = s.UpdateScanBlockNumber(ctx, chain.ID, lastBlockNumber)
								if err != nil {
									return errs.NewError(err)
								}
							}
							return nil
						}(networkID)
						if err != nil {
							_ = s.UpdateScanBlockErrorForNetwork(ctx, networkID, err)
							return errs.NewError(err)
						}
					}
				default:
					{
						err := func(networkID uint64) error {
							for {
								chain, err := s.dao.FirstBlockScanInfo(
									daos.GetDBMainCtx(ctx),
									map[string][]any{
										"type = ?":              {"evm"},
										"network_id = ?":        {networkID},
										"last_block_number > 0": {},
									},
									map[string][]any{},
									[]string{},
								)
								if err != nil {
									return errs.NewError(err)
								}
								ethClient := s.GetEthereumClient(ctx, networkID)
								if chain == nil {
									chain = &models.BlockScanInfo{
										Type:      "evm",
										NetworkID: networkID,
										NumBlocks: 100,
									}
									lastBlockNumber, err := ethClient.GetLastBlockNumber()
									if err != nil {
										return errs.NewError(err)
									}
									if lastBlockNumber <= 0 {
										lastBlockNumber = 1
									}
									chain.LastBlockNumber = lastBlockNumber
									err = s.dao.Create(
										daos.GetDBMainCtx(ctx),
										chain,
									)
									if err != nil {
										return errs.NewError(err)
									}
								}
								if !chain.Enabled || chain.LastBlockNumber == 0 {
									break
								}
								addrs, err := s.GetFilterAddrs(ctx, chain.NetworkID)
								if err != nil {
									return errs.NewError(err)
								}
								startBlocks := chain.LastBlockNumber + 1
								endBlocks := (chain.LastBlockNumber + chain.NumBlocks - 1)
								eventResp, err := ethClient.ScanEvents(addrs, startBlocks, endBlocks)
								if err != nil {
									return errs.NewError(err)
								}
								if eventResp != nil {
									err = s.MemeEventsByTransactionEventResp(ctx, chain.NetworkID, eventResp, true)
									if err != nil {
										return errs.NewError(err)
									} else {
										lastBlockNumber := endBlocks
										if endBlocks > eventResp.LastBlockNumber {
											lastBlockNumber = eventResp.LastBlockNumber
										}
										err = s.UpdateScanBlockNumber(ctx, chain.ID, lastBlockNumber)
										if err != nil {
											return errs.NewError(err)
										}
										if endBlocks >= eventResp.LastBlockNumber {
											break
										}
									}
								} else {
									break
								}
							}
							return nil
						}(networkID)
						if err != nil {
							_ = s.UpdateScanBlockErrorForNetwork(ctx, networkID, err)
							return errs.NewError(err)
						}
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

// func (s *Service) JobScanEventsByChain(ctx context.Context) error {
// 	err := s.JobRunCheck(
// 		ctx, "JobScanEventsByChain",
// 		func() error {
// 			var retErr error
// 			var networkIDs []uint64
// 			for networkIDStr := range s.conf.Networks {
// 				networkID, _ := strconv.ParseUint(networkIDStr, 10, 64)
// 				networkIDs = append(networkIDs, networkID)
// 			}
// 			// random networkIDs
// 			rand.Shuffle(len(networkIDs), func(i, j int) {
// 				networkIDs[i], networkIDs[j] = networkIDs[j], networkIDs[i]
// 			})
// 			for _, networkID := range networkIDs {
// 				err := s.ScanEventsByChain(ctx, networkID)
// 				if err != nil {
// 					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, networkID))
// 				}
// 			}
// 			return retErr
// 		},
// 	)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	return nil
// }
