package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// func (s *Service) InfraTwitterAppAuthenInstall(ctx context.Context, installCode string, installUri string) (string, error) {
// 	err := func() error {
// 		if installCode == "" {
// 			return errs.NewError(errs.ErrBadRequest)
// 		}
// 		infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
// 			daos.GetDBMainCtx(ctx),
// 			map[string][]any{
// 				"install_code = ?": {installCode},
// 			},
// 			map[string][]any{}, []string{},
// 		)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}
// 		if infraTwitterApp == nil {
// 			var res struct {
// 				Result string `json:"result"`
// 			}
// 			err = helpers.CurlURL(s.conf.InfraTwitterApp.InfraAuthUri+"?code="+installCode, http.MethodGet, nil, nil, &res)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 			if res.Result == "" {
// 				return errs.NewError(errs.ErrBadRequest)
// 			}
// 			infraTwitterApp = &models.InfraTwitterApp{
// 				Address:     res.Result,
// 				InstallCode: installCode,
// 			}
// 		}
// 		err = s.dao.Save(daos.GetDBMainCtx(ctx), infraTwitterApp)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}
// 		return nil
// 	}()
// 	if err != nil {
// 		return helpers.BuildUri(
// 			installUri,
// 			map[string]string{
// 				"error": err.Error(),
// 			},
// 		), nil
// 	}
// 	redirectUri := helpers.BuildUri(
// 		s.conf.InfraTwitterApp.RedirectUri,
// 		map[string]string{
// 			"install_code": installCode,
// 			"install_uri":  installUri,
// 		},
// 	)
// 	return helpers.BuildUri(
// 		"https://twitter.com/i/oauth2/authorize",
// 		map[string]string{
// 			"client_id":             s.conf.InfraTwitterApp.OauthClientId,
// 			"state":                 "state",
// 			"response_type":         "code",
// 			"code_challenge":        "challenge",
// 			"code_challenge_method": "plain",
// 			"scope":                 "offline.access+tweet.read+tweet.write+users.read+follows.write+like.write+like.read+users.read",
// 			"redirect_uri":          redirectUri,
// 		},
// 	), nil
// }

// func (s *Service) InfraTwitterAppAuthenCallback(ctx context.Context, installCode string, installUri string, code string) (string, error) {
// 	if installCode == "" || code == "" {
// 		return "", errs.NewError(errs.ErrBadRequest)
// 	}
// 	infraTwitterApp, err := func() (*models.InfraTwitterApp, error) {
// 		redirectUri := helpers.BuildUri(
// 			s.conf.InfraTwitterApp.RedirectUri,
// 			map[string]string{
// 				"install_code": installCode,
// 				"install_uri":  installUri,
// 			},
// 		)
// 		respOauth, err := s.twitterAPI.TwitterOauthCallbackForSampleApp(
// 			s.conf.InfraTwitterApp.OauthClientId, s.conf.InfraTwitterApp.OauthClientSecret, code, redirectUri)
// 		if err != nil {
// 			return nil, errs.NewError(err)
// 		}
// 		if respOauth != nil && respOauth.AccessToken != "" {
// 			twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
// 			if err != nil {
// 				return nil, errs.NewError(err)
// 			}
// 			twitterInfo, err := s.dao.FirstTwitterInfo(
// 				daos.GetDBMainCtx(ctx),
// 				map[string][]any{
// 					"twitter_id = ?": {twitterUser.ID},
// 				},
// 				map[string][]any{},
// 				false,
// 			)
// 			if err != nil {
// 				return nil, errs.NewError(err)
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
// 			twitterInfo.OauthClientId = s.conf.InfraTwitterApp.OauthClientId
// 			twitterInfo.OauthClientSecret = s.conf.InfraTwitterApp.OauthClientSecret
// 			twitterInfo.Description = twitterUser.Description
// 			twitterInfo.RefreshError = "OK"
// 			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
// 			twitterInfo.ExpiredAt = &expiredAt
// 			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
// 			if err != nil {
// 				return nil, errs.NewError(err)
// 			}
// 			infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
// 				daos.GetDBMainCtx(ctx),
// 				map[string][]any{
// 					"install_code = ?": {installCode},
// 				},
// 				map[string][]any{}, []string{},
// 			)
// 			if err != nil {
// 				return nil, errs.NewError(err)
// 			}
// 			if infraTwitterApp == nil {
// 				return nil, errs.NewError(errs.ErrBadRequest)
// 			}
// 			infraTwitterApp.TwitterInfoID = twitterInfo.ID
// 			infraTwitterApp.TwitterInfo = twitterInfo
// 			err = s.dao.Save(daos.GetDBMainCtx(ctx), infraTwitterApp)
// 			if err != nil {
// 				return nil, errs.NewError(err)
// 			}
// 			return infraTwitterApp, nil
// 		}
// 		return nil, errs.NewError(errs.ErrBadRequest)
// 	}()
// 	if err != nil {
// 		return helpers.BuildUri(
// 			installUri,
// 			map[string]string{
// 				"install_code": installCode,
// 				"error":        err.Error(),
// 			},
// 		), nil
// 	}
// 	params := map[string]string{
// 		"address":          infraTwitterApp.Address,
// 		"twitter_id":       infraTwitterApp.TwitterInfo.TwitterID,
// 		"twitter_username": infraTwitterApp.TwitterInfo.TwitterUsername,
// 		"twitter_name":     infraTwitterApp.TwitterInfo.TwitterName,
// 	}
// 	returnData := base64.StdEncoding.EncodeToString([]byte(helpers.ConvertJsonString(params)))
// 	return helpers.BuildUri(
// 		installUri,
// 		map[string]string{
// 			"install_code": installCode,
// 			"return_data":  returnData,
// 		},
// 	), nil
// }

