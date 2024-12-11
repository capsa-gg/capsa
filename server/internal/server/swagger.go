package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/swagger"
)

// @title Capsa Server
// @version 1
// @description Storage server for Project Capsa, a ground-breaking Unreal Engine logging solution

// @contact.name Luciano Nooijen
// @contact.url https://capsa.gg
// @contact.email capsa@lucianonooijen.com

// NOTE: there is securitydefinitions.bearerauth, but this does not allow setting the bearer value, this is better.
// @securitydefinitions.apikey JwtClient
// @in header
// @name Authorization
// @description Header value should be "Bearer ClientJwtString"

// @securitydefinitions.apikey JwtUser
// @in header
// @name Authorization
// @description Header value should be "Bearer UserJwtString", JwtAdmin also works for this endpoint

// @securitydefinitions.apikey JwtAdmin
// @in header
// @name Authorization
// @description Header value should be "Bearer UserJwtString", this JWT is similar to JwtUser, but contains the "Admin" role

// @host example.com
// @schemes http https
// @BasePath /v1

func swaggerDoc(conf *entities.Config) func(c *gin.Context) {
	if strings.Contains(conf.ServerHostname, "local") {
		serverHostname := conf.ServerHostname + ":" + strconv.Itoa(conf.ServerPort)

		swagger.SwaggerInfo.Host = serverHostname
	} else {
		swagger.SwaggerInfo.Host = conf.ServerHostname
	}

	return func(c *gin.Context) {
		c.String(http.StatusOK, swagger.SwaggerInfo.ReadDoc())
	}
}

func swaggerRedirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/v1/swagger/index.html")
}

func registerSwagger(r *gin.RouterGroup, conf *entities.Config) {
	// Disable Swagger in production mode
	if !conf.IsDevMode {
		return
	}

	r.GET("/swagger.json", swaggerDoc(conf))
	r.GET("/swagger", swaggerRedirect)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.URL("/v1/swagger.json"),
	))
}
