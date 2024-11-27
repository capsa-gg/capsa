package domainerror

// ErrorType is used to differentiate between different types of domain errors.
type ErrorType string

const (
	// InvalidArgument indicates that validation on the arguments passed in has failed.
	InvalidArgument = ErrorType("error-invalid-argument")

	// NoPermission indicates that requester does not have permission to access the requested data .
	NoPermission = ErrorType("error-invalid-no-permission")

	// NotFound indicates that a requested resource was not found.
	NotFound = ErrorType("error-not-found")

	// Conflict indicates that a data conflict occurred.
	Conflict = ErrorType("error-conflict")

	// Unexpected indicates that an unexpected error occurred.
	Unexpected = ErrorType("error-unexpected")
)
