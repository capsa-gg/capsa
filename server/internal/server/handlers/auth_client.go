package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/constants"
	"github.com/lucianonooijen/capsa/server/internal/domain/logs"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"
)

// ClientAuth returns a 201 status code if a client has correctly authenticated
// @Summary 	Client authentication handler
// @Tags        ClientUnauthenticated
// @Accept 		json
// @Produce     json
// @Param		creation_request 	body 	bodies.ClientLogCreationRequest 	true 	"ClientLogCreationRequest"
// @Description	Allows clients to create a log session and receive a token to send logs with
// @Success		201		{object}	bodies.ClientLogCreationResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     401		{object}	bodies.ErrorResponse
// @Failure     403		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     409		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/client/auth [post]
func (h Handlers) ClientAuth(c *gin.Context) {
	log := h.logger.Named("ClientAuth")

	req, err := extractBodyJSON[bodies.ClientLogCreationRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	logType, err := constants.LogTypeFromString(req.Type)
	if err != nil {
		log.Infof("error getting logtype from string: %s", err)

		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{
			Error: fmt.Sprintf("%s is not a valid log type", req.Type),
		})

		return
	}

	sesInfo, err := logs.CreateNewLogSession(h.services, req.Key, req.Platform, logType)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	res := bodies.ClientLogCreationResponse{
		Token:   sesInfo.ClientJWT,
		LogID:   sesInfo.UUID,
		Expiry:  sesInfo.TokenExpiry,
		LinkWeb: "[UNIMPLEMENTED]", // TODO: implement
	}

	c.JSON(http.StatusCreated, res)
}
