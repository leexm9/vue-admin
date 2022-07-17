package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var encCfg = zapcore.EncoderConfig{
	TimeKey:        `timestamp`,
	LevelKey:       "serverity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var logger *zap.Logger

func GetInstance() *zap.Logger {
	return logger
}

func init() {
	stdout := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zapcore.DebugLevel))
	cores := []zapcore.Core{stdout}
	core := zapcore.NewTee(cores...)
	hostname, _ := os.Hostname()
	fields := zap.Fields(zap.String("host", hostname))
	logger = zap.New(core, zap.AddCaller(), fields, zap.AddStacktrace(zapcore.InfoLevel))
}
