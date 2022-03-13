package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"mysqld_export/config"
)

func ReadConfig() *config.Options {
	options, err := config.ParseConfig("./etc/export.yaml")
	if err != nil {
		logrus.Error(err)
	}

	logger := &lumberjack.Logger{
		Filename:   options.Logger.Filename,
		MaxSize:    options.Logger.MaxSize,
		MaxAge:     options.Logger.MaxAge,
		MaxBackups: options.Logger.MaxBackups,
		Compress:   options.Logger.Compress,
	}
	defer logger.Close()
	logrus.SetOutput(logger)

	switch options.Logger.LoggerLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}

	switch options.Logger.Formats {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	return options

}
