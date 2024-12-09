package bodies

import (
	"time"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
)

// LogOverview is used to show information about available logs.
type LogOverview struct {
	HasMore bool      `json:"hasMore"`
	Logs    []LogInfo `json:"logs"`
}

// LogInfo contains the data for log stored log.
type LogInfo struct {
	LogUUID            uuid.UUID         `json:"id"`
	Platform           string            `json:"platform"`
	LogType            constants.LogType `json:"logType"`
	Title              string            `json:"title"`
	Environment        string            `json:"environment"`
	LineCount          int64             `json:"lineCount"`
	ChunkCount         int64             `json:"chunkCount"`
	LinkedLogCount     int64             `json:"linkedLogCount"`
	TimestampFirstLine *time.Time        `json:"tsFirstLine"`
	TimestampLastLine  *time.Time        `json:"tsLastLine"`
	CategoriesCounts   map[string]int    `json:"categoriesCounts"`
	SeveritiesCounts   map[string]int    `json:"severitiesCounts"`
}

// LogMetadata contains the metadata for a stored log.
type LogMetadata struct {
	LogData            LogInfo                 `json:"logData"`
	AdditionalMetadata []LogAdditionalMetadata `json:"additionalMetadata"`
	Links              []LogLink               `json:"linkedLogs"`
}

// LogAdditionalMetadata contains the arbitrary metadata that can be set for logs.
type LogAdditionalMetadata struct {
	SavedOn  time.Time      `json:"savedOn"`
	Metadata map[string]any `json:"metadata"`
}

// LogLink shows the data for logs linked to the source log.
type LogLink struct {
	LinkedLog   uuid.UUID `json:"linkedLog"`
	Description string    `json:"description"`
}
