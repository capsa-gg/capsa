package logs

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// SaveLogMetadata saves the log metadata and attempts to link the linked logs.
func SaveLogMetadata(ctx context.Context, s *interactor.Services, logUUID uuid.UUID, metadata map[string]any, linkedLogs map[uuid.UUID]string) error {
	log := s.GetDomainLogger("logs", "SaveLogMetadata").With("log_uuid", logUUID.String())

	// Get log
	logInfo, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("log_id", logInfo.ID)
	log.Debug("found log in database")

	if len(metadata) > 0 {
		err = storeLogMetadata(s, logInfo.ID, metadata)

		if err != nil {
			return err // Error type formed in storeLogMetadata
		}
	} else {
		log.Info("no metadata stored for log, map is empty")
	}

	// Store linked logs, will not execute if map is empty
	for link, description := range linkedLogs {
		attemptSaveLogLink(s, logInfo.ID, link, description)
	}

	return nil
}

// storeLogMetadata stores the log metadata to the database.
func storeLogMetadata(s *interactor.Services, logID int64, metadata map[string]any) error {
	ctx := context.TODO()

	// Convert map to JSON
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		return entities.NewDomainError(entities.DomainErrorUnexpected, "cannot convert metadata to json", err)
	}

	// Store metadata
	err = s.Database.AddLogMetadata(ctx, database.AddLogMetadataParams{
		Log:      logID,
		Metadata: jsonData,
	})

	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	return nil
}

// attemptSaveLogLink attempts to create a link between two logs, and drops a link if an error occurs.
func attemptSaveLogLink(s *interactor.Services, logID int64, link uuid.UUID, description string) {
	log := s.GetDomainLogger("logs", "saveLogLinks").With("log_id", logID).With("link_uuid", link)
	ctx := context.TODO()

	linkedLog, err := s.Database.GetLogByUuid(ctx, link)
	if err != nil {
		log.Warnf("error getting log by uuid, link dropped: %s", err)

		return
	}

	log = log.With("link_id", linkedLog.ID)
	log.Debug("linked log found in database")

	err = s.Database.AddLogLink(ctx, database.AddLogLinkParams{
		Source:      logID,
		Link:        linkedLog.ID,
		Description: description,
	})

	if err != nil {
		log.Errorf("could nog link logs, link dropped: %s", err)
	}

	log.With("description", description).Info("log link stored")
}
