package bodies

// ErrorResponse contains the data for reporting errors to the client.
type ErrorResponse struct {
	// Error contains the error string.
	Error string `json:"error"`

	// Details contains details about the error.
	// This field is mostly for development, but is used in production for adding request body validation errors.
	Details string `json:"details,omitempty"`

	// RawError contains the raw error.
	// This field is only sent in development mode.
	RawError string `json:"raw_error,omitempty"`
}
