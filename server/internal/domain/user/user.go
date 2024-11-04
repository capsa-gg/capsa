package user

import (
	"context"
	"fmt"

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
