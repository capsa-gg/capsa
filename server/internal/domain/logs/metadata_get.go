package logs

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
	"github.com/capsa-gg/capsa/server/internal/util"
)

// GetMetadataForLog fetches the metadata for a given log.
func GetMetadataForLog(ctx context.Context, s *interactor.Services, logUUID uuid.UUID) (*bodies.LogMetadata, error) { //nolint:funlen,gocyclo // For now this is fine, we might want to abstract the struct conversion later
	log := s.GetDomainLogger("logs", "GetMetadataForLog").With("log_uuid", logUUID)

	// Get from database
	logData, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("log_id", logData.ID)
	log.Debug("fetched log from database")

	// Get additional log metadata
	additionalMetadataDD, err := s.Database.GetMetadataForLog(ctx, logData.ID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("additional_metadata_count", len(additionalMetadataDD))
	log.Debug("fetched additional metadata from database")

	additionalMetadata := make([]bodies.LogAdditionalMetadata, len(additionalMetadataDD))

	// Convert metadata to return value
	for i := range additionalMetadataDD {
		additionalMetadata[i].SavedOn = additionalMetadataDD[i].SavedOn

		err = json.Unmarshal(additionalMetadataDD[i].Metadata, &additionalMetadata[i].Metadata)
		if err != nil {
			return nil, domainerror.New(domainerror.Unexpected, "cannot extract metadata", err)
		}
	}

	// Get linked logs
	linkedLogsDB, err := s.Database.GetLinkedLogsForLog(ctx, logData.ID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("linked_logs_count", len(linkedLogsDB))
	log.Debug("fetched linked logs from database")

	linkedLogs := make([]bodies.LogLink, len(linkedLogsDB))

	for i := range linkedLogsDB {
		linkedLogs[i].LinkedLog = linkedLogsDB[i].LinkedLog
		linkedLogs[i].Description = linkedLogsDB[i].Description
	}

	// Assemble the log metadata
	metadata := bodies.LogMetadata{
		AdditionalMetadata: additionalMetadata,
		Links:              linkedLogs,
	}

	// Get log data from database
	rows, err := s.Database.ListAvailableLogs(ctx, database.ListAvailableLogsParams{
		FilterByLogUuid: &logUUID,
		Fetchlimit:      1,
	})
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	if len(rows) != 1 {
		log.Errorf("logData query yielded %d results, hasError 1", len(rows))
	} else {
		logDataResponse := rows[0]

		metadata.LogData.LogUUID = logDataResponse.LogUuid
		metadata.LogData.LogType = constants.LogType(logDataResponse.LogType) // safe conversion
		metadata.LogData.Title = logDataResponse.Title
		metadata.LogData.Environment = logDataResponse.Environment
		metadata.LogData.Platform = logDataResponse.Platform
		metadata.LogData.LineCount = logDataResponse.LineCount
		metadata.LogData.ChunkCount = logDataResponse.ChunkCount
		metadata.LogData.LinkedLogCount = logDataResponse.LinkCount

		earliest := logDataResponse.Earliest

		if earliest != nil {
			earliestTS, err := util.ExtractTimeFromAny(earliest)
			if err != nil {
				log.Errorf("could not convert earlist value %#v to time: %s", earliest, err)
			} else {
				metadata.LogData.TimestampFirstLine = &earliestTS
			}
		}

		last := logDataResponse.Last

		if last != nil {
			lastTS, err := util.ExtractTimeFromAny(last)
			if err != nil {
				log.Errorf("could not convert last value %#v to time: %s", last, err)
			} else {
				metadata.LogData.TimestampLastLine = &lastTS
			}
		}

		if err = json.Unmarshal(logDataResponse.CategoriesCount, &metadata.LogData.CategoriesCounts); err != nil {
			log.Errorf("cannot convert CategoriesCount to map[string]int: %s", err)
		}

		if err = json.Unmarshal(logDataResponse.SeveritiesCount, &metadata.LogData.SeveritiesCounts); err != nil {
			log.Errorf("cannot convert SeveritiesCounts to map[string]int: %s", err)
		}
	}

	log.Info("metadata fetching completed")

	return &metadata, nil
}
