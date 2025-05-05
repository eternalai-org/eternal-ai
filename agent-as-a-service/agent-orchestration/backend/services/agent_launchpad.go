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
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) DeployMultisend(ctx context.Context, networkID uint64) (string, error) {
	address, _, err := s.GetEthereumClient(context.Background(), networkID).
		DeployMultisend(
			s.GetAddressPrk(s.conf.GetConfigKeyString(networkID, "meme_pool_address")),
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	return address, nil
}

func (s *Service) ProxyAdminDAOUpgrade(ctx context.Context, networkID uint64, proxyAddress string) (string, error) {
	txHash, err := s.GetEthereumClient(context.Background(), networkID).
		ProxyAdminUpgrade(
			s.conf.GetConfigKeyString(networkID, "proxy_admin_address"),
			s.GetAddressPrk(s.conf.GetConfigKeyString(networkID, "meme_pool_address")),
			helpers.HexToAddress(proxyAddress),
			helpers.HexToAddress(s.conf.GetConfigKeyString(networkID, "dao_treasury_logic_address")),
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	return txHash, nil
}

func (s *Service) DeployDAOTreasuryLogic(ctx context.Context, networkID uint64) (string, error) {
	address, _, err := s.GetEthereumClient(context.Background(), networkID).
		DeployDAOTreasury(
			s.GetAddressPrk(s.conf.GetConfigKeyString(networkID, "meme_pool_address")),
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	return address, nil
}

func (s *Service) DeployDAOTreasuryAddress(ctx context.Context, networkID uint64) (string, error) {
	data, err := s.GetEthereumClient(context.Background(), networkID).DAOTreasuryInitializeData(
		helpers.HexToAddress(s.conf.GetConfigKeyString(networkID, "uniswap_position_mamanger_address")),
		helpers.HexToAddress(s.conf.GetConfigKeyString(networkID, "eai_contract_address")),
	)
	if err != nil {
		return "", errs.NewError(err)
	}
	address, txHash, err := s.GetEthereumClient(context.Background(), networkID).
		DeployTransparentUpgradeableProxy(
			s.GetAddressPrk(s.conf.GetConfigKeyString(networkID, "meme_pool_address")),
			helpers.HexToAddress(s.conf.GetConfigKeyString(networkID, "dao_treasury_logic_address")),
			helpers.HexToAddress(s.conf.GetConfigKeyString(networkID, "proxy_admin_address")),
			data,
		)
	if err != nil {
		return "", errs.NewError(err)
	}
	time.Sleep(10 * time.Second)
	err = s.GetEthereumClient(context.Background(), networkID).TransactionConfirmed(txHash)
	if err != nil {
		return "", errs.NewError(err)
	}
	return address, nil
}

func (s *Service) JobScanAgentTwitterPostForCreateLaunchpad(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobScanAgentTwitterPostForCreateLaunchpad",
		func() error {
			agent, err := s.dao.FirstAgentInfoByID(
				daos.GetDBMainCtx(ctx),
				s.conf.LaunchpadAgentInfoId,
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {agent.TwitterID},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if twitterInfo != nil {
				err = func() error {
					tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(twitterInfo.TwitterID, "", twitterInfo.AccessToken, 25)
					if err != nil {
						return errs.NewError(err)
					}
					err = s.CreateAgentTwitterPostForCreateLaunchpad(daos.GetDBMainCtx(ctx), agent.ID, agent.TwitterUsername, tweetMentions)
					if err != nil {
						return errs.NewError(err)
					}
					return nil
				}()
				if err != nil {
					s.UpdateAgentScanEventError(ctx, agent.ID, err)
					return err
				} else {
					err = s.UpdateAgentScanEventSuccess(ctx, agent.ID, nil, "")
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

func (s *Service) CreateAgentTwitterPostForCreateLaunchpad(tx *gorm.DB, agentInfoID uint, twitterUsername string, tweetMentions *twitter.UserTimeline) error {
	if tweetMentions != nil {
		agentInfo, err := s.dao.FirstAgentInfoByID(
			tx,
			agentInfoID,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if agentInfo == nil {
			return errs.NewError(errs.ErrBadRequest)
		}
		twitterInfo, err := s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(errs.ErrBadRequest)
		}
		for _, item := range tweetMentions.Tweets {
			if !strings.EqualFold(item.AuthorID, agentInfo.TwitterID) {
				var rs bool
				err := s.RedisCached(
					fmt.Sprintf("CreateAgentTwitterPostForCreateLaunchpad_%s", item.ID),
					true,
					12*time.Hour,
					&rs,
					func() (interface{}, error) {
						err := func() error {
							author, err := s.CreateUpdateUserTwitter(tx, item.AuthorID)
							if err != nil {
								return errs.NewError(errs.ErrBadRequest)
							}
							if author != nil {
								twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, []string{item.ID})
								if err != nil {
									return errs.NewError(err)
								}
								if twitterDetail != nil {
									for k, v := range *twitterDetail {
										if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
											if strings.EqualFold(k, item.ID) {
												existPosts, err := s.dao.FirstAgentTwitterPost(
													tx,
													map[string][]interface{}{
														"twitter_post_id = ?": {v.Tweet.ID},
													},
													map[string][]interface{}{},
													[]string{},
												)
												if err != nil {
													return errs.NewError(err)
												}
												if existPosts == nil {
													fullText := v.Tweet.NoteTweet.Text
													if fullText == "" {
														fullText = v.Tweet.Text
													}
													var prjDesc string
													if v.Tweet.ConversationID == "" || v.Tweet.ConversationID == v.Tweet.ID {
														prjInfo, err := s.GetAgentCreateLaunchpad(context.Background(), v.User.UserName, fullText)
														if err != nil {
															return errs.NewError(err)
														}
														prjDesc = prjInfo.Description
													}
													postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
													if prjDesc != "" {
														m := &models.AgentTwitterPost{
															NetworkID:             agentInfo.NetworkID,
															AgentInfoID:           agentInfo.ID,
															TwitterID:             v.User.ID,
															TwitterUsername:       v.User.UserName,
															TwitterName:           v.User.Name,
															TwitterPostID:         v.Tweet.ID,
															Content:               fullText,
															Status:                models.AgentTwitterPostStatusNew,
															PostAt:                postedAt,
															TwitterConversationId: v.Tweet.ConversationID,
															PostType:              models.AgentSnapshotPostActionTypeCreateLaunchpad,
															IsMigrated:            true,
															TokenDesc:             prjDesc,
														}
														m.OwnerTwitterID = m.TwitterID
														m.OwnerUsername = m.TwitterUsername
														err = s.dao.Create(tx, m)
														if err != nil {
															return errs.NewError(err)
														}
													} else {
														m := &models.AgentTwitterPost{
															NetworkID:             agentInfo.NetworkID,
															AgentInfoID:           agentInfo.ID,
															TwitterID:             v.User.ID,
															TwitterUsername:       v.User.UserName,
															TwitterName:           v.User.Name,
															TwitterPostID:         v.Tweet.ID,
															Content:               fullText,
															Status:                models.AgentTwitterPostStatusInvalid,
															PostAt:                postedAt,
															TwitterConversationId: v.Tweet.ConversationID,
															PostType:              models.AgentSnapshotPostActionTypeUnknown,
															IsMigrated:            true,
														}
														m.OwnerTwitterID = m.TwitterID
														m.OwnerUsername = m.TwitterUsername
														err = s.dao.Create(tx, m)
														if err != nil {
															return errs.NewError(err)
														}
													}
													_, _ = s.CreateUpdateUserTwitter(tx, v.User.ID)
												}
											}
										}
									}
								}
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
			}
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) JobAgentTwitterPostCreateLaunchpad(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostCreateLaunchpad",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.LaunchpadAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostStatusNew},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeCreateLaunchpad},
					},
					map[string][]interface{}{},
					[]string{
						"rand()",
					},
					0,
					2,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterPost := range twitterPosts {
					err = s.AgentTwitterPostCreateLaunchpad(ctx, twitterPost.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterPost.ID))
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

func (s *Service) AgentTwitterPostCreateLaunchpad(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwitterPostCreateAgent_%d", twitterPostID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterPost, err := s.dao.FirstAgentTwitterPostByID(
						tx,
						twitterPostID,
						map[string][]interface{}{
							"AgentInfo":             {},
							"AgentInfo.TwitterInfo": {},
						},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if twitterPost.Status == models.AgentTwitterPostStatusNew &&
						twitterPost.PostType == models.AgentSnapshotPostActionTypeCreateLaunchpad {
						lp, err := s.dao.FirstLaunchpad(
							tx,
							map[string][]interface{}{
								"twitter_post_id = ?": {twitterPost.ID},
							},
							map[string][]interface{}{},
							[]string{},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if lp == nil {
							lp = &models.Launchpad{
								NetworkID:       models.BASE_CHAIN_ID,
								TwitterPostID:   twitterPost.ID,
								TweetId:         twitterPost.TwitterPostID,
								TwitterId:       twitterPost.TwitterID,
								TwitterUsername: twitterPost.TwitterUsername,
								TwitterName:     twitterPost.TwitterName,
								Description:     twitterPost.Content,
								Name:            twitterPost.TokenDesc,
								Status:          models.LaunchpadStatusNew,
								MaxFundBalance:  numeric.NewBigFloatFromString("105000"),
								TgeBalance:      numeric.NewBigFloatFromString("800000000"),
								TotalSupply:     numeric.NewBigFloatFromString("1000000000"),
							}
							err = s.dao.Create(tx, lp)
							if err != nil {
								return errs.NewError(err)
							}
							tier1, _ := models.MulBigFloats(&lp.MaxFundBalance.Float, big.NewFloat(0.02)).Float64()
							tier2, _ := models.MulBigFloats(&lp.MaxFundBalance.Float, big.NewFloat(0.01)).Float64()
							tier3, _ := models.MulBigFloats(&lp.MaxFundBalance.Float, big.NewFloat(0.005)).Float64()
							//create mission template
							agentInfo := twitterPost.AgentInfo
							mission := &models.AgentSnapshotMission{}
							mission.NetworkID = agentInfo.NetworkID
							mission.AgentInfoID = agentInfo.ID
							mission.UserPrompt = fmt.Sprintf(`Project Description: %s

Task: Analyze the data provided for the specified Twitter user (note: this data belongs to the user and is not associated with your Twitter account). Predict the percentage value of their potential contribution to the project. Based on this prediction, classify the user into one of the following tiers:
	•	Tier 1: Contribution percentage over 80%% (with a maximum allocation of %.0f eai for investment)
	•	Tier 2: Contribution percentage between 51%% and 80%% (with a maximum allocation of %.0f eai for investment)
	•	Tier 3: Contribution percentage 50%% or below (with a maximum allocation of %.0f eai for investment)

The final output should clearly indicate the tier to which the user belongs. Submit the tier and message (including the tier and maximum allocation, something like that: After analyzing your Twitter account and predicting your potential contribution to the project, we have placed you in Tier 2, with a maximum allocation of 1000 EAI for investment.) through the submit_result API.`, lp.Description, tier1, tier2, tier3)
							mission.ToolSet = models.ToolsetTypeLaunchpadJoin
							mission.NotDelay = true
							mission.Enabled = false
							mission.IsTesting = true
							mission.ReplyEnabled = true
							mission.AgentType = models.AgentInfoAgentTypeNormal
							toolList := `[{"description":"API to get twitter tweets ","executor":"https://agent.api.eternalai.org/api/internal/twitter/user/recent-info?id=%s","headers":{"api-key": "%s"},"label":"query","method":"GET","name":"get_twitter_tweets","params":[]},{"description":"API to submit result","executor":"https://agent.api.eternalai.org/api/internal/launchpad/%d/tier/%d","headers":{"api-key": "%s"},"label":"action","method":"POST","name":"submit_result","params":[{"name":"tier","dtype":"string"},{"name":"message","dtype":"string"}]}]`
							mission.ToolList = toolList
							if mission.ToolList != "" {
								mission.ReactMaxSteps = 2
							}
							//
							err = s.dao.Save(tx, mission)
							if err != nil {
								return errs.NewError(err)
							}
							//
							lp.AgentSnapshotMissionID = mission.ID
							err = s.dao.Save(tx, lp)
							if err != nil {
								return errs.NewError(err)
							}
						}
						twitterPost.Status = models.AgentTwitterPostStatusReplied
						err = s.dao.Save(tx, twitterPost)
						if err != nil {
							return errs.NewError(err)
						}
						err = s.ReplyAferAutoCreateLaunchpad(tx, twitterPost.ID, lp.ID)
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

func (s *Service) GetAgentCreateLaunchpad(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	info := &models.TweetParseInfo{}
	fullText = strings.ReplaceAll(fullText, "@"+userName, "")
	userPrompt := fmt.Sprintf(`
Detect DAO Fund Creation Requests
This is the user conversation: %s

From this conversation determine if the user is requesting assistance with fundraising, look for a direct and unambiguous statement that explicitly asks to assistance with fundraising. This statement must be clear, concise, and isolated from any surrounding context that may alter its meaning.

If yes, extract or generate the following information:

Answer ("yes" or "no")
Project information (generate if not provided, make sure it not empty and not referencing "EAI" or "Eternal AI")

Return a JSON response with the following format:
{"answer": "yes/no", "project-information": ""}

Respond with only the JSON string, without any additional explanation.
	`, fullText)
	aiStr, err := s.openais["Lama"].ChatMessage(strings.TrimSpace(userPrompt))
	if err != nil {
		return info, nil
	}
	if aiStr != "" {
		mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
		if mapInfo != nil {
			answer := "no"
			if v, ok := mapInfo["answer"]; ok {
				answer = fmt.Sprintf(`%v`, v)
			}
			if strings.EqualFold(answer, "yes") {
				info.IsCreateAgent = true
				if v, ok := mapInfo["project-information"]; ok {
					info.Description = fmt.Sprintf(`%v`, v)
				}
			}

		}
	}
	return info, nil
}

func (s *Service) ReplyAferAutoCreateLaunchpad(tx *gorm.DB, twitterPostID, launchpadId uint) error {
	if twitterPostID > 0 && launchpadId > 0 {
		twitterPost, err := s.dao.FirstAgentTwitterPostByID(
			tx,
			twitterPostID,
			map[string][]interface{}{
				"AgentInfo.TwitterInfo": {},
			},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		launchpad, err := s.dao.FirstLaunchpadByID(
			tx,
			launchpadId,
			map[string][]interface{}{},
			true,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if twitterPost != nil && launchpad != nil && twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil && twitterPost.ReplyPostId == "" {
			if launchpad.Address == "" {
				var address string
				for i := 0; i < 3; i++ {
					address, err = s.DeployDAOTreasuryAddress(context.Background(), launchpad.NetworkID)
					if err != nil {
						continue
					}
				}
				launchpad.Address = address
				err = s.dao.Save(tx, launchpad)
				if err != nil {
					return errs.NewError(err)
				}
				if launchpad.Address == "" {
					err = tx.Model(twitterPost).Updates(
						map[string]interface{}{
							"error": "failed create dao address",
						},
					).Error
					if err != nil {
						return errs.NewError(err)
					}
					return nil
				}
			}
			replyContent := fmt.Sprintf(`
We're thrilled to announce our new Dao Fund initiative, Dao %s! This visionary project empowers decentralized AI innovation through the power of community-owned compute resources.

📥 Funding Address: %s
🚀 Whitelist Applications: Now Open! Reply to this message with your Base address to apply.

Join us in shaping the future of decentralized AI!
						`, launchpad.Name, launchpad.Address)
			replyContent = strings.TrimSpace(replyContent)
			refId, err := helpers.ReplyTweetByToken(twitterPost.AgentInfo.TwitterInfo.AccessToken, replyContent, twitterPost.TwitterPostID, "")
			if err != nil {
				_ = tx.Model(twitterPost).Updates(
					map[string]interface{}{
						"reply_content": replyContent,
						"error":         err.Error(),
					},
				).Error
			} else {
				launchpad.StartTweetId = refId
				launchpad.StartAt = helpers.TimeNow()
				launchpad.EndAt = helpers.TimeAdd(time.Now(), 7*24*time.Hour)
				launchpad.Status = models.LaunchpadStatusRunning
				err = s.dao.Save(tx, launchpad)
				if err != nil {
					return errs.NewError(err)
				}
				_ = tx.Model(twitterPost).Updates(
					map[string]interface{}{
						"reply_content": replyContent,
						"reply_post_at": helpers.TimeNow(),
						"reply_post_id": refId,
						"error":         "",
					},
				).Error
			}
		}
	}
	return nil
}

func (s *Service) JobScanRepliesByLaunchpadTweetID(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobScanRepliesByLaunchpadTweetID",
		func() error {
			var retErr error
			{
				launchpads, err := s.dao.FindLaunchpad(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"status = ?": {models.LaunchpadStatusRunning},
					},
					map[string][]interface{}{
						"AgentSnapshotMission": {},
					},
					[]string{},
					0,
					1000,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, launchpad := range launchpads {
					err = daos.WithTransaction(
						daos.GetDBMainCtx(ctx),
						func(tx *gorm.DB) error {
							l, _ := s.dao.FirstLaunchpadByID(tx, launchpad.ID, map[string][]interface{}{}, true)
							if l != nil {
								meta, err := s.ScanTwitterTweetByParentID(ctx, launchpad)
								if err != nil {
									return err
								}
								if meta != nil && meta.Meta.NewestID != "" {
									l.LastScanID = meta.Meta.NewestID
									err = s.dao.Save(tx, l)
									if err != nil {
										return err
									}
								}
							}
							return nil
						},
					)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, launchpad.ID))
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

func (s *Service) ExecuteLaunchpadTier(ctx context.Context, launchpadID, memberID uint, req *serializers.TierReq) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			lp, err := s.dao.FirstLaunchpadByID(tx, launchpadID, map[string][]interface{}{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if lp != nil {
				member, err := s.dao.FirstLaunchpadMemberByID(tx, memberID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
				if member != nil && member.LaunchpadID == lp.ID {
					tier1 := models.MulBigFloats(&lp.MaxFundBalance.Float, numeric.NewFloatFromString("0.02"))
					tier2 := models.MulBigFloats(&lp.MaxFundBalance.Float, numeric.NewFloatFromString("0.01"))
					tier3 := models.MulBigFloats(&lp.MaxFundBalance.Float, numeric.NewFloatFromString("0.005"))
					member.Tier = models.LaunchpadTier(req.Tier)
					member.ReplyContent = req.Message
					if member.Tier == models.LaunchpadTier1 {
						member.MaxFundBalance = numeric.BigFloat{*tier1}
					} else if member.Tier == models.LaunchpadTier2 {
						member.MaxFundBalance = numeric.BigFloat{*tier2}
					} else if member.Tier == models.LaunchpadTier3 {
						member.MaxFundBalance = numeric.BigFloat{*tier3}
					}
					err = s.dao.Save(tx, member)
					if err != nil {
						return errs.NewError(err)
					}
					s.ReplyAfterJoinLaunchpad(tx, lp.TwitterPostID, lp.ID, member.ID, member.ReplyContent)
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

func (s *Service) ReplyAfterJoinLaunchpad(tx *gorm.DB, twitterPostID, launchpadId uint, memberID uint, replyContent string) error {
	if twitterPostID > 0 && launchpadId > 0 && memberID > 0 {
		twitterPost, err := s.dao.FirstAgentTwitterPostByID(
			tx,
			twitterPostID,
			map[string][]interface{}{
				"AgentInfo.TwitterInfo": {},
			},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		launchpad, err := s.dao.FirstLaunchpadByID(
			tx,
			launchpadId,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		member, err := s.dao.FirstLaunchpadMemberByID(
			tx,
			memberID,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if twitterPost != nil && launchpad != nil && twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil && member.ReplyPostID == "" {
			replyContent = strings.TrimSpace(replyContent)
			refId, err := helpers.ReplyTweetByToken(twitterPost.AgentInfo.TwitterInfo.AccessToken, replyContent, member.TweetID, "")
			if err != nil {
				_ = tx.Model(member).Updates(
					map[string]interface{}{
						"reply_content": replyContent,
						"error":         err.Error(),
					},
				).Error
			} else {
				_ = tx.Model(member).Updates(
					map[string]interface{}{
						"reply_content": replyContent,
						"reply_post_at": helpers.TimeNow(),
						"reply_post_id": refId,
						"error":         "",
					},
				).Error
			}
		}
	}
	return nil
}
