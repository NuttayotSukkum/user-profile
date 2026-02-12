package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(env string) {
	var cfg zap.Config

	if env == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.TimeKey = "TIME"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Log = logger
}

func Infof(format string, args ...any) {
	Log.Sugar().Infof(format, args...)
}

func Errorf(format string, args ...any) {
	Log.Sugar().Errorf(format, args...)
}

func Warnf(format string, args ...any) {
	Log.Sugar().Warnf(format, args...)
}
