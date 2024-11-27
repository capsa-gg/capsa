package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// UpdateUser updates a user in the database.
func UpdateUser(ctx context.Context, s *interactor.Services, userUUID uuid.UUID, userInfo *bodies.UserUpdateRequest) error {
	log := s.GetDomainLogger("user", "UpdateUser").With("user_uuid", userUUID).With("user_info", userInfo)

	log.Debugf("starting update user")

	if userInfo == nil {
		return domainerror.New(domainerror.Unexpected, "nil arrgument", errors.New("userInfo is nil"))
	}

	// Get user that needs to be updated
	user, err := s.Database.GetUserByUuid(ctx, userUUID)
	if err != nil {
		return domainerror.NewFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID)
	log.Debug("user retrieved")

	// Convert to user role
	role, err := constants.UserRoleFromString(userInfo.Role)
	if err != nil {
		return domainerror.New(domainerror.InvalidArgument, "invalid role", fmt.Errorf("invalid role: %s", userInfo.Role))
	}

	_, err = s.Database.UpdateUser(ctx, database.UpdateUserParams{
		ID:        user.ID,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		UserRole:  database.UserRoles(role),
	})
	if err != nil {
		return domainerror.NewFromDatabaseError(err)
	}

	log.Info("user updated")

	return nil
}
