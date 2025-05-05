package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	sync "sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/openai"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
	openai2 "github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

func (s *Service) UpdateAgentScanEventSuccess(ctx context.Context, agentInfoID uint, lastTimeEvent *time.Time, lastId string) error {
	if lastTimeEvent == nil {
		lastTimeEvent = helpers.TimeNow()
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.AgentInfo{}).
		Where(
			"id = ?", agentInfoID,
		).
		Updates(
			map[string]interface{}{
				"scan_latest_time": lastTimeEvent,
				"scan_latest_id":   lastId,
				"scan_error":       "OK",
			},
		).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateAgentScanEventError(ctx context.Context, agentInfoID uint, errData error) error {
	if strings.Contains(errData.Error(), "not found") {
		return nil
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.AgentInfo{}).
		Where(
			"id = ?", agentInfoID,
		).
		Updates(
			map[string]interface{}{
				"scan_error": errData.Error(),
			},
		).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) GenerateTipAddress(ctx context.Context, agentInfoID uint) error {
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent.TipBtcAddress == "" {
		address, err := s.CreateBTCAddress(ctx)
		if err != nil {
			return errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_btc_address", address).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	if agent.TipEthAddress == "" {
		address, err := s.CreateETHAddress(ctx)
		if err != nil {
			return errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_eth_address", address).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	if agent.TipSolAddress == "" {
		address, err := s.CreateSOLAddress(ctx)
		if err != nil {
			return errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_sol_address", address).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) TwitterOauthCallbackV1(ctx context.Context, callbackUrl, address, code, agentID, clientID string) (map[string]string, error) {
	var err error
	var resp map[string]string
	if agentID == "" {
		err = s.TwitterOauthCallbackForRelink(ctx, callbackUrl, address, code, clientID)
		return resp, err
	} else if agentID == "0" {
		err = s.TwitterOauthCallbackForApiSubscription(ctx, callbackUrl, address, code, clientID)
		return resp, err
	} else if agentID == "1" {
		err = s.TwitterOauthCallbackForCreateAgent(ctx, callbackUrl, address, code, clientID)
		return resp, err
	} else if agentID == "2" {
		resp, err = s.TwitterOauthCallbackForClaimVideoReward(ctx, callbackUrl, address, code, clientID)
		return resp, err
	}

	agentInfo, err := s.SyncAgentInfoDetailByAgentID(ctx, agentID)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if agentInfo != nil {
		isFirstLinked := false
		// isAdvance := false
		oauthClientId := ""
		oauthClientSecret := ""
		if strings.EqualFold(clientID, s.conf.Twitter.OauthClientId) {
			oauthClientId = s.conf.Twitter.OauthClientId
			oauthClientSecret = s.conf.Twitter.OauthClientSecret
		} else {
			// isAdvance = true
			oauthClientId = agentInfo.OauthClientId
			oauthClientSecret = agentInfo.OauthClientSecret
		}

		respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKey(
			oauthClientId, oauthClientSecret,
			code, callbackUrl, address, agentID)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if respOauth != nil && respOauth.AccessToken != "" {
			twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
			if err != nil {
				return nil, errs.NewError(err)
			}

			user, err := s.GetUser(daos.GetDBMainCtx(ctx), agentInfo.NetworkID, address, true)
			if err != nil {
				return nil, errs.NewError(err)
			}

			user.TwitterID = twitterUser.ID
			user.TwitterAvatar = twitterUser.ProfileImageURL
			user.TwitterName = twitterUser.Name
			user.TwitterUsername = twitterUser.UserName

			err = s.dao.Save(daos.GetDBMainCtx(ctx), user)
			if err != nil {
				return nil, errs.NewError(err)
			}

			//
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]interface{}{}, false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo == nil {
				twitterInfo = &models.TwitterInfo{
					TwitterID: twitterUser.ID,
				}
				isFirstLinked = true
			}
			twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
			twitterInfo.TwitterName = twitterUser.Name
			twitterInfo.TwitterUsername = twitterUser.UserName
			twitterInfo.AccessToken = respOauth.AccessToken
			twitterInfo.RefreshToken = respOauth.RefreshToken
			twitterInfo.ExpiresIn = respOauth.ExpiresIn
			twitterInfo.Scope = respOauth.Scope
			twitterInfo.TokenType = respOauth.TokenType
			twitterInfo.OauthClientId = oauthClientId
			twitterInfo.OauthClientSecret = oauthClientSecret
			twitterInfo.Description = twitterUser.Description
			twitterInfo.RefreshError = "OK"

			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
			twitterInfo.ExpiredAt = &expiredAt
			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
			if err != nil {
				return nil, errs.NewError(err)
			}
			//

			updateFields := map[string]interface{}{
				"twitter_info_id":  twitterInfo.ID,
				"twitter_id":       twitterInfo.TwitterID,
				"twitter_username": twitterInfo.TwitterUsername,
				"scan_enabled":     true,
				"reply_enabled":    true,
			}

			err = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
				updateFields,
			).Error
			if err != nil {
				return nil, errs.NewError(err)
			}

			if isFirstLinked {
				// off default follow
				// go s.FollowListDefaultTwitters(ctx, agentInfo.ID)
				go s.AgentCreateMissionDefault(context.Background(), agentInfo.ID)
			}

		}
	}
	return nil, nil
}

func (s *Service) AgentCreateMissionDefault(ctx context.Context, agentInfoID uint) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	switch agentInfo.AgentType {
	case models.AgentInfoAgentTypeNormal,
		models.AgentInfoAgentTypeReasoning,
		models.AgentInfoAgentTypeKnowledgeBase:
		{
		}
	default:
		{
			return nil
		}
	}
	switch agentInfo.NetworkID {
	case models.SHARDAI_CHAIN_ID:
		{
			_ = daos.GetDBMainCtx(ctx).Exec(
				`
				INSERT INTO agent_snapshot_missions (created_at, updated_at, deleted_at, network_id, agent_info_id, user_prompt, interval_sec, reply_enabled, enabled, tool_set)
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Check and follow Twitter accounts that look interesting to you.',
					7200,
					1,
					1,
					'follow'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'follow'
					)
				and agent_infos.id = ?
						`,
				agentInfo.ID,
			).Error
		}
	case models.LOCAL_CHAIN_ID:
		{
		}
	default:
		{
			_ = daos.GetDBMainCtx(ctx).Exec(
				`
				INSERT INTO agent_snapshot_missions (created_at, updated_at, deleted_at, network_id, agent_info_id, user_prompt, interval_sec, reply_enabled, enabled, tool_set)
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Provide a single message to join the following conversation. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don''t include any instructions or extra words, just the raw message ready to post.',
					7200,
					1,
					1,
					'reply_mentions'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'reply_mentions'
					)
				and agent_infos.id = ?
				union
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Browse Twitter and choose ONE post and reply it. Keep the reply concise (under 128 chars), NO hashtags, links or emojis, and don''t include any instructions or extra words, just the raw reply ready to post. IMPORTANT: Immediately stop after replying one post. DO NOT REPLY YOUR OWN TWEET.',
					7200,
					1,
					1,
					'reply_non_mentions'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'reply_non_mentions'
					)
				and agent_infos.id = ?
				union
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Check and follow Twitter accounts that look interesting to you.',
					7200,
					1,
					1,
					'follow'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'follow'
					)
				and agent_infos.id = ?
				union
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Choose ONE topic that you like or dislike, and tweet about it with your own perspective.',
					7200,
					1,
					1,
					'post_search'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'post_search'
					)
				and agent_infos.id = ?;
						`,
				agentInfo.ID,
				agentInfo.ID,
				agentInfo.ID,
				agentInfo.ID,
			).Error
		}
	}
	return nil
}

func (s *Service) TwitterOauthCallbackForInternalData(ctx context.Context, callbackUrl, code string) error {
	respOauth, err := s.twitterAPI.TwitterOauthCallbackForInternalData(
		s.conf.Twitter.OauthClientIdForTwitterData, s.conf.Twitter.OauthClientSecretForTwitterData, code, callbackUrl)
	if err != nil {
		return errs.NewError(err)
	}

	if respOauth != nil && respOauth.AccessToken != "" {
		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
		if err != nil {
			return errs.NewError(err)
		}

		twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"twitter_id = ?": {twitterUser.ID},
			},
			map[string][]interface{}{}, false,
		)
		if err != nil {
			return errs.NewError(err)
		}

		if twitterInfo == nil {
			twitterInfo = &models.TwitterInfo{
				TwitterID: twitterUser.ID,
			}
		}
		twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
		twitterInfo.TwitterName = twitterUser.Name
		twitterInfo.TwitterUsername = twitterUser.UserName
		twitterInfo.AccessToken = respOauth.AccessToken
		twitterInfo.RefreshToken = respOauth.RefreshToken
		twitterInfo.ExpiresIn = respOauth.ExpiresIn
		twitterInfo.Scope = respOauth.Scope
		twitterInfo.TokenType = respOauth.TokenType
		twitterInfo.OauthClientId = s.conf.Twitter.OauthClientIdForTwitterData
		twitterInfo.OauthClientSecret = s.conf.Twitter.OauthClientSecretForTwitterData
		twitterInfo.Description = twitterUser.Description
		twitterInfo.RefreshError = "OK"

		expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
		twitterInfo.ExpiredAt = &expiredAt
		err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
		if err != nil {
			return errs.NewError(err)
		}
	}

	return nil
}

