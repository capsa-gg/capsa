package entities

import (
	"time"

	"github.com/google/uuid"

	"github.com/lucianonooijen/capsa/server/constants"
)

// TitleEnvironment contains the data about an environment in a given title.
type TitleEnvironment struct {
	Title           string    `json:"title"`
	EnvironmentName string    `json:"environment_name"`
	EnvironmentKey  uuid.UUID `json:"environment_key"`
}

// LogOverview contains the data for log stored log.
type LogOverview struct {
	LogUUID            uuid.UUID         `json:"id"`
	Platform           string            `json:"platform"`
	LogType            constants.LogType `json:"logType"`
	LineCount          int64             `json:"lineCount"`
	TimestampFirstLine *time.Time        `json:"tsFirstLine"`
	TimestampLastLine  *time.Time        `json:"tsLastLine"`
}

// LogMetadata contains the metadata for a stored log.
type LogMetadata struct {
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
