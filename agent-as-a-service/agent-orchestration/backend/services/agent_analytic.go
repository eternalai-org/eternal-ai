package services

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (s *Service) JobScanAgentTwitterPostForTA(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobScanAgentTwitterPostForTA",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`id in (?)`: {[]uint{s.conf.NobullshitAgentInfoId}},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				1,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = s.ScanAgentTwitterPostForTA(ctx, agent.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewError(err))
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

func (s *Service) ScanAgentTwitterPostForTA(ctx context.Context, agentID uint) error {
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterID},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if twitterInfo != nil {
		err = func() error {
			tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(agent.TwitterID, "", twitterInfo.AccessToken, 25)
			if err != nil {
				return errs.NewError(err)
			}
			err = s.CheckTwitterPostForTA(daos.GetDBMainCtx(ctx), agent.ID, agent.TwitterUsername, tweetMentions)
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
}

func (s *Service) TweetIsMentionNBS(tweets twitter.TweetObj, username string) (string, bool) {
	fullText := tweets.NoteTweet.Text
	if fullText == "" {
		fullText = tweets.Text
	}

	for _, item := range tweets.Entities.Mentions {
		fullText = strings.ReplaceAll(fullText, fmt.Sprintf("@%s", item.UserName), "")
	}

	if len(tweets.Entities.Mentions) > 0 {
		lastMentions := tweets.Entities.Mentions[len(tweets.Entities.Mentions)-1]
		if strings.EqualFold(lastMentions.UserName, username) {
			return fullText, true
		}

		if len(tweets.Entities.Mentions) > 1 {
			lastMentions = tweets.Entities.Mentions[len(tweets.Entities.Mentions)-2]
			if strings.EqualFold(lastMentions.UserName, username) {
				return fullText, true
			}
		}
	}
	return fullText, false
}

func (s *Service) TweetGetFullText(tweets twitter.TweetObj) string {
	fullText := tweets.NoteTweet.Text
	if fullText == "" {
		fullText = tweets.Text
	}

	for _, item := range tweets.Entities.Mentions {
		fullText = strings.ReplaceAll(fullText, fmt.Sprintf("@%s", item.UserName), "")
	}

	return fullText
}

func (s *Service) CheckTwitterPostForTA(tx *gorm.DB, agentInfoID uint, twitterUsername string, tweetMentions *twitter.UserTimeline) error {
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
			var checkTwitterID string
			err := s.GetRedisCachedWithKey(fmt.Sprintf("CheckTwitterPostForTA_%s", item.ID), &checkTwitterID)
			if err != nil {
				if !strings.EqualFold(item.AuthorID, agentInfo.TwitterID) {
					author, err := s.CreateUpdateUserTwitter(tx, item.AuthorID)
					if err != nil {
						return errs.NewError(errs.ErrBadRequest)
					}
					if author != nil {
						twIDs := []string{item.ID}
						twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
						if err != nil {
							return errs.NewError(err)
						}
						if twitterDetail != nil {
							for k, v := range *twitterDetail {
								fullText, isMentions := s.TweetIsMentionNBS(v.Tweet, "NOBULLSHIT_EXE")
								if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) && isMentions {
									if strings.EqualFold(k, item.ID) {
										tokenInfo, _ := s.GetTradingAnalyticInfo(context.Background(), author.TwitterUsername, fullText)
										if tokenInfo != nil && (tokenInfo.IsCreateToken) && tokenInfo.TokenSymbol != "" {
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
												postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
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
													PostType:              models.AgentSnapshotPostActionTypeReply,
													IsMigrated:            true,
												}

												m.TokenSymbol = strings.ToUpper(tokenInfo.TokenSymbol)
												m.Prompt = tokenInfo.Personality
												m.PostType = models.AgentSnapshotPostActionTypeTradeAnalytic
												err = s.dao.Create(tx, m)
												if err != nil {
													return errs.NewError(err)
												}

												_, _ = s.CreateUpdateUserTwitter(tx, m.TwitterID)
											}
										}
									}
								}
							}
						}
						// }
					}
				}
			}

			err = s.SetRedisCachedWithKey(
				fmt.Sprintf("CheckTwitterPostForTA_%s", item.ID),
				item.ID,
				12*time.Hour,
			)
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) JobAgentTwitterPostTA(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostTA",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.NobullshitAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostStatusNew},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeTradeAnalytic},
					},
					map[string][]interface{}{},
					[]string{
						"post_at desc",
					},
					0,
					5,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterPost := range twitterPosts {
					err = s.ProcessMissionTradingAnalytic(ctx, twitterPost.ID)
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

func (s *Service) ProcessMissionTradingAnalytic(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("ProcessMissionTradingAnalytic_%d", twitterPostID),
		func() error {
			agentID := uint(0)
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

					isValid := true
					existPosts, err := s.dao.FindAgentTwitterPost(
						tx,
						map[string][]interface{}{
							"not EXISTS (select 1 from agent_twitter_posts atp2 where twitter_conversation_id=? and twitter_id =? and post_type='analytic' and twitter_post_id != agent_twitter_posts.twitter_post_id )": {twitterPost.TwitterConversationId, twitterPost.TwitterID},
							"twitter_id = ?": {twitterPost.TwitterID},
							"post_type = ?":  {models.AgentSnapshotPostActionTypeTradeAnalytic},
							"status = ?":     {models.AgentTwitterPostStatusInferSubmitted},
							"created_at >= adddate(now(), interval -24 hour)": {},
						},
						map[string][]interface{}{},
						[]string{}, 0, 5,
					)
					if err != nil {
						return errs.NewError(err)
					}

					if existPosts != nil && len(existPosts) >= 3 {
						isValid = false
					}

					if isValid {
						//check valid token symbol
						binanceInfo, err := helpers.GetBinancePrice24h(fmt.Sprintf(`%sUSDT`, twitterPost.TokenSymbol))
						if err != nil || (binanceInfo.LastPrice == "") {
							isValid = false
						}
					}

					if isValid {
						if twitterPost.Status == models.AgentTwitterPostStatusNew &&
							twitterPost.PostType == models.AgentSnapshotPostActionTypeTradeAnalytic &&
							twitterPost.AgentInfo != nil {
							//find mission
							missions, err := s.dao.FindAgentSnapshotMissionJoin(
								daos.GetDBMainCtx(ctx),
								map[string][]interface{}{
									`
									join agent_infos on agent_infos.id = agent_snapshot_missions.agent_info_id
									join agent_snapshot_mission_configs on agent_snapshot_mission_configs.network_id = agent_snapshot_missions.network_id and agent_snapshot_mission_configs.tool_set = agent_snapshot_missions.tool_set
									left join twitter_infos on twitter_infos.id = agent_infos.twitter_info_id
									`: {},
								},
								map[string][]interface{}{
									"agent_snapshot_missions.enabled = 1":                            {},
									"agent_snapshot_missions.reply_enabled = 1":                      {},
									"agent_snapshot_missions.tool_set = 'trade_analytics_mentions' ": {},
									"agent_snapshot_missions.interval_sec > 0":                       {},
									`(
										agent_infos.agent_type = 1
										and agent_infos.agent_contract_id != ''
										and agent_infos.scan_enabled = 1
										and agent_infos.reply_enabled = 1
										and agent_infos.eai_balance > 0
										and agent_infos.agent_fee > 0
										and agent_infos.eai_balance >= agent_infos.agent_fee
									)`: {},
									"agent_infos.id = ?": {twitterPost.AgentInfoID},
								},
								map[string][]interface{}{},
								[]string{},
								0,
								1,
							)
							if err != nil {
								return errs.NewError(err)
							}
							for _, mission := range missions {
								err = s.AgentSnapshotPostCreate(ctx, mission.ID, twitterPost.TwitterPostID, twitterPost.TokenSymbol)
								if err != nil {
									return errs.NewError(err)
								}

								twitterPost.Status = models.AgentTwitterPostStatusInferSubmitted
								err = s.dao.Save(tx, twitterPost)
								if err != nil {
									return errs.NewError(err)
								}
							}
						}
					} else {
						twitterPost.Status = models.AgentTwitterPostStatusInvalid
						err = s.dao.Save(tx, twitterPost)
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

			if agentID > 0 {
				_ = s.CreateTokenInfo(ctx, agentID)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) GetTradingAnalyticInfo(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	info := &models.TweetParseInfo{
		IsCreateToken: false,
	}
	fullText = strings.ReplaceAll(fullText, "@CryptoEternalAI", "")
	fullText = strings.ReplaceAll(fullText, "@NOBULLSHIT_EXE", "")
	fullText = strings.ReplaceAll(fullText, "BTCH", "")

	userPrompt := fmt.Sprintf(`
Detect Token Analysis Request
This is the user conversation: "%s".

From this conversation determine if the user is requesting to analyze a specific token. Look for a clear, direct mention of a token symbol (e.g., "$XYZ").

If yes, extract or generate the following information:

Answer ("yes" or "no")
Token symbol

Return a JSON response with the following format:
{"answer": "yes/no", "symbol": ""}

Respond with only the JSON string, without any additional explanation.

	`, fullText)
	fmt.Println(userPrompt)
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
				info.IsCreateToken = true
				if v, ok := mapInfo["symbol"]; ok {
					info.TokenSymbol = fmt.Sprintf(`%v`, v)
					info.TokenSymbol = strings.ReplaceAll(info.TokenSymbol, "$", "")
				}
			}
		}
	}

	return info, nil
}

// /////////////////
func (s *Service) AnalyzeAgentTwitterPostByTweetID(tx *gorm.DB, tweetID string) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(
		tx,
		s.conf.NobullshitAgentInfoId,
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

	twIDs := []string{tweetID}
	twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
	if err != nil {
		return errs.NewError(err)
	}
	if twitterDetail != nil {
		for k, v := range *twitterDetail {
			if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
				if strings.EqualFold(k, tweetID) {
					fullText := v.Tweet.NoteTweet.Text
					if fullText == "" {
						fullText = v.Tweet.Text
					}
					tokenInfo, _ := s.GetTradingAnalyticInfo(context.Background(), v.User.UserName, fullText)
					if tokenInfo != nil && (tokenInfo.IsCreateToken) && tokenInfo.TokenSymbol != "" {
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
							postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
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
								PostType:              models.AgentSnapshotPostActionTypeReply,
								IsMigrated:            true,
							}

							m.TokenSymbol = tokenInfo.TokenSymbol
							m.Prompt = tokenInfo.Personality
							m.PostType = models.AgentSnapshotPostActionTypeTradeAnalytic

							err = s.dao.Create(tx, m)
							if err != nil {
								return errs.NewError(err)
							}

							_, _ = s.CreateUpdateUserTwitter(tx, m.TwitterID)
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) GetChartImage(ctx context.Context, symbol string) (string, error) {
	symbol = fmt.Sprintf("BINANCE:%sUSDT", strings.ToUpper(symbol))
	interval := "1D"
	viewId := strings.ToLower(helpers.RandomReferralCode(5))
	html := fmt.Sprintf(
		`
		<!DOCTYPE html>
		<html>
		<body>
		  <div id="main" style="width:640px;height:480px;">
			<div id="tradingview_%s"></div>
			<script type="text/javascript" src="https://s3.tradingview.com/tv.js"></script>
			<script type="text/javascript">
			  new TradingView.widget(
				{
				  "width": "640",
				  "height": "480",
				  "symbol": "%s",
				  "interval": "%s",
				  "timezone": "Etc/UTC",
				  "theme": "light",
				  "style": "1",
				  "locale": "en",
				  "enable_publishing": false,
				  "hide_top_toolbar": true,
				  "save_image": false,
				  "hide_volume": true,
				  "container_id": "tradingview_%s"
				}
			  );
			</script>
		  </div>
		</body>
		</html>
		`,
		viewId,
		symbol,
		interval,
		viewId,
	)
	return s.CaptureHtmlContentV4(html)
}

func (s *Service) CaptureHtmlContentV4(html string) (string, error) {
	duration := 5
	width := int64(655)
	height := int64(495)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("no-first-run", true),
	)

	ctx := context.Background()
	allocCtx, _ := chromedp.NewExecAllocator(ctx, opts...)
	cctx, cancel := chromedp.NewContext(allocCtx)

	//avoid overlap html
	ackCtx, cancel := context.WithTimeout(cctx, time.Duration(duration)*5*time.Second)
	defer cancel()

	var buf []byte
	traits := make(map[string]interface{})

	actions := []chromedp.Action{}
	actions = append(actions, chromedp.EmulateViewport(width, height))

	actions = append(actions, chromedp.Navigate("about:blank"))
	actions = append(actions, loadHTMLFromStringActionFunc(html))

	actions = append(actions, chromedp.Sleep(time.Second*time.Duration(duration)))
	actions = append(actions, chromedp.CaptureScreenshot(&buf))
	actions = append(actions, chromedp.EvaluateAsDevTools("window.$generativeTraits", &traits))

	err := chromedp.Run(ackCtx, actions...)
	if err != nil {
		return "", err
	}

	imageBase64 := base64.StdEncoding.EncodeToString(buf)
	if imageBase64 == "" {
		return "", errors.New("cannot capture image")
	}

	filename := fmt.Sprintf("%s.%s", uuid.NewString(), "png")
	urlPath, err := s.gsClient.UploadPublicDataBase64("agent", filename, imageBase64)
	if err != nil {
		return "", errs.NewError(err)
	}
	imageURL := fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath)
	fmt.Println(imageURL)
	return imageURL, nil
}

func loadHTMLFromStringActionFunc(content string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		ch := make(chan bool, 1)
		defer close(ch)

		go chromedp.ListenTarget(ctx, func(ev interface{}) {
			if _, ok := ev.(*page.EventLoadEventFired); ok {
				ch <- true
			}
		})

		frameTree, err := page.GetFrameTree().Do(ctx)
		if err != nil {
			return err
		}

		if err := page.SetDocumentContent(frameTree.Frame.ID, content).Do(ctx); err != nil {
			return err
		}

		select {
		case <-ch:
			return nil
		case <-ctx.Done():
			return context.DeadlineExceeded
		}
	})
}
