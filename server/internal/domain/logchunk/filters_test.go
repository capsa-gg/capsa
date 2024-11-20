package logchunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var shouldStreamChunkTests = map[string]struct {
	filters       LogStreamLineFilters
	chunkMetadata logChunkMetadata
	expected      bool
}{
	"no filters": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{},
			IncludedCategories: []string{},
			ExcludedCategories: []string{},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Log": 1, "Error": 2},
			CategoriesCount: map[string]int{"LogInclude": 3, "LogExclude": 4},
		},
		expected: true,
	},
	"only included severities with match": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{"Log", "Warning"},
			IncludedCategories: []string{"LogInclude"},
			ExcludedCategories: []string{},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Warning": 2},
			CategoriesCount: map[string]int{"LogInclude": 1, "LogOther": 1},
		},
		expected: true,
	},

	"only included severities without match": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{"Log", "Warning"},
			IncludedCategories: []string{},
			ExcludedCategories: []string{"LogInclude"},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Error": 2},
			CategoriesCount: map[string]int{"LogOther": 2},
		},
		expected: false,
	},
	"only excluded categories no stream": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{},
			IncludedCategories: []string{},
			ExcludedCategories: []string{"LogExclude"},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Log": 1},
			CategoriesCount: map[string]int{"LogExclude": 1},
		},
		expected: false,
	},
	"only excluded categories with stream": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{"Log"},
			IncludedCategories: []string{},
			ExcludedCategories: []string{"LogExclude"},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Log": 1},
			CategoriesCount: map[string]int{"LogInclude": 1, "LogExclude": 2},
		},
		expected: true,
	},
	"both included and excluded categories": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{},
			IncludedCategories: []string{"LogInclude"},
			ExcludedCategories: []string{"LogExclude"},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Log": 1},
			CategoriesCount: map[string]int{"LogInclude": 1, "LogExclude": 2},
		},
		expected: true,
	},
	"mixed no stream": {
		filters: LogStreamLineFilters{
			IncludedSeverities: []string{"Fatal"},
			IncludedCategories: []string{"LogInclude"},
			ExcludedCategories: []string{"LogExclude"},
		},
		chunkMetadata: logChunkMetadata{
			SeveritiesCount: map[string]int{"Log": 1},
			CategoriesCount: map[string]int{"LogInclude": 1, "LogExclude": 2},
		},
		expected: false,
	},
}

func TestLogStreamLineFilters_ShouldStreamChunk(t *testing.T) {
	for name, tt := range shouldStreamChunkTests {
		t.Run(name, func(t *testing.T) {
			result := tt.filters.shouldStreamChunk(tt.chunkMetadata)
			assert.Equal(t, tt.expected, result)
		})
	}
}
