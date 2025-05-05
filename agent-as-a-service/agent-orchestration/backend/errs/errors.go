package errs

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/getsentry/raven-go"
	"go.uber.org/zap"
)

var (
	ErrInvalidApiKey         = &Error{Code: -1000, Message: "Invalid API Key."}
	ErrSystemError           = &Error{Code: -1001, Message: "Something went wrong. Please try again."}
	ErrInvalidCredentials    = &Error{Code: -1002, Message: "Invalid credentials. Please try again"}
	ErrBadRequest            = &Error{Code: -1003, Message: "Something went wrong. Please try again."}
	ErrBadContent            = &Error{Code: -1004, Message: "Something went wrong. Please try again."}
	ErrInvalidRecaptcha      = &Error{Code: -1005, Message: "Invalid reCAPTCHA. Please refresh the page and try again."}
	ErrPermissionDenied      = &Error{Code: -1006, Message: "Something went wrong. Please try again."}
	ErrUserNotFound          = &Error{Code: -1007, Message: "Something went wrong. Please try again."}
	ErrAuthorizationExistsed = &Error{Code: -1008, Message: "Something went wrong. Please try again."}
	ErrInvalidSignature      = &Error{Code: -1009, Message: "Invalid signature. Please check your wallet address and try again."}
	ErrRecordNotFound        = &Error{Code: -1010, Message: "Something went wrong. Please try again."}
	ErrTwitterIdNotFound     = &Error{Code: -1011, Message: "Something went wrong. Please try again."}
	ErrRewardNotFound        = &Error{Code: -1012, Message: "Something went wrong. Please try again."}
	ErrUnAuthorization       = &Error{Code: -1013, Message: "Something went wrong. Please try again."}
	ErrInvalidRequest        = &Error{Code: -1014, Message: "Invalid request. Please try again."}
	ErrInternalServerError   = &Error{Code: -1015, Message: "Something went wrong. Please try again."}

	ErrAlreadyPurchase         = &Error{Code: -2001, Message: "Already purchase apps. Please try again."}
	ErrAppNotFound             = &Error{Code: -2002, Message: "Something went wrong. Please try again."}
	ErrPurchaseError           = &Error{Code: -2003, Message: "Something went wrong. Please try again."}
	ErrInvalidOwner            = &Error{Code: -2004, Message: "Invalid Owner. Please try again."}
	ErrAppNotInstalled         = &Error{Code: -2005, Message: "This app haven't installed. Please try again."}
	ErrReferralCodeExistsed    = &Error{Code: -2006, Message: "Something went wrong. Please try again."}
	ErrTokenNotFound           = &Error{Code: -2007, Message: "Something went wrong. Please try again."}
	ErrNetworkNotFound         = &Error{Code: -2008, Message: "Something went wrong. Please try again."}
	ErrTxHashExisted           = &Error{Code: -2009, Message: "TxHash already existed"}
	ErrNameExisted             = &Error{Code: -2010, Message: "Name already existed"}
	ErrUserNotExist            = &Error{Code: -2008, Message: "Something went wrong. Please try again."}
	ErrBadBalance              = &Error{Code: -2011, Message: "Insufficient balance. Please check your balance or open orders."}
	ErrQuestionDuplicate       = &Error{Code: -2012, Message: "Question already existed. Please try again."}
	ErrPostExisted             = &Error{Code: -2013, Message: "Post already existed. Please try again."}
	ErrAgentNotFound           = &Error{Code: -2014, Message: "Agent Not Found"}
	ErrTwitterIDExistsed       = &Error{Code: -2015, Message: "TwitterID Existsed"}
	ErrTwitterUsernameNotFound = &Error{Code: -2015, Message: "Twitter Username Not Found"}
	ErrInsufficientBalance     = &Error{Code: -2016, Message: "Insufficient Balance"}
	ErrApiUrlNotHealth         = &Error{Code: -2016, Message: "Api Url not Health"}

	ErrApiKeyRateLimited = &Error{Code: -2016, Message: "API Key rate limit"}
	ErrReferralCodeUsed  = &Error{Code: -2017, Message: "Referral code already used"}

	ErrAgentUtilityNotAuthen       = &Error{Code: -3000, Message: "Please authentican with this link and try again"}
	ErrAgentUtilityInvalidBalance  = &Error{Code: -3001, Message: "Insufficient balance! You need at least 1 $EAI to proceed with address {address} in Base."}
	ErrAgentUtilityPostTweetFailed = &Error{Code: -3002, Message: "Post tweet failed. Please try again."}
	ErrAgentUtilitySystemError     = &Error{Code: -3003, Message: "Something went wrong. Please try again."}
	ErrAgentUtilityNotFound        = &Error{Code: -3004, Message: "Please authentican with this link and try again"}
)

type Error struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Trace      string `json:"trace"`
	stacktrace string
	extra      []interface{}
}

func (e *Error) SetStacktrace(stacktrace string) {
	e.Trace = stacktrace
	e.stacktrace = stacktrace
}

func (e *Error) Stacktrace() string {
	return e.stacktrace
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) SetExtra(extra []interface{}) {
	e.extra = extra
}

