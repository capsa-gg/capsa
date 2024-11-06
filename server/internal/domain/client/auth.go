package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/lucianonooijen/capsa/server/internal/infrastructure/token"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// ValidateClientJwt validates if the token passed in is valid for a client.
func ValidateClientJwt(s *interactor.Services, tok string) (*token.JwtClaims, error) {
	log := s.GetDomainLogger("client", "ValidateClientJwt")
	ctx := context.TODO()

	log.Debug("starting client jwt validation")

	claims, err := s.Token.ValidateJwt(tok)
	if err != nil {
		return nil, fmt.Errorf("error validating token: %w", err)
	}

	log = log.With("jwt_aud", claims.Audience, "jwt_sub", claims.Subject)
	log.Debug("token is valid, claims need verification")

	if claims.Audience != token.AudienceClient {
		return nil, fmt.Errorf("token audience required is %s, but token contains %s: %w", token.AudienceClient, claims.Audience, token.ErrorJwtInvalidAudience)
	}

	log.Debug("token is valid and audience claim is correct")

	logUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("error parsing subject claim to uuid: %w", err)
	}

	log = log.With("log_uuid", logUUID)
	log.Debug("log uuid parsed")

	ses, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		return nil, fmt.Errorf("error fetching log from database based on uuid: %w", err)
	}

	log.With("environment", ses.Environment).Info("log found in database")

	return claims, nil
}
