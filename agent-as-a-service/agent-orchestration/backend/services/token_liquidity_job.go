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
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/apechainnonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/arbitrumnonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/avaxnonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/basenonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/bscnonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/celononfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/memenonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/polygonnonfungiblepositionmanager"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/zksyncnonfungiblepositionmanager"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) JobAgentDeployToken(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentDeployToken",
		func() error {
			memes, err := s.dao.FindMemeJoin(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"join agent_infos on memes.agent_info_id = agent_infos.id": {},
				},
				map[string][]any{
					"memes.status = ?": {models.MemeStatusNew},
					"memes.fee = 0 or agent_infos.eai_balance >= memes.fee or agent_infos.ref_tweet_id > 0": {},
					"memes.num_retries < 3": {},
					"agent_infos.agent_type in (?)": {
						[]models.AgentInfoAgentType{
							models.AgentInfoAgentTypeNormal,
							models.AgentInfoAgentTypeReasoning,
							models.AgentInfoAgentTypeKnowledgeBase,
							models.AgentInfoAgentTypeEliza,
							models.AgentInfoAgentTypeZerepy,
						},
					},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				5,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.AgentDeployToken(ctx, meme.ID)
				if err != nil {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": gorm.Expr("num_retries + ?", 1),
								"updated_at":  time.Now(),
								"err":         err.Error(),
							},
						).Error
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, meme.ID))
				} else {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": 0,
								"updated_at":  time.Now(),
							},
						).Error
					s.DeleteFilterAddrs(ctx, meme.NetworkID)
				}
				time.Sleep(10 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentDeployToken(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentDeployToken_%d", memeID),
		func() error {
			m, err := s.dao.FirstMemeByID(
				daos.GetDBMainCtx(ctx),
				memeID,
				map[string][]any{
					"AgentInfo": {},
				},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if m == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			if m.AgentInfo.IsVibeAgent() {
				return errs.NewError(errs.ErrBadRequest)
			}
			if m.TokenAddress == "" {
				switch m.NetworkID {
				case models.BASE_CHAIN_ID,
					models.ARBITRUM_CHAIN_ID,
					models.BSC_CHAIN_ID,
					models.APE_CHAIN_ID,
					models.AVALANCHE_C_CHAIN_ID,
					models.SOLANA_CHAIN_ID,
					models.CELO_CHAIN_ID:
					{
						if m.AgentInfoID > 0 && m.Fee.Float.Cmp(big.NewFloat(0)) > 0 {
							agent, err := s.dao.FirstAgentInfoByID(
								daos.GetDBMainCtx(ctx),
								m.AgentInfoID,
								map[string][]any{},
								false,
							)
							if err != nil {
								return errs.NewError(err)
							}
							if agent.RefTweetID > 0 || m.NetworkID == models.BITTENSOR_CHAIN_ID {
								m.Fee = numeric.NewBigFloatFromString("0")
								err = daos.GetDBMainCtx(ctx).Model(&m).
									Updates(
										map[string]any{
											"fee": m.Fee,
										},
									).Error
								if err != nil {
									return errs.NewError(err)
								}
							}
							if agent.EaiBalance.Float.Cmp(&m.Fee.Float) < 0 {
								return errs.NewError(errs.ErrBadRequest)
							}
						}
					}
				default:
					{
						return errs.NewError(errs.ErrBadRequest)
					}
				}
				switch m.NetworkID {
				case models.BASE_CHAIN_ID,
					models.ARBITRUM_CHAIN_ID,
					models.BSC_CHAIN_ID,
					models.APE_CHAIN_ID,
					models.AVALANCHE_C_CHAIN_ID,
					models.SOLANA_CHAIN_ID,
					models.CELO_CHAIN_ID:
					{
						tokenSupply := &m.TotalSuply.Float
						if tokenSupply.Cmp(big.NewFloat(1)) <= 0 {
							return errs.NewError(errs.ErrBadRequest)
						}
						var tokenAddress string
						status := models.MemeStatusCreated
						switch m.NetworkID {
						case models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID:
							{
								memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(m.NetworkID, "meme_pool_address"))
								tokenAddress, _, err = s.GetEthereumClient(ctx, m.NetworkID).
									DeployAGENTToken(
										s.GetAddressPrk(memePoolAddress),
										m.Name,
										m.Ticker,
										models.ConvertBigFloatToWei(tokenSupply, 18),
										memePoolAddress,
									)
								if err != nil {
									return errs.NewError(err)
								}
								status = models.MemeStatusCreated
							}
						case models.SOLANA_CHAIN_ID:
							{
								agentTokenAdminAddress := s.conf.GetConfigKeyString(models.GENERTAL_NETWORK_ID, "agent_token_admin_address")
								base64Str, err := helpers.CurlBase64String(models.GetImageUrl(m.Image))
								if err != nil {
									return errs.NewError(err)
								}
								pumfunResp, err := s.blockchainUtils.SolanaCreatePumpfunToken(
									&blockchainutils.SolanaCreatePumpfunTokenReq{
										Address:     agentTokenAdminAddress,
										Name:        m.Name,
										Symbol:      m.Ticker,
										Description: m.Description,
										Telegram:    "",
										Website:     "",
										Amount:      0,
										ImageBase64: base64Str,
									},
								)
								if err != nil {
									return errs.NewError(err)
								}
								tokenAddress = pumfunResp.Mint
								status = models.MemeStatusAddPoolLevel2
							}
						}
						// check token address exist in db
						{
							memeCheck, err := s.dao.FirstMeme(
								daos.GetDBMainCtx(ctx),
								map[string][]any{
									"token_address = ?": {strings.ToLower(tokenAddress)},
								},
								map[string][]any{},
								false,
							)
							if err != nil {
								return errs.NewError(err)
							}
							if memeCheck != nil {
								return errs.NewError(errs.ErrBadRequest)
							}
						}
						err = daos.GetDBMainCtx(ctx).Model(&m).
							Updates(
								map[string]any{
									"token_address": strings.ToLower(tokenAddress),
									"total_suply":   numeric.NewBigFloatFromFloat(tokenSupply),
									"status":        status,
									"num_retries":   0,
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
						if m.AgentInfoID > 0 {
							err = daos.GetDBMainCtx(ctx).
								Model(m.AgentInfo).
								Updates(
									map[string]any{
										"token_address": strings.ToLower(tokenAddress),
										"eai_balance":   gorm.Expr("eai_balance - ?", m.Fee),
									},
								).Error
							if err != nil {
								return errs.NewError(err)
							}
							if m.Fee.Cmp(big.NewFloat(0)) > 0 {
								_ = s.dao.Create(
									daos.GetDBMainCtx(ctx),
									&models.AgentEaiTopup{
										NetworkID:      m.AgentInfo.NetworkID,
										EventId:        fmt.Sprintf("agent_token_fee_%d", m.ID),
										AgentInfoID:    m.AgentInfoID,
										Type:           models.AgentEaiTopupTypeSpent,
										Amount:         m.Fee,
										Status:         models.AgentEaiTopupStatusDone,
										DepositAddress: m.AgentInfo.ETHAddress,
										ToAddress:      m.AgentInfo.ETHAddress,
										Toolset:        "token_fee",
									},
								)
							}
							// _ = s.ReplyAferAutoCreateAgent(daos.GetDBMainCtx(ctx), m.AgentInfo.RefTweetID, m.AgentInfo.ID)
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

func (s *Service) JobRetryAgentDeployToken(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobRetryAgentDeployToken",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"updated_at <= ?":       {time.Now().Add(-120 * time.Minute)},
					"status = ?":            {models.MemeStatusCreated},
					"add_pool1_tx_hash = ?": {""},
					"num_retries < 6":       {},
					"network_id in (?)": {
						[]uint64{
							models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				5,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.RetryAgentDeployToken(ctx, meme.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, meme.ID))
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": gorm.Expr("num_retries + ?", 1),
								"updated_at":  time.Now(),
								"err":         err.Error(),
							},
						).Error
				} else {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": 0,
								"updated_at":  time.Now(),
							},
						).Error
					s.DeleteFilterAddrs(ctx, meme.NetworkID)
				}
				time.Sleep(10 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) RetryAgentDeployToken(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("RetryAgentDeployToken_%d", memeID),
		func() error {
			m, err := s.dao.FirstMemeByID(
				daos.GetDBMainCtx(ctx),
				memeID,
				map[string][]any{
					"AgentInfo": {},
				},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if m == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			if m.TokenAddress != "" && m.Status == models.MemeStatusCreated {
				switch m.NetworkID {
				case models.BASE_CHAIN_ID,
					models.ARBITRUM_CHAIN_ID,
					models.BSC_CHAIN_ID,
					models.APE_CHAIN_ID,
					models.AVALANCHE_C_CHAIN_ID,
					models.CELO_CHAIN_ID:
					{
						if m.AgentInfo.IsVibeAgent() {
							return errs.NewError(errs.ErrBadRequest)
						}
						isContact, err := s.GetEthereumClient(ctx, m.NetworkID).IsContract(m.TokenAddress)
						if err != nil {
							return errs.NewError(err)
						}
						if !isContact {
							tokenSupply := &m.TotalSuply.Float
							if tokenSupply.Cmp(big.NewFloat(1)) <= 0 {
								return errs.NewError(errs.ErrBadRequest)
							}
							memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(m.NetworkID, "meme_pool_address"))
							tokenAddress, _, err := s.GetEthereumClient(ctx, m.NetworkID).
								DeployAGENTToken(
									s.GetAddressPrk(memePoolAddress),
									m.Name,
									m.Ticker,
									models.ConvertBigFloatToWei(tokenSupply, 18),
									memePoolAddress,
								)
							if err != nil {
								return errs.NewError(err)
							}
							// check token address exist in db
							{
								memeCheck, err := s.dao.FirstMeme(
									daos.GetDBMainCtx(ctx),
									map[string][]any{
										"token_address = ?": {strings.ToLower(tokenAddress)},
									},
									map[string][]any{},
									false,
								)
								if err != nil {
									return errs.NewError(err)
								}
								if memeCheck != nil {
									return errs.NewError(errs.ErrBadRequest)
								}
							}
							err = daos.GetDBMainCtx(ctx).Model(&m).
								Updates(
									map[string]any{
										"token_address": strings.ToLower(tokenAddress),
										"total_suply":   numeric.NewBigFloatFromFloat(tokenSupply),
										"status":        models.MemeStatusCreated,
									},
								).Error
							if err != nil {
								return errs.NewError(err)
							}
							// _ = s.ReplyAferAutoCreateAgent(daos.GetDBMainCtx(ctx), m.AgentInfo.RefTweetID, m.AgentInfo.ID)
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

func (s *Service) JobMemeAddPositionInternal(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobMemeAddPositionInternal",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"network_id in (?)": {
						[]uint64{
							models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
					"status = ?":             {models.MemeStatusCreated},
					"add_pool1_tx_hash = ''": {},
					"num_retries < 3":        {},
				},
				map[string][]any{},
				[]string{},
				0,
				999999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.MemeAddPositionInternal(ctx, meme.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, meme.ID))
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": gorm.Expr("num_retries + ?", 1),
								"updated_at":  time.Now(),
								"err":         err.Error(),
							},
						).Error
				} else {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": 0,
								"updated_at":  time.Now(),
							},
						).Error
					s.DeleteFilterAddrs(ctx, meme.NetworkID)
				}
				time.Sleep(10 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) MemeAddPositionInternal(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("MemeAddPositionInternal_%d", memeID),
		func() error {
			meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), memeID, map[string][]any{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme.Status == models.MemeStatusCreated && meme.AddPool1TxHash == "" {
				// 1.000.000.000 token
				checkBalance := models.ConvertBigFloatToWei(models.MulBigFloats(&meme.TotalSuply.Float, big.NewFloat(0.999)), 18)
				if checkBalance.Cmp(big.NewInt(0)) <= 0 {
					return errs.NewError(errs.ErrBadRequest)
				}
				memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(meme.NetworkID, "meme_pool_address"))
				{
					memeBalance, err := s.GetEthereumClient(ctx, meme.NetworkID).Erc20Balance(meme.TokenAddress, memePoolAddress)
					if err != nil {
						return errs.NewError(err)
					}
					if memeBalance.Cmp(checkBalance) >= 0 {
						baseTokenPrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), meme.BaseTokenSymbol)
						if baseTokenPrice.Cmp(big.NewFloat(0)) <= 0 {
							return errs.NewError(errs.ErrBadRequest)
						}
						memeBalanceFloat := models.ConvertWeiToBigFloat(memeBalance, 18)
						lowerPrice, _ := models.QuoBigFloats(
							big.NewFloat(2800),
							memeBalanceFloat,
							baseTokenPrice,
						).Float64()
						upperPrice, _ := models.QuoBigFloats(
							big.NewFloat(1000000000),
							memeBalanceFloat,
							baseTokenPrice,
						).Float64()
						baseToken := s.GetMemeBaseToken(daos.GetDBMainCtx(ctx), meme.NetworkID, meme.BaseTokenSymbol)
						memeToken := strings.ToLower(meme.TokenAddress)
						var token0, token1 string
						var amount0, amount1, tickLower, tickUpper, sqrtPriceX96 *big.Int
						poolFee := int64(20000)
						tickLowerVar := models.PriceToTick(lowerPrice, 1200)
						tickUpperVar := models.PriceToTick(upperPrice, 1200)
						tickCurr := int64(0)
						if strings.Compare(memeToken, baseToken) < 0 {
							token0 = memeToken
							token1 = baseToken
							amount0 = memeBalance
							amount1 = big.NewInt(0)
							tickCurr = tickLowerVar - 1
							sqrtPriceX96 = helpers.GetSqrtRatioAtTick(tickCurr)
							tickLower = big.NewInt(tickLowerVar)
							tickUpper = big.NewInt(tickUpperVar)
						} else {
							token1 = memeToken
							token0 = baseToken
							amount1 = memeBalance
							amount0 = big.NewInt(0)
							tickCurr = -tickLowerVar + 1
							sqrtPriceX96 = helpers.GetSqrtRatioAtTick(tickCurr)
							tickLower = big.NewInt(-tickUpperVar)
							tickUpper = big.NewInt(-tickLowerVar)
						}
						err = daos.GetDBMainCtx(ctx).
							Model(meme).
							Updates(
								map[string]any{
									"tick": tickCurr,
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
						addPoolTxHash, err := s.GetEthereumClient(ctx, meme.NetworkID).MemeNonfungiblePositionManagerMint(
							s.conf.GetConfigKeyString(meme.NetworkID, "memeswap_position_mamanger_address"),
							s.GetAddressPrk(
								memePoolAddress,
							),
							helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
							sqrtPriceX96,
							&memenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
								Token0:         helpers.HexToAddress(token0),
								Token1:         helpers.HexToAddress(token1),
								Fee:            big.NewInt(poolFee),
								TickLower:      tickLower,
								TickUpper:      tickUpper,
								Amount0Desired: amount0,
								Amount1Desired: amount1,
								Amount0Min:     big.NewInt(0),
								Amount1Min:     big.NewInt(0),
								Deadline:       big.NewInt(time.Now().Add(60 * time.Second).Unix()),
							},
						)
						if err != nil {
							return errs.NewError(err)
						}
						err = daos.GetDBMainCtx(ctx).
							Model(meme).
							Updates(
								map[string]any{
									"pool_fee":          poolFee,
									"tick":              tickCurr,
									"add_pool1_tx_hash": addPoolTxHash,
									"price":             numeric.NewBigFloatFromFloat(big.NewFloat(lowerPrice)),
									"price_last24h":     numeric.NewBigFloatFromFloat(big.NewFloat(lowerPrice)),
									"price_usd":         numeric.NewBigFloatFromFloat(models.MulBigFloats(big.NewFloat(lowerPrice), baseTokenPrice)),
									"tick_lower":        tickLower,
									"tick_upper":        tickUpper,
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
						//update cache pair detail
						go func() {
							time.Sleep(10 * time.Second)
							_ = s.MemeEventsByTransaction(context.Background(), meme.NetworkID, addPoolTxHash)
						}()
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

func (s *Service) JobMemeRemovePositionInternal(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobMemeRemovePositionInternal",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"status = ?":                {models.MemeStatusReachedMC},
					"remove_pool1_tx_hash = ''": {},
					`network_id != ?`:           {models.BITTENSOR_CHAIN_ID},
					"num_retries < 3":           {},
					"not_graduated = false":     {},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				5,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.MemeRemovePositionInternal(ctx, meme.ID)
				if err != nil {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": gorm.Expr("num_retries + ?", 1),
								"updated_at":  time.Now(),
								"err":         err.Error(),
							},
						).Error
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, meme.ID))
				} else {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": 0,
								"updated_at":  time.Now(),
							},
						).Error
					s.DeleteFilterAddrs(ctx, meme.NetworkID)
				}
				time.Sleep(10 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) MemeRemovePositionInternal(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("MemeRemovePositionInternal_%d", memeID),
		func() error {
			meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), memeID, map[string][]any{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme.Status == models.MemeStatusReachedMC {
				networkID := meme.NetworkID
				positinInfo, err := s.GetEthereumClient(ctx, networkID).MemeNonfungiblePositionManagerPositionInfo(
					s.conf.GetConfigKeyString(networkID, "memeswap_position_mamanger_address"),
					big.NewInt(int64(meme.PositionID)),
				)
				if err != nil {
					return errs.NewError(err)
				}
				meme.PositionLiquidity = numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(positinInfo.Liquidity, 18))
				meme.TickLower = positinInfo.TickLower.Int64()
				meme.TickUpper = positinInfo.TickUpper.Int64()
				err = daos.GetDBMainCtx(ctx).
					Model(meme).
					Updates(
						map[string]any{
							"position_liquidity": meme.PositionLiquidity,
							"tick_lower":         meme.TickLower,
							"tick_upper":         meme.TickUpper,
						},
					).Error
				if err != nil {
					return errs.NewError(err)
				}
				memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(meme.NetworkID, "meme_pool_address"))
				removePoolTxHash, err := s.GetEthereumClient(ctx, networkID).MemeNonfungiblePositionManagerBurn(
					s.conf.GetConfigKeyString(networkID, "memeswap_position_mamanger_address"),
					s.GetAddressPrk(
						memePoolAddress,
					),
					helpers.HexToAddress(s.conf.GetConfigKeyString(networkID, "weth9_contract_address")),
					big.NewInt(int64(meme.PositionID)),
				)
				if err != nil {
					return errs.NewError(err)
				}
				err = daos.GetDBMainCtx(ctx).
					Model(meme).
					Updates(
						map[string]any{
							"remove_pool1_tx_hash": removePoolTxHash,
							"status":               models.MemeStatusRemovePoolLelve1,
						},
					).Error
				if err != nil {
					return errs.NewError(err)
				}
				// snapshot token holder
				go func() {
					time.Sleep(10 * time.Second)
					_ = s.MemeEventsByTransaction(context.Background(), meme.NetworkID, removePoolTxHash)
				}()
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) JobMemeAddPositionUniswap(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobMemeAddPositionUniswap",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"network_id in (?)": {
						[]uint64{
							models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
					"status = ?":             {models.MemeStatusRemovePoolLelve1},
					"add_pool2_tx_hash = ''": {},
					`network_id != ?`:        {models.BITTENSOR_CHAIN_ID},
					"num_retries < 3":        {},
					"not_graduated = false":  {},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				5,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.MemeAddPositionUniswap(ctx, meme.ID)
				if err != nil {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": gorm.Expr("num_retries + ?", 1),
								"updated_at":  time.Now(),
								"err":         err.Error(),
							},
						).Error
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, meme.ID))
				} else {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": 0,
								"updated_at":  time.Now(),
							},
						).Error
					s.DeleteFilterAddrs(ctx, meme.NetworkID)
				}
				time.Sleep(10 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) MemeAddPositionUniswap(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("MemeAddPositionUniswap_%d", memeID),
		func() error {
			meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), memeID, map[string][]any{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme.Status == models.MemeStatusRemovePoolLelve1 {
				memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(meme.NetworkID, "meme_pool_address"))
				{
					memeBalance, err := s.GetEthereumClient(ctx, meme.NetworkID).Erc20Balance(meme.TokenAddress, memePoolAddress)
					if err != nil {
						return errs.NewError(err)
					}
					if memeBalance.Cmp(models.Number2BigInt("100000000", 18)) > 0 {
						baseTokenPrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), meme.BaseTokenSymbol)
						if baseTokenPrice.Cmp(big.NewFloat(0)) <= 0 {
							return errs.NewError(errs.ErrBadRequest)
						}
						baseToken := s.GetMemeBaseToken(daos.GetDBMainCtx(ctx), meme.NetworkID, meme.BaseTokenSymbol)
						memeToken := meme.TokenAddress
						var token0, token1 string
						var amount0, amount1, tickLower, tickUpper, sqrtPriceX96 *big.Int
						poolFee := int64(3000)
						tickLowerVar := meme.TickLower
						tickUpperVar := meme.TickUpper
						switch meme.NetworkID {
						case models.BSC_CHAIN_ID:
							{
								poolFee = int64(2500)
							}
						}
						liquidity := models.ConvertBigFloatToWei(&meme.PositionLiquidity.Float, 18)
						if strings.Compare(memeToken, baseToken) < 0 {
							token0 = memeToken
							token1 = baseToken
							amount0 = memeBalance
							sqrtPriceX96 = models.GetSqrtPriceX96ForLiquidityAndAmount0(
								helpers.GetSqrtRatioAtTick(tickLowerVar),
								helpers.GetSqrtRatioAtTick(tickUpperVar),
								liquidity,
								memeBalance,
							)
							amount1 = helpers.GetAmount1ForLiquidity(
								helpers.GetSqrtRatioAtTick(tickLowerVar),
								sqrtPriceX96,
								liquidity,
							)
							tickLower = big.NewInt(tickLowerVar)
							tickUpper = big.NewInt(tickUpperVar)
						} else {
							token1 = memeToken
							token0 = baseToken
							amount1 = memeBalance
							sqrtPriceX96 = models.GetSqrtPriceX96ForLiquidityAndAmount1(
								helpers.GetSqrtRatioAtTick(tickLowerVar),
								helpers.GetSqrtRatioAtTick(tickUpperVar),
								liquidity,
								memeBalance,
							)
							amount0 = helpers.GetAmount0ForLiquidity(
								sqrtPriceX96,
								helpers.GetSqrtRatioAtTick(tickUpperVar),
								liquidity,
							)
							tickLower = big.NewInt(tickLowerVar)
							tickUpper = big.NewInt(tickUpperVar)
						}
						priceF, _ := models.ConvertSqrtPriceX96ToPriceEx(sqrtPriceX96, 18, meme.ZeroForOne).Float64()
						tickCurr := models.PriceToTick(priceF, 1)
						if meme.ZeroForOne {
							tickCurr = models.PriceToTick(1/priceF, 1)
						}
						err = daos.GetDBMainCtx(ctx).
							Model(meme).
							Updates(
								map[string]any{
									"tick":         tickCurr,
									"add_pool2_at": time.Now(),
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
						var addPoolTxHash string
						switch meme.NetworkID {
						case models.BASE_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).BaseNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										Fee:            big.NewInt(poolFee),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
										Recipient:      helpers.HexToAddress(memePoolAddress),
									},
								)
							}
						case models.ARBITRUM_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).ArbitrumNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&arbitrumnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
									},
								)
							}
						case models.BSC_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).BscNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&bscnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										Fee:            big.NewInt(poolFee),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
									},
								)
							}
						case models.APE_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).ApechainNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&apechainnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
									},
								)
							}
						case models.AVALANCHE_C_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).AvaxNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&avaxnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										Fee:            big.NewInt(poolFee),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
										Recipient:      helpers.HexToAddress(memePoolAddress),
									},
								)
							}
						case models.POLYGON_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).PolygonNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&polygonnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										Fee:            big.NewInt(poolFee),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
										Recipient:      helpers.HexToAddress(memePoolAddress),
									},
								)
							}
						case models.ZKSYNC_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetZkClient(ctx, meme.NetworkID).ZksyncNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&zksyncnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										Fee:            big.NewInt(poolFee),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
										Recipient:      helpers.HexToAddress(memePoolAddress),
									},
								)
							}
						case models.CELO_CHAIN_ID:
							{
								addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).CeloNonfungiblePositionManagerMint(
									s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
									s.GetAddressPrk(
										memePoolAddress,
									),
									helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
									sqrtPriceX96,
									&celononfungiblepositionmanager.INonfungiblePositionManagerMintParams{
										Token0:         helpers.HexToAddress(token0),
										Token1:         helpers.HexToAddress(token1),
										Fee:            big.NewInt(poolFee),
										TickLower:      tickLower,
										TickUpper:      tickUpper,
										Amount0Desired: models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(9999)), big.NewInt(10000)),
										Amount1Desired: models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(9999)), big.NewInt(10000)),
										Amount0Min:     models.QuoBigInts(models.MulBigInts(amount0, big.NewInt(99)), big.NewInt(100)),
										Amount1Min:     models.QuoBigInts(models.MulBigInts(amount1, big.NewInt(99)), big.NewInt(100)),
										Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
										Recipient:      helpers.HexToAddress(memePoolAddress),
									},
								)
							}
						default:
							{
								return errs.NewError(errs.ErrBadRequest)
							}
						}
						if err != nil {
							return errs.NewError(err)
						}
						err = daos.GetDBMainCtx(ctx).
							Model(meme).
							Updates(
								map[string]any{
									"pool_fee":          poolFee,
									"tick":              tickCurr,
									"add_pool2_tx_hash": addPoolTxHash,
									"price":             numeric.NewBigFloatFromFloat(big.NewFloat(priceF)),
									"price_usd":         numeric.NewBigFloatFromFloat(models.MulBigFloats(big.NewFloat(priceF), baseTokenPrice)),
									"tick_lower":        tickLower,
									"tick_upper":        tickUpper,
									"add_pool2_at":      time.Now(),
								},
							).Error
						if err != nil {
							return errs.NewError(err)
						}
						//update cache pair detail
						go func() {
							_ = s.CreateMemeNotifications(daos.GetDBMainCtx(ctx), 0, meme.ID, 0, models.NotiTypeNewMeme, fmt.Sprintf("%s_%d", models.NotiTypeNewMeme, meme.ID))
							_ = s.CacheMemeDetail(daos.GetDBMainCtx(ctx), meme.TokenAddress)
							time.Sleep(10 * time.Second)
							_ = s.MemeEventsByTransaction(context.Background(), meme.NetworkID, addPoolTxHash)
						}()
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