func (s *Service) TwitterOauthCallbackForRelink(ctx context.Context, callbackUrl, address, code, clientID string) error {
	oauthClientId := s.conf.Twitter.OauthClientId
	oauthClientSecret := s.conf.Twitter.OauthClientSecret

	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForRelink(
		oauthClientId, oauthClientSecret,
		code, callbackUrl, address)
	if err != nil {
		return errs.NewError(err)
	}

	if respOauth != nil && respOauth.AccessToken != "" {
		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
		if err != nil {
			return errs.NewError(err)
		}

		if twitterUser != nil {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]interface{}{}, false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if twitterInfo == nil {
				twitterInfo = &models.TwitterInfo{
					TwitterID: twitterUser.ID,
				}
			}
			twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
			twitterInfo.TwitterName = twitterUser.Name
			twitterInfo.TwitterUsername = twitterUser.UserName
			twitterInfo.AccessToken = respOauth.AccessToken
			twitterInfo.RefreshToken = respOauth.RefreshToken
			twitterInfo.ExpiresIn = respOauth.ExpiresIn
			twitterInfo.Scope = respOauth.Scope
			twitterInfo.TokenType = respOauth.TokenType
			twitterInfo.OauthClientId = oauthClientId
			twitterInfo.OauthClientSecret = oauthClientSecret
			twitterInfo.Description = twitterUser.Description
			twitterInfo.RefreshError = "OK"

			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
			twitterInfo.ExpiredAt = &expiredAt
			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
			if err != nil {
				return errs.NewError(err)
			}

			agentInfos, err := s.dao.FindAgentInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`twitter_id = ?`: {twitterUser.ID},
				},
				map[string][]interface{}{},
				[]string{},
				0, 100,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfos != nil {
				for _, agentInfo := range agentInfos {
					updateFields := map[string]interface{}{
						"twitter_info_id":  twitterInfo.ID,
						"twitter_id":       twitterInfo.TwitterID,
						"twitter_username": twitterInfo.TwitterUsername,
						"scan_enabled":     true,
					}

					err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
						updateFields,
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

func (s *Service) TwitterOauthCallbackForCreateAgent(ctx context.Context, callbackUrl, address, code, clientID string) error {
	oauthClientId := s.conf.Twitter.OauthClientId
	oauthClientSecret := s.conf.Twitter.OauthClientSecret

	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForCreateAgent(
		oauthClientId, oauthClientSecret,
		code, callbackUrl, address)
	if err != nil {
		return errs.NewError(err)
	}

	if respOauth != nil && respOauth.AccessToken != "" {
		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
		if err != nil {
			return errs.NewError(err)
		}

		if twitterUser != nil {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]interface{}{}, false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if twitterInfo == nil {
				twitterInfo = &models.TwitterInfo{
					TwitterID: twitterUser.ID,
				}
			}
			twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
			twitterInfo.TwitterName = twitterUser.Name
			twitterInfo.TwitterUsername = twitterUser.UserName
			twitterInfo.AccessToken = respOauth.AccessToken
			twitterInfo.RefreshToken = respOauth.RefreshToken
			twitterInfo.ExpiresIn = respOauth.ExpiresIn
			twitterInfo.Scope = respOauth.Scope
			twitterInfo.TokenType = respOauth.TokenType
			twitterInfo.OauthClientId = oauthClientId
			twitterInfo.OauthClientSecret = oauthClientSecret
			twitterInfo.Description = twitterUser.Description
			twitterInfo.RefreshError = "OK"

			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
			twitterInfo.ExpiredAt = &expiredAt
			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
			if err != nil {
				return errs.NewError(err)
			}

			agentInfos, err := s.dao.FindAgentInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`tmp_twitter_id = ?`:                    {twitterUser.ID},
					`(twitter_id is null or twitter_id="")`: {},
				},
				map[string][]interface{}{},
				[]string{},
				0, 100,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfos != nil {
				user, err := s.GetUser(daos.GetDBMainCtx(ctx), models.GENERTAL_NETWORK_ID, strings.ToLower(address), false)
				if err != nil {
					return errs.NewError(err)
				}

				if user != nil {
					user.TwitterID = twitterInfo.TwitterID
					user.TwitterName = twitterInfo.TwitterName
					user.TwitterUsername = twitterInfo.TwitterUsername
					user.TwitterAvatar = twitterInfo.TwitterAvatar
				}

				for _, agentInfo := range agentInfos {
					updateFields := map[string]interface{}{
						"creator":      strings.ToLower(address),
						"scan_enabled": true,
					}

					err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
						updateFields,
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

func (s *Service) TwitterOauthCallbackForClaimVideoReward(ctx context.Context, callbackUrl, address, code, clientID string) (map[string]string, error) {
	mapResp := map[string]string{}
	oauthClientId := s.conf.Twitter.OauthClientId
	oauthClientSecret := s.conf.Twitter.OauthClientSecret

	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForVideoReward(
		oauthClientId, oauthClientSecret,
		code, callbackUrl, address)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if respOauth != nil && respOauth.AccessToken != "" {
		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if twitterUser != nil {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]interface{}{}, false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo == nil {
				twitterInfo = &models.TwitterInfo{
					TwitterID: twitterUser.ID,
				}
			}
			twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
			twitterInfo.TwitterName = twitterUser.Name
			twitterInfo.TwitterUsername = twitterUser.UserName
			twitterInfo.AccessToken = respOauth.AccessToken
			twitterInfo.RefreshToken = respOauth.RefreshToken
			twitterInfo.ExpiresIn = respOauth.ExpiresIn
			twitterInfo.Scope = respOauth.Scope
			twitterInfo.TokenType = respOauth.TokenType
			twitterInfo.OauthClientId = oauthClientId
			twitterInfo.OauthClientSecret = oauthClientSecret
			twitterInfo.Description = twitterUser.Description
			twitterInfo.RefreshError = "OK"

			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
			twitterInfo.ExpiredAt = &expiredAt
			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
			if err != nil {
				return nil, errs.NewError(err)
			}

			//user address
			privyWallet, err := s.dao.FirstPrivyWallet(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				}, map[string][]interface{}{}, []string{})
			if err != nil {
				return nil, errs.NewError(err)
			}

			if privyWallet == nil {
				return nil, errs.NewError(errs.ErrRecordNotFound)
			}

			updateFields := map[string]interface{}{
				"user_address": strings.ToLower(address),
			}

			err = daos.GetDBMainCtx(ctx).Model(privyWallet).Updates(
				updateFields,
			).Error
			if err != nil {
				return nil, errs.NewError(err)
			}
			mapResp["address"] = privyWallet.Address
			mapResp["twitter_id"] = twitterUser.ID
		}

	}

	return mapResp, nil
}

// func (s *Service) TwitterOauthCallbackForClaimVideoReward(ctx context.Context, callbackUrl, address, code, clientID string) error {
// 	oauthClientId := s.conf.Twitter.OauthClientId
// 	oauthClientSecret := s.conf.Twitter.OauthClientSecret

// 	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForVideoReward(
// 		oauthClientId, oauthClientSecret,
// 		code, callbackUrl, address)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}

// 	if respOauth != nil && respOauth.AccessToken != "" {
// 		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}

// 		if twitterUser != nil {
// 			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
// 				map[string][]interface{}{
// 					"twitter_id = ?": {twitterUser.ID},
// 				},
// 				map[string][]interface{}{}, false,
// 			)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}

// 			if twitterInfo == nil {
// 				twitterInfo = &models.TwitterInfo{
// 					TwitterID: twitterUser.ID,
// 				}
// 			}
// 			twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
// 			twitterInfo.TwitterName = twitterUser.Name
// 			twitterInfo.TwitterUsername = twitterUser.UserName
// 			twitterInfo.AccessToken = respOauth.AccessToken
// 			twitterInfo.RefreshToken = respOauth.RefreshToken
// 			twitterInfo.ExpiresIn = respOauth.ExpiresIn
// 			twitterInfo.Scope = respOauth.Scope
// 			twitterInfo.TokenType = respOauth.TokenType
// 			twitterInfo.OauthClientId = oauthClientId
// 			twitterInfo.OauthClientSecret = oauthClientSecret
// 			twitterInfo.Description = twitterUser.Description
// 			twitterInfo.RefreshError = "OK"

// 			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
// 			twitterInfo.ExpiredAt = &expiredAt
// 			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}

// 			insts, err := s.dao.FindAgentVideo(daos.GetDBMainCtx(ctx),
// 				map[string][]interface{}{
// 					`owner_twitter_id = ?`: {twitterUser.ID},
// 				},
// 				map[string][]interface{}{},
// 				[]string{},
// 				0, 9999,
// 			)

// 			if err != nil {
// 				return errs.NewError(err)
// 			}

// 			if insts != nil && len(insts) > 0 {
// 				user, err := s.GetUser(daos.GetDBMainCtx(ctx), models.GENERTAL_NETWORK_ID, strings.ToLower(address), false)
// 				if err != nil {
// 					return errs.NewError(err)
// 				}

// 				if user != nil {
// 					user.TwitterID = twitterInfo.TwitterID
// 					user.TwitterName = twitterInfo.TwitterName
// 					user.TwitterUsername = twitterInfo.TwitterUsername
// 					user.TwitterAvatar = twitterInfo.TwitterAvatar
// 					err = s.dao.Save(daos.GetDBMainCtx(ctx), user)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}
// 				}

// 				for _, inst := range insts {
// 					updateFields := map[string]interface{}{
// 						"user_address": strings.ToLower(address),
// 					}

// 					err := daos.GetDBMainCtx(ctx).Model(inst).Updates(
// 						updateFields,
// 					).Error
// 					if err != nil {
// 						return errs.NewError(err)
// 					}
// 				}
// 			}

// 		}

// 	}

// 	return nil
// }

func (s *Service) CreateUpdateUserTwitter(tx *gorm.DB, userTwitterID string) (*models.TwitterUser, error) {
	tweetUser, err := s.dao.FirstTwitterUser(tx,
		map[string][]interface{}{
			"twitter_id = ?": {userTwitterID},
		}, map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	info, err := s.twitterAPI.GetTwitterUserInfoID(userTwitterID)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if tweetUser == nil {
		tweetUser = &models.TwitterUser{
			TwitterID:       info.TwitterID,
			TwitterUsername: info.TwitterUsername,
			Name:            info.Name,
			ProfileUrl:      info.ProfileUrl,
			FollowersCount:  info.FollowersCount,
			FollowingsCount: info.FollowingsCount,
			IsBlueVerified:  info.IsBlueVerified,
			JoinedAt:        info.CreatedAt,
		}
	} else {
		tweetUser.TwitterUsername = info.TwitterUsername
		tweetUser.Name = info.Name
		tweetUser.ProfileUrl = info.ProfileUrl
	}

	err = s.dao.Save(tx, tweetUser)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tweetUser, nil
}

func (s *Service) CreateUpdateUserTwitterByUserName(tx *gorm.DB, username string) (*models.TwitterUser, error) {
	tweetUser, err := s.dao.FirstTwitterUser(tx,
		map[string][]interface{}{
			"twitter_username = ?": {username},
		}, map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if tweetUser == nil {
		info, err := s.twitterAPI.GetTwitterByUserName(username)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if info != nil {
			tweetUser = &models.TwitterUser{
				TwitterID:       info.ID,
				TwitterUsername: info.UserName,
				Name:            info.Name,
				ProfileUrl:      info.ProfileImageURL,
			}

			err = s.dao.Create(tx, tweetUser)
			if err != nil {
				return nil, errs.NewError(err)
			}
		}
	}
	return tweetUser, nil
}

func (s *Service) JobUpdateTwitterAccessToken(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateTwitterAccessToken",
		func() error {
			var retErr error
			twitterInfos, err := s.dao.FindTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"expired_at <= now()":                             {},
					"expired_at >= adddate(now(), interval -24 hour)": {},
				},
				map[string][]interface{}{},
				[]string{
					"updated_at asc",
				}, 0, 50,
			)
			if err != nil {
				return errs.NewError(err)
			}
			for _, twitterInfo := range twitterInfos {
				err := s.UpdateTwitterAccessToken(ctx, twitterInfo.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterInfo.ID))
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

func (s *Service) UpdateTwitterAccessToken(ctx context.Context, twitterInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("UpdateTwitterAccessToken_%d", twitterInfoID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterInfo, err := s.dao.FirstTwitterInfoByID(tx, twitterInfoID, map[string][]interface{}{}, true)
					if err != nil {
						return errs.NewError(err)
					}
					var authInfo *twitter.TwitterTokenResponse

					if twitterInfo.OauthClientId != "" && twitterInfo.OauthClientSecret != "" {
						authInfo, err = s.twitterAPI.GetTwitterAccessTokenWithKey(twitterInfo.OauthClientId,
							twitterInfo.OauthClientSecret, twitterInfo.RefreshToken)
						// if err != nil {
						// 	return errs.NewError(err)
						// }
					} else {
						authInfo, err = s.twitterAPI.GetTwitterAccessToken(twitterInfo.RefreshToken)
						// if err != nil {
						// 	return errs.NewError(err)
						// }
					}

					if authInfo != nil && err == nil {
						twitterInfo.AccessToken = authInfo.AccessToken
						twitterInfo.RefreshToken = authInfo.RefreshToken
						twitterInfo.ExpiresIn = authInfo.ExpiresIn
						twitterInfo.Scope = authInfo.Scope
						twitterInfo.TokenType = authInfo.TokenType
						expiredAt := time.Now().Add(time.Second * time.Duration(authInfo.ExpiresIn-(60*20)))
						twitterInfo.ExpiredAt = &expiredAt
						twitterInfo.RefreshError = "OK"
						err = s.dao.Save(tx, twitterInfo)
						if err != nil {
							return errs.NewError(err)
						}
					} else {
						twitterInfo.RefreshError = err.Error()
						err = s.dao.Save(tx, twitterInfo)
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
	_ = s.UpdateAgentTwitterInfo(ctx, twitterInfoID)
	return nil
}

func (s *Service) UpdateAgentTwitterInfo(ctx context.Context, twitterInfoID uint) error {
	agents, err := s.dao.FindAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_info_id = ?": {twitterInfoID},
		},
		map[string][]interface{}{
			"TwitterInfo": {},
		},
		[]string{},
		0,
		999999,
	)
	if err != nil {
		return errs.NewError(err)
	}
	for _, agent := range agents {
		func(agent *models.AgentInfo) error {
			if agent.TwitterInfo != nil {
				userMe, err := helpers.GetTwitterUserMe(agent.TwitterInfo.AccessToken)
				if err != nil {
					return errs.NewError(err)
				}
				if userMe != nil && userMe.Data.UserName != "" && userMe.Data.Name != "" {
					err = daos.GetDBMainCtx(ctx).Model(agent).
						UpdateColumn("twitter_username", userMe.Data.UserName).
						UpdateColumn("agent_name", userMe.Data.Name).
						UpdateColumn("twitter_verified", userMe.Data.Verified).Error
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).Model(agent.TwitterInfo).
						UpdateColumn("twitter_username", userMe.Data.UserName).
						UpdateColumn("twitter_name", userMe.Data.Name).
						UpdateColumn("twitter_avatar", userMe.Data.ProfileImageURL).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		}(agent)
	}
	return nil
}

func (s *Service) GetTwitterVerified(tx *gorm.DB, agentInfoID uint) (bool, error) {
	m, err := s.dao.FirstAgentInfoByID(
		tx,
		agentInfoID,
		map[string][]interface{}{
			"TwitterInfo": {},
		},
		false,
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	if m == nil {
		return false, errs.NewError(errs.ErrBadRequest)
	}
	return m.TwitterVerified, nil
}

func (s *Service) GetTwitterPostMaxChars(tx *gorm.DB, agentInfoID uint) (uint, error) {
	verified, err := s.GetTwitterVerified(tx, agentInfoID)
	if err != nil {
		return 0, errs.NewError(err)
	}
	if verified {
		return 4000, nil
	}
	return 280, nil
}

func (s *Service) GetAgentTokenInfoFromContractAddress(ctx context.Context, tokenAddress string) (string, string, error) {
	tokenAddress = strings.ToLower(tokenAddress)
	if tokenAddress != "" && tokenAddress != "no" && tokenAddress != "yes" && tokenAddress != "pending" {
		tokenMetaData, err := s.blockchainUtils.SolanaTokenMetaData(tokenAddress)
		if err != nil {
			return "", "", nil
		}
		return tokenMetaData.Name, tokenMetaData.Symbol, nil
	}
	return "", "", nil
}

func (s *Service) UpdateAgentFarcasterInfo(ctx context.Context, agentID string, fID string, fUsername string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}
			if agentInfo != nil {
				agentInfo, _ = s.dao.FirstAgentInfoByID(tx, agentInfo.ID, map[string][]interface{}{}, true)
				agentInfo.FarcasterID = fID
				agentInfo.FarcasterUsername = fUsername
				err = s.dao.Save(tx, agentInfo)
				if err != nil {
					return errs.NewError(err)
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) PreviewAgentSystemPromp(ctx context.Context, personality, question string) (string, error) {
	aiStr, err := s.openais["Agent"].TestAgentPersinality(personality, question, s.conf.AgentOffchainChatUrl)
	if err != nil {
		return "", errs.NewError(err)
	}
	return aiStr, nil
}

func (s *Service) RetrieveKnowledge(agentModel string, messages []openai2.ChatCompletionMessage,
	knowledgeBases []*models.KnowledgeBase, topK *int, threshold *float64) (string, error) {
	if len(knowledgeBases) == 0 {
		return "", errs.NewError(errors.New("knowledge bases is empty"))
	}
	if agentModel == "" {
		agentModel = "NousResearch/Hermes-3-Llama-3.1-70B-FP8"
	}
	systemPrompt := openai.GetSystemPromptFromLLMMessage(messages)
	isKbAgent := false
	if systemPrompt == "" {
		isKbAgent = true
		systemPrompt = "You are a helpful assistant."
	}
	_ = isKbAgent

	userPromptInput := openai.LastUserPrompt(messages)
	retrieveQuery := userPromptInput

	topKQuery := 5
	if topK != nil {
		topKQuery = *topK
	}
	th := 0.4
	if threshold != nil {
		th = *threshold
	}

	request := serializers.RetrieveKnowledgeBaseRequest{
		Query: retrieveQuery,
		TopK:  topKQuery,
		Kb: []string{
			knowledgeBases[0].KbId,
		},
		Threshold: th,
	}

	// retry
	var (
		body string
		err  error
	)
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		body, err = helpers.CurlURLString(
			s.conf.KnowledgeBaseConfig.QueryServiceUrl,
			"POST",
			map[string]string{},
			&request,
		)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	response := &serializers.RetrieveKnowledgeBaseResponse{}
	err = json.Unmarshal([]byte(body), response)
	if err != nil {
		return "", errs.NewError(err)
	}

	searchResult := ""
	for _, item := range response.Result {
		searchResult = searchResult + item.Content + "\n\n"
	}
	options := map[string]interface{}{}
	userPrompt := fmt.Sprintf("Use the following pieces of retrieved context to answer the question. If there is not enough information in the retrieved context to answer the question, just say something you know about the topic."+
		"\n\nQuestion: %v\nContext: \n\n%vAnswer:", userPromptInput, searchResult)
	//answer prompt
	payloadAgentChat := []openai2.ChatCompletionMessage{
		{
			Role:    openai2.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai2.ChatMessageRoleUser,
			Content: userPrompt,
		},
	}
	if agentModel == "DeepSeek-R1-Distill-Llama-70B" {
		options = map[string]interface{}{
			"temperature": 0.7,
			"max_tokens":  4096,
		}
	} else {
		options = map[string]interface{}{
			"temperature": 0,
			"top_p":       0.01,
			"max_tokens":  1024,
		}
	}

	messageCallLLM, _ := json.Marshal(&payloadAgentChat)
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
	}
	stringResp, err := s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), agentModel, url, options)
	if err != nil {
		return "", errs.NewError(err)
	}
	return stringResp, nil
}

func (s *Service) StreamRetrieveKnowledge(ctx context.Context, agentModel string, messages []openai2.ChatCompletionMessage,
	knowledgeBases []*models.KnowledgeBase, topK *int, threshold *float64,
	outputChan chan *models.ChatCompletionStreamResponse,
	errChan chan error, doneChan chan bool) {
	if len(knowledgeBases) == 0 {
		errChan <- errs.NewError(errors.New("knowledge bases is empty"))
		return
	}
	if agentModel == "" {
		agentModel = "NousResearch/Hermes-3-Llama-3.1-70B-FP8"
	}
	systemPrompt := openai.GetSystemPromptFromLLMMessage(messages)
	isKbAgent := false
	if systemPrompt == "" {
		isKbAgent = true
		systemPrompt = "You are a helpful assistant."
	}
	_ = isKbAgent
	go func() {
		outputChan <- &models.ChatCompletionStreamResponse{
			Message: "Start generating the query.",
			Code:    http.StatusProcessing,
		}
	}()
	idRequest := time.Now().UnixMicro()
	retrieveQuery, errGenerateQuery, conversation := s.GenerateKnowledgeQuery(agentModel, messages)
	_ = conversation
	if errGenerateQuery != nil {
		errChan <- errs.NewError(errors.New("ERROR_GENERATE_QUERY"))
		return
	}
	logger.Info("stream_retrieve_knowledge", "generate query", zap.Any("id_request", idRequest), zap.Any("retrieveQuery", retrieveQuery), zap.Any("input", messages))
	if retrieveQuery == nil {
		str := ""
		retrieveQuery = &str
	}

	analysedResultChanel := make(chan string)
	go func() {
		topKQuery := 20
		if topK != nil {
			topKQuery = *topK
		}
		th := 0.2
		if threshold != nil {
			th = *threshold
		}

		request := serializers.RetrieveKnowledgeBaseRequest{
			Query: *retrieveQuery,
			TopK:  topKQuery,
			Kb: []string{
				knowledgeBases[0].KbId,
			},
			Threshold: th,
		}
		go func() {
			outputChan <- &models.ChatCompletionStreamResponse{
				Message: "Start searching query in the RAG system.",
				Code:    http.StatusProcessing,
			}
		}()
		searchResponse, err := s.GetResultFromRagSearch(&request)
		if err != nil {
			errChan <- err
			return
		}

		searchedResult := []string{}
		for _, item := range searchResponse.Result {
			searchedResult = append(searchedResult, item.Content)
		}

		go func() {
			outputChan <- &models.ChatCompletionStreamResponse{
				Message: "Start analyzing the search result.",
				Code:    http.StatusProcessing,
			}
		}()
		logger.Info("stream_retrieve_knowledge", "searched result", zap.Any("id_request", idRequest), zap.Any("searchedResult", searchedResult), zap.Any("input", request))
		analysedResult, err := s.AnalyseSearchResults(agentModel, systemPrompt, *retrieveQuery, searchedResult)
		if err != nil {
			errChan <- err
			return
		}

		logger.Info("stream_retrieve_knowledge", "analyze result", zap.Any("id_request", idRequest), zap.Any("query", retrieveQuery), zap.Any("analyzed Result", analysedResult))
		analysedResultChanel <- analysedResult
	}()
	toolCallData := ""
	if knowledgeBases[0].ID == 211 { //eth denver agent
		var err error
		toolCallData, err = s.GetResultFromToolCall(*retrieveQuery)
		if err != nil {
			errChan <- err
			return
		}
	}
	logger.Info("stream_retrieve_knowledge", "tool call data", zap.Any("id_request", idRequest), zap.Any("query", retrieveQuery), zap.Any("tool call data", toolCallData))
	// wait finish get analysedResult
	analysedResult := <-analysedResultChanel
	options := map[string]interface{}{}
	//answer prompt
	question := openai.GetQuestionFromLLMMessage(messages)
	payloadAgentChat := []openai2.ChatCompletionMessage{
		{
			Role:    openai2.ChatMessageRoleSystem,
			Content: systemPrompt,
		}}

	if len(messages) > 0 {
		payloadAgentChat = messages[:len(messages)-1]
	}
	questionPrompt := fmt.Sprintf("Generate a response to the user's query based strictly on the user question and the provided information.\n\n### Guidelines:\n- Prioritize database data over website data when answering.\n- The response must be concise and directly relevant.\n- No external knowledge should be introduced beyond the provided sources.\n- Ensure clarity and alignment with ETHDenver-related context.\n- Prefer structured lists over paragraphs whenever possible to enhance readability.\n- If the response involves listing events, ensure they are formatted as follows:\n\nRequired Format for Events:\n```\n<Event name> (<Speaker 1>; <Speaker 2>; ...; <Speaker n>) - <Local start time> - <Stage/Location name>\n```\n\n- Speakers should be listed in the order provided. If they have an affiliation, include it exactly as given.\n- The local start time must be preserved in its original format.\n- The stage or location name should appear at the end.\n- If only one speaker is listed, follow the same format without modification.\n- If no speaker is listed, ignore the speaker listing part of the format.\n- If multiple events are listed, each should follow the format on a new line.\n\nExample of correct event with speakers output:\n- Easy-to-Miss Solidity Bugs (Jonathan Mevs - Quantstamp; Michael Boyle - Quantstamp) - February 24, 2025 at 10:50 AM - Captain Ethereum Stage\n\nExample of correct event without speakers output:\n- Messari - Feb 27, 2025 at 1:30 PM - BUIDL Event Hall\n\n### User Question:\n%v\n\n### Relevant Information from Database (Primary Source):\n%v\n\n### Relevant Information from the Website:\n%v\n\n### Final Answer:\n(Provide a precise, ETHDenver-relevant response following these guidelines.)\n",
		question, toolCallData, analysedResult)
	if knowledgeBases[0].ID != 211 { // not is eth denver agent
		questionPrompt = fmt.Sprintf("Use the following context from the conversation to answer the question. If the context is insufficient, you may draw from external knowledge to provide a relevant answer.\n\nContext: \n%v\n\nQuestion: \n%v\n\nAnswer:",
			analysedResult, question)
	}

	payloadAgentChat = append(payloadAgentChat, openai2.ChatCompletionMessage{
		Role:    openai2.ChatMessageRoleUser,
		Content: questionPrompt,
	})

	url := s.conf.AgentOffchainChatUrl
	apiKey := s.conf.KnowledgeBaseConfig.OnchainAPIKey

	if s.conf.KnowledgeBaseConfig.OnChainUrl != "" {
		url = s.conf.KnowledgeBaseConfig.OnChainUrl
	}
	seedAnswer := 1234
	chainID := fmt.Sprintf("%v", knowledgeBases[0].NetworkID)
	llmRequest := openai2.ChatCompletionRequest{
		Stream:      true,
		Messages:    payloadAgentChat,
		Model:       agentModel,
		Temperature: 0.01,
		MaxTokens:   4096,
		Seed:        &seedAnswer,
		Metadata: map[string]string{
			"chain_id":         chainID,
			"onchain_internal": "1",
		},
	}
	go func() {
		outputChan <- &models.ChatCompletionStreamResponse{
			Message: "Start generating the result.",
			Code:    http.StatusProcessing,
		}
	}()
	logger.Info("stream_retrieve_knowledge", "start call finish result", zap.Any("id_request", idRequest), zap.Any("payloadAgentChat", payloadAgentChat), zap.Any("options", options))
	s.openais["Agent"].CallStreamOnchainEternalLLM(ctx, url, apiKey, llmRequest, outputChan, errChan, doneChan)
}

func (s *Service) GenerateKnowledgeQuery(baseModel string, histories []openai2.ChatCompletionMessage) (*string, error, string) {
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
	}

	today := time.Now().Format("Jan 02, 2006")
	conversation := fmt.Sprintf("Remember that today is: %v\n\n", today)
	for _, item := range histories {
		if item.Role == openai2.ChatMessageRoleSystem {
			continue
		}

		conversation += fmt.Sprintf("%v: %v\n", strings.ToTitle(item.Role), item.Content)
	}
	systemPrompt := openai.GetSystemPromptFromLLMMessage(histories)
	if systemPrompt == "" {
		systemPrompt = "You are a helpfully assistant"
	}
	question := openai.GetQuestionFromLLMMessage(histories)
	if question == "" {
		question = "Hi"
	}
	type historyMsg struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	historiesPrompt := []historyMsg{}
	for i := 1; i < len(histories)-1; i++ {
		historiesPrompt = append(historiesPrompt, historyMsg{
			Role:    strings.ToLower(histories[i].Role),
			Content: histories[i].Content,
		})
	}
	generateQueryPrefix := "Based on the conversation history and the user question below, generate a concise query that can be used to retrieve relevant information.\n\n### Instruction:\n- Generate a concise query based on the conversation history and the user question.\n- Output the query in **stringified JSON format** with a key `\"query\"`.\n- Do not include additional explanations or comments—just the JSON.\n\n### Example\n\n**Conversation History:**\n[{{\"role\":\"user\",\"content\":\"What is French cuisine?\"}},{{\"role\":\"assistant\",\"content\":\"French cuisine refers to the traditional cooking styles of France, famous for its rich flavors and varied dishes.\"}}]\n\n**User Question:**\nWhat is the most popular?\n\n**Output:**  \n```json\n{{\n    \"query\": \"popular French cuisine\"\n}}\n```\n\n### Input\n\nRemember that today is: %v\n\n**Conversation History:**\n%v\n\n**User Question:**\n%v\n\n### Answer\n"
	userPrompt := fmt.Sprintf(generateQueryPrefix, today, historiesPrompt, question)
	messages := []openai2.ChatCompletionMessage{
		{
			Role:    openai2.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai2.ChatMessageRoleUser,
			Content: userPrompt,
		},
	}

	maxRetry := 10
	messageCallLLM, _ := json.Marshal(&messages)

	var queryStringResp *string
	for i := 1; i <= maxRetry; i++ {
		if i > 1 {
			time.Sleep(time.Second)
		}

		stringResp, err := s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), baseModel, url, map[string]interface{}{
			"temperature": 0.7,
			"max_tokens":  4096,
		})
		if err != nil || stringResp == "" {
			continue
		}
		// find { and } in stringResp
		startIndex := strings.Index(stringResp, "{")
		endIndex := strings.LastIndex(stringResp, "}")
		if startIndex == -1 || endIndex == -1 {
			continue
		}
		queryStringRespJson := stringResp[startIndex : endIndex+1]
		queryStringRespMap := map[string]string{}
		err = json.Unmarshal([]byte(queryStringRespJson), &queryStringRespMap)
		if err != nil {
			continue
		}
		if _, ok := queryStringRespMap["query"]; !ok {
			continue
		}

		val := queryStringRespMap["query"]
		queryStringResp = &val
		break
	}

	return queryStringResp, nil, conversation
}
func (s *Service) AnalyseSearchResults(baseModel string, systemPrompt string, query string, searchedResult []string) (string, error) {
	batchSize := 5
	start := 0
	analyzeResult := ""
	countRequest := len(searchedResult) / batchSize
	if len(searchedResult)%batchSize != 0 {
		countRequest++
	}
	listAnalyzeResult := make([]string, countRequest)
	wg := sync.WaitGroup{}
	wg.Add(countRequest)
	for i := 0; i < countRequest; i++ {
		go func(index int) {
			defer wg.Done()
			start = index * batchSize
			end := start + batchSize
			if end > len(searchedResult) {
				end = len(searchedResult)
			}
			searchResult := ""
			for _, item := range searchedResult[start:end] {
				searchResult = searchResult + item + "\n\n"
			}
			url := s.conf.AgentOffchainChatUrl
			if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
				url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
			}

			generateQueryPrefix := "Act as a critical information analyst, skilled in extracting only the most essential insights from search results.\n\n## Task:\nAnalyze the provided search results and extract only the most critical insights directly relevant to the given question.\n\n## Instructions:\n- Strictly extract only essential information. No introductions, summaries, or extra context—only the key insights.\n- Ensure accuracy. Verify time, location, and context before including any insight.\n- Be concise and precise. Remove redundant details, filler content, and tangential information.\n- No assumptions or external knowledge. Only use information explicitly stated in the search results.\n- Maintain neutrality and clarity. Present insights objectively without speculation.\n\n## Input:\nQuestion: %v\nSearch Results: %v\n\n## Response Format (Strictly Follow This):\n- Directly list the critical insights only—no introductions or explanations.\n- If multiple key insights exist, format them as bullet points.\n- Each point should be as concise as possible while preserving meaning.\n\n## Example Output:\n- [Critical insight 1]\n- [Critical insight 2]\n- [Critical insight 3]\n"
			userPrompt := fmt.Sprintf(generateQueryPrefix, query, searchResult)
			messages := []openai2.ChatCompletionMessage{
				{
					Role:    openai2.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai2.ChatMessageRoleUser,
					Content: userPrompt,
				},
			}

			maxRetry := 10
			messageCallLLM, _ := json.Marshal(&messages)
			stringResp := ""
			var err error
			for i := 1; i <= maxRetry; i++ {
				if i > 1 {
					time.Sleep(time.Second)
				}

				stringResp, err = s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), baseModel, url, map[string]interface{}{
					"temperature": 0.7,
					"max_tokens":  4096,
				})
				if err != nil || stringResp == "" {
					continue
				}
				break
			}
			listAnalyzeResult[index] = stringResp
		}(i)
	}
	wg.Wait()
	for _, result := range listAnalyzeResult {
		analyzeResult = analyzeResult + result + "\n"
	}
	return analyzeResult, nil
}

