package bodies

// ErrorResponse contains the data for reporting errors to the client.
type ErrorResponse struct {
	Error    string `json:"error"`
	RawError string `json:"raw_error,omitempty"`
}
