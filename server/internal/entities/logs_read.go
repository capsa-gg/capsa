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
