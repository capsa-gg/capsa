package server

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

// GinServer is the Server instance struct.
type GinServer struct {
	port   int
	Router *gin.Engine
}

// Start the server instance.
func (s GinServer) Start() error {
	return s.Router.Run(":" + strconv.Itoa(s.port))
}

// New creates a new Server instance with middleware and handlers added.
// Use Server.Start() to run the server.
func New(services *interactor.Services) (*GinServer, error) {
	c := services.Config
	log := c.RootLogger.Named("HttpInit").Sugar()

	debug := c.IsDevMode
	port := c.ServerPort

	if debug {
		log.Debug("Detected debug mode for Gin")
		gin.SetMode(gin.DebugMode)
	} else {
		log.Debug("Detected production mode for Gin")
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = io.Discard // Not so clean way to don't get the "you are using Gin debug" error
	server := GinServer{
		Router: gin.New(),
		port:   port,
	}
	gin.DefaultWriter = os.Stdout // Reset to the default writer

	registerMiddleware(server.Router, c.RootLogger)

	initializedHandlers, err := handlers.New(services)
	if err != nil {
		return nil, fmt.Errorf("error initializing handlers: %w", err)
	}

	setGinRouteLogger(log.Named("HttpRoutes")) // Print the Gin routes using our own logger

	log.Debug("Registering routes")

	server.Router.GET(".well-known/jwks.json", initializedHandlers.Jwks)

	routerGroup := server.Router.Group("v1")
	registerRoutes(routerGroup, initializedHandlers, services)
	registerSwagger(routerGroup, c)

	log.Debugf("Routes registered, ready to start on port %d", port)

	return &server, nil
}
