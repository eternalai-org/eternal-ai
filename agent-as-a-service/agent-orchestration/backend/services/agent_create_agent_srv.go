package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	io "io"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
	"github.com/sashabaranov/go-openai"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"go.uber.org/zap"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/nfnt/resize"
)

// func (s *Service) JobScanAgentTwitterPostForCreateAgent(ctx context.Context) error {
// 	err := s.JobRunCheck(
// 		ctx, "JobScanAgentTwitterPostForCreateAgent",
// 		func() error {
// 			agents, err := s.dao.FindAgentInfo(
// 				daos.GetDBMainCtx(ctx),
// 				map[string][]interface{}{
// 					`id in (?)`: {[]uint{s.conf.EternalAiAgentInfoId}},
// 				},
// 				map[string][]interface{}{},
// 				[]string{
// 					"scan_latest_time asc",
// 				},
// 				0,
// 				2,
// 			)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 			var retErr error
// 			for _, agent := range agents {
// 				err = s.ScanAgentTwitterPostFroCreateAgent(ctx, agent.ID)
// 				if err != nil {
// 					retErr = errs.MergeError(retErr, errs.NewError(err))
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

func (s *Service) JobScanAgentTwitterPostForGenerateVideo(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobScanAgentTwitterPostForGenerateVideo",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`id in (?)`: {[]uint{s.conf.VideoAiAgentInfoId}},
				},
				map[string][]interface{}{},
				[]string{
					"scan_latest_time asc",
				},
				0,
				2,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = s.ScanAgentTwitterPostForGenerateVideo(ctx, agent.ID)
				if err != nil {
					logger.Info("ScanAgentTwitterPostForGenerateVideo", "err", zap.Any("err", err), zap.Any("agent.id", agent.ID))
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

// func (s *Service) ScanAgentTwitterPostFroCreateAgent(ctx context.Context, agentID uint) error {
// 	agent, err := s.dao.FirstAgentInfoByID(
// 		daos.GetDBMainCtx(ctx),
// 		agentID,
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
// 		map[string][]interface{}{
// 			"twitter_id = ?": {s.conf.TokenTwiterID},
// 		},
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	if twitterInfo != nil {
// 		err = func() error {
// 			tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(agent.TwitterID, "", twitterInfo.AccessToken, 50)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 			err = s.CreateAgentTwitterPostForCreateAgent(daos.GetDBMainCtx(ctx), agent.ID, agent.TwitterUsername, tweetMentions)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 			return nil
// 		}()
// 		if err != nil {
// 			s.UpdateAgentScanEventError(ctx, agent.ID, err)
// 			return err
// 		} else {
// 			err = s.UpdateAgentScanEventSuccess(ctx, agent.ID, nil, "")
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 		}
// 	}
// 	return nil
// }

func (s *Service) ScanAgentTwitterPostForGenerateVideo(ctx context.Context, agentID uint) error {
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
			tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(agent.TwitterID, "", twitterInfo.AccessToken, 100)
			if err != nil {
				return errs.NewError(err)
			}
			recentSearch, err := s.twitterWrapAPI.SearchRecentTweet(fmt.Sprintf("@%v", agent.TwitterUsername), "", twitterInfo.AccessToken, 100)
			mentionIds := map[string]bool{}
			for _, v := range tweetMentions.Tweets {
				mentionIds[v.ID] = true
			}
			if err == nil {
				for _, v := range recentSearch.LookUps {
					if len(v.Tweet.ReferencedTweets) > 0 {
						retweeted := false
						for _, tweetType := range v.Tweet.ReferencedTweets {
							if tweetType.Type == "retweeted" {
								retweeted = true
								break
							}
						}
						if retweeted {
							continue
						}
					}
					if v.Tweet.AuthorID != twitterInfo.TwitterID && !mentionIds[v.Tweet.ID] {
						tweetMentions.Tweets = append(tweetMentions.Tweets, v.Tweet)
						mentionIds[v.Tweet.ID] = true
					}
				}
			}
			sort.Slice(tweetMentions.Tweets, func(i, j int) bool {
				return tweetMentions.Tweets[i].CreatedAt < tweetMentions.Tweets[j].CreatedAt
			})

			err = s.CreateAgentTwitterPostForGenerateVideo(daos.GetDBMainCtx(ctx), agent.ID, tweetMentions)
			if err != nil {
				logger.Info("CreateAgentTwitterPostForGenerateVideo", "err", zap.Any("err", err), zap.Any("agent.id", agent.ID))
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

// func (s *Service) CreateAgentTwitterPostForCreateAgent(tx *gorm.DB, agentInfoID uint, twitterUsername string, tweetMentions *twitter.UserTimeline) error {
// 	if tweetMentions == nil {
// 		return nil
// 	}

// 	agentInfo, err := s.dao.FirstAgentInfoByID(
// 		tx,
// 		agentInfoID,
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	if agentInfo == nil {
// 		return errs.NewError(errs.ErrBadRequest)
// 	}

// 	twitterInfo, err := s.dao.FirstTwitterInfo(tx,
// 		map[string][]interface{}{
// 			"twitter_id = ?": {s.conf.TokenTwiterID},
// 		},
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return errs.NewError(errs.ErrBadRequest)
// 	}

// 	for _, item := range tweetMentions.Tweets {
// 		var checkTwitterID string
// 		err := s.GetRedisCachedWithKey(fmt.Sprintf("CheckedForCreateAgent_%s", item.ID), &checkTwitterID)
// 		if err != nil {
// 			if !strings.EqualFold(item.AuthorID, agentInfo.TwitterID) {
// 				author, err := s.CreateUpdateUserTwitter(tx, item.AuthorID)
// 				if err != nil {
// 					return errs.NewError(errs.ErrBadRequest)
// 				}
// 				if author != nil {
// 					// listContext, err := s.GetConversionHistory(tx, item.ID)
// 					// if err != nil {
// 					// 	return errs.NewError(errs.ErrBadRequest)
// 					// }

// 					// jsonString, _ := json.Marshal(listContext)

// 					// tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), author.TwitterUsername, string(jsonString))
// 					// if tokenInfo != nil && (tokenInfo.IsCreateAgent) {
// 					twIDs := []string{item.ID}
// 					twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}
// 					if twitterDetail != nil {
// 						for k, v := range *twitterDetail {
// 							if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
// 								if strings.EqualFold(k, item.ID) {
// 									fullText := v.Tweet.NoteTweet.Text
// 									if fullText == "" {
// 										fullText = v.Tweet.Text
// 									}

// 									tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), author.TwitterUsername, fullText)
// 									if tokenInfo != nil && (tokenInfo.IsCreateAgent) {
// 										existPosts, err := s.dao.FirstAgentTwitterPost(
// 											tx,
// 											map[string][]interface{}{
// 												"twitter_post_id = ?": {v.Tweet.ID},
// 											},
// 											map[string][]interface{}{},
// 											[]string{},
// 										)
// 										if err != nil {
// 											return errs.NewError(err)
// 										}

// 										if existPosts == nil {
// 											postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
// 											m := &models.AgentTwitterPost{
// 												NetworkID:             agentInfo.NetworkID,
// 												AgentInfoID:           agentInfo.ID,
// 												TwitterID:             v.User.ID,
// 												TwitterUsername:       v.User.UserName,
// 												TwitterName:           v.User.Name,
// 												TwitterPostID:         v.Tweet.ID,
// 												Content:               fullText,
// 												Status:                models.AgentTwitterPostStatusNew,
// 												PostAt:                postedAt,
// 												TwitterConversationId: v.Tweet.ConversationID,
// 												PostType:              models.AgentSnapshotPostActionTypeReply,
// 												IsMigrated:            true,
// 											}

// 											m.TokenSymbol = tokenInfo.TokenSymbol
// 											m.TokenName = tokenInfo.TokenName
// 											m.TokenDesc = tokenInfo.TokenDesc
// 											m.Prompt = tokenInfo.Personality
// 											m.AgentChain = tokenInfo.ChainName
// 											m.PostType = models.AgentSnapshotPostActionTypeCreateAgent
// 											m.OwnerTwitterID = m.TwitterID
// 											m.OwnerUsername = m.TwitterUsername
// 											if strings.EqualFold(tokenInfo.ChainName, "base") && tokenInfo.IsIntellect {
// 												m.ExtractContent = "PrimeIntellect/INTELLECT-1-Instruct"
// 											}

// 											// if tokenInfo.Owner != "" {
// 											// 	twUser, _ := s.CreateUpdateUserTwitterByUserName(tx, tokenInfo.Owner)
// 											// 	if twUser != nil {
// 											// 		m.OwnerTwitterID = twUser.TwitterID
// 											// 		m.OwnerUsername = twUser.TwitterUsername
// 											// 	}
// 											// }

// 											err = s.dao.Create(tx, m)
// 											if err != nil {
// 												return errs.NewError(err)
// 											}

// 											_, _ = s.CreateUpdateUserTwitter(tx, m.TwitterID)
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 					// }
// 				}
// 			}
// 		}

// 		err = s.SetRedisCachedWithKey(
// 			fmt.Sprintf("CheckedForCreateAgent_%s", item.ID),
// 			item.ID,
// 			12*time.Hour,
// 		)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}
// 	}

// 	return nil
// }

// func (s *Service) JobAgentTwitterPostCreateAgent(ctx context.Context) error {
// 	err := s.JobRunCheck(
// 		ctx,
// 		"JobAgentTwitterPostCreateAgent",
// 		func() error {
// 			var retErr error
// 			{
// 				twitterPosts, err := s.dao.FindAgentTwitterPost(
// 					daos.GetDBMainCtx(ctx),
// 					map[string][]interface{}{
// 						"agent_info_id in (?)": {[]uint{s.conf.EternalAiAgentInfoId}},
// 						"status = ?":           {models.AgentTwitterPostStatusNew},
// 						"post_type = ?":        {models.AgentSnapshotPostActionTypeCreateAgent},
// 					},
// 					map[string][]interface{}{},
// 					[]string{
// 						"post_at desc",
// 					},
// 					0,
// 					5,
// 				)
// 				if err != nil {
// 					return errs.NewError(err)
// 				}
// 				for _, twitterPost := range twitterPosts {
// 					err = s.AgentTwitterPostCreateAgent(ctx, twitterPost.ID)
// 					if err != nil {
// 						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterPost.ID))
// 					}
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

func (s *Service) GetGenerateVideoCheckTweetHandledRedisKey(tweetId string) string {
	return fmt.Sprintf("CheckedForTweetGenerateVideo_V3_%s", tweetId)
}

type HandleGenerateVideoRequest struct {
	TweetID            string
	AgentInfoMentionID uint
	AgentInfo          *models.AgentInfo
	TwitterInfo        *models.TwitterInfo

	DecideToHandle *bool
	PromptToHandle *string
}

func (s *Service) HandleGenerateVideoWithSpecificTweet(tx *gorm.DB, handleRequest *HandleGenerateVideoRequest, source string) (*models.AgentTwitterPost, error) {
	var err error
	if handleRequest == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	tweetId := handleRequest.TweetID
	agentInfoMentionID := handleRequest.AgentInfoMentionID
	agentInfo := handleRequest.AgentInfo
	twitterInfo := handleRequest.TwitterInfo

	if agentInfo == nil {
		agentInfo, err = s.dao.FirstAgentInfoByID(
			tx,
			agentInfoMentionID,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}
		if agentInfo == nil {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
	}

	if twitterInfo == nil {
		twitterInfo, err = s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
		if twitterInfo == nil {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
	}

	var checkTweetID string
	if source != "admin_api" {
		redisKeyToCheckHandled := s.GetGenerateVideoCheckTweetHandledRedisKey(tweetId)
		err = s.GetRedisCachedWithKey(redisKeyToCheckHandled, &checkTweetID)
		//err := errors.New("redis:nil")
		if err == nil && checkTweetID != "" {
			return nil, nil // already handled
		}
	}
	twIDs := []string{tweetId}
	twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
	if err != nil {
		s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[ERROR] LookupUserTweets error %v", err))
		return nil, errs.NewError(err)
	}

	if twitterDetail == nil || len(*twitterDetail) == 0 {
		s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[ERROR] twitterDetail :%v is nil", tweetId))
		return nil, errors.New("twitterDetail is nil")
	}
	var existPosts *models.AgentTwitterPost
	v, exist := (*twitterDetail)[tweetId]
	if !exist {
		return nil, errors.New("twitterDetail is nil")
	}
	if strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
		return nil, fmt.Errorf("poster is not match v.user.id:%v , agent.twitterid:%v", v.User.ID, agentInfo.ID)
	}

	_, err = s.CreateUpdateUserTwitter(tx, v.User.ID)
	if err != nil {
		s.SendTeleVideoActivitiesAlert("[ERROR] CreateUpdateUserTwitter error")
		return nil, errs.NewError(err)
	}

	existPosts, err = s.dao.FirstAgentTwitterPost(
		tx,
		map[string][]interface{}{
			"twitter_post_id = ?": {v.Tweet.ID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if existPosts != nil {
		return existPosts, nil
	}

	fullText := v.Tweet.GetAllFullText()
	fullText = strings.Replace(fullText, fmt.Sprintf("@%s", agentInfo.TwitterUsername), "", -1)
	fullText = strings.TrimSpace(fullText)
	re := regexp.MustCompile(`^(@[\w_]+\s+)+`)
	fullText = re.ReplaceAllString(fullText, "")
	fullText = strings.TrimSpace(fullText)
	tweetParseInfo, _ := s.ValidateTweetContentGenerateVideo(context.Background(), agentInfo.TwitterUsername, fullText)

	if handleRequest.DecideToHandle != nil && *handleRequest.DecideToHandle {
		if tweetParseInfo == nil {
			tweetParseInfo = &models.TweetParseInfo{}
		}
		tweetParseInfo.IsGenerateVideo = true
		tweetParseInfo.GenerateVideoContent = fullText

		if handleRequest.PromptToHandle != nil && *handleRequest.PromptToHandle != "" {
			tweetParseInfo.GenerateVideoContent = *handleRequest.PromptToHandle
		}
	}

	if tweetParseInfo == nil || tweetParseInfo.IsGenerateVideo == false {
		if s.conf.DetectVideoLLMVersion == "v2" {
			tweetParseInfo, err = s.ValidateTweetContentGenerateVideoWithLLM2(context.Background(), fullText)
		} else {
			tweetParseInfo, err = s.ValidateTweetContentGenerateVideoWithLLM(context.Background(), agentInfo.TwitterUsername, fullText)
		}
		if err != nil {
			return nil, err
		}
	}
	if tweetParseInfo == nil || tweetParseInfo.IsGenerateVideo == false {
		s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[FAIL_SYNTAX] a requirement gen fail syntax tw_id=%v, tw_user=%v,  full_text:%v ",
			tweetId, v.User.UserName, fullText), s.conf.VideoFailSyntaxTelegramAlert)
		return nil, nil
	}
	if source != "admin_api" {
		limitPost, err := s.dao.FindAgentTwitterPost(
			tx,
			map[string][]interface{}{
				"twitter_id = ?":  {v.User.ID},
				"created_at >= ?": {time.Now().Add(-24 * time.Hour)},
			},
			map[string][]interface{}{},
			[]string{},
			0, 10,
		)
		if err != nil {
			return nil, err
		}
		if len(limitPost) >= 10 {
			s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[SKIP_LIMIT_GEN_VIDEO] gen video, "+
				"twitter_user=%v tweet_id=%v,  total_24h :%v,post :%v ",
				v.User.UserName, v.Tweet.ID, fullText, len(limitPost)))
			return nil, nil
		}
	}

	imageToVideoInfo := s.DetectTweetIsImageToVideo(twitterInfo, v)
	entityType := models.AgentTwitterPostTypeText2Video
	extractMediaContent := ""
	if imageToVideoInfo.IsImageToVideo {
		entityType = models.AgentTwitterPostTypeImage2video
		extractMediaContent = imageToVideoInfo.LighthouseImageUrl
		for {
			lastIndexOfHTTPs := strings.LastIndex(tweetParseInfo.GenerateVideoContent, "https://")
			if lastIndexOfHTTPs > 0 {
				tweetParseInfo.GenerateVideoContent = tweetParseInfo.GenerateVideoContent[:lastIndexOfHTTPs]
				tweetParseInfo.GenerateVideoContent = strings.TrimSpace(tweetParseInfo.GenerateVideoContent)
			} else {
				break
			}
		}
		extractMediaContent, err = s.ModifyTwitterImageRatio(extractMediaContent)
		if err != nil {
			return nil, err
		}
	}

	postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
	m := &models.AgentTwitterPost{
		NetworkID:             agentInfo.NetworkID,
		AgentInfoID:           agentInfo.ID,
		TwitterID:             v.User.ID,
		TwitterUsername:       v.User.UserName,
		TwitterName:           v.User.Name,
		TwitterPostID:         v.Tweet.ID, // bai reply
		Content:               fullText,
		ExtractContent:        tweetParseInfo.GenerateVideoContent,
		ExtractMediaContent:   extractMediaContent,
		Status:                models.AgentTwitterPostWaitSubmitVideoInfer,
		PostAt:                postedAt,
		TwitterConversationId: v.Tweet.ConversationID, // bai goc cua conversation
		PostType:              models.AgentSnapshotPostActionTypeGenerateVideo,
		IsMigrated:            true,
		InferId:               "20250",
		Type:                  entityType,
	}

	// now support only image to video
	if entityType != models.AgentTwitterPostTypeImage2video {
		s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[TEXT_TO_VIDEO_FOUND_SKIP] a requirement gen video, db_id=%v,"+
			" twitter_user=%v tweet_id=%v, post :%v ",
			m.ID, m.TwitterUsername, m.TwitterPostID, fullText))

		return nil, nil
	}

	err = s.dao.Create(tx, m)
	if err != nil {
		return nil, errs.NewError(err)
	}

	_, err = s.CreateUpdateUserTwitter(tx, m.TwitterID)
	if err != nil {
		s.SendTeleVideoActivitiesAlert("[ERROR] CreateUpdateUserTwitter error")
		return nil, errs.NewError(err)
	}
	s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[FOUND][%v] a requirement gen video, db_id=%v, twitter_user=%v, tweet_id=%v, post :%v ",
		source, m.ID, m.TwitterUsername, m.TwitterPostID, fullText))
	return m, nil
}