func (s *Service) InfraTwitterAppAuthenInstall(ctx context.Context, userAddress string) (string, error) {
	err := func() error {
		infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
			daos.GetDBMainCtx(ctx),
			map[string][]any{
				"address = ?": {strings.ToLower(userAddress)},
			},
			map[string][]any{}, []string{},
		)
		if err != nil {
			return errs.NewError(err)
		}
		if infraTwitterApp == nil {
			ethAddress, err := s.CreateETHAddress(ctx)
			if err != nil {
				return errs.NewError(err)
			}
			infraTwitterApp = &models.InfraTwitterApp{
				Address:    strings.ToLower(userAddress),
				ETHAddress: strings.ToLower(ethAddress),
			}
			err = s.dao.Create(daos.GetDBMainCtx(ctx), infraTwitterApp)
			if err != nil {
				return errs.NewError(err)
			}
		}

		return nil
	}()

	if err != nil {
		return "", errs.NewError(err)
	}

	redirectUri := helpers.BuildUri(
		s.conf.InfraTwitterApp.RedirectUri,
		map[string]string{
			"address": userAddress,
		},
	)

	return helpers.BuildUri(
		"https://twitter.com/i/oauth2/authorize",
		map[string]string{
			"redirect_uri":          redirectUri,
			"client_id":             s.conf.InfraTwitterApp.OauthClientId,
			"state":                 "state",
			"response_type":         "code",
			"code_challenge":        "challenge",
			"code_challenge_method": "plain",
			"scope":                 "offline.access+tweet.read+tweet.write+users.read+follows.write+like.write+like.read",
		},
	), nil
}

