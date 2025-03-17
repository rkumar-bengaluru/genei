package utils

import (
	"context"
	"os"
	"strconv"

	"example.com/rest-api/logger"
	"go.uber.org/zap"
)

func ReadStr(ctx context.Context, key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logger.Get(ctx).Fatal(key + " is missing")
	}
	return val
}

func ReadIntWithDefault(ctx context.Context, key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		logger.Get(ctx).With(zap.Error(err)).Error(key + " parsing failed")
		return defaultVal
	}
	return valInt
}

func ReadInt64WithDefault(ctx context.Context, key string, defaultVal int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	valInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		logger.Get(ctx).With(zap.Error(err)).Error(key + " parsing failed")
		return defaultVal
	}
	return valInt
}
