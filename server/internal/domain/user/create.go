package user

import (
	"context"
	"fmt"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/entities"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// AddNewUser adds a new user and initializes the flow to set their password.
func AddNewUser(ctx context.Context, s *interactor.Services, email, firstName, lastName string, role constants.UserRole) error {
	log := s.GetDomainLogger("user", "AddNewUser").
		With("email", email, "first_name", firstName, "last_name", lastName)

	log.Debugf("attempting to add new user")

	err := s.Database.AddUser(ctx, database.AddUserParams{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserRole:  database.UserRoles(role),
	})

	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	user, err := s.Database.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID)

	err = s.Database.InitializeUserPasswordReset(ctx, user.ID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Debug("password flow initialized")

	reset, err := s.Database.GetPasswordResetByUserId(ctx, user.ID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("reset_code", reset.ResetToken.String())
	log.Debug("reset code retrieved")

	err = s.Emails.SendAccountSetPassword(user.Email, user.FirstName, reset.ResetToken.String())
	if err != nil {
		return entities.NewDomainError(entities.DomainErrorUnexpected, "error sending password set email to user", err)
	}

	log.Infof("user added and email sent")

	return nil
}

// AddNewUserWithPassword adds a new user with a password set.
// This should only be used with the CLI for development.
// With development mode disabled, this function returns an error.
func AddNewUserWithPassword(s *interactor.Services, email, firstName, lastName, password string) (*uuid.UUID, error) {
	if !s.Config.IsDevMode {
		return nil, fmt.Errorf("adding a user with a password set is only available in development mode")
	}

	log := s.GetDomainLogger("user", "AddNewUser").
		With("email", email, "first_name", firstName, "last_name", lastName)
	ctx := context.Background()

	log.Debugf("attempting to add new user")

	passHash, err := s.Passhash.PlainTextToHash(password)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot generate password hash", err)
	}

	userUUID, err := s.Database.AddUserWithPassHash(ctx, database.AddUserWithPassHashParams{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: &passHash,
	})

	if err != nil {
		log.Warnf("cannot add user: %s", err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_uuid", userUUID)
	log.Infof("user added")

	return &userUUID, nil
}
