package apis

import (
	"fmt"
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

// func (s *Server) InfraTwitterAppAuthenInstall(c *gin.Context) {
// 	ctx := s.requestContext(c)
// 	authUrl, err := s.nls.InfraTwitterAppAuthenInstall(
// 		ctx,
// 		s.stringFromContextQuery(c, "install_code"),
// 		s.stringFromContextQuery(c, "install_uri"),
// 	)
// 	if err != nil {
// 		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}
// 	c.Redirect(http.StatusFound, authUrl)
// }

func (s *Server) InfraTwitterAppAuthenCallback(c *gin.Context) {
	ctx := s.requestContext(c)
	_, err := s.nls.InfraTwitterAppAuthenCallback(
		ctx, s.stringFromContextQuery(c, "address"),
		s.stringFromContextQuery(c, "code"),
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("https://eternalai.org/%d", s.conf.InfraTwitterApp.AgentID))
}

func (s *Server) UtilityPostTwitter(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AgentUtilityTwitterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress := s.getUserAddress(c)
	resp, err := s.nls.UtilityPostTwitter(ctx, userAddress, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err), Result: &resp})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

// func (s *Server) UtilityTwitterVerifyDeposit(c *gin.Context) {
// 	ctx := s.requestContext(c)
// 	var req serializers.AgentUtilityTwitterReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}
// 	userAddress := s.getUserAddress(c)
// 	resp, err := s.nls.UtilityTwitterVerifyDeposit(ctx, userAddress, req.TxHash)
// 	if err != nil {
// 		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err), Result: &resp})
// 		return
// 	}
// 	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
// }

func (s *Server) InfraTwitterAppSearchRecentTweet(c *gin.Context) {
	ctx := s.requestContext(c)
	query := s.stringFromContextQuery(c, "query")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.InfraTwitterAppSearchRecentTweet(ctx, query, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetInfraTwitterAppInfo(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress := s.getUserAddress(c)
	info, err := s.nls.GetInfraTwitterAppInfo(ctx, userAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: *serializers.NewInfraTwitterAppInfoResp(info)})
}
