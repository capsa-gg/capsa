package server

import (
	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/handlers"
	"github.com/capsa-gg/capsa/server/internal/server/middleware"
)

// NOTE: the routes defined here should only have a single handler attached to them.
// If you need to add more than one handler, make a router group and add middleware.
// The reason for this is that the handlers from the handler package don't call c.Abort().

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
		clientLogs.POST("/chunk", h.LogStoreChunk)
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
		// Log routes
		userRoutes.GET("/logs", h.LogsList)
		userRoutes.GET("/logs/:loguuid/log", h.StreamLogChunks)
		userRoutes.GET("/logs/:loguuid/metadata", h.LogGetMetadata)

		// Environments
		userRoutes.GET("/environments", h.EnvironmentsList)
	}
}
