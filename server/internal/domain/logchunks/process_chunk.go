package logchunk

import (
	"errors"
	"time"

	"go.uber.org/zap"
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

const timestampParseLayout = "2006.01.02-15.04.05.999"

type logChunkLineMetadata struct {
	Timestamp time.Time
	Category  string
	Severity  string
}

func (l logChunkLineMetadata) isComplete() bool {
	return l.Severity != "" && l.Category != "" && !l.Timestamp.IsZero()
}

type logChunkMetadata struct {
	Start           time.Time
	End             time.Time
	LineCount       int32
	CategoriesCount map[string]int // Used to get a string[] of categories
	SeveritiesCount map[string]int
}

// Should only be called when the line processing was successful and the line .isComplete().
func (lcm *logChunkMetadata) addLineMetadata(lineMetadata logChunkLineMetadata) {
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

		lineMetadata, err := extractMetadataFromLine(lineContents)

		switch {
		case err != nil: // Metadata is complete and should be processed
			log.With(
				zap.Int("line", i),
				zap.ByteString("lineContents", lineContents),
				zap.Error(err),
			).Warn("error extracting metadata from line, not processing")

			unprocessedLines = append(unprocessedLines, string(lineContents))
		case !lineMetadata.isComplete(): // Don't process incomplete lines
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

//nolint:funlen,nakedret,gocognit,gocyclo // This is easier to comprehend as a single function with naked returns, the function is somewhat complex, but hard to simplify, the alternative is a lexer/parser
func extractMetadataFromLine(line []byte) (lineMetadata logChunkLineMetadata, err error) {
	// Check if we have content
	if len(line) == 0 {
		err = errorLineEmpty

		return
	}

	// Lines must start with a `[` to be considered valid
	if line[0] != '[' {
		err = errorLineNoValidStart

		return
	}

	isInMetadataBlock := false
	metadataBlockCount := 0                      // We use 1-based indexing for this in the logic below
	metadataBlockContents := make([]byte, 0, 64) // Assuming we don't go over 64 characters

	// Loop over characters, not using regex for performance
charLoop:
	for i, char := range line {
		hasNext := len(line) > i+1

		if char == '[' {
			// If we get `[` without a `]` first, we cannot parse the metadata correctly: failing example: `[Timestamp][[Severity]`
			if isInMetadataBlock {
				err = errorLineInvalidMetadataChar

				return
			}

			// Start reading the data for the metadata block
			isInMetadataBlock = true
			metadataBlockCount++

			continue // Do not add the character to metadataBlockContents
		}

		// End of metadata block
		if char == ']' {
			isInMetadataBlock = false

			// Metadata contents cannot be empty, failing example: `[TimestampValue][][LogCategory]`
			if len(metadataBlockContents) == 0 {
				err = errorLineEmptyMetadataBlock

				return
			}

			// Metadata blocks must be `][` without any other characters between them, failing example: `[Timestamp] [Severity]`
			// The len check here is necessary to prevent index out of range panic
			if metadataBlockCount < 3 && (!hasNext || line[i+1] != '[') { // peek next character
				err = errorLineInvalidMetadataString

				return
			}

			// Valid metadata block completion
			switch metadataBlockCount {
			case 1: // Timestamp
				ts, errTS := time.Parse(timestampParseLayout, string(metadataBlockContents))

				if errTS != nil {
					err = errorLineTimestampParsing

					return
				}

				lineMetadata.Timestamp = ts
			case 2: // Severity
				lineMetadata.Severity = string(metadataBlockContents)
			case 3: // Category
				lineMetadata.Category = string(metadataBlockContents)

				// After the metadata, there should be a `:` character, failing example: ``
				// The len check here is necessary to prevent index out of range panic
				if !hasNext || line[i+1] != ':' {
					err = errorLineMissingEndOfMetadataChar

					return
				}

				// We have three metadata blocks, so we break the loop for final validation
				break charLoop
			}

			// Reset metadataBlockContents
			metadataBlockContents = make([]byte, 0, 64)

			continue // Do not add the character to metadataBlockContents
		}

		// Append character to the contents
		metadataBlockContents = append(metadataBlockContents, char)
	}

	// Make sure the metadata was correctly closed
	if isInMetadataBlock {
		err = errorLineUnexpectedEnd

		return
	}

	return
}
