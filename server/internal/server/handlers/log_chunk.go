package handlers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/logchunk"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// LogStoreChunk allows clients to upload log chunks
// @Summary 	Log chunk storage
// @Tags        ClientAuthenticated
// @Accept     	plain
// @Accept     	application/gzip
// @Description	Allows clients to upload log chunks for their log sessions. To test this endpoint, please run the upload locally to add the correct request body.
// @Security	JwtClient
// @Param 		log		body 		string 		true 	"Log contents, in line with the Content-Type specified"
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

	var chunkText []byte

	// Handle plain text stream
	switch contentType {
	case "text/plain":
		var err error
		chunkText, err = io.ReadAll(c.Request.Body)

		if err != nil {
			log.Error("cannot extract payload from plain text chunkText")

			c.JSON(http.StatusInternalServerError, bodies.ErrorResponse{Error: "cannot extract log chunkText from request chunkText"})

			return
		}
	case "application/gzip":
		decodedBody, err := decodeGzipBody(c)
		log.Errorf("gzip: %s", decodedBody)

		if err != nil {
			h.sendErrorResponse(c, err)

			return
		}

		chunkText = decodedBody
	default:
		c.JSON(http.StatusUnsupportedMediaType, bodies.ErrorResponse{
			Error: fmt.Sprintf("content-type %s is not supported", contentType),
		})

		return
	}

	// Arbitrary number
	if len(chunkText) < 10 {
		log.Infof("chunkText length %d too short for processing", len(chunkText))

		c.Status(http.StatusNotModified)

		return
	}

	_, logID, ok := extractClientJwtClaimsFromContext(c, log)
	if !ok {
		return // Response has been sent by extractClientJwtClaimsFromContext
	}

	log = log.With("log_uuid", logID)
	log.Debug("attempting to store log chunkText")

	err := logchunk.StoreLogChunk(h.services, logID, chunkText)

	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Info("chunkText stored")

	c.Status(http.StatusCreated)
}

func decodeGzipBody(c *gin.Context) ([]byte, error) {
	gzipContents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, domainerror.New(domainerror.Unexpected, "cannot extract request body", err)
	}

	buf := bytes.NewBuffer(gzipContents)

	reader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, domainerror.New(domainerror.InvalidArgument, "cannot decompress request body", err)
	}

	defer reader.Close() //nolint:errcheck // Best effort is fine here

	logContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, domainerror.New(domainerror.Unexpected, "cannot read decompressed log data", err)
	}

	return logContents, nil
}
