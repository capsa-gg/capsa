package server

import (
	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/internal/interactor"
	"github.com/lucianonooijen/capsa/server/internal/server/handlers"
)

func registerRoutes(r *gin.RouterGroup, h *handlers.Handlers, _ *interactor.Services) {
	// Status route for health checks
	r.GET("/status", h.Status)

	// Unauthenticated client routes
	r.POST("/client/auth", h.ClientAuth)
}

// TO IMPLEMENT WITH ROUTER GROUPS AND MIDDLEWARE
// ==============================================
// Authenticated client routes /client/logs
// Log chunk POST /chunk
// Metadata POST /metadata
//
// Unauthenticated user routes /user/auth
// Login POST /login
// Request password reset GET /reset-password
// Set password POST /reset-password
//
// Authenticated user routes
// Everything else