func (s *Service) CreateAgentTwitterPostForGenerateVideo(tx *gorm.DB, agentInfoID uint, tweetMentions *twitter.UserTimeline) error {
	if tweetMentions == nil {
		return nil
	}

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

	handledTweetID := make([]string, 0, len(tweetMentions.Tweets))

	for _, item := range tweetMentions.Tweets {
		_, err = s.HandleGenerateVideoWithSpecificTweet(tx, &HandleGenerateVideoRequest{
			TweetID:            item.ID,
			AgentInfoMentionID: agentInfoID,
			AgentInfo:          agentInfo,
			TwitterInfo:        twitterInfo,
		}, "scan")

		if err == nil {
			handledTweetID = append(handledTweetID, item.ID)
		}
	}

	for _, tweetId := range handledTweetID {
		redisKeyToCheckHandled := s.GetGenerateVideoCheckTweetHandledRedisKey(tweetId)
		_ = s.SetRedisCachedWithKey(
			redisKeyToCheckHandled,
			tweetId,
			1*time.Hour,
		)
	}

	return nil
}

func (s *Service) JobAgentTwitterScanResultGenerateVideo(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostCreateAgent",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"status = ?": {models.AgentTwitterPostStatusInferSubmitted},
						"infer_id IS NOT NULL  AND infer_id <> ? ":           {""},
						"infer_tx_hash IS NOT NULL  AND infer_tx_hash <> ? ": {""},
					},
					map[string][]interface{}{},
					[]string{},
					0,
					100,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterPost := range twitterPosts {
					url := fmt.Sprintf("%v?infer_id=%v&tx_hash=%v", s.conf.GetResultInferUrl, twitterPost.InferId, twitterPost.InferTxHash)
					body, _, code, err := helpers.HttpRequest(url, "GET", nil, nil)
					if err != nil {
						continue
					}
					if code != http.StatusOK {
						continue
					}
					type WorkerProcessHistory struct {
						CID            string `bson:"cid" json:"cid" json:"cid,omitempty"`
						ResultLink     string `bson:"result_link" json:"result_link" json:"result_link,omitempty"` // link to download result for all model: TEXT AND IMAGE
						ChainID        string `bson:"chain_id" json:"chain_id" json:"chain_id,omitempty"`
						TxHash         string `bson:"tx_hash" json:"tx_hash" json:"tx_hash,omitempty"`
						InferenceInput string `bson:"inference_input" json:"inference_input,omitempty"`
					}
					type Response struct {
						Data WorkerProcessHistory `json:"data"`
					}
					response := &Response{}
					err = json.Unmarshal(body, &response)
					if err != nil {
						continue
					}
					if len(response.Data.TxHash) == 0 {
						continue
					}
					err = daos.WithTransaction(
						daos.GetDBMainCtx(ctx),
						func(tx *gorm.DB) error {
							twitterPost.Status = models.AgentTwitterPostStatusNew
							twitterPost.SubmitSolutionTxHash = response.Data.TxHash
							twitterPost.ImageUrl = strings.ReplaceAll(response.Data.ResultLink, "ipfs://", "https://gateway.lighthouse.storage/ipfs/")
							err = s.dao.Save(tx, twitterPost)
							if err != nil {
								return errs.NewError(err)
							}
							prompt := twitterPost.ExtractContent
							if twitterPost.Type == models.AgentTwitterPostTypeImage2video {
								inferInput := map[string]interface{}{}
								json.Unmarshal([]byte(response.Data.InferenceInput), &inferInput)
								prompt = inferInput["prompt"].(string)
							}

							s.SendTeleVideoActivitiesAlert(fmt.Sprintf("success scan result gen video db_id:%v \n infer_id :%v \n result :%v \n \n user prompt :%v \n magic prompt :%v ", twitterPost.ID, twitterPost.InferId, twitterPost.ImageUrl, twitterPost.ExtractContent, prompt))
							return nil
						})
					if err != nil {
						continue
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

func (s *Service) JobAgentTwitterScanResultGenerateVideoMagicPrompt(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterScanResultGenerateVideoMagicPrompt",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"status = ?": {models.AgentTwitterPostStatusReplied},
						"infer_id IS NOT NULL  AND infer_id <> ? ":                       {""},
						"infer_tx_hash IS NOT NULL  AND infer_tx_hash <> ? ":             {""},
						"infer_magic_id IS NOT NULL  AND infer_magic_id <> ? ":           {""},
						"infer_magic_tx_hash IS NOT NULL  AND infer_magic_tx_hash <> ? ": {""},
						"submit_solution_magic_tx_hash = ? ":                             {""},
					},
					map[string][]interface{}{},
					[]string{},
					0,
					5,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterPost := range twitterPosts {
					url := fmt.Sprintf("%v?infer_id=%v&tx_hash=%v", s.conf.GetResultInferUrl, twitterPost.InferMagicId, twitterPost.InferMagicTxHash)
					body, _, code, err := helpers.HttpRequest(url, "GET", nil, nil)
					if err != nil {
						continue
					}
					if code != http.StatusOK {
						continue
					}
					type WorkerProcessHistory struct {
						CID            string `bson:"cid" json:"cid" json:"cid,omitempty"`
						ResultLink     string `bson:"result_link" json:"result_link" json:"result_link,omitempty"` // link to download result for all model: TEXT AND IMAGE
						ChainID        string `bson:"chain_id" json:"chain_id" json:"chain_id,omitempty"`
						TxHash         string `bson:"tx_hash" json:"tx_hash" json:"tx_hash,omitempty"`
						InferenceInput string `bson:"inference_input" json:"inference_input,omitempty"`
					}
					type Response struct {
						Data WorkerProcessHistory `json:"data"`
					}
					response := &Response{}
					err = json.Unmarshal(body, &response)
					if err != nil {
						continue
					}
					if len(response.Data.TxHash) == 0 {
						continue
					}
					daos.WithTransaction(
						daos.GetDBMainCtx(ctx),
						func(tx *gorm.DB) error {
							twitterPost.SubmitSolutionMagicTxHash = response.Data.TxHash
							err = s.dao.Save(tx, twitterPost)
							return s.dao.Save(tx, twitterPost)
						})
					inferInput := map[string]interface{}{}
					json.Unmarshal([]byte(response.Data.InferenceInput), &inferInput)
					s.SendTeleMagicVideoActivitiesAlert(fmt.Sprintf("user prompt  :%v \n :%v ", twitterPost.ExtractContent, twitterPost.ImageUrl))
					s.SendTeleMagicVideoActivitiesAlert(fmt.Sprintf("magic prompt :%v \n https://gateway.lighthouse.storage/ipfs/%v ", inferInput["prompt"],
						response.Data.CID))
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

func (s *Service) JobAgentTwitterPostCreateClankerToken(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostCreateClankerToken",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.VideoAiAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostStatusNew},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeGenerateVideo},
						"token_address = ?":    {""},
					},
					map[string][]interface{}{
						"AgentInfo":             {},
						"AgentInfo.TwitterInfo": {},
					},
					[]string{
						"post_at desc",
					},
					0,
					5,
				)
				if err != nil {
					return errs.NewError(err)
				}
				if len(twitterPosts) > 0 {
					for _, twitterPost := range twitterPosts {
						var err error
						if s.conf.Clanker.IsCreateToken {
							err = s.CreateClankerTokenForVideoByPostID(ctx, twitterPost.ID)
							if err != nil {
								retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterPost.ID))
							}
						}
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

func (s *Service) JobAgentTwitterPostGenerateVideo(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostGenerateVideo",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.VideoAiAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostStatusNew},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeGenerateVideo},
						"token_address != '' and token_address is not null": {},
					},
					map[string][]interface{}{
						"AgentInfo":             {},
						"AgentInfo.TwitterInfo": {},
					},
					[]string{
						"post_at desc",
					},
					0,
					5,
				)
				if err != nil {
					return errs.NewError(err)
				}
				if len(twitterPosts) > 0 {
					s.UpdateTwitterAccessToken(ctx, twitterPosts[0].AgentInfo.TwitterInfo.ID)

					for _, twitterPost := range twitterPosts {
						err := s.AgentTwitterPostGenerateVideoByUserTweetId(ctx, twitterPost.ID)
						if err != nil {
							retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterPost.ID))
						}

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

func (s *Service) AgentTwitterPostGenerateVideoByUserTweetId(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwitterPostGenerateVideoByUserTweetId_%d", twitterPostID),
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
						twitterPost.PostType == models.AgentSnapshotPostActionTypeGenerateVideo &&
						twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil &&
						twitterPost.TokenAddress != "" {

						videoUrl := twitterPost.ImageUrl
						mediaID := ""
						if videoUrl != "" {
							mediaID, err = s.twitterAPI.UploadVideo(models.GetImageUrl(videoUrl), []string{twitterPost.AgentInfo.TwitterID})
							if err != nil {
								s.SendTeleVideoActivitiesAlert(fmt.Sprintf("[FAIL] upload video to twitter db_id:%v , err: %v ", twitterPost.ID, err))
							}
						}

						if mediaID != "" {
							refId, err := func() (string, error) {
								var err error
								for i := 0; i < 5; i++ {
									time.Sleep(time.Duration(i*5) * time.Second)

									contentReply := fmt.Sprintf("Hey @%v, here is your decentralized video.\n\nOnchain Prompt: basescan.org/tx/%v\nOnchain Video: basescan.org/tx/%v",
										twitterPost.TwitterUsername, twitterPost.InferTxHash, twitterPost.SubmitSolutionTxHash)

									if s.conf.Clanker.IsCreateToken && twitterPost.TokenAddress != "" &&
										twitterPost.TokenName != "" && twitterPost.TokenSymbol != "" {
										contentReply = fmt.Sprintf("Hey @%v, here is your decentralized video.\n\nOnchain Video: basescan.org/tx/%v\n\nTicker $%s has been deployed.\n\n Contract address: %s\n\n Trade here: https://www.clanker.world/clanker/%s",
											twitterPost.TwitterUsername, twitterPost.SubmitSolutionTxHash, twitterPost.TokenSymbol, twitterPost.TokenAddress, twitterPost.TokenAddress)
									}

									refId, _err := helpers.ReplyTweetByToken(twitterPost.AgentInfo.TwitterInfo.AccessToken, contentReply, twitterPost.TwitterPostID, mediaID)
									if _err == nil {
										return refId, nil
									} else if strings.Contains(_err.Error(), "You attempted to reply to a Tweet that is deleted or not visible to you") {
										return "", _err
									}
									err = _err
								}

								return "", err
							}()

							if err != nil {
								s.SendTeleVideoActivitiesAlert(fmt.Sprintf("fail when reply video:\n db_id:%v \n prompt: %v \n err:%v ",
									twitterPost.ID, twitterPost.ExtractContent, err.Error()))

								if strings.Contains(err.Error(), "You attempted to reply to a Tweet that is deleted or not visible to you") {
									twitterPost.Status = models.AgentTwitterPostStatusInvalid
									err = s.dao.Save(tx, twitterPost)
									if err != nil {
										return errs.NewError(err)
									}
								}

								return errs.NewError(err)
							}
							twitterPost.ImageUrl = videoUrl
							twitterPost.ReplyPostId = refId
							twitterPost.Status = models.AgentTwitterPostStatusReplied
							err = s.dao.Save(tx, twitterPost)
							if err != nil {
								return errs.NewError(err)
							}
							s.SendTeleVideoActivitiesAlert(fmt.Sprintf("success gen video reply twitter id=%v,\n, Prompt:%v \n https://x.com/%v/status/%v \n process time :%v",
								twitterPost.ID, twitterPost.ExtractContent, twitterPost.TwitterUsername, twitterPost.TwitterPostID, time.Since(twitterPost.CreatedAt)))

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

func (s *Service) JobAgentTwitterPostSubmitVideoInfer(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostSubmitVideoInfer",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.VideoAiAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostWaitSubmitVideoInfer},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeGenerateVideo},
					},
					map[string][]interface{}{},
					[]string{
						"post_at asc",
					},
					0,
					5,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterPost := range twitterPosts {
					err = s.AgentTwitterPostSubmitVideoInferByID(ctx, twitterPost.ID)
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

func (s *Service) AgentTwitterPostSubmitVideoInferByID(ctx context.Context, agentTwitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwitterPostSubmitVideoInferByID_%d", agentTwitterPostID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterPost, err := s.dao.FirstAgentTwitterPostByID(
						tx,
						agentTwitterPostID,
						map[string][]interface{}{
							"AgentInfo":             {},
							"AgentInfo.TwitterInfo": {},
						},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}

					isValid := twitterPost.IsValidSubmitVideoInfer()

					if isValid {
						if twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil {

							model := "wan"
							prompt := twitterPost.ExtractContent
							if twitterPost.Type == models.AgentTwitterPostTypeImage2video {
								model = "Wan-I2V"
								/*if !strings.Contains(twitterPost.Content, "create video:") &&
									!strings.Contains(twitterPost.Content, "create video :") {
									videoMagicPrompt, err := s.GetVideoMagicPromptFromImage(ctx, twitterPost.ExtractContent, twitterPost.ExtractMediaContent)
									if err != nil {
										videoMagicPrompt = prompt
									}
									prompt = videoMagicPrompt
								}*/
								promptByte, _ := json.Marshal(map[string]interface{}{
									"prompt": prompt,
									"url":    twitterPost.ExtractMediaContent,
								})
								prompt = string(promptByte)
							}
							response, _, code, err := helpers.HttpRequest(s.conf.KnowledgeBaseConfig.OnChainUrl, "POST",
								map[string]string{
									"Authorization": fmt.Sprintf("Bearer %v", s.conf.KnowledgeBaseConfig.OnchainAPIKey),
								}, map[string]interface{}{
									"chain_id":          "8453",
									"model":             model,
									"prompt":            prompt,
									"only_create_infer": true,
								})
							if err != nil {
								return err
							}
							if code != http.StatusOK {
								s.SendTeleVideoActivitiesAlert(fmt.Sprintf("fail when submit infer db_id:%v \n response :%v \n code :%v ", twitterPost.ID, string(response), code))
								return fmt.Errorf("agent submit video infer response code %d", code)
							}
							type SubmitTaskResponse struct {
								InferID uint64
								TxHash  string
							}
							type DataResponse struct {
								Data  SubmitTaskResponse `json:"data"`
								Error string             `json:"error"`
							}

							dataResponse := DataResponse{}
							err = json.Unmarshal(response, &dataResponse)
							if err != nil {
								return err
							}
							if dataResponse.Error != "" {
								return fmt.Errorf("agent submit video infer response error: %s", dataResponse.Error)
							}
							if len(dataResponse.Data.TxHash) == 0 {
								return fmt.Errorf("agent submit video infer response empty tx hash")
							}
							twitterPost.InferTxHash = dataResponse.Data.TxHash
							twitterPost.InferId = strconv.FormatUint(dataResponse.Data.InferID, 10)
							now := time.Now().UTC()
							twitterPost.InferAt = &now
							twitterPost.Status = models.AgentTwitterPostStatusInferSubmitted
							err = s.dao.Save(tx, twitterPost)
							if err != nil {
								return errs.NewError(err)
							}
							s.SendTeleVideoActivitiesAlert(fmt.Sprintf("success submit infer gen video db_id:%v \n infer id :%v \n tx :%v ", twitterPost.ID, twitterPost.InferId, twitterPost.InferTxHash))
						}
					} else {
						twitterPost.Status = models.AgentTwitterConversationInvalid
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

			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

// func (s *Service) AgentTwitterPostCreateAgent(ctx context.Context, twitterPostID uint) error {
// 	err := s.JobRunCheck(
// 		ctx,
// 		fmt.Sprintf("AgentTwitterPostCreateAgent_%d", twitterPostID),
// 		func() error {
// 			agentID := uint(0)
// 			err := daos.WithTransaction(
// 				daos.GetDBMainCtx(ctx),
// 				func(tx *gorm.DB) error {
// 					twitterPost, err := s.dao.FirstAgentTwitterPostByID(
// 						tx,
// 						twitterPostID,
// 						map[string][]interface{}{
// 							"AgentInfo":             {},
// 							"AgentInfo.TwitterInfo": {},
// 						},
// 						false,
// 					)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}

// 					isValid := true
// 					existPosts, err := s.dao.FindAgentTwitterPost(
// 						tx,
// 						map[string][]interface{}{
// 							"not EXISTS (select 1 from agent_twitter_posts atp2 where twitter_conversation_id=? and owner_twitter_id =? and post_type='create_agent' and twitter_post_id != agent_twitter_posts.twitter_post_id )": {twitterPost.TwitterConversationId, twitterPost.OwnerTwitterID},
// 							"owner_twitter_id = ?": {twitterPost.OwnerTwitterID},
// 							"post_type = ?":        {models.AgentSnapshotPostActionTypeCreateAgent},
// 							"status = ?":           {models.AgentTwitterPostStatusReplied},
// 							"created_at >= adddate(now(), interval -24 hour)": {},
// 						},
// 						map[string][]interface{}{},
// 						[]string{}, 0, 5,
// 					)
// 					if err != nil {
// 						return errs.NewError(err)
// 					}

// 					if existPosts != nil && len(existPosts) >= 3 {
// 						isValid = false
// 					}

// 					if isValid {
// 						if twitterPost.Status == models.AgentTwitterPostStatusNew &&
// 							twitterPost.PostType == models.AgentSnapshotPostActionTypeCreateAgent &&
// 							twitterPost.AgentInfo != nil {
// 							networkID := models.BASE_CHAIN_ID
// 							if twitterPost.AgentChain != "" {
// 								networkID = models.GetChainID(twitterPost.AgentChain)
// 							}

// 							tokenNetworkID := models.BASE_CHAIN_ID
// 							if networkID == models.APE_CHAIN_ID {
// 								tokenNetworkID = networkID
// 							}

// 							creator := strings.ToLower(s.conf.GetConfigKeyString(models.BASE_CHAIN_ID, "meme_pool_address"))
// 							user, _ := s.dao.FirstUser(
// 								tx,
// 								map[string][]interface{}{
// 									"network_id = ?": {models.GENERTAL_NETWORK_ID},
// 									"twitter_id = ?": {twitterPost.GetOwnerTwitterID()},
// 								},
// 								map[string][]interface{}{},
// 								false,
// 							)

// 							if user != nil {
// 								creator = user.Address
// 							}

// 							agentInfo := &models.AgentInfo{
// 								NetworkID:      networkID,
// 								NetworkName:    models.GetChainName(networkID),
// 								SystemPrompt:   twitterPost.Prompt,
// 								AgentName:      twitterPost.TokenName,
// 								TokenMode:      string(models.TokenSetupEnumAutoCreate),
// 								AgentType:      models.AgentInfoAgentTypeReasoning,
// 								TmpTwitterID:   twitterPost.GetOwnerTwitterID(),
// 								TokenNetworkID: tokenNetworkID,
// 								Version:        "2",
// 								AgentID:        helpers.RandomBigInt(12).Text(16),
// 								ScanEnabled:    true,
// 								Creator:        creator,
// 							}

// 							if networkID == models.BASE_CHAIN_ID && twitterPost.ExtractContent != "" {
// 								agentInfo.AgentBaseModel = twitterPost.ExtractContent
// 							}

// 							if agentInfo.AgentBaseModel == "" {
// 								agentInfo.AgentBaseModel = s.GetModelDefaultByChainID(agentInfo.NetworkID)
// 							}

// 							ethAddress, err := s.CreateETHAddress(ctx)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}
// 							agentInfo.ETHAddress = strings.ToLower(ethAddress)
// 							agentInfo.TronAddress = trxapi.AddrEvmToTron(ethAddress)
// 							solAddress, err := s.CreateSOLAddress(ctx)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}
// 							agentInfo.SOLAddress = solAddress
// 							err = s.dao.Create(tx, agentInfo)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}
// 							agentInfo.TokenMode = string(models.TokenSetupEnumAutoCreate)
// 							agentInfo.TokenName = twitterPost.TokenName
// 							agentInfo.TokenSymbol = twitterPost.TokenSymbol
// 							agentInfo.TokenDesc = twitterPost.TokenDesc
// 							agentInfo.TokenNetworkID = tokenNetworkID
// 							agentInfo.SystemPrompt = twitterPost.Prompt
// 							agentInfo.MetaData = twitterPost.Prompt
// 							agentInfo.TokenStatus = "pending"
// 							agentInfo.EaiBalance = numeric.NewBigFloatFromString("50")
// 							agentInfo.Status = models.AssistantStatusPending

// 							agentTokenInfo := &models.AgentTokenInfo{}
// 							agentTokenInfo.AgentInfoID = agentInfo.ID
// 							agentTokenInfo.NetworkID = tokenNetworkID
// 							agentTokenInfo.NetworkName = models.GetChainName(agentTokenInfo.NetworkID)
// 							err = s.dao.Create(tx, agentTokenInfo)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}

// 							agentInfo.TokenInfoID = agentTokenInfo.ID
// 							agentInfo.RefTweetID = twitterPost.ID
// 							err = s.dao.Save(tx, agentInfo)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}

// 							twitterPost.Status = models.AgentTwitterPostStatusReplied
// 							err = s.dao.Save(tx, twitterPost)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}
// 							agentID = agentInfo.ID
// 						}
// 					} else {
// 						twitterPost.Status = models.AgentTwitterConversationInvalid
// 						err = s.dao.Save(tx, twitterPost)
// 						if err != nil {
// 							return errs.NewError(err)
// 						}
// 					}

// 					return nil
// 				},
// 			)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}

// 			if agentID > 0 {
// 				_ = s.CreateTokenInfo(ctx, agentID)
// 				// _ = s.AgentMintNft(ctx, agentID)
// 			}
// 			return nil
// 		},
// 	)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	return nil
// }

// func (s *Service) getContentTwiterForCreateAgent(ownerName, agentName, tokenSymbol, tokenDesc, tokenAddress string) string {
// 	replyContent := fmt.Sprintf(`
// Hey @%s, your Eternal AI agent $%s is live. Be the first to buy its AI coin:

// https://eternalai.org/agent/%s

// %s ($%s): %s

// PS: You can activate its autonomous tweeting ability (and soon DeFi trading ability) at https://eternalai.org/connect-x
// `, ownerName, tokenSymbol, tokenAddress, agentName, tokenSymbol, tokenDesc)
// 	return strings.TrimSpace(replyContent)
// }

// func (s *Service) getContentTwiterForCreateAgentV1(ownerName, agentName, tokenSymbol, tokenDesc, tokenAddress, chainName string, agentID uint) string {
// 	replyContent := fmt.Sprintf(`
// Hey @%s, your Eternal AI agent %s is now live on %s!

// Be the first to buy its token:

// https://eternalai.org/%d

// %s ($%s): %s

// `, ownerName, agentName, chainName, agentID, agentName, tokenSymbol, tokenDesc)
// 	return strings.TrimSpace(replyContent)
// }

// func (s *Service) ReplyAferAutoCreateAgent(tx *gorm.DB, twitterPostID, agentInfoId uint) error {
// 	if twitterPostID > 0 && agentInfoId > 0 {
// 		twitterPost, err := s.dao.FirstAgentTwitterPostByID(
// 			tx,
// 			twitterPostID,
// 			map[string][]interface{}{
// 				"AgentInfo":             {},
// 				"AgentInfo.TwitterInfo": {},
// 			},
// 			false,
// 		)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}

// 		agentInfo, err := s.dao.FirstAgentInfoByID(
// 			tx,
// 			agentInfoId,
// 			map[string][]interface{}{},
// 			false,
// 		)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}

// 		if twitterPost != nil && agentInfo != nil && twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil && twitterPost.ReplyPostId == "" {
// 			// replyContent := s.getContentTwiterForCreateAgent(twitterPost.GetAgentOnwerName(), agentInfo.AgentName, agentInfo.TokenSymbol, agentInfo.TokenDesc, agentInfo.TokenAddress)
// 			replyContent := s.getContentTwiterForCreateAgentV1(twitterPost.GetAgentOnwerName(), agentInfo.AgentName, agentInfo.TokenSymbol,
// 				agentInfo.TokenDesc, agentInfo.TokenAddress, agentInfo.NetworkName, agentInfo.ID)
// 			refId, err := helpers.ReplyTweetByToken(twitterPost.AgentInfo.TwitterInfo.AccessToken, replyContent, twitterPost.TwitterPostID, "")
// 			if err != nil {
// 				tx.Model(twitterPost).Updates(
// 					map[string]interface{}{
// 						"error": err.Error(),
// 					},
// 				)
// 			} else {
// 				_ = tx.Model(twitterPost).Updates(
// 					map[string]interface{}{
// 						"reply_post_at": helpers.TimeNow(),
// 						"reply_post_id": refId,
// 						"error":         "",
// 					},
// 				).Error

// 				//noti tele
// 				bot, err := telego.NewBot(s.conf.Telebot.Alert.Botkey, telego.WithDefaultDebugLogger())
// 				if err != nil {
// 					return errs.NewError(err)
// 				}
// 				title := "🟡🟡🟡 New Agent Alert! 🟡🟡🟡"
// 				msg := fmt.Sprintf(`just created agent %s https://eternalai.org/agent/%s 👀`, agentInfo.TokenSymbol, agentInfo.TokenAddress)
// 				_, err = bot.SendMessage(
// 					&telego.SendMessageParams{
// 						ChatID: telego.ChatID{
// 							ID: s.conf.Telebot.Alert.ChatID,
// 						},
// 						Text: strings.TrimSpace(
// 							fmt.Sprintf(
// 								`
// %s

// %s (@%s) %s

// Hey, @JohnEnt, let's do something about this!
// 	`,
// 								title,
// 								twitterPost.TwitterName,
// 								twitterPost.TwitterUsername,
// 								msg,
// 							),
// 						),
// 					},
// 				)
// 				if err != nil {
// 					return errs.NewError(err)
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func (s *Service) GetImageUrlFromTokenInfo(tokenSymbol, tokenName, tokenDesc string) (string, error) {
// 	stringBase64 := s.GenerateTokenImageBase64(context.Background(), tokenSymbol, tokenName, tokenDesc)
// 	if stringBase64 != "" {
// 		filename := fmt.Sprintf("%s.%s", uuid.NewString(), "jpg")
// 		urlPath, err := s.gsClient.UploadPublicDataBase64("agent", filename, stringBase64)
// 		if err != nil {
// 			return "", errs.NewError(err)
// 		}
// 		return fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath), nil
// 	}
// 	return "", errs.NewError(errs.ErrBadRequest)
// }

func (s *Service) GetGifImageUrlFromTokenInfo(tokenSymbol, tokenName, tokenDesc string) (string, error) {
	stringBase64 := s.GenerateTokenImageBase64Gif(context.Background(), tokenSymbol, tokenName, tokenDesc)
	if stringBase64 != "" {
		filename := fmt.Sprintf("%s.%s", uuid.NewString(), "gif")
		urlPath, err := s.gsClient.UploadPublicDataBase64("agent", filename, stringBase64)
		if err != nil {
			return "", errs.NewError(err)
		}
		url := fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath)
		return url, nil
	}
	return "", errs.NewError(errs.ErrBadRequest)
}

func (s *Service) ValidateTweetContentGenerateVideo(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	isGenerateVideo := false
	fullText = strings.TrimSpace(fullText)
	inferContent := strings.ToLower(fullText)

	if strings.Contains(inferContent, fmt.Sprintf("%v", "create video:")) {
		isGenerateVideo = true
		index := strings.Index(inferContent, "create video:")
		inferContent = fullText[index+len("create video:"):]
	} else if strings.Contains(inferContent, fmt.Sprintf("%v", "create video :")) {
		isGenerateVideo = true
		index := strings.Index(inferContent, "create video :")
		inferContent = fullText[index+len("create video :"):]
	} else if strings.Contains(inferContent, fmt.Sprintf("%v", "create video")) {
		isGenerateVideo = true
		index := strings.Index(inferContent, "create video")
		inferContent = fullText[index+len("create video"):]
	}
	inferContent = strings.TrimSpace(inferContent)
	if len(inferContent) == 0 && isGenerateVideo {
		inferContent = fullText
	}
	return &models.TweetParseInfo{
		IsGenerateVideo:      isGenerateVideo,
		GenerateVideoContent: strings.TrimSpace(inferContent),
	}, nil
}

func (s *Service) ValidateTweetContentGenerateVideoWithLLM(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	listSystemPrompts := []string{
		"You are an advanced AI tasked with detecting if a tweet on X (formerly Twitter) is **asking for action** related to generating or creating a **video**. Only tweets that contain **imperative verbs** (commands) requesting video creation or generation should return `true`.\\r\\n\\r\\n### **Important Criteria for Detection:**\\r\\n\\r\\n1. **Imperative Verb**: The verb must be used in the **imperative form** to indicate a **command** or direct action (e.g., \\\"Create video,\\\" \\\"Generate video\\\").\\r\\n   \\r\\n2. **Context of Video Creation**: The verb must refer specifically to **video creation** and not general content creation or unrelated activities (e.g., \\\"create content,\\\" \\\"make content\\\").\\r\\n\\r\\n3. **Descriptive, Informational Content**: Tweets that mention **features**, **capabilities**, or **future releases** like \\\"**video generation tools**\\\" or \\\"**AI video generation**\\\" should return `false`. These are not actionable requests and are simply **descriptive** or **informational**.\\r\\n\\r\\n4. **No Imperative Action**: If the tweet doesn\\u2019t contain an imperative verb asking the reader to take action related to video creation (e.g., \\\"**create a video**,\\\" \\\"**generate video**\\\"), then it should return `false`.\\r\\n\\r\\n### **Valid Phrases:**\\r\\n- \\\"create video\\\"\\r\\n- \\\"generate video\\\"\\r\\n- \\\"make video\\\"\\r\\n- \\\"build video\\\"\\r\\n- \\\"creat video\\\"\\r\\n- \\\"generaet video\\\"\\r\\n\\r\\nFor snake_case and hyphenated versions:\\r\\n- \\\"create_video\\\"\\r\\n- \\\"generate_video\\\"\\r\\n- \\\"make_video\\\"\\r\\n- \\\"build_video\\\"\\r\\n- \\\"create-video\\\"\\r\\n- \\\"generate-video\\\"\\r\\n\\r\\nThe key point is that the phrase should include a **clear imperative verb** asking for **video generation**.\\r\\n\\r\\n### **Do Not Flag Descriptive Tweets**: \\r\\nDo not flag tweets that describe features, releases, or technologies **without issuing an actionable command**. For example:\\r\\n- \\\"**Video generation tools**\\\" in the context of describing an upcoming release is **not actionable**.\\r\\n- **Example of what to avoid**: \\\"Tomorrow's release includes: Video generation tools.\\\" \\r\\n  - This is **informative**, not a **call to action**.\\r\\n  - **Correct action**: **Return `false`** for tweets like this.\\r\\n\\r\\n### **Default to `false` if Unclear**:\\r\\nIf the tweet content does not contain a **valid imperative phrase** (e.g., \\\"create video\\\") or if the model is **unsure** whether the tweet is asking for an action, **return `false`**. When in doubt, default to `false`.\\r\\n\\r\\n---\\r\\n\\r\\n### **Response Format**:\\r\\n\\r\\n- The result should always be returned in the following **JSON format**:\\r\\n```json\\r\\n{\\r\\n  \\\"is_generate_video\\\": true\\/false\\r\\n}\\r\\n```\\r\\n\\r\\n### **Clarification Example for \\\"Video Generation Tools\\\" Tweet:**\\r\\n\\r\\nIn the case of the tweet you provided, the phrase \\\"**Video generation tools**\\\" is part of an **informational** description about upcoming features, and it is **not actionable**. Therefore, the model should return:\\r\\n\\r\\n```json\\r\\n{\\r\\n  \\\"is_generate_video\\\": false\\r\\n}\\r\\n```",
		"You are an advanced AI tasked with detecting if a tweet on X (formerly Twitter) is **asking for action** related to generating or creating a **video**. Only tweets that contain **imperative verbs** (commands) requesting video creation or generation should return `true`.\\r\\n\\r\\n### **Important Criteria for Detection:**\\r\\n\\r\\n1. **Imperative Verb**: The verb must be used in the **imperative form** to indicate a **command** or direct action (e.g., \\\"Create video,\\\" \\\"Generate video\\\").\\r\\n   \\r\\n2. **Context of Video Creation**: The verb must refer specifically to **video creation** and not general content creation or unrelated activities (e.g., \\\"create content,\\\" \\\"make content\\\").\\r\\n\\r\\n3. **Descriptive, Informational Content**: Tweets that mention **features**, **capabilities**, or **future releases** like \\\"**video generation tools**\\\" or \\\"**AI video generation**\\\" should return `false`. These are not actionable requests and are simply **descriptive** or **informational**.\\r\\n\\r\\n4. **No Imperative Action**: If the tweet doesn\\u2019t contain an imperative verb asking the reader to take action related to video creation (e.g., \\\"**create a video**,\\\" \\\"**generate video**\\\"), then it should return `false`.\\r\\n\\r\\n### **Valid Phrases:**\\r\\n- \\\"create video\\\"\\r\\n- \\\"generate video\\\"\\r\\n- \\\"make video\\\"\\r\\n- \\\"build video\\\"\\r\\n- \\\"creat video\\\"\\r\\n- \\\"generaet video\\\"\\r\\n\\r\\nFor snake_case and hyphenated versions:\\r\\n- \\\"create_video\\\"\\r\\n- \\\"generate_video\\\"\\r\\n- \\\"make_video\\\"\\r\\n- \\\"build_video\\\"\\r\\n- \\\"create-video\\\"\\r\\n- \\\"generate-video\\\"\\r\\n\\r\\nThe key point is that the phrase should include a **clear imperative verb** asking for **video generation**.\\r\\n\\r\\n### **Do Not Flag Descriptive Tweets**: \\r\\nDo not flag tweets that describe features, releases, or technologies **without issuing an actionable command**. For example:\\r\\n- \\\"**Video generation tools**\\\" in the context of describing an upcoming release is **not actionable**.\\r\\n- **Example of what to avoid**: \\\"Tomorrow's release includes: Video generation tools.\\\" \\r\\n  - This is **informative**, not a **call to action**.\\r\\n  - **Correct action**: **Return `false`** for tweets like this.\\r\\n\\r\\n### **Default to `false` if Unclear**:\\r\\nIf the tweet content does not contain a **valid imperative phrase** (e.g., \\\"create video\\\") or if the model is **unsure** whether the tweet is asking for an action, **return `false`**. When in doubt, default to `false`.\\r\\n\\r\\n---\\r\\n\\r\\n### **Response Format**:\\r\\n\\r\\n- The result should always be returned in the following **JSON format**:\\r\\n```json\\r\\n{\\r\\n  \\\"is_generate_video\\\": true\\/false\\r\\n}\\r\\n```\\r\\n\\r\\n### **Clarification Example for \\\"Video Generation Tools\\\" Tweet:**\\r\\n\\r\\nIn the case of the tweet you provided, the phrase \\\"**Video generation tools**\\\" is part of an **informational** description about upcoming features, and it is **not actionable**. Therefore, the model should return:\\r\\n\\r\\n```json\\r\\n{\\r\\n  \\\"is_generate_video\\\": false\\r\\n}\\r\\n```",
		"You are an advanced AI tasked with detecting if a tweet on X (formerly Twitter) is **asking for action** related to generating or creating a **video**. Only tweets that contain **imperative verbs** (commands) requesting video creation or generation should return `true`.\\r\\n\\r\\n### **Important Criteria for Detection:**\\r\\n\\r\\n1. **Imperative Verb**: The verb must be used in the **imperative form** to indicate a **command** or direct action (e.g., \\\"Create video,\\\" \\\"Generate video\\\").\\r\\n   \\r\\n2. **Context of Video Creation**: The verb must refer specifically to **video creation** and not general content creation or unrelated activities (e.g., \\\"create content,\\\" \\\"make content\\\").\\r\\n\\r\\n3. **Descriptive, Informational Content**: Tweets that mention **features**, **capabilities**, or **future releases** like \\\"**video generation tools**\\\" or \\\"**AI video generation**\\\" should return `false`. These are not actionable requests and are simply **descriptive** or **informational**.\\r\\n\\r\\n4. **No Imperative Action**: If the tweet doesn\\u2019t contain an imperative verb asking the reader to take action related to video creation (e.g., \\\"**create a video**,\\\" \\\"**generate video**\\\"), then it should return `false`.\\r\\n\\r\\n### **Valid Phrases:**\\r\\n- \\\"create video\\\"\\r\\n- \\\"generate video\\\"\\r\\n- \\\"make video\\\"\\r\\n- \\\"build video\\\"\\r\\n- \\\"creat video\\\"\\r\\n- \\\"generaet video\\\"\\r\\n\\r\\nFor snake_case and hyphenated versions:\\r\\n- \\\"create_video\\\"\\r\\n- \\\"generate_video\\\"\\r\\n- \\\"make_video\\\"\\r\\n- \\\"build_video\\\"\\r\\n- \\\"create-video\\\"\\r\\n- \\\"generate-video\\\"\\r\\n\\r\\nThe key point is that the phrase should include a **clear imperative verb** asking for **video generation**.\\r\\n\\r\\n### **Do Not Flag Descriptive Tweets**: \\r\\nDo not flag tweets that describe features, releases, or technologies **without issuing an actionable command**. For example:\\r\\n- \\\"**Video generation tools**\\\" in the context of describing an upcoming release is **not actionable**.\\r\\n- **Example of what to avoid**: \\\"Tomorrow's release includes: Video generation tools.\\\" \\r\\n  - This is **informative**, not a **call to action**.\\r\\n  - **Correct action**: **Return `false`** for tweets like this.\\r\\n\\r\\n### **Default to `false` if Unclear**:\\r\\nIf the tweet content does not contain a **valid imperative phrase** (e.g., \\\"create video\\\") or if the model is **unsure** whether the tweet is asking for an action, **return `false`**. When in doubt, default to `false`.\\r\\n\\r\\n---\\r\\n\\r\\n### **Response Format**:\\r\\n\\r\\n- The result should always be returned in the following **JSON format**:\\r\\n```json\\r\\n{\\r\\n  \\\"is_generate_video\\\": true\\/false\\r\\n}\\r\\n```\\r\\n\\r\\n### **Clarification Example for \\\"Video Generation Tools\\\" Tweet:**\\r\\n\\r\\nIn the case of the tweet you provided, the phrase \\\"**Video generation tools**\\\" is part of an **informational** description about upcoming features, and it is **not actionable**. Therefore, the model should return:\\r\\n\\r\\n```json\\r\\n{\\r\\n  \\\"is_generate_video\\\": false\\r\\n}\\r\\n```",
	}
	listStatus := make([]bool, 0, len(listSystemPrompts))
	for _, systemPrompt := range listSystemPrompts {
		request := openai.ChatCompletionRequest{
			Model: "Llama3.3",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: systemPrompt,
				},
				{
					Role:    "user",
					Content: fullText,
				},
			},
		}
		isGenerateVideo := false
		response, _, code, err := helpers.HttpRequest(s.conf.KnowledgeBaseConfig.DirectServiceUrl, "POST",
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %v", s.conf.KnowledgeBaseConfig.OnchainAPIKey),
			}, request)
		if err != nil {
			isGenerateVideo = false
		}
		if code != http.StatusOK {
			isGenerateVideo = false
		}
		res := openai.ChatCompletionResponse{}
		err = json.Unmarshal(response, &res)
		if err != nil {
			isGenerateVideo = false
		}

		if len(res.Choices) > 0 {
			result := map[string]bool{}
			res.Choices[0].Message.Content = strings.Replace(res.Choices[0].Message.Content, "```json", "", 1)
			res.Choices[0].Message.Content = strings.Replace(res.Choices[0].Message.Content, "```", "", 1)
			err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &result)
			if err != nil {
				isGenerateVideo = false
			}
			isGenerateVideo = result["is_generate_video"]
		}
		listStatus = append(listStatus, isGenerateVideo)
	}
	score := 0
	for _, status := range listStatus {
		if status {
			score++
		}
	}
	isGenerateVideo := false
	if score*3 >= len(listStatus)*2 {
		isGenerateVideo = true
	}
	/*if isGenerateVideo {
		return s.GetPromptFromTweetContentGenerateVideoWithLLM(ctx, userName, fullText)
	}*/
	return &models.TweetParseInfo{
		IsGenerateVideo:      isGenerateVideo,
		GenerateVideoContent: fullText,
	}, nil
}

func (s *Service) ValidateTweetContentGenerateVideoWithLLM2(ctx context.Context, fullText string) (*models.TweetParseInfo, error) {
	msg, _ := json.Marshal([]openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "You are an AI assistant that verifies whether an input is an video generation prompt.",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Please think carefully and determine whether the content of the tweet/prompt is relevant to a request to create a video/animation/GIF. A tweet/prompt is considered a request to create a video when the content of the tweet/prompt clearly states the INTENT TO CREATE A VIDEO/ANIMATION/GIF for a specific content. NOTE: Link do not affect the content of the tweet/prompt.\n\n- If not, return: 'NONE'\n- If yes, extract only the most important content in the tweet/prompt while keeping the original content intact. If the content cannot be extracted, return to the original prompt:\n\n**RULES:**\n1.  **Focus on the core subject:** Identify the main object or scene in the prompt.\n2.  **Remove unnecessary descriptions:** Eliminate vague, redundant, or irrelevant words that do not significantly affect the generated video.\n3.  **Retain critical elements:** Keep essential aspects such as objects, style, lighting, colors, and composition.\n4.  **Use concise and clear language:** Rewrite the prompt in a clear and effective way that maximizes model accuracy.\n5.  The content of the extracted tweet MUST be the same as the original tweet content.\n6.  If the specific content cannot be extracted from the prompt, return the original prompt.\n7.  Return the result in JSON format with the \"optimized_prompt\" key.\n\n**EXAMPLE:**\nINPUT 1: \"create video the man in the photo smile in the rain\"\n\nOUTPUT 1: {\"optimized_prompt\": \"Man in the photo smiling in the rain\"}\n\n\nINPUT 2: \"the character fighting the monster is a woman\"\nOUTPUT 2: {\"optimized_prompt\": \"NONE\"}\n\nApply these rules to optimize prompts effectively.\n\n**Response Format**:\n- The answer should always be returned in the following **JSON format**:\n{\n  \"optimized_prompt\": \"\"\n}\r\nUser Query: %v", fullText),
		},
	})

	outputChan := make(chan *models.ChatCompletionStreamResponse, 1)
	errChan := make(chan error, 1)
	doneChan := make(chan bool, 1)
	go func() {
		s.openais["Agent"].CallStreamDirectlyEternalLLM(ctx, string(msg), "Qwen/QwQ-32B",
			s.conf.KnowledgeBaseConfig.DirectServiceUrl, map[string]interface{}{
				"temperature": 0,
			}, outputChan, errChan, doneChan)
	}()
	output := ""
	thinking := false
	done := false
	for !done {
		select {
		case <-doneChan:
			done = true
		case err := <-errChan:
			return nil, err
		case res := <-outputChan:
			if res.Choices[0].Delta.Content != "" {
				if strings.Contains(res.Choices[0].Delta.Content, "<think>") {
					thinking = true
				} else if strings.Contains(res.Choices[0].Delta.Content, "</think>") {
					thinking = false
					continue
				}
				if !thinking {
					output += res.Choices[0].Delta.Content
				}
			}
		case <-time.After(5 * time.Minute):
			return nil, errors.New("timeout")
		}
	}

	type Result struct {
		OptimizedPrompt string `json:"optimized_prompt"`
	}
	var result Result
	err := json.Unmarshal([]byte(output), &result)
	if err != nil {
		return nil, err
	}
	if result.OptimizedPrompt == "NONE" {
		return &models.TweetParseInfo{
			IsGenerateVideo:      false,
			GenerateVideoContent: result.OptimizedPrompt,
		}, nil
	}
	return &models.TweetParseInfo{
		IsGenerateVideo:      true,
		GenerateVideoContent: result.OptimizedPrompt,
	}, nil
}

