package constants

import "fmt"

// LogType is used to determine a log type.
type LogType string

const (
	// LogTypeClient indicates a client as defined in log_client_type.
	LogTypeClient = "Client"

	// LogTypeServer indicates a server as defined in log_client_type.
	LogTypeServer = "Server"
)

// LogTypeFromString parses a string and validates if it's a valid LogType.
func LogTypeFromString(s string) (LogType, error) {
	if s == LogTypeClient {
		return LogTypeClient, nil
	}

	if s == LogTypeServer {
		return LogTypeServer, nil
	}

	return "", fmt.Errorf("%s is not a valid LogType value", s)
}
