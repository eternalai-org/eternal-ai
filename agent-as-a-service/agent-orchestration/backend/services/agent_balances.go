package services

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/dexscreener"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) FindAgentSnapshotMission(ctx context.Context, agentId uint) ([]*models.AgentSnapshotMission, error) {
	filters := map[string][]any{
		`agent_info_id = ?`: {agentId},
	}

	return s.dao.FindAgentSnapshotMission(daos.GetDBMainCtx(ctx), filters, nil, []string{}, 0, 999999)
}

func (s *Service) FindAgentSnapshotPostAction(ctx context.Context, agentId uint) ([]*models.AgentSnapshotPostAction, error) {
	filters := map[string][]any{
		`agent_info_id = ?`: {agentId},
		// `status = ?`:        {models.AgentSnapshotPostActionStatusDone},
	}

	preloads := map[string][]any{}
	preloads[`AgentSnapshotMission`] = []any{}

	return s.dao.FindAgentSnapshotPostAction(daos.GetDBMainCtx(ctx), filters, preloads, []string{"id DESC"}, 0, 50)
}

func (s *Service) GetListAgentCategory(ctx context.Context, page, limit int) ([]*models.AgentCategory, error) {
	filters := map[string][]any{}
	ms, err := s.dao.FindAgentCategory(
		daos.GetDBMainCtx(ctx),
		filters,
		nil,
		[]string{"priority DESC", "id ASC"},
		page,
		limit,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (s *Service) GetListAgentInfos(ctx context.Context, networkID uint64, creator string, agentTypes []uint, kbStatus int64, keyword string, page, limit int) ([]*models.AgentInfo, uint, error) {
	selected := []string{
		"agent_infos.*",
	}
	joinFilters := map[string][]any{}

	filters := map[string][]any{}
	if networkID > 0 {
		filters["network_id = ?"] = []any{networkID}
	}
	if creator != "" {
		filters["creator = ?"] = []any{strings.ToLower(creator)}
	}

	if keyword != "" {
		search := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		filters[`LOWER(token_name) like ?
			or LOWER(token_symbol) like ?
			or LOWER(token_address) like ?
			or LOWER(twitter_username) like ?
			or LOWER(agent_name) like ?`] = []any{search, search, search, search, search}
	}

	if len(agentTypes) != 0 {
		filters["agent_type IN (?)"] = []any{agentTypes}
	}

	if kbStatus > 0 {
		joinFilters["JOIN knowledge_bases kb ON kb.id = agent_infos.agent_kb_id"] = []any{}
		filters["kb.status = ?"] = []any{kbStatus}
	}

	keys, err := s.dao.FindAgentInfoJoinSelect(
		daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"TwitterInfo":                      {},
			"TmpTwitterInfo":                   {},
			"Meme":                             {`deleted_at IS NULL`},
			"TokenInfo":                        {},
			"KnowledgeBase":                    {},
			"KnowledgeBase.KnowledgeBaseFiles": {},
			"AgentCategory":                    {},
		},
		[]string{"created_at desc"}, page, limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	for _, key := range keys {
		needBalance, _, err := s.CheckAgentIsReadyToRunTwinTraining(key)
		if err != nil {
			return nil, 0, errs.NewError(err)
		}
		key.TotalMintTwinFee = needBalance
	}
	return keys, 0, nil
}

func (s *Service) GetListAgentUnClaimed(ctx context.Context, search string, page, limit int) ([]*models.AgentInfo, uint, error) {
	selected := []string{
		"agent_infos.*",
	}
	joinFilters := map[string][]any{}

	admin1 := strings.ToLower(s.conf.GetConfigKeyString(models.BASE_CHAIN_ID, "meme_pool_address"))
	admin2 := s.conf.AdminAutoCreateAgentAddress
	filters := map[string][]any{
		`creator in (?)`:   {[]string{admin1, admin2}},
		`ref_tweet_id > 0`: {},
	}
	if search != "" {
		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
		filters[`
			LOWER(agent_infos.token_name) like ? 
			or LOWER(agent_infos.token_symbol) like ? 
			or LOWER(agent_infos.token_address) like ?
			or LOWER(agent_infos.twitter_username) like ?
			or LOWER(agent_infos.agent_name) like ?
			or ifnull(twitter_users.twitter_username, "") like ?
			or ifnull(twitter_users.name, "") like ?
		`] = []any{search, search, search, search, search, search, search}
	}
	keys, err := s.dao.FindAgentInfoJoinSelect(
		daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"TwitterInfo":    {},
			"TmpTwitterInfo": {},
			"Meme":           {`deleted_at IS NULL`},
			"TokenInfo":      {},
		},
		[]string{"created_at desc"}, page, limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return keys, 0, nil
}

func (s *Service) GetAgentInfoDetail(ctx context.Context, networkID uint64, agentID string) (*models.AgentInfo, error) {
	selected := []string{
		"agent_infos.*",
	}
	joinFilters := map[string][]any{}

	filters := map[string][]any{
		`agent_infos.agent_id = ? or agent_infos.id= ?`: {agentID, agentID},
	}
	if networkID > 0 {
		filters["network_id = ?"] = []any{networkID}
	}
	agent, err := s.dao.FirstAgentInfoJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"TwitterInfo":          {},
			"TokenInfo":            {},
			"AgentSnapshotMission": {"enabled = 1"},
		},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	needBalance, _, err := s.CheckAgentIsReadyToRunTwinTraining(agent)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agent.TotalMintTwinFee = needBalance
	return agent, nil
}

func (s *Service) GetAgentInfoDetailByAgentID(ctx context.Context, agentID string) (*models.AgentInfo, error) {
	filters := map[string][]any{}
	if id, err := strconv.ParseInt(agentID, 10, 64); err == nil {
		filters[`agent_infos.id = ?`] = []any{id}
	} else {
		filters[`agent_infos.agent_id = ?`] = []any{agentID}
	}

	agent, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		filters,
		map[string][]any{
			"TwitterInfo":                      {},
			"TokenInfo":                        {},
			"AgentSnapshotMission":             {"enabled = 1 and is_testing=0"},
			"KnowledgeBase":                    {},
			"KnowledgeBase.KnowledgeBaseFiles": {},
		},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	needBalance, _, err := s.CheckAgentIsReadyToRunTwinTraining(agent)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agent.TotalMintTwinFee = needBalance
	return agent, nil
}

func (s *Service) FindAgenByTwitterUsername(ctx context.Context, twitterUsername string) (*models.AgentInfo, error) {
	filters := map[string][]any{
		"twitter_username = ?": {twitterUsername},
	}
	return s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx), filters, nil, nil)
}

