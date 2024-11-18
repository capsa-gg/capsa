package main

import (
	"fmt"
	"time"
)

const tsFormat = "2006.01.02-15.04.05.000"

// Running this package will generate a log file with 1M lines for testing
// Run it with go run . > chunk_long.log
func main() {
	ts := time.Date(2024, 11, 15, 22, 00, 00, 0, time.UTC)
	incrementMs := 42
	lines := 1_000_000

	for i := 0; i < lines; i++ {
		line := genLine(ts)
		fmt.Println(line)

		ts = ts.Add(time.Duration(incrementMs * 1_000_000)) // 1ms=1_000_000ns
	}
}

func genLine(ts time.Time) string {
	return fmt.Sprintf("[" + ts.Format(tsFormat) + "][Log][LogCategoryExample]: This is an entry in a long log file") // No Sprintf for brevity
}
