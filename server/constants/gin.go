package constants

const (
	// GinContextKeyValidatedClient is the Gin context key to store the validated jwt claims for clients.
	GinContextKeyValidatedClient = "validated-client"

	// GinContextKeyValidatedUser is the Gin context key to store the validated jwt claims for users.
	GinContextKeyValidatedUser = "validated-user"
)
