package zap

import (
	"go.uber.org/zap"
	"time"
	"testing"
)

func TestZap(t *testing.T)  {
	logger, _ := zap.NewProduction()
	url:="baidu.com"
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}