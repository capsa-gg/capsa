package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// Status returns a 200 status code if everything is fine
// @Summary 	Status handler
// @Tags        Common
// @Description	To be used with for status pings to check readiness and liveliness of the server
// @Produce     json
// @Success		200		{object}	bodies.StatusResponse
// @Failure     500		{object}	bodies.StatusResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/status [get]
func (h Handlers) Status(c *gin.Context) {
	err := h.services.DBConn.Ping()

	if err != nil {
		statusBody := bodies.StatusResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error pinnging database: %s", err),
			Version: constants.Version,
		}

		c.JSON(statusBody.Code, statusBody)

		return
	}

	statusBody := bodies.StatusResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Version: constants.Version,
	}

	c.JSON(statusBody.Code, statusBody)
}
