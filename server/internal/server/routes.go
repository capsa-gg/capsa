package server

import (
	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/internal/interactor"
	"github.com/lucianonooijen/capsa/server/internal/server/handlers"
)

func registerRoutes(r *gin.RouterGroup, h *handlers.Handlers, _ *interactor.Services) {
	// STATUS AND VERSION CHECK ROUTES
	r.GET("/status", h.Status)
}
