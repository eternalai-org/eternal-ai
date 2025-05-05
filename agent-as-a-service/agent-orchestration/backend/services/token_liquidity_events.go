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
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentshares"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/vibetokenfactory"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetMemeBaseToken(tx *gorm.DB, networkID uint64, baseSymbol string) string {
	var memeBaseTokenAddress string
	if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
		switch baseSymbol {
		case string(models.BaseTokenSymbolETH):
			{
				memeBaseTokenAddress = strings.ToLower(s.conf.GetConfigKeyString(networkID, "weth9_contract_address"))
			}
		case string(models.BaseTokenSymbolEAI):
			{
				memeBaseTokenAddress = strings.ToLower(s.conf.GetConfigKeyString(networkID, "eai_contract_address"))
			}
		}
	}
	return memeBaseTokenAddress
}

func (s *Service) GetMemeBaseTokenByTokens01(tx *gorm.DB, networkID uint64, token0, token1 string) string {
	var memeBaseTokenAddress string
	if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
		baseTokenETH := strings.ToLower(s.conf.GetConfigKeyString(networkID, "weth9_contract_address"))
		baseTokenEAI := strings.ToLower(s.conf.GetConfigKeyString(networkID, "eai_contract_address"))
		if strings.EqualFold(token0, strings.ToLower(baseTokenETH)) ||
			strings.EqualFold(token1, strings.ToLower(baseTokenETH)) {
			memeBaseTokenAddress = strings.ToLower(baseTokenETH)
		} else if strings.EqualFold(token0, strings.ToLower(baseTokenEAI)) ||
			strings.EqualFold(token1, strings.ToLower(baseTokenEAI)) {
			memeBaseTokenAddress = strings.ToLower(baseTokenEAI)
		} else {
			return ""
		}
	}
	return memeBaseTokenAddress
}

