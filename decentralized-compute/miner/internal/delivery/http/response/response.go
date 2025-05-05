package response

import (
	"context"
	"encoding/json"
	"net/http"
	"solo/pkg/logger"

	"github.com/gorilla/mux"
)

type IResponse interface {
	SetID(string)
	GetID() string
}

type BaseResponse struct {
	ID string `json:"id"`
}

func (p *BaseResponse) SetID(ID string) {
	p.ID = ID
}

func (p BaseResponse) GetID() string {
	return p.ID
}

type IHttpResponse interface {
	RespondWithError(w http.ResponseWriter, httpCode int, appCode int, payload error)
	RespondSuccess(w http.ResponseWriter, httpCode int, appCode int, payload interface{}, customerMessage string)
	RespondWithoutContainer(w http.ResponseWriter, httpCode int, payload interface{})
}

type JsonResponse struct {
	Error  *RespondErr `json:"error"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type RespondErr struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"code"`
}

type httpResponse struct{}

type StreamResponse struct {
	Data        interface{}
	IsNotStream bool
}

func NewHttpResponse() *httpResponse {
	return new(httpResponse)
}

func (h *httpResponse) RespondWithError(w http.ResponseWriter, httpCode int, appCode int, payload error) {
	h.respondWithJSON(w, payload, httpCode, appCode, payload, "")
}

func (h *httpResponse) RespondSuccess(w http.ResponseWriter, httpCode int, appCode int, payload interface{}, customerMessage string) {
	h.respondWithJSON(w, nil, httpCode, appCode, payload, customerMessage)
}

func (h *httpResponse) respondWithJSON(w http.ResponseWriter, respErr error, httpCode int, appCode int, payload interface{}, _ string) {
	code := ResponseMessage[appCode].Code

	jsr := JsonResponse{Data: payload, Status: true}

	if respErr != nil {
		errMessage := &RespondErr{}
		errMessage.Message = respErr.Error()
		errMessage.ErrorCode = code
		jsr.Error = errMessage
		jsr.Status = false
	}

	response, _ := json.Marshal(jsr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	if _, err := w.Write(response); err != nil {
		logger.AtLog.Fatal(err)
	}
}

func (h *httpResponse) RespondWithoutContainer(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	if _, err := w.Write(response); err != nil {
		logger.AtLog.Fatal(err)
	}
}

// HandlerFunc --
type HandlerFunc func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error)

type restHandlerTemplate struct {
	handlerFunc HandlerFunc
	httpResp    *httpResponse
}

type streamHandlerTemplate struct {
	handlerFunc HandlerFunc
	httpResp    *httpResponse
}

func NewRESTHandlerTemplate(handlerFunc HandlerFunc) http.Handler {
	return &restHandlerTemplate{
		handlerFunc: handlerFunc,
		httpResp:    NewHttpResponse(),
	}
}

func NewStreamHandlerTemplate(handlerFunc HandlerFunc) http.Handler {
	return &streamHandlerTemplate{
		handlerFunc: handlerFunc,
		httpResp:    NewHttpResponse(),
	}
}

// user Devices server has prefix with "8"
const (
	//Success has prefix with 8"0"
	Success = 1

	//Error has prefix with 8"1"
	Error = -1
)

// Message
var ResponseMessage = map[int]struct {
	Code    int
	Message string
}{
	Success: {Success, "Success"},
	Error:   {Error, "Failed."},
}

func (h *restHandlerTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	vars := mux.Vars(r)

	item, err := h.handlerFunc(ctx, r, vars)
	if err != nil {
		h.httpResp.RespondWithError(w, http.StatusBadRequest, Error, err)
		return
	}

	if item == nil {
		item = map[string]interface{}{
			"status": true,
		}
	}

	h.httpResp.RespondSuccess(w, http.StatusOK, Success, item, "")
}

func (h *streamHandlerTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	item, err := h.handlerFunc(ctx, r, vars)
	resp := item.(StreamResponse)
	if resp.IsNotStream {
		if err != nil {
			h.httpResp.RespondWithError(w, http.StatusBadRequest, Error, err)
			return
		}
		if item == nil {
			item = map[string]interface{}{
				"status": true,
			}
		}
		h.httpResp.RespondSuccess(w, http.StatusOK, Success, resp.Data, "")
	}

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r) // Call the next handler
	})
}
