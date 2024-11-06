package server

import (
	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/internal/interactor"
	"github.com/lucianonooijen/capsa/server/internal/server/handlers"
	"github.com/lucianonooijen/capsa/server/internal/server/middleware"
)

//nolint:gocritic // We use blocks to show nested routes more cleanly
func registerRoutes(r *gin.RouterGroup, h *handlers.Handlers, s *interactor.Services) {
	// Status route for health checks
	r.GET("/status", h.Status)

	// Unauthenticated client routes
	r.POST("/client/auth", h.ClientAuth)

	// Authenticated client routes
	clientLogs := r.Group("/client/log")
	clientLogs.Use(middleware.AuthClientMiddleware(s)) // Add authentication middleware
	{
		clientLogs.POST("/metadata", h.LogMetadataSave)
		clientLogs.POST("/chunk", h.Status) // TODO: implement
	}

	// User authentication routes
	userAuth := r.Group("/user/auth")
	{
		userAuth.POST("/login", h.UserLogin)
		userAuth.GET("/password-reset", h.UserPasswordRequest)
		userAuth.POST("/password-reset", h.UserPasswordComplete)
	}

	// Authenticated user routes
	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.AuthUserMiddleware(s)) // Add authentication middleware
	{
		userRoutes.GET("/logs", h.Status)                 // TODO: implement
		userRoutes.GET("/logs/:logid/log", h.Status)      // TODO: implement
		userRoutes.GET("/logs/:logid/metadata", h.Status) // TODO: implement
	}
}
