package bodies

import "github.com/google/uuid"

// LogMetadataSaveRequest contains the data to for a user to log in.
type LogMetadataSaveRequest struct {
	LogLinks           map[uuid.UUID]string `json:"log_links"`
	AdditionalMetadata map[string]any       `json:"additional_metadata"`
}
