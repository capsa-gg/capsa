package user

import (
	"context"
	"fmt"

	"github.com/lucianonooijen/capsa/server/internal/entities"

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
		return nil, entities.NewDomainError(entities.DomainErrorInvalidArgument, "token validation failed", err)
	}

	log = log.With("jwt_aud", claims.Audience, "jwt_sub", claims.Subject)
	log.Debug("token is valid, claims need verification")

	if claims.Audience != token.AudienceUser {
		message := fmt.Sprintf("token audience required is %s, but token contains %s", token.AudienceUser, claims.Audience)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, message, token.ErrorJwtInvalidAudience)
	}

	log.Debug("token is valid and audience claim is correct")

	logUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot parse subject claim to uuid", err)
	}

	log = log.With("user_uuid", logUUID)
	log.Debugf("log uuid parsed")

	user, err := s.Database.GetUserByUuid(ctx, logUUID)
	if err != nil {
		log.Warnf("cannot get log by uuid %s: %s", logUUID, err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log.With("user_id", user.ID).Info("user found in database")

	return claims, nil
}
