package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() error {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	Logger = zapLogger
	return nil
}