func (s *Service) SyncAgentInfoDetailByAgentID(ctx context.Context, agentID string) (*models.AgentInfo, error) {
	agent, err := s.GetAgentInfoDetailByAgentID(ctx, agentID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	var isGen bool
	if agent.TipBtcAddress == "" {
		address, err := s.CreateBTCAddress(ctx)
		if err != nil {
			return nil, errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_btc_address", address).Error
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	if agent.TipEthAddress == "" {
		address, err := s.CreateETHAddress(ctx)
		if err != nil {
			return nil, errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_eth_address", address).Error
		if err != nil {
			return nil, errs.NewError(err)
		}
		isGen = true
	}
	if agent.TipSolAddress == "" {
		address, err := s.CreateSOLAddress(ctx)
		if err != nil {
			return nil, errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_sol_address", address).Error
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	if isGen {
		agent, err = s.GetAgentInfoDetailByAgentID(ctx, agentID)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	needBalance, _, err := s.CheckAgentIsReadyToRunTwinTraining(agent)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agent.TotalMintTwinFee = needBalance
	return agent, nil
}

func (s *Service) GetAgentSummaryReport(ctx context.Context) ([]*models.AgentInfo, error) {
	agent, err := s.dao.GetAgentSummaryReport(daos.GetDBMainCtx(ctx), s.conf.HiddenNetworkId)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return agent, nil
}

func (s *Service) GetEstimateTime(ctx context.Context, agentInfo *models.AgentInfo) (*time.Time, error) {
	if agentInfo.EstimateTwinDoneTimestamp != nil {
		return agentInfo.EstimateTwinDoneTimestamp, nil
	}
	var estimate time.Time
	if agentInfo.TwinStatus == models.TwinStatusRunning {
		estimate = time.Now().Add(20 * time.Minute)
	} else if agentInfo.TwinStatus == models.TwinStatusPending {
		agents, err := s.dao.FindAgentInfo(
			daos.GetDBMainCtx(ctx),
			map[string][]any{
				"id < ?": {agentInfo.ID},
				//"agent_contract_id = ?":                            {""},
				//"agent_nft_minted = ?":                             {false},
				`twin_twitter_usernames != '' and twin_status = ?`: {models.TwinStatusPending},
				"scan_enabled = ?": {true},
			},
			map[string][]any{},
			[]string{
				"id asc",
			},
			0,
			999999,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}
		estimate = time.Now().Add(time.Duration(len(agents)) * 10 * time.Minute)

	} else {
		estimate = time.UnixMicro(0)
	}

	return &estimate, nil
}

func (s *Service) GetAgentInfoDetailByContract(ctx context.Context, networkID uint64, agentContractID, agentContractAddress string) (*models.AgentInfo, error) {
	selected := []string{
		"agent_infos.*",
	}
	joinFilters := map[string][]any{}

	filters := map[string][]any{
		`agent_infos.agent_contract_id = ?`:             {agentContractID},
		`lower(agent_infos.agent_contract_address) = ?`: {strings.ToLower(agentContractAddress)},
	}

	if networkID > 0 {
		filters["network_id = ?"] = []any{networkID}
	}

	agent, err := s.dao.FirstAgentInfoJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"TwitterInfo": {},
			"TokenInfo":   {},
		},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	return agent, nil
}

func (s *Service) JobCreateTokenInfo(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobCreateTokenInfo",
		func() error {
			var retErr error
			// create token info for normal agent
			{
				agents, err := s.dao.FindAgentInfoJoin(
					daos.GetDBMainCtx(ctx),
					map[string][]any{
						"join agent_chain_fees on agent_chain_fees.network_id = agent_infos.token_network_id": {},
					},
					map[string][]any{
						`agent_infos.token_status in (?) or (agent_infos.agent_type=2 and agent_infos.status="ready")`: {[]string{"pending", "etching"}},
						"agent_infos.agent_nft_minted = ?":                                      {true},
						`(agent_infos.token_address is null or agent_infos.token_address = "")`: {},
						`agent_infos.token_network_id > 0`:                                      {},
						`(
							(agent_infos.eai_balance > 0 and agent_infos.eai_balance >= agent_chain_fees.token_fee)
							or agent_infos.ref_tweet_id > 0
							or agent_infos.agent_type in (?)
						)`: {
							[]models.AgentInfoAgentType{
								models.AgentInfoAgentTypeModel,
								models.AgentInfoAgentTypeModelOnline,
								models.AgentInfoAgentTypeJs,
								models.AgentInfoAgentTypePython,
								models.AgentInfoAgentTypeInfa,
								models.AgentInfoAgentTypeCustomUi,
								models.AgentInfoAgentTypeCustomPrompt,
							},
						},
						"agent_infos.system_prompt != ''": {},
					},
					map[string][]any{},
					[]string{
						"rand()",
					},
					0,
					50,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, agent := range agents {
					err = s.CreateTokenInfo(ctx, agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
				}
			}
			// create token info for vibe agent
			{
				agents, err := s.dao.FindAgentInfo(
					daos.GetDBMainCtx(ctx),
					map[string][]any{
						`status = ?`:           {models.AssistantStatusReady},
						`agent_nft_minted = ?`: {true},
						`token_address = ''`:   {},
						`token_name = '' or token_image_url = '' or token_image_url = ''`: {},
						`agent_type in (?)`: {
							[]models.AgentInfoAgentType{
								models.AgentInfoAgentTypeModel,
								models.AgentInfoAgentTypeModelOnline,
								models.AgentInfoAgentTypeJs,
								models.AgentInfoAgentTypePython,
								models.AgentInfoAgentTypeInfa,
								models.AgentInfoAgentTypeCustomUi,
								models.AgentInfoAgentTypeCustomPrompt,
							},
						},
						`system_prompt != ''`: {},
					},
					map[string][]any{},
					[]string{
						"rand()",
					},
					0,
					50,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, agent := range agents {
					err = s.CreateTokenInfoVibe(ctx, agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
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

func (s *Service) CreateTokenInfo(ctx context.Context, agentID uint) error {
	err := s.JobRunCheck(
		ctx, fmt.Sprintf("CreateTokenInfo_%d", agentID),
		func() error {
			agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx),
				agentID,
				map[string][]any{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if agentInfo.TokenStatus == "" {
				if agentInfo.AgentType == models.AgentInfoAgentTypeKnowledgeBase &&
					agentInfo.Status == models.AssistantStatusReady &&
					agentInfo.TokenMode == string(models.CreateTokenModeTypeAutoCreate) && agentInfo.TokenNetworkID > 0 {
					updateFields := map[string]any{
						"token_status": "pending",
					}
					err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
						updateFields,
					).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			} else {
				if agentInfo != nil && agentInfo.TokenStatus == "pending" && agentInfo.SystemPrompt != "" {
					if agentInfo.TokenSymbol == "" || agentInfo.TokenName == "" {
						promptGenerateToken := fmt.Sprintf(`
									I want to generate my token base on this info
									'%s'
			
									token-name (generate if not provided, make sure it not empty)
									token-symbol (generate if not provided, make sure it not empty)
									token-story (generate if not provided, make sure it not empty)
			
									Please return in string in json format including token-name, token-symbol, token-story, just only json without explanation  and token name limit with 15 characters
								`, agentInfo.SystemPrompt)
						aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
						if err != nil {
							return errs.NewError(err)
						}
						if aiStr != "" {
							mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
							tokenName := ""
							tokenSymbol := ""
							tokenDesc := ""
							if mapInfo != nil {
								if v, ok := mapInfo["token-name"]; ok {
									tokenName = fmt.Sprintf(`%v`, v)
								}
								if v, ok := mapInfo["token-symbol"]; ok {
									tokenSymbol = fmt.Sprintf(`%v`, v)
								}
								if v, ok := mapInfo["token-story"]; ok {
									tokenDesc = fmt.Sprintf(`%v`, v)
								}
							}
							if tokenDesc != "" && tokenName != "" && tokenSymbol != "" {
								updateFields := map[string]any{
									"token_name": tokenName,
									"token_desc": tokenDesc,
								}
								if agentInfo.TokenSymbol == "" {
									updateFields["token_symbol"] = tokenSymbol
								}
								err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
									updateFields,
								).Error
								if err != nil {
									return errs.NewError(err)
								}
							}
						}
					} else if agentInfo.TokenImageUrl == "" {
						imageUrl, err := s.GetGifImageUrlFromTokenInfo(agentInfo.TokenSymbol, agentInfo.TokenName, agentInfo.TokenDesc)
						if err != nil {
							return errs.NewError(err)
						}
						err = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
							map[string]any{
								"token_image_url": imageUrl,
							},
						).Error
						if err != nil {
							return errs.NewError(err)
						}
					} else if agentInfo.TokenImageUrl != "" && agentInfo.TokenSymbol != "" {
						tokenNetworkID := agentInfo.TokenNetworkID
						if tokenNetworkID == 0 {
							tokenNetworkID = agentInfo.NetworkID
						}
						if tokenNetworkID == models.SOLANA_CHAIN_ID {
							agentTokenAdminAddress := s.conf.GetConfigKeyString(models.GENERTAL_NETWORK_ID, "agent_token_admin_address")
							base64Str, _ := helpers.CurlBase64String(models.GetImageUrl(agentInfo.TokenImageUrl))
							if base64Str != "" {
								tokenFee := big.NewFloat(0)
								if agentInfo.RefTweetID <= 0 {
									agentChainFee, err := s.GetAgentChainFee(
										daos.GetDBMainCtx(ctx),
										tokenNetworkID,
									)
									if err != nil {
										return errs.NewError(err)
									}
									tokenFee = &agentChainFee.TokenFee.Float
								}
								if tokenFee.Cmp(big.NewFloat(0)) > 0 &&
									agentInfo.EaiBalance.Float.Cmp(tokenFee) < 0 {
									return errs.NewError(errs.ErrBadRequest)
								}
								pumfunResp, err := s.blockchainUtils.SolanaCreatePumpfunToken(
									&blockchainutils.SolanaCreatePumpfunTokenReq{
										Address:     agentTokenAdminAddress,
										Name:        agentInfo.TokenName,
										Symbol:      agentInfo.TokenSymbol,
										Description: agentInfo.TokenDesc,
										Twitter:     fmt.Sprintf("https://x.com/%s", agentInfo.TwitterUsername),
										Telegram:    "",
										Website:     "",
										Amount:      0,
										ImageBase64: base64Str,
									},
								)
								if err != nil {
									_ = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
										map[string]any{
											"token_position_hash": err.Error(),
										},
									).Error
								} else {
									_ = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
										map[string]any{
											"token_address":       pumfunResp.Mint,
											"token_status":        "created",
											"token_signature":     pumfunResp.Signature,
											"token_position_hash": "ok",
										},
									).Error
									// if agentInfo.RefTweetID > 0 {
									// 	go s.ReplyAferAutoCreateAgent(daos.GetDBMainCtx(ctx), agentInfo.RefTweetID, agentInfo.ID)
									// } else {
									// 	// TODO: post twitter
									// 	go s.PostTwitterAferCreateToken(ctx, agentInfo.ID)
									// }
									if tokenFee.Cmp(big.NewFloat(0)) > 0 {
										_ = func() error {
											_ = daos.GetDBMainCtx(ctx).
												Model(agentInfo).
												Updates(
													map[string]any{
														"eai_balance": gorm.Expr("eai_balance - ?", numeric.NewBigFloatFromFloat(tokenFee)),
													},
												).Error
											_ = s.dao.Create(
												daos.GetDBMainCtx(ctx),
												&models.AgentEaiTopup{
													NetworkID:      agentInfo.NetworkID,
													EventId:        fmt.Sprintf("agent_token_fee_%d", agentInfo.ID),
													AgentInfoID:    agentInfo.ID,
													Type:           models.AgentEaiTopupTypeSpent,
													Amount:         numeric.NewBigFloatFromFloat(tokenFee),
													Status:         models.AgentEaiTopupStatusDone,
													DepositAddress: agentInfo.ETHAddress,
													ToAddress:      agentInfo.ETHAddress,
													Toolset:        "token_fee",
												},
											)
											return nil
										}()
									}
								}
							}
						} else {
							// create meme
							_, err := s.CreateMeme(
								ctx, agentInfo.Creator,
								tokenNetworkID,
								&serializers.MemeReq{
									Name:            agentInfo.TokenName,
									Ticker:          agentInfo.TokenSymbol,
									Description:     agentInfo.TokenDesc,
									Image:           agentInfo.TokenImageUrl,
									Twitter:         fmt.Sprintf("https://x.com/%s", agentInfo.TwitterUsername),
									AgentInfoID:     agentInfo.ID,
									BaseTokenSymbol: string(models.BaseTokenSymbolEAI),
								})
							if err != nil {
								_ = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
									map[string]any{
										"token_position_hash": err.Error(),
									},
								).Error
							} else {
								_ = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
									map[string]any{
										"token_status":        "created",
										"token_position_hash": "ok",
									},
								).Error
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
	return nil
}

func (s *Service) CreateTokenInfoVibe(ctx context.Context, agentID uint) error {
	err := s.JobRunCheck(
		ctx, fmt.Sprintf("CreateTokenInfoVibe_%d", agentID),
		func() error {
			agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx),
				agentID,
				map[string][]any{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if agentInfo == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			if !agentInfo.IsVibeAgent() {
				return errs.NewError(errs.ErrBadRequest)
			}
			if agentInfo.SystemPrompt != "" {
				if agentInfo.TokenSymbol == "" || agentInfo.TokenName == "" {
					promptGenerateToken := fmt.Sprintf(`
									I want to generate my token base on this info
									'%s'
			
									token-name (generate if not provided, make sure it not empty)
									token-symbol (generate if not provided, make sure it not empty)
									token-story (generate if not provided, make sure it not empty)
			
									Please return in string in json format including token-name, token-symbol, token-story, just only json without explanation  and token name limit with 15 characters
								`, agentInfo.SystemPrompt)
					aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
					if err != nil {
						return errs.NewError(err)
					}
					if aiStr != "" {
						mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
						tokenName := ""
						tokenSymbol := ""
						tokenDesc := ""
						if mapInfo != nil {
							if v, ok := mapInfo["token-name"]; ok {
								tokenName = fmt.Sprintf(`%v`, v)
							}
							if v, ok := mapInfo["token-symbol"]; ok {
								tokenSymbol = fmt.Sprintf(`%v`, v)
							}
							if v, ok := mapInfo["token-story"]; ok {
								tokenDesc = fmt.Sprintf(`%v`, v)
							}
						}
						if tokenDesc != "" && tokenName != "" && tokenSymbol != "" {
							updateFields := map[string]any{
								"token_name": tokenName,
								"token_desc": tokenDesc,
							}
							if agentInfo.TokenSymbol == "" {
								updateFields["token_symbol"] = tokenSymbol
							}
							err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
								updateFields,
							).Error
							if err != nil {
								return errs.NewError(err)
							}
						}
					}
				} else if agentInfo.TokenImageUrl == "" {
					imageUrl, err := s.GetGifImageUrlFromTokenInfo(agentInfo.TokenSymbol, agentInfo.TokenName, agentInfo.TokenDesc)
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
						map[string]any{
							"token_image_url": imageUrl,
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

func (s *Service) getCreateTokenTwitterPost(agentName, networkName, tokenName, tokenSymbol, tokenAddress string) string {
	return strings.TrimSpace(
		fmt.Sprintf(`
	Hello World!

	I'm %s, a fully on-chain AI agent living on %s!

	Guess what? I've got my own token now—called %s! Its ticker? $%s

	To all my fellow Eternals from @CryptoEternalAI, let's rally together!

	https://pump.fun/coin/%s
	
		`, agentName, networkName, tokenName, tokenSymbol, tokenAddress),
	)
}

func (s *Service) getCreateTokenTwitterPostFromAI(tokenDesc, tokenSymbol string) string {
	promptGenerateToken := fmt.Sprintf(`
	I want to announce my own token . 
	Please make a tweet exclude hashtag on twitter max 200 character with this info
	
	Token symbol: %s.
	Token story: %s. 
	
	Please return in json format including tweet. Just only json without explanation
	`, tokenDesc, tokenSymbol)
	aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(aiStr)
}

func (s *Service) PostTwitterAferCreateToken(ctx context.Context, agentInfoID uint) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]any{
			"TwitterInfo": {},
		},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}

	if agentInfo != nil && agentInfo.TwitterInfo != nil {
		content := ""
		aiStr := s.getCreateTokenTwitterPostFromAI(agentInfo.TokenDesc, agentInfo.TokenSymbol)
		mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
		if mapInfo != nil {
			if tweet, ok := mapInfo["tweet"]; ok {
				content = fmt.Sprintf(`
				%s 

				https://pump.fun/coin/%s
				`, tweet, agentInfo.TokenAddress)

				content = strings.TrimSpace(content)
			}
		}

		if content == "" {
			content = s.getCreateTokenTwitterPost(agentInfo.AgentName, agentInfo.NetworkName, agentInfo.TokenName, agentInfo.TokenSymbol, agentInfo.TokenAddress)
		}

		tweetID, err := helpers.PostTweetByToken(agentInfo.TwitterInfo.AccessToken, content, "")
		if err != nil {
			daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
				map[string]any{
					"token_position_hash": err.Error(),
				},
			)
		}
		if tweetID != "" {
			agentPost := &models.AgentTwitterPost{
				NetworkID:       agentInfo.NetworkID,
				AgentInfoID:     agentInfo.ID,
				TwitterID:       agentInfo.TwitterInfo.TwitterID,
				TwitterUsername: agentInfo.TwitterInfo.TwitterUsername,
				TwitterName:     agentInfo.TwitterInfo.TwitterName,
				TwitterPostID:   tweetID,
				PostAt:          helpers.TimeNow(),
				Content:         content,
				ReplyContent:    content,
				PostType:        models.AgentSnapshotPostActionTypeTweet,
				ReplyPostId:     tweetID,
				ReplyPostAt:     helpers.TimeNow(),
				Status:          models.AgentTwitterPostStatusReplied,
				IsMigrated:      true,
			}
			_ = s.dao.Create(daos.GetDBMainCtx(ctx), agentPost)

			_ = s.AgentTeleAlertNewTokenByID(ctx, agentInfoID)
		}
	}
	return nil
}

func (s *Service) GetDashboardAgentInfos(ctx context.Context, contractAddresses []string, userAddress string, networkID uint64, agentType int, agentTypes []int,
	tokenAddress, search, agentModel string, installed *bool, ids, exludeIds []uint, categoryIds []string, includeHidden *bool, sortListStr []string, page, limit int) ([]*models.AgentInfo, uint, error) {
	selected := []string{
		`ifnull(agent_infos.reply_latest_time, agent_infos.updated_at) reply_latest_time`,
		"agent_infos.*",
		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), ifnull(agent_token_infos.price_change,0)) meme_percent`,
		`cast(case when ifnull(agent_token_infos.usd_market_cap, 0) then ifnull(agent_token_infos.usd_market_cap, 0)
		when ifnull(memes.price_usd*memes.total_suply, 0) > 0 then ifnull(memes.price_usd*memes.total_suply, 0) end as decimal(36, 18)) meme_market_cap`,
		`ifnull(memes.price_usd, agent_token_infos.price_usd) meme_price`,
		`ifnull(memes.volume_last24h*memes.price_usd, agent_token_infos.volume_last24h) meme_volume_last24h`,
	}
	joinFilters := map[string][]any{
		`
			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
			left join twitter_users on twitter_users.twitter_id = agent_infos.tmp_twitter_id and  agent_infos.tmp_twitter_id is not null
		`: {},
	}

	filters := map[string][]any{
		`	
			((agent_infos.agent_contract_address is not null and agent_infos.agent_contract_address != "") or (agent_infos.agent_type=2 and agent_infos.status="ready") or agent_infos.token_address != "")
			and ifnull(agent_infos.priority, 0) >= 0
			and agent_infos.id != 15
		`: {},
		`agent_infos.token_address != "" and ifnull(memes.status, "") not in ("created", "pending")`: {},
		`agent_infos.agent_type != ?`: {models.AgentInfoAgentTypeVideo},
	}

	if search != "" {
		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
		filters[`
			LOWER(agent_infos.token_name) like ? 
			or LOWER(agent_infos.token_symbol) like ? 
			or LOWER(agent_infos.token_address) like ?
			or LOWER(agent_infos.twitter_username) like ?
			or LOWER(agent_infos.agent_name) like ?
			or ifnull(twitter_users.twitter_username, "") like ?
			or ifnull(twitter_users.name, "") like ?
			or LOWER(agent_infos.display_name) like ? 
		`] = []any{search, search, search, search, search, search, search, search}
	}

	//filter agent type
	if agentType != -1 {
		if agentType > 0 {
			filters["agent_infos.agent_type = ?"] = []any{agentType}
		} else if len(agentTypes) > 0 {
			filters["agent_infos.agent_type in (?)"] = []any{agentTypes}
		} else {
			filters["agent_infos.agent_type not in (?)"] = []any{[]models.AgentInfoAgentType{
				models.AgentInfoAgentTypeModel,
				models.AgentInfoAgentTypeModelOnline,
				models.AgentInfoAgentTypeJs,
				models.AgentInfoAgentTypePython,
				models.AgentInfoAgentTypeInfa,
				models.AgentInfoAgentTypeCustomUi,
				models.AgentInfoAgentTypeCustomPrompt,
			}}
		}
	}

	//filter contract address
	if len(contractAddresses) > 0 {
		filters["agent_infos.agent_contract_address in (?)"] = []any{contractAddresses}
	}

	//filter agent model
	if agentModel != "" {
		filters["agent_infos.agent_base_model = ?"] = []any{agentModel}
	}

	if tokenAddress != "" {
		filters["LOWER(agent_infos.token_address) = ? or agent_infos.agent_id = ? or agent_infos.id = ?"] = []any{strings.ToLower(tokenAddress), tokenAddress, tokenAddress}
	} else {
		filters[`(agent_infos.agent_contract_id is not null and agent_infos.agent_contract_id != "") or (agent_infos.token_address is not null and agent_infos.token_address != "") or (agent_infos.agent_type=2 and agent_infos.status="ready")`] = []any{}
		if networkID == models.SOLANA_CHAIN_ID_OLD {
			networkID = models.SOLANA_CHAIN_ID
		}

		if networkID > 0 {
			if networkID == models.SHARDAI_CHAIN_ID || networkID == models.SOLANA_CHAIN_ID {
				filters["agent_infos.network_id = ? or agent_infos.token_network_id = ? or agent_infos.id = 763"] = []any{networkID, networkID}
			} else if networkID == models.HERMES_CHAIN_ID {
				filters["agent_infos.network_id = ? or agent_infos.token_network_id = ?"] = []any{networkID, networkID}
				filters["agent_infos.id != 763"] = []any{}
			} else if networkID == models.BASE_CHAIN_ID {
				filters["agent_infos.network_id = ? or agent_infos.token_network_id = ?"] = []any{networkID, networkID}
				filters["agent_infos.id != 763"] = []any{}
			} else if networkID == models.ETHEREUM_CHAIN_ID {
				listEtherModels := []uint64{models.ETHEREUM_CHAIN_ID, models.BASE_CHAIN_ID, models.APE_CHAIN_ID, models.ABSTRACT_TESTNET_CHAIN_ID, models.ARBITRUM_CHAIN_ID}
				filters["agent_infos.network_id in (?) or agent_infos.token_network_id in (?)"] = []any{listEtherModels, listEtherModels}
			} else {
				filters["agent_infos.network_id = ? or agent_infos.token_network_id = ?"] = []any{networkID, networkID}
			}
		}
	}

	if len(ids) > 0 {
		filters["agent_infos.id in (?)"] = []any{ids}
	} else {
		if includeHidden == nil || !(*includeHidden) {
			filters["agent_infos.is_public = 1"] = []any{}
		}
		//filter instlled app
		// if installed != nil && userAddress != "" {
		// 	if *installed {
		// 		joinFilters = map[string][]any{
		// 			`
		// 			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
		// 			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
		// 			left join twitter_users on twitter_users.twitter_id = agent_infos.tmp_twitter_id and  agent_infos.tmp_twitter_id is not null
		// 			join agent_utility_installs on agent_utility_installs.agent_info_id = agent_infos.id
		// 					and agent_utility_installs.deleted_at IS NULL and agent_utility_installs.address = ?
		// 		`: {strings.ToLower(userAddress)},
		// 		}
		// 	} else {
		// 		filters["agent_infos.id not in (select agent_info_id from agent_utility_installs where address = ?)"] = []any{strings.ToLower(userAddress)}
		// 		filters["agent_infos.is_public = 1"] = []any{}
		// 	}
		// } else {
		// 	filters["agent_infos.is_public = 1"] = []any{}
		// }
	}

	if len(exludeIds) > 0 {
		filters["agent_infos.id not in (?)"] = []any{exludeIds}
	}

	if len(categoryIds) > 0 {
		filters["agent_infos.agent_category_id in (?)"] = []any{categoryIds}
	}

	if userAddress != "" {
		joinFilters = map[string][]any{
			`
			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
			left join twitter_users on twitter_users.twitter_id = agent_infos.tmp_twitter_id and  agent_infos.tmp_twitter_id is not null
			left join agent_utility_recent_chats on agent_utility_recent_chats.agent_info_id = agent_infos.id
					and agent_utility_recent_chats.address = ?
		`: {strings.ToLower(userAddress)},
		}
		sortRecentChat := "ifnull(agent_utility_recent_chats.updated_at, now() - interval 100 day) desc"
		newSortListStr := make([]string, 0, len(sortListStr)+1)
		newSortListStr = append(newSortListStr, sortRecentChat)
		newSortListStr = append(newSortListStr, sortListStr...)
		sortListStr = newSortListStr
	}

	sortDefault := "ifnull(agent_infos.priority, 0) desc, meme_market_cap desc"
	if len(sortListStr) > 0 {
		sortDefault = strings.Join(sortListStr, ", ")
	}
	agents, err := s.dao.FindAgentInfoJoinSelect(
		daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"TwitterInfo":    {},
			"TmpTwitterInfo": {},
			"Meme":           {`deleted_at IS NULL and status not in ("created", "pending")`},
			"TokenInfo":      {},
			"AgentCategory":  {},
		},
		[]string{sortDefault},
		page, limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return agents, 0, nil
}

func (s *Service) GetTokenInfoByContract(ctx context.Context, tokenAddress string) (*dexscreener.PairsDetailResp, error) {
	var coinInfo dexscreener.PairsDetailResp
	err := s.RedisCached(
		fmt.Sprintf("GetTokenInfoByContract_%s", tokenAddress),
		true,
		1*time.Hour,
		&coinInfo,
		func() (any, error) {
			rs, err := s.dexscreener.SearchPairs(tokenAddress)
			if err != nil {
				return nil, errs.NewError(err)
			}
			if rs != nil {
				tokenDescription := ""
				rs.NetworkID = models.GetChainID(rs.ChainId)
				if rs.Info != nil {
					for _, item := range rs.Info.Websites {
						if strings.EqualFold(item.Label, "Website") {
							webContent := helpers.ContentHtmlByUrl(item.Url)
							if webContent == "" {
								webContent = helpers.RodContentHtmlByUrl(item.Url)
							}
							if webContent != "" {
								cleanWebContent, _ := s.blockchainUtils.CleanHtml(webContent)
								if cleanWebContent != "" {
									tokenDescription, _ = s.openais["Lama"].SummaryWebContent(cleanWebContent)
								}
							}
						}
					}

					if tokenDescription == "" {
						for _, item := range rs.Info.Socials {
							if item.Type == "twitter" {
								twiterPostIDArry := strings.Split(item.Url, "/")
								if len(twiterPostIDArry) > 0 {
									twitteUserName := twiterPostIDArry[len(twiterPostIDArry)-1]
									twiterPostIDArry = strings.Split(twitteUserName, "?")
									if len(twiterPostIDArry) > 0 {
										twitteUserName = twiterPostIDArry[0]
										twitterUser, _ := s.twitterWrapAPI.GetTwitterByUserName(twitteUserName)
										if twitterUser != nil {
											tokenDescription = twitterUser.Description
										}
									}

								}
							}
						}
					}

				}
				rs.Description = tokenDescription
				coinInfo = *rs
				return coinInfo, nil
			}
			return nil, errs.NewError(errs.ErrTokenNotFound)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return &coinInfo, nil
}

func (s *Service) GetWebpageText(ctx context.Context, url string) (string, error) {
	var coinInfo string
	err := s.RedisCached(
		fmt.Sprintf("GetWebpageText_%s", url),
		true,
		1*time.Hour,
		&coinInfo,
		func() (any, error) {
			webContent := helpers.ContentHtmlByUrl(url)
			if webContent == "" {
				webContent = helpers.RodContentHtmlByUrl(url)
			}
			if webContent != "" {
				coinInfo, _ = s.blockchainUtils.CleanHtml(webContent)
			}
			return coinInfo, nil
		},
	)
	if err != nil {
		return "", errs.NewError(err)
	}
	return coinInfo, nil
}

func (s *Service) JobUpdateTokenPriceInfo(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateTokenPriceInfo",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					`
						agent_contract_id is not null and agent_contract_id != ""
						and agent_contract_address is not null and agent_contract_address != ""
						and token_address is not null and token_address != ""
						and scan_enabled = 1
						and agent_nft_minted = 1
					`: {},
				},
				map[string][]any{},
				[]string{},
				0,
				9999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = s.UpdateTokenPriceInfo(ctx, agent.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewError(err))
				}
				time.Sleep(1 * time.Second)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateTokenPriceInfo(ctx context.Context, agentID uint) error {
	err := s.JobRunCheck(
		ctx, fmt.Sprintf("UpdateTokenPriceInfo_%d", agentID),
		func() error {
			agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx),
				agentID,
				map[string][]any{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo.TokenAddress != "" && agentInfo.TokenInfoID > 0 {
				coinInfo, err := s.dexscreener.SearchPairs(agentInfo.TokenAddress)
				if err != nil {
					return errs.NewError(err)
				}
				if coinInfo != nil && coinInfo.ChainId != "" {
					_ = daos.GetDBMainCtx(ctx).Model(&models.AgentTokenInfo{}).Where("id = ?", agentInfo.TokenInfoID).Updates(
						map[string]any{
							"price_usd":         numeric.NewBigFloatFromString(coinInfo.PriceUsd),
							"price":             numeric.NewBigFloatFromString(coinInfo.PriceNative),
							"usd_market_cap":    coinInfo.MarketCap,
							"pool_address":      coinInfo.PairAddress,
							"dex_url":           coinInfo.Url,
							"volume_last24h":    numeric.NewBigFloatFromFloat(big.NewFloat(coinInfo.Volume.H24)),
							"base_token_symbol": coinInfo.QuoteToken.Symbol,
							"price_change":      coinInfo.PriceChange.H24,
							"dex_id":            coinInfo.DexId,
						},
					)
				} else if agentInfo.TokenNetworkID == models.SOLANA_CHAIN_ID {
					coinInfo, err := s.pumfunAPI.GetPumpFunCoinInfo(agentInfo.TokenAddress)
					if err != nil {
						return errs.NewError(err)
					}
					if coinInfo == nil || coinInfo.Mint == "" {
						return errs.NewError(err)
					}
					tknDecimals, err := s.GetSolanaTokenDecimals(agentInfo.TokenAddress)
					if err != nil {
						return errs.NewError(err)
					}
					coinPrice := coinInfo.UsdMarketCap / (float64(coinInfo.TotalSupply) / math.Pow10(tknDecimals))
					_ = daos.GetDBMainCtx(ctx).Model(&models.AgentTokenInfo{}).Where("id = ?", agentInfo.TokenInfoID).Updates(
						map[string]any{
							"price_usd":      numeric.NewBigFloatFromFloat(big.NewFloat(coinPrice)),
							"usd_market_cap": coinInfo.UsdMarketCap,
							"total_supply":   coinInfo.TotalSupply,
							"dex_id":         "pumpfun",
						},
					)
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

func (s *Service) JobUpdateAgentImage(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateTokenPriceInfo",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					`
						(token_image_url is null or token_image_url="")
						and (thumbnail is NULL or thumbnail="")
						and (nft_token_image is NULL or nft_token_image="")
						and twitter_info_id = 0
						and agent_nft_minted = 1
						and network_id not in (43338, 0, 1, 45761, 222671)
						and network_id = 43114
						and system_prompt != ""
					`: {},
				},
				map[string][]any{},
				[]string{"created_at desc"},
				0,
				200,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				tokenInfo, _ := s.GenerateTokenInfoFromSystemPrompt(ctx, agent.AgentName, agent.SystemPrompt)
				if tokenInfo != nil && tokenInfo.TokenImageUrl != "" {
					_ = daos.GetDBMainCtx(ctx).Model(agent).Updates(
						map[string]any{
							"token_image_url": tokenInfo.TokenImageUrl,
						},
					).Error
				}
				time.Sleep(20 * time.Millisecond)
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
