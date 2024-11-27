package logchunk

import (
	"context"
	"slices"
	"strconv"
	"strings"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// ChunkStreamer is the type used to stream chunks.
type ChunkStreamer func(chunk string) (int, error)

// StreamUnfilteredLogChunks fetches the log chunks and streams them without filtering
// NOTE: the ID is not validated, this should be done in the calling function.
func StreamUnfilteredLogChunks(ctx context.Context, s *interactor.Services, logID int64, streamChunk ChunkStreamer) error {
	log := s.GetDomainLogger("logs", "StreamUnfilteredLogChunks").With("log_id", logID)

	log.Debug("start chunk streaming")

	// Get chunks from database
	chunks, err := s.Database.GetLogChunksForLog(ctx, logID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	// Loop over log, and stream chunks
	for i, c := range chunks {
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

// StreamFilteredLogChunks fetches the log chunks and streams them after filtering them with the setting in LogStreamLineFilters, the absolute line numbers are added as prefixes with {int}.
// NOTE: the ID is not validated, this should be done in the calling function.
func StreamFilteredLogChunks(ctx context.Context, s *interactor.Services, logID int64, filters LogStreamLineFilters, streamChunk ChunkStreamer) error {
	log := s.GetDomainLogger("logs", "StreamLogChunks").With("log_id", logID)

	log.Debug("start chunk streaming")

	// Get chunks from database
	chunks, err := s.Database.GetLogChunksForLog(ctx, logID)
	if err != nil {
		return entities.NewDomainErrorFromDatabaseError(err)
	}

	lineCounter := 0 // Line counts start at 1, but we increment when we find a \n from 0 to 1

	// Loop over log, and stream chunks
	for i, c := range chunks {
		logLoop := log.With("i_chunk", i, "blob_path", c.BlobPath, "created_on", c.CreatedOn)

		shouldStreamChunk := filters.shouldStreamChunk(logChunkMetadata{
			// Note: we omit some fields here, as they are not used
			SeveritiesCount: c.SeverityCounts,
			CategoriesCount: c.CategoryCounts,
		})

		if !shouldStreamChunk {
			lineCounter += int(c.LineCount)

			logLoop.Info("shouldStreamChunk returned false, skipping chunk processing/streaming")

			continue
		}

		chunkText, err := s.LogChunks.GetChunk(c.BlobPath)
		if err != nil {
			logLoop.Errorf("error fetching chunk: %s", err)

			return entities.NewDomainError(entities.DomainErrorUnexpected, "error getting chunk from storage", err)
		}

		logLoop.Debug("fetched chunk contents")

		filteredLines := filterLinesForChunk(chunkText, filters, &lineCounter)

		if len(filteredLines) == 0 {
			logLoop.Warn("no contents after filtering chunk, not streaming data")

			continue
		}

		bytesWritten, err := streamChunk(string(filteredLines))

		if err != nil {
			logLoop.Errorf("error streaming filtered chunk: %s", err)

			return entities.NewDomainError(entities.DomainErrorUnexpected, "error streaming chunk", err)
		}

		logLoop.Debugf("streamed %d bytes", bytesWritten)
	}

	return nil
}

func filterLinesForChunk(chunk []byte, filters LogStreamLineFilters, lineCounter *int) []byte {
	filteredLines := make([]byte, 0, 16*1024)
	currentLineContents := make([]byte, 0, 1024)

	// Loop over the chunk in O(1), when we find a \n character, we handle the line
	for _, c := range chunk {
		currentLineContents = append(currentLineContents, c) // The \n characters get added as well, so we don't have to add those manually

		// If we don't encounter a new line, continue loop
		if c != '\n' {
			continue
		}

		// We encountered a \n, so handle the line
		*lineCounter++

		lineMetadata, err := extractMetadataFromLine(currentLineContents)

		includeLine := err == nil && lineMetadata.isComplete() && shouldIncludeLineBasedOnFilters(lineMetadata, filters)

		// Build prefix with the absolute line number
		sb := strings.Builder{}
		sb.WriteByte('{')
		sb.WriteString(strconv.Itoa(*lineCounter))
		sb.WriteByte('}')

		prefix := []byte(sb.String())

		if includeLine {
			filteredLines = append(filteredLines, prefix...)              // Write the absolute line prefix
			filteredLines = append(filteredLines, currentLineContents...) // Linebreak is in currentLineContents
		}

		currentLineContents = make([]byte, 0, 1024)
	}

	return filteredLines
}

func shouldIncludeLineBasedOnFilters(lineMetadata logChunkLineMetadata, filters LogStreamLineFilters) bool {
	// If severity filters are enabled, we filter out lines that don't fit the filters
	if len(filters.IncludedSeverities) > 0 && !slices.Contains(filters.IncludedSeverities, lineMetadata.Severity) {
		return false
	}

	// If we have included categories, we ignore excluded categories
	if len(filters.IncludedCategories) > 0 {
		return slices.Contains(filters.IncludedCategories, lineMetadata.Category)
	}

	// So we don't have included categories, we might have excluded categories
	if len(filters.ExcludedCategories) > 0 && slices.Contains(filters.ExcludedCategories, lineMetadata.Category) {
		return false
	}

	// We don't have excluded categories, or the category is not excluded, so we can return true because all other checks passed
	return true
}