type TweetImageToVideo struct {
	IsImageToVideo     bool
	LighthouseImageUrl string
}

func (s *Service) UploadImageUrlToLighthouse(ctx context.Context, imageUrl string) (string, error) {
	filename := fmt.Sprintf("%s", uuid.NewString())
	resp, err := http.Get(imageUrl)
	if err != nil {
		return "", errs.NewError(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", errs.NewError(err)
	}

	cid, err := lighthouse.UploadDataWithRetry(s.conf.Lighthouse.Apikey, filename, body)
	if err != nil {
		return "", errs.NewError(err)
	}

	return cid, nil
}

func (s *Service) DetectTweetIsImageToVideo(twitterInfo *models.TwitterInfo, item twitter.TweetLookup) *TweetImageToVideo {
	isImageToVideo := false
	cid := ""
	var err error
	if len(item.AttachmentMedia) > 0 {
		fmt.Println("medias", item.AttachmentMedia)
		firstMedia := item.AttachmentMedia[0]
		if firstMedia.Type == "photo" {
			isImageToVideo = true

			// TODO first
			return &TweetImageToVideo{
				IsImageToVideo:     isImageToVideo,
				LighthouseImageUrl: firstMedia.URL,
			}

			cid, err = s.UploadImageUrlToLighthouse(context.Background(), firstMedia.URL)
			if err == nil && cid != "" {
				isImageToVideo = true
			}
		}
	}

	if len(item.Tweet.ReferencedTweets) > 0 {
		for _, refTweet := range item.Tweet.ReferencedTweets {
			if refTweet.Type == "quoted" {
				refTweetID := refTweet.ID
				refTweetDetails, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, []string{refTweetID})
				if err == nil && refTweetDetails != nil {
					for k, v := range *refTweetDetails {
						if k != refTweetID {
							continue
						}

						if len(v.AttachmentMedia) > 0 {
							firstMedia := v.AttachmentMedia[0]
							if firstMedia.Type == "photo" {
								isImageToVideo = true

								// TODO first
								return &TweetImageToVideo{
									IsImageToVideo:     isImageToVideo,
									LighthouseImageUrl: firstMedia.URL,
								}

								cid, err = s.UploadImageUrlToLighthouse(context.Background(), firstMedia.URL)
								if err == nil && cid != "" {
									isImageToVideo = true
								}
							}
						}
					}
				}
			}

			if refTweet.Type == "replied_to" {
				resp, err := s.GetFirstImageFromTweet(twitterInfo, refTweet.ID)
				if err == nil && resp != nil {
					return resp
				}
			}
		}
	}

	return &TweetImageToVideo{
		IsImageToVideo:     isImageToVideo,
		LighthouseImageUrl: lighthouse.IPFSGateway + cid,
	}
}

