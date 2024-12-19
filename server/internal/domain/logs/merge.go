package logs

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

const mergeNoNextLineIndex = -1

// StreamMergedLog accepts entities.MergedLogInput as a handler to get log lines, then merges the logs.
// The input for entities.MergedLogInput should return the filtered log lines, if applicable.
// Merging is done based on the timestamp for each line.
func StreamMergedLog(ctx context.Context, s *interactor.Services, input entities.MergedLogInput, maxLinesPerChunkStream uint64, streamChunk entities.ChunkStreamer) error { //nolint:gocyclo // For now, this is fine
	if maxLinesPerChunkStream < 1 {
		return domainerror.New(domainerror.Unexpected, fmt.Sprintf("maxLinesPerChunkStream value %d not valid", maxLinesPerChunkStream), errors.New("maxLinesPerChunkStream must be >= 1"))
	}

	log := s.GetDomainLogger("logs", "StreamMergedLog")
	chunkLineCount := uint64(0)
	chunkBuffer := ""

	flush := func() error { // bool indicates success, stop execution on failure
		_, err := streamChunk(chunkBuffer)
		if err != nil {
			return err
		}

		chunkLineCount = 0
		chunkBuffer = ""

		return nil
	}

	// Loop until we manually break
	for {
		select {
		case <-ctx.Done():
			log.Warn("streaming canceled, due to ctx.Done()")

			return nil
		default:
			nextLineLogIndex, err := getNextLineLogKey(&input)

			if err != nil {
				log.Errorf("error getting next log line key: %s", err)

				return domainerror.New(domainerror.Unexpected, "error getting next line log key", err)
			}

			if nextLineLogIndex == mergeNoNextLineIndex { // No more lines
				return flush()
			}

			nextLogLine := input[nextLineLogIndex]

			lineRaw, err := nextLogLine.Loader.GetNextLine()
			if err != nil {
				return domainerror.New(domainerror.Unexpected, "error reading next line metadata", err)
			}

			if lineRaw == nil {
				return domainerror.New(domainerror.Unexpected, "lineRaw is nil\"", err)
			}

			// Build line: (<Key>)<Line>\n
			sb := strings.Builder{}
			sb.WriteByte('(')
			sb.WriteString(nextLogLine.Key)
			sb.WriteByte(')')
			sb.WriteString(*lineRaw)
			sb.WriteByte('\n')

			// Add to buffer and increment count
			chunkBuffer += sb.String()
			chunkLineCount++

			// Check if we need to stream the chunk buffer
			if chunkLineCount >= maxLinesPerChunkStream {
				flushError := flush()
				if flushError != nil {
					return flushError
				}
			}
		}
	}
}

// Returns nil if no lines are available.
func getNextLineLogKey(input *entities.MergedLogInput) (int, error) {
	var lowestTSValue *time.Time

	var lowestTSIndex = -1

	for i, l := range *input {
		logKey := l.Key
		loader := l.Loader

		hasNext, err := loader.HasNextLine()

		if err != nil {
			return mergeNoNextLineIndex, fmt.Errorf("cannot check if next line is available for logKey %s", logKey)
		}

		if !hasNext {
			continue
		}

		nextLineMetadata, err := loader.ReadNextLineMetadata()
		if err != nil {
			return mergeNoNextLineIndex, fmt.Errorf("error reading next line metadata: %w", err)
		}

		if nextLineMetadata == nil {
			return mergeNoNextLineIndex, fmt.Errorf("nil metadata returned for log key %s", logKey)
		}

		if !nextLineMetadata.IsComplete() {
			return mergeNoNextLineIndex, fmt.Errorf("nextLineMetadata for log with key %s is not complete (%#v)", logKey, nextLineMetadata)
		}

		// Set values if not set already
		if lowestTSValue == nil {
			lowestTSValue = &nextLineMetadata.Timestamp
			lowestTSIndex = i

			continue
		} // The values are set from here

		if lowestTSValue.After(nextLineMetadata.Timestamp) {
			lowestTSValue = &nextLineMetadata.Timestamp
			lowestTSIndex = i
		}
	}

	return lowestTSIndex, nil
}
