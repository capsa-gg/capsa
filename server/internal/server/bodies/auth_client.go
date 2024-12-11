package bodies

import (
	"time"

	"github.com/google/uuid"
)

// ClientLogCreationRequest contains the data to request the creation of a new log session.
type ClientLogCreationRequest struct {
	Key      uuid.UUID `json:"key" validate:"required"`
	Platform string    `json:"platform" validate:"required,max=64"`
	Type     string    `json:"type" validate:"required,max=32"` // Needs manual validation for enum
}

// ClientLogCreationResponse contains the after a client has successfully authenticated and a log session has been created.
type ClientLogCreationResponse struct {
	Token   string    `json:"token"`
	LogID   uuid.UUID `json:"logId"`
	LinkWeb string    `json:"linkWeb"`
	Expiry  time.Time `json:"expiry"`
}