func (s *Service) ModifyTwitterImageRatio(imageUrl string) (string, error) {
	res, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New(res.Status)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	imageInfo, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return imageUrl, nil
	}
	width := float64(imageInfo.Bounds().Dx())
	height := float64(imageInfo.Bounds().Dy())
	ratio := height / width
	if ratio <= 2.5 && ratio >= 0.4 {
		return imageUrl, nil
	}
	newWidth := imageInfo.Bounds().Dx()
	newHeight := newWidth * 2 / 5 // height = width/2.5
	if width < height {
		newHeight = imageInfo.Bounds().Dy()
		newWidth = newHeight * 2 / 5 // width = height/2.5
	}
	resizedImg := resize.Resize(uint(newWidth), uint(newHeight), imageInfo, resize.Lanczos3)

	filledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	dstRect := image.Rect(0, 0, resizedImg.Bounds().Dx(), resizedImg.Bounds().Dy())
	draw.Draw(filledImg, filledImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src) // Fill with white
	draw.Draw(filledImg, dstRect, resizedImg, image.Point{}, draw.Over)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, filledImg, nil)
	if err != nil {
		return "", err
	}
	response, err := http.Post(s.conf.LighthouseUploadBinaryUrl, "", bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errors.New(res.Status)
	}
	output, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	type LightHouseResult struct {
		Error string `json:"error"`
		Data  string `json:"data"`
	}
	var result LightHouseResult
	err = json.Unmarshal(output, &result)
	if err != nil {
		return "", err
	}
	if len(result.Error) > 0 {
		return "", errors.New(result.Error)
	}
	if len(result.Data) == 0 {
		return "", errors.New("can not upload image to lighthouse")
	}
	return strings.ReplaceAll(result.Data, "ipfs://", "https://gateway.lighthouse.storage/ipfs/"), err
}

