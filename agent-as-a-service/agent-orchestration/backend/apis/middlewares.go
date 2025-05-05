package apis

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/getsentry/raven-go"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const (
	CONTEXT_USER_DATA       = "context_user_data"
	CONTEXT_ERROR_DATA      = "context_error_data"
	CONTEXT_STACKTRACE_DATA = "context_stacktrace_data"
)

const (
	userIDKey         = "id"
	userEmailKey      = "email"
	otp               = "OTP"
	otpToken          = "OTPToken"
	otpPhone          = "OTPPhone"
	userClient        = "client"
	userPermission    = "permission"
	userData          = "userData"
	userCtxData       = "userCtxData"
	userEmail         = "userEmail"
	userId            = "userId"
	errorData         = "errorData"
	stacktraceData    = "stacktraceData"
	stacktraceExtra   = "stacktraceExtra"
	contextBodyLogFun = "contextBodyLogFun"
)

func ctxJSON(c *gin.Context, respCode int, resp interface{}) {
	if respCode != http.StatusOK {
		WrapRespError(c, resp)
	}
	c.JSON(respCode, resp)
}

func ctxSTRING(c *gin.Context, respCode int, resp string) {
	if respCode != http.StatusOK {
		WrapRespError(c, resp)
	}
	c.String(respCode, resp)
}

func ctxData(c *gin.Context, respCode int, contentType string, resp []byte) {
	if respCode != http.StatusOK {
		WrapRespError(c, resp)
	}
	c.Data(respCode, contentType, resp)
}

func ctxAbortWithStatusJSON(c *gin.Context, respCode int, resp interface{}) {
	if respCode != http.StatusOK {
		WrapRespError(c, resp)
	}
	c.AbortWithStatusJSON(respCode, resp)
}

func WrapRespError(c *gin.Context, resp interface{}) {
	var retErr *errs.Error
	switch resp.(type) {
	case *serializers.Resp:
		{
			retResp, ok := resp.(*serializers.Resp)
			if ok &&
				retResp.Error != nil {
				retResp.Error = errs.NewError(retResp.Error)
				retErr, _ = retResp.Error.(*errs.Error)
			}
		}
	}
	if retErr != nil {
		c.Set(CONTEXT_ERROR_DATA, retErr.Error())
		c.Set(CONTEXT_STACKTRACE_DATA, retErr.Stacktrace())
		if strings.ToLower(strings.TrimSpace(c.Query("stacktrace"))) != "true" {
			retErr.SetStacktrace("")
		}
	}
}

func (s *Server) requestContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

func (s *Server) getRequestIP(c *gin.Context) string {
	var ipStr string
	if ipStr == "" {
		ipStr = c.Request.Header.Get("X-Original-Forwarded-For")
	}
	if ipStr == "" {
		ipStr = c.Request.Header.Get("Cf-Connecting-Ip")
	}
	if ipStr == "" {
		ipStr = c.Request.Header.Get("ip")
	}
	if ipStr == "" {
		ipStr = c.ClientIP()
	}
	return ipStr
}

func (s *Server) getUserAgent(c *gin.Context) string {
	return c.Request.UserAgent()
}

func (s *Server) contextBVMUserExtra(c *gin.Context) string {
	var userData interface{}
	json.Unmarshal([]byte(c.Request.Header.Get("user-data")), &userData)
	return helpers.ConvertJsonString(
		map[string]interface{}{
			"ip":         s.getRequestIP(c),
			"ip_country": c.Request.Header.Get("Cf-IpLocation"),
			"agent":      c.Request.Header.Get("user-agent"),
			"user-data":  userData,
		},
	)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (s *Server) loggerDisabledBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("log_body", false)
		c.Next()
	}
}

func (s *Server) slackErrorBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("slack_error", true)
		c.Next()
	}
}

