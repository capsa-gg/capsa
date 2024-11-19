package logchunk

import (
	"testing"
)

var testIncludeAllFilter = LogStreamLineFilters{
	IncludedSeverities: []string{"Fatal", "Error", "Warning", "Log", "Display", "Verbose", "VeryVerbose"},
	IncludedCategories: []string{"LogCategoryExample"},
	ExcludedCategories: nil,
}

func BenchmarkFilterLinesForChunk_IncludeAll_1k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		counter := 0

		filterLinesForChunk(oneThousandLineChunk, testIncludeAllFilter, &counter)
	}
}

func BenchmarkFilterLinesForChunk_IncludeAll_10k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		counter := 0

		filterLinesForChunk(tenThousandsLineChunk, testIncludeAllFilter, &counter)
	}
}

func BenchmarkFilterLinesForChunk_IncludeAll_100k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		counter := 0

		filterLinesForChunk(hundredThousandsLineChunk, testIncludeAllFilter, &counter)
	}
}

var testExcludeAllFilter = LogStreamLineFilters{
	IncludedSeverities: []string{"Fatal", "Error"},
	IncludedCategories: nil,
	ExcludedCategories: []string{"LogCategoryExample"},
}

func BenchmarkFilterLinesForChunk_ExcludeAll_1k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		counter := 0

		filterLinesForChunk(oneThousandLineChunk, testExcludeAllFilter, &counter)
	}
}

func BenchmarkFilterLinesForChunk_ExcludeAll_10k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		counter := 0

		filterLinesForChunk(tenThousandsLineChunk, testExcludeAllFilter, &counter)
	}
}

func BenchmarkFilterLinesForChunk_ExcludeAll_100k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		counter := 0

		filterLinesForChunk(hundredThousandsLineChunk, testExcludeAllFilter, &counter)
	}
}
