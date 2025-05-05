package apis

import (
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) AddVibeWhiteList(c *gin.Context) {
	ctx := s.requestContext(c)
	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	_ = s.nls.AddVibeWhiteList(ctx, req.Email)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) VibeValidateReferralCode(c *gin.Context) {
	ctx := s.requestContext(c)
	var req struct {
		RefCode     string `json:"ref_code"`
		UserAddress string `json:"user_address"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	valid, err := s.nls.ValidateReferralCode(ctx, req.RefCode, req.UserAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Result: valid, Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: valid})
}

func (s *Server) AgentComment(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.uintFromContextParam(c, "id")
	req := &serializers.AgentCommentReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.AgentComment(ctx, userAddress, agentID, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) GetListAgentComment(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	agentID := s.uintFromContextParam(c, "id")
	agentUserComments, err := s.nls.GetListAgentComment(ctx, agentID, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentUserCommentArray(agentUserComments)})
}

func (s *Server) GetVibeDashBoards(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	chain := s.chainFromContextQuery(c)
	sortStr := s.agentSortListFromContext(c)
	search := s.stringFromContextQuery(c, "search")
	categoryIds := s.stringArrayFromContextQuery(c, "category_ids")
	agentTypesInt, _ := s.intArrayFromContextQuery(c, "agent_types")
	contractAddresses := s.stringArrayFromContextQuery(c, "contract_addresses")
	ids, _ := s.uintArrayFromContextQuery(c, "ids")
	exludeIds, _ := s.uintArrayFromContextQuery(c, "exlude_ids")
	includeHidden, _ := s.boolFromContextQuery(c, "include_hidden")
	userAddress, _ := s.getUserAddressFromTK1Token(c)
	ms, count, err := s.nls.GetVibeDashboard(ctx,
		includeHidden,
		contractAddresses,
		userAddress, chain, agentTypesInt, search,
		ids, exludeIds, categoryIds, sortStr, page, limit)

	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfoRespArry(ms), Count: &count})
}

func (s *Server) GetVibeDashBoardsDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.stringFromContextParam(c, "agent_id")
	ms, err := s.nls.GetVibeDashboardsDetail(ctx, agentID)
	if err != nil || ms == nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfoResp(ms)})
}

func (s *Server) VibeTokenGetDeployInfo(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.uintFromContextParam(c, "agent_id")
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	resp, err := s.nls.VibeTokenGetDeployInfo(ctx, userAddress, agentID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}
