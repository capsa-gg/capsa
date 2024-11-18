package logchunk

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		lineMetadata logChunkLineMetadata
		filters      LogStreamLineFilters
		expected     bool
	}{
		"no filters": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{},
			expected:     true,
		},
		"included severity match": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{IncludedSeverities: []string{"Log"}},
			expected:     true,
		},
		"included severity no match": {
			lineMetadata: logChunkLineMetadata{Severity: "Error", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{IncludedSeverities: []string{"Log", "Verbose"}},
			expected:     false,
		},
		"included category match": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{IncludedCategories: []string{"LogTestExample"}},
			expected:     true,
		},
		"included category no match": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogWorldMetrics"},
			filters:      LogStreamLineFilters{IncludedCategories: []string{"LogTestExample"}},
			expected:     false,
		},
		"excluded category match": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters:      LogStreamLineFilters{ExcludedCategories: []string{"LogTestExample"}},
			expected:     false,
		},
		"excluded category match overwritten by included category": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogTestExample"},
			filters: LogStreamLineFilters{
				IncludedCategories: []string{"LogCapsaCore"},
				ExcludedCategories: []string{"LogTestExample"},
			},
			expected: false,
		},
		"excluded category no match": {
			lineMetadata: logChunkLineMetadata{Severity: "Log", Category: "LogWorldMetrics"},
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
		includedLines       []int
	}{
		"no filters": {
			filters:             LogStreamLineFilters{},
			expectedFilteredLog: input,
			includedLines:       []int{1, 2, 3, 4, 5, 6},
		},
		"severity only": {
			filters: LogStreamLineFilters{IncludedSeverities: []string{"Error", "Log", "Verbose"}},
			expectedFilteredLog: []byte(`[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.919][Log][LogTestOne]: This is a test log line
[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
`),
			includedLines: []int{1, 3, 5},
		},
		"included category": {
			filters: LogStreamLineFilters{IncludedCategories: []string{"LogTestTwo"}},
			expectedFilteredLog: []byte(`[2024.11.14-16.08.23.955][Display][LogTestTwo]: This is a test log line
[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
[2024.11.14-16.08.25.493][VeryVerbose][LogTestTwo]: This is a test log line
`),
			includedLines: []int{4, 5, 6},
		},
		"excluded category": {
			filters: LogStreamLineFilters{ExcludedCategories: []string{"LogTestTwo"}},
			expectedFilteredLog: []byte(`[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.919][Warning][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.919][Log][LogTestOne]: This is a test log line
`),
			includedLines: []int{1, 2, 3},
		},
		"included category overriding excluded category": {
			filters: LogStreamLineFilters{
				IncludedCategories: []string{"LogTestTwo"},
				ExcludedCategories: []string{"LogTestTwo"},
			},
			expectedFilteredLog: []byte(`[2024.11.14-16.08.23.955][Display][LogTestTwo]: This is a test log line
[2024.11.14-16.08.25.493][Verbose][LogTestTwo]: This is a test log line
[2024.11.14-16.08.25.493][VeryVerbose][LogTestTwo]: This is a test log line
`),
			includedLines: []int{4, 5, 6},
		},
		"mixed": {
			filters: LogStreamLineFilters{
				IncludedSeverities: []string{"Error", "Warning", "Verbose"},
				IncludedCategories: []string{"LogTestOne"},
				ExcludedCategories: []string{"LogTestTwo"},
			},
			expectedFilteredLog: []byte(`[2024.11.14-16.08.23.916][Error][LogTestOne]: This is a test log line
[2024.11.14-16.08.23.919][Warning][LogTestOne]: This is a test log line
`),
			includedLines: []int{1, 2},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			counter := 0
			includedLines := []int{}
			result := filterLinesForChunk(input, tt.filters, &counter, &includedLines)

			require.Equal(t, tt.expectedFilteredLog, result)
			require.Equal(t, tt.includedLines, includedLines)
		})
	}
}
