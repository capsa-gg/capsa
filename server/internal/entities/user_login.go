package entities

import (
	"time"
)

// UserLoginResult contains the result after a successful user login.
type UserLoginResult struct {
	Token       string    `json:"token"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	UserUUID    string    `json:"userUUID"`
	TokenExpiry time.Time `json:"tokenExpiry"`
}
