package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

type CoreType int

const (
	Console CoreType = iota
	File
	Both
)

type LogConfig struct {
	Core     CoreType
	FileName string
}

func New(logConf *LogConfig) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	logLevel := zapcore.InfoLevel

	cores := make([]zapcore.Core, 0)
	switch logConf.Core {
	case Console:
		cores = append(cores, createConsoleCore(config, logLevel))
	case File:
		cores = append(cores, createFileCore(config, logConf.FileName, logLevel))
	case Both:
		cores = append(cores, createConsoleCore(config, logLevel))
		cores = append(cores, createFileCore(config, logConf.FileName, logLevel))
	}

	core := zapcore.NewTee(cores...)
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func createConsoleCore(config zapcore.EncoderConfig, loglevel zapcore.Level) zapcore.Core {
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), loglevel)
	return core
}

func createFileCore(config zapcore.EncoderConfig, FileName string, loglevel zapcore.Level) zapcore.Core {
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	writer := zapcore.AddSync(logFile)
	core := zapcore.NewCore(fileEncoder, writer, loglevel)
	return core
}
