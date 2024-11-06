package user

import (
	"context"
	"errors"
	"time"

	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// Login validates a user's password and returns the login result if everything is valid.
func Login(s *interactor.Services, email, password string) (*entities.UserLoginResult, error) {
	log := s.GetDomainLogger("user", "CreateNewLogSession").With("email", email)
	ctx := context.TODO()

	log.Debug("attempting to log in user")

	// Get user
	user, err := s.Database.GetUserByEmail(ctx, email)
	if err != nil {
		log.Warn("user not found")

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_uuid", user.UserUuid)
	log.Debug("user found")

	if !user.PasswordHash.Valid || !user.PasswordUuid.Valid {
		log.Warnf("user password hash (%#v) or password uuid (%#v) not found", user.PasswordHash, user.PasswordUuid)

		return nil, entities.NewDomainError(entities.DomainErrorNotFound, "user password not set", errors.New("password hash or uuid not valid"))
	}

	name := user.FirstName + " " + user.LastName

	log = log.With("name", name)
	log.Debug("user password is set")

	// Validate password
	err = s.Passhash.ComparePassToHash(password, user.PasswordHash.String)
	if err != nil {
		log.Warnf("error validating user password hash: %s", err)

		return nil, entities.NewDomainError(entities.DomainErrorNoPermission, "password validation failed", err)
	}

	log.Debug("user password validation succeeded")

	// Generate JWT
	jwt, err := s.Token.GenerateUserJwt(user.UserUuid.String(), user.PasswordUuid.UUID.String(), name)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot generate jwt for log session", err)
	}

	log.Debug("token generated")

	// Get JWT claims
	jwtClaims, err := s.Token.ValidateJwt(jwt)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot get jwt claims for user token", err)
	}

	log.Debug("token parsed")

	// Assemble information
	logSession := entities.UserLoginResult{
		Token:       jwt,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		UserUUID:    user.UserUuid.String(),
		TokenExpiry: time.Unix(jwtClaims.Expiry, 0),
	}

	log.Info("user successfully authenticated")

	return &logSession, nil
}
