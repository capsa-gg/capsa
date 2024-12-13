package logchunk

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

func TestHasFilters(t *testing.T) {
	tests := map[string]struct {
		filters  LogStreamLineFilters
		expected bool
	}{
		"no filters": {
			filters:  LogStreamLineFilters{},
			expected: false,
		},
		"included severities": {
			filters:  LogStreamLineFilters{IncludedSeverities: []string{"Log"}},
			expected: true,
		},
		"included categories": {
			filters:  LogStreamLineFilters{IncludedCategories: []string{"LogTestExample"}},
			expected: true,
		},
		"excluded categories": {
			filters:  LogStreamLineFilters{ExcludedCategories: []string{"LogTestExample"}},
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := tc.filters.HasFilters()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestShouldIncludeLineBasedOnFilters(t *testing.T) {
	tests := map[string]struct {
		lineMetadata entities.LogChunkLineMetadata
		filters      LogStreamLineFilters
		expected     bool
	}{
		"no filters": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{},
			expected:     true,
		},
		"included severity match": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{IncludedSeverities: []string{"Log"}},
			expected:     true,
		},
		"included severity no match": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Error", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{IncludedSeverities: []string{"Log", "Verbose"}},
			expected:     false,
		},
		"included category match": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{IncludedCategories: []string{"LogTestExample"}},
			expected:     true,
		},
		"included category no match": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogWorldMetrics"},
			filters:      LogStreamLineFilters{IncludedCategories: []string{"LogTestExample"}},
			expected:     false,
		},
		"excluded category match": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{ExcludedCategories: []string{"LogTestExample"}},
			expected:     false,
		},
		"excluded category match overwritten by included category": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters: LogStreamLineFilters{
				IncludedCategories: []string{"LogCapsaCore"},
				ExcludedCategories: []string{"LogTestExample"},
			},
			expected: false,
		},
		"excluded category no match": {
			lineMetadata: entities.LogChunkLineMetadata{Severity: "Log", Category: "LogWorldMetrics"},
			filters:      LogStreamLineFilters{ExcludedCategories: []string{"LogTestExample"}},
			expected:     true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := shouldIncludeLineBasedOnFilters(tt.lineMetadata, tt.filters)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestFilterLinesForChunk(t *testing.T) {
	input := []byte(`[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.919][Warning][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.919][Log][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.955][Display][LogTestTwo]: This is a test log line
[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
[2024.11.14-16.08.25.493][VeryVerbose][LogTestTwo]: This is a test log line
`)
	tests := map[string]struct {
		filters             LogStreamLineFilters
		expectedFilteredLog []byte
	}{
		"no effective filters": {
			filters: LogStreamLineFilters{ExcludedCategories: []string{"LogTestNonExistent"}},
			expectedFilteredLog: []byte(`{1}[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
{2}[2024.11.14-16.08.23.919][Warning][LogTestOne]: This is a test log line
{3}[2024.11.14-16.08.23.919][Log][LogTestOne]: This is a test log line
{4}[2024.11.14-16.08.23.955][Display][LogTestTwo]: This is a test log line
{5}[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
{6}[2024.11.14-16.08.25.493][VeryVerbose][LogTestTwo]: This is a test log line
`),
		},
		"severity only": {
			filters: LogStreamLineFilters{IncludedSeverities: []string{"Error", "Log", "Verbose"}},
			expectedFilteredLog: []byte(`{1}[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
{3}[2024.11.14-16.08.23.919][Log][LogTestOne]: This is a test log line
{5}[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
`),
		},
		"included category": {
			filters: LogStreamLineFilters{IncludedCategories: []string{"LogTestTwo"}},
			expectedFilteredLog: []byte(`{4}[2024.11.14-16.08.23.955][Display][LogTestTwo]: This is a test log line
{5}[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
{6}[2024.11.14-16.08.25.493][VeryVerbose][LogTestTwo]: This is a test log line
`),
		},
		"excluded category": {
			filters: LogStreamLineFilters{ExcludedCategories: []string{"LogTestTwo"}},
			expectedFilteredLog: []byte(`{1}[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
{2}[2024.11.14-16.08.23.919][Warning][LogTestOne]: This is a test log line
{3}[2024.11.14-16.08.23.919][Log][LogTestOne]: This is a test log line
`),
		},
		"included category overriding excluded category": {
			filters: LogStreamLineFilters{
				IncludedCategories: []string{"LogTestTwo"},
				ExcludedCategories: []string{"LogTestTwo"},
			},
			expectedFilteredLog: []byte(`{4}[2024.11.14-16.08.23.955][Display][LogTestTwo]: This is a test log line
{5}[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
{6}[2024.11.14-16.08.25.493][VeryVerbose][LogTestTwo]: This is a test log line
`),
		},
		"mixed": {
			filters: LogStreamLineFilters{
				IncludedSeverities: []string{"Error", "Warning", "Verbose"},
				IncludedCategories: []string{"LogTestOne"},
				ExcludedCategories: []string{"LogTestTwo"},
			},
			expectedFilteredLog: []byte(`{1}[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
{2}[2024.11.14-16.08.23.919][Warning][LogTestOne]: This is a test log line
`),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			counter := 0
			result := filterLinesForChunk(input, tt.filters, &counter)

			require.Equal(t, tt.expectedFilteredLog, result)
		})
	}
}
