package logging

import (
	"fmt"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func NewLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("error at creating logger: %v", err)
	}

	Logger = logger
	return nil
}
