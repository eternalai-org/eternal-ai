package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"go.uber.org/zap"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/aidojo"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (s *Service) GetAgentChainFee(tx *gorm.DB, networkID uint64) (*models.AgentChainFee, error) {
	agentChainFee, err := s.dao.FirstAgentChainFee(
		tx,
		map[string][]interface{}{
			"network_id = ?": {networkID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentChainFee == nil {
		agentChainFee = &models.AgentChainFee{
			NetworkID: networkID,
		}
		err = s.dao.Create(tx, agentChainFee)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return agentChainFee, nil
}

func (s *Service) JobAgentSnapshotPostCreate(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentSnapshotPostCreate",
		func() error {
			missions, err := s.dao.FindAgentSnapshotMissionJoin(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`
					join agent_infos on agent_infos.id = agent_snapshot_missions.agent_info_id
					join agent_snapshot_mission_configs on agent_snapshot_mission_configs.network_id = agent_snapshot_missions.network_id and agent_snapshot_mission_configs.tool_set = agent_snapshot_missions.tool_set
					join agent_chain_fees on agent_chain_fees.network_id = agent_infos.network_id
					left join twitter_infos on twitter_infos.id = agent_infos.twitter_info_id
					`: {},
				},
				map[string][]interface{}{
					"agent_snapshot_missions.enabled = 1":       {},
					"agent_snapshot_missions.reply_enabled = 1": {},
					"agent_snapshot_missions.tool_set != '' and agent_snapshot_missions.tool_set != 'trade_analytics_mentions' ": {},
					"agent_snapshot_missions.interval_sec > 0": {},
					`(
						agent_infos.agent_type = 1
						and agent_infos.agent_contract_id != ''
						and agent_infos.scan_enabled = 1
						and agent_infos.reply_enabled = 1
						and agent_infos.eai_balance > 0
						and agent_chain_fees.infer_fee > 0
						and agent_infos.eai_balance >= agent_chain_fees.infer_fee
						and (
							(
								agent_infos.twitter_info_id > 0 and twitter_infos.expired_at > adddate(now(), interval -15 minute) 
								and (
									agent_snapshot_mission_configs.platform = 'twitter'
									or agent_snapshot_mission_configs.platform = 'defi'
								)
							)
							or (agent_infos.farcaster_id is not null and agent_infos.farcaster_id != '' and agent_snapshot_mission_configs.platform = 'farcaster')
							or agent_snapshot_missions.agent_store_mission_id > 0
						)
						and agent_infos.network_id in (?)
					)`: {
						[]uint64{
							models.SHARDAI_CHAIN_ID,
							models.ETHEREUM_CHAIN_ID,
							models.BASE_CHAIN_ID,
							models.HERMES_CHAIN_ID,
							models.ARBITRUM_CHAIN_ID,
							models.ZKSYNC_CHAIN_ID,
							models.POLYGON_CHAIN_ID,
							models.SOLANA_CHAIN_ID,
							models.BSC_CHAIN_ID,
							models.APE_CHAIN_ID,
							models.AVALANCHE_C_CHAIN_ID,
							models.ABSTRACT_TESTNET_CHAIN_ID,
							models.BITTENSOR_CHAIN_ID,
							models.DUCK_CHAIN_ID,
							models.TRON_CHAIN_ID,
							models.MODE_CHAIN_ID,
							models.ZETA_CHAIN_ID,
							models.STORY_CHAIN_ID,
							models.HYPE_CHAIN_ID,
							models.MONAD_TESTNET_CHAIN_ID,
							models.MEGAETH_TESTNET_CHAIN_ID,
							models.CELO_CHAIN_ID,
							models.BASE_SEPOLIA_CHAIN_ID,
						},
					},
					`agent_snapshot_missions.infer_at is null
						or agent_snapshot_missions.infer_at <= adddate(now(), interval -agent_snapshot_missions.interval_sec second)
					`: {},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				999999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			if len(missions) > 0 {
				for _, mission := range missions {
					err = s.AgentSnapshotPostCreate(ctx, mission.ID, "", "")
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, mission.ID))
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

func (s *Service) JobAgentSnapshotPostActionExecuted(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentSnapshotPostActionExecuted",
		func() error {
			var retErr error
			actions, err := s.dao.FindAgentSnapshotPostActionJoin(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"join agent_snapshot_missions on agent_snapshot_missions.id = agent_snapshot_post_actions.agent_snapshot_mission_id": {},
					"join agent_infos on agent_infos.id = agent_snapshot_post_actions.agent_info_id":                                     {},
				},
				map[string][]interface{}{
					"agent_snapshot_post_actions.status = ?":                                     {models.AgentSnapshotPostActionStatusNew},
					"agent_snapshot_post_actions.schedule_at < now()":                            {},
					"agent_snapshot_post_actions.schedule_at > adddate(now(), interval -6 hour)": {},
					`(
						agent_infos.agent_type = 1
						and agent_infos.reply_enabled = true
						and (
								agent_infos.reply_latest_time is null
								or agent_infos.reply_latest_time <= adddate(now(), interval -agent_infos.action_delayed second)
								or agent_snapshot_post_actions.tool_set in ('follow', 'reply_mentions', 'inscribe_tweet', 'post', 'trading', 'post_search', 'trade_analytics', 'trade_analytics_twitter', 'trade_analytics_mentions', 'lucky_moneys')
								or agent_snapshot_post_actions.type in ('reply_multi_unlimited')
								or agent_snapshot_missions.not_delay = true
							)
					)`: {},
					`(
						agent_snapshot_missions.enabled = 1
						and agent_snapshot_missions.reply_enabled = 1
						and agent_snapshot_missions.interval_sec > 0
						and agent_snapshot_missions.is_testing = 0
						and agent_snapshot_missions.deleted_at is null
					)`: {},
				},
				map[string][]interface{}{},
				[]string{
					"rand()",
				},
				0,
				100,
			)
			if err != nil {
				return errs.NewError(err)
			}
			for _, action := range actions {
				err = s.AgentSnapshotPostActionExecuted(ctx, action.ID)
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) JobAgentSnapshotPostActionDupplicated(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentSnapshotPostActionDupplicated",
		func() error {
			var retErr error
			actions, err := s.dao.FindAgentSnapshotPostAction(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"status = ?": {models.AgentSnapshotPostActionStatusNew},
					`(
						(
							type in (
								'reply', 'retweet', 'quote_tweet'
							) and exists(
									select 1
									from agent_snapshot_post_actions checks
									where agent_snapshot_post_actions.agent_info_id = checks.agent_info_id
									and agent_snapshot_post_actions.type = checks.type
									and agent_snapshot_post_actions.conversation_id = checks.conversation_id
									and agent_snapshot_post_actions.id > checks.id
									and checks.status in ('new', 'done')
								)
							)
						or (
							type in (
								'tweet'
							) and exists(
									select 1
									from agent_snapshot_post_actions checks
									where agent_snapshot_post_actions.agent_info_id = checks.agent_info_id
									and agent_snapshot_post_actions.agent_snapshot_mission_id = checks.agent_snapshot_mission_id
									and agent_snapshot_post_actions.type = checks.type
									and agent_snapshot_post_actions.id < checks.id
									and checks.status in ('new')
								)
							)
						or (
							type in (
								'follow'
							) and exists(
									select 1
									from agent_snapshot_post_actions checks
									where agent_snapshot_post_actions.agent_info_id = checks.agent_info_id
									and agent_snapshot_post_actions.type = checks.type
									and agent_snapshot_post_actions.target_username = checks.target_username
									and agent_snapshot_post_actions.id > checks.id
									and checks.status in ('new', 'done')
								)
							)
						or (
							type in (
								'create_token'
							) and exists(
									select 1
									from agent_snapshot_post_actions checks
									where agent_snapshot_post_actions.agent_info_id = checks.agent_info_id
									and agent_snapshot_post_actions.type = checks.type
									and agent_snapshot_post_actions.id < checks.id
									and checks.status in ('new')
								)
							)
					)`: {},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				999999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			for _, action := range actions {
				err = daos.
					GetDBMainCtx(ctx).
					Model(action).
					Where("status = ?", models.AgentSnapshotPostActionStatusNew).
					UpdateColumn("status", models.AgentSnapshotPostActionStatusDoneDuplicated).
					Error
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) JobAgentSnapshotPostActionCancelled(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentSnapshotPostActionCancelled",
		func() error {
			var retErr error
			actions, err := s.dao.FindAgentSnapshotPostAction(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"status = ?": {models.AgentSnapshotPostActionStatusNew},
					"schedule_at < adddate(now(), interval -6 hour)": {},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				999999,
			)
			if err != nil {
				return errs.NewError(err)
			}
			for _, action := range actions {
				err = daos.GetDBMainCtx(ctx).
					Model(action).
					Where("status = ?", models.AgentSnapshotPostActionStatusNew).
					UpdateColumn("status", models.AgentSnapshotPostActionStatusDoneCancelled).
					Error
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) AgentSnapshotPostCreate(ctx context.Context, missionID uint, orgTweetID, tokenSymbol string) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentSnapshotPostCreate_%d", missionID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					mission, err := s.dao.FirstAgentSnapshotMissionByID(
						tx,
						missionID,
						map[string][]interface{}{
							"AgentInfo":         {},
							"AgentStore":        {},
							"AgentStoreMission": {},
						},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					agentInfo := mission.AgentInfo
					if agentInfo == nil {
						return errs.NewError(errs.ErrBadRequest)
					}
					var headSystemPrompt string
					if agentInfo.AgentName != "" && agentInfo.TwitterUsername != "" {
						headSystemPrompt = headSystemPrompt + "Your Twitter name is <twitter_name>. Your Twitter username is @<twitter_username>. People refer to you as <twitter_name>, @<twitter_username>, <twitter_username>.\n\n"
					}
					metaDataReq := &aidojo.AgentMetadataRequest{}
					if agentInfo.TokenSymbol != "" &&
						agentInfo.TokenAddress != "" &&
						agentInfo.TwitterUsername != "0xIntellect" {
						tokenNetworkName := cases.Title(language.English).String(models.GetChainName(agentInfo.TokenNetworkID))
						headSystemPrompt = headSystemPrompt + "You have a token. Your token name is <token_name>. Your token ticker is $<token_ticker>. People refer to your token as <token_name> or $<token_ticker>. Your token address is <token_address>. Your token was deployed on " + tokenNetworkName + ".\n\n"
						metaDataReq.TokenInfo.Name = agentInfo.TokenName
						metaDataReq.TokenInfo.Symbol = agentInfo.TokenSymbol
						metaDataReq.TokenInfo.Address = agentInfo.TokenAddress
						metaDataReq.TokenInfo.Chain = tokenNetworkName
					}
					headSystemPrompt = strings.TrimSpace(headSystemPrompt)
					headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<twitter_name>", agentInfo.AgentName)
					headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<twitter_username>", agentInfo.TwitterUsername)
					headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<token_name>", agentInfo.TokenName)
					headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<token_ticker>", agentInfo.TokenSymbol)
					headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<token_address>", agentInfo.TokenAddress)

					inferTxHash := helpers.RandomBigInt(12).Text(16)
					if mission.ToolList != "" &&
						(mission.ToolSet == models.ToolsetTypeTradeAnalytics || mission.ToolSet == models.ToolsetTypeTradeAnalyticsOnTwitter || mission.ToolSet == models.ToolsetTypeTradeAnalyticsMentions) {
						if mission.Tokens == "" {
							mission.Tokens = tokenSymbol
						}
						mission.UserPrompt = strings.ReplaceAll(mission.UserPrompt, "{token_symbol}", mission.Tokens)
						mission.ToolList = strings.ReplaceAll(mission.ToolList, "{ref_id}", inferTxHash)
						mission.ToolList = strings.ReplaceAll(mission.ToolList, "{token_symbol}", mission.Tokens)
					}

					inferItems := []*models.UserAgentInferDataItem{
						{
							Role:    "user",
							Content: strings.TrimSpace(mission.UserPrompt),
						},
					}
					agentChainFee, err := s.GetAgentChainFee(
						tx,
						agentInfo.NetworkID,
					)
					if err != nil {
						return errs.NewError(err)
					}
					inferFee := agentChainFee.InferFee
					missionStoreFee := numeric.NewBigFloatFromFloat(big.NewFloat(0))
					toolList := mission.ToolList
					if mission.AgentStoreMissionID > 0 {
						missionStoreFee = mission.AgentStoreMission.Price
						inferFee = numeric.NewBigFloatFromFloat(models.AddBigFloats(&inferFee.Float, &missionStoreFee.Float))
						agentStoreInstall, err := s.dao.FirstAgentStoreInstall(
							tx,
							map[string][]interface{}{
								"type = ?":           {models.AgentStoreInstallTypeAgent},
								"agent_store_id = ?": {mission.AgentStoreID},
								"agent_info_id = ?":  {mission.AgentInfoID},
								"status = ?":         {models.AgentStoreInstallStatusDone},
							},
							map[string][]interface{}{
								"User": {},
							},
							[]string{"id desc"},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agentStoreInstall == nil {
							return errs.NewError(errs.ErrBadRequest)
						}
						//
						agentStoreMission, err := s.dao.FirstAgentStoreMissionByID(
							tx,
							mission.AgentStoreMissionID,
							map[string][]interface{}{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agentStoreMission == nil {
							return errs.NewError(errs.ErrBadRequest)
						}
						agentStoreMission.NumUsed++
						err = s.dao.Save(tx, agentStoreMission)
						if err != nil {
							return errs.NewError(err)
						}
						params := map[string]interface{}{}
						err = helpers.ConvertJsonObject(agentStoreInstall.CallbackParams, &params)
						if err != nil {
							return errs.NewError(err)
						}
						params["install_code"] = agentStoreInstall.Code
						if agentStoreInstall.User != nil {
							params["user_address"] = agentStoreInstall.User.Address
						}
						inferItems[0].Content, err = helpers.GenerateTemplateContent(mission.AgentStoreMission.UserPrompt, params)
						if err != nil {
							return errs.NewError(err)
						}
						if mission.UserPrompt != "" {
							inferItems[0].Content, err = helpers.GenerateTemplateContent(mission.UserPrompt, params)
							if err != nil {
								return errs.NewError(err)
							}
						}
						toolList, err = helpers.GenerateTemplateContent(mission.AgentStoreMission.ToolList, params)
						if err != nil {
							return errs.NewError(err)
						}
					}
					inferData, err := json.Marshal(inferItems)
					if err != nil {
						return errs.NewError(err)
					}
					inferPost := &models.AgentSnapshotPost{
						NetworkID:              agentInfo.NetworkID,
						AgentInfoID:            agentInfo.ID,
						AgentSnapshotMissionID: mission.ID,
						InferData:              string(inferData),
						InferAt:                helpers.TimeNow(),
						Status:                 models.AgentSnapshotPostStatusInferSubmitted,
						Fee:                    inferFee,
						UserPrompt:             inferItems[0].Content,
						HeadSystemPrompt:       headSystemPrompt,
						AgentMetaData:          helpers.ConvertJsonString(metaDataReq),
						ToolList:               toolList,
						SystemPrompt:           agentInfo.SystemPrompt,
						SystemReminder:         agentInfo.SystemReminder,
						Toolset:                string(mission.ToolSet),
						AgentBaseModel:         mission.AgentBaseModel,
						ReactMaxSteps:          mission.ReactMaxSteps,
						InferTxHash:            inferTxHash,
						OrgTweetID:             orgTweetID,
						Token:                  tokenSymbol,
						AgentStoreMissionID:    mission.AgentStoreMissionID,
						AgentStoreID:           mission.AgentStoreID,
						IsRated:                false,
						AgentStoreMissionFee:   missionStoreFee,
					}
					if inferPost.AgentBaseModel == "" {
						inferPost.AgentBaseModel = agentInfo.AgentBaseModel
					}
					if inferPost.Fee.Float.Cmp(big.NewFloat(0)) <= 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					if agentInfo.EaiBalance.Float.Cmp(&inferPost.Fee.Float) < 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					err = tx.Model(agentInfo).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance - ?", inferPost.Fee)).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					agentInfo, err = s.dao.FirstAgentInfoByID(
						tx,
						agentInfo.ID,
						map[string][]interface{}{},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if agentInfo.EaiBalance.Float.Cmp(big.NewFloat(0)) < 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					if mission.ToolSet != models.ToolsetTypeLuckyMoneys {
						inferPost, err = s.BatchPromptItemV2(ctx, agentInfo, inferPost)
						if err != nil {
							inferPost.Error = err.Error()
							inferPost.Status = models.AgentSnapshotPostStatusInferError
							err = s.dao.Create(
								tx,
								inferPost,
							)
							if err != nil {
								return errs.NewError(err)
							}
							err = tx.Model(agentInfo).
								UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", inferPost.Fee)).
								Error
							if err != nil {
								return errs.NewError(err)
							}
							inferPost.Fee = numeric.NewBigFloatFromString("0")
							err = s.dao.Save(
								tx,
								inferPost,
							)
							if err != nil {
								return errs.NewError(err)
							}
							return nil
						}
					}
					err = s.dao.Create(
						tx,
						inferPost,
					)
					if err != nil {
						return errs.NewError(err)
					}
					err = tx.Model(mission).
						UpdateColumn("infer_at", helpers.TimeNow()).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					if inferPost.Fee.Cmp(big.NewFloat(0)) > 0 {
						_ = s.dao.Create(
							tx,
							&models.AgentEaiTopup{
								NetworkID:      agentInfo.NetworkID,
								EventId:        fmt.Sprintf("agent_trigger_%d", inferPost.ID),
								AgentInfoID:    agentInfo.ID,
								Type:           models.AgentEaiTopupTypeSpent,
								Amount:         inferPost.Fee,
								Status:         models.AgentEaiTopupStatusDone,
								DepositAddress: agentInfo.ETHAddress,
								ToAddress:      agentInfo.ETHAddress,
								Toolset:        string(mission.ToolSet),
							},
						)
					}
					{
						if inferPost.AgentStoreMissionID > 0 {
							_ = tx.Model(&models.AgentStoreMission{}).
								Where("id = ?", inferPost.AgentStoreMissionID).
								Updates(
									map[string]interface{}{
										"num_used": gorm.Expr("num_used + ?", 1),
										"volume":   gorm.Expr("volume + ?", inferPost.AgentStoreMissionFee),
									},
								).
								Error
						}
						if inferPost.AgentStoreID > 0 {
							_ = tx.Model(&models.AgentStore{}).
								Where("id = ?", inferPost.AgentStoreID).
								Updates(
									map[string]interface{}{
										"num_usage": gorm.Expr("num_usage + ?", 1),
										"volume":    gorm.Expr("volume + ?", inferPost.AgentStoreMissionFee),
									},
								).
								Error
						}
					}
					return nil
				},
			)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentSnapshotPostCreateForUser(ctx context.Context, networkID uint64, userAddress string, userPrompt string, agentBaseModel string, agentStoreMissionID uint) (*models.AgentSnapshotPost, error) {
	var inferPost *models.AgentSnapshotPost
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentSnapshotPostCreateForUser_%s", userAddress),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					user, err := s.GetUser(
						tx,
						0,
						userAddress,
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					agentStoreMission, err := s.dao.FirstAgentStoreMissionByID(
						tx,
						agentStoreMissionID,
						map[string][]interface{}{},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					agentStoreMission.NumUsed++
					var headSystemPrompt string
					metaDataReq := &aidojo.AgentMetadataRequest{}
					inferTxHash := helpers.RandomBigInt(12).Text(16)
					inferItems := []*models.UserAgentInferDataItem{
						{
							Role:    "user",
							Content: strings.TrimSpace(agentStoreMission.UserPrompt),
						},
					}
					agentChainFee, err := s.GetAgentChainFee(
						tx,
						networkID,
					)
					if err != nil {
						return errs.NewError(err)
					}
					inferFee := agentChainFee.InferFee
					missionStoreFee := numeric.NewBigFloatFromFloat(big.NewFloat(0))
					missionStoreFee = agentStoreMission.Price
					inferFee = numeric.NewBigFloatFromFloat(models.AddBigFloats(&inferFee.Float, &missionStoreFee.Float))
					params := map[string]interface{}{}
					agentStoreInstall, err := s.dao.FirstAgentStoreInstall(
						tx,
						map[string][]interface{}{
							"type = ?":           {models.AgentStoreInstallTypeUser},
							"user_id = ?":        {user.ID},
							"agent_store_id = ?": {agentStoreMission.AgentStoreID},
							"status = ?":         {models.AgentStoreInstallStatusDone},
						},
						map[string][]interface{}{
							"User": {},
						},
						[]string{"id desc"},
					)
					if err != nil {
						return errs.NewError(err)
					}
					if agentStoreInstall != nil {
						err = helpers.ConvertJsonObject(agentStoreInstall.CallbackParams, &params)
						if err != nil {
							return errs.NewError(err)
						}
						params["install_code"] = agentStoreInstall.Code
						if agentStoreInstall.User != nil {
							params["user_address"] = agentStoreInstall.User.Address
						}
					}
					inferItems[0].Content, err = helpers.GenerateTemplateContent(agentStoreMission.UserPrompt, params)
					if err != nil {
						return errs.NewError(err)
					}
					toolList, err := helpers.GenerateTemplateContent(agentStoreMission.ToolList, params)
					if err != nil {
						return errs.NewError(err)
					}
					inferData, err := json.Marshal(inferItems)
					if err != nil {
						return errs.NewError(err)
					}
					prompt := inferItems[0].Content
					if userPrompt != "" {
						prompt = userPrompt
					}

					inferPost = &models.AgentSnapshotPost{
						NetworkID:        networkID,
						UserID:           user.ID,
						InferData:        string(inferData),
						InferAt:          helpers.TimeNow(),
						Status:           models.AgentSnapshotPostStatusInferSubmitted,
						Fee:              inferFee,
						UserPrompt:       prompt,
						HeadSystemPrompt: headSystemPrompt,
						AgentMetaData:    helpers.ConvertJsonString(metaDataReq),
						ToolList:         toolList,
						// SystemPrompt:         systemPrompt,
						SystemReminder:       "",
						Toolset:              "mission_store",
						AgentBaseModel:       agentBaseModel,
						ReactMaxSteps:        3,
						InferTxHash:          inferTxHash,
						AgentStoreMissionID:  agentStoreMission.ID,
						AgentStoreID:         agentStoreMission.AgentStoreID,
						IsRated:              false,
						AgentStoreMissionFee: missionStoreFee,
					}
					if inferPost.Fee.Float.Cmp(big.NewFloat(0)) <= 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					if user.EaiBalance.Float.Cmp(&inferPost.Fee.Float) < 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					err = tx.Model(user).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance - ?", inferPost.Fee)).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					user, err = s.dao.FirstUserByID(
						tx,
						user.ID,
						map[string][]interface{}{},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if user.EaiBalance.Float.Cmp(big.NewFloat(0)) < 0 {
						return errs.NewError(errs.ErrBadRequest)
					}
					inferPost, err = s.BatchPromptItemV2(ctx, nil, inferPost)
					if err != nil {
						inferPost.Error = err.Error()
						inferPost.Status = models.AgentSnapshotPostStatusInferError
						err = s.dao.Create(
							tx,
							inferPost,
						)
						if err != nil {
							return errs.NewError(err)
						}
						err = tx.Model(user).
							UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", inferPost.Fee)).
							Error
						if err != nil {
							return errs.NewError(err)
						}
						inferPost.Fee = numeric.NewBigFloatFromString("0")
						err = s.dao.Save(
							tx,
							inferPost,
						)
						if err != nil {
							return errs.NewError(err)
						}
						return nil
					}
					err = s.dao.Create(
						tx,
						inferPost,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if inferPost.Fee.Cmp(big.NewFloat(0)) > 0 {
						_ = s.dao.Create(
							tx,
							&models.UserTransaction{
								EventId: fmt.Sprintf("agent_trigger_%d", inferPost.ID),
								UserID:  user.ID,
								Type:    models.UserTransactionTypeTriggerFee,
								Amount:  numeric.NewBigFloatFromFloat(models.NegativeBigFloat(&inferPost.Fee.Float)),
								Status:  models.UserTransactionStatusDone,
							},
						)
					}
					{
						if inferPost.AgentStoreMissionID > 0 {
							_ = tx.Model(&models.AgentStoreMission{}).
								Where("id = ?", inferPost.AgentStoreMissionID).
								Updates(
									map[string]interface{}{
										"num_used": gorm.Expr("num_used + ?", 1),
										"volume":   gorm.Expr("volume + ?", inferPost.AgentStoreMissionFee),
									},
								).
								Error
						}
						if inferPost.AgentStoreID > 0 {
							_ = tx.Model(&models.AgentStore{}).
								Where("id = ?", inferPost.AgentStoreID).
								Updates(
									map[string]interface{}{
										"num_usage": gorm.Expr("num_usage + ?", 1),
										"volume":    gorm.Expr("volume + ?", inferPost.AgentStoreMissionFee),
									},
								).
								Error
						}
					}
					// add history
					agentStoreTry, _ := s.dao.FirstAgentStoreTry(
						tx,
						map[string][]interface{}{
							"user_id = ?":        {user.ID},
							"agent_store_id = ?": {agentStoreMission.AgentStoreID},
						},
						map[string][]interface{}{},
						false,
					)
					if agentStoreTry == nil {
						agentStoreTry = &models.AgentStoreTry{
							UserID:       user.ID,
							AgentStoreID: agentStoreMission.AgentStoreID,
						}
						s.dao.Create(
							tx,
							agentStoreTry,
						)
					}
					if agentStoreTry != nil && agentStoreTry.ID > 0 {
						history := &models.AgentStoreTryDetail{
							AgentStoreTryID: agentStoreTry.ID,
							FromUser:        true,
							Content:         userPrompt,
						}
						s.dao.Create(
							tx,
							history,
						)

						history = &models.AgentStoreTryDetail{
							AgentStoreTryID:     agentStoreTry.ID,
							FromUser:            false,
							Content:             "",
							AgentSnapshotPostID: inferPost.ID,
						}
						s.dao.Create(
							tx,
							history,
						)
					}
					return nil
				},
			)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return inferPost, nil
}

func (s *Service) AgentSnapshotPostActionExecuted(ctx context.Context, twitterPostID uint) error {
	var snapshotPostAction *models.AgentSnapshotPostAction
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentSnapshotPostActionExecuted_%d", twitterPostID),
		func() error {
			contentLines := []string{}
			var accessToken, missionToolSet string
			postIds := []string{}
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					var err error
					snapshotPostAction, err = s.dao.FirstAgentSnapshotPostActionByID(
						tx,
						twitterPostID,
						map[string][]interface{}{},
						true,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if snapshotPostAction.Status == models.AgentSnapshotPostActionStatusNew {
						if snapshotPostAction.ScheduleAt.Before(time.Now().Add(-6 * time.Hour)) {
							snapshotPostAction.Status = models.AgentSnapshotPostActionStatusDoneCancelled
							snapshotPostAction.ExecutedAt = helpers.TimeNow()
							err = s.dao.Save(
								tx,
								snapshotPostAction,
							)
							if err != nil {
								return errs.NewError(err)
							}
							return nil
						}
						agent, err := s.dao.FirstAgentInfoByID(
							tx,
							snapshotPostAction.AgentInfoID,
							map[string][]interface{}{
								"TwitterInfo": {},
							},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						agentSnapshotMission, err := s.dao.FirstAgentSnapshotMissionByID(
							tx,
							snapshotPostAction.AgentSnapshotMissionID,
							map[string][]interface{}{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						var isPassed bool
						if agentSnapshotMission.NotDelay {
							isPassed = true
						} else {
							switch snapshotPostAction.ToolSet {
							case models.ToolsetTypeReplyMentions,
								models.ToolsetTypeFollow,
								models.ToolsetTypeInscribeTweet,
								models.ToolsetTypePost,
								models.ToolsetTypeTrading,
								models.ToolsetTypeTradeAnalytics,
								models.ToolsetTypeTradeAnalyticsOnTwitter,
								models.ToolsetTypeTradeAnalyticsMentions,
								models.ToolsetTypeLuckyMoneys,
								models.ToolsetType("post_search"):
								{
									isPassed = true
								}
							default:
								{
									if agent.ReplyLatestTime == nil ||
										agent.ReplyLatestTime.Before(time.Now().Add(-time.Duration(agent.ActionDelayed)*time.Second)) {
										isPassed = true
									}
								}
							}
						}
						missionToolSet = string(agentSnapshotMission.ToolSet)
						switch snapshotPostAction.Type {
						case models.AgentSnapshotPostActionTypeReplyMultiUnlimited:
							{
								isPassed = true
							}
						}
						if isPassed &&
							agent.ReplyEnabled {
							//
							var refId, inscribeId, errText string
							var tokenAddress, tokenHash string
							isDupplicated := false
							switch snapshotPostAction.Type {
							case models.AgentSnapshotPostActionTypeFollow:
								{
									if snapshotPostAction.TargetTwitterId == "" {
										targetUsername := strings.TrimPrefix(snapshotPostAction.TargetUsername, "@")
										if targetUsername == "" {
											errText = "targetUsername not found"
										} else {
											var targetTwitterId string
											var followerCount uint
											for i := 0; i < 3; i++ {
												obj, _ := s.twitterAPI.GetTwitterByUserName(targetUsername)
												if obj != nil {
													followerCount = uint(obj.PublicMetrics.Followers)
													targetTwitterId = obj.ID
													break
												}
											}
											if targetTwitterId == "" {
												twitterUser, _ := s.dao.FirstTwitterUser(
													tx,
													map[string][]interface{}{
														"twitter_username = ?": {targetUsername},
													},
													map[string][]interface{}{},
													[]string{},
												)
												if twitterUser != nil {
													followerCount = twitterUser.FollowersCount
													targetTwitterId = twitterUser.TwitterID
												}
											}
											if targetTwitterId == "" {
												errText = "targetTwitterId not found"
											} else {
												snapshotPostAction.TargetTwitterId = targetTwitterId
												snapshotPostAction.FollowerCount = followerCount
											}
										}
									}
									if snapshotPostAction.AgentTwitterId == snapshotPostAction.TargetTwitterId {
										isDupplicated = true
									} else {
										m, err := s.dao.FirstTwitterFollowing(
											tx,
											map[string][]interface{}{
												"owner_twitter_id = ?": {snapshotPostAction.AgentTwitterId},
												"twitter_id = ?":       {snapshotPostAction.TargetTwitterId},
											},
											map[string][]interface{}{},
											[]string{},
										)
										if err != nil {
											return errs.NewError(err)
										}
										if m != nil {
											isDupplicated = true
										}
									}
								}
							case models.AgentSnapshotPostActionTypeReply,
								models.AgentSnapshotPostActionTypeRetweet,
								models.AgentSnapshotPostActionTypeQuoteTweet:
								{
									m, err := s.dao.FirstAgentSnapshotPostAction(
										tx,
										map[string][]interface{}{
											"agent_info_id = ?":   {snapshotPostAction.AgentInfoID},
											"type = ?":            {snapshotPostAction.Type},
											"conversation_id = ?": {snapshotPostAction.ConversationId},
											"status = ?":          {models.AgentSnapshotPostActionStatusDone},
										},
										map[string][]interface{}{},
										[]string{},
									)
									if err != nil {
										return errs.NewError(err)
									}
									if m != nil {
										isDupplicated = true
									}
								}
							case models.AgentSnapshotPostActionTypeCreateToken:
								{
									m, err := s.dao.FirstAgentSnapshotPostAction(
										tx,
										map[string][]interface{}{
											"agent_info_id = ?": {snapshotPostAction.AgentInfoID},
											"type = ?":          {snapshotPostAction.Type},
											"token_symbol = ?":  {snapshotPostAction.TokenSymbol},
											"status = ?":        {models.AgentSnapshotPostActionStatusDone},
										},
										map[string][]interface{}{},
										[]string{},
									)
									if err != nil {
										return errs.NewError(err)
									}
									if m != nil {
										isDupplicated = true
									}
								}
							}
							if !isDupplicated {
								err = tx.Model(agent).
									UpdateColumn("reply_latest_time", time.Now()).
									Error
								if err != nil {
									return errs.NewError(err)
								}
								if errText == "" {
									accessToken = agent.TwitterInfo.AccessToken
									switch snapshotPostAction.Type {
									case models.AgentSnapshotPostActionTypeFollow:
										{
											err = helpers.TwitterFollowUserCreate(agent.TwitterInfo.AccessToken, agent.TwitterID, snapshotPostAction.TargetTwitterId)
											if err != nil {
												errText = err.Error()
											}
										}
									case models.AgentSnapshotPostActionTypeTweet:
										{
											maxChars, err := s.GetTwitterPostMaxChars(tx, snapshotPostAction.AgentInfoID)
											if err != nil {
												return errs.NewError(err)
											}
											if snapshotPostAction.ToolSet == models.ToolsetTypeTrading {
												maxChars = 250
											}
											contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(snapshotPostAction.Content, int(maxChars))
											if len(contentLines) <= 0 {
												return errs.NewError(errs.ErrBadRequest)
											}
											refId, err = helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, contentLines[0], "")
											if err != nil {
												errText = err.Error()
											}
											postIds = append(postIds, refId)
										}
									case models.AgentSnapshotPostActionTypeReply,
										models.AgentSnapshotPostActionTypeReplyMulti,
										models.AgentSnapshotPostActionTypeReplyMultiUnlimited:
										{
											if snapshotPostAction.ToolSet == models.ToolsetTypeTradeAnalyticsMentions {
												err = json.Unmarshal([]byte(snapshotPostAction.Content), &contentLines)
												if err != nil {
													return errs.NewError(err)
												}
												contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(strings.Join(contentLines, ". "), 250)
												if len(contentLines) <= 0 {
													return errs.NewError(errs.ErrBadRequest)
												}
											} else {
												maxChars, err := s.GetTwitterPostMaxChars(tx, snapshotPostAction.AgentInfoID)
												if err != nil {
													return errs.NewError(err)
												}
												contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(snapshotPostAction.Content, int(maxChars))
												if len(contentLines) <= 0 {
													return errs.NewError(errs.ErrBadRequest)
												}
											}

											mediaID := ""
											if snapshotPostAction.TokenImageUrl != "" {
												mediaID, _ = s.twitterAPI.UploadImage(models.GetImageUrl(snapshotPostAction.TokenImageUrl), []string{snapshotPostAction.AgentTwitterId})
												// if err != nil {
												// 	return errs.NewError(err)
												// }
											}

											refId, err = helpers.ReplyTweetByToken(agent.TwitterInfo.AccessToken, strings.TrimPrefix(contentLines[0], "."), snapshotPostAction.Tweetid, mediaID)
											if err != nil {
												errText = err.Error()
											}
											postIds = append(postIds, refId)
											cacheKey := fmt.Sprintf("CacheCheckIsTweetReplied_%s", snapshotPostAction.Tweetid)
											_ = s.SetRedisCachedWithKey(cacheKey,
												true,
												24*time.Hour,
											)
										}
									case models.AgentSnapshotPostActionTypeRetweet:
										{
											refId, err = helpers.RepostTweetByToken(agent.TwitterInfo.AccessToken, agent.TwitterID, snapshotPostAction.Tweetid)
											if err != nil {
												errText = err.Error()
											}
										}
									case models.AgentSnapshotPostActionTypeQuoteTweet:
										{
											maxChars, err := s.GetTwitterPostMaxChars(tx, snapshotPostAction.AgentInfoID)
											if err != nil {
												return errs.NewError(err)
											}
											contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(snapshotPostAction.Content, int(maxChars))
											if len(contentLines) <= 0 {
												return errs.NewError(errs.ErrBadRequest)
											}
											refId, err = helpers.QuoteTweetByToken(agent.TwitterInfo.AccessToken, strings.TrimPrefix(contentLines[0], "."), snapshotPostAction.Tweetid)
											if err != nil {
												errText = err.Error()
											}
											postIds = append(postIds, refId)
										}
									case models.AgentSnapshotPostActionTypeInscribeTweet:
										{
											inscribeId, err = s.dojoAPI.AgentInscribe(
												&aidojo.AgentInscribeReq{
													Content:     snapshotPostAction.Content,
													MysqlID:     snapshotPostAction.ID,
													TweetID:     "",
													AgentID:     snapshotPostAction.AgentInfoID,
													Type:        "tweet",
													CreatedAt:   *helpers.TimeNow(),
													PostTweetAt: *s.GetPostTimeByTweetID(tx, snapshotPostAction.Tweetid),
												},
											)
											if err != nil {
												errText = err.Error()
											}
										}
									case models.AgentSnapshotPostActionTypeTweetV2:
										{
											maxChars, err := s.GetTwitterPostMaxChars(tx, snapshotPostAction.AgentInfoID)
											if err != nil {
												return errs.NewError(err)
											}
											contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(snapshotPostAction.Content, int(maxChars))
											if len(contentLines) <= 0 {
												return errs.NewError(errs.ErrBadRequest)
											}
											mediaID, err := s.twitterAPI.UploadImage(models.GetImageUrl(snapshotPostAction.TokenImageUrl), []string{snapshotPostAction.AgentTwitterId})
											if err != nil {
												return errs.NewError(err)
											}
											refId, err = helpers.PostTweetWithMediaByToken(agent.TwitterInfo.AccessToken, contentLines[0], mediaID)
											if err != nil {
												errText = err.Error()
											}
											postIds = append(postIds, refId)
										}
									case models.AgentSnapshotPostActionTypeCreateToken:
										{
											if snapshotPostAction.TokenImageUrl == "" {
												imageUrl, err := s.GetGifImageUrlFromTokenInfo(snapshotPostAction.TokenSymbol, snapshotPostAction.TokenName, snapshotPostAction.Description)
												if err != nil {
													return errs.NewError(err)
												}
												if imageUrl != "" {
													snapshotPostAction.TokenImageUrl = imageUrl
													err = s.dao.Save(
														tx,
														snapshotPostAction,
													)
													if err != nil {
														return errs.NewError(err)
													}
													return nil
												}
											} else {
												agentTokenAdminAddress := s.conf.GetConfigKeyString(models.GENERTAL_NETWORK_ID, "agent_token_admin_address")
												base64Str, _ := helpers.CurlBase64String(models.GetImageUrl(snapshotPostAction.TokenImageUrl))
												if base64Str != "" {
													pumfunResp, err := s.blockchainUtils.SolanaCreatePumpfunToken(
														&blockchainutils.SolanaCreatePumpfunTokenReq{
															Address:     agentTokenAdminAddress,
															Name:        snapshotPostAction.TokenName,
															Symbol:      snapshotPostAction.TokenSymbol,
															Description: snapshotPostAction.Content,
															Twitter:     fmt.Sprintf("https://x.com/%s", agent.TwitterUsername),
															Telegram:    "",
															Website:     "",
															Amount:      0,
															ImageBase64: base64Str,
														},
													)
													if err != nil {
														errText = err.Error()
													} else {
														tokenAddress = pumfunResp.Mint
														tokenHash = pumfunResp.Signature
														content := fmt.Sprintf(
															"%s\n\n%s",
															snapshotPostAction.Content,
															fmt.Sprintf("https://pump.fun/coin/%s", tokenAddress),
														)
														contentLines = []string{
															content,
														}
														refId, err = helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, contentLines[0], "")
														if err != nil {
															errText = err.Error()
														}
														postIds = append(postIds, refId)
													}
												} else {
													errText = "base64Str empty data"
												}
											}
										}
									case models.AgentSnapshotPostActionTypeTweetMulti:
										{
											err = json.Unmarshal([]byte(snapshotPostAction.Content), &contentLines)
											if err != nil {
												return errs.NewError(err)
											}
											if snapshotPostAction.ToolSet == models.ToolsetTypeTrading || snapshotPostAction.ToolSet == models.ToolsetTypeTradeAnalyticsOnTwitter {
												contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(strings.Join(contentLines, ". "), 250)
												if len(contentLines) <= 0 {
													return errs.NewError(errs.ErrBadRequest)
												}
												refId, err = helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, contentLines[0], "")
												if err != nil {
													errText = err.Error()
												}
												postIds = append(postIds, refId)
											} else {
												refId, err = helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, contentLines[0], "")
												if err != nil {
													errText = err.Error()
												}
												postIds = append(postIds, refId)
											}
										}
									case models.AgentSnapshotPostActionTypeTradeHold, models.AgentSnapshotPostActionTypeTradeBuy, models.AgentSnapshotPostActionTypeTradeSell, models.AgentSnapshotPostActionTypeTradeAnalytic:
										{
											switch agentSnapshotMission.ToolSet {
											case models.ToolsetTypeTradeAnalytics:
												{
													if agentSnapshotMission.TeleChatID != "" {
														teleMsg :=
															fmt.Sprintf(`
		%s
		
		Action: %s
		`, snapshotPostAction.Content, snapshotPostAction.Type)

														//TODO: send telegram
														refId, err = s.SendTeleMsgToChatID(ctx, teleMsg, agentSnapshotMission.TeleChatID)
														if err != nil {
															errText = err.Error()
														}
													}
												}
											case models.ToolsetTypeTradeAnalyticsOnTwitter:
												{
													err = json.Unmarshal([]byte(snapshotPostAction.Content), &contentLines)
													if err != nil {
														return errs.NewError(err)
													}
													contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(strings.Join(contentLines, ". "), 250)
													if len(contentLines) <= 0 {
														return errs.NewError(errs.ErrBadRequest)
													}
													refId, err = helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, contentLines[0], "")
													if err != nil {
														errText = err.Error()
													}
													postIds = append(postIds, refId)
												}
											case models.ToolsetTypeLaunchpadJoin:
												{
													//TODO
													// err = json.Unmarshal([]byte(snapshotPostAction.Content), &contentLines)
													// if err != nil {
													// 	return errs.NewError(err)
													// }
													// contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(strings.Join(contentLines, ". "), 250)
													// if len(contentLines) <= 0 {
													// 	return errs.NewError(errs.ErrBadRequest)
													// }
													// refId, err = helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, contentLines[0], "")
													// if err != nil {
													// 	errText = err.Error()
													// }
													// postIds = append(postIds, refId)
												}
											default:
												{
													errText = "not supported"
												}
											}
										}
									default:
										{
											errText = "not supported"
										}
									}
								}
								snapshotPostAction.RefId = refId
								snapshotPostAction.RefIds = refId
								snapshotPostAction.InscribeId = inscribeId
								snapshotPostAction.Error = errText
								snapshotPostAction.TokenAddress = tokenAddress
								snapshotPostAction.TokenHash = tokenHash
								if errText == "" {
									snapshotPostAction.Status = models.AgentSnapshotPostActionStatusDone
								} else {
									// if strings.Contains(errText, "suspended") || strings.Contains(errText, "locked") {
									// 	err = tx.Model(agent).
									// 		UpdateColumn("reply_enabled", false).
									// 		Error
									// 	if err != nil {
									// 		return errs.NewError(err)
									// 	}
									// }
									snapshotPostAction.Status = models.AgentSnapshotPostActionStatusDoneError
								}
							} else {
								snapshotPostAction.Status = models.AgentSnapshotPostActionStatusDoneDuplicated
							}
							snapshotPostAction.ExecutedAt = helpers.TimeNow()
							err = s.dao.Save(
								tx,
								snapshotPostAction,
							)
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
			if len(postIds) > 0 && len(contentLines) > 1 {
				for i := 1; i < len(contentLines); i++ {
					var postId string
					switch snapshotPostAction.Type {
					case models.AgentSnapshotPostActionTypeTweet,
						models.AgentSnapshotPostActionTypeQuoteTweet,
						models.AgentSnapshotPostActionTypeCreateToken,
						models.AgentSnapshotPostActionTypeTweetV2:
						{
							postId, err = helpers.ReplyTweetByToken(accessToken, contentLines[i], postIds[0], "")
							if err != nil {
								return errs.NewError(err)
							}
						}
					case models.AgentSnapshotPostActionTypeReply,
						models.AgentSnapshotPostActionTypeTweetMulti,
						models.AgentSnapshotPostActionTypeReplyMulti:
						{
							postId, _ = helpers.ReplyTweetByToken(accessToken, contentLines[i], postIds[len(postIds)-1], "")
						}
					case models.AgentSnapshotPostActionTypeTradeHold, models.AgentSnapshotPostActionTypeTradeBuy, models.AgentSnapshotPostActionTypeTradeSell, models.AgentSnapshotPostActionTypeTradeAnalytic:
						{
							switch missionToolSet {
							case string(models.ToolsetTypeTradeAnalyticsOnTwitter):
								{
									postId, _ = helpers.ReplyTweetByToken(accessToken, contentLines[i], postIds[len(postIds)-1], "")
								}
							}
						}
					}
					if postId != "" {
						postIds = append(postIds, postId)
						err = daos.
							GetDBMainCtx(ctx).
							Model(&models.AgentSnapshotPostAction{}).
							Where("id = ?", twitterPostID).
							UpdateColumn("ref_ids", strings.Join(postIds, ",")).
							Error
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
	if snapshotPostAction != nil {
		_ = s.UpdateOffchainAutoOutputV2(ctx, snapshotPostAction.AgentSnapshotPostID)
	}
	return nil
}

//	func (s *Service) RetryAgentSnapshotPostActionExecuted(ctx context.Context, twitterPostID uint) error {
//		err := s.JobRunCheck(
//			ctx,
//			fmt.Sprintf("RetryAgentSnapshotPostActionExecuted_%d", twitterPostID),
//			func() error {
//				var snapshotPostAction *models.AgentSnapshotPostAction
//				contentLines := []string{}
//				var accessToken, missionToolSet string
//				postIds := []string{}
//				err := daos.WithTransaction(
//					daos.GetDBMainCtx(ctx),
//					func(tx *gorm.DB) error {
//						var err error
//						snapshotPostAction, err = s.dao.FirstAgentSnapshotPostActionByID(
//							tx,
//							twitterPostID,
//							map[string][]interface{}{},
//							true,
//						)
//						if err != nil {
//							return errs.NewError(err)
//						}
//						if snapshotPostAction.Status == models.AgentSnapshotPostActionStatusDone {
//							agent, err := s.dao.FirstAgentInfoByID(
//								tx,
//								snapshotPostAction.AgentInfoID,
//								map[string][]interface{}{
//									"TwitterInfo": {},
//								},
//								false,
//							)
//							if err != nil {
//								return errs.NewError(err)
//							}
//							//
//							var refId, errText string
//							agentSnapshotMission, err := s.dao.FirstAgentSnapshotMissionByID(
//								tx,
//								snapshotPostAction.AgentSnapshotMissionID,
//								map[string][]interface{}{},
//								false,
//							)
//							if err != nil {
//								return errs.NewError(err)
//							}
//							missionToolSet = string(agentSnapshotMission.ToolSet)
//							accessToken = agent.TwitterInfo.AccessToken
//							switch snapshotPostAction.Type {
//							case models.AgentSnapshotPostActionTypeTradeHold, models.AgentSnapshotPostActionTypeTradeBuy, models.AgentSnapshotPostActionTypeTradeSell, models.AgentSnapshotPostActionTypeTradeAnalytic:
//								{
//									switch missionToolSet {
//									case string(models.ToolsetTypeTradeAnalyticsOnTwitter):
//										{
//											err = json.Unmarshal([]byte(snapshotPostAction.Content), &contentLines)
//											if err != nil {
//												return errs.NewError(err)
//											}
//											contentLines = helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(strings.Join(contentLines, ". "), 250)
//											if len(contentLines) <= 0 {
//												return errs.NewError(errs.ErrBadRequest)
//											}
//											refId = snapshotPostAction.RefId
//											postIds = append(postIds, refId)
//										}
//									default:
//										{
//											return errs.NewError(errs.ErrBadRequest)
//										}
//									}
//								}
//							default:
//								{
//									return errs.NewError(errs.ErrBadRequest)
//								}
//							}
//							_ = errText
//						}
//						return nil
//					},
//				)
//				if err != nil {
//					return errs.NewError(err)
//				}
//				if len(postIds) > 0 && len(contentLines) > 1 {
//					for i := 1; i < len(contentLines); i++ {
//						var postId string
//						switch snapshotPostAction.Type {
//						case models.AgentSnapshotPostActionTypeTweet,
//							models.AgentSnapshotPostActionTypeQuoteTweet,
//							models.AgentSnapshotPostActionTypeCreateToken,
//							models.AgentSnapshotPostActionTypeTweetV2:
//							{
//								postId, err = helpers.ReplyTweetByToken(accessToken, contentLines[i], postIds[0])
//								if err != nil {
//									return errs.NewError(err)
//								}
//							}
//						case models.AgentSnapshotPostActionTypeReply,
//							models.AgentSnapshotPostActionTypeTweetMulti:
//							{
//								postId, _ = helpers.ReplyTweetByToken(accessToken, contentLines[i], postIds[len(postIds)-1])
//							}
//						case models.AgentSnapshotPostActionTypeTradeHold, models.AgentSnapshotPostActionTypeTradeBuy, models.AgentSnapshotPostActionTypeTradeSell, models.AgentSnapshotPostActionTypeTradeAnalytic:
//							{
//								switch missionToolSet {
//								case string(models.ToolsetTypeTradeAnalyticsOnTwitter):
//									{
//										postId, _ = helpers.ReplyTweetByToken(accessToken, contentLines[i], postIds[len(postIds)-1])
//									}
//								}
//							}
//						}
//						if postId != "" {
//							postIds = append(postIds, postId)
//							err = daos.
//								GetDBMainCtx(ctx).
//								Model(&models.AgentSnapshotPostAction{}).
//								Where("id = ?", twitterPostID).
//								UpdateColumn("ref_ids", strings.Join(postIds, ",")).
//								Error
//							if err != nil {
//								return errs.NewError(err)
//							}
//						}
//					}
//				}
//				return nil
//			},
//		)
//		if err != nil {
//			return errs.NewError(err)
//		}
//		_ = s.UpdateOffchainAutoOutputV2(ctx, twitterPostID)
//		return nil
//	}
func (s *Service) UpdateOffchainAutoOutputV2ForId(ctx context.Context, snapshotPostID uint) error {
	agentSnapshotPost, err := s.dao.FirstAgentSnapshotPostByID(
		daos.GetDBMainCtx(ctx),
		snapshotPostID,
		map[string][]interface{}{
			"AgentStore": {},
		},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agentSnapshotPost != nil {
		if agentSnapshotPost.ResponseId != "" {
			if agentSnapshotPost.Status == models.AgentSnapshotPostStatusInferSubmitted {
				if agentSnapshotPost.CreatedAt.Before(time.Now().Add(-30 * time.Hour)) {
					err = daos.GetDBMainCtx(ctx).
						Model(agentSnapshotPost).
						UpdateColumn("status", models.AgentSnapshotPostStatusInferFailed).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					return nil
				}
				offchainAutoAgentOutput, err := s.dojoAPI.OffchainAutoAgentOutput(s.conf.AgentOffchain.Url, agentSnapshotPost.ResponseId, s.conf.AgentOffchain.ApiKey)
				if err != nil {
					return errs.NewError(err)
				}
				aiOutput := map[string]interface{}{}
				err = helpers.ConvertJsonObject(offchainAutoAgentOutput, &aiOutput)
				if err != nil {
					return errs.NewError(err)
				}
				inferOutputData := helpers.ConvertJsonString(
					struct {
						Data struct {
							ResponseId string                 `json:"response_id"`
							Toolset    string                 `json:"toolset"`
							Task       string                 `json:"task"`
							AIOutput   map[string]interface{} `json:"ai_output"`
						} `json:"data"`
					}{
						Data: struct {
							ResponseId string                 "json:\"response_id\""
							Toolset    string                 "json:\"toolset\""
							Task       string                 "json:\"task\""
							AIOutput   map[string]interface{} "json:\"ai_output\""
						}{
							ResponseId: agentSnapshotPost.ResponseId,
							Toolset:    agentSnapshotPost.Toolset,
							Task:       agentSnapshotPost.Task,
							AIOutput:   aiOutput,
						},
					},
				)
				if len(inferOutputData) > len(agentSnapshotPost.InferOutputData) || !strings.EqualFold(inferOutputData, agentSnapshotPost.InferOutputData) {
					err = daos.GetDBMainCtx(ctx).
						Model(agentSnapshotPost).
						UpdateColumn("infer_output_data", inferOutputData).
						Error
					if err != nil {
						return errs.NewError(err)
					}
				}

				//cache result
				agentSnapshotPost.InferOutputData = inferOutputData
				err = s.CacheAgentSnapshotPost(agentSnapshotPost)
				if err != nil {
					return errs.NewError(err)
				}

				state, ok := aiOutput["state"]
				if ok {
					if state.(string) == "done" {
						err = daos.WithTransaction(
							daos.GetDBMainCtx(ctx),
							func(tx *gorm.DB) error {
								err = tx.
									Model(agentSnapshotPost).
									UpdateColumn("status", models.AgentSnapshotPostStatusInferResolved).
									Error
								if err != nil {
									return errs.NewError(err)
								}
								if agentSnapshotPost.AgentStore != nil &&
									agentSnapshotPost.AgentStoreMissionFee.Float.Cmp(big.NewFloat(0)) > 0 {
									user, err := s.GetUser(tx, models.GENERTAL_NETWORK_ID, agentSnapshotPost.AgentStore.OwnerAddress, false)
									if err != nil {
										return errs.NewError(err)
									}
									err = tx.Model(user).
										UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", agentSnapshotPost.AgentStoreMissionFee)).
										Error
									if err != nil {
										return errs.NewError(err)
									}
									_ = s.dao.Create(
										tx,
										&models.UserTransaction{
											NetworkID: agentSnapshotPost.NetworkID,
											EventId:   fmt.Sprintf("agent_store_fee_%d", agentSnapshotPost.ID),
											UserID:    user.ID,
											Type:      models.UserTransactionTypeAgentStoreFee,
											Amount:    agentSnapshotPost.AgentStoreMissionFee,
											Status:    models.UserTransactionStatusDone,
										},
									)
								}
								return nil
							},
						)
						if err != nil {
							return errs.NewError(err)
						}
						go s.UpdateDataMissionTradeAnalytics(context.Background(), agentSnapshotPost.ID)
					} else if state.(string) == "error" {
						err = daos.GetDBMainCtx(ctx).
							Model(agentSnapshotPost).
							UpdateColumn("status", models.AgentSnapshotPostStatusInferFailed).
							Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				}
			}
		} else {
			if agentSnapshotPost.Status == models.AgentSnapshotPostStatusInferSubmitted {
				inferOutputData, err := s.dojoAPI.OffchainAgentOutput(agentSnapshotPost.InferTxHash)
				if err != nil {
					return errs.NewError(err)
				}
				if len(inferOutputData) > len(agentSnapshotPost.InferOutputData) || !strings.EqualFold(inferOutputData, agentSnapshotPost.InferOutputData) {
					err = daos.GetDBMainCtx(ctx).
						Model(agentSnapshotPost).
						UpdateColumn("infer_output_data", inferOutputData).
						Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) UpdateOffchainAutoOutputV2(ctx context.Context, snapshotPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("UpdateOffchainAutoOutputV2x_%d", snapshotPostID),
		func() error {
			var rs bool
			err := s.RedisCached(
				fmt.Sprintf("UpdateOffchainAutoOutputV2x_%d", snapshotPostID),
				true,
				15*time.Minute,
				&rs,
				func() (interface{}, error) {
					err := func() error {
						err := s.UpdateOffchainAutoOutputV2ForId(ctx, snapshotPostID)
						if err != nil {
							return errs.NewError(err)
						}
						return nil
					}()
					if err != nil {
						return false, errs.NewError(err)
					}
					return true, nil
				},
			)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateDataMissionTradeAnalytics(ctx context.Context, snapshotPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("UpdateDataMissionTradeAnalytics_%d", snapshotPostID),
		func() error {
			err := func() error {
				snapshotPost, err := s.dao.FirstAgentSnapshotPostByID(
					daos.GetDBMainCtx(ctx),
					snapshotPostID,
					map[string][]interface{}{
						"AgentSnapshotMission": {},
						"AgentInfo":            {},
					},
					false,
				)
				if err != nil {
					return errs.NewError(err)
				}
				if snapshotPost != nil && snapshotPost.AgentInfo != nil &&
					snapshotPost.ResponseId != "" &&
					snapshotPost.AgentSnapshotMission != nil &&
					(snapshotPost.AgentSnapshotMission.ToolSet == models.ToolsetTypeTradeAnalyticsOnTwitter ||
						snapshotPost.AgentSnapshotMission.ToolSet == models.ToolsetTypeTradeAnalyticsMentions ||
						snapshotPost.AgentSnapshotMission.ToolSet == models.ToolsetTypeLuckyMoneys ||
						snapshotPost.AgentSnapshotMission.ToolSet == models.ToolsetTypeMissionStore) {
					if snapshotPost.AgentSnapshotMission.ToolSet == models.ToolsetTypeMissionStore || snapshotPost.AgentStoreMissionID > 0 {
						//cache result
						err = s.CacheAgentSnapshotPost(snapshotPost)
						if err != nil {
							return errs.NewError(err)
						}
					} else if snapshotPost.Status == models.AgentSnapshotPostStatusInferResolved {
						var rs struct {
							Data struct {
								ResponseId string `json:"response_id"`
								Toolset    string `json:"toolset"`
								Task       string `json:"task"`
								AIOutput   struct {
									Scratchpad []struct {
										Task        string `json:"task"`
										Action      string `json:"action"`
										ActionInput string `json:"action_input"`
										Thought     string `json:"thought"`
										FinalAnswer string `json:"final_answer"`
									} `json:"scratchpad"`
								} `json:"ai_output"`
							} `json:"data"`
						}

						err := helpers.ConvertJsonObject(snapshotPost.InferOutputData, &rs)
						if err != nil {
							return errs.NewError(err)
						}

						status := models.AgentSnapshotPostActionStatusNew
						if snapshotPost.AgentSnapshotMission != nil {
							if snapshotPost.AgentSnapshotMission.IsTesting {
								status = models.AgentSnapshotPostActionStatusTesting
							}
						}

						listThough := []string{}
						for i, item := range rs.Data.AIOutput.Scratchpad {
							if i == 0 || i == 1 {
								continue
							}
							if item.Thought != "" {
								listThough = append(listThough, item.Thought)
							}

							if item.FinalAnswer != "" {
								listThough = append(listThough, item.FinalAnswer)
							}
						}

						j, err := json.Marshal(listThough)

						imageUrl := ""
						if snapshotPost.Token != "" {
							imageUrl, _ = s.GetChartImage(ctx, snapshotPost.Token)
						}
						action := &models.AgentSnapshotPostAction{
							NetworkID:              snapshotPost.AgentInfo.NetworkID,
							AgentInfoID:            snapshotPost.AgentInfo.ID,
							AgentSnapshotPostID:    snapshotPostID,
							AgentSnapshotMissionID: snapshotPost.AgentSnapshotMissionID,
							AgentTwitterId:         snapshotPost.AgentInfo.TwitterID,
							Content:                string(j),
							Status:                 status,
							ScheduleAt:             helpers.TimeNow(),
							ReqRefID:               snapshotPost.InferTxHash,
							ToolSet:                snapshotPost.AgentSnapshotMission.ToolSet,
							Tweetid:                snapshotPost.OrgTweetID,
							TokenImageUrl:          imageUrl,
						}
						if snapshotPost.OrgTweetID != "" {
							action.ConversationId = s.GetConversationIdByTweetID(daos.GetDBMainCtx(ctx), snapshotPost.OrgTweetID)
						}

						switch snapshotPost.AgentSnapshotMission.ToolSet {
						case models.ToolsetTypeTradeAnalyticsOnTwitter:
							{
								action.Type = models.AgentSnapshotPostActionTypeTweetMulti
							}
						case models.ToolsetTypeTradeAnalyticsMentions:
							{
								if snapshotPost.OrgTweetID == "" {
									action.Status = models.AgentSnapshotPostActionStatusInvalid
								}
								action.Type = models.AgentSnapshotPostActionTypeReplyMulti
							}
						case models.ToolsetTypeLuckyMoneys:
							{
								action.Type = models.AgentSnapshotPostActionTypeTweet
							}
						case models.ToolsetTypeLaunchpadJoin:
							{
								action.Type = models.AgentSnapshotPostActionTypeLaunchpadJoin
							}
						}

						err = s.dao.Create(
							daos.GetDBMainCtx(ctx), action,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if action.Type == models.AgentSnapshotPostActionTypeLaunchpadJoin {
							err = s.AgentSnapshotPostActionExecuted(ctx, action.ID)
							if err != nil {
								return errs.NewError(err)
							}
						}
					}
				}
				return nil
			}()
			return err
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) JobUpdateOffchainAutoOutputForMission(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateOffchainAutoOutputForMission",
		func() error {
			var retErr error
			{
				err := s.UpdateOffchainAutoOutputForMission(ctx)
				if err != nil {
					retErr = errs.MergeError(retErr, err)
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

func (s *Service) JobUpdateOffchainAutoOutput(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateOffchainAutoOutput",
		func() error {
			var retErr error
			{
				ms, err := s.dao.FindAgentSnapshotPostAction(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at <= adddate(now(), interval -5 minute)":  {},
						"created_at >= adddate(now(), interval -30 minute)": {},
						"agent_snapshot_post_id > 0":                        {},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.UpdateOffchainAutoOutputV2(ctx, m.AgentSnapshotPostID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.AgentSnapshotPostID))
					}
				}
			}
			{
				ms, err := s.dao.FindAgentWalletAction(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at <= adddate(now(), interval -5 minute)":  {},
						"created_at >= adddate(now(), interval -30 minute)": {},
						"agent_snapshot_post_id > 0":                        {},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.UpdateOffchainAutoOutputV2(ctx, m.AgentSnapshotPostID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
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

func (s *Service) UpdateOffchainAutoOutputForMission(ctx context.Context) error {
	var retErr error
	joinFilters := map[string][]interface{}{
		`
				left join agent_snapshot_missions on agent_snapshot_missions.id = agent_snapshot_posts.agent_snapshot_mission_id
			`: {},
	}

	selected := []string{
		"agent_snapshot_posts.*",
	}
	ms, err := s.dao.FindAgentSnapshotPostJoinSelect(daos.GetDBMainCtx(ctx),
		selected, joinFilters,
		map[string][]interface{}{
			"agent_snapshot_posts.created_at >= adddate(now(), interval -12 hour)":                       {},
			"agent_snapshot_missions.tool_set in (?) or agent_snapshot_posts.agent_store_mission_id > 0": {[]models.ToolsetType{models.ToolsetTypeTradeAnalyticsOnTwitter, models.ToolsetTypeTradeAnalyticsMentions, models.ToolsetTypeLuckyMoneys, models.ToolsetTypeMissionStore}},
			"agent_snapshot_posts.status = ?":                                                            {models.AgentSnapshotPostStatusInferSubmitted},
		},
		map[string][]interface{}{},
		[]string{}, 1, 50,
	)
	if err != nil {
		return errs.NewError(err)
	}

	for _, m := range ms {
		err := s.UpdateOffchainAutoOutputV2(ctx, m.ID)
		if err != nil {
			retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
		}
	}
	return retErr
}

func (s *Service) JobUpdateOffchainAutoOutput3Hour(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateOffchainAutoOutput3Hour",
		func() error {
			var retErr error
			{
				ms, err := s.dao.FindAgentSnapshotPost(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at <= adddate(now(), interval -2 hour)": {},
						"created_at >= adddate(now(), interval -3 hour)": {},
						"status = ?": {models.AgentSnapshotPostStatusInferSubmitted},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.UpdateOffchainAutoOutputV2(ctx, m.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
					}
				}
			}
			{
				ms, err := s.dao.FindAgentSnapshotPost(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at <= adddate(now(), interval -24 hour)": {},
						"created_at >= adddate(now(), interval -36 hour)": {},
						"status = ?": {models.AgentSnapshotPostStatusInferSubmitted},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.UpdateOffchainAutoOutputV2(ctx, m.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
					}
				}
			}
			{
				ms, err := s.dao.FindAgentSnapshotPostAction(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at <= adddate(now(), interval -2 hour)": {},
						"created_at >= adddate(now(), interval -3 hour)": {},
						"agent_snapshot_post_id > 0":                     {},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.UpdateOffchainAutoOutputV2(ctx, m.AgentSnapshotPostID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
					}
				}
			}
			{
				ms, err := s.dao.FindAgentWalletAction(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at <= adddate(now(), interval -2 hour)": {},
						"created_at >= adddate(now(), interval -3 hour)": {},
						"agent_snapshot_post_id > 0":                     {},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.UpdateOffchainAutoOutputV2(ctx, m.AgentSnapshotPostID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
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

func (s *Service) getTaskToolSet(assistant *models.AgentInfo, taskReq string) (string, string) {
	task := taskReq
	toolset := ""
	switch taskReq {
	case "reply_mentions":
		task = "reply"
		toolset = ""
	case "reply_non_mentions":
		task = "react_agent"
		toolset = "reply_non_mentions"
	case "follow":
		task = "react_agent"
		toolset = "follow"
	case "post":
		task = "react_agent"
		toolset = "post"
	case "quote_tweet":
		task = "react_agent"
		toolset = "quote_tweet"
	case "inscribe_tweet":
		task = "react_agent"
		toolset = "inscribe_tweet"
	case "shadow_reply":
		task = "shadow_reply"
		toolset = ""
	case "issue_token":
		task = "react_agent"
		toolset = "issue_token"
	case "tweet_news":
		task = "tweet_news"
		toolset = ""
	case "lucky_moneys":
		task = "lucky_moneys"
		toolset = "lucky_moneys"
	}
	if assistant != nil {
		if assistant.NetworkID == models.HERMES_CHAIN_ID && assistant.AgentContractID == "2" {
			// @thetickerisbot
			switch taskReq {
			case "quote_tweet":
				task = "quote_tweet"
				toolset = ""
			}
		}
		if assistant.NetworkID == models.BASE_CHAIN_ID && assistant.AgentContractID == "40" {
			// @thetickerisbot
			switch taskReq {
			case "quote_tweet":
				task = "quote_tweet"
				toolset = ""
			}
		}
	}
	return task, toolset
}

func (s *Service) callWakeup(logRequest *models.AgentSnapshotPost, assistant *models.AgentInfo) (string, error) {
	if assistant != nil {
		logRequest.AgentBaseModel = assistant.AgentBaseModel
	}
	var agentMetaDataRequest models.AgentMetadataRequest
	err := helpers.ConvertJsonObject(logRequest.AgentMetaData, &agentMetaDataRequest)
	if err != nil {
		return "", errs.NewError(err)
	}
	agentMetaDataRequest.KbAgents = []models.AgentWakeupKnowledgeBase{}
	request := &models.CallWakeupRequest{
		Toolkit:       []interface{}{},
		Prompt:        logRequest.UserPrompt,
		SystemPrompt:  logRequest.SystemPrompt,
		Task:          logRequest.Task,
		Toolset:       logRequest.Toolset,
		Model:         logRequest.AgentBaseModel,
		AgentMetaData: agentMetaDataRequest,
		ToolList:      logRequest.ToolList,
		MetaData: models.WakeupRequestMetadata{
			RefID:          logRequest.InferTxHash,
			SystemReminder: logRequest.SystemReminder, // DONT USE system reminder from assistant
			Params: models.ParamWakeupRequest{
				QuoteUsername: "cryptopunksbot",
				ReactMaxSteps: logRequest.ReactMaxSteps,
			},
		},
	}
	if assistant != nil {
		request.MetaData.TwitterId = assistant.TwitterID
		request.MetaData.TwitterUsername = assistant.TwitterUsername
		request.MetaData.AgentContractId = assistant.AgentContractID
		request.MetaData.ChainId = strconv.Itoa(int(assistant.NetworkID))
		request.MetaData.KnowledgeBaseId = assistant.KnowledgeBaseID
		knowledgeAgentsUsed, _ := s.KnowledgeUsecase.GetKBAgentsUsedOfSocialAgent(context.Background(), assistant.ID)
		if len(knowledgeAgentsUsed) > 0 {
			for _, item := range knowledgeAgentsUsed {
				itemAdd := models.AgentWakeupKnowledgeBase{
					KbId: item.KbId,
				}
				if item.AgentInfo != nil {
					itemAdd.ChainId = fmt.Sprintf("%v", item.AgentInfo.NetworkID)
				}
				if request.AgentMetaData.KbAgents == nil {
					request.AgentMetaData.KbAgents = []models.AgentWakeupKnowledgeBase{}
				}

				request.AgentMetaData.KbAgents = append(request.AgentMetaData.KbAgents, itemAdd)
			}
		}
		request.MetaData.TwitterUsername = assistant.TwitterUsername
	} else {
		request.MetaData.ChainId = strconv.Itoa(int(logRequest.NetworkID))
		request.MetaData.AgentContractId = "1"
	}
	body, err := helpers.CurlURLString(
		s.conf.AgentOffchain.Url+"/async/enqueue",
		"POST",
		map[string]string{
			"x-token": s.conf.AgentOffchain.ApiKey,
		},
		&request,
	)
	input, _ := json.Marshal(request)
	logger.Info("callWakeup", "async_enqueue",
		zap.Any("url", s.conf.AgentOffchain.Url+"/async/enqueue"),
		zap.Any("input", string(input)),
		zap.Any("output", body),
		zap.Any("err", err))
	if err != nil {
		return "", errs.NewError(err)
	}
	return body, err
}

func (s *Service) BatchPromptItemV2(ctx context.Context, agentInfo *models.AgentInfo, request *models.AgentSnapshotPost) (*models.AgentSnapshotPost, error) {
	if agentInfo != nil && agentInfo.TwitterUsername == "" && agentInfo.FarcasterID == "" {
		return request, errs.NewError(errs.ErrBadRequest)
	}
	if request.ToolList != "" {
		request.Toolset = "react_agent"
	}
	request.Task, request.Toolset = s.getTaskToolSet(agentInfo, string(request.Toolset))
	if len(request.HeadSystemPrompt) > 0 {
		request.SystemPrompt = request.HeadSystemPrompt + "\n\n" + request.SystemPrompt
	}
	if request.InferTxHash == "" {
		request.InferTxHash = helpers.RandomBigInt(12).Text(16)
	}
	body, err := s.callWakeup(request, agentInfo)
	if err != nil {
		return request, errs.NewError(err)
	}
	output := make(map[string]interface{})
	if err := json.Unmarshal([]byte(body), &output); err != nil {
		return request, errs.NewError(err)
	}
	id, ok := output["id"].(string)
	if ok {
		request.ResponseId = id
	}
	return request, nil
}

func (s *Service) JobAgentSnapshotPostStatusInferRefund(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobAgentSnapshotPostStatusInferRefund",
		func() error {
			var retErr error
			{
				ms, err := s.dao.FindAgentSnapshotPost(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"status = ?":                    {models.AgentSnapshotPostStatusInferFailed},
						"agent_info_id > 0":             {},
						"agent_snapshot_mission_id > 0": {},
					},
					map[string][]interface{}{},
					[]string{}, 0, 999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.AgentSnapshotPostStatusInferRefund(ctx, m.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
					}
				}
				// ms, err = s.dao.FindAgentSnapshotPost(daos.GetDBMainCtx(ctx),
				// 	map[string][]interface{}{
				// 		"status = ?":                    {models.AgentSnapshotPostStatusInferError},
				// 		"agent_info_id > 0":             {},
				// 		"agent_snapshot_mission_id > 0": {},
				// 		"fee > 0":                       {},
				// 	},
				// 	map[string][]interface{}{},
				// 	[]string{}, 0, 999999,
				// )
				// if err != nil {
				// 	return errs.NewError(err)
				// }
				// for _, m := range ms {
				// 	err := s.AgentSnapshotPostStatusInferRefund(ctx, m.ID)
				// 	if err != nil {
				// 		retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
				// 	}
				// }
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentSnapshotPostStatusInferRefund(ctx context.Context, snapshotPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentSnapshotPostStatusInferRefund_%d", snapshotPostID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					inferPost, err := s.dao.FirstAgentSnapshotPostByID(
						tx,
						snapshotPostID,
						map[string][]interface{}{
							"AgentInfo":        {},
							"AgentStore.Owner": {},
							"User":             {},
						},
						true,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if inferPost != nil &&
						inferPost.Status == models.AgentSnapshotPostStatusInferFailed &&
						inferPost.AgentInfo != nil &&
						inferPost.AgentSnapshotMissionID > 0 {
						//
						toolSet, err := s.dao.GetMissionToolset(tx, inferPost.AgentSnapshotMissionID)
						if err != nil {
							return errs.NewError(err)
						}
						//
						err = tx.Model(inferPost).
							UpdateColumn("status", models.AgentSnapshotPostStatusInferRefund).
							Error
						if err != nil {
							return errs.NewError(err)
						}
						if inferPost.AgentInfoID <= 0 {
							if inferPost.UserID > 0 {
								user := inferPost.User
								err = tx.Model(user).
									UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", inferPost.Fee)).
									Error
								if err != nil {
									return errs.NewError(err)
								}
								_ = s.dao.Create(
									tx,
									&models.UserTransaction{
										NetworkID: inferPost.NetworkID,
										EventId:   fmt.Sprintf("agent_trigger_refund_%d", inferPost.ID),
										UserID:    user.ID,
										Type:      models.UserTransactionTypeTriggerRefundFee,
										Amount:    inferPost.Fee,
										Status:    models.UserTransactionStatusDone,
									},
								)
							}
						} else {
							agentInfo := inferPost.AgentInfo
							err = tx.Model(agentInfo).
								UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", inferPost.Fee)).
								Error
							if err != nil {
								return errs.NewError(err)
							}
							_ = s.dao.Create(
								tx,
								&models.AgentEaiTopup{
									NetworkID:      agentInfo.NetworkID,
									EventId:        fmt.Sprintf("agent_trigger_refund_%d", inferPost.ID),
									AgentInfoID:    agentInfo.ID,
									Type:           models.AgentEaiTopupTypeRefund,
									Amount:         inferPost.Fee,
									Status:         models.AgentEaiTopupStatusDone,
									DepositAddress: agentInfo.ETHAddress,
									ToAddress:      agentInfo.ETHAddress,
									Toolset:        toolSet,
								},
							)
						}
					}
					if inferPost != nil &&
						inferPost.Status == models.AgentSnapshotPostStatusInferError &&
						inferPost.AgentInfo != nil &&
						inferPost.AgentSnapshotMissionID > 0 {
						agentInfo := inferPost.AgentInfo
						err = tx.Model(agentInfo).
							UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", inferPost.Fee)).
							Error
						if err != nil {
							return errs.NewError(err)
						}
						err = tx.Model(inferPost).Updates(
							map[string]interface{}{
								"fee": numeric.NewBigFloatFromString("0"),
							},
						).
							Error
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
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
