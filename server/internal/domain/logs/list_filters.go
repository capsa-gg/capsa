package logs

import (
	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
)

// ListFilters contains the filtering settings for listing logs.
type ListFilters struct {
	Environment *uuid.UUID
	Platform    *string
	LogType     *constants.LogType
}
