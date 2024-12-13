package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domain/logchunk"
	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/entities"
)

const (
	logModeHeaderName       = "X-Capsa-Log-Mode"
	logModeSingleUnfiltered = "SingleUnfiltered"
	logModeSingleFiltered   = "SingleFiltered"
	logModeMergedFiltered   = "MergedFiltered"

	maxLinesPerMergedChunk = uint64(100)
)

// StreamLogChunks allows users to stream all uploaded chunks for a given log id
// @Summary 	Log chunk storage
// @Tags        UserLogs
// @Produce    	plain
// @Produce    	json
// @Param		included_severities	query	string 		false 		"Included log line severities, optional"
// @Param		included_categories	query	string 		false 		"Included log categories, optional"
// @Param		excluded_categories	query	string 		false 		"Excluded log categories, will be ignored if included_categories is set, optional"
// @Param		merge_logs			query	[]string 	false 		"Comma separated log UUIDs that should be merged"
// @Description	Allows users to stream all uploaded chunks for a given log id. If no query parameters are set, the whole log will be fetched.
// @Security	JwtUser
// @Security	JwtAdmin
// @Success		200		{string}	string 							"Log chunk stream"
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Header		all		{string} 	X-Capsa-Log-Mode				"Indicates the log mode, which can change the log content response, possible values: SingleUnfiltered|SingleFiltered"
// @Header		500		{string} 	X-Capsa-Error					"Server error information"
// @Router 		/user/logs/{logid}/log [get]
func (h Handlers) StreamLogChunks(c *gin.Context) { //nolint:funlen // This is fine
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
		h.sendErrorResponse(c, domainerror.NewFromDatabaseError(err))

		return
	}

	log = log.With("log_id", logInfo.ID)
	log.Debug("processing log stream request")

	streamer := c.Writer.WriteString

	// Check if we need to merge logs or not
	mergeLogsArg := c.Query("merge_logs")
	if mergeLogsArg != "" {
		c.Header(logModeHeaderName, logModeMergedFiltered)

		mergeLogs := strings.Split(mergeLogsArg, ",")

		h.streamMergedLogs(c, &logInfo, mergeLogs, filters, streamer)

		return
	}

	hasFilters := filters.HasFilters()

	log.With("has_filters", hasFilters)

	// Set response headers
	c.Header("Content-Type", "text/plain")
	c.Header("Transfer-Encoding", "chunked")

	if hasFilters {
		c.Header(logModeHeaderName, logModeSingleFiltered)
		err = logchunk.StreamFilteredLogChunks(c, h.services, logInfo.ID, filters, streamer)
	} else {
		c.Header(logModeHeaderName, logModeSingleUnfiltered)
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

func (h Handlers) streamMergedLogs(c *gin.Context, baseLog *database.Log, mergeLogs []string, filters logchunk.LogStreamLineFilters, streamChunk entities.ChunkStreamer) {
	log := h.logger.Named("streamMergedLogs")

	mergeLogsInfo := []database.Log{}

	for _, ml := range mergeLogs {
		mlUUID, err := uuid.Parse(ml)
		if err != nil {
			h.sendErrorResponse(c, domainerror.New(domainerror.InvalidArgument, ml+" cannot be parsed to uuid", err))

			return
		}

		l, err := h.services.Database.GetLogByUuid(c, mlUUID)
		if err != nil {
			h.sendErrorResponse(c, domainerror.NewFromDatabaseError(err))

			return
		}

		mergeLogsInfo = append(mergeLogsInfo, l)
	}

	mergeInput, err := h.getMergedLogLoaders(c, baseLog, mergeLogsInfo, filters)
	if err != nil {
		h.sendErrorResponse(c, domainerror.NewFromDatabaseError(err))

		return
	}

	// Set response headers
	c.Header("Content-Type", "text/plain")
	c.Header("Transfer-Encoding", "chunked")

	// Start streaming
	err = logs.StreamMergedLog(c, h.services, mergeInput, maxLinesPerMergedChunk, streamChunk)
	if err != nil {
		log.Errorf("error streaming log chunks: %s", err)

		c.Header("X-Capsa-Error", "error streaming merged logs")
		c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(http.StatusOK)
}

func (h Handlers) getMergedLogLoaders(c *gin.Context, baseLog *database.Log, mergeLogs []database.Log, filters logchunk.LogStreamLineFilters) (entities.MergedLogInput, error) {
	if baseLog == nil {
		return nil, domainerror.New(domainerror.Unexpected, "baseLog nil", errors.New("baseLog nil"))
	}

	typeCounts := map[database.LogClientType]int{
		database.LogClientTypeEditor: 0,
		database.LogClientTypeGame:   0,
		database.LogClientTypeClient: 0,
		database.LogClientTypeServer: 0,
	}

	// Add base log as first argument
	baseLoader, err := logchunk.GenerateFilteredLineLoaderForLog(c, h.services, baseLog.ID, filters)
	if err != nil {
		return nil, err
	}

	mergedLogInput := entities.MergedLogInput{
		{Key: "--", Loader: baseLoader},
	}

	// Add the rest of the merged logs
	for _, ml := range mergeLogs {
		typeCounts[ml.LogType]++
		num := typeCounts[ml.LogType]
		key := fmt.Sprintf("%s%d", logClientPrefix(ml.LogType), num)

		loader, err := logchunk.GenerateFilteredLineLoaderForLog(c, h.services, ml.ID, filters)
		if err != nil {
			return nil, fmt.Errorf("error generating loader for log with ID %d: %w", ml.ID, err)
		}

		mergedLogInput = append(mergedLogInput, entities.MergedLogInputData{
			Key:    key,
			Loader: loader,
		})
	}

	return mergedLogInput, err
}

func logClientPrefix(lct database.LogClientType) string {
	switch lct {
	case database.LogClientTypeEditor:
		return "E"
	case database.LogClientTypeGame:
		return "G"
	case database.LogClientTypeClient:
		return "C"
	case database.LogClientTypeServer:
		return "S"
	default:
		return "?"
	}
}
