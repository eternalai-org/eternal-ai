package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/clanker"
	"github.com/jinzhu/gorm"
)

func (s *Service) CreateClankerTokenForVideoByPostID(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("CreateClankerTokenForVideoByPostID_%d", twitterPostID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterPost, err := s.dao.FirstAgentTwitterPostByID(
						tx,
						twitterPostID,
						map[string][]interface{}{},
						true,
					)
					if err != nil {
						return errs.NewError(err)
					}

					if twitterPost == nil {
						return errs.NewError(errs.ErrRecordNotFound)
					}

					if twitterPost.Status == models.AgentTwitterPostStatusNew &&
						twitterPost.PostType == models.AgentSnapshotPostActionTypeGenerateVideo && twitterPost.TokenAddress == "" {
						// check if privy wallet already exists
						privyWallet, err := s.dao.FirstPrivyWallet(tx,
							map[string][]interface{}{
								"twitter_id = ?": {twitterPost.TwitterID},
							}, map[string][]interface{}{}, []string{})
						if err != nil {
							return errs.NewError(err)
						}

						if privyWallet == nil {
							ethAddress, err := s.CreateETHAddress(ctx)
							if err != nil {
								return errs.NewError(err)
							}

							privyWallet = &models.PrivyWallet{
								TwitterID:  twitterPost.TwitterID,
								Address:    strings.ToLower(ethAddress),
								WalletType: models.WalletTypeInternal,
								PrivyID:    helpers.RandomBigInt(12).Text(25),
							}
							err = s.dao.Create(tx, privyWallet)
							if err != nil {
								return errs.NewError(err)
							}

							// var privyResp *privy.CreateUserResp
							// var appID string
							// for _, item := range s.conf.Privy.AppIDList {
							// 	//call api privy
							// 	privyResp, err = s.privyClient.CreateUserEx(&privy.CreateUserReq{
							// 		CreateEthereumWallet: true,
							// 		LinkedAccounts: []privy.LinkedAccountReq{
							// 			{
							// 				Type:     "twitter_oauth",
							// 				Subject:  twitterPost.TwitterID,
							// 				Name:     twitterPost.TwitterName,
							// 				Username: twitterPost.TwitterUsername,
							// 			},
							// 		},
							// 	})

							// 	if err == nil && privyResp != nil && privyResp.LinkedAccounts != nil && len(privyResp.LinkedAccounts) > 0 {
							// 		appID = item
							// 		break
							// 	}
							// }

							// if privyResp != nil && appID != "" && privyResp.LinkedAccounts != nil && len(privyResp.LinkedAccounts) > 0 {
							// 	walletAddress := ""
							// 	for _, wallet := range privyResp.LinkedAccounts {
							// 		if wallet.Type == "wallet" {
							// 			walletAddress = wallet.Address
							// 			break
							// 		}
							// 	}
							// 	//create privy wallet
							// 	privyWallet = &models.PrivyWallet{
							// 		TwitterID: twitterPost.TwitterID,
							// 		Address:   walletAddress,
							// 		PrivyID:   privyResp.ID,
							// 	}
							// 	err = s.dao.Create(tx, privyWallet)
							// 	if err != nil {
							// 		return errs.NewError(err)
							// 	}
							// }
						}

						inst, err := s.dao.FirstClankerVideoToken(
							tx,
							map[string][]interface{}{
								"agent_twitter_post_id = ?": {twitterPost.ID},
							},
							map[string][]interface{}{},
							[]string{},
						)
						if err != nil {
							return errs.NewError(err)
						}

						if inst == nil && privyWallet != nil {
							inst := &models.ClankerVideoToken{
								TokenImageUrl:      twitterPost.ExtractMediaContent,
								OwnerTwitterID:     twitterPost.TwitterID,
								AgentTwitterPostID: twitterPost.ID,
								TokenDesc:          twitterPost.ExtractContent,
								TokenStatus:        "pending",
								VideoUrl:           twitterPost.ImageUrl,
								RequestorAddress:   privyWallet.Address,
								RequestKey:         helpers.RandomStringWithLength(32),
							}

							isGenImage := false
							if twitterPost.ExtractMediaContent == "" {
								isGenImage = true
							}
							tokenInfo, _ := s.GenerateTokenInfoFromVideoPrompt(ctx, twitterPost.ExtractContent, isGenImage)
							if tokenInfo != nil && tokenInfo.TokenSymbol != "" {
								inst.TokenName = tokenInfo.TokenName
								inst.TokenSymbol = tokenInfo.TokenSymbol
								if isGenImage {
									inst.TokenImageUrl = tokenInfo.TokenImageUrl
								}
							}

							if inst.TokenName != "" && inst.TokenSymbol != "" {
								//call api clanker
								tokenResp, err := s.clanker.DeployToken(&clanker.DeployTokenReq{
									Name:                     inst.TokenName,
									Symbol:                   inst.TokenSymbol,
									Description:              inst.TokenDesc,
									Image:                    inst.TokenImageUrl,
									RequestKey:               inst.RequestKey,
									RequestorAddress:         inst.RequestorAddress,
									SocialMediaUrls:          []string{fmt.Sprintf("https://x.com/%s/status/%s", twitterPost.TwitterUsername, twitterPost.TwitterPostID)},
									Platform:                 "Eternal AI",
									CreatorRewardsPercentage: s.conf.Clanker.CreatorRewardsPercentage,
									CreatorRewardsAdmin:      s.conf.Clanker.CreatorRewardsAdmin,
								})

								if err != nil {
									inst.Error = err.Error()
								} else if tokenResp != nil {
									inst.TokenAddress = tokenResp.ContractAddress
									inst.TxHash = tokenResp.TxHash
									inst.TokenStatus = "done"
									inst.PairAddress = tokenResp.PoolAddress
									inst.Error = ""

									//update status twitter post
									twitterPost.TokenName = inst.TokenName
									twitterPost.TokenSymbol = inst.TokenSymbol
									twitterPost.TokenAddress = tokenResp.ContractAddress
									err = s.dao.Save(tx, twitterPost)
									if err != nil {
										return errs.NewError(err)
									}
								}
								err = s.dao.Create(tx, inst)
								if err != nil {
									return errs.NewError(err)
								}
							} else {
								return errs.NewError(errs.ErrBadRequest, "Token name or symbol is empty")
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
		},
	)

	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) GenerateTokenInfoFromVideoPrompt(ctx context.Context, sysPrompt string, isGenImage bool) (*models.TweetParseInfo, error) {
	info := &models.TweetParseInfo{}
	sysPrompt = strings.ReplaceAll(sysPrompt, "@CryptoEternalAI", "")
	promptGenerateToken := fmt.Sprintf(`
						I want to generate my token base on this info
						'%s'

						token-name (generate if not provided, make sure it not empty, limit with 15 characters)
						token-symbol (generate if not provided, make sure it not empty, limit with 5 characters)

						Please return in string in json format including token-name, token-symbol, just only json without explanation  and token name limit with 15 characters
					`, sysPrompt)
	message := []openai.ChatCompletionMessage{
		openai.ChatCompletionMessage{
			Role:    "system",
			Content: "You are a helpful assistant",
		},
		openai.ChatCompletionMessage{
			Role:    "user",
			Content: promptGenerateToken,
		},
	}
	aiStr, err := s.openais["Agent"].CallStreamDirectlyEternalLLMV2(ctx, message, "Qwen/QwQ-32B", s.conf.KnowledgeBaseConfig.DirectServiceUrl, nil)
	if err != nil {
		return nil, errs.NewError(err)
	}
	fmt.Println(aiStr)
	if aiStr != "" {
		mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
		tokenSymbol := ""
		tokenName := ""
		if mapInfo != nil {

			if v, ok := mapInfo["token-symbol"]; ok {
				tokenSymbol = fmt.Sprintf(`%v`, v)
			}

			if v, ok := mapInfo["token-name"]; ok {
				tokenName = fmt.Sprintf(`%v`, v)
			}

			if tokenName == "" {
				tokenName = tokenSymbol
			}

			info.TokenName = tokenName
			info.TokenSymbol = tokenSymbol
			if isGenImage {
				info.TokenImageUrl = s.GenerateTokenImageBase64Gif(ctx, tokenSymbol, tokenName, sysPrompt)
			}
		}
	}

	return info, nil
}

func (s *Service) GetListUserVideo(ctx context.Context, userAddres, search string) ([]*models.ClankerVideoToken, error) {
	filters := map[string][]interface{}{
		"lower(requestor_address) = ? or lower(owner_twitter_id) = ?": {strings.ToLower(userAddres), strings.ToLower(userAddres)},
	}

	if search != "" {
		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
		filters[`
			LOWER(token_name) like ?
			or LOWER(token_symbol) like ?
			or LOWER(token_address) like ?
		`] = []any{search, search, search}
	}

	res, err := s.dao.FindClankerVideoToken(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{},
		filters,
		map[string][]interface{}{
			"AgentTwitterPost": {},
		},
		[]string{"id desc"},
		0,
		1000,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) GetVideoUserInfo(ctx context.Context, twitterID string) (*models.PrivyWallet, error) {
	insts, err := s.dao.FirstPrivyWallet(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ? or address = ?": {twitterID, strings.ToLower(twitterID)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return insts, nil
}

func (s *Service) ExportUserPrivateKeyForClaimVideoReward(ctx context.Context, userAddress string) (string, error) {
	privyWallet, err := s.dao.FirstPrivyWallet(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"user_address = ?": {strings.ToLower(userAddress)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return "", errs.NewError(err)
	}
	if privyWallet == nil {
		return "", errs.NewError(errs.ErrRecordNotFound)
	}
	prk := s.GetAddressPrk(privyWallet.Address)

	return prk, nil
}

// package services

// import (
// 	"context"
// 	"fmt"
// 	"strings"

// 	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
// 	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
// 	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
// 	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
// 	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
// 	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
// )

// func (s *Service) CreateAgentVideoByPostID(ctx context.Context, twitterPostID uint) error {
// 	err := s.JobRunCheck(
// 		ctx,
// 		fmt.Sprintf("CreateAgentVideoByPostID_%d", twitterPostID),
// 		func() error {
// 			var reqMeme *serializers.MemeReq
// 			creator := strings.ToLower(s.conf.GetConfigKeyString(models.BASE_CHAIN_ID, "meme_pool_address"))
// 			twitterPost, err := s.dao.FirstAgentTwitterPostByID(
// 				daos.GetDBMainCtx(ctx),
// 				twitterPostID,
// 				map[string][]interface{}{},
// 				true,
// 			)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}

// 			if twitterPost != nil && twitterPost.Status == models.AgentTwitterPostStatusReplied {
// 				agentInfo, err := s.dao.FirstAgentInfo(
// 					daos.GetDBMainCtx(ctx),
// 					map[string][]interface{}{
// 						"ref_tweet_id = ?": {twitterPost.ID},
// 						"agent_type = ?":   {models.AgentInfoAgentTypeVideo},
// 					},
// 					map[string][]interface{}{},
// 					[]string{},
// 				)
// 				if err != nil {
// 					return errs.NewError(err)
// 				}

// 				if agentInfo == nil {
// 					user, _ := s.dao.FirstUser(
// 						daos.GetDBMainCtx(ctx),
// 						map[string][]interface{}{
// 							"network_id = ?": {models.GENERTAL_NETWORK_ID},
// 							"twitter_id = ?": {twitterPost.GetOwnerTwitterID()},
// 						},
// 						map[string][]interface{}{},
// 						false,
// 					)

// 					if user != nil {
// 						creator = user.Address
// 					}

// 					agentInfo := &models.AgentInfo{
// 						NetworkID:      models.BASE_CHAIN_ID,
// 						NetworkName:    models.GetChainName(models.BASE_CHAIN_ID),
// 						SystemPrompt:   twitterPost.Prompt,
// 						AgentName:      twitterPost.TokenName,
// 						TokenMode:      string(models.TokenSetupEnumAutoCreate),
// 						AgentType:      models.AgentInfoAgentTypeVideo,
// 						TmpTwitterID:   twitterPost.GetOwnerTwitterID(),
// 						TokenNetworkID: models.BASE_CHAIN_ID,
// 						Version:        "2",
// 						AgentID:        helpers.RandomBigInt(12).Text(16),
// 						ScanEnabled:    true,
// 						Creator:        creator,
// 						RefTweetID:     twitterPost.ID,
// 						TokenImageUrl:  twitterPost.ImageUrl,
// 					}

// 					agentInfo.TokenMode = string(models.TokenSetupEnumAutoCreate)
// 					tokenInfo, _ := s.GenerateTokenInfoFromVideoPrompt(ctx, twitterPost.ExtractContent)
// 					if tokenInfo != nil && tokenInfo.TokenSymbol != "" {
// 						agentInfo.TokenName = tokenInfo.TokenName
// 						agentInfo.TokenSymbol = tokenInfo.TokenSymbol
// 					}

// 					agentInfo.TokenDesc = twitterPost.ExtractContent
// 					agentInfo.TokenNetworkID = models.BASE_CHAIN_ID
// 					agentInfo.SystemPrompt = twitterPost.ExtractContent
// 					agentInfo.MetaData = twitterPost.ExtractContent
// 					agentInfo.TokenStatus = "pending"
// 					agentInfo.EaiBalance = numeric.NewBigFloatFromString("50")
// 					agentInfo.Status = models.AssistantStatusReady

// 					err = s.dao.Create(daos.GetDBMainCtx(ctx), agentInfo)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}

// 					agentTokenInfo := &models.AgentTokenInfo{}
// 					agentTokenInfo.AgentInfoID = agentInfo.ID
// 					agentTokenInfo.NetworkID = models.BASE_CHAIN_ID
// 					agentTokenInfo.NetworkName = models.GetChainName(agentTokenInfo.NetworkID)
// 					err = s.dao.Create(daos.GetDBMainCtx(ctx), agentTokenInfo)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}

// 					agentInfo.TokenInfoID = agentTokenInfo.ID
// 					err = s.dao.Save(daos.GetDBMainCtx(ctx), agentInfo)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}

// 					reqMeme = &serializers.MemeReq{
// 						Name:            agentInfo.TokenName,
// 						Ticker:          agentInfo.TokenSymbol,
// 						Description:     agentInfo.TokenDesc,
// 						Image:           agentInfo.TokenImageUrl,
// 						Twitter:         fmt.Sprintf("https://x.com/%s", agentInfo.TwitterUsername),
// 						AgentInfoID:     agentInfo.ID,
// 						BaseTokenSymbol: string(models.BaseTokenSymbolEAI),
// 						NotGraduated:    true,
// 					}
// 					_, err = s.CreateMeme(ctx, creator, models.BASE_CHAIN_ID, reqMeme)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}

// 					twitterPost.Status = models.AgentTwitterPostStatusDone
// 					err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterPost)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}
// 				}
// 			}

// 			return nil
// 		},
// 	)

// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	return nil
// }

// func (s *Service) GenerateTokenInfoFromVideoPrompt(ctx context.Context, sysPrompt string) (*models.TweetParseInfo, error) {
// 	info := &models.TweetParseInfo{}
// 	sysPrompt = strings.ReplaceAll(sysPrompt, "@CryptoEternalAI", "")
// 	promptGenerateToken := fmt.Sprintf(`
// 						I want to generate my token base on this info
// 						'%s'

// 						token-name (generate if not provided, make sure it not empty)
// 						token-symbol (generate if not provided, make sure it not empty)

// 						Please return in string in json format including token-name, token-symbol, just only json without explanation  and token name limit with 15 characters
// 					`, sysPrompt)
// 	aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
// 	if err != nil {
// 		return nil, errs.NewError(err)
// 	}
// 	fmt.Println(aiStr)
// 	if aiStr != "" {
// 		mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
// 		tokenSymbol := ""
// 		tokenName := ""
// 		if mapInfo != nil {

// 			if v, ok := mapInfo["token-symbol"]; ok {
// 				tokenSymbol = fmt.Sprintf(`%v`, v)
// 			}

// 			if v, ok := mapInfo["token-name"]; ok {
// 				tokenName = fmt.Sprintf(`%v`, v)
// 			}

// 			if tokenName == "" {
// 				tokenName = tokenSymbol
// 			}
// 		}
// 		info = &models.TweetParseInfo{
// 			TokenSymbol: tokenSymbol,
// 			TokenName:   tokenName,
// 		}
// 	}

// 	return info, nil
// }

// func (s *Service) GetDashboardAgentVideo(ctx context.Context, userAddress string, tokenAddress, search string, sortListStr []string, page, limit int,
// ) ([]*models.AgentInfo, uint, error) {
// 	sortDefault := "ifnull(agent_infos.priority, 0) desc, meme_market_cap desc"
// 	if len(sortListStr) > 0 {
// 		sortDefault = strings.Join(sortListStr, ", ")
// 	}

// 	selected := []string{
// 		`ifnull(agent_infos.reply_latest_time, agent_infos.updated_at) reply_latest_time`,
// 		"agent_infos.*",
// 		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), ifnull(agent_token_infos.price_change,0)) meme_percent`,
// 		`cast(case when ifnull(agent_token_infos.usd_market_cap, 0) then ifnull(agent_token_infos.usd_market_cap, 0)
// 		when ifnull(memes.price_usd*memes.total_suply, 0) > 0 then ifnull(memes.price_usd*memes.total_suply, 0) end as decimal(36, 18)) meme_market_cap`,
// 		`ifnull(memes.price_usd, agent_token_infos.price_usd) meme_price`,
// 		`ifnull(memes.volume_last24h*memes.price_usd, agent_token_infos.volume_last24h) meme_volume_last24h`,
// 	}
// 	joinFilters := map[string][]any{
// 		`
// 			left join memes on agent_infos.id = memes.agent_info_id and memes.deleted_at IS NULL
// 			left join agent_token_infos on agent_token_infos.id = agent_infos.token_info_id
// 			left join twitter_users on twitter_users.twitter_id = agent_infos.tmp_twitter_id and  agent_infos.tmp_twitter_id is not null
// 		`: {},
// 	}

// 	filters := map[string][]any{
// 		`
// 			((agent_infos.agent_contract_address is not null and agent_infos.agent_contract_address != "") or (agent_infos.agent_type=2 and agent_infos.status="ready") or agent_infos.token_address != "")
// 			and ifnull(agent_infos.priority, 0) >= 0
// 			and agent_infos.id != 15
// 		`: {},
// 		`agent_infos.token_address != "" and ifnull(memes.status, "") not in ("created", "pending")`: {},
// 		`agent_infos.agent_type = ?`: {models.AgentInfoAgentTypeVideo},
// 	}

// 	if search != "" {
// 		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
// 		filters[`
// 			LOWER(agent_infos.token_name) like ?
// 			or LOWER(agent_infos.token_symbol) like ?
// 			or LOWER(agent_infos.token_address) like ?
// 			or LOWER(agent_infos.twitter_username) like ?
// 			or LOWER(agent_infos.agent_name) like ?
// 			or ifnull(twitter_users.twitter_username, "") like ?
// 			or ifnull(twitter_users.name, "") like ?
// 		`] = []any{search, search, search, search, search, search, search}
// 	}

// 	if tokenAddress != "" {
// 		filters["LOWER(agent_infos.token_address) = ? or agent_infos.agent_id = ? or agent_infos.id = ?"] = []any{strings.ToLower(tokenAddress), tokenAddress, tokenAddress}
// 	}

// 	//filter instlled app
// 	agents, err := s.dao.FindAgentInfoJoinSelect(
// 		daos.GetDBMainCtx(ctx),
// 		selected,
// 		joinFilters,
// 		filters,
// 		map[string][]any{
// 			"TmpTwitterInfo": {},
// 			"Meme":           {`deleted_at IS NULL and status not in ("created", "pending")`},
// 			"TokenInfo":      {},
// 		},
// 		[]string{sortDefault},
// 		page, limit,
// 	)
// 	if err != nil {
// 		return nil, 0, errs.NewError(err)
// 	}

// 	return agents, 0, nil
// }
