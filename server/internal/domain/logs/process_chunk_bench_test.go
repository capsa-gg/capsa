package logs

import (
	"testing"

	"go.uber.org/zap"
)

var noopBenchmarkLogger = zap.NewNop().Sugar()

func BenchmarkExtractMetadataFromChunk_1k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		extractMetadataFromChunk(noopBenchmarkLogger, oneThousandLineChunk)
	}
}

func BenchmarkExtractMetadataFromChunk_10k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		extractMetadataFromChunk(noopBenchmarkLogger, tenThousandsLineChunk)
	}
}

func BenchmarkExtractMetadataFromChunk_100k(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		extractMetadataFromChunk(noopBenchmarkLogger, hundredThousandsLineChunk)
	}
}

// Uncomment to run 1m, see testdata/README.md for instructions on creating the testing chunk

/*
//go:embed testdata/chunk_long.log
var longChunk []byte

var nopLogger = zap.NewNop().Sugar()

func BenchmarkProcessChunk_Long(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		extractMetadataFromChunk(nopLogger, longChunk)
	}
}
*/