func (s *Service) InfraTwitterAppAuthenCallback(ctx context.Context, address string, code string) (string, error) {
	_, err := func() (*models.InfraTwitterApp, error) {
		redirectUri := helpers.BuildUri(
			s.conf.InfraTwitterApp.RedirectUri,
			map[string]string{
				"address": address,
			},
		)
		respOauth, err := s.twitterAPI.TwitterOauthCallbackForSampleApp(
			s.conf.InfraTwitterApp.OauthClientId, s.conf.InfraTwitterApp.OauthClientSecret, code, redirectUri)
		if err != nil {
			return nil, errs.NewError(err)
		}
		if respOauth != nil && respOauth.AccessToken != "" {
			twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
			if err != nil {
				return nil, errs.NewError(err)
			}
			twitterInfo, err := s.dao.FirstTwitterInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]any{},
				false,
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
			twitterInfo.OauthClientId = s.conf.InfraTwitterApp.OauthClientId
			twitterInfo.OauthClientSecret = s.conf.InfraTwitterApp.OauthClientSecret
			twitterInfo.Description = twitterUser.Description
			twitterInfo.RefreshError = "OK"
			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
			twitterInfo.ExpiredAt = &expiredAt
			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
			if err != nil {
				return nil, errs.NewError(err)
			}
			infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
				daos.GetDBMainCtx(ctx),
				map[string][]any{
					"address = ?": {strings.ToLower(address)},
				},
				map[string][]any{}, []string{},
			)
			if err != nil {
				return nil, errs.NewError(err)
			}
			if infraTwitterApp == nil {
				return nil, errs.NewError(errs.ErrBadRequest)
			}
			if infraTwitterApp.ETHAddress == "" {
				ethAddress, err := s.CreateETHAddress(ctx)
				if err != nil {
					return nil, errs.NewError(err)
				}
				infraTwitterApp.ETHAddress = strings.ToLower(ethAddress)
			}
			infraTwitterApp.TwitterInfoID = twitterInfo.ID
			infraTwitterApp.TwitterInfo = twitterInfo
			err = s.dao.Save(daos.GetDBMainCtx(ctx), infraTwitterApp)
			if err != nil {
				return nil, errs.NewError(err)
			}
			return infraTwitterApp, nil
		}
		return nil, errs.NewError(errs.ErrBadRequest)
	}()

	if err != nil {
		return "", errs.NewError(err)
	}

	// if err != nil {
	// 	return helpers.BuildUri(
	// 		installUri,
	// 		map[string]string{
	// 			"address": address,
	// 			"error":   err.Error(),
	// 		},
	// 	), nil
	// }
	// params := map[string]string{
	// 	"address":          infraTwitterApp.Address,
	// 	"twitter_id":       infraTwitterApp.TwitterInfo.TwitterID,
	// 	"twitter_username": infraTwitterApp.TwitterInfo.TwitterUsername,
	// 	"twitter_name":     infraTwitterApp.TwitterInfo.TwitterName,
	// }
	// returnData := base64.StdEncoding.EncodeToString([]byte(helpers.ConvertJsonString(params)))
	return "", nil
}

func (s *Service) UtilityPostTwitter(ctx context.Context, userAddress string, req *serializers.AgentUtilityTwitterReq) (*serializers.AgentUtilityTwitterResp, error) {
	resp := &serializers.AgentUtilityTwitterResp{}
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
				tx,
				map[string][]any{
					"address = ?": {strings.ToLower(userAddress)},
				},
				map[string][]any{
					"TwitterInfo": {},
				},
				[]string{},
			)

			if err != nil {
				resp.AuthUrl, _ = s.InfraTwitterAppAuthenInstall(ctx, userAddress)
				return errs.NewError(errs.ErrAgentUtilityNotFound)
			}

			if infraTwitterApp == nil || (infraTwitterApp != nil && infraTwitterApp.TwitterInfo == nil) ||
				(infraTwitterApp != nil && infraTwitterApp.TwitterInfo != nil && infraTwitterApp.TwitterInfo.RefreshError != "OK") {
				resp.AuthUrl, _ = s.InfraTwitterAppAuthenInstall(ctx, userAddress)
				// resp.Message = errs.ErrAgentUtilityNotAuthen.Message
				return errs.NewError(errs.ErrAgentUtilityNotAuthen)
			}

			if s.conf.InfraTwitterApp.Fee > 0 {
				if infraTwitterApp != nil && infraTwitterApp.RemainRequest <= 0 {
					resp.Message = strings.ReplaceAll(errs.ErrAgentUtilityInvalidBalance.Message, "{address}", s.conf.InfraTwitterApp.AgentAddress)
					return errs.NewError(errs.ErrAgentUtilityInvalidBalance)
				}
			}

			if infraTwitterApp != nil && infraTwitterApp.TwitterInfo != nil && infraTwitterApp.TwitterInfo.RefreshError == "OK" {
				if s.conf.InfraTwitterApp.Fee > 0 {
					eventId := uuid.New().String()
					feePerRequest := numeric.NewBigFloatFromString(fmt.Sprintf(`%d`, s.conf.InfraTwitterApp.Fee))
					topupTx := &models.InfraTwitterTopupTx{
						InfraTwitterAppID: infraTwitterApp.ID,
						NetworkID:         models.BASE_CHAIN_ID,
						EventId:           eventId,
						Type:              models.AgentEaiTopupTypeSpent,
						DepositAddress:    infraTwitterApp.Address,
						ToAddress:         infraTwitterApp.Address,
						TxHash:            eventId,
						Amount:            feePerRequest,
					}
					err = s.dao.Save(tx, topupTx)
					if err != nil {
						return errs.NewError(err)
					}
					err = tx.Model(infraTwitterApp).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance - ?", feePerRequest)).
						Error
					if err != nil {
						return errs.NewError(errs.ErrAgentUtilitySystemError)
					}
				}
				tweetId, err := helpers.PostTweetByToken(infraTwitterApp.TwitterInfo.AccessToken, req.Content, "")
				if err != nil {
					return errs.NewError(errs.ErrAgentUtilityPostTweetFailed)
				}
				resp.Message = fmt.Sprintf(`https://x.com/%s/status/%s`, infraTwitterApp.TwitterInfo.TwitterUsername, tweetId)
				return nil
			}

			return nil
		},
	)
	if err != nil {
		if resp.Message == "" {
			resp.Message = err.Error()
		}
		return resp, errs.NewError(err)
	}
	return resp, nil
}

