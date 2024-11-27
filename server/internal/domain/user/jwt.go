package user

import (
	"context"
	"fmt"
	"slices"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/entities"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// ValidateUserJwt validates if the token passed in is valid for a user and validates if a user's role is in the allowed list.
func ValidateUserJwt(ctx context.Context, s *interactor.Services, tok string, allowedRoles []constants.UserRole) (*token.JwtClaims, error) {
	log := s.GetDomainLogger("client", "ValidateClientJwt")

	log.Debugf("starting user jwt validation")

	// Check validity for JWT
	claims, err := s.Token.ValidateJwt(tok)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorInvalidArgument, "token validation failed", err)
	}

	log = log.With("jwt_aud", claims.Audience, "jwt_sub", claims.Subject)
	log.Debug("token is valid, claims need verification")

	// Check that we have the correct audience, client audience cannot be used
	if claims.Audience != token.AudienceUser {
		message := fmt.Sprintf("token audience required is %s, but token contains %s", token.AudienceUser, claims.Audience)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, message, token.ErrorJwtInvalidAudience)
	}

	log.Debug("audience claim is correct")

	userUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot parse subject claim to uuid", err)
	}

	log = log.With("user_uuid", userUUID)
	log.Debug("log uuid parsed")

	// Get user for further validation
	user, err := s.Database.GetUserByUuid(ctx, userUUID)
	if err != nil {
		log.Warnf("cannot get log by uuid %s: %s", userUUID, err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID)
	log.Debug("user found in database")

	// Fail validation for deactivated users
	if user.DeactivatedOn != nil {
		log.Warnf("user for jwt is is deactivated on %s", *user.DeactivatedOn)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, "user is deactivated", fmt.Errorf("user deactivated on %s", *user.DeactivatedOn))
	}

	role, err := constants.UserRoleFromString(string(user.UserRole))
	if err != nil {
		log.Errorf("user role %s cannot be converted", user.UserUuid)

		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot extract user role to type", err)
	}

	// Check user role, see if it is in the allowed list
	if !slices.Contains(allowedRoles, role) {
		log.Infof("user role '%s' not in the allowed list '%#v'", role, allowedRoles)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, "user role is not allowed to access this", fmt.Errorf("user role '%s' not in the allowed list '%#v'", role, allowedRoles))
	}

	userPassUUID := user.PasswordUuid.String()

	// If a user has changed their password, don't succeed the validation
	if claims.JwtID != userPassUUID {
		log.Warnf("jwtid (%s) and password uuid (%s) do not match", claims.JwtID, userPassUUID)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, "token id does not match expected value", token.ErrorJwtIncorrectJwtID)
	}

	log.Debug("jwt validation succeeded")

	return claims, nil
}
