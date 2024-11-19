package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/logchunk"
	"github.com/capsa-gg/capsa/server/internal/entities"
)

const (
	logModeHeaderName       = "X-Capsa-Log-Mode"
	logModeSingleUnfiltered = "SingleUnfiltered"
	logModeSingleFiltered   = "SingleFiltered"
)

// StreamLogChunks allows users to stream all uploaded chunks for a given log id
// @Summary 	Log chunk storage
// @Tags        UserLogs
// @Produce    	plain
// @Produce    	json
// @Param		included_severities	query	string 		false 		"Included log line severities, optional"
// @Param		included_categories	query	string 		false 		"Included log categories, optional"
// @Param		excluded_categories	query	string 		false 		"Excluded log categories, will be ignored if included_categories is set, optional"
// @Description	Allows users to stream all uploaded chunks for a given log id. If no query parameters are set, the whole log will be fetched.
// @Security	JwtUser
// @Success		200		{string}	string 							"Log chunk stream"
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Header		all		{string} 	X-Capsa-Log-Mode				"Indicates the log mode, which can change the log content response, possible values: SingleUnfiltered|SingleFiltered"
// @Header		500		{string} 	X-Capsa-Error					"Server error information"
// @Router 		/user/logs/{logid}/log [get]
func (h Handlers) StreamLogChunks(c *gin.Context) {
	log := h.logger.Named("LogStoreChunk")

	logUUID, ok := getLogUUIDFromURI(c)
	if !ok {
		return // Response sent by getLogUUIDFromURI
	}

	log = log.With("log_uuid", logUUID)

	// Build the log line filters
	filters := logchunk.LogStreamLineFilters{}

	includedSeverities := c.Query("included_severities")
	if includedSeverities != "" {
		filters.IncludedSeverities = strings.Split(includedSeverities, ",")
	}

	includedCategories := c.Query("included_categories")
	if includedCategories != "" {
		filters.IncludedCategories = strings.Split(includedCategories, ",")
	}

	excludedCategories := c.Query("excluded_categories")
	if excludedCategories != "" {
		filters.ExcludedCategories = strings.Split(excludedCategories, ",")
	}

	log.Debugf("filters: %#v", filters)

	// Validate if log exists, and get id from uuid
	// NOTE: this logic is usually done in the domain logic, but due to the streaming later in the handler, it's done here
	logInfo, err := h.services.Database.GetLogByUuid(c, logUUID)
	if err != nil {
		h.sendErrorResponse(c, entities.NewDomainErrorFromDatabaseError(err))

		return
	}

	log = log.With("log_id", logInfo.ID)
	log.Debug("processing log stream request")

	// Set response headers
	c.Header("Content-Type", "text/plain")
	c.Header("Transfer-Encoding", "chunked")

	streamer := c.Writer.WriteString

	hasFilters := filters.HasFilters()

	log.With("has_filters", hasFilters)

	if hasFilters {
		c.Header(logModeHeaderName, logModeSingleUnfiltered)
		err = logchunk.StreamFilteredLogChunks(c, h.services, logInfo.ID, filters, streamer)
	} else {
		c.Header(logModeHeaderName, logModeSingleFiltered)
		err = logchunk.StreamUnfilteredLogChunks(c, h.services, logInfo.ID, streamer)
	}

	if err != nil {
		log.Errorf("error streaming log chunks: %s", err)

		c.Header("X-Capsa-Error", "error streaming logs")
		c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(http.StatusOK)

	log.Debug("finished streaming log")
}
