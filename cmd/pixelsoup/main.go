package main

import (
	"github.com/robertrenzorudio/pixelsoup/config"
	"github.com/robertrenzorudio/pixelsoup/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	conf := config.New()
	logSettings := createLogSetting(conf)
	logger.New(logSettings)

	logger.Log.Info("Server started", zap.String("Environment", conf.Env))
}

func createLogSetting(conf *config.Config) *logger.LogConfig {
	const logFileName = "log.json"
	logSettings := &logger.LogConfig{}
	switch conf.Env {
	case "development":
		logSettings.Core = logger.Console
	case "production":
		logSettings.Core = logger.File
		logSettings.FileName = logFileName
	}

	return logSettings
}