func (s *Server) logApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("log", true)
		c.Set("log_body", true)
		c.Set("slack_error", false)
		start := time.Now()
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		var bodyRequest string
		if (c.Request.Method == http.MethodPost ||
			c.Request.Method == http.MethodPut) &&
			strings.LastIndex(strings.ToLower(c.GetHeader("content-type")), "application/json") >= 0 {
			buf, bodyErr := ioutil.ReadAll(c.Request.Body)
			if bodyErr == nil {
				rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
				bodyRequest = string(buf)
				c.Request.Body = rdr2
			}
		}
		c.Next()
		if c.GetBool("log") {
			end := time.Now()
			latency := end.Sub(start).Seconds()
			ipStr := c.Request.Header.Get("ip")
			if ipStr == "" {
				ipStr = c.ClientIP()
			}
			var errText, stacktraceText, bodyResponse string
			v, ok := c.Get(CONTEXT_ERROR_DATA)
			if ok {
				errText = v.(string)
			}
			v, ok = c.Get(CONTEXT_STACKTRACE_DATA)
			if ok {
				stacktraceText = v.(string)
			}
			bodyResponse = bodyLogWriter.body.String()
			if !c.GetBool("log_body") {
				bodyRequest = ""
			}
			address := s.stringFromContextQuery(c, "address")
			if bodyResponse != "" {
				var rsBody struct {
					Address string `json:"address"`
				}
				if rsBody.Address != "" {
					address = rsBody.Address
				}
			}
			logger.Info(
				"api_response_time",
				"request info",
				zap.Any("referer", c.Request.Referer()),
				zap.Any("ip", ipStr),
				zap.Any("method", c.Request.Method),
				zap.Any("path", c.Request.URL.Path),
				zap.Any("raw_query", c.Request.URL.RawQuery),
				zap.Any("latency", latency),
				zap.Any("status", c.Writer.Status()),
				zap.Any("user_agent", c.Request.UserAgent()),
				zap.Any("platform", c.Request.Header.Get("platform")),
				zap.Any("os", c.Request.Header.Get("os")),
				zap.Any("country", c.Request.Header.Get("country")),
				zap.Any("error_text", errText),
				zap.Any("stacktrace", stacktraceText),
				zap.Any("body_request", helpers.SubStringBodyResponse(bodyRequest, 10000)),
				zap.Any("body_response", helpers.SubStringBodyResponse(bodyResponse, 10000)),
				zap.Any("address", address),
				zap.Any("ip_country", c.Request.Header.Get("CF-IPLocation")),
				zap.Any("connecting_ip", c.Request.Header.Get("Cf-Connecting-Ip")),
			)
			if os.Getenv("DEV") == "true" {
				fmt.Println(stacktraceText)
			}
		}
	}
}

func (s *Server) recoveryMiddleware(client *raven.Client, onlyCrashes bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			flags := map[string]string{
				"endpoint": c.Request.RequestURI,
			}
			if rval := recover(); rval != nil {
				rvalStr := fmt.Sprint(rval)
				client.CaptureMessage(rvalStr, flags, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)),
					raven.NewHttp(c.Request))
				ctxAbortWithStatusJSON(c, http.StatusInternalServerError, &serializers.Resp{
					Result: nil,
					Error:  errs.NewError(errors.New(rvalStr)),
				})
			}
			if !onlyCrashes {
				for _, item := range c.Errors {
					client.CaptureMessage(item.Error(), flags, &raven.Message{
						Message: item.Error(),
						Params:  []interface{}{item.Meta},
					},
						raven.NewHttp(c.Request))
				}
			}
		}()
		c.Next()
	}
}

func (s *Server) pagingFromContext(c *gin.Context) (int, int) {
	var (
		pageS  = c.DefaultQuery("page", "1")
		limitS = c.DefaultQuery("limit", "50")
		page   int
		limit  int
		err    error
	)

	page, err = strconv.Atoi(pageS)
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(limitS)
	if err != nil {
		limit = 50
	}

	if limit > 500 {
		limit = 500
	}

	return page, limit
}

