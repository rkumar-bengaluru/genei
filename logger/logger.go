package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var defaultLogger *zap.Logger

type loggerCtx string

const loggerCtxKey loggerCtx = "logger"

func MillisecondDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d Millisecond", d.Milliseconds()))
}

func init() {
	var cfg zap.Config
	log.Println("init logger with IS_LOCAL=", isLocal())
	if isLocal() {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeDuration = MillisecondDurationEncoder
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defaultLogger = logger
}

// Gets the zap logger from context, creates a new one if does not exists
func Get(ctx context.Context) *zap.Logger {
	logger := logger(ctx)
	if logger == nil {
		return defaultLogger
	}
	return logger
}

func WithLogger(ctx context.Context, l *zap.Logger) context.Context {
	newCtx := context.WithValue(ctx, loggerCtxKey, l)
	return newCtx
}

func logger(ctx context.Context) *zap.Logger {
	logger := ctx.Value(loggerCtxKey)
	if logger == nil {
		return nil
	}
	return logger.(*zap.Logger)
}

func isLocal() bool {
	isLocal, err := strconv.ParseBool(os.Getenv("IS_LOCAL"))
	if err != nil {
		return false
	}
	return isLocal
}