func (s *Service) GetFirstImageFromTweet(twitterInfo *models.TwitterInfo, tweetID string) (*TweetImageToVideo, error) {
	var getImage func(tweetID string) (*TweetImageToVideo, error)
	getImage = func(tweetID string) (*TweetImageToVideo, error) {
		tweetDetails, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, []string{tweetID})
		if err != nil {
			return nil, err
		}
		if tweetDetails == nil {
			return nil, fmt.Errorf("no tweet found")
		}

		for _, tweet := range *tweetDetails {
			if len(tweet.AttachmentMedia) > 0 {
				firstMedia := tweet.AttachmentMedia[0]
				if firstMedia.Type == "photo" {
					return &TweetImageToVideo{
						IsImageToVideo:     true,
						LighthouseImageUrl: firstMedia.URL,
					}, nil
				}
			}

			for _, refTweet := range tweet.Tweet.ReferencedTweets {
				if refTweet.Type == "replied_to" {
					return getImage(refTweet.ID)
				}
			}
		}

		return nil, fmt.Errorf("no image found in tweet or its references")
	}

	return getImage(tweetID)
}

func (s *Service) GetPromptFromTweetContentGenerateVideoWithLLM(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	request := openai.ChatCompletionRequest{
		Model: "Llama3.3",
		Messages: []openai.ChatCompletionMessage{
			openai.ChatCompletionMessage{
				Role:    "system",
				Content: "You are an advanced AI tasked with accurately detecting if a tweet on X (formerly Twitter) mentions generating or creating a video. Pay close attention to the phrasing used by the user, including common misspellings, variations, case-insensitivity, and the use of snake_case or hyphenated notation. Before answering, carefully evaluate the content of the tweet, ensuring that it directly references video creation or generation. Respond with the highest accuracy possible and provide a relevant video prompt.\n\nReturn the response in the following JSON format:\n{\n  \"prompt\": \"\"\n}\nYour response should reflect whether the tweet is related to video generation or not. Think critically about context and phrasing to ensure the most accurate determination.",
			}, openai.ChatCompletionMessage{
				Role:    "user",
				Content: fullText,
			},
		},
	}
	response, _, code, err := helpers.HttpRequest(s.conf.KnowledgeBaseConfig.DirectServiceUrl, "POST",
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", s.conf.KnowledgeBaseConfig.OnchainAPIKey),
		}, request)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("err while get response code:%v ,body:%v", code, string(response))
	}
	res := openai.ChatCompletionResponse{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, err
	}
	prompt := ""
	if len(res.Choices) > 0 {
		result := map[string]string{}
		err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &result)
		if err != nil {
			return nil, err
		}
		prompt = result["prompt"]
	}
	return &models.TweetParseInfo{
		IsGenerateVideo:      true,
		GenerateVideoContent: strings.TrimSpace(prompt),
	}, nil
}