func (s *Server) ipfsProxyMiddleware(hostPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.Param("hash")
		r := c.Request
		w := c.Writer
		director := func(req *http.Request) {
			hostURL, err := url.Parse(hostPath)
			if err != nil {
				ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
				return
			}
			req.URL.Scheme = hostURL.Scheme
			req.URL.Host = hostURL.Host
			req.Host = hostURL.Host
			req.URL.Path = hostPath + "/" + hash + ".mp4"
		}
		proxy := &httputil.ReverseProxy{
			Director: director,
		}
		proxy.ServeHTTP(w, r)
	}
}

func (s *Server) uintFromContextParam(c *gin.Context, param string) uint {
	val := strings.TrimSpace(c.Param(param))
	if val == "" {
		panic(errors.New("invalid param"))
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		panic(errors.New("invalid param"))
	}
	return uint(num)
}

func (s *Server) uint64FromContextParam(c *gin.Context, param string) (uint64, error) {
	val := strings.TrimSpace(c.Param(param))
	if val == "" {
		return uint64(0), nil
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(num), nil
}

func (s *Server) uintFromContextQuery(c *gin.Context, query string) uint {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return uint(0)
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(num)
}

func (s *Server) intFromContextQuery(c *gin.Context, query string) int {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return int(0)
	}
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func (s *Server) float64FromContextQuery(c *gin.Context, query string) (float64, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return 0, nil
	}
	num, err := strconv.ParseFloat(val, 10)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (s *Server) ptrUintFromContextQuery(c *gin.Context, query string) (*uint, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return nil, nil
	}

	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return nil, err
	}
	v := uint(num)
	return &v, nil
}

func (s *Server) uint64FromContextQuery(c *gin.Context, query string) (uint64, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return uint64(0), nil
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(num), nil
}

func (s *Server) stringFromContextQuery(c *gin.Context, query string) string {
	return strings.TrimSpace(c.Query(query))
}

func (s *Server) stringFromContextQueryDefault(c *gin.Context, query string, df string) string {
	return strings.TrimSpace(c.DefaultQuery(query, df))
}

func (s *Server) stringFromContextParam(c *gin.Context, query string) string {
	return strings.TrimSpace(c.Param(query))
}

func (s *Server) maxResultFromContextQuery(c *gin.Context) int {
	maxResults := s.intFromContextQuery(c, "max_results")
	if maxResults == 0 {
		maxResults = 25
	}
	if maxResults < 5 {
		maxResults = 5
	}
	return maxResults
}

func (s *Server) stringArrayFromContextQuery(c *gin.Context, query string) []string {
	val := strings.ToLower(strings.TrimSpace(c.Query(query)))
	if val == "" {
		return []string{}
	}
	return strings.Split(val, ",")
}

func (s *Server) uintArrayFromContextQuery(c *gin.Context, query string) ([]uint, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return []uint{}, nil
	}
	vals := strings.Split(val, ",")
	rets := []uint{}
	for _, val := range vals {
		num, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return []uint{}, err
		}
		rets = append(rets, uint(num))
	}
	return rets, nil
}

func (s *Server) intArrayFromContextQuery(c *gin.Context, query string) ([]int, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return []int{}, nil
	}
	vals := strings.Split(val, ",")
	rets := []int{}
	for _, val := range vals {
		num, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return []int{}, err
		}
		rets = append(rets, int(num))
	}
	return rets, nil
}

