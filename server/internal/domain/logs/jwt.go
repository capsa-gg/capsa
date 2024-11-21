package logs

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// ValidateClientJwt validates if the token passed in is valid for a client.
func ValidateClientJwt(ctx context.Context, s *interactor.Services, tok string) (*token.JwtClaims, error) {
	log := s.GetDomainLogger("logs", "ValidateClientJwt")

	log.Debug("starting client jwt validation")

	// Validate claims
	claims, err := s.Token.ValidateJwt(tok)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorInvalidArgument, "token validation failed", err)
	}

	log = log.With("jwt_aud", claims.Audience, "jwt_sub", claims.Subject)
	log.Debug("token is valid, claims need verification")

	// Check that we have the correct audience, user tokens will not work
	if claims.Audience != token.AudienceClient {
		message := fmt.Sprintf("token audience required is %s, but token contains %s", token.AudienceClient, claims.Audience)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, message, token.ErrorJwtInvalidAudience)
	}

	log.Debug("token is valid and audience claim is correct")

	logUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot parse subject claim to uuid", err)
	}

	log = log.With("log_uuid", logUUID)
	log.Debug("log uuid parsed")

	// Get log from database for further validation
	ses, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		log.Warnf("cannot get log by uuid %s: %s", logUUID, err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log.With("environment", ses.Environment).Info("log found in database")

	return claims, nil
}
