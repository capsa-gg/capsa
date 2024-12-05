package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// LogsList allows users to fetch available logs from the database
// @Summary 	Log listing
// @Tags        UserLogs
// @Produce    	json
// @Param		title	query	string 		false 		"Title for which to fetch logs, required if 'env' is set"
// @Param		env		query	string 		false 		"Environment for which to fetch logs"
// @Description	Allows users to fetch available logs from the database, limiting the results to 1000.
// @Security	JwtUser
// @Security	JwtAdmin
// @Success		200		{object}	bodies.LogOverview
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/user/logs [get]
func (h Handlers) LogsList(c *gin.Context) {
	log := h.logger.Named("LogsList")

	filters, err := extractLogFilterSettings(c)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	fetchedLogs, hasMore, err := logs.GetLogs(c, h.services, filters)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	res := bodies.LogOverview{
		HasMore: hasMore,
		Logs:    fetchedLogs,
	}

	log.Debugf("fetched %d items, hasMore: %v", len(fetchedLogs), hasMore)

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

func extractLogFilterSettings(c *gin.Context) (logs.ListFilters, error) {
	filters := logs.ListFilters{}

	if env := c.Query("env"); env != "" {
		envUUID, err := uuid.Parse(env)

		if err != nil {
			return filters, domainerror.New(domainerror.InvalidArgument, "env query could not be parsed to uuid", err)
		}

		filters.Environment = &envUUID
	}

	return filters, nil
}
