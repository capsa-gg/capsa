package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/swagger"
)

// @title Capsa Server
// @version 1
// @description Storage server for Project Capsa, a ground-breaking Unreal Engine logging solution

// @contact.name Luciano Nooijen
// @contact.url https://lucianonooijen.com
// @contact.email capsa@lucianonooijen.com

// NOTE: there is securitydefinitions.bearerauth, but this does not allow setting the bearer value, this is better.
// @securitydefinitions.apikey JwtClient
// @in header
// @name Authorization
// @description Header value should be "Bearer ClientJwtString"

// @securitydefinitions.apikey JwtUser
// @in header
// @name Authorization
// @description Header value should be "Bearer UserJwtString"

// @host example.com
// @BasePath /v1

func swaggerDoc(conf *entities.Config) func(c *gin.Context) {
	if conf.IsDevMode {
		httpsReplacer := strings.NewReplacer(
			"https://", "",
			"http://", "")

		serverHostname := httpsReplacer.Replace(conf.ServerHostname) + ":" + strconv.Itoa(conf.ServerPort)

		swagger.SwaggerInfo.Host = serverHostname
	} else {
		swagger.SwaggerInfo.Host = strings.TrimSuffix(conf.ServerHostname, "/")
	}

	return func(c *gin.Context) {
		c.String(http.StatusOK, swagger.SwaggerInfo.ReadDoc())
	}
}

func swaggerRedirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/v1/swagger/index.html")
}

// TODO: do we want to disable Swagger for non-dev builds?
func registerSwagger(r *gin.RouterGroup, conf *entities.Config) {
	r.GET("/swagger.json", swaggerDoc(conf))
	r.GET("/swagger", swaggerRedirect)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.URL("/v1/swagger.json"),
	))
}
