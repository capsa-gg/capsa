//nolint:nilnil // the interface has documented that the value can be nil
package logs_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/domain/logchunk"
	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// NOTE: when calling the API endpoint, the log lines would also include the absolute line number.
// For these tests, we are calling the logchunk.ExtractMetadataFromLine function, which would return an error with absolute line numbers present.
// This current set setup is fine, as the merge function does not care about what the lines look like, it just needs the metadata and raw line.
// We are simply using logchunk.ExtractMetadataFromLine for convenience.

var logsServer = []string{
	"[2024.12.11-20.00.00.000][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.200][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.300][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.500][Log][LogCapsaTest]: TestLogLine",
}

var logsClientOne = []string{
	"[2024.12.11-20.00.00.100][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.200][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.400][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.600][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.750][Log][LogCapsaTest]: TestLogLine",
}

var logsClientTwo = []string{
	"[2024.12.11-20.00.00.200][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.400][Log][LogCapsaTest]: TestLogLine",
	"[2024.12.11-20.00.00.700][Log][LogCapsaTest]: TestLogLine",
}

const logsOutputExpected = `(S1)[2024.12.11-20.00.00.000][Log][LogCapsaTest]: TestLogLine
(C1)[2024.12.11-20.00.00.100][Log][LogCapsaTest]: TestLogLine
(S1)[2024.12.11-20.00.00.200][Log][LogCapsaTest]: TestLogLine
(C1)[2024.12.11-20.00.00.200][Log][LogCapsaTest]: TestLogLine
(C2)[2024.12.11-20.00.00.200][Log][LogCapsaTest]: TestLogLine
(S1)[2024.12.11-20.00.00.300][Log][LogCapsaTest]: TestLogLine
(C1)[2024.12.11-20.00.00.400][Log][LogCapsaTest]: TestLogLine
(C2)[2024.12.11-20.00.00.400][Log][LogCapsaTest]: TestLogLine
(S1)[2024.12.11-20.00.00.500][Log][LogCapsaTest]: TestLogLine
(C1)[2024.12.11-20.00.00.600][Log][LogCapsaTest]: TestLogLine
(C2)[2024.12.11-20.00.00.700][Log][LogCapsaTest]: TestLogLine
(C1)[2024.12.11-20.00.00.750][Log][LogCapsaTest]: TestLogLine
`

type filteredLineLoader struct {
	nextLine int
	input    []string
}

func (f *filteredLineLoader) HasNextLine() (bool, error) {
	return len(f.input) > f.nextLine, nil
}

func (f *filteredLineLoader) ReadNextLineMetadata() (lineMetadata *entities.LogChunkLineMetadata, err error) {
	hasNext, _ := f.HasNextLine()

	if !hasNext {
		return nil, nil
	}

	line := f.input[f.nextLine]

	metadata, err := logchunk.ExtractMetadataFromLine([]byte(line))

	// We should not be getting any extraction errors, the input is valid
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (f *filteredLineLoader) GetNextLine() (logLine *string, err error) {
	hasNext, _ := f.HasNextLine()

	if !hasNext {
		return nil, nil
	}

	line := f.input[f.nextLine]

	f.nextLine++

	return &line, nil
}

func getTestFilteredLineLoader(input []string) entities.FilteredLineLoader {
	return &filteredLineLoader{0, input}
}

func getMergedLogInput() entities.MergedLogInput {
	return entities.MergedLogInput{
		{Key: "S1", Loader: getTestFilteredLineLoader(logsServer)},
		{Key: "C1", Loader: getTestFilteredLineLoader(logsClientOne)},
		{Key: "C2", Loader: getTestFilteredLineLoader(logsClientTwo)},
	}
}

var testInteractor = &interactor.Services{
	Config: &entities.Config{
		RootLogger: zap.NewNop(),
	},
}

func TestStreamMergedLog(t *testing.T) {
	tests := map[string]struct {
		input            entities.MergedLogInput
		maxLinesPerChunk uint64
		expected         string
	}{
		"OneChunk": {
			// Test for the base logic working
			input:            getMergedLogInput(),
			maxLinesPerChunk: 1_000_000,
			expected:         logsOutputExpected,
		},
		"SmallChunks": {
			// Test if all lines get flushed correctly
			input:            getMergedLogInput(),
			maxLinesPerChunk: 2,
			expected:         logsOutputExpected,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			output := ""

			var chunkStreamer = func(chunk string) (int, error) {
				output += chunk
				return 0, nil
			}

			err := logs.StreamMergedLog(context.Background(), testInteractor, tt.input, tt.maxLinesPerChunk, chunkStreamer)

			require.NoError(t, err)
			require.Equal(t, tt.expected, output)
		})
	}
}
