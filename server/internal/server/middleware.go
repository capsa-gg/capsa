package server

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/server/middleware"
)

func registerMiddleware(router *gin.Engine, logger *zap.Logger) {
	router.Use(gin.Recovery())
	router.Use(middleware.GinLogger(logger))
	router.Use(middleware.ServerVersionMiddleware())
	router.Use(middleware.SecurityHeadersMiddleware())

	// TODO: Configure properly
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.ExposeHeaders = []string{
		"X-Capsa-Log-Mode",
		"X-Capsa-Server-Version",
		"X-Capsa-Error",
	}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"accept",
		"origin",
		"Cache-Control",
		"Authorization",
	}

	router.Use(cors.New(config))

	if err := router.SetTrustedProxies([]string{}); err != nil { // TODO: Set this when needed
		logger.Named("registerMiddleware.SetTrustedProxies").Sugar().Fatalf("error from SetTrustedProxies: %s", err)
	}
}

func setGinRouteLogger(logger *zap.SugaredLogger) {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		handlerReplacer := strings.NewReplacer(
			"github.com/capsa-gg/capsa/server/internal/server/handlers.", "",
			"github.com/gin-gonic/", "")
		handlerShort := handlerReplacer.Replace(handlerName)

		logger.Debug(fmt.Sprintf("Route registered: %-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerShort, nuHandlers))
	}
}