func (s *Service) GetResultFromToolCall(query string) (string, error) {
	if len(query) == 0 {
		return "", nil
	}
	url := fmt.Sprintf("%v?question=%v", s.conf.KnowledgeBaseConfig.ToolCallServiceUrl, query)
	for i := 0; i < 10; i++ {
		body, err := helpers.CurlURLString(
			url,
			"GET",
			map[string]string{},
			nil,
		)
		if err != nil {
			continue
		}
		res := make(map[string]interface{})
		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			return "[]", nil
		}
		if res["result"] == nil {
			return "[]", nil
		}
		result, ok := res["result"].(map[string]interface{})
		if !ok {
			return "[]", nil
		}
		data, _ := json.Marshal(result["data"])
		return string(data), nil
	}
	return "", fmt.Errorf("can not get tool call result with query %v", query)
}

func (s *Service) GetResultFromRagSearch(request *serializers.RetrieveKnowledgeBaseRequest) (*serializers.RetrieveKnowledgeBaseResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("empty query")
	}

	for i := 1; i <= 10; i++ {
		body, err := helpers.CurlURLString(
			s.conf.KnowledgeBaseConfig.QueryServiceUrl,
			"POST",
			map[string]string{},
			&request,
		)
		if err != nil {
			continue
		}
		var response serializers.RetrieveKnowledgeBaseResponse
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			break
		}
		return &response, nil
	}
	return nil, fmt.Errorf("can not get search result with request %v", request)
}

