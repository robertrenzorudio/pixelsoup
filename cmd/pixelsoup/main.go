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
	logSettings := &logger.LogConfig{}
	if conf.Env == "production" || conf.Env == "prod" {
		logSettings.Core = logger.File
		logSettings.FileName = config.LogFileName
	} else {
		logSettings.Core = logger.Console
	}
	return logSettings
}
