package entities

import "time"

// LogChunkMetadata contains metadata for a log chunk
// This can contain either the severities or the categories.
type LogChunkMetadata map[string]int

// ChunkStreamer is the type used to stream chunks.
type ChunkStreamer func(chunk string) (int, error)

// MergedLogInput is used to call the log merging function.
// This needs to be a slice to keep the log output ordered, as maps don't keep the same order.
type MergedLogInput []MergedLogInputData

// MergedLogInputData allows the merger to abstractly get log data and contains the string to prefix content the lines with.
type MergedLogInputData struct {
	Key    string
	Loader FilteredLineLoader
}

// LogChunkLineMetadata contains the metadata for a given line in a chunk.
type LogChunkLineMetadata struct {
	Timestamp time.Time
	Category  string
	Severity  string
}

// IsComplete returns whether the LogChunkLineMetadata is complete or not.
func (l LogChunkLineMetadata) IsComplete() bool {
	return l.Severity != "" && l.Category != "" && !l.Timestamp.IsZero()
}