func (s *Service) PreviewAgentSystemPrompV1(ctx context.Context,
	messages string, agentId *uint, kbIdFromKnowledegeBase *string, modelNameFromRequest *string) (string, error) {
	var agentInfo *models.AgentInfo
	baseModel := "NousResearch/Hermes-3-Llama-3.1-70B-FP8"
	if agentId != nil {
		agentInfo, _ = s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx), *agentId, map[string][]interface{}{}, false)
	}
	llmMessage := []openai2.ChatCompletionMessage{}
	err := json.Unmarshal([]byte(messages), &llmMessage)
	if err != nil {
		return "", errs.NewError(errors.New("invalid message request"))
	}

	systemContent := openai.GetSystemPromptFromLLMMessage(llmMessage)
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
		if agentInfo != nil && agentInfo.AgentBaseModel != "" {
			baseModel = agentInfo.AgentBaseModel
		}
	}
	if modelNameFromRequest != nil && *modelNameFromRequest != "" {
		baseModel = *modelNameFromRequest
	}

	{
		if s.conf.KnowledgeBaseConfig.EnableSimulation && agentInfo != nil {
			// get last knowledge base of agent
			var knowledgeBaseUse *models.KnowledgeBase
			agentInfoKnowledgeBase, _ := s.dao.FirstAgentInfoKnowledgeBaseByAgentInfoID(
				daos.GetDBMainCtx(ctx),
				agentInfo.ID,
				map[string][]interface{}{
					"KnowledgeBase": {}, // must preload
				},
				[]string{"id desc"},
			)
			if agentInfoKnowledgeBase != nil && agentInfoKnowledgeBase.KnowledgeBase != nil {
				knowledgeBaseUse = agentInfoKnowledgeBase.KnowledgeBase
			}

			if kbIdFromKnowledegeBase != nil {
				knowledgeBaseFromQuery, _ := s.dao.FirstKnowledgeBase(daos.GetDBMainCtx(ctx), map[string][]interface{}{
					"kb_id = ?": {*kbIdFromKnowledegeBase},
				}, map[string][]interface{}{}, []string{"id desc"}, false)
				if knowledgeBaseFromQuery != nil {
					knowledgeBaseUse = knowledgeBaseFromQuery
				}
			}
			if knowledgeBaseUse != nil {
				retrieveKnowledgeBaseResponse, err := s.RetrieveKnowledge(baseModel, llmMessage, []*models.KnowledgeBase{
					knowledgeBaseUse,
				}, nil, nil)

				if err != nil {
					return "", err
				}

				return retrieveKnowledgeBaseResponse, nil
			}
		}
	}

	llmMessage = openai.UpdateSystemPromptInLLMRequest(llmMessage, systemContent)
	messageCallLLM, _ := json.Marshal(&llmMessage)
	aiStr, err := s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), baseModel, url,
		map[string]interface{}{
			"top_p":      0.01,
			"max_tokens": 4096,
		})
	if err != nil {
		return "", errs.NewError(err)
	}
	return aiStr, nil
}

