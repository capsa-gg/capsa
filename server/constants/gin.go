package constants

const (
	// APIStaticPath is the path used after the /v1 prefix to serve static files.
	APIStaticPath = "/static"

	// GinContextKeyValidatedClient is the Gin context key to store the validated jwt claims for clients.
	GinContextKeyValidatedClient = "validated-client"

	// GinContextKeyValidatedUser is the Gin context key to store the validated jwt claims for users.
	GinContextKeyValidatedUser = "validated-user"
)
