package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// ClientAuth returns a 201 status code if a client has correctly authenticated
// @Summary 	Client authentication handler
// @Tags        Client
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

	if req == nil {
		h.sendErrorResponse(c, errors.New("cannot extract body json from request"))

		return
	}

	logType, err := constants.LogTypeFromString(req.Type)
	if err != nil {
		log.Infof("error getting logtype from string: %s", err)

		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{
			Error: "not a valid log type: " + req.Type,
		})

		return
	}

	sesInfo, err := logs.CreateNewLogSession(c, h.services, req.Key, req.Platform, logType)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	// Note: we force https to be used
	linkWeb := fmt.Sprintf("https://%s/logs/%s", h.services.Config.WebappHostname, sesInfo.UUID)

	res := bodies.ClientLogCreationResponse{
		Token:   sesInfo.ClientJWT,
		LogID:   sesInfo.UUID,
		Expiry:  sesInfo.TokenExpiry,
		LinkWeb: linkWeb,
	}

	c.JSON(http.StatusCreated, res)
}
