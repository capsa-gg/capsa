package bodies

import (
	"time"

	"github.com/google/uuid"

	"github.com/capsa-gg/capsa/server/constants"
)

// UserInfoResponse contains the information about a user, as visible for admins.
type UserInfoResponse struct {
	UserUUID       uuid.UUID          `json:"userUuid"`
	Email          string             `json:"email"`
	FirstName      string             `json:"firstName"`
	LastName       string             `json:"lastName"`
	HasPasswordSet bool               `json:"hasPasswordSet"`
	Role           constants.UserRole `json:"role"`
	DeactivatedTS  *time.Time         `json:"deactivatedTs"`
	CreatedAt      time.Time          `json:"createdAt"`
}

// UserCreateRequest contains the information to create a user.
type UserCreateRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Role      string `json:"role" validate:"oneof=User Admin"` // Conversion to UserRole happens in handler
}

// UserUpdateRequest contains the information to create a user.
type UserUpdateRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Role      string `json:"role" validate:"oneof=User Admin"` // Conversion to UserRole happens in domain
}
