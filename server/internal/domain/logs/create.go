package logs

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/domainerror"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// LogCreatedResult contains the result of a created log.
type LogCreatedResult struct {
	UUID        uuid.UUID
	ClientJWT   string
	TokenExpiry time.Time
}

// CreateNewLogSession registers a new log session, returning the data about the created log.
func CreateNewLogSession(ctx context.Context, s *interactor.Services, envKey uuid.UUID, platform string, logType constants.LogType) (*LogCreatedResult, error) {
	log := s.GetDomainLogger("logs", "CreateNewLogSession").With("env_key", envKey, "log_type", logType)

	log.Debug("attempting to register a new log session")

	// Get environment
	env, err := s.Database.GetEnvironmentByKey(ctx, envKey)
	if err != nil {
		log.Warnf("cannot get environment: %s", err)

		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("env_name", env.Name, "env_title", env.Title)
	log.Debug("found matching environment")

	// Create log session in database
	sesKey, err := s.Database.AddNewLogSession(ctx, database.AddNewLogSessionParams{
		Environment: env.ID,
		LogType:     database.LogClientType(logType),
		Platform:    platform,
	})

	if err != nil {
		log.Warnf("cannot create new log session: %s", err)

		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("session_key", sesKey)
	log.Debug("log session added")

	// Generate JWT
	jwt, err := s.Token.GenerateClientJwt(sesKey.String())
	if err != nil {
		return nil, domainerror.New(domainerror.Unexpected, "cannot generate jwt for log session", err)
	}

	log.Debug("token generated")

	// Get JWT claims
	jwtClaims, err := s.Token.ValidateJwt(jwt)
	if err != nil {
		return nil, domainerror.New(domainerror.Unexpected, "cannot get jwt claims for log session token", err)
	}

	log.Debug("token parsed")

	// Assemble information
	logSession := LogCreatedResult{
		UUID:        sesKey,
		ClientJWT:   jwt,
		TokenExpiry: time.Unix(jwtClaims.Expiry, 0),
	}

	log.Info("log session creation succeeded")

	return &logSession, nil
}
