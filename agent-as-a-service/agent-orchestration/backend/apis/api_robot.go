package apis

import (
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	blockchainutils "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/blockchain_utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) GenerateRobotSaleWallet(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.RobotSaleWalletReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	robotSaleWallet, err := s.nls.GenerateRobotSaleWallet(ctx, req.ProjectID, req.UserAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewRobotSaleWalletResp(robotSaleWallet)})
}

func (s *Server) GetRobotSaleWallet(c *gin.Context) {
	ctx := s.requestContext(c)
	projectID := s.stringFromContextQuery(c, "project_id")
	userAddress := s.stringFromContextQuery(c, "user_address")
	robotSaleWallet, err := s.nls.GetRobotSaleWallet(ctx, projectID, userAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewRobotSaleWalletResp(robotSaleWallet)})
}

func (s *Server) GetRobotProject(c *gin.Context) {
	ctx := s.requestContext(c)
	projectID := s.stringFromContextQuery(c, "project_id")
	robotProject, err := s.nls.GetRobotProject(ctx, projectID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewRobotProjectResp(robotProject)})
}

func (s *Server) RobotCreateToken(c *gin.Context) {
	ctx := s.requestContext(c)
	var req blockchainutils.SolanaCreateTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	robotProject, err := s.nls.RobotCreateToken(ctx, req.ProjectID, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewRobotProjectResp(robotProject)})
}

func (s *Server) RobotTransferToken(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.RobotTokenTransferReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	transfer, err := s.nls.RobotTransferToken(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewRobotTokenTransferResp(transfer)})
}

func (s *Server) GetRobotProjectLeaderBoards(c *gin.Context) {
	ctx := s.requestContext(c)
	projectID := s.stringFromContextQuery(c, "project_id")
	userAddress := s.stringFromContextQuery(c, "user_address")
	page, limit := s.pagingFromContext(c)
	lstResp, err := s.nls.GetRobotProjectLeaderBoards(ctx, userAddress, projectID, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewRobotSaleWalletRespList(lstResp)})
}