// func (s *Service) UtilityTwitterVerifyDeposit(ctx context.Context, userAddress, txHash string) (bool, error) {
// 	err := daos.WithTransaction(
// 		daos.GetDBMainCtx(ctx),
// 		func(tx *gorm.DB) error {
// 			eventResp, err := s.GetEthereumClient(ctx, models.BASE_CHAIN_ID).Erc20EventsByTransaction(txHash)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}

// 			if eventResp != nil && len(eventResp.Transfer) > 0 {
// 				txEvent := eventResp.Transfer[0]
// 				if !strings.EqualFold(s.conf.InfraTwitterApp.AgentAddress, txEvent.To) {
// 					return errs.NewError(err)
// 				}

// 				eventID := fmt.Sprintf(`%s_%d`, strings.ToLower(txHash), txEvent.TxIndex)
// 				topupTx, err := s.dao.FirstInfraTwitterTopupTx(
// 					tx,
// 					map[string][]any{
// 						"event_id = ?": {eventID},
// 					},
// 					map[string][]any{},
// 					[]string{},
// 				)

// 				if err != nil {
// 					return errs.NewError(err)
// 				}

// 				if topupTx == nil {
// 					infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
// 						tx,
// 						map[string][]any{
// 							"eth_address = ?": {strings.ToLower(txEvent.From)},
// 						},
// 						map[string][]any{},
// 						[]string{},
// 					)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}
// 					if infraTwitterApp == nil {
// 						return errs.NewError(errs.ErrAgentUtilityNotAuthen)
// 					}
// 					fBalance := models.ConvertWeiToBigFloat(txEvent.Value, 18)
// 					topupTx := &models.InfraTwitterTopupTx{
// 						InfraTwitterAppID: infraTwitterApp.ID,
// 						NetworkID:         models.BASE_CHAIN_ID,
// 						EventId:           eventID,
// 						Type:              models.AgentEaiTopupTypeDeposit,
// 						DepositAddress:    txEvent.From,
// 						ToAddress:         txEvent.To,
// 						TxHash:            txHash,
// 						Amount:            numeric.NewBigFloatFromFloat(fBalance),
// 					}
// 					err = s.dao.Save(tx, topupTx)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}
// 					tmpBlance, _ := fBalance.Float64()
// 					newRequest := int(math.Round(tmpBlance))
// 					err = tx.Model(infraTwitterApp).
// 						UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", fBalance)).
// 						UpdateColumn("total_request", gorm.Expr("total_request + ?", newRequest)).
// 						UpdateColumn("remain_request", gorm.Expr("remain_request + ?", newRequest)).Error
// 					if err != nil {
// 						return errs.NewError(errs.ErrBadRequest)
// 					}
// 				}
// 			}
// 			return nil
// 		},
// 	)
// 	if err != nil {
// 		return false, errs.NewError(err)
// 	}
// 	return true, nil
// }

