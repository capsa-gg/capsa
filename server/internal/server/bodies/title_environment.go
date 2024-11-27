package bodies

import "github.com/google/uuid"

// TitleEnvironment contains the data about an environment in a given title.
type TitleEnvironment struct {
	Title           string    `json:"title"`
	EnvironmentName string    `json:"environmentName"`
	EnvironmentKey  uuid.UUID `json:"environmentKey"`
}
