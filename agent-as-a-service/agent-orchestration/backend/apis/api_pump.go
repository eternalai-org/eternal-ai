package apis

import (
	"fmt"
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetPumpBalance(c *gin.Context) {
	apiKey := c.GetHeader("api-key")
	if apiKey == "" {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{
			Error: fmt.Errorf("api-key header is required"),
		})
		return
	}

	response, err := s.nls.GetPumpBalance(c.Request.Context(), apiKey)
	if err != nil {
		ctxJSON(c, http.StatusInternalServerError, &serializers.Resp{
			Error: fmt.Errorf("failed to get pump balance: %w", err),
		})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{
		Data: response,
	})
}

func (s *Server) GetListTrendingTokens(c *gin.Context) {
	ctx := s.requestContext(c)
	tokens, err := s.nls.GetListTrendingTokens(ctx)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewTrendingTokenRespArray(tokens)})
}

func (s *Server) GetPumpOrderHistory(c *gin.Context) {
	ctx := s.requestContext(c)
	apiKey := c.Request.Header.Get("api-key")
	orderHistory, err := s.nls.GetPumpOrderHistory(ctx, apiKey)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: orderHistory})
}

func (s *Server) NotifyChangePricePump(c *gin.Context) {
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}
