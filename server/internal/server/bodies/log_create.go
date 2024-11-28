package bodies

import (
	"time"

	"github.com/google/uuid"
)

// LogCreatedResult contains the result of a created log.
type LogCreatedResult struct {
	UUID        uuid.UUID `json:"uuid"`
	ClientJWT   string    `json:"client_jwt"`
	TokenExpiry time.Time `json:"token_expiry"`
}
