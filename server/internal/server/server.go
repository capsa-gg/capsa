package server

import (
	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/internal/entities"
)

// Start is some example code.
func Start(logger *zap.SugaredLogger, config *entities.Config) error {
	logger.Infof("server on port %d", config.ServerPort)

	return nil
}
