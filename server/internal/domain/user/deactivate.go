package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// DeactivateUser marks a user as deactivated and prevents them from accessing endpoints.
func DeactivateUser(ctx context.Context, s *interactor.Services, userUUID uuid.UUID) error {
	log := s.GetDomainLogger("user", "DeactivateUser").With("user_uuid", userUUID)

	log.Debugf("attempting to deactivate user")

	// Get user that needs to be deleted
	user, err := s.Database.GetUserByUuid(ctx, userUUID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID)
	log.Debug("user retrieved")

	// Delete outstanding password resets
	err = s.Database.DeletePasswordResetForUser(ctx, user.ID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Mark as deleted
	err = s.Database.DeactivateUser(ctx, user.ID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Info("user deactivated")

	return nil
}

// ReactivateUser marks the user as not deactivated and sends out an email to set a password.
func ReactivateUser(ctx context.Context, s *interactor.Services, userUUID uuid.UUID) error {
	log := s.GetDomainLogger("user", "ReactivateUser").With("user_uuid", userUUID)

	log.Debugf("attempting to reactivate user")

	// Get user that needs to be deleted
	user, err := s.Database.GetUserByUuid(ctx, userUUID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID)
	log.Debug("user retrieved")

	// Remove deactivation
	err = s.Database.RemoveUserDeactivation(ctx, user.ID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Init reset flow
	err = s.Database.InitializeUserPasswordReset(ctx, user.ID)
	if err != nil {
		log.Warnf("password reset could not be initialized: %s", err)

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Get reset data
	reset, err := s.Database.GetPasswordResetByUserId(ctx, user.ID)
	if err != nil {
		log.Warnf("password reset data could not be fetched: %s", err)

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Send email
	err = s.Emails.SendPasswordResetToken(user.Email, user.FirstName, reset.ResetToken.String())
	if err != nil {
		log.Warnf("password reset email could not be sent: %s", err)

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	return nil
}
