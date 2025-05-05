package apis

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/moralis"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/opensea"
	"github.com/gin-gonic/gin"
)

func (s *Server) AgentCreateAgentStudio(c *gin.Context) {
	ctx := s.requestContext(c)
	req := &serializers.StudioReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.AgentCreateAgentStudio(ctx, userAddress, req.GraphData)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantRespArry(resp)})
}

func (s *Server) AgentUpdateAgentStudio(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.StudioReq

	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	agentID := s.stringFromContextParam(c, "id")
	resp, err := s.nls.AgentUpdateAgentStudio(ctx, userAddress, agentID, req.GraphData)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantResp(resp)})
}

func (s *Server) AgentCreateAgentAssistant(c *gin.Context) {
	ctx := s.requestContext(c)
	req := &serializers.AssistantsReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.AgentCreateAgentAssistant(ctx, userAddress, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantResp(resp)})
}

func (s *Server) AgentUpdateAgentAssistant(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AssistantsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.AgentUpdateAgentAssistant(ctx, userAddress, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantResp(resp)})
}

func (s *Server) AgentUpdateAgentAssistantInContract(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.UpdateAgentAssistantInContractRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.UpdateAgentInfoInContract(ctx, userAddress, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantResp(resp)})
}

func (s *Server) UpdateTwinStatus(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.UpdateTwinStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	resp, err := s.nls.UpdateTwinStatus(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantResp(resp)})
}

func (s *Server) UploadDataToLightHouse(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.DataUploadToLightHouse

	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.UploadDataToLightHouse(ctx, userAddress, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Data: resp})
}

func (s *Server) GetNftOpenseaCollections(c *gin.Context) {
	ctx := s.requestContext(c)
	chain := s.stringFromContextQuery(c, "chain")
	creatorUsername := s.stringFromContextQuery(c, "creator_username")
	includeHidden := s.stringFromContextQuery(c, "include_hidden")
	limit := s.stringFromContextQuery(c, "limit")
	next := s.stringFromContextQuery(c, "next")
	orderBy := s.stringFromContextQuery(c, "order_by")
	inscriptionStr := s.stringFromContextQuery(c, "inscription")

	f := opensea.OpenSeaFilterCollections{}
	if chain != "" {
		f.Chain = chain
	}

	if creatorUsername != "" {
		f.CreatorUsername = creatorUsername
	}

	if includeHidden != "" {
		_includeHidden, err1 := strconv.ParseBool(includeHidden)
		if err1 == nil {
			f.IncludeHidden = &_includeHidden
		}

	}

	if limit != "" {
		_limit, err1 := strconv.Atoi(limit)
		if err1 == nil {
			f.Limit = _limit
		}

	}

	if next != "" {
		f.Next = next
	}

	if orderBy != "" {
		f.OrderBy = orderBy
	}

	if inscriptionStr != "" {
		inscription, err1 := strconv.ParseBool(inscriptionStr)
		if err1 == nil {
			f.Inscription = &inscription
		}
	}
	result, err := s.nls.OpenseaCollections(ctx, f)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Data: result})
}

func (s *Server) GetNftCollectionsDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	inscription := s.stringFromContextQuery(c, "inscription")
	address := s.stringFromContextParam(c, "address")

	if inscription == "true" {
		result, err := s.nls.GetInscriptonInfo(ctx, address, opensea.OpenSeaFilterCollections{})
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		ctxJSON(c, http.StatusOK, &serializers.Resp{Data: result})
	} else {
		cursor := c.Query("cursor")
		pageSize := c.Query("limit")

		result, err := s.nls.GetNftCollectionMetadataByContract(ctx, address, cursor, pageSize, moralis.MoralisFilter{})
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		ctxJSON(c, http.StatusOK, &serializers.Resp{Data: result})
	}
}

func (s *Server) GetNftCollectionsByTokenID(c *gin.Context) {
	ctx := s.requestContext(c)
	inscription := s.stringFromContextQuery(c, "inscription")
	address := s.stringFromContextParam(c, "address")
	tokenId := s.stringFromContextParam(c, "tokenId")

	if inscription == "true" {
		if !strings.Contains(tokenId, "i0") {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
			return
		}
		result, err := s.nls.GetInscriptonInfoByTokenID(ctx, address, tokenId, opensea.OpenSeaFilterCollections{})
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		ctxJSON(c, http.StatusOK, &serializers.Resp{Data: result})
	} else {

		result, err := s.nls.GetNftCollectionMetadataByTokenID(ctx, address, tokenId, moralis.MoralisFilter{})
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		ctxJSON(c, http.StatusOK, &serializers.Resp{Data: result})
	}
}

func (s *Server) AgentCreateAgentAssistantForLocal(c *gin.Context) {
	ctx := s.requestContext(c)
	req := &serializers.AssistantsReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	resp, err := s.nls.AgentCreateAgentAssistantForLocal(ctx, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAssistantResp(resp)})
}

func (s *Server) MarkInstalledUtilityAgent(c *gin.Context) {
	ctx := s.requestContext(c)
	req := &serializers.AgentActionReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.MarkInstalledUtilityAgent(ctx, userAddress, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) MarkRecentChatUtilityAgent(c *gin.Context) {
	ctx := s.requestContext(c)
	req := &serializers.AgentActionReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.MarkRecentChatUtilityAgent(ctx, userAddress, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) MarkPromptCountUtilityAgent(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.uintFromContextParam(c, "id")

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.MarkPromptCountUtilityAgent(ctx, userAddress, agentID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) LikeAgent(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.uintFromContextParam(c, "id")

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.LikeAgent(ctx, userAddress, agentID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) CheckAgentLiked(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.uintFromContextParam(c, "id")

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.CheckAgentLiked(ctx, userAddress, agentID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) PublicAgent(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.uintFromContextParam(c, "id")

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	resp, err := s.nls.PublicAgent(ctx, userAddress, agentID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}
