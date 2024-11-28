package domainerror

// ErrorType is used to differentiate between different types of domain errors.
type ErrorType string

const (
	// NoModifications indicates that no modifications are made for a resource.
	// Results in a 304 response.
	NoModifications = ErrorType("error-no-modifications")

	// InvalidArgument indicates that validation on the arguments passed in has failed.
	// Results in a 400 response.
	InvalidArgument = ErrorType("error-invalid-argument")

	// NoPermission indicates that requester does not have permission to access the requested data .
	// Results in a 401 response.
	NoPermission = ErrorType("error-invalid-no-permission")

	// NotFound indicates that a requested resource was not found.
	// Results in a 404 response.
	NotFound = ErrorType("error-not-found")

	// Conflict indicates that a data conflict occurred.
	// Results in a 409 response.
	Conflict = ErrorType("error-conflict")

	// Unexpected indicates that an unexpected error occurred.
	// Results in a 500 response.
	Unexpected = ErrorType("error-unexpected")
)