// func (s *Service) GetAgentInfoInContent(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
// 	info := &models.TweetParseInfo{
// 		IsCreateToken: false,
// 		IsCreateAgent: false,
// 	}
// 	fullText = strings.ReplaceAll(fullText, "@CryptoEternalAI", "")
// 	if strings.Contains(fullText, `🍌`) {
// 		info.IsCreateAgent = true
// 		info.ChainName = "apechain"

// 		promptGenerateToken := fmt.Sprintf(`
// 						I want to generate my agent infomation base on this info
// 						'%s'

// 						Agent name (generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI" or "CryptoEternalAI" or "Crypto Eternal AI)
// 						Agent token symbol (generate if not provided, generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI" or "CryptoEternalAI" or "Crypto Eternal AI)
// 						Agent backstory (generate if not provided, generate if not provided, make sure it not empty and not referencing "EAI" or "Eternal AI" or "CryptoEternalAI" or "Crypto Eternal AI)
// 						Agent personality (predefined instruction to guide the Agent's behavior during a conversation or task, generate if not provided)

// 						Return a JSON response with the following format:
// 						{"name": "", "symbol": "", "story": "", "personality": ""}

// 						Respond with only the JSON string, without any additional explanation.
// 					`, fullText)
// 		aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
// 		if err != nil {
// 			return info, nil
// 		}
// 		fmt.Println(aiStr)
// 		if aiStr != "" {
// 			mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
// 			if mapInfo != nil {
// 				if v, ok := mapInfo["personality"]; ok {
// 					info.Personality = fmt.Sprintf(`%v`, v)
// 				}

