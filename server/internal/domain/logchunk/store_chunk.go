package logchunk

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

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

	chunkMetadata, unprocessedLines := extractMetadataFromChunk(log, logData)
	if len(unprocessedLines) > 0 {
		log.With("len_unprocessed_lines").Warn("some lines are not processed")
	}

	log = log.With("category_counts_len", len(chunkMetadata.CategoriesCount), "severities_counts_len", len(chunkMetadata.SeveritiesCount))
	log.Debug("chunk metadata extracted")

	// Assemble database params
	// TODO: Store the unprocessed line count in database
	addLogChunkParams := database.AddLogChunkParams{
		Log:            logInfo.ID,
		BlobPath:       fileName,
		LineCount:      chunkMetadata.LineCount,
		ChunkStart:     pgtype.Timestamp{Time: chunkMetadata.Start},
		ChunkEnd:       pgtype.Timestamp{Time: chunkMetadata.End},
		CategoryCounts: chunkMetadata.CategoriesCount,
		SeverityCounts: chunkMetadata.SeveritiesCount,
	}

	// Store log chunk metadata in database
	err = s.Database.AddLogChunk(ctx, addLogChunkParams)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.With("chunk_ts_start", chunkMetadata.Start.String()).With("chunk_ts_end", chunkMetadata.End.String()).Debug("timestamps sent to database for processing")

	return nil
}
