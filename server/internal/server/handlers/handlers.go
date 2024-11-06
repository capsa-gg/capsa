package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/constants"
	"github.com/lucianonooijen/capsa/server/internal/infrastructure/token"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"

	"github.com/lucianonooijen/capsa/server/internal/interactor"
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
		logger:   services.Config.RootLogger.Named("http_handlers").Sugar(),
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
