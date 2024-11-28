package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

var (
	jsonBodyExtractionError = domainerror.New(domainerror.Unexpected, "error extracting json body", errors.New("json body extraction yielded nil"))
)

func (h Handlers) sendErrorResponse(c *gin.Context, err error) {
	// Domain Errors
	var domainError domainerror.Error
	if errors.As(err, &domainError) {
		h.sendDomainError(c, domainError)
		return
	}

	// Unexpected error - default to 500
	res := bodies.ErrorResponse{Error: "unexpected error"}

	if h.services.Config.IsDevMode {
		res.RawError = err.Error()
	}

	h.logger.Named("sendErrorResponse").Errorf("unhandled internal server error: %s", err)

	c.JSON(http.StatusInternalServerError, res)
}

func (h Handlers) sendDomainError(c *gin.Context, domainError domainerror.Error) {
	log := h.logger.Named("sendDomainError")

	res := bodies.ErrorResponse{
		Error: domainError.Error(),
	}

	if h.services.Config.IsDevMode {
		res.Details = domainError.Details
		res.RawError = domainError.RawError.Error()
	}

	switch domainError.Type {
	case domainerror.NoModifications:
		if h.services.Config.IsDevMode { // In dev mode, add details to header
			details, err := json.Marshal(res)
			if err == nil {
				c.Header("X-Capsa-Error", string(details))
			}
		}

		c.Status(http.StatusNotModified)

		return
	case domainerror.InvalidArgument:
		c.JSON(http.StatusBadRequest, res)
		return
	case domainerror.NoPermission:
		c.JSON(http.StatusForbidden, res)
		return
	case domainerror.NotFound:
		c.JSON(http.StatusNotFound, res)
		return
	case domainerror.Conflict:
		c.JSON(http.StatusConflict, res)
		return
	case domainerror.Unexpected:
		log.Errorf("unexpected domain error: %#v", domainError)

		c.JSON(http.StatusInternalServerError, res)

		return
	default:
		log.Errorf("unhandled domain error: %#v", domainError)

		c.JSON(http.StatusInternalServerError, res)

		return
	}
}
