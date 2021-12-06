package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is the default logging instance
var Log, err = zap.Config{
	Level:            zap.NewAtomicLevel(),
	Encoding:         "json",
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey:     "message",
		CallerKey:      "caller",
		LevelKey:       "level",
		StacktraceKey:  "stacktrace",
		TimeKey:        "ts",
		FunctionKey:    "function",
		LineEnding:     zapcore.DefaultLineEnding,
		NameKey:        "name",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.EpochMillisTimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	},
}.Build()
