package logs

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// GetMetadataForLog fetches the metadata for a given log.
func GetMetadataForLog(s *interactor.Services, logUUID uuid.UUID) (*entities.LogMetadata, error) {
	log := s.GetDomainLogger("logs", "GetMetadataForLog")
	ctx := context.TODO()

	// Get from database
	logData, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("log_id", logData.ID)
	log.Debug("fetched log from database")

	// Get additional log metadata
	additionalMetadataDD, err := s.Database.GetMetadataForLog(ctx, logData.ID)
	if err != nil {
		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("additional_metadata_count", len(additionalMetadataDD))
	log.Debug("fetched additional metadata from database")

	additionalMetadata := make([]entities.LogAdditionalMetadata, len(additionalMetadataDD))

	// Convert metadata to return value
	for i := range additionalMetadataDD {
		additionalMetadata[i].SavedOn = additionalMetadataDD[i].SavedOn

		err = json.Unmarshal(additionalMetadataDD[i].Metadata, &additionalMetadata[i].Metadata)
		if err != nil {
			return nil, entities.NewDomainError(entities.DomainErrorUnexpected, "cannot extract metadata", err)
		}
	}

	// Get linked logs
	linkedLogsDB, err := s.Database.GetLinkedLogsForLog(ctx, logData.ID)
	if err != nil {
		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("linked_logs_count", len(linkedLogsDB))
	log.Debug("fetched linked logs from database")

	linkedLogs := make([]entities.LogLink, len(linkedLogsDB))

	for i := range linkedLogsDB {
		linkedLogs[i].LinkedLog = linkedLogsDB[i].LinkedLog
		linkedLogs[i].Description = linkedLogsDB[i].Description
	}

	// Assemble the log metadata
	metadata := entities.LogMetadata{
		AdditionalMetadata: additionalMetadata,
		Links:              linkedLogs,
	}

	log.Info("metadata fetching completed")

	return &metadata, nil
}
