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
	counter := 0
	includedLines := []int{}

	for i := 0; i <= b.N; i++ {
		filterLinesForChunk(oneThousandLineChunk, testIncludeAllFilter, &counter, &includedLines)
	}
}

func BenchmarkFilterLinesForChunk_IncludeAll_10k(b *testing.B) {
	counter := 0
	includedLines := []int{}

	for i := 0; i <= b.N; i++ {
		filterLinesForChunk(tenThousandsLineChunk, testIncludeAllFilter, &counter, &includedLines)
	}
}

func BenchmarkFilterLinesForChunk_IncludeAll_100k(b *testing.B) {
	counter := 0
	includedLines := []int{}

	for i := 0; i <= b.N; i++ {
		filterLinesForChunk(hundredThousandsLineChunk, testIncludeAllFilter, &counter, &includedLines)
	}
}

var testExcludeAllFilter = LogStreamLineFilters{
	IncludedSeverities: []string{"Fatal", "Error"},
	IncludedCategories: nil,
	ExcludedCategories: []string{"LogCategoryExample"},
}

func BenchmarkFilterLinesForChunk_ExcludeAll_1k(b *testing.B) {
	counter := 0
	includedLines := []int{}

	for i := 0; i <= b.N; i++ {
		filterLinesForChunk(oneThousandLineChunk, testExcludeAllFilter, &counter, &includedLines)
	}
}

func BenchmarkFilterLinesForChunk_ExcludeAll_10k(b *testing.B) {
	counter := 0
	includedLines := []int{}

	for i := 0; i <= b.N; i++ {
		filterLinesForChunk(tenThousandsLineChunk, testExcludeAllFilter, &counter, &includedLines)
	}
}

func BenchmarkFilterLinesForChunk_ExcludeAll_100k(b *testing.B) {
	counter := 0
	includedLines := []int{}

	for i := 0; i <= b.N; i++ {
		filterLinesForChunk(hundredThousandsLineChunk, testExcludeAllFilter, &counter, &includedLines)
	}
}
