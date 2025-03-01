package infra

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	Logger = logger
}
