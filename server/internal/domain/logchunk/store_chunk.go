package logchunk

import (
	"context"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// StoreLogChunk stores log chunks.
// TODO: can logData be passed or do we need a reader?
func StoreLogChunk(s *interactor.Services, logUUID uuid.UUID, logData []byte) error {
	log := s.GetDomainLogger("logs", "StoreLogChunk")
	ctx := context.Background() // We don't want to cancel the log storing procedure if the connection gets lost

	log.Debug("storing chunk")

	// Get log from database
	logInfo, err := s.Database.GetLogByUuid(ctx, logUUID)
	if err != nil {
		return domainerror.NewFromDatabaseError(err)
	}

	log = log.
		With("log_id", logInfo.ID).
		With("log_type", logInfo.LogType).
		With("platform", logInfo.Platform)

	log.Debug("log found")

	// Store log in storage
	fileName, err := s.LogChunks.SaveChunk(logData)
	if err != nil {
		return domainerror.New(domainerror.Unexpected, "cannot store chunk in storage", err)
	}

	chunkMetadata, unprocessedLines := extractMetadataFromChunk(log, logData)

	totalLines := chunkMetadata.LineCount + int32(len(unprocessedLines)) //nolint:gosec // We will not have more than 2.1B unprocessed lines (G115: integer overflow conversion int -> int32)

	log = log.
		With("len_unprocessed_lines", len(unprocessedLines)).
		With("len_total_lines", totalLines).
		With("ts_first_line", chunkMetadata.Start).
		With("ts_last_line", chunkMetadata.End).
		With("category_counts_len", len(chunkMetadata.CategoriesCount)).
		With("severities_counts_len", len(chunkMetadata.SeveritiesCount))

	if len(unprocessedLines) > 0 {
		log.Warn("some lines not processed") // Logged out in extractMetadataFromChunk
	} else {
		log.Debug("metadata extracted with all lines processed")
	}

	// Assemble database params
	addLogChunkParams := database.AddLogChunkParams{
		Log:            logInfo.ID,
		BlobPath:       fileName,
		LineCount:      chunkMetadata.LineCount,
		ChunkStart:     &chunkMetadata.Start,
		ChunkEnd:       &chunkMetadata.End,
		CategoryCounts: chunkMetadata.CategoriesCount,
		SeverityCounts: chunkMetadata.SeveritiesCount,
	}

	// Store log chunk metadata in database
	err = s.Database.AddLogChunk(ctx, addLogChunkParams)
	if err != nil {
		return domainerror.NewFromDatabaseError(err)
	}

	log.Info("log chunk processed")

	return nil
}
