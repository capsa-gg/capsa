package entities

import (
	"time"

	"github.com/google/uuid"
)

// LogCreatedResult contains the result of a created log.
type LogCreatedResult struct {
	UUID        uuid.UUID
	ClientJWT   string
	TokenExpiry time.Time
}
