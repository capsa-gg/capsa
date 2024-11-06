package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"
)

func (h Handlers) sendErrorResponse(c *gin.Context, err error) {
	// Domain Errors
	var domainError entities.DomainError
	if errors.As(err, &domainError) {
		h.sendDomainError(c, domainError)
		return
	}

	// Unexpected error - default to 500
	res := bodies.ErrorResponse{Error: "unexpected error"}

	if h.services.Config.IsDevMode {
		res.RawError = err.Error()
	}

	c.JSON(http.StatusInternalServerError, res)
}

func (h Handlers) sendDomainError(c *gin.Context, domainError entities.DomainError) {
	res := bodies.ErrorResponse{
		Error: domainError.Error(),
	}

	if h.services.Config.IsDevMode {
		res.Details = domainError.Details
		res.RawError = domainError.RawError.Error()
	}

	switch domainError.Type {
	case entities.DomainErrorInvalidArgument:
		c.JSON(http.StatusBadRequest, res)
		return
	case entities.DomainErrorNoPermission:
		c.JSON(http.StatusForbidden, res)
		return
	case entities.DomainErrorNotFound:
		c.JSON(http.StatusNotFound, res)
		return
	case entities.DomainErrorConflict:
		c.JSON(http.StatusConflict, res)
		return
	case entities.DomainErrorUnexpected:
		c.JSON(http.StatusInternalServerError, res)
		return
	default:
		return
	}
}