func (s *Server) dateFromContextQuery(c *gin.Context, query string) (*time.Time, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Server) timeFromContextQuery(c *gin.Context, query string) (*time.Time, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Server) boolFromContextQuery(c *gin.Context, query string) (*bool, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return nil, nil
	}
	ret, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Server) boolValueFromContextQuery(c *gin.Context, query string) (bool, error) {
	val := strings.TrimSpace(c.Query(query))
	if val == "" {
		return false, nil
	}
	ret, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func (s *Server) proxyMiddleware(prefixPath string, host string, headers map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		director := func(req *http.Request) {
			hostURL, err := url.Parse(host)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
				return
			}
			req.URL.Scheme = hostURL.Scheme
			req.URL.Host = hostURL.Host
			req.Host = hostURL.Host
			req.URL.Path = strings.TrimPrefix(req.URL.Path, prefixPath)
			query := req.URL.Query()
			req.URL.RawQuery = query.Encode()
			for k := range r.Header {
				v := c.GetHeader(k)
				req.Header.Set(k, v)
			}
			for k, v := range headers {
				req.Header.Set(k, v)
			}
			if os.Getenv("DEV") == "true" {
				fmt.Printf("%s -> %s\n", r.URL.String(), req.URL.String())
			}
		}
		proxy := &httputil.ReverseProxy{
			Director: director,
		}
		proxy.ServeHTTP(w, r)
	}
}

const (
	SCOPE_TRANSFER string = "constant-transfer"

	// this const is used to force the user re-login to get newest token
	VALID_TOKEN = "2019-03-25"

	PERMISSION_AUTHORIZE = "authorize"
	PERMISSION_TOKEN     = "token"
)

func (s *Server) bodyLogMiddleware(fn func(c *gin.Context, req string, res string) (string, string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(contextBodyLogFun, fn)
		c.Next()
	}
}

func (s *Server) loginBodyLogMiddleware() gin.HandlerFunc {
	return s.bodyLogMiddleware(func(c *gin.Context, req, res string) (string, string) {
		rs := struct {
			Email string
		}{}
		err := json.Unmarshal([]byte(req), &rs)
		if err != nil {
			return req, ""
		}
		{
			c.Set(userEmail, rs.Email)
		}
		reqBytes, err := json.Marshal(&rs)
		if err != nil {
			return "", ""
		}
		return string(reqBytes), ""
	})
}

func (s *Server) sortListFromContext(c *gin.Context) []string {
	var (
		sortCol  = c.DefaultQuery("sort_col", "")
		sortType = c.DefaultQuery("sort_type", "0")
		sortStr  []string
	)

	arrayCol := strings.Split(sortCol, ",")
	arrayType := strings.Split(sortType, ",")

	whiteListSortCol := map[string]string{
		"price":        "1",
		"market_cap":   "1",
		"created_at":   "1",
		"total_volume": "1",
		"percent":      "1",
		"reply_count":  "1",
		"last_reply":   "1",
	}

	for i, item := range arrayCol {
		item = strings.TrimSpace(item)
		if strings.EqualFold(item, "following") {
			item = "price"
		}

		_, ok := whiteListSortCol[item]
		if item != "" && ok {
			if strings.Contains(item, " ") {
				panic("bad request")
			}
			sortTypeStr := "desc"
			sortNum, _ := strconv.Atoi(arrayType[i])
			if sortNum > 0 {
				sortTypeStr = "asc"
			}

			sortStr = append(sortStr, fmt.Sprintf(`%s %s`, item, sortTypeStr))
		}
	}
	return sortStr
}

func (s *Server) chainFromContextQuery(c *gin.Context) uint64 {
	chain, err := s.uint64FromContextQuery(c, "chain")
	if err != nil {
		chain = 0
	}
	return chain
}

func (s *Server) internalApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		internalKey := s.stringFromContextQuery(c, "api_key")
		if internalKey == "" {
			internalKey = c.GetHeader("api-key")
			if internalKey == "" {
				authHeader := c.GetHeader("Authorization")
				internalKey = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if !strings.EqualFold(internalKey, s.conf.InternalApiKey) {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrInvalidApiKey)})
			return
		}
		c.Next()
	}
}

