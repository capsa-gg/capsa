package server

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/handlers"
	"github.com/capsa-gg/capsa/server/internal/server/middleware"
	"github.com/capsa-gg/capsa/server/static"
)

// NOTE: the routes defined here should only have a single handler attached to them.
// If you need to add more than one handler, make a router group and add middleware.
// The reason for this is that the handlers from the handler package don't call c.Abort().

//nolint:gocritic // We use blocks to show nested routes more cleanly
func registerRoutes(r *gin.RouterGroup, h *handlers.Handlers, s *interactor.Services) {
	// STATIC FILES
	r.StaticFS(constants.APIStaticPath, mustFS(s))

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
	userRoutes.Use(middleware.AuthUserMiddleware(s, constants.AllUserRoles)) // Add authentication middleware
	{
		// Log routes
		userRoutes.GET("/logs", h.LogsList)
		userRoutes.GET("/logs/:loguuid/log", h.StreamLogChunks)
		userRoutes.GET("/logs/:loguuid/metadata", h.LogGetMetadata)

		// Environments
		userRoutes.GET("/environments", h.EnvironmentsList)
	}

	// Admin routes
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthUserMiddleware(s, []constants.UserRole{constants.UserRoleAdmin}))
	{
		adminRoutes.GET("/users", h.ListAllUsers)
		adminRoutes.POST("/users", h.CreateUser)
		adminRoutes.PUT("/users/:useruuid", h.UpdateUser)
		adminRoutes.POST("/users/:useruuid/activation", h.ReactivateUser)
		adminRoutes.DELETE("/users/:useruuid/activation", h.DeactivateUser)
	}
}

func mustFS(s *interactor.Services) http.FileSystem {
	sub, err := fs.Sub(static.FS, ".")

	if err != nil {
		s.Config.RootLogger.
			Named("HTTP").
			Named("registerRoutes").
			Named("mustFS").
			With(zap.Error(err)).
			Fatal("error generating http filesystem")
	}

	return http.FS(sub)
}
