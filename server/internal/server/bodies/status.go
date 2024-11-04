package bodies

// StatusResponse contains the data for reporting the status of the running server.
type StatusResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"ok"`
	Version string `json:"version" example:"0.0.1"`
}
