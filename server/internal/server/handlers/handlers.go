package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// Handlers contains Gin request handlers as methods.
type Handlers struct {
	services *interactor.Services
	logger   *zap.SugaredLogger
}

// New returns Handlers instance.
func New(services *interactor.Services) (*Handlers, error) {
	handlers := Handlers{
		services: services,
		logger:   services.Config.RootLogger.Named("HttpHandlers").Sugar(),
	}

	validate := validator.New()
	err := validate.Struct(handlers)

	return &handlers, err
}

// extractUserJwtClaimsFromContext extracts the client claims from the context
// The boolean is "ok", if this is false, just return the handler, the response has already been sent.
func extractClientJwtClaimsFromContext(c *gin.Context, log *zap.SugaredLogger) (*token.JwtClaims, uuid.UUID, bool) {
	return extractJwtClaimsFromContext(c, log, constants.GinContextKeyValidatedClient)
}

// extractUserJwtClaimsFromContext extracts the user claims from the context
// The boolean is "ok", if this is false, just return the handler, the response has already been sent.
func extractUserJwtClaimsFromContext(c *gin.Context, log *zap.SugaredLogger) (*token.JwtClaims, uuid.UUID, bool) { // nolint:unused // will be used very soon
	return extractJwtClaimsFromContext(c, log, constants.GinContextKeyValidatedUser)
}

// extractJwtClaimsFromContext extracts the data from the Gin contexts, should only be used by the other two extraction helpers.
func extractJwtClaimsFromContext(c *gin.Context, log *zap.SugaredLogger, contextKey string) (*token.JwtClaims, uuid.UUID, bool) {
	jwtClaimsRaw := c.MustGet(contextKey)

	jwtClaims, ok := jwtClaimsRaw.(*token.JwtClaims)
	if !ok {
		log.Named("extractJwtClaimsFromContext").Error("could not convert jwt claims to struct")

		c.JSON(http.StatusInternalServerError, bodies.ErrorResponse{Error: "error extracting jwt claims from request context"})

		return nil, uuid.Nil, false
	}

	logIDStr := jwtClaims.Subject
	log = log.With("log_id", logIDStr)

	logID, err := uuid.Parse(logIDStr)
	if err != nil { // TODO: Helper func
		log.Named("extractJwtClaimsFromContext").Error("could not parse jwt subject to uuid")

		c.JSON(http.StatusInternalServerError, bodies.ErrorResponse{Error: "error converting jwt subject to uuid"})

		return nil, uuid.Nil, false
	}

	return jwtClaims, logID, true
}

// getLogUUIDFromURI extracts the "loguuid" from the parameters.
// If the boolean return argument is false, stop handler execution. The response has already been sent.
func getLogUUIDFromURI(c *gin.Context) (uuid.UUID, bool) {
	return getUUIDFromURI(c, "loguuid")
}

// getUUIDFromURI extracts the uuid for a given parameter from the context parameters, parses it and returns the UUID.
// If the boolean return argument is false, stop handler execution. The response has already been sent.
// For re-used parameters, create a wrapper to not duplicate the Param name.
func getUUIDFromURI(c *gin.Context, param string) (uuid.UUID, bool) {
	val := c.Param(param)

	if val == "" || len(val) > 40 {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: param + " is a required uuid value"})

		return uuid.Nil, false
	}

	valUUID, err := uuid.Parse(val)
	if err != nil {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: param + " cannot be parsed to uuid"})

		return uuid.Nil, false
	}

	return valUUID, true
}
