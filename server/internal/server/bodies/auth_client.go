package bodies

import (
	"time"

	"github.com/google/uuid"
)

// ClientLogCreationRequest contains the data to request the creation of a new log session.
type ClientLogCreationRequest struct {
	Key      uuid.UUID `json:"key" validation:"required"`
	Platform string    `json:"platform" validation:"required"`
	Type     string    `json:"type" validation:"required"` // Needs manual validation
}

// ClientLogCreationResponse contains the after a client has successfully authenticated and a log session has been created.
type ClientLogCreationResponse struct {
	Token   string    `json:"token"`
	LogID   uuid.UUID `json:"log_id"`
	LinkWeb string    `json:"link_web"`
	Expiry  time.Time `json:"expiry"`
}
