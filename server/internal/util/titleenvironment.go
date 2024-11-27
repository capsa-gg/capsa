package util

import "regexp"

var titleEnvRegex = regexp.MustCompile("^[a-zA-Z0-9]+$")

// IsValidTitleOrEnvironmentName validates a string to see if it's a valid title or environment title.
func IsValidTitleOrEnvironmentName(name string) bool {
	return titleEnvRegex.MatchString(name)
}