func (s *Service) PreviewStreamAgentSystemPromptV1(ctx context.Context, writerResponse gin.ResponseWriter,
	outputChan chan *models.ChatCompletionStreamResponse,
	errChan chan error, doneChan chan bool) {
	needFakeResponse := true
	for {
		select {
		case <-ctx.Done():
			return
		case <-doneChan:
			if _, err := writerResponse.Write(serializers.HttpEventStreamResponse{Data: serializers.DoneResponseStreamData}.ToOutPut()); err != nil {
				return
			}
			writerResponse.Flush()
			return
		case <-time.Tick(time.Minute * 20):
			if _, err := writerResponse.Write(serializers.HttpEventStreamResponse{Data: serializers.TimeoutResponseStreamData}.ToOutPut()); err != nil {
				return
			}
			writerResponse.Flush()
			return
		case err := <-errChan:
			data, _ := json.Marshal(models.ChatCompletionStreamResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest})
			if _, err := writerResponse.Write(serializers.HttpEventStreamResponse{Data: data}.ToOutPut()); err != nil {
				return
			}
			writerResponse.Flush()
			logger.Error("stream_retrieve_knowledge", "return with err", zap.Any("err", err))
			return
		case output := <-outputChan:
			if output.Code != http.StatusProcessing {
				needFakeResponse = false
			}
			data, _ := json.Marshal(output)
			if _, err := writerResponse.Write(serializers.HttpEventStreamResponse{Data: data}.ToOutPut()); err != nil {
				return
			}
			writerResponse.Flush()
		default:
			if needFakeResponse {
				if _, err := writerResponse.Write(serializers.HttpEventStreamResponse{Data: serializers.FakeResponseStreamData}.ToOutPut()); err != nil {
					return
				}
				writerResponse.Flush()
				time.Sleep(2 * time.Second)
			}
		}
	}
}

