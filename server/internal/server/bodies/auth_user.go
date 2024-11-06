package bodies

// UserLoginRequest contains the data to for a user to log in.
type UserLoginRequest struct {
	Email    string `json:"email" validation:"required"`
	Password string `json:"password" validation:"required"` // Needs manual validation
}
