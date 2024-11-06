package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lucianonooijen/capsa/server/internal/entities"

	"github.com/google/uuid"

	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// AddNewUser adds a new user and initializes the flow to set their password.
func AddNewUser(s *interactor.Services, email, firstName, lastName string) error {
	log := s.GetDomainLogger("user", "AddNewUser").
		With("email", email, "first_name", firstName, "last_name", lastName)
	ctx := context.TODO()

	log.Debugf("attempting to add new user")

	err := s.Database.AddUser(ctx, database.AddUserParams{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	})

	if err != nil {
		return fmt.Errorf("error creating new user: %w", err)
	}

	user, err := s.Database.GetUserByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("error getting created user by email: %w", err)
	}

	log = log.With("user_id", user.ID)

	err = s.Database.InitializeUserPasswordReset(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error initializing password setting flow for user: %w", err)
	}

	// TODO: Send out email for verification

	log.Infof("user added")

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
	ctx := context.TODO()

	log.Debugf("attempting to add new user")

	passHash, err := s.Passhash.PlainTextToHash(password)
	if err != nil {
		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot generate password hash", err)
	}

	userUUID, err := s.Database.AddUserWithPassHash(ctx, database.AddUserWithPassHashParams{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: sql.NullString{String: passHash, Valid: true},
	})

	if err != nil {
		log.Warnf("cannot add user: %s", err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_uuid", userUUID)
	log.Infof("user added")

	return &userUUID, nil
}