func (s *Service) ProcessStreamAgentSystemPromptV1(ctx context.Context,
	req *serializers.PreviewRequest, outputChan chan *models.ChatCompletionStreamResponse,
	errChan chan error, doneChan chan bool) {
	var agentInfo *models.AgentInfo
	baseModel := "NousResearch/Hermes-3-Llama-3.1-70B-FP8"
	if req.AgentID != nil {
		agentInfo, _ = s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx), *req.AgentID, map[string][]interface{}{}, false)
	}
	var llmMessage []openai2.ChatCompletionMessage
	err := json.Unmarshal([]byte(req.Messages), &llmMessage)
	if err != nil {
		errChan <- errs.NewError(errors.New("invalid message request"))
		return
	}

	systemContent := openai.GetSystemPromptFromLLMMessage(llmMessage)
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
		if agentInfo != nil && agentInfo.AgentBaseModel != "" {
			baseModel = agentInfo.AgentBaseModel
		}
	}
	if req.ModelName != nil && len(*req.ModelName) > 0 {
		baseModel = *req.ModelName
	}

	if s.conf.KnowledgeBaseConfig.EnableSimulation && agentInfo != nil {
		// get last knowledge base of agent
		var knowledgeBaseUse *models.KnowledgeBase
		agentInfoKnowledgeBase, _ := s.dao.FirstAgentInfoKnowledgeBaseByAgentInfoID(
			daos.GetDBMainCtx(ctx),
			agentInfo.ID,
			map[string][]interface{}{
				"KnowledgeBase": {}, // must preload
			},
			[]string{"id desc"},
		)
		if agentInfoKnowledgeBase != nil && agentInfoKnowledgeBase.KnowledgeBase != nil {
			knowledgeBaseUse = agentInfoKnowledgeBase.KnowledgeBase
		}

		if req.KbId != nil {
			knowledgeBaseFromQuery, _ := s.dao.FirstKnowledgeBase(daos.GetDBMainCtx(ctx), map[string][]interface{}{
				"kb_id = ?": {*req.KbId},
			}, map[string][]interface{}{}, []string{"id desc"}, false)
			if knowledgeBaseFromQuery != nil {
				knowledgeBaseUse = knowledgeBaseFromQuery
			}
		}
		if knowledgeBaseUse != nil {
			s.StreamRetrieveKnowledge(ctx, baseModel, llmMessage, []*models.KnowledgeBase{
				knowledgeBaseUse,
			}, nil, nil, outputChan, errChan, doneChan)
			return
		}
	}

	llmMessage = openai.UpdateSystemPromptInLLMRequest(llmMessage, systemContent)
	messageCallLLM, _ := json.Marshal(&llmMessage)
	s.openais["Agent"].CallStreamDirectlyEternalLLM(ctx, string(messageCallLLM), baseModel, url, map[string]interface{}{
		"top_p":      0.01,
		"max_tokens": 4096,
	}, outputChan, errChan, doneChan)
}

func (s *Service) AgentChatSupport(ctx context.Context, msg string) (string, error) {
	aiStr, err := s.openais["Lama"].ChatMessage(msg)
	if err != nil {
		return "", errs.NewError(err)
	}
	return aiStr, nil
}

