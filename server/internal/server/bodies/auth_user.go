package bodies

import (
	"time"

	"github.com/google/uuid"
)

// UserLoginRequest contains the data to for a user to log in.
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,ascii,min=12,max=64"`
}

// UserPasswordResetCompleteRequest contains the reset token and new password.
type UserPasswordResetCompleteRequest struct {
	ResetToken uuid.UUID `json:"resetToken" validate:"required,uuid"`
	Password   string    `json:"password" validate:"required,ascii,min=12,max=64"`
}

// UserLoginResult contains the result after a successful user login.
type UserLoginResult struct {
	Token       string    `json:"token"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	UserUUID    string    `json:"userUUID"`
	Role        string    `json:"role"`
	TokenExpiry time.Time `json:"tokenExpiry"`
}