func (e *Error) Extra() []interface{} {
	return e.extra
}

func (e *Error) ExtraJson() string {
	return helpers.ConvertJsonString(e.extra)
}

func NewErrorWithId(err error, id interface{}) error {
	if err != nil {
		msg := err.Error()
		err = NewError(err)
		err.(*Error).Message = fmt.Sprintf("%v : %s", id, msg)
	}
	return err
}

func NewError(err error, extras ...interface{}) error {
	if err == nil {
		return nil
	}
	_, ok := err.(*Error)
	if ok {
		sterr := err.(*Error).Stacktrace()
		retErr := &Error{
			Code:    err.(*Error).Code,
			Message: err.(*Error).Message,
		}
		if sterr == "" {
			retErr.SetStacktrace(fmt.Sprintf("%s\n\n%s", err.Error(), NewStacktraceString(extras...)))
			err.(*Error).SetExtra(extras)
		} else {
			retErr.SetStacktrace(sterr)
		}
		return retErr
	}
	retErr := &Error{
		Code:    ErrSystemError.Code,
		Message: err.Error(),
	}
	retErr.SetStacktrace(fmt.Sprintf("%s\n\n%s", err.Error(), NewStacktraceString(extras...)))
	return retErr
}

func NewTwitterError(err error, extras ...interface{}) error {
	return NewError(err, extras...)
}

func NewStacktraceString(extras ...interface{}) string {
	var rets []string
	if len(extras) > 0 {
		rets = append(rets, fmt.Sprintf("Extras -> %s", helpers.ConvertJsonString(extras)))
	}
	st := raven.NewStacktrace(1, 3, nil)
	for i := len(st.Frames) - 1; i >= 0; i-- {
		frame := st.Frames[i]
		if strings.TrimSpace(frame.Filename) != "" {
			rets = append(rets, fmt.Sprintf("%s\t%s\t%d", frame.Filename, frame.Function, frame.Lineno))
			rets = append(rets, fmt.Sprintf("\t%s", strings.Join(frame.PreContext, "\n\t")))
			rets = append(rets, fmt.Sprintf("%d.\t%s", frame.Lineno, frame.ContextLine))
			rets = append(rets, fmt.Sprintf("\t%s", strings.Join(frame.PostContext, "\n\t")))
		}
	}
	return strings.Join(rets, "\n")
}

func MergeError(err1 error, errss ...error) error {
	var msgs, sterrs []string
	if err1 != nil {
		err1 = NewError(err1)
		_, ok := err1.(*Error)
		if ok {
			msgs = append(msgs, strings.TrimSpace(err1.Error()))
			sterrs = append(sterrs,
				err1.(*Error).Stacktrace(),
			)
		}
	}
	for _, err := range errss {
		if err != nil {
			err = NewError(err)
			_, ok := err.(*Error)
			if ok {
				msgs = append(msgs, strings.TrimSpace(err.Error()))
				sterrs = append(sterrs,
					fmt.Sprintf(
						"------------------------------------------------------------------------------------------------------------------------------------\n\n%s\n\n%s",
						strings.TrimSpace(err.Error()),
						err.(*Error).Stacktrace()),
				)
			}
		}
	}
	if len(msgs) <= 0 {
		return nil
	}
	err := &Error{
		Code:    ErrSystemError.Code,
		Message: strings.Join(msgs, "\n"),
	}
	err.SetStacktrace(
		strings.Join(
			sterrs,
			"\n\n",
		),
	)
	return err
}

func LoggerFunc(fn func() error, path string, userID uint, email string, extras ...interface{}) {
	var err error
	start := time.Now()
	defer func() {
		end := time.Now()
		latency := end.Sub(start).Seconds()
		if rval := recover(); rval != nil {
			if rval := recover(); rval != nil {
				err = NewError(errors.New(fmt.Sprint(rval)))
			}
		}
		if path == "" {
			path = "default"
		}
		path = fmt.Sprintf("nft-marketet-api-fun-%s", path)
		var stacktrace, errText string
		errCode := 200
		if err != nil {
			errCode = 400
			err = NewError(err)
			errText = err.Error()
			retErr, ok := err.(*Error)
			if ok {
				stacktrace = retErr.Stacktrace()
			}
		}
		logger.Info(
			"logger_func_error",
			"msg info",
			zap.Any("referer", ""),
			zap.Any("ip", ""),
			zap.Any("method", "FUN"),
			zap.Any("path", path),
			zap.Any("raw_query", ""),
			zap.Any("latency", latency),
			zap.Any("status", errCode),
			zap.Any("user_agent", ""),
			zap.Any("platform", ""),
			zap.Any("os", ""),
			zap.Any("country", ""),
			zap.Any("email", email),
			zap.Any("user_id", userID),
			zap.Any("error_text", errText),
			zap.Any("stacktrace", stacktrace),
			zap.Any("body_request", helpers.ConvertJsonString(extras)),
			zap.Any("body_response", ""),
		)
		if os.Getenv("DEV") == "true" {
			if stacktrace != "" {
				fmt.Println(stacktrace)
			}
		}
	}()
	err = fn()
}