// 				if v, ok := mapInfo["name"]; ok {
// 					info.TokenName = fmt.Sprintf(`%v`, v)
// 				}

// 				if v, ok := mapInfo["symbol"]; ok {
// 					info.TokenSymbol = fmt.Sprintf(`%v`, v)
// 				}

// 				if v, ok := mapInfo["story"]; ok {
// 					info.TokenDesc = fmt.Sprintf(`%v`, v)
// 				}

// 				if v, ok := mapInfo["personality"]; ok {
// 					info.Personality = fmt.Sprintf(`%v`, v)
// 				}
// 			}
// 		}
// 	} else {
// 		userPrompt := fmt.Sprintf(`
// 	Detect Agent Creation Request
// 	This is the user conversation: "%s".

// 	From this conversation determine if the user is requesting you to create an agent, also referred as a decentralized agent (dagent), look for a direct and unambiguous statement that explicitly asks to create an agent. This statement must be clear, concise, and isolated from any surrounding context that may alter its meaning.

// 	If yes, extract or generate the following information:

// 	Answer ("yes" or "no")
// 	Owner (who is the owner of the agent)
// 	Agent name (generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI")
// 	Agent token symbol (generate if not provided, generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI")
// 	Agent backstory (generate if not provided, generate if not provided, make sure it not empty and not referencing "EAI" or "Eternal AI")
// 	Blockchain ("base" if not provided, "base" or "arbitrum" or "bsc" or "bnbchain" or "binancechain" or "polygon" or "avax" or "avalanche" or "apechain")
// 	Is Intellect Model ("yes" or "no")
// 	Agent personality (predefined instruction to guide the Agent's behavior during a conversation or task, generate if not provided)

// 	Return a JSON response with the following format:
// 	{"answer": "yes/no", "owner": "", "name": "", "symbol": "", "story": "", "blockchain": "", "personality": "", , "is_intellect": ""}

// 	Respond with only the JSON string, without any additional explanation.
// 	`, fullText)
// 		fmt.Println(userPrompt)
// 		aiStr, err := s.openais["Lama"].ChatMessage(strings.TrimSpace(userPrompt))
// 		if err != nil {
// 			return info, nil
// 		}

// 		if aiStr != "" {
// 			mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
// 			if mapInfo != nil {
// 				answer := "no"
// 				if v, ok := mapInfo["answer"]; ok {
// 					answer = fmt.Sprintf(`%v`, v)
// 				}

// 				if strings.EqualFold(answer, "yes") {
// 					info.IsCreateAgent = true
// 					if v, ok := mapInfo["personality"]; ok {
// 						info.Personality = fmt.Sprintf(`%v`, v)
// 					}

// 					if v, ok := mapInfo["name"]; ok {
// 						info.TokenName = fmt.Sprintf(`%v`, v)
// 					}

// 					if v, ok := mapInfo["symbol"]; ok {
// 						info.TokenSymbol = fmt.Sprintf(`%v`, v)
// 					}

// 					if v, ok := mapInfo["story"]; ok {
// 						info.TokenDesc = fmt.Sprintf(`%v`, v)
// 					}

// 					if v, ok := mapInfo["blockchain"]; ok {
// 						info.ChainName = strings.ToLower(fmt.Sprintf(`%v`, v))
// 					}

// 					if v, ok := mapInfo["owner"]; ok {
// 						info.Owner = strings.ToLower(fmt.Sprintf(`%v`, v))
// 					}

// 					if v, ok := mapInfo["is_intellect"]; ok {
// 						if strings.ToLower(fmt.Sprintf(`%v`, v)) == "yes" {
// 							info.IsIntellect = true
// 						}
// 					}
// 				}

// 			}
// 		}
// 	}

// 	return info, nil
// }

