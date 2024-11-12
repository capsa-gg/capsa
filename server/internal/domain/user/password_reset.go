package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// PasswordResetStart starts the password reset flow for a user.
func PasswordResetStart(s *interactor.Services, email string) error {
	log := s.GetDomainLogger("user", "PasswordResetStart").With("email", email)
	ctx := context.TODO()

	log.Debug("attempting to fetch user by email")

	// Get user
	user, err := s.Database.GetUserByEmail(ctx, email)
	if err != nil {
		log.Warn("user not found")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_id", user.ID).With("user_uuid", user.UserUuid)
	log.Debug("user found")

	// Init reset flow
	err = s.Database.InitializeUserPasswordReset(ctx, user.ID)
	if err != nil {
		log.Warn("password reset could not be initialized")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Get reset data
	reset, err := s.Database.GetPasswordResetByUserId(ctx, user.ID)
	if err != nil {
		log.Warn("password reset data could not be fetched")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Send email
	err = s.Emails.SendPasswordResetToken(user.Email, user.FirstName, reset.ResetToken.String())
	if err != nil {
		log.Warn("password reset email could not be sent")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	return nil
}

// PasswordResetComplete checks and consumes a password reset token and sets the new password.
func PasswordResetComplete(s *interactor.Services, resetToken uuid.UUID, newPassword string) error {
	log := s.GetDomainLogger("user", "PasswordResetComplete").With("reset_token", resetToken.String())
	ctx := context.TODO()

	log.Debug("attempting to fetch user by email")

	// Get reset information
	reset, err := s.Database.GetPasswordResetByResetToken(ctx, resetToken)
	if err != nil {
		log.Warn("password reset data could not be fetched")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("user_id", reset.UserID)

	// Get user
	user, err := s.Database.GetUserByID(ctx, reset.UserID)
	if err != nil {
		log.Warn("user not found")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Debug("user found")

	// Generate new password
	passHash, err := s.Passhash.PlainTextToHash(newPassword)
	if err != nil {
		log.Warn("user password hash could not be generated")

		return entities.NewDomainError(entities.DomainErrorUnexpected, "error generating password hash", err)
	}

	// Remove token as it has been used
	err = s.Database.DeletePasswordResetForUser(ctx, user.ID)
	if err != nil {
		log.Warn("user password reset could not be removed")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Store new password hash, set new password_uuid
	err = s.Database.UpdateUserPassword(ctx, database.UpdateUserPasswordParams{
		PasswordHash: sql.NullString{String: passHash, Valid: true},
		ID:           user.ID,
	})

	if err != nil {
		log.Warn("user password could not be updated")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Send email
	err = s.Emails.SendPasswordResetConfirmation(user.Email, user.FirstName)
	if err != nil {
		log.Warn("password reset email could not be sent")

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	return nil
}
