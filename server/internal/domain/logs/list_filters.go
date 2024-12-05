package logs

import "github.com/google/uuid"

// ListFilters contains the filtering settings for listing logs.
type ListFilters struct {
	Environment *uuid.UUID
}