func (s *Service) GetImageUrlForBase64(stringBase64 string) (string, error) {
	if stringBase64 != "" {
		filename := fmt.Sprintf("%s.%s", uuid.NewString(), "jpg")
		urlPath, err := s.gsClient.UploadPublicDataBase64("tweetv2", filename, stringBase64)
		if err != nil {
			return "", errs.NewError(err)
		}
		return fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath), nil
	}
	return "", errs.NewError(errs.ErrBadRequest)
}

// /////////////////
// func (s *Service) CreateAgentTwitterPostByTweetID(tx *gorm.DB, tweetID string) error {
// 	agentInfo, err := s.dao.FirstAgentInfoByID(
// 		tx,
// 		s.conf.EternalAiAgentInfoId,
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	if agentInfo == nil {
// 		return errs.NewError(errs.ErrBadRequest)
// 	}

// 	twitterInfo, err := s.dao.FirstTwitterInfo(tx,
// 		map[string][]interface{}{
// 			"twitter_id = ?": {s.conf.TokenTwiterID},
// 		},
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return errs.NewError(errs.ErrBadRequest)
// 	}

// 	twIDs := []string{tweetID}
// 	twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
// 	if err != nil {
// 		return errs.NewError(err)
// 	}
// 	if twitterDetail != nil {
// 		for k, v := range *twitterDetail {
// 			if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
// 				if strings.EqualFold(k, tweetID) {
// 					fullText := v.Tweet.NoteTweet.Text
// 					if fullText == "" {
// 						fullText = v.Tweet.Text
// 					}
// 					// listContext, err := s.GetConversionHistory(tx, v.Tweet.ID)
// 					// if err != nil {
// 					// 	return errs.NewError(errs.ErrBadRequest)
// 					// }

// 					// jsonString, _ := json.Marshal(listContext)
// 					// tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), v.User.UserName, string(jsonString))
// 					tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), v.User.UserName, fullText)
// 					// tokenInfo := &models.TweetParseInfo{
// 					// 	IsCreateAgent: true,
// 					// 	TokenName:     "GrowkAI",
// 					// 	TokenSymbol:   "GROWK",
// 					// 	ChainName:     "arbitrum",
// 					// 	TokenDesc:     "Growk is a frog based regen meme that will forever change the way we think of memes and public goods",
// 					// 	Personality:   "Be friendly and helpful, and provide information about the Growk meme and its community",
// 					// 	Owner:         v.User.UserName,
// 					// }
// 					if tokenInfo != nil && (tokenInfo.IsCreateAgent) {
// 						isValid := true
// 						existPosts, err := s.dao.FirstAgentTwitterPost(
// 							tx,
// 							map[string][]interface{}{
// 								"twitter_post_id = ?": {v.Tweet.ID},
// 							},
// 							map[string][]interface{}{},
// 							[]string{},
// 						)
// 						if err != nil {
// 							return errs.NewError(err)
// 						}

// 						if existPosts != nil {
// 							isValid = false
// 						}

// 						if isValid {
// 							postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
// 							m := &models.AgentTwitterPost{
// 								NetworkID:             agentInfo.NetworkID,
// 								AgentInfoID:           agentInfo.ID,
// 								TwitterID:             v.User.ID,
// 								TwitterUsername:       v.User.UserName,
// 								TwitterName:           v.User.Name,
// 								TwitterPostID:         v.Tweet.ID,
// 								Content:               fullText,
// 								Status:                models.AgentTwitterPostStatusNew,
// 								PostAt:                postedAt,
// 								TwitterConversationId: v.Tweet.ConversationID,
// 								PostType:              models.AgentSnapshotPostActionTypeReply,
// 								IsMigrated:            true,
// 							}

// 							m.TokenSymbol = tokenInfo.TokenSymbol
// 							m.TokenName = tokenInfo.TokenName
// 							m.TokenDesc = tokenInfo.TokenDesc
// 							m.Prompt = tokenInfo.Personality
// 							m.AgentChain = tokenInfo.ChainName
// 							m.PostType = models.AgentSnapshotPostActionTypeCreateAgent

// 							m.OwnerTwitterID = m.TwitterID
// 							m.OwnerUsername = m.TwitterUsername

// 							if tokenInfo.Owner != "" {
// 								twUser, _ := s.CreateUpdateUserTwitterByUserName(tx, tokenInfo.Owner)
// 								if twUser != nil {
// 									m.OwnerTwitterID = twUser.TwitterID
// 									m.OwnerUsername = twUser.TwitterUsername
// 								}
// 							}

// 							err = s.dao.Create(tx, m)
// 							if err != nil {
// 								return errs.NewError(err)
// 							}

// 							_, _ = s.CreateUpdateUserTwitter(tx, m.TwitterID)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func (s *Service) BuildConversionHistory(tx *gorm.DB, tweetID string) (string, error) {
// 	userPrompt := ""
// 	listContext, _ := s.GetConversionHistory(tx, tweetID)
// 	if len(listContext) > 0 {
// 		for _, item := range listContext {
// 			userPrompt += fmt.Sprintf(`
// 				@%s: %s
// 			`, item["twitter_username"], item["text"])
// 		}
// 	}
// 	return userPrompt, nil
// }

// func (s *Service) GetConversionHistory(tx *gorm.DB, tweetID string) ([]map[string]string, error) {
// 	listContext := []map[string]string{}
// 	twitterInfo, err := s.dao.FirstTwitterInfo(tx,
// 		map[string][]interface{}{
// 			"twitter_id = ?": {s.conf.TokenTwiterID},
// 		},
// 		map[string][]interface{}{},
// 		false,
// 	)
// 	if err != nil {
// 		return listContext, errs.NewError(err)
// 	}
// 	if twitterInfo != nil {
// 		twIDs := []string{tweetID}
// 		twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
// 		if err != nil {
// 			return listContext, errs.NewError(err)
// 		}

// 		if twitterDetail != nil {
// 			for k, v := range *twitterDetail {
// 				if strings.EqualFold(k, tweetID) {
// 					context := map[string]string{}
// 					context["user"] = v.User.UserName
// 					context["message"] = v.Tweet.NoteTweet.Text
// 					if context["message"] == "" {
// 						context["message"] = v.Tweet.Text
// 					}
// 					listContext = append([]map[string]string{context}, listContext...)

// 					isValid := true
// 					referencedTweets := v.ReferencedTweets
// 					i := 1
// 					for {
// 						if len(referencedTweets) > 0 {
// 							refTw := referencedTweets[0]
// 							contextRef := map[string]string{}
// 							contextRef["user"] = refTw.User.UserName
// 							contextRef["message"] = refTw.Tweet.NoteTweet.Text
// 							if contextRef["message"] == "" {
// 								contextRef["message"] = refTw.Tweet.Text
// 							}
// 							listContext = append([]map[string]string{contextRef}, listContext...)

// 							twIDRefs := []string{refTw.Tweet.ID}
// 							twitterDetailRef, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDRefs)
// 							if err != nil {
// 								isValid = false
// 							}

// 							if twitterDetailRef != nil {
// 								for kr, vr := range *twitterDetailRef {
// 									if strings.EqualFold(kr, refTw.Tweet.ID) {
// 										if len(vr.ReferencedTweets) > 0 {
// 											referencedTweets = vr.ReferencedTweets
// 										} else {
// 											isValid = false
// 										}
// 									}
// 								}
// 							} else {
// 								isValid = false
// 							}
// 						}

// 						i += 1
// 						if !isValid || i >= 4 {
// 							break
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return listContext, nil
// }

// func (s *Service) GenerateTokenImageBase64(ctx context.Context, tokenSymbol, tokenName, tokenDesc string) string {
// 	imagePrompt := fmt.Sprintf(`
// 		I want to create image for a token base on this info
// 		Token Symbol: %s
// 		Token name: %s
// 		Token Description: %s
// 	`, tokenSymbol, tokenName, tokenDesc)
// 	base64Str, _ := s.dojoAPI.GenerateImage(imagePrompt, s.conf.GenerateImageUrl)
// 	return base64Str
// }

func (s *Service) GenerateTokenImageBase64Gif(ctx context.Context, tokenSymbol, tokenName, tokenDesc string) string {
	imagePrompt := fmt.Sprintf(`
		I want to create image for a token base on this info
		Token Symbol: %s
		Token name: %s
		Token Description: %s
	`, tokenSymbol, tokenName, tokenDesc)
	base64Str, _ := s.dojoAPI.GenerateImage(imagePrompt, s.conf.GenerateGifImageUrl)
	return base64Str
}

func (s *Service) GetConversationIdByTweetID(tx *gorm.DB, tweetID string) string {
	conversationId := ""
	m, err := s.dao.FirstAgentTwitterPost(
		tx,
		map[string][]interface{}{
			"twitter_post_id = ? ": {tweetID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return tweetID
	}
	if m != nil {
		conversationId = m.TwitterConversationId
	}
	if conversationId == "" {
		twitterInfo, err := s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return tweetID
		}
		if twitterInfo != nil {
			tweetDetail, err := s.twitterWrapAPI.LookupTweetsByID(twitterInfo.AccessToken, tweetID)
			if err != nil {
				return tweetID
			}
			if tweetDetail != nil {
				conversationId = tweetDetail.ConversationID
			}
		}
	}
	if conversationId == "" {
		return tweetID
	}
	return conversationId
}

func (s *Service) GetPostTimeByTweetID(tx *gorm.DB, tweetID string) *time.Time {
	var postTime *time.Time
	m, err := s.dao.FirstAgentTwitterPost(
		tx,
		map[string][]interface{}{
			"twitter_post_id = ? ": {tweetID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		panic(err)
	}
	if m != nil {
		postTime = m.PostAt
	}
	if postTime == nil {
		twitterInfo, err := s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			panic(err)
		}
		if twitterInfo != nil {
			tweetDetail, _ := s.twitterWrapAPI.LookupTweetsByID(twitterInfo.AccessToken, tweetID)
			if tweetDetail != nil {
				createdAt, err := time.Parse(time.RFC3339, tweetDetail.CreatedAt)
				if err != nil {
					panic(err)
				}
				postTime = &createdAt
			}
		}
	}
	if postTime == nil {
		postTime = helpers.TimeNow()
	}
	return postTime
}

func (s *Service) TestVideo(ctx context.Context) {
	videoUrl := "https://gateway.lighthouse.storage/ipfs/bafybeia7y5xp74komdtmiisunemiod56tqhotglzkke4ym66tvx4ywz7u4"
	mediaID := ""
	var err error
	if videoUrl != "" {
		mediaID, err = s.twitterAPI.UploadVideo(models.GetImageUrl(videoUrl), []string{s.conf.TokenTwiterID})
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if mediaID != "" {
		twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)

		// post truc tiep reply, luu lai reply_id
		refId, err := helpers.ReplyTweetByToken(twitterInfo.AccessToken, "DONE", "1895369759690219863", mediaID)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(refId)
	}
}
