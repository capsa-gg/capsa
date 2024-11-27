package user

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// ListAllUsers lists all users from the database.
func ListAllUsers(ctx context.Context, s *interactor.Services) ([]bodies.UserInfoResponse, error) {
	log := s.GetDomainLogger("user", "ListAllUsers")

	log.Debugf("attempting to fetch users")

	// Get users
	usersDB, err := s.Database.ListUsers(ctx, nil)
	if err != nil {
		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	users := make([]bodies.UserInfoResponse, 0, len(usersDB))

	for i := range usersDB {
		userDB := usersDB[i]

		user, err := dbUserToDomainUser(&userDB)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUserByUUID gets a user by UUID.
func GetUserByUUID(ctx context.Context, s *interactor.Services, userUUID uuid.UUID) (*bodies.UserInfoResponse, error) {
	log := s.GetDomainLogger("user", "GetUserByEmail").With("user_uuid", userUUID)

	log.Debugf("attempting to fetch user")

	// Get users
	usersDB, err := s.Database.ListUsers(ctx, &userUUID)
	if err != nil {
		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	if len(usersDB) != 1 {
		log.With("len_usersdb", len(usersDB)).Error("unexpected users length")

		return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "database error fetching user", errors.New("unexpected usersdb length"))
	}

	userDB := usersDB[0]

	user, err := dbUserToDomainUser(&userDB)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func dbUserToDomainUser(userDB *database.ListUsersRow) (bodies.UserInfoResponse, error) {
	userRole, err := constants.UserRoleFromString(string(userDB.Role))
	if err != nil {
		return bodies.UserInfoResponse{}, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot convert user role to string", err)
	}

	user := bodies.UserInfoResponse{
		UserUUID:       userDB.UserUuid,
		Email:          userDB.Email,
		FirstName:      userDB.FirstName,
		LastName:       userDB.LastName,
		HasPasswordSet: userDB.HasPasswordSet,
		Role:           userRole,
		DeactivatedTS:  userDB.DeactivatedTs,
		CreatedAt:      userDB.CreatedAt,
	}

	return user, nil
}