func (s *Service) UtilityTwitterHandleDeposit(tx *gorm.DB, networkID uint64, event *ethapi.Erc20TokenTransferEventResp) error {
	eaiContractAddress := s.conf.GetConfigKeyString(networkID, "eai_contract_address")
	if !strings.EqualFold(eaiContractAddress, event.ContractAddress) {
		return nil
	}
	infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
		tx,
		map[string][]any{
			"eth_address = ?": {strings.ToLower(event.To)},
		},
		map[string][]any{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if infraTwitterApp != nil {
		eventId := fmt.Sprintf(`%d_%s_%d`, networkID, strings.ToLower(event.TxHash), event.Index)
		topupTx, err := s.dao.FirstInfraTwitterTopupTx(
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
		if topupTx == nil {
			fBalance := models.ConvertWeiToBigFloat(event.Value, 18)
			topupTx := &models.InfraTwitterTopupTx{
				InfraTwitterAppID: infraTwitterApp.ID,
				NetworkID:         models.BASE_CHAIN_ID,
				EventId:           eventId,
				Type:              models.AgentEaiTopupTypeDeposit,
				DepositAddress:    event.From,
				ToAddress:         event.To,
				TxHash:            event.TxHash,
				Amount:            numeric.NewBigFloatFromFloat(fBalance),
			}
			err = s.dao.Save(tx, topupTx)
			if err != nil {
				return errs.NewError(err)
			}
			err = tx.Model(infraTwitterApp).
				Updates(map[string]any{
					"eai_balance": gorm.Expr("eai_balance + ?", fBalance),
				}).
				Error
			if err != nil {
				return errs.NewError(errs.ErrBadRequest)
			}
		}
	}
	return nil
}

func (s *Service) GetInfraTwitterAppInfo(ctx context.Context, userAddress string) (*models.InfraTwitterApp, error) {
	infraTwitterApp, err := s.dao.FirstInfraTwitterApp(daos.GetDBMainCtx(ctx),
		map[string][]any{
			"address = ?": {strings.ToLower(userAddress)},
		},
		map[string][]any{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if infraTwitterApp == nil {
		return nil, errs.NewError(errs.ErrAgentUtilityNotFound)
	}

	if infraTwitterApp.ETHAddress == "" {
		ethAddress, err := s.CreateETHAddress(ctx)
		if err != nil {
			return nil, errs.NewError(err)
		}
		infraTwitterApp.ETHAddress = strings.ToLower(ethAddress)
		err = daos.GetDBMainCtx(ctx).
			Model(infraTwitterApp).
			Update("eth_address", ethAddress).
			Error
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return infraTwitterApp, nil
}

func (s *Service) TestSignature(ctx context.Context) {
	// address := "0x7c9d59cD31F27c7cBEEde2567c9fa377537bdDE0"
	timestamp := time.Now().UTC().Unix()
	fmt.Println(timestamp)
	prk := "068a7653ddda56556baadbc81ece67c6ad7d3caf7929d3908d2deb52f7a31f51"
	signature, _ := s.GetEthereumClient(ctx, models.ETHEREUM_CHAIN_ID).GetSignatureTimestamp(prk, timestamp)
	fmt.Println(signature)
	//	1741061954
	//
	// eab561ca350c4ae7c3a020b2444840ca402e29a006da937ba9817fd90d42a6b755caf4f341b322ff8c299cb93ff37d83127ce20490bbcfdeac46c7a4ed8be5251c
}

func (s *Service) InfraTwitterAppSearchRecentTweet(ctx context.Context, query, paginationToken string, maxResults int) (*twitter.TweetRecentSearch, error) {
	var tweetRecentSearch twitter.TweetRecentSearch
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]any{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]any{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		tweetRecentSearch, err := s.twitterWrapAPI.SearchRecentTweet(query, paginationToken, twitterInfo.AccessToken, maxResults)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return tweetRecentSearch, nil
	}

	return &tweetRecentSearch, nil
}
