package logchunk

import (
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

var (
	errorLineEmpty                    = errors.New("line is empty")
	errorLineNoValidStart             = errors.New("character is not a valid start of a line, required: '['")
	errorLineInvalidMetadataChar      = errors.New("metadata parsing failed due to invalid character")
	errorLineUnexpectedEnd            = errors.New("metadata parsing failed due to unexpected end of data")
	errorLineEmptyMetadataBlock       = errors.New("metadata parsing failed due to empty metadata block")
	errorLineInvalidMetadataString    = errors.New("metadata blocks should have no characters between them")
	errorLineTimestampParsing         = errors.New("timestamp cannot be parsed due to incorrect formatting")
	errorLineMissingEndOfMetadataChar = errors.New("metadata blocks should be ended with a `:` character")
)

const timestampParseLayout = "2006.01.02-15.04.05.000"

type logChunkMetadata struct {
	Start           time.Time
	End             time.Time
	LineCount       int32
	CategoriesCount map[string]int // Used to get a string[] of categories
	SeveritiesCount map[string]int
}

// Should only be called when the line processing was successful and the line .IsComplete().
func (lcm *logChunkMetadata) addLineMetadata(lineMetadata entities.LogChunkLineMetadata) {
	if lcm.Start.IsZero() || lineMetadata.Timestamp.Before(lcm.Start) {
		lcm.Start = lineMetadata.Timestamp
	}

	if lineMetadata.Timestamp.After(lcm.End) {
		lcm.End = lineMetadata.Timestamp
	}

	lcm.CategoriesCount[lineMetadata.Category]++
	lcm.SeveritiesCount[lineMetadata.Severity]++
}

// extractMetadataFromChunk takes in a log chunk as []byte and returns the metadata,
// the unprocessed/unprocessable lines are logged to logger.Warnf() as well as returned to the caller.
func extractMetadataFromChunk(logger *zap.SugaredLogger, logChunk []byte) (logChunkMetadata, []string) { //nolint:gocritic // This is fine and documented in a comment
	log := logger.Named("extractMetadataFromChunk").Desugar() // Desugar for a bit better performance
	unprocessedLines := []string{}
	chunkMetadata := logChunkMetadata{
		CategoriesCount: map[string]int{},
		SeveritiesCount: map[string]int{},
	}

	lineContents := make([]byte, 0, 512)
	lineCount := int32(0)

	// Loop over the chunk in O(1), when we find a \n character, we handle the line
	for i, c := range logChunk {
		// If we don't encounter a new line, add to lineContents and continue loop
		if c != '\n' {
			lineContents = append(lineContents, c)

			continue
		}

		// We encountered a \n, so handle the line and try and get the metadata for it
		lineCount++

		lineMetadata, err := ExtractMetadataFromLine(lineContents)

		switch {
		case err != nil: // Metadata is complete and should be processed
			log.With(
				zap.Int("line", i),
				zap.ByteString("lineContents", lineContents),
				zap.Error(err),
			).Warn("error extracting metadata from line, not processing")

			unprocessedLines = append(unprocessedLines, string(lineContents))
		case !lineMetadata.IsComplete(): // Don't process incomplete lines
			log.With(
				zap.Int("line", i),
				zap.Time("timestamp", lineMetadata.Timestamp),
				zap.String("severity", lineMetadata.Severity),
				zap.String("category", lineMetadata.Severity),
				zap.ByteString("lineContents", lineContents),
				zap.Error(err),
			).Warn("metadata from line is not complete, not processing")

			unprocessedLines = append(unprocessedLines, string(lineContents))
		default:
			chunkMetadata.addLineMetadata(lineMetadata)
		}

		// Reset lineContents
		lineContents = make([]byte, 0, 1024)
	}

	// Set the total line count, then return data
	chunkMetadata.LineCount = lineCount

	return chunkMetadata, unprocessedLines
}