func (s *Server) externalApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		externalKey := strings.TrimPrefix(authHeader, "TK1 ")
		if externalKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (s *Server) agentSortListFromContext(c *gin.Context) []string {
	var (
		sortCol  = c.DefaultQuery("sort_col", "")
		sortType = c.DefaultQuery("sort_type", "0")
		sortStr  []string
	)

	arrayCol := strings.Split(sortCol, ",")
	arrayType := strings.Split(sortType, ",")

	whiteListSortCol := map[string]string{
		"meme_market_cap":     "1",
		"meme_percent":        "1",
		"reply_latest_time":   "1",
		"meme_price":          "1",
		"meme_volume_last24h": "1",
		"created_at":          "1",
		"prompt_calls":        "1",
		"installed_count":     "1",
		"recent_chat_time":    "1",
	}

	for i, item := range arrayCol {
		item = strings.TrimSpace(item)
		_, ok := whiteListSortCol[item]
		if item != "" && ok {
			if strings.Contains(item, " ") {
				panic("bad request")
			}
			sortTypeStr := "desc"
			sortNum, _ := strconv.Atoi(arrayType[i])
			if sortNum > 0 {
				sortTypeStr = "asc"
			}

			sortStr = append(sortStr, fmt.Sprintf(`%s %s`, item, sortTypeStr))
		}
	}
	return sortStr
}

func (s *Server) authCheckTK1TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		authHeader = strings.TrimPrefix(authHeader, "TK1 ")
		if authHeader == "" {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		authData, err := helpers.DecryptAndVerifyAuthToken(authHeader, s.conf.EncryptAuthenKey)
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		if authData != nil {
			c.Set("userData", &models.User{
				Address: authData.Address,
			})
		} else {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		c.Next()
	}
}

func (s *Server) getUserAddressFromTK1Token(c *gin.Context) (string, error) {
	authToken := c.GetHeader("Authorization")
	authToken = strings.TrimPrefix(authToken, "TK1 ")
	authData, err := helpers.DecryptAndVerifyAuthToken(authToken, s.conf.EncryptAuthenKey)
	if err != nil {
		return "", err
	}
	return authData.Address, nil
}

func (s *Server) authCheckSignatureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAddress := c.GetHeader("XXX-Address")
		userMessage := c.GetHeader("XXX-Message")
		signature := c.GetHeader("XXX-Signature")

		if signature == "" || userAddress == "" || userMessage == "" {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		i, err := strconv.ParseInt(userMessage, 10, 64)
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		timeSignature := helpers.TimeFromUnix(i)
		timeValid := helpers.TimeNow().Add(-15 * time.Minute)
		if timeSignature.Before(timeValid) {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		err = s.nls.VerifyAddressSignature(context.Background(), models.ETHEREUM_CHAIN_ID, userAddress, userMessage, signature)
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusUnauthorized, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
			return
		}

		c.Set("userData", &models.User{
			Address: strings.ToLower(userAddress),
		})

		c.Next()
	}
}

func (s *Server) getUserAddress(c *gin.Context) string {
	return c.GetHeader("XXX-Address")
}

func (s *Server) middlewareApiLimit(numRequest int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := s.getRequestIP(c)
		apiPath := c.Request.URL.Path
		limitKey := fmt.Sprintf("api_limit_%s_%s_%s", c.Request.Method, ip, apiPath)
		var limit int
		err := s.nls.GetRedisCachedWithKey(limitKey, &limit)
		if err != nil {
			// Only set limit to 0 if key doesn't exist, not on other redis errors
			if err.Error() == "redis: nil" {
				limit = 0
			} else {
				ctxAbortWithStatusJSON(c, http.StatusInternalServerError, &serializers.Resp{Error: errs.NewError(err)})
				return
			}
		}
		if limit >= numRequest {
			ctxAbortWithStatusJSON(c, http.StatusTooManyRequests, &serializers.Resp{Error: errs.NewError(errors.New("too many requests"))})
			return
		}

		// Calculate expiration time from start of current window
		now := time.Now()
		windowStart := now.Truncate(duration)
		expireTime := windowStart.Add(duration)

		// Only increment if we can set the value
		err = s.nls.SetRedisCachedWithKey(limitKey, limit+1, time.Until(expireTime))
		if err != nil {
			ctxAbortWithStatusJSON(c, http.StatusInternalServerError, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		c.Next()
	}
}
