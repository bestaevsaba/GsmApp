package logger

import (
	"log"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func init() {
	var err error
	zapLogger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Не удалось инициализировать zap logger: %v", err)
	}
}

func GetLogger() *zap.Logger {
	return zapLogger
}