// func (s *Service) MemeAddPositionUniswap(ctx context.Context, memeID uint) error {
// 	meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), memeID, map[string][]any{}, false)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	if meme.Status == models.MemeStatusRemovePoolLelve1 && meme.AddPool2TxHash == "" {
// 		// 1.000.000.000 token
// 		checkBalance := models.ConvertBigFloatToWei(models.MulBigFloats(&meme.TotalSuply.Float, big.NewFloat(0.999)), 18)
// 		if checkBalance.Cmp(big.NewInt(0)) <= 0 {
// 			return errs.NewError(errs.ErrBadRequest)
// 		}
// 		memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(meme.NetworkID, "meme_pool_address"))
// 		{
// 			memeBalance, err := s.GetEthereumClient(ctx, meme.NetworkID).Erc20Balance(meme.TokenAddress, memePoolAddress)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 			if memeBalance.Cmp(checkBalance) >= 0 {
// 				baseTokenPrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), meme.BaseTokenSymbol)
// 				if baseTokenPrice.Cmp(big.NewFloat(0)) <= 0 {
// 					return errs.NewError(errs.ErrBadRequest)
// 				}
// 				memeBalanceFloat := models.ConvertWeiToBigFloat(memeBalance, 18)
// 				lowerPrice, _ := models.QuoBigFloats(
// 					big.NewFloat(2800),
// 					memeBalanceFloat,
// 					baseTokenPrice,
// 				).Float64()
// 				upperPrice, _ := models.QuoBigFloats(
// 					big.NewFloat(1000000000),
// 					memeBalanceFloat,
// 					baseTokenPrice,
// 				).Float64()
// 				baseToken := s.GetMemeBaseToken(daos.GetDBMainCtx(ctx), meme.NetworkID, meme.BaseTokenSymbol)
// 				memeToken := strings.ToLower(meme.TokenAddress)
// 				var token0, token1 string
// 				var amount0, amount1, tickLower, tickUpper, sqrtPriceX96 *big.Int
// 				poolFee := int64(3000)
// 				tickLowerVar := models.PriceToTick(lowerPrice, 60)
// 				tickUpperVar := models.PriceToTick(upperPrice, 60)
// 				switch meme.NetworkID {
// 				case models.BSC_CHAIN_ID:
// 					{
// 						poolFee = int64(2500)
// 						tickLowerVar = models.PriceToTick(lowerPrice, 50)
// 						tickUpperVar = models.PriceToTick(upperPrice, 50)
// 					}
// 				}
// 				tickCurr := int64(0)
// 				if strings.Compare(memeToken, baseToken) < 0 {
// 					token0 = memeToken
// 					token1 = baseToken
// 					amount0 = memeBalance
// 					amount1 = big.NewInt(0)
// 					tickCurr = tickLowerVar - 1
// 					sqrtPriceX96 = helpers.GetSqrtRatioAtTick(tickCurr)
// 					tickLower = big.NewInt(tickLowerVar)
// 					tickUpper = big.NewInt(tickUpperVar)
// 				} else {
// 					token1 = memeToken
// 					token0 = baseToken
// 					amount1 = memeBalance
// 					amount0 = big.NewInt(0)
// 					tickCurr = -tickLowerVar + 1
// 					sqrtPriceX96 = helpers.GetSqrtRatioAtTick(tickCurr)
// 					tickLower = big.NewInt(-tickUpperVar)
// 					tickUpper = big.NewInt(-tickLowerVar)
// 				}
// 				err = daos.GetDBMainCtx(ctx).
// 					Model(meme).
// 					Updates(
// 						map[string]any{
// 							"tick": tickCurr,
// 						},
// 					).Error
// 				if err != nil {
// 					return errs.NewError(err)
// 				}
// 				var addPoolTxHash string
// 				switch meme.NetworkID {
// 				case models.BASE_CHAIN_ID:
// 					{
// 						addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).BaseNonfungiblePositionManagerMint(
// 							s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
// 							s.GetAddressPrk(
// 								memePoolAddress,
// 							),
// 							helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
// 							sqrtPriceX96,
// 							&basenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
// 								Token0:         helpers.HexToAddress(token0),
// 								Token1:         helpers.HexToAddress(token1),
// 								Fee:            big.NewInt(poolFee),
// 								TickLower:      tickLower,
// 								TickUpper:      tickUpper,
// 								Amount0Desired: amount0,
// 								Amount1Desired: amount1,
// 								Amount0Min:     big.NewInt(0),
// 								Amount1Min:     big.NewInt(0),
// 								Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
// 							},
// 						)
// 						if err != nil {
// 							return errs.NewError(err)
// 						}
// 					}
// 				case models.ARBITRUM_CHAIN_ID:
// 					{
// 						addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).ArbitrumNonfungiblePositionManagerMint(
// 							s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
// 							s.GetAddressPrk(
// 								memePoolAddress,
// 							),
// 							helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
// 							sqrtPriceX96,
// 							&arbitrumnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
// 								Token0:         helpers.HexToAddress(token0),
// 								Token1:         helpers.HexToAddress(token1),
// 								TickLower:      tickLower,
// 								TickUpper:      tickUpper,
// 								Amount0Desired: amount0,
// 								Amount1Desired: amount1,
// 								Amount0Min:     big.NewInt(0),
// 								Amount1Min:     big.NewInt(0),
// 								Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
// 							},
// 						)
// 						if err != nil {
// 							return errs.NewError(err)
// 						}
// 					}
// 				case models.BSC_CHAIN_ID:
// 					{
// 						addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).BscNonfungiblePositionManagerMint(
// 							s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
// 							s.GetAddressPrk(
// 								memePoolAddress,
// 							),
// 							helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
// 							sqrtPriceX96,
// 							&bscnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
// 								Token0:         helpers.HexToAddress(token0),
// 								Token1:         helpers.HexToAddress(token1),
// 								Fee:            big.NewInt(poolFee),
// 								TickLower:      tickLower,
// 								TickUpper:      tickUpper,
// 								Amount0Desired: amount0,
// 								Amount1Desired: amount1,
// 								Amount0Min:     big.NewInt(0),
// 								Amount1Min:     big.NewInt(0),
// 								Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
// 							},
// 						)
// 						if err != nil {
// 							return errs.NewError(err)
// 						}
// 					}
// 				case models.APE_CHAIN_ID:
// 					{
// 						addPoolTxHash, err = s.GetEthereumClient(ctx, meme.NetworkID).ApechainNonfungiblePositionManagerMint(
// 							s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
// 							s.GetAddressPrk(
// 								memePoolAddress,
// 							),
// 							helpers.HexToAddress(s.conf.GetConfigKeyString(meme.NetworkID, "weth9_contract_address")),
// 							sqrtPriceX96,
// 							&apechainnonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
// 								Token0:         helpers.HexToAddress(token0),
// 								Token1:         helpers.HexToAddress(token1),
// 								TickLower:      tickLower,
// 								TickUpper:      tickUpper,
// 								Amount0Desired: amount0,
// 								Amount1Desired: amount1,
// 								Amount0Min:     big.NewInt(0),
// 								Amount1Min:     big.NewInt(0),
// 								Deadline:       big.NewInt(time.Now().Add(120 * time.Second).Unix()),
// 							},
// 						)
// 						if err != nil {
// 							return errs.NewError(err)
// 						}
// 					}
// 				default:
// 					{
// 						return errs.NewError(errs.ErrBadRequest)
// 					}
// 				}
// 				err = daos.GetDBMainCtx(ctx).
// 					Model(meme).
// 					Updates(
// 						map[string]any{
// 							"pool_fee":          poolFee,
// 							"tick":              tickCurr,
// 							"add_pool2_tx_hash": addPoolTxHash,
// 							"price":             numeric.NewBigFloatFromFloat(big.NewFloat(lowerPrice)),
// 							"price_last24h":     numeric.NewBigFloatFromFloat(big.NewFloat(lowerPrice)),
// 							"price_usd":         numeric.NewBigFloatFromFloat(models.MulBigFloats(big.NewFloat(lowerPrice), baseTokenPrice)),
// 							"tick_lower":        tickLower,
// 							"tick_upper":        tickUpper,
// 						},
// 					).Error
// 				if err != nil {
// 					return errs.NewError(err)
// 				}
// 				//update cache pair detail
// 				go func() {
// 					time.Sleep(10 * time.Second)
// 					_ = s.MemeEventsByTransaction(context.Background(), meme.NetworkID, addPoolTxHash)
// 				}()
// 			}
// 		}
// 	}
// 	return nil
// }

