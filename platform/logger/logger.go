package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wreckitral/kafka-chat-app/pkg/config"
)

type Logger struct {
    *logrus.Logger
}

var logger = &Logger{}

func SetUpLogger() {
    logger = &Logger{logrus.New()}
    logger.Formatter = &logrus.JSONFormatter{}
    logger.SetOutput(os.Stdout)

    if config.AppCfg().Debug {
        logger.SetLevel(logrus.DebugLevel)
    }
}

func GetLogger() *Logger {
    return logger
}
