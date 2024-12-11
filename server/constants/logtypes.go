package constants

import "fmt"

// LogType is used to determine a log type.
type LogType string

const (
	// LogTypeClient indicates a client (no server code in build) as defined in log_client_type.
	LogTypeClient = "Client"

	// LogTypeGame indicates a game (client with server code in build) as defined in log_client_type.
	LogTypeGame = "Game"

	// LogTypeEditor indicates an editor build as defined in log_client_type.
	LogTypeEditor = "Editor"

	// LogTypeServer indicates a dedicated server as defined in log_client_type.
	LogTypeServer = "Server"
)

// LogTypeFromString parses a string and validates if it's a valid LogType.
func LogTypeFromString(s string) (LogType, error) {
	if s == LogTypeClient {
		return LogTypeClient, nil
	}

	if s == LogTypeGame {
		return LogTypeGame, nil
	}

	if s == LogTypeEditor {
		return LogTypeEditor, nil
	}

	if s == LogTypeServer {
		return LogTypeServer, nil
	}

	return "", fmt.Errorf("%s is not a valid LogType value", s)
}