func (s *Service) GetExtractDataFromPost(content string) (string, error) {
	extractData := ""
	link, isTwitterPost := helpers.ExtractLinks(content)
	if link != "" {
		if isTwitterPost {
			twiterPostIDArry := strings.Split(link, "/")
			if len(twiterPostIDArry) > 0 {
				twitterPostID := twiterPostIDArry[len(twiterPostIDArry)-1]
				fmt.Println(twitterPostID)
				tweetDetail, err := s.rapid.GetTweetDetailByID(twitterPostID)
				if err != nil {
					return extractData, nil
				}
				if tweetDetail != nil {
					extractData = tweetDetail.FullText
				}
			}
		} else {
			webContent := helpers.ContentHtmlByUrl(link)
			if webContent == "" {
				webContent = helpers.RodContentHtmlByUrl(link)
			}
			webContent, err := s.blockchainUtils.CleanHtml(webContent)
			if err != nil {
				return extractData, nil
			}

			if webContent != "" {
				summary, err := s.openais["Lama"].SummaryWebContent(webContent)
				if err != nil {
					return extractData, nil
				}
				extractData = summary
			}
		}
	}
	return extractData, nil
}

func (s *Service) CheckAgentIsReadyToRunTwinTraining(agentInfo *models.AgentInfo) (float64, bool, error) {
	needBalance := 0.0
	switch agentInfo.NetworkID {
	case models.SHARDAI_CHAIN_ID:
		needBalance = 5999.9
	case models.ETHEREUM_CHAIN_ID:
		needBalance = 299.9
	case models.SOLANA_CHAIN_ID:
		needBalance = 174.9
	default:
		needBalance = 0.99
	}

	if agentInfo.TwinTwitterUsernames != "" {
		arr := strings.Split(agentInfo.TwinTwitterUsernames, ",")
		// 300 EAI for each twitter username
		needBalance += float64(len(arr)) * 300
	}

	agentEaiBalance, _ := agentInfo.EaiBalance.Float64()
	if agentEaiBalance >= needBalance {
		return needBalance, true, nil
	}

	return needBalance, false, nil
}

