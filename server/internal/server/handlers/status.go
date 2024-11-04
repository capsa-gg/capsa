package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"
	"net/http"
)

// Status returns a 200 status code if everything is fine
// @Summary 	Status handler
// @Tags        Common
// @Description	To be used with for status pings to check readiness and liveliness of the server
// @Produce     json
// @Success		200		{object}	bodies.StatusResponse
// @Failure     500		{object}	bodies.StatusResponse
// @Router 		/status [get]
func (h Handlers) Status(c *gin.Context) {
	err := h.services.DBConn.Ping()

	if err != nil {
		statusBody := bodies.StatusResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error pinnging database: %s", err),
		}

		c.JSON(statusBody.Code, statusBody)

		return
	}

	statusBody := bodies.StatusResponse{
		Code:    http.StatusOK,
		Message: "ok",
	}

	c.JSON(statusBody.Code, statusBody)
}
