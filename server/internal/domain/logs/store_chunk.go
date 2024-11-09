package logs

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// StoreLogChunk stores log chunks.
// TODO: make this ProcessLogChunk and add processing.
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

	// Assemble database params
	addLogChunkParams := database.AddLogChunkParams{
		Log:      logInfo.ID,
		BlobPath: fileName,

		// TODO: add processing
		LineCount:      0,
		ChunkStart:     sql.NullTime{},
		ChunkEnd:       sql.NullTime{},
		CategoryCounts: []byte("{}"),
		SeverityCounts: []byte("{}"),
	}

	// Store log in database
	err = s.Database.AddLogChunk(ctx, addLogChunkParams)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	return nil
}
