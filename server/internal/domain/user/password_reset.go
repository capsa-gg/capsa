package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// PasswordResetStart starts the password reset flow for a user.
func PasswordResetStart(ctx context.Context, s *interactor.Services, email string) error {
	log := s.GetDomainLogger("user", "PasswordResetStart").With("email", email)

	log.Debug("attempting to fetch user by email")

	// Get user
	user, err := s.Database.GetUserByEmail(ctx, email)
	if err != nil {
		log.Warn("user not found")

		return domainerror.NewFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID).With("user_uuid", user.UserUuid)
	log.Debug("user found")

	// Don't allow password reset for deactivated users
	if user.DeactivatedOn != nil {
		log.Warnf("user tried to reset password, but is deactivated on %s", *user.DeactivatedOn)

		return domainerror.New(domainerror.NotFound, "no active user found", fmt.Errorf("user deactivated on %s", *user.DeactivatedOn))
	}

	// Init reset flow
	err = s.Database.InitializeUserPasswordReset(ctx, user.ID)
	if err != nil {
		log.Warnf("password reset could not be initialized: %s", err)

		return domainerror.NewFromDatabaseError(err)
	}

	// Get reset data
	reset, err := s.Database.GetPasswordResetByUserId(ctx, user.ID)
	if err != nil {
		log.Warnf("password reset data could not be fetched: %s", err)

		return domainerror.NewFromDatabaseError(err)
	}

	// Send email
	err = s.Emails.SendPasswordResetToken(user.Email, user.FirstName, reset.ResetToken.String())
	if err != nil {
		log.Warn("password reset email could not be sent")

		return domainerror.NewFromDatabaseError(err)
	}

	return nil
}

// PasswordResetComplete checks and consumes a password reset token and sets the new password.
func PasswordResetComplete(ctx context.Context, s *interactor.Services, resetToken uuid.UUID, newPassword string) error {
	log := s.GetDomainLogger("user", "PasswordResetComplete").With("reset_token", resetToken.String())

	log.Debug("attempting to fetch user by email")

	// Get reset information
	reset, err := s.Database.GetPasswordResetByResetToken(ctx, resetToken)
	if err != nil {
		log.Warn("password reset data could not be fetched")

		return domainerror.NewFromDatabaseError(err)
	}

	log = log.With("user_id", reset.UserID)

	// Get user
	user, err := s.Database.GetUserByID(ctx, reset.UserID)
	if err != nil {
		log.Warn("user not found")

		return domainerror.NewFromDatabaseError(err)
	}

	log.Debug("user found")

	// Generate new password
	passHash, err := s.Passhash.PlainTextToHash(newPassword)
	if err != nil {
		log.Warn("user password hash could not be generated")

		return domainerror.New(domainerror.Unexpected, "error generating password hash", err)
	}

	// Remove token as it has been used
	err = s.Database.DeletePasswordResetForUser(ctx, user.ID)
	if err != nil {
		log.Warn("user password reset could not be removed")

		return domainerror.NewFromDatabaseError(err)
	}

	// Store new password hash, set new password_uuid
	err = s.Database.UpdateUserPassword(ctx, database.UpdateUserPasswordParams{
		PasswordHash: &passHash,
		ID:           user.ID,
	})

	if err != nil {
		log.Warn("user password could not be updated")

		return domainerror.NewFromDatabaseError(err)
	}

	// Send email
	err = s.Emails.SendPasswordResetConfirmation(user.Email, user.FirstName)
	if err != nil {
		log.Warn("password reset email could not be sent")

		return domainerror.NewFromDatabaseError(err)
	}

	return nil
}
