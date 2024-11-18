package handlers

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/logchunks"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

var allowedLogChunkUploadContentTypes = []string{
	"text/plain",
}

// LogStoreChunk allows clients to upload log chunks
// @Summary 	Log chunk storage
// @Tags        ClientAuthenticated
// @Accept     	plain
// @Description	Allows clients to upload log chunks for their log sessions. To test this endpoint, please run the upload locally to add the correct request body.
// @Security	JwtClient
// @Param 		log		body 		string 		true 	"Plain text log"
// @Success		201
// @Success		304
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     415		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/client/log/chunk [post]
func (h Handlers) LogStoreChunk(c *gin.Context) {
	log := h.logger.Named("LogStoreChunk")

	// Validate content type
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "" {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: "content-type header is required"})

		return
	}

	if !slices.Contains(allowedLogChunkUploadContentTypes, strings.ToLower(contentType)) {
		c.JSON(http.StatusUnsupportedMediaType, bodies.ErrorResponse{
			Error: fmt.Sprintf("content-type %s is not supported", contentType),
		})

		return
	}

	_, logID, ok := extractClientJwtClaimsFromContext(c, log)
	if !ok {
		return // Response has been sent by extractClientJwtClaimsFromContext
	}

	// Extract payload, plain text
	chunk, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("cannot extract payload from plain text chunk")

		c.JSON(http.StatusInternalServerError, bodies.ErrorResponse{Error: "cannot extract log chunk from request chunk"})

		return
	}

	// Arbitrary number
	if len(chunk) < 10 {
		log.Infof("chunk length %d too short for processing", len(chunk))

		c.Status(http.StatusNotModified)

		return
	}

	log.Debug("attempting to store log chunk")

	err = logchunk.StoreLogChunk(h.services, logID, chunk)

	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Info("chunk stored")

	c.Status(http.StatusCreated)
}
