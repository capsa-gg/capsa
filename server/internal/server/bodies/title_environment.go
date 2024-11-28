package bodies

import (
	"time"

	"github.com/google/uuid"
)

// TitleResponse contains information about a title.
type TitleResponse struct {
	Title     string    `json:"title"`
	CreatedOn time.Time `json:"createdOn"`
}

// TitleEnvironmentResponse contains the data about an environment in a given title.
type TitleEnvironmentResponse struct {
	Title           string    `json:"title"`
	EnvironmentName string    `json:"environmentName"`
	EnvironmentKey  uuid.UUID `json:"environmentKey"`
}

// AddTitleRequest contains the information to add a new title.
type AddTitleRequest struct {
	TitleName string `json:"title" validate:"required,ascii,min=3,max=24"`
}

// AddEnvironmentRequest contains the information to add a new environment for a title.
type AddEnvironmentRequest struct {
	TitleName       string `json:"title" validate:"required,ascii,min=3,max=24"`
	EnvironmentName string `json:"environment" validate:"required,ascii,min=3,max=24"`
}
