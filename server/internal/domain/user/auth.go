package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/lucianonooijen/capsa/server/internal/infrastructure/token"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// ValidateUserJwt validates if the token passed in is valid for a user.
func ValidateUserJwt(s *interactor.Services, tok string) (*token.JwtClaims, error) {
	log := s.GetDomainLogger("client", "ValidateClientJwt")
	ctx := context.TODO()

	log.Debugf("starting user jwt validation")

	claims, err := s.Token.ValidateJwt(tok)
	if err != nil {
		return nil, fmt.Errorf("error validating token: %w", err)
	}

	log = log.With("jwt_aud", claims.Audience, "jwt_sub", claims.Subject)
	log.Debug("token is valid, claims need verification")

	if claims.Audience != token.AudienceUser {
		return nil, fmt.Errorf("token audience required is %s, but token contains %s: %w", token.AudienceUser, claims.Audience, token.ErrorJwtInvalidAudience)
	}

	log.Debug("token is valid and audience claim is correct")

	logUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("error parsing subject claim to uuid: %w", err)
	}

	log = log.With("user_uuid", logUUID)
	log.Debugf("log uuid parsed")

	user, err := s.Database.GetUserByUuid(ctx, logUUID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user from database based on uuid: %w", err)
	}

	log.With("user_id", user.ID).Info("user found in database")

	return claims, nil
}
