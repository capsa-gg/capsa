package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/logs"
)

// LogsList allows users to fetch available logs from the database
// @Summary 	Log listing
// @Tags        UserLogs
// @Produce    	json
// @Description	Allows users to fetch available logs from the database
// @Security	JwtUser
// @Security	JwtAdmin
// @Success		200		{array}		bodies.LogOverview
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/user/logs [get]
func (h Handlers) LogsList(c *gin.Context) {
	log := h.logger.Named("LogsList")

	res, err := logs.GetAllLogsOverview(c, h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(res))

	c.JSON(http.StatusOK, res)
}

// LogGetMetadata allows users to fetch metadata for a log
// @Summary 	Log metadata
// @Tags        UserLogs
// @Produce    	json
// @Description	Allows users to fetch metadata for a log
// @Security	JwtUser
// @Security	JwtAdmin
// @Success		200		{object}	bodies.LogMetadata
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/user/logs/{logid}/metadata [get]
func (h Handlers) LogGetMetadata(c *gin.Context) {
	log := h.logger.Named("LogGetMetadata")

	logUUID, ok := getLogUUIDFromURI(c)
	if !ok {
		return // Response sent by getLogUUIDFromURI
	}

	res, err := logs.GetMetadataForLog(c, h.services, logUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debug("metadata fetched")

	c.JSON(http.StatusOK, res)
}
