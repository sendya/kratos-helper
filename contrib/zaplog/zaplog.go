package zaplog

import (
	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"go.uber.org/zap"
)

func NewZapLogger() *kratoszap.Logger {
	zapconf := zap.NewProductionConfig()
	zapconf.EncoderConfig.MessageKey = ""
	zapconf.EncoderConfig.CallerKey = ""
	zapconf.EncoderConfig.TimeKey = ""
	zapconf.OutputPaths = []string{"stdout"}
	zapconf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zlog, _ := zapconf.Build()
	return kratoszap.NewLogger(zlog)
}