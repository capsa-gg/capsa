package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// LogMetadataSave allows clients to store linked logs and additional metadata for the log session
// @Summary 	Log metadata storage
// @Tags        ClientAuthenticated
// @Accept     	json
// @Param		metadata 	body 	bodies.LogMetadataSaveRequest 	true 	"LogMetadataSaveRequest"
// @Description	Allows clients to store linked logs and additional metadata for their log session
// @Security	JwtClient
// @Success		201
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/client/log/metadata [post]
func (h Handlers) LogMetadataSave(c *gin.Context) {
	log := h.logger.Named("LogMetadataSave")

	_, logID, ok := extractClientJwtClaimsFromContext(c, log)
	if !ok {
		return // Response has been sent by extractClientJwtClaimsFromContext
	}

	req, err := extractBodyJSON[bodies.LogMetadataSaveRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	if req == nil {
		h.sendErrorResponse(c, errors.New("cannot extract body json from request"))

		return
	}

	log.Debug("body extracted")

	if len(req.AdditionalMetadata) == 0 && len(req.LogLinks) == 0 {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: "metadata and log links are both empty"})

		return
	}

	log.Debug("attempting to store log metadata")

	err = logs.SaveLogMetadata(c, h.services, logID, req.AdditionalMetadata, req.LogLinks)

	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Info("metadata stored")

	c.Status(http.StatusCreated)
}
