package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/constants"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"
)

// TODO: Finish this

// ClientAuth returns a 201 status code if a client has correctly authenticated
// @Summary 	Client authentication handler
// @Tags        Client
// @Description	Allows clients to create a log session and receive a token to send logs with
// @Produce     json
// @Success		201		{object}	bodies.StatusResponse
// @Failure     500		{object}	bodies.StatusResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/client/auth [post]
func (h Handlers) ClientAuth(c *gin.Context) {
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
