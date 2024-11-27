package user

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/domainerror"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// AddNewUser adds a new user and initializes the flow to set their password.
func AddNewUser(ctx context.Context, s *interactor.Services, email, firstName, lastName string, role constants.UserRole) (*uuid.UUID, error) {
	log := s.GetDomainLogger("user", "AddNewUser").
		With("email", email, "first_name", firstName, "last_name", lastName)

	log.Debugf("attempting to add new user")

	userUUID, err := s.Database.AddUser(ctx, database.AddUserParams{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserRole:  database.UserRoles(role),
	})

	log = log.With("user_uuid", userUUID)
	log.Info("user added to database")

	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	user, err := s.Database.GetUserByUuid(ctx, userUUID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID)

	err = s.Database.InitializeUserPasswordReset(ctx, user.ID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log.Debug("password flow initialized")

	reset, err := s.Database.GetPasswordResetByUserId(ctx, user.ID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("reset_code", reset.ResetToken.String())
	log.Debug("reset code retrieved")

	err = s.Emails.SendAccountSetPassword(user.Email, user.FirstName, reset.ResetToken.String())
	if err != nil {
		return nil, domainerror.New(domainerror.Unexpected, "error sending password set email to user", err)
	}

	log.Infof("user added and email sent")

	return &userUUID, nil
}

// AddNewUserWithPassword adds a new user with a password set.
// This should only be used with the CLI for development.
// With development mode disabled, this function returns an error.
func AddNewUserWithPassword(s *interactor.Services, email, firstName, lastName, password string) (*uuid.UUID, error) {
	if !s.Config.IsDevMode {
		return nil, errors.New("adding a user with a password set is only available in development mode")
	}

	log := s.GetDomainLogger("user", "AddNewUser").
		With("email", email, "first_name", firstName, "last_name", lastName)
	ctx := context.Background()

	log.Debugf("attempting to add new user")

	passHash, err := s.Passhash.PlainTextToHash(password)
	if err != nil {
		return nil, domainerror.New(domainerror.Unexpected, "cannot generate password hash", err)
	}

	userUUID, err := s.Database.AddUserWithPassHash(ctx, database.AddUserWithPassHashParams{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: &passHash,
	})

	if err != nil {
		log.Warnf("cannot add user: %s", err)

		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("user_uuid", userUUID)
	log.Infof("user added")

	return &userUUID, nil
}
