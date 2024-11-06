package logs

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/lucianonooijen/capsa/server/constants"
	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// CreateNewLogSession registers a new log session, returning the uuid that will be used to identify the log.
func CreateNewLogSession(s *interactor.Services, envKey uuid.UUID, logType constants.LogType) (*uuid.UUID, error) {
	log := s.GetDomainLogger("logs", "CreateNewLogSession").With("env_key", envKey, "log_type", logType)
	ctx := context.TODO()

	log.Debug("attempting to register a new log session")

	env, err := s.Database.GetEnvironmentByKey(ctx, envKey)
	if err != nil {
		return nil, fmt.Errorf("error getting environment from database: %w", err)
	}

	log = log.With("env_name", env.Name, "env_title", env.Title)
	log.Debug("found matching environment")

	sesKey, err := s.Database.AddNewLogSession(ctx, database.AddNewLogSessionParams{
		Environment: env.ID,
		LogType:     database.LogClientType(logType),
	})

	if err != nil {
		return nil, fmt.Errorf("error creating new log session: %w", err)
	}

	log = log.With("session_key", sesKey)
	log.Debug("log session added")

	return &sesKey, nil
}
