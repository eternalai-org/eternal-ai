package apis

import (
	"context"
	"fmt"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Server) AdminTriggerJob(c *gin.Context) {
	request := &struct {
		JobName string `json:"job_name"`
	}{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if request.JobName == "" {
		c.JSON(400, gin.H{"error": "job_name is required"})
		return
	}

	var err error

	switch request.JobName {
	case "JobScanAgentTwitterPostForGenerateVideo":
		err = s.nls.JobScanAgentTwitterPostForGenerateVideo(context.Background())
	case "JobAgentTwitterPostSubmitVideoInfer":
		err = s.nls.JobAgentTwitterPostSubmitVideoInfer(context.Background())
	case "JobAgentTwitterScanResultGenerateVideo":
		err = s.nls.JobAgentTwitterScanResultGenerateVideo(context.Background())
	case "JobAgentTwitterPostGenerateVideo":
		err = s.nls.JobAgentTwitterPostGenerateVideo(context.Background())

	default:
		c.JSON(400, gin.H{"error": "job_name is not implemented at admin internal"})
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "job triggered"})
}

func (s *Server) AdminHandleVideo(c *gin.Context) {
	var request struct {
		TweetID string `json:"tweet_id"`
	}

	request.TweetID = s.stringFromContextQuery(c, "tweet_id")

	if request.TweetID == "" {
		c.JSON(400, gin.H{"error": "tweet_id is required"})
		return
	}

	handleRequest := &services.HandleGenerateVideoRequest{
		TweetID:            request.TweetID,
		AgentInfoMentionID: s.conf.VideoAiAgentInfoId,
	}

	isHandle := s.stringFromContextQuery(c, "is_handle")
	if isHandle == "100" {
		decideHandle := true
		handleRequest.DecideToHandle = &decideHandle
	}

	agentTwitterPost, err := s.nls.HandleGenerateVideoWithSpecificTweet(daos.GetDBMainCtx(context.Background()), handleRequest, "admin_api")
	if err != nil {
		logger.Error("handle_video", "internal_api_key error", zap.Error(err))
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if agentTwitterPost == nil {
		c.JSON(200, gin.H{"message": "agent twitter post is nil"})
		return
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("SUCCESS insert agent twitter post id: %d", agentTwitterPost.ID)})
	return
}
