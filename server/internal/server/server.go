package server

import (
	"github.com/lucianonooijen/capsa/server/internal/entities"
	"go.uber.org/zap"
)

// Start is some example code
func Start(logger *zap.SugaredLogger, config *entities.Config) error {
	logger.Infof("server on port %d", config.ServerPort)

	return nil
}
