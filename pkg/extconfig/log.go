package extconfig

import (
	_ "beryju.io/gravity/pkg/extconfig/log_iml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (e *ExtConfig) Logger() *zap.Logger {
	return e.logger
}

func (e *ExtConfig) BuildLogger() *zap.Logger {
	l, err := zapcore.ParseLevel(e.LogLevel)
	if err != nil {
		l = zapcore.InfoLevel
	}
	if e.Debug {
		l = zapcore.DebugLevel
	}
	return e.BuildLoggerWithLevel(l)
}

func (e *ExtConfig) BuildLoggerWithLevel(l zapcore.Level) *zap.Logger {
	config := zap.Config{
		Encoding:         "json",
		Development:      false,
		OutputPaths:      []string{"stderr", "gravity-in-memory://"},
		ErrorOutputPaths: []string{"stderr", "gravity-in-memory://"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}
	config.Level = zap.NewAtomicLevelAt(l)
	config.DisableCaller = !e.Debug
	if e.Debug {
		config.Development = true
		config.Encoding = "console"
		config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	config.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	log, err := config.Build()
	if err != nil {
		panic(err)
	}
	return log.With(
		zap.String("app", "gravity.beryju.io"),
		zap.String("instance", e.Instance.Identifier),
		zap.String("version", FullVersion()),
	)
}
