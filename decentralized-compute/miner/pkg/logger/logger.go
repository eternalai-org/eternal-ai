package logger

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// AppLogger is a utility struct for logging data in an extremely high performance system.
// We can use both Logger and SugarLog for logging. For more information,
// just visit https://godoc.org/go.uber.org/zap
type AppLogger struct {
	// Sugar for logging
	*zap.SugaredLogger
	// configuration
	config map[string]interface{}
	// Logger for logging
	Logger *zap.Logger
}

func (atl *AppLogger) Print(args ...interface{}) {
	atl.Info(args...)
}

func (atl *AppLogger) Printf(f string, args ...interface{}) {
	atl.Infof(f, args...)
}

func (atl *AppLogger) Println(args ...interface{}) {
	atl.Info(args)
}

func (atl *AppLogger) Log(ctx context.Context, level logging.Level, msg string, fields ...any) {
	f := make([]zap.Field, 0, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		key := fields[i]
		value := fields[i+1]
		switch v := value.(type) {
		case string:
			f = append(f, zap.String(key.(string), v))
		case int:
			f = append(f, zap.Int(key.(string), v))
		case bool:
			f = append(f, zap.Bool(key.(string), v))
		default:
			f = append(f, zap.Any(key.(string), v))
		}
	}
	log := atl.WithOptions(zap.AddCallerSkip(1)).With(atl.WithContext(ctx), f)
	switch level {
	case logging.LevelDebug:
		log.Debug(msg)
	case logging.LevelInfo:
		log.Info(msg)
	case logging.LevelWarn:
		log.Warn(msg)
	case logging.LevelError:
		log.Error(msg)
	default:
		log.Warn(msg)
	}
}

// Return fields DataDog traceid
func (atl *AppLogger) WithContext(ctx context.Context) []zapcore.Field {
	fields := []zapcore.Field{}
	span, found := tracer.SpanFromContext(ctx)
	if found {
		fields = append(fields,
			zap.Uint64("trace.traceid", span.Context().TraceID()),
			zap.Uint64("trace.spanid", span.Context().SpanID()))
	}
	return fields
}

// LogRoundTrip prints the information about request and response.
func (atl *AppLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	var (
		nReq int64
		nRes int64
	)

	// Count number of bytes in request and response.
	if req != nil && req.Body != nil && req.Body != http.NoBody {
		nReq, _ = io.Copy(io.Discard, req.Body)
	}
	if res != nil && res.Body != nil && res.Body != http.NoBody {
		nRes, _ = io.Copy(io.Discard, res.Body)
	}

	fields := []zap.Field{
		zap.String("method", req.Method),
		zap.Int("status_code", res.StatusCode),
		zap.Duration("duration", dur),
		zap.Int64("req_bytes", nReq),
		zap.Int64("res_bytes", nRes),
	}

	// Set error level.
	switch {
	case err != nil:
		atl.Logger.With(fields...).Error(req.URL.String())
	case res != nil && res.StatusCode > 0 && res.StatusCode < 300:
		atl.Logger.With(fields...).Debug(req.URL.String())
	case res != nil && res.StatusCode > 299 && res.StatusCode < 500:
		atl.Logger.With(fields...).Warn(req.URL.String())
	case res != nil && res.StatusCode > 499:
		atl.Logger.With(fields...).Error(req.URL.String())
	default:
		atl.Logger.With(fields...).Error(req.URL.String())
	}

	return nil
}

// RequestBodyEnabled makes the client pass request body to logger
func (atl *AppLogger) RequestBodyEnabled() bool { return true }

// RequestBodyEnabled makes the client pass response body to logger
func (atl *AppLogger) ResponseBodyEnabled() bool { return true }

// AtLog is logger
var AtLog *AppLogger

func init() {
	InitLoggerDefaultDev()
}

// InitLoggerDefault -- format json
func InitLoggerDefault(enableDebug bool) {
	// init production encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.MessageKey = "message"
	// init production config
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}
	if enableDebug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	// build logger
	logger, _ := cfg.Build()

	sugarLog := logger.Sugar()
	cfgParams := make(map[string]interface{})
	AtLog = &AppLogger{sugarLog, cfgParams, logger}
}

// InitLoggerDefaultDev -- format text
func InitLoggerDefaultDev() {
	// init development encoder config
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	// init development config
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{"stdout"}
	// build logger
	logger, _ := cfg.Build()

	sugarLog := logger.Sugar()
	cfgParams := make(map[string]interface{})
	AtLog = &AppLogger{sugarLog, cfgParams, logger}
}

// GetLoggerInstanceFromContext returns the logger instance from context
func GetLoggerInstanceFromContext(ctx context.Context) *zap.Logger {
	return AtLog.Logger.With(AtLog.WithContext(ctx)...)
}
