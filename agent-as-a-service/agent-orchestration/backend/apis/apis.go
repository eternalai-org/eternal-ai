package apis

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/url"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) ClearCacheKey(c *gin.Context) {
	key := s.stringFromContextQuery(c, "key")
	prefix, err := s.boolValueFromContextQuery(c, "prefix")
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	if prefix {
		err = s.nls.DeleteRedisCachedWithPrefix(key)
	} else {
		err = s.nls.DeleteRedisCachedWithKey(key)
	}
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) DisableJobs() {
	s.nls.DisableJobs()
}

func (s *Server) RunJobs() {
	go s.nls.RunJobs(context.Background())
}

func (s *Server) RunTeleBotJob() {
	go s.nls.RunTeleBotJob(context.Background())
}

func (s *Server) TwitterOauthCallback(c *gin.Context) {
	ctx := s.requestContext(c)
	callbackUrl := s.stringFromContextQuery(c, "callback")
	code := s.stringFromContextQuery(c, "code")
	address := s.stringFromContextQuery(c, "address")
	agentID := s.stringFromContextQuery(c, "agent_id")
	clientID := s.stringFromContextQuery(c, "client_id")

	mapResp, err := s.nls.TwitterOauthCallbackV1(ctx, callbackUrl, address, code, agentID, clientID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	if mapResp != nil {
		callbackUrl = helpers.BuildUri(
			callbackUrl,
			mapResp,
		)
	}

	location := url.URL{Path: callbackUrl}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func (s *Server) TwitterOauthCallbackForInternalData(c *gin.Context) {
	ctx := s.requestContext(c)
	callbackUrl := s.stringFromContextQuery(c, "callback")
	code := s.stringFromContextQuery(c, "code")

	err := s.nls.TwitterOauthCallbackForInternalData(ctx, callbackUrl, code)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	location := url.URL{Path: callbackUrl}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func (s *Server) GetEAISupplyTotal(c *gin.Context) {
	ctxJSON(c, http.StatusOK, 1000000000)
}

func (s *Server) GetEAISupplyCirculating(c *gin.Context) {
	ctx := s.requestContext(c)
	supply := big.NewFloat(1000000000)
	balance, err := s.nls.GetEthereumClient(ctx, models.ETERNAL_AI_CHAIN_ID).Balance("0x6fe0B7d6Ea2597946CCbbD198A79C95ADbb7228E")
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	fBalance := models.ConvertWeiToBigFloat(balance, 18)
	if fBalance.Cmp(big.NewFloat(0)) == 0 {
		fBalance = big.NewFloat(798099380)
	}
	circulatingSupply := models.SubBigFloats(supply, fBalance)
	resp, _ := circulatingSupply.Float64()
	ctxJSON(c, http.StatusOK, resp)
}

func (s *Server) GetAllConfigsExplorer(c *gin.Context) {
	cfs, err := s.nls.GetAllConfigsExplorer(
		s.requestContext(c),
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: cfs})
}

func (s *Server) GetTokenPrice(c *gin.Context) {
	mapTokenPrice := s.nls.GetMapTokenPrice(c)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: mapTokenPrice})
}

func (s *Server) GetListBubbleCrypto(c *gin.Context) {

	resp, err := http.Get("https://cryptobubbles.net/backend/data/bubbles1000.usd.json")
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	result := make([]map[string]interface{}, 0)
	for _, item := range data {
		name, ok2 := item["name"].(string)
		symbol, ok4 := item["symbol"].(string)
		rank, ok7 := item["rank"].(float64)
		price, ok9 := item["price"].(float64)
		marketcap, ok10 := item["marketcap"].(float64)
		volume, ok11 := item["volume"].(float64)
		performance, ok13 := item["performance"].(map[string]interface{})

		if ok2 && ok4 && ok7 && ok9 && ok10 && ok11 && ok13 {
			result = append(result, map[string]interface{}{
				"name":        name,
				"symbol":      symbol,
				"rank":        rank,
				"price":       price,
				"marketcap":   marketcap,
				"volume":      volume,
				"performance": performance,
			})
		}
	}

	if len(result) > 100 {
		result = result[:100]
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: result})
}

func (s *Server) MemeEventsByTransaction(c *gin.Context) {
	ctx := s.requestContext(c)
	txHash := s.stringFromContextQuery(c, "tx_hash")
	networkID, err := s.uint64FromContextQuery(c, "network_id")
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err = s.nls.MemeEventsByTransaction(ctx, networkID, txHash)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}
