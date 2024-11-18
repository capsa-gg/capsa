package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/logchunks"
	"github.com/capsa-gg/capsa/server/internal/entities"
)

// StreamLogChunks allows users to stream all uploaded chunks for a given log id
// @Summary 	Log chunk storage
// @Tags        UserLogs
// @Produce    	plain
// @Produce    	json
// @Description	Allows users to stream all uploaded chunks for a given log id
// @Security	JwtUser
// @Success		200		{string}	string 							"Log chunk stream"
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Header		500		{string} 	X-Capsa-Error					"Server error information"
// @Router 		/user/logs/{logid}/log [get]
func (h Handlers) StreamLogChunks(c *gin.Context) {
	log := h.logger.Named("LogStoreChunk")

	logUUID, ok := getLogUUIDFromURI(c)
	if !ok {
		return // Response sent by getLogUUIDFromURI
	}

	log = log.With("log_uuid", logUUID)

	// Validate if log exists, and get id from uuid
	// NOTE: this logic is usually done in the domain logic, but due to the streaming later in the handler, it's done here
	logInfo, err := h.services.Database.GetLogByUuid(context.TODO(), logUUID)
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

	err = logchunk.StreamLogChunks(h.services, logInfo.ID, streamer)
	if err != nil {
		log.Errorf("error streaming log chunks: %s", err)

		c.Header("X-Capsa-Error", "error streaming logs")
		c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(http.StatusOK)
}
