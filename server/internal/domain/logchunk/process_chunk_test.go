package logchunk

import (
	_ "embed"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

//go:embed testdata/chunk_short.log
var shortChunk []byte

//go:embed testdata/chunk_1k.log
var oneThousandLineChunk []byte

//go:embed testdata/chunk_10k.log
var tenThousandsLineChunk []byte

//go:embed testdata/chunk_100k.log
var hundredThousandsLineChunk []byte

// TODO(lucianonooijen): Add test cases for
//   - having multi line log lines (so missing ts)
//   - empty chunk
//   - malformed chunk
//   - benchmarking logs of various sizes

func TestExtractMetadataFromChunk(t *testing.T) { //nolint:funlen // This is fine
	tests := map[string]struct {
		input            []byte
		start            string
		end              string
		lineCount        int32
		unprocessedLines []string
		categories       map[string]int
		severities       map[string]int
	}{
		"ShortFullLog": {
			input:            shortChunk,
			start:            "2024.11.14-16.08.23.317",
			end:              "2024.11.14-16.08.25.652",
			lineCount:        43,
			unprocessedLines: []string{},
			categories: map[string]int{
				"LogStats":            1,
				"LogTextureExample":   1,
				"LogWorldMetrics":     4,
				"LogInit":             2,
				"LogAnalytics":        1,
				"LogAudio":            6,
				"LogAudioMixer":       6,
				"SourceControl":       1,
				"LogUnrealEdMisc":     2,
				"Cmd":                 2,
				"MapCheck":            1,
				"LogWorld":            1,
				"LogSlate":            4,
				"LogWorldPartition":   1,
				"LogUObjectHash":      1,
				"LogAssetRegistry":    1,
				"LogStall":            2,
				"LogNetVersion":       1,
				"LogLoad":             1,
				"LogContentStreaming": 1,
				"LogCsvProfiler":      1,
				"LogCapsaCore":        1,
				"LogLiveCoding":       1,
			},
			severities: map[string]int{
				"Log":         23,
				"Display":     17,
				"Warning":     1,
				"Verbose":     1,
				"VeryVerbose": 1,
			},
		},
		"Chunk_1k": {
			input:            oneThousandLineChunk,
			start:            "2024.11.15-22.00.00.000",
			end:              "2024.11.15-22.00.41.958",
			lineCount:        1_000,
			unprocessedLines: []string{},
			categories:       map[string]int{"LogCategoryExample": 1_000},
			severities:       map[string]int{"Log": 1_000},
		},
		"Chunk_10k": {
			input:            tenThousandsLineChunk,
			start:            "2024.11.15-22.00.00.000",
			end:              "2024.11.15-22.06.59.958",
			lineCount:        10_000,
			unprocessedLines: []string{},
			categories:       map[string]int{"LogCategoryExample": 10_000},
			severities:       map[string]int{"Log": 10_000},
		},
		"Chunk_100k": {
			input:            hundredThousandsLineChunk,
			start:            "2024.11.15-22.00.00.000",
			end:              "2024.11.15-23.09.59.958",
			lineCount:        100_000,
			unprocessedLines: []string{},
			categories:       map[string]int{"LogCategoryExample": 100_000},
			severities:       map[string]int{"Log": 100_000},
		},
	}

	for tt, tData := range tests {
		t.Run(tt, func(t *testing.T) {
			expectedStart, err := time.Parse(timestampParseLayout, tData.start)
			require.NoError(t, err)

			expectedEnd, err := time.Parse(timestampParseLayout, tData.end)
			require.NoError(t, err)

			metadata, unprocessedLines := extractMetadataFromChunk(zap.NewNop().Sugar(), tData.input)

			require.Equal(t, expectedStart, metadata.Start)
			require.Equal(t, expectedEnd, metadata.End)
			require.Equal(t, tData.lineCount, metadata.LineCount)
			require.Equal(t, tData.categories, metadata.CategoriesCount)
			require.Equal(t, tData.severities, metadata.SeveritiesCount)
			require.Equal(t, tData.unprocessedLines, unprocessedLines)
		})
	}
}

//nolint:funlen // Long due to tests declarations
func TestExtractMetadataFromLine(t *testing.T) {
	tests := map[string]struct {
		input      []byte
		err        error
		isComplete bool
		timestamp  time.Time
		severity   string
		category   string
	}{
		"Valid": {
			input:      []byte("[2024.11.14-16.08.25.609][Log][LogCapsaCore]: Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        nil,
			isComplete: true,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "Log",
			category:   "LogCapsaCore",
		},
		"Invalid_EmptyLine": {
			input:      []byte(""),
			err:        errorLineEmpty,
			isComplete: false,
			timestamp:  time.Time{},
			severity:   "",
			category:   "",
		},
		"Invalid_IncorrectStart_Text": {
			input:      []byte("Invalid log line"),
			err:        errorLineNoValidStart,
			isComplete: false,
			timestamp:  time.Time{},
			severity:   "",
			category:   "",
		},
		"Invalid_IncorrectStart_Space": {
			input:      []byte("    === some info continued from the previous line"),
			err:        errorLineNoValidStart,
			isComplete: false,
			timestamp:  time.Time{},
			severity:   "",
			category:   "",
		},
		"Invalid_ExtraOpenBracket": {
			input:      []byte("[2024.11.14-16.08.25.609][[Log][LogCapsaCore]: Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineInvalidMetadataChar,
			isComplete: false,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "",
			category:   "",
		},
		"Invalid_CutOff": {
			input:      []byte("[2024.11.14-16.08.25.609][Log][LogCaps"),
			err:        errorLineUnexpectedEnd,
			isComplete: false,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "Log",
			category:   "",
		},
		"Invalid_EmptyMetadata": {
			input:      []byte("[2024.11.14-16.08.25.609][Log][]: Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineEmptyMetadataBlock,
			isComplete: false,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "Log",
			category:   "",
		},
		"Invalid_SpacePaddingBetweenMetadata": {
			input:      []byte("[2024.11.14-16.08.25.609][Log] [LogCapsaCore]: Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineInvalidMetadataString,
			isComplete: false,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "",
			category:   "",
		},
		"Invalid_CharactersBetweenMetadataBlocks": {
			input:      []byte("[2024.11.14-16.08.25.609]Log[LogCapsaCore]: Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineInvalidMetadataString,
			isComplete: false,
			timestamp:  time.Time{},
			severity:   "",
			category:   "",
		},
		"Invalid_IncorrectTimeFormat": {
			input:      []byte("[2024-11-14.16-08-25-609][Log][LogCapsaCore]: Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineTimestampParsing,
			isComplete: false,
			timestamp:  time.Time{},
			severity:   "",
			category:   "",
		},
		"Invalid_IncorrectEndOfMetadata_MissingSemicolon": {
			input:      []byte("[2024.11.14-16.08.25.609][Log][LogCapsaCore] Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineMissingEndOfMetadataChar,
			isComplete: true,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "Log",
			category:   "LogCapsaCore",
		},
		"Invalid_IncorrectEndOfMetadata_SpaceBeforeSemicolon": {
			input:      []byte("[2024.11.14-16.08.25.609][Log][LogCapsaCore] : Capsa ID: 00000000-0000-0000-0000-000000000000 | CapsaLogURL: https://example.dev/logs/00000000-0000-0000-0000-000000000000"),
			err:        errorLineMissingEndOfMetadataChar,
			isComplete: true,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "Log",
			category:   "LogCapsaCore",
		},
		"Invalid_IncorrectEndOfMetadata_OnlyMetadata": {
			input:      []byte("[2024.11.14-16.08.25.609][Log][LogCapsaCore]"),
			err:        errorLineMissingEndOfMetadataChar,
			isComplete: true,
			timestamp:  time.Date(2024, 11, 14, 16, 8, 25, 609000000, time.UTC),
			severity:   "Log",
			category:   "LogCapsaCore",
		},
	}

	for tt, tData := range tests {
		t.Run(tt, func(t *testing.T) {
			got, err := extractMetadataFromLine(tData.input)

			if tData.err != nil {
				require.ErrorIs(t, tData.err, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tData.timestamp, got.Timestamp)
			require.Equal(t, tData.severity, got.Severity)
			require.Equal(t, tData.category, got.Category)
			require.Equal(t, tData.isComplete, got.isComplete())
		})
	}
}
