package logs

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// StoreLogChunk stores log chunks.
// TODO: can logData be passed or do we need a reader?
func StoreLogChunk(s *interactor.Services, logUUID uuid.UUID, logData []byte) error {
	log := s.GetDomainLogger("logs", "StoreLogChunk")
	ctx := context.TODO()

	log.Debug("storing chunk")

	// Get log from database
	logInfo, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.
		With("log_id", logInfo.ID).
		With("log_type", logInfo.LogType).
		With("platform", logInfo.Platform)

	log.Debug("log found")

	// Store log in storage
	fileName, err := s.LogChunks.SaveChunk(logData)
	if err != nil {
		return entities.NewDomainError(entities.DomainErrorUnexpected, "cannot store chunk in storage", err)
	}

	chunkMetadata, unprocessedLinkes := extractMetadataFromChunk(log, logData)
	if len(unprocessedLinkes) > 0 {
		log.With("len_unprocessed_lines").Warn("some lines are not processed")
	}

	categoryCounts, err := json.Marshal(chunkMetadata.CategoriesCount)
	if err != nil {
		log.With("error", err).Error("cannot convert categories count to json")
	}

	severitiesCount, err := json.Marshal(chunkMetadata.SeveritiesCount)
	if err != nil {
		log.With("error", err).Error("cannot convert severities count to json")
	}

	// Assemble database params
	addLogChunkParams := database.AddLogChunkParams{
		Log:            logInfo.ID,
		BlobPath:       fileName,
		LineCount:      chunkMetadata.LineCount,
		ChunkStart:     sql.NullTime{Valid: true, Time: chunkMetadata.Start},
		ChunkEnd:       sql.NullTime{Valid: true, Time: chunkMetadata.End},
		CategoryCounts: categoryCounts,
		SeverityCounts: severitiesCount,
	}

	// Store log chunk metadata in database
	err = s.Database.AddLogChunk(ctx, addLogChunkParams)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Update the log timestamps based on chunk metadata, logic for whether fields get updated is in the sql
	err = s.Database.UpdateLogTimestamps(ctx, database.UpdateLogTimestampsParams{
		LogUuid:  logInfo.LogUuid,
		LogStart: chunkMetadata.Start,
		LogEnd:   chunkMetadata.End,
	})
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.With("chunk_ts_start", chunkMetadata.Start.String()).With("chunk_ts_end", chunkMetadata.End.String()).Debug("timestamps sent to database for processing")

	return nil
}
