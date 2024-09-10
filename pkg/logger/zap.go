package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() error {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	file, err := os.Create("logfile.log")
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}

	fileEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.DebugLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config.EncoderConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		fileCore,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}

func GetLogger() *zap.Logger {
	if Logger == nil {
		if err := InitLogger(); err != nil {
			//nolint:forbidigo // need to panic
			panic(fmt.Sprintf("failed to initialize logger: %v", err))
		}
	}

	return Logger
}

func SyncLogger() {
	if err := Logger.Sync(); err != nil {
		Logger.Error("Failed to sync logger", zap.Error(err))
	}
}
