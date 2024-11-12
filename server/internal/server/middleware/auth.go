package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/domain/user"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// AuthClientMiddleware is middleware to validate that a client is correctly authenticated.
func AuthClientMiddleware(s *interactor.Services) gin.HandlerFunc { //nolint:dupl // Not similar enough for a layer of abstraction
	log := s.Config.RootLogger.Named("AuthClientMiddleware").Sugar()

	return func(c *gin.Context) {
		// Get the token from the header
		tok := extractTokenFromHeader(c)
		if tok == "" {
			return // Helper function sent response and aborted request
		}

		// Validate the token
		claims, err := logs.ValidateClientJwt(s, tok)
		if err != nil {
			sendValidationError(c, err)

			log.Infof("received request with invalid token, error: %s, claims: %#v", err, claims)

			c.Abort()

			return
		}

		log = log.With("jwt_subject", claims.Subject)
		log.Info("user jwt claims added to request context")

		c.Set(constants.GinContextKeyValidatedClient, claims)

		c.Next()
	}
}

// AuthUserMiddleware is middleware to validate that a user is correctly authenticated.
func AuthUserMiddleware(s *interactor.Services) gin.HandlerFunc { //nolint:dupl // Not similar enough for a layer of abstraction
	log := s.Config.RootLogger.Named("AuthUserMiddleware").Sugar()

	return func(c *gin.Context) {
		// Get the token from the header
		tok := extractTokenFromHeader(c)
		if tok == "" {
			return // Helper function sent response and aborted request
		}

		// Validate the token
		claims, err := user.ValidateUserJwt(s, tok)
		if err != nil {
			sendValidationError(c, err)

			log.Infof("received request with invalid token, error: %s, claims: %#v", err, claims)

			c.Abort()

			return
		}

		log = log.With("jwt_subject", claims.Subject)
		log.Info("user jwt claims added to request context")

		c.Set(constants.GinContextKeyValidatedUser, claims)

		c.Next()
	}
}

// if extractTokenFromHeader returns an empty string, the c.Abort() call has been made and the middleware should exit.
func extractTokenFromHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")

	// Reject requests that do not contain authorization information
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, bodies.ErrorResponse{Error: "Authorization header required"})

		c.Abort()

		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: "Authorization header should be formatted like 'Bearer <jwt>', only Bearer tokens are accepted"})

		c.Abort()

		return ""
	}

	tok := parts[1]

	// 10 is a magic number, but JWTs will never be this small
	if len(tok) < 10 {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: "Authorization header Bearer token is invalid"})

		c.Abort()

		return ""
	}

	return tok
}

func sendValidationError(c *gin.Context, err error) {
	if errors.Is(err, token.ErrorJwtInvalidAudience) {
		c.JSON(http.StatusForbidden, bodies.ErrorResponse{Error: "Provided token does not contain required audience"})

		return
	}

	if errors.Is(err, token.ErrorJwtNotValidYet) {
		c.JSON(http.StatusForbidden, bodies.ErrorResponse{Error: "Provided token cannot be used yet"})

		return
	}

	if errors.Is(err, token.ErrorJwtExpired) {
		c.JSON(http.StatusUnauthorized, bodies.ErrorResponse{Error: "Provided token is expired"})

		return
	}

	c.JSON(http.StatusUnauthorized, bodies.ErrorResponse{Error: "Provided token is not valid"})
}
