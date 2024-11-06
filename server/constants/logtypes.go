package constants

// LogType is used to determine a log type.
type LogType string

const (
	// LogTypeClient indicates a client as defined in log_client_type.
	LogTypeClient = "Client"

	// LogTypeServer indicates a server as defined in log_client_type.
	LogTypeServer = "Server"
)