func (s *Service) JobAgentTwinTrain(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobAgentTwinTrain",
		func() error {
			// Count pending twin training
			twinTrainingAgents, err := s.dao.FindAgentInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twin_twitter_usernames != '' and twin_status = ?": {models.TwinStatusRunning},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				999999,
			)
			if err == nil && len(twinTrainingAgents) > 5 {
				logger.Info("twin_training_jobs", "twin training is running maximum ===> skip", zap.Any("len_training_agents", len(twinTrainingAgents)))
				return nil
			}

			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					//"agent_id != ''":                                   {},
					//"agent_contract_id = ?":                            {""},
					//"agent_nft_minted = ?":                             {false},
					`twin_twitter_usernames != '' and twin_status = ?`: {models.TwinStatusPending},
					"scan_enabled = ?": {true},
				},
				map[string][]interface{}{},
				[]string{
					"id asc",
				},
				0,
				999999,
			)

			if err != nil {
				return errs.NewError(err)
			}

			var retErr error
			wg := sync.WaitGroup{}
			for _, agent := range agents {
				wg.Add(1)
				go func(_agent *models.AgentInfo) {
					defer wg.Done()
					err = s.AgentTwinTrain(ctx, _agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
				}(agent)
			}
			wg.Wait()
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateUpdateAgentSnapshotMission(ctx context.Context, agentID string, authHeader string, req []*serializers.AgentSnapshotMissionInfo) (*models.AgentInfo, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo != nil {
				defataulUserPromp := ""
				listID := []uint{}

				for _, item := range req {
					if item.ID > 0 {
						listID = append(listID, item.ID)
					}
				}

				listTestToolSet := strings.Split(s.conf.ListTestToolSet, ",")
				if len(listID) > 0 {
					err = tx.Where("agent_info_id = ? and id not in (?) and (mission_store_id = 0 or mission_store_id is NULL) and tool_set not in (?)", agentInfo.ID, listID, listTestToolSet).Delete(&models.AgentSnapshotMission{}).Error
					if err != nil {
						return errs.NewError(err)
					}
				} else {
					err = tx.Where("agent_info_id = ? and (mission_store_id = 0 or mission_store_id is NULL) and tool_set not in (?)", agentInfo.ID, listTestToolSet).Delete(&models.AgentSnapshotMission{}).Error
					if err != nil {
						return errs.NewError(err)
					}
				}

				for _, item := range req {
					if defataulUserPromp == "" {
						defataulUserPromp = item.UserPrompt
					}

					mission := &models.AgentSnapshotMission{}
					if item.ID > 0 {
						mission, err = s.dao.FirstAgentSnapshotMissionByID(tx, item.ID,
							map[string][]interface{}{}, false,
						)
						if err != nil {
							return errs.NewError(err)
						}
					}
					mission.NetworkID = agentInfo.NetworkID
					mission.AgentInfoID = agentInfo.ID
					mission.UserPrompt = item.UserPrompt
					mission.IntervalSec = item.Interval
					mission.ToolSet = item.ToolSet
					mission.Enabled = true
					mission.ReplyEnabled = true
					mission.AgentType = item.AgentType
					mission.UserTwitterIds = item.UserTwitterIDs
					mission.Tokens = item.Tokens
					mission.AgentBaseModel = item.AgentBaseModel
					mission.Topics = item.Topics
					mission.IsBingSearch = item.IsBingSearch
					mission.IsTwitterSearch = item.IsTwitterSearch
					mission.RewardAmount = item.RewardAmount
					mission.RewardUser = item.RewardUser
					mission.MinTokenHolding = item.MinTokenHolding
					mission.LookupInterval = item.LookupInterval
					//farcaster
					if mission.ToolSet == models.ToolsetTypePostFarcaster {
						toolList := fmt.Sprintf(s.conf.ToolLists.FarcasterPost, agentInfo.FarcasterID, authHeader, agentInfo.AgentID)
						mission.ToolList = toolList
					} else if mission.ToolSet == models.ToolsetTypeReplyMentionsFarcaster {
						toolList := fmt.Sprintf(s.conf.ToolLists.FarcasterReply, agentInfo.FarcasterID, authHeader, agentInfo.AgentID)
						mission.ToolList = toolList
					} else if mission.ToolSet == models.ToolsetTypeTradeNews {
						toolList := fmt.Sprintf(s.conf.ToolLists.TradeNews, s.conf.InternalApiKey, agentInfo.ID)
						mission.ToolList = toolList
						mission.UserPrompt = "Analyze the coin price fluctuations in the past 24 hours, suggest which coin to buy or sell and post it on twitter"
					} else if mission.ToolSet == models.ToolsetTypeTradeAnalytics || mission.ToolSet == models.ToolsetTypeTradeAnalyticsOnTwitter {
						toolList := s.conf.ToolLists.TradeAnalytic
						if item.Tokens == "" {
							return errs.NewError(errs.ErrTokenNotFound)
						}
						mission.UserPrompt = fmt.Sprintf(`Conduct a technical analysis of $%s price data. Based on your findings, provide a recommended buy price and sell price to maximize potential returns.`, item.Tokens)
						toolList = strings.ReplaceAll(toolList, "{api_key}", s.conf.InternalApiKey)
						toolList = strings.ReplaceAll(toolList, "{token_symbol}", item.Tokens)

						mission.ToolList = toolList
					} else if mission.ToolSet == models.ToolsetTypeLuckyMoneys {
						if item.RewardAmount.Cmp(big.NewFloat(0)) <= 0 || item.RewardUser <= 0 {
							return errs.NewError(errs.ErrBadRequest)
						}
					} else if item.AgentStoreMissionID > 0 {
						agentStoreMission, err := s.dao.FirstAgentStoreMissionByID(tx, item.AgentStoreMissionID, map[string][]interface{}{}, false)
						if err != nil {
							return errs.NewError(err)
						}
						if agentStoreMission == nil {
							return errs.NewError(errs.ErrBadRequest)
						}
						agentStoreInstall, err := s.dao.FirstAgentStoreInstall(
							tx,
							map[string][]interface{}{
								"agent_store_id = ?": {agentStoreMission.AgentStoreID},
								"agent_info_id = ?":  {mission.AgentInfoID},
							},
							map[string][]interface{}{},
							[]string{"id desc"},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agentStoreInstall == nil {
							return errs.NewError(errs.ErrBadRequest)
						}
						mission.AgentStoreMissionID = agentStoreMission.ID
						mission.AgentStoreID = agentStoreMission.AgentStoreID
						mission.ReactMaxSteps = 5
						mission.ToolSet = "mission_store"
					} else if item.ToolList != "" {
						mission.ToolList = item.ToolList
					}
					if mission.ToolList != "" {
						mission.ReactMaxSteps = 5
					}
					//
					err = s.dao.Save(tx, mission)
					if err != nil {
						return errs.NewError(err)
					}
				}
				if defataulUserPromp != "" {
					updateAgentFields := map[string]interface{}{
						"user_prompt": defataulUserPromp,
					}

					err = tx.Model(agentInfo).Updates(updateAgentFields).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}

			return nil
		},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}
	return s.GetAgentInfoDetailByAgentID(ctx, agentID)
}

func (s *Service) DeleteAgentSnapshotMission(ctx context.Context, missionID uint, userAddress string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			missionInfo, err := s.dao.FirstAgentSnapshotMissionByID(tx,
				missionID,
				map[string][]interface{}{
					"AgentInfo": {},
				}, false,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if missionInfo != nil && missionInfo.AgentInfo != nil && strings.EqualFold(missionInfo.AgentInfo.Creator, userAddress) {
				err = tx.Delete(missionInfo).Error
				if err != nil {
					return errs.NewError(err)
				}
			}

			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) FollowListDefaultTwitters(ctx context.Context, agentID uint) error {
	listFollow, err := s.dao.GetListTwitterDefaultFollow(daos.GetDBMainCtx(ctx))
	if err != nil {
		return errs.NewError(err)
	}
	for _, v := range listFollow {
		agent, err := s.dao.FirstAgentInfoByID(
			daos.GetDBMainCtx(ctx),
			agentID,
			map[string][]interface{}{
				"TwitterInfo": {},
			},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if agent != nil {
			helpers.TwitterFollowUserCreate(agent.TwitterInfo.AccessToken, agent.TwitterID, v)
			time.Sleep(20 * time.Second)
		}
	}
	return nil
}

func (s *Service) UpdateTwinStatus(ctx context.Context, req *serializers.UpdateTwinStatusRequest) (*models.AgentInfo, error) {
	agentIDInt, err := strconv.Atoi(req.AgentID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agentInfoEntity, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"id = ?": {agentIDInt},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	agentInfoEntity.TwinStatus = models.TwinStatus(req.TwinStatus)
	agentInfoEntity.KnowledgeBaseID = req.KnowledgeBaseID
	agentInfoEntity.SystemPrompt = req.SystemPrompt
	agentInfoEntity.TwinTrainingProgress = req.TwinTrainingProgress
	if agentInfoEntity.TwinTrainingProgress > 100 {
		agentInfoEntity.TwinTrainingProgress = 100
	}
	if agentInfoEntity.TwinTrainingProgress < 0 {
		agentInfoEntity.TwinTrainingProgress = 0
	}

	if req.TwinStatus == string(models.TwinStatusDoneError) || req.TwinStatus == string(models.TwinStatusDoneSuccess) {
		endAt := time.Now().UTC()
		agentInfoEntity.TwinTrainingMessage = req.TwinTrainingMessage
		agentInfoEntity.TwinEndTrainingAt = &endAt
	}

	err = s.dao.Save(daos.GetDBMainCtx(ctx), agentInfoEntity)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if req.TwinStatus == string(models.TwinStatusDoneError) {
		eventId := fmt.Sprintf("twin_train_refund_%d", agentInfoEntity.ID)
		checkRefunded, err := s.dao.FirstAgentEaiTopup(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"event_id = ?": {eventId},
			},
			map[string][]interface{}{},
			[]string{},
		)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewError(err)
		}
		if checkRefunded != nil {
			return agentInfoEntity, nil
		}

		// increase balance and add refund history
		arr := strings.Split(agentInfoEntity.TwinTwitterUsernames, ",")
		twinFee := numeric.NewFloatFromString(fmt.Sprintf("%v", float64(len(arr))*300))
		_ = daos.WithTransaction(daos.GetDBMainCtx(ctx), func(dbTx *gorm.DB) error {
			err = daos.GetDBMainCtx(ctx).
				Model(agentInfoEntity).
				Updates(
					map[string]interface{}{
						"eai_balance": gorm.Expr("eai_balance + ?", numeric.NewBigFloatFromFloat(twinFee)),
					},
				).
				Error
			if err != nil {
				return errs.NewError(err)
			}

			return s.dao.Create(
				daos.GetDBMainCtx(ctx),
				&models.AgentEaiTopup{
					NetworkID:      agentInfoEntity.NetworkID,
					EventId:        fmt.Sprintf("twin_train_refund_%d", agentInfoEntity.ID),
					AgentInfoID:    agentInfoEntity.ID,
					Type:           models.AgentEaiTopupTypeRefundTrainFail,
					Amount:         numeric.NewBigFloatFromFloat(twinFee),
					Status:         models.AgentEaiTopupStatusDone,
					DepositAddress: agentInfoEntity.ETHAddress,
					ToAddress:      agentInfoEntity.ETHAddress,
					Toolset:        "twin_train_refund",
				},
			)
		})
	}

	return agentInfoEntity, nil
}

func (s *Service) UnlinkAgentTwitterInfo(ctx context.Context, agentID string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo != nil {
				updateFields := map[string]interface{}{
					"twitter_info_id":  0,
					"twitter_id":       "",
					"twitter_username": "",
				}

				err := tx.Model(agentInfo).Updates(
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
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) AgentChats(ctx context.Context, agentID string, messages serializers.AgentChatMessageReq) (*openai.ChatResponse, error) {
	var aiStr *openai.ChatResponse
	agentInfo, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"agent_id = ?": {agentID},
		},
		map[string][]interface{}{},
		[]string{},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	if agentInfo != nil {
		aiStr, err = s.openais["Agent"].AgentChats(agentInfo.GetSystemPrompt(), s.conf.AgentOffchainChatUrl, messages)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return aiStr, nil
}

func (s *Service) PauseAgent(ctx context.Context, agentID string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo != nil {
				isPause := !agentInfo.ReplyEnabled

				updateFields := map[string]interface{}{
					"reply_enabled": isPause,
				}

				err := tx.Model(agentInfo).Updates(
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
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) UpdateAgentExternalInfo(ctx context.Context, agentID string, req *serializers.AgentExternalInfoReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}
			if agentInfo != nil {
				externalInfo, err := s.dao.FirstAgentExternalInfo(tx,
					map[string][]interface{}{
						"agent_info_id = ?": {agentInfo.ID},
						"network_id = ?":    {agentInfo.NetworkID},
						"type = ?":          {req.Type},
					},
					map[string][]interface{}{},
					[]string{})
				if err != nil {
					return errs.NewError(err)
				}
				if externalInfo != nil {
					externalInfo.ExternalID = req.ExternalID
					externalInfo.ExternalUsername = req.ExternalUsername
					externalInfo.ExternalName = req.ExternalName
					err = s.dao.Save(tx, externalInfo)
					if err != nil {
						return errs.NewError(err)
					}
				} else {
					err = s.dao.Create(tx, &models.AgentExternalInfo{
						NetworkID:        agentInfo.NetworkID,
						Type:             models.ExternalAgentType(req.Type),
						AgentInfoID:      agentInfo.ID,
						ExternalID:       req.ExternalID,
						ExternalUsername: req.ExternalUsername,
						ExternalName:     req.ExternalName,
					})
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) CreateAgentInfoInstallCode(ctx context.Context, userAddress string, agentInfoID uint) (*models.AgentInfoInstall, error) {
	if agentInfoID > 0 {
		agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx), agentInfoID, map[string][]interface{}{}, true)
		if err != nil {
			return nil, errs.NewError(err)
		}
		if agentInfo == nil {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
	}
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agentInfoInstall, err := s.dao.FirstAgentInfoInstall(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"user_id = ?":       {user.ID},
			"agent_info_id = ?": {agentInfoID},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentInfoInstall == nil {
		code := helpers.RandomStringWithLength(64)
		params := map[string]string{"code": code}
		paramstr, _ := json.Marshal(params)
		agentInfoInstall = &models.AgentInfoInstall{
			Code:           code,
			AgentInfoID:    agentInfoID,
			Status:         models.AgenInfoInstallStatusNew,
			UserID:         user.ID,
			CallbackParams: string(paramstr),
		}
		err = s.dao.Create(daos.GetDBMainCtx(ctx), agentInfoInstall)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return agentInfoInstall, nil
}

func (s *Service) GetAgentInfoInstall(ctx context.Context, code string) (*models.AgentInfoInstall, error) {
	res, err := s.dao.FirstAgentInfoInstall(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"code = ?": {code},
		},
		map[string][]interface{}{
			"User": {},
		},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) CheckNameExist(ctx context.Context, networkID uint64, name string) (bool, error) {
	obj, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?": {networkID},
			"agent_name = ?": {name},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	if obj != nil {
		return true, nil
	}
	return false, nil
}
