package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
	"go.uber.org/zap"
)

// Handlers contains Gin request handlers as methods.
type Handlers struct {
	services *interactor.Services
	logger   *zap.SugaredLogger
}

// New returns Handlers instance.
func New(services *interactor.Services) (*Handlers, error) {
	handlers := Handlers{
		services: services,
		logger:   services.Config.RootLogger.Named("http_handlers").Sugar(),
	}

	validate := validator.New()
	err := validate.Struct(handlers)

	return &handlers, err
}
