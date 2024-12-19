package logchunk

import (
	"time"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

// ExtractMetadataFromLine extracts the metadata the input line.
// In case a line cannot be parsed, the err value will be populated.
//
//nolint:funlen,nakedret,gocognit,gocyclo // This is easier to comprehend as a single function with naked returns, the function is somewhat complex, but hard to simplify, the alternative is a lexer/parser
func ExtractMetadataFromLine(line []byte) (lineMetadata entities.LogChunkLineMetadata, err error) {
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
			case 2: // Severity, we don't check for validity
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