func (s *Service) CreateMemePool(ctx context.Context, networkID uint64, event *ethapi.UniswapPoolCreatedEventResp, syncInfo bool) error {
	tokenAddress := ""
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if s.conf.ExistsedConfigKey(networkID, "memeswap_factory_contract_address") {
				swapFactoryAddr := s.conf.GetConfigKeyString(networkID, "memeswap_factory_contract_address")
				if strings.EqualFold(swapFactoryAddr, event.ContractAddress) {
					token, err := s.dao.FirstMeme(tx,
						map[string][]any{
							"pool = ?": {strings.ToLower(event.Pool)},
						},
						map[string][]any{},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if token == nil {
						memeBaseTokenAddress := s.GetMemeBaseTokenByTokens01(tx, networkID, event.Token0, event.Token1)
						if memeBaseTokenAddress == "" {
							return nil
						}
						baseIndex := int(1)
						tokenAddress = strings.ToLower(event.Token0)
						if strings.EqualFold(event.Token0, memeBaseTokenAddress) {
							tokenAddress = strings.ToLower(event.Token1)
							baseIndex = 0
						}
						token, err = s.dao.FirstMeme(tx,
							map[string][]any{
								"token_address = ?": []any{tokenAddress},
							},
							map[string][]any{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if token != nil {
							if token.AddPool1TxHash == "" {
								return errs.NewError(errs.ErrBadRequest)
							}
							token, err = s.dao.FirstMemeByID(tx, token.ID, map[string][]any{}, true)
							if err != nil {
								return errs.NewError(err)
							}
							if token.AddPool1TxHash == "" {
								return errs.NewError(errs.ErrBadRequest)
							}
							if strings.EqualFold(token.AddPool1TxHash, event.TxHash) {
								token.Status = models.MemeStatusAddPoolLevel1
								token.Pool = strings.ToLower(event.Pool)
								token.Token0Address = strings.ToLower(event.Token0)
								token.Token1Address = strings.ToLower(event.Token1)
								token.BaseTokenIndex = baseIndex
								token.ZeroForOne = token.TokenAddress != strings.ToLower(event.Token0)
								if token.PriceUsd.Cmp(big.NewFloat(0)) == 0 {
									basePrice := s.GetTokenMarketPrice(tx, token.BaseTokenSymbol)
									token.PriceUsd = numeric.BigFloat{*models.MulBigFloats(&token.Price.Float, basePrice)}
								}
								err = s.dao.Save(tx, token)
								if err != nil {
									return errs.NewError(err)
								}
							}
						}
					}
				}
			}

			if s.conf.ExistsedConfigKey(networkID, "uniswap_factory_contract_address") {
				swapFactoryAddr := s.conf.GetConfigKeyString(networkID, "uniswap_factory_contract_address")
				if strings.EqualFold(swapFactoryAddr, event.ContractAddress) {
					token, err := s.dao.FirstMeme(tx,
						map[string][]any{
							"uniswap_pool = ?": []any{strings.ToLower(event.Pool)},
						},
						map[string][]any{},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if token == nil || token.Status == models.MemeStatusCreated || token.Status == models.MemeStatusRemovePoolLelve1 {
						memeBaseTokenAddress := s.GetMemeBaseTokenByTokens01(tx, networkID, event.Token0, event.Token1)
						if memeBaseTokenAddress == "" {
							return nil
						}
						baseIndex := int(1)
						tokenAddress = strings.ToLower(event.Token0)
						if strings.EqualFold(event.Token0, memeBaseTokenAddress) {
							tokenAddress = strings.ToLower(event.Token1)
							baseIndex = 0
						}
						token, err = s.dao.FirstMeme(tx,
							map[string][]any{
								"token_address = ?": []any{tokenAddress},
							},
							map[string][]any{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if token != nil {
							if token.AddPool2TxHash == "" {
								token.AddPool2TxHash = event.TxHash
							}
							token, err = s.dao.FirstMemeByID(tx, token.ID, map[string][]any{}, true)
							if err != nil {
								return errs.NewError(err)
							}
							if strings.EqualFold(token.AddPool2TxHash, event.TxHash) {
								token.Status = models.MemeStatusAddPoolLevel2
								token.UniswapPool = strings.ToLower(event.Pool)
								token.Token0Address = strings.ToLower(event.Token0)
								token.Token1Address = strings.ToLower(event.Token1)
								token.BaseTokenIndex = baseIndex
								token.ZeroForOne = token.TokenAddress != strings.ToLower(event.Token0)
								if token.PriceUsd.Cmp(big.NewFloat(0)) == 0 {
									basePrice := s.GetTokenMarketPrice(tx, token.BaseTokenSymbol)
									token.PriceUsd = numeric.BigFloat{*models.MulBigFloats(&token.Price.Float, basePrice)}
								}
								err = s.dao.Save(tx, token)
								if err != nil {
									return errs.NewError(err)
								}
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
	if tokenAddress != "" {
		s.DeleteFilterAddrs(ctx, networkID)
	}
	return nil
}

func (s *Service) CreateMemeTradeHistory(ctx context.Context, event *ethapi.UniswapSwapEventResp) error {
	memeID := uint(0)
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			meme, err := s.dao.FirstMeme(tx,
				map[string][]any{
					"pool = ? or uniswap_pool = ?": {strings.ToLower(event.ContractAddress), strings.ToLower(event.ContractAddress)},
				},
				map[string][]any{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme != nil {
				eventHash := fmt.Sprintf("%s_%d", event.TxHash, event.Index)
				var err error
				history, err := s.dao.FirstMemeHistory(
					tx,
					map[string][]any{
						"event_id = ?": {eventHash},
					},
					map[string][]any{},
					[]string{},
				)
				if err != nil {
					return errs.NewError(err)
				}
				if history == nil {
					eventTime, err := s.GetEthereumClient(ctx, meme.NetworkID).GetBlockTime(event.BlockNumber)
					if err != nil {
						return errs.NewError(err)
					}
					history = &models.MemeTradeHistory{}
					history.EventId = eventHash
					history.NetworkID = meme.NetworkID
					history.ContractAddress = strings.ToLower(event.ContractAddress)
					history.TxHash = strings.ToLower(event.TxHash)
					history.TxAt = time.Unix(int64(eventTime), 0)
					history.Amount0 = numeric.BigFloat{*models.ConvertWeiToBigFloatNegative(event.Amount0, 18)}
					history.Amount1 = numeric.BigFloat{*models.ConvertWeiToBigFloatNegative(event.Amount1, 18)}
					history.SqrtPriceX96 = event.SqrtPriceX96.String()
					history.Liquidity = numeric.BigFloat{*models.ConvertWeiToBigFloat(event.Liquidity, 18)}
					history.Tick = event.Tick.Int64()
					history.MemeTokenAddress = meme.TokenAddress
					history.MemeID = meme.ID
					history.AmountIn = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
					history.TokenInAddress = meme.Token1Address
					history.AmountOut = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
					history.TokenOutAddress = meme.Token0Address
					history.Tick = event.Tick.Int64()
					if history.Amount0.Cmp(big.NewFloat(0)) > 0 {
						history.AmountIn = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
						history.TokenInAddress = meme.Token0Address
						history.AmountOut = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
						history.TokenOutAddress = meme.Token1Address
					}
					fromAddress, err := s.GetEthereumClient(ctx, meme.NetworkID).GetFromFromHash(event.TxHash)
					if err != nil {
						return errs.NewError(err)
					}
					recipient, err := s.GetUser(tx, meme.NetworkID, fromAddress, false)
					if err != nil {
						return errs.NewError(err)
					}
					history.RecipientAddress = recipient.Address
					history.RecipientUserID = recipient.ID
					if meme.BaseTokenIndex == 0 {
						history.BaseAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
						history.TokenAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
						if history.Amount0.Cmp(big.NewFloat(0)) > 0 {
							history.IsBuy = true
						}
					} else {
						history.BaseAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
						history.TokenAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
						if history.Amount1.Cmp(big.NewFloat(0)) > 0 {
							history.IsBuy = true
						}
					}
					history.BaseTokenSymbol = meme.BaseTokenSymbol
					history.Price = numeric.BigFloat{*models.ConvertSqrtPriceX96ToPriceEx(event.SqrtPriceX96, 18, meme.ZeroForOne)}
					basePrice := s.GetTokenMarketPrice(tx, history.BaseTokenSymbol)
					history.BaseTokenPrice = numeric.BigFloat{*basePrice}
					err = s.dao.Create(tx, history)
					if err != nil {
						return errs.NewError(err)
					}
					err = tx.
						Model(meme).
						Updates(
							map[string]any{
								"req_sync_at": time.Now(),
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
					memeID = meme.ID
				}
				poolSwappedUpdated := helpers.GetTimeIndex(uint(event.BlockNumber), event.TxIndex, event.Index)
				if meme.PoolSwappedUpdated < poolSwappedUpdated {
					err = tx.
						Model(meme).
						Updates(
							map[string]any{
								"pool_swapped_updated": poolSwappedUpdated,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if memeID > 0 {
		go func() {
			s.UpdateMemeInfo(context.Background(), memeID)
		}()
	}
	return nil
}

func (s *Service) CreateMemeShareTradeHistory(ctx context.Context, networkID uint64, event *agentshares.AgentSharesTrade) error {
	memeID := uint(0)
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			swapFactoryAddr := s.conf.GetConfigKeyString(networkID, "agentshares_contract_address")
			if strings.EqualFold(swapFactoryAddr, event.Raw.Address.Hex()) {
				meme, err := s.dao.FirstMeme(tx,
					map[string][]any{
						"token_id = ?": {event.TokenId.String()},
					},
					map[string][]any{}, false)
				if err != nil {
					return errs.NewError(err)
				}
				if meme != nil {
					eventHash := fmt.Sprintf("%s_%d", event.Raw.TxHash.Hex(), event.Raw.Index)
					var err error
					history, err := s.dao.FirstMemeHistory(
						tx,
						map[string][]any{
							"event_id = ?": {eventHash},
						},
						map[string][]any{},
						[]string{},
					)
					if err != nil {
						return errs.NewError(err)
					}
					if history == nil {
						eventTime, err := s.GetEthereumClient(ctx, meme.NetworkID).GetBlockTime(event.Raw.BlockNumber)
						if err != nil {
							return errs.NewError(err)
						}
						memeBaseTokenAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "eai_contract_address"))
						history = &models.MemeTradeHistory{}
						history.EventId = eventHash
						history.NetworkID = meme.NetworkID
						history.ContractAddress = strings.ToLower(event.Raw.Address.Hex())
						history.TxHash = strings.ToLower(event.Raw.TxHash.Hex())
						history.TxAt = time.Unix(int64(eventTime), 0)
						history.TokenId = meme.TokenId
						history.MemeID = meme.ID
						history.Amount0 = numeric.BigFloat{*models.ConvertWeiToBigFloatNegative(event.ShareAmount, 1)}
						history.Amount1 = numeric.BigFloat{*models.ConvertWeiToBigFloatNegative(event.EthAmount, 18)}
						history.AmountIn = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
						history.Liquidity = numeric.BigFloat{*models.ConvertWeiToBigFloatNegative(event.Supply, 1)}
						history.TokenInAddress = memeBaseTokenAddress
						history.AmountOut = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
						history.TokenOutAddress = meme.TokenId
						history.BaseAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
						history.TokenAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
						history.IsBuy = true
						if !event.IsBuy {
							history.AmountIn = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
							history.TokenInAddress = meme.TokenId
							history.AmountOut = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
							history.TokenOutAddress = memeBaseTokenAddress
							history.BaseAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount0.Float)}
							history.TokenAmount = numeric.BigFloat{*models.AbsBigFloat(&history.Amount1.Float)}
							history.IsBuy = false
						}
						recipient, err := s.GetUser(tx, meme.NetworkID, event.Trader.Hex(), false)
						if err != nil {
							return errs.NewError(err)
						}
						history.RecipientAddress = recipient.Address
						history.RecipientUserID = recipient.ID
						history.BaseTokenSymbol = meme.BaseTokenSymbol
						supply, _ := history.Liquidity.Float64()
						history.Price = numeric.NewBigFloatFromFloat(big.NewFloat((supply + 1) * supply / 1000))
						basePrice := s.GetTokenMarketPrice(tx, history.BaseTokenSymbol)
						history.BaseTokenPrice = numeric.BigFloat{*basePrice}
						err = s.dao.Create(tx, history)
						if err != nil {
							return errs.NewError(err)
						}
						err = tx.
							Model(meme).
							Updates(
								map[string]any{
									"req_sync_at": time.Now(),
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
						memeID = meme.ID
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if memeID > 0 {
		go func() {
			s.UpdateMemeInfo(context.Background(), memeID)
		}()
	}
	return nil
}

func (s *Service) UpdateMemeInfo(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("UpdateMemeInfo_%d", memeID),
		func() error {
			meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), memeID, map[string][]any{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme.PositionID > 0 {
				tmpResp, err := s.dao.GetMemeTradeHistoryInfo(daos.GetDBMainCtx(ctx), memeID)
				if err != nil {
					return errs.NewError(err)
				}
				if tmpResp != nil {
					updateFields := map[string]any{
						"volume_last24h": tmpResp.VolumeLast24h,
						"total_volume":   tmpResp.TotalVolume,
						"price":          tmpResp.Price,
						"tick":           tmpResp.Tick,
					}
					if tmpResp.PriceLast24h.Cmp(big.NewFloat(0)) > 0 {
						updateFields["price_last24h"] = tmpResp.PriceLast24h
					}
					baseTokenPrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), meme.BaseTokenSymbol)
					updateFields["price_usd"] = numeric.BigFloat{*models.MulBigFloats(&tmpResp.Price.Float, baseTokenPrice)}
					err = daos.GetDBMainCtx(ctx).
						Model(meme).
						Updates(
							updateFields,
						).Error
					if err != nil {
						return errs.NewError(err)
					}

					s.CacheMemeCandleDataChart(daos.GetDBMainCtx(ctx), memeID)
					s.CacheMemeTradeHistoryLatest(daos.GetDBMainCtx(ctx), meme.TokenAddress)
					s.CacheMemeHolders(daos.GetDBMainCtx(ctx), meme.TokenAddress)
					s.CacheMemeDetail(daos.GetDBMainCtx(ctx), meme.TokenAddress)
					err = daos.GetDBMainCtx(ctx).
						Model(meme).
						Updates(
							map[string]any{
								"sync_at": time.Now(),
							},
						).Error
					if err != nil {
						return errs.NewError(err)
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

func (s *Service) UpdateMemeLiquidityPosition(ctx context.Context, networkID uint64, event *ethapi.UniswapPositionLiquidity) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if s.conf.ExistsedConfigKey(networkID, "memeswap_position_mamanger_address") {
				memeSwapPositionMamangerAddress := s.conf.GetConfigKeyString(networkID, "memeswap_position_mamanger_address")
				if strings.EqualFold(memeSwapPositionMamangerAddress, event.ContractAddress) {
					positionInfo, _ := s.GetEthereumClient(ctx, networkID).MemeNonfungiblePositionManagerPositionInfo(
						memeSwapPositionMamangerAddress, big.NewInt(event.TokenId.Int64()),
					)
					if positionInfo != nil {
						meme, err := s.dao.FirstMeme(tx,
							map[string][]any{
								"token_address in (?)": {[]string{strings.ToLower(positionInfo.Token0), strings.ToLower(positionInfo.Token1)}},
								"position_id = 0":      {},
							},
							map[string][]any{},
							false,
						)
						if meme != nil && meme.PositionID <= 0 {
							if meme.AddPool1TxHash == "" || strings.EqualFold(event.TxHash, meme.AddPool1TxHash) {
								meme, err = s.dao.FirstMemeByID(tx, meme.ID, map[string][]any{}, true)
								if err != nil {
									return errs.NewError(err)
								}
								if meme.PositionID <= 0 {
									meme.Liquidity = numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Liquidity, 18))
									meme.PositionID = event.TokenId.Int64()
									err = s.dao.Save(tx, meme)
									if err != nil {
										return errs.NewError(err)
									}
								}
							}
							if err != nil {
								return errs.NewError(err)
							}
						}
					}
				}
			}
			if s.conf.ExistsedConfigKey(networkID, "uniswap_position_mamanger_address") {
				uniswapSwapPositionMamangerAddress := s.conf.GetConfigKeyString(networkID, "uniswap_position_mamanger_address")
				if strings.EqualFold(uniswapSwapPositionMamangerAddress, event.ContractAddress) {
					var positionInfo *ethapi.UniV3SwapPositionInfo
					switch networkID {
					case models.ARBITRUM_CHAIN_ID, models.APE_CHAIN_ID:
						{
							positionInfo, _ = s.GetEthereumClient(ctx, networkID).CamelotNonfungiblePositionManagerPositionInfo(
								uniswapSwapPositionMamangerAddress, big.NewInt(event.TokenId.Int64()),
							)
						}
					default:
						{
							positionInfo, _ = s.GetEthereumClient(ctx, networkID).MemeNonfungiblePositionManagerPositionInfo(
								uniswapSwapPositionMamangerAddress, big.NewInt(event.TokenId.Int64()),
							)
						}
					}
					if positionInfo != nil {
						meme, err := s.dao.FirstMeme(tx,
							map[string][]any{
								"token_address in (?)":    {[]string{strings.ToLower(positionInfo.Token0), strings.ToLower(positionInfo.Token1)}},
								"uniswap_position_id = 0": {},
							},
							map[string][]any{},
							false,
						)
						if meme != nil && meme.UniswapPositionID <= 0 {
							if meme.AddPool2TxHash == "" || strings.EqualFold(event.TxHash, meme.AddPool2TxHash) {
								meme, err = s.dao.FirstMemeByID(tx, meme.ID, map[string][]any{}, true)
								if err != nil {
									return errs.NewError(err)
								}
								if meme.UniswapPositionID <= 0 {
									meme.Liquidity = numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Liquidity, 18))
									meme.UniswapPositionID = event.TokenId.Int64()
									err = s.dao.Save(tx, meme)
									if err != nil {
										return errs.NewError(err)
									}
								}
							}
						}
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

func (s *Service) GeUserByAddress(ctx context.Context, networkID uint64, address string) (*models.User, error) {
	var err error
	user := &models.User{}
	if strings.HasPrefix(address, "0x") {
		user, err = s.GetUser(daos.GetDBMainCtx(ctx), networkID, address, false)
	} else {
		user, err = s.dao.FirstUser(daos.GetDBMainCtx(ctx),
			map[string][]any{
				"address = ?":    {strings.ToLower(address)},
				"network_id = ?": {networkID},
			},
			map[string][]any{},
			false,
		)
	}

	return user, err
}

func (s *Service) VibeTokenFactoryTokenDeployedEvent(ctx context.Context, networkID uint64, event *vibetokenfactory.VibeTokenFactoryTokenDeployed) error {
	if !s.conf.ExistsedConfigKey(networkID, "vibe_token_factory_address") {
		return nil
	}
	vibeTokenFactoryAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "vibe_token_factory_address"))
	if strings.EqualFold(vibeTokenFactoryAddress, event.Raw.Address.Hex()) {
		agentID := big.NewInt(0).SetBytes(event.Nonce[:]).Text(16)
		agentInfo, err := s.dao.FirstAgentInfo(
			daos.GetDBMainCtx(ctx),
			map[string][]any{
				"agent_id = ?": {agentID},
			},
			map[string][]any{},
			[]string{},
		)
		if err != nil {
			return errs.NewError(err)
		}
		if agentInfo != nil {
			if agentInfo.IsVibeAgent() {
				meme, err := s.dao.FirstMeme(
					daos.GetDBMainCtx(ctx),
					map[string][]any{
						"agent_info_id = ?": {agentInfo.ID},
					},
					map[string][]any{},
					false,
				)
				if err != nil {
					return errs.NewError(err)
				}
				if meme != nil {
					err = daos.GetDBMainCtx(ctx).
						Model(meme).
						Updates(
							map[string]any{
								"token_address":     strings.ToLower(event.Token.Hex()),
								"add_pool2_tx_hash": event.Raw.TxHash.Hex(),
								"factory_address":   strings.ToLower(event.Raw.Address.Hex()),
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) VibeTokenFactoryFeesCollectedEvent(ctx context.Context, networkID uint64, event *vibetokenfactory.VibeTokenFactoryFeesCollected) error {
	if !s.conf.ExistsedConfigKey(networkID, "vibe_token_factory_address") {
		return nil
	}
	vibeTokenFactoryAddress := strings.ToLower(s.conf.GetConfigKeyString(networkID, "vibe_token_factory_address"))
	if strings.EqualFold(vibeTokenFactoryAddress, event.Raw.Address.Hex()) {
		err := daos.WithTransaction(
			daos.GetDBMainCtx(ctx),
			func(tx *gorm.DB) error {
				meme, err := s.dao.FirstMeme(
					tx,
					map[string][]any{
						"token_address = ?": {strings.ToLower(event.Token.Hex())},
					},
					map[string][]any{},
					false,
				)
				if err != nil {
					return errs.NewError(err)
				}
				if meme != nil {
					eventId := fmt.Sprintf("%s_%d", event.Raw.TxHash.Hex(), event.Raw.Index)
					memeFeesCollected, err := s.dao.FirstMemeFeesCollected(
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
					if memeFeesCollected == nil {
						memeFeesCollected = &models.MemeFeesCollected{
							EventId: eventId,
							MemeID:  meme.ID,
							TxHash:  event.Raw.TxHash.Hex(),
							Amount0: numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Amount0, 18)),
							Amount1: numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Amount1, 18)),
						}
						err = s.dao.Create(
							tx,
							memeFeesCollected,
						)
						if err != nil {
							return errs.NewError(err)
						}
						err = tx.
							Model(meme).
							Updates(
								map[string]any{
									"fees0_collected_balance": gorm.Expr("fees0_collected_balance + ?", memeFeesCollected.Amount0),
									"fees1_collected_balance": gorm.Expr("fees1_collected_balance + ?", memeFeesCollected.Amount1),
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
					}
					feesCollectedUpdated := helpers.GetTimeIndex(uint(event.Raw.BlockNumber), event.Raw.TxIndex, event.Raw.Index)
					if meme.FeesCollectedUpdated < feesCollectedUpdated {
						err = tx.
							Model(meme).
							Where("fees_collected_updated < ?", feesCollectedUpdated).
							Updates(
								map[string]any{
									"fees_collected_updated": feesCollectedUpdated,
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				}
				return nil
			},
		)
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}
