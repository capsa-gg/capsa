package logs

// Uncomment to run benchmarks, see testdata/README.md for instructions on creating the testing chunk

/*
import (
	_ "embed"
	"testing"

	"go.uber.org/zap"
)

//go:embed testdata/chunk_long.log
var longChunk []byte

var nopLogger = zap.NewNop().Sugar()

func BenchmarkProcessChunk_Long(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		extractMetadataFromChunk(nopLogger, longChunk)
	}
}
*/
