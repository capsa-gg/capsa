package logchunk

import (
	"context"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// ChunkStreamer is the type used to stream chunks.
type ChunkStreamer func(chunk string) (int, error)

// StreamLogChunks fetches the log chunks and streams them.
// NOTE: the ID is not validated, this should be done in the calling function.
func StreamLogChunks(ctx context.Context, s *interactor.Services, logID int64, streamChunk ChunkStreamer) error {
	log := s.GetDomainLogger("logs", "StreamLogChunks").With("log_id", logID)

	log.Debug("start chunk streaming")

	// Get chunks from database
	chunks, err := s.Database.GetLogChunksForLog(ctx, logID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Loop over log, and stream chunks
	for i, c := range chunks { //nolint:gocritic // 144 bytes, for now this is fine
		logLoop := log.With("i_chunk", i, "blob_path", c.BlobPath, "created_on", c.CreatedOn)

		chunkText, err := s.LogChunks.GetChunk(c.BlobPath)
		if err != nil {
			logLoop.Errorf("error fetching chunk: %s", err)

			return entities.NewDomainError(entities.DomainErrorUnexpected, "error getting chunk from storage", err)
		}

		logLoop.Debug("fetched chunk contents")

		bytesWritten, err := streamChunk(string(chunkText))
		if err != nil {
			logLoop.Errorf("error streaming chunk: %s", err)

			return entities.NewDomainError(entities.DomainErrorUnexpected, "error streaming chunk", err)
		}

		logLoop.Debugf("streamed %d bytes", bytesWritten)
	}

	return nil
}