func (s *Service) JobCheckMemeReachMarketCap(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobCheckMemeReachMarketCap",
		func() error {
			filters := map[string][]any{
				`price_usd * supply >= 69000`: {},
				`status = ?`:                  {models.MemeStatusAddPoolLevel1},
				`network_id != ?`:             {models.BITTENSOR_CHAIN_ID},
				"not_graduated = false":       {},
			}
			memes, err := s.dao.FindMeme(daos.GetDBMainCtx(ctx),
				filters,
				map[string][]any{},
				[]string{},
				0,
				5,
			)
			if err != nil {
				return errs.NewError(err)
			}
			for _, meme := range memes {
				updateFields := map[string]any{
					"status": models.MemeStatusReachedMC,
				}
				err := daos.GetDBMainCtx(ctx).
					Model(meme).
					Where("status = ?", models.MemeStatusAddPoolLevel1).
					Updates(
						updateFields,
					).Error
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) JobUpdateMemeUsdPrice(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobUpdateMemeUsdPrice",
		func() error {
			{
				basePrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), string(models.BaseTokenSymbolETH))
				if basePrice.Cmp(big.NewFloat(0)) <= 0 {
					return errs.NewError(errs.ErrBadRequest)
				}
				err := daos.GetDBMainCtx(ctx).
					Model(&models.Meme{}).
					Where("base_token_symbol = ?", models.BaseTokenSymbolETH).
					UpdateColumn(
						"price_usd", gorm.Expr("price * ?", numeric.NewBigFloatFromFloat(basePrice)),
					).Error
				if err != nil {
					return errs.NewError(err)
				}
			}
			{
				basePrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), string(models.BaseTokenSymbolEAI))
				if basePrice.Cmp(big.NewFloat(0)) <= 0 {
					return errs.NewError(errs.ErrBadRequest)
				}
				err := daos.GetDBMainCtx(ctx).
					Model(&models.Meme{}).
					Where("base_token_symbol = ?", models.BaseTokenSymbolEAI).
					UpdateColumn(
						"price_usd", gorm.Expr("price * ?", numeric.NewBigFloatFromFloat(basePrice)),
					).Error
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) JobRetryAddPool1(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobRetryAddPool1",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"updated_at <= ?":        {time.Now().Add(-120 * time.Minute)},
					"status = ?":             {models.MemeStatusCreated},
					"add_pool1_tx_hash != ?": {""},
					"network_id in (?)": {
						[]uint64{
							models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				10,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				fmt.Println(meme.AddPool1TxHash)
				err = s.GetEVMClient(ctx, meme.NetworkID).TransactionConfirmed(meme.AddPool1TxHash)
				if err != nil {
					fmt.Println(err.Error())
					if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "transaction is not Successful") {
						err = daos.GetDBMainCtx(ctx).
							Model(meme).
							Updates(
								map[string]any{
									"add_pool1_tx_hash": "",
								},
							).
							Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				} else {
					s.MemeEventsByTransaction(ctx, meme.NetworkID, meme.AddPool1TxHash)
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) JobRetryAddPool2(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobRetryAddPool1",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"updated_at <= ?":        {time.Now().Add(-120 * time.Minute)},
					"status = ?":             {models.MemeStatusRemovePoolLelve1},
					"add_pool2_tx_hash != ?": {""},
					"network_id in (?)": {
						[]uint64{
							models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
				},
				map[string][]any{},
				[]string{
					"rand()",
				},
				0,
				10,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.GetEVMClient(ctx, meme.NetworkID).TransactionConfirmed(meme.AddPool2TxHash)
				if err != nil {
					fmt.Println(err.Error())
					if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "transaction is not Successful") {
						err = daos.GetDBMainCtx(ctx).
							Model(meme).
							Updates(
								map[string]any{
									"add_pool2_tx_hash": "",
								},
							).
							Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				} else {
					s.MemeEventsByTransaction(ctx, meme.NetworkID, meme.AddPool2TxHash)
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) JobMemeBurnPositionUniswap(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobMemeBurnPositionUniswap",
		func() error {
			memes, err := s.dao.FindMeme(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"factory_address is null or factory_address != ''": {},
					"add_pool2_at <= ?":       {time.Now().Add(-24 * time.Hour)},
					"status = ?":              {models.MemeStatusAddPoolLevel2},
					"uniswap_position_id > 0": {},
					"burn_pool2_at is null":   {},
					"network_id in (?)": {
						[]uint64{
							models.BASE_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.CELO_CHAIN_ID,
						},
					},
					"num_retries < 3": {},
				},
				map[string][]any{},
				[]string{
					"add_pool2_at asc",
				},
				0,
				2,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, meme := range memes {
				err = s.MemeBurnPositionUniswap(ctx, meme.ID)
				if err != nil {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": gorm.Expr("num_retries + ?", 1),
								"updated_at":  time.Now(),
								"err":         err.Error(),
							},
						).Error
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, meme.ID))
				} else {
					_ = daos.GetDBMainCtx(ctx).
						Model(&meme).
						Updates(
							map[string]any{
								"num_retries": 0,
								"updated_at":  time.Now(),
							},
						).Error
				}
				time.Sleep(10 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) MemeBurnPositionUniswap(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("MemeBurnPositionUniswap_%d", memeID),
		func() error {
			meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), memeID, map[string][]any{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme.Status == models.MemeStatusAddPoolLevel2 &&
				meme.UniswapPositionID > 0 &&
				meme.AddPool2At.Before(time.Now().Add(24*time.Hour)) {
				memePoolAddress := strings.ToLower(s.conf.GetConfigKeyString(meme.NetworkID, "meme_pool_address"))
				{
					burnPoolTxHash, err := s.GetEVMClient(ctx, meme.NetworkID).Erc721Transfer(
						s.conf.GetConfigKeyString(meme.NetworkID, "uniswap_position_mamanger_address"),
						s.GetAddressPrk(
							memePoolAddress,
						),
						models.BURN_ADDRESS,
						big.NewInt(meme.UniswapPositionID),
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).
						Model(meme).
						Updates(
							map[string]any{
								"burn_pool2_at":      time.Now(),
								"burn_pool2_tx_hash": burnPoolTxHash,
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
