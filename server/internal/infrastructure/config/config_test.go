package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectServerHostname(t *testing.T) {
	tests := map[string]struct {
		hostname string
		expected string
	}{
		"HTTP Prefix with Trailing Slash": {
			hostname: "http://example.com/",
			expected: "example.com",
		},
		"HTTPS Prefix with Trailing Slash": {
			hostname: "https://example.com/",
			expected: "example.com",
		},
		"HTTP Prefix without Trailing Slash": {
			hostname: "http://example.com",
			expected: "example.com",
		},
		"HTTPS Prefix without Trailing Slash": {
			hostname: "https://example.com",
			expected: "example.com",
		},
		"No Prefix with Trailing Slash": {
			hostname: "example.com/",
			expected: "example.com",
		},
		"No Prefix without Trailing Slash": {
			hostname: "example.com",
			expected: "example.com",
		},
		"Empty String": {
			hostname: "",
			expected: "",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, correctServerHostname(tt.hostname))
		})
	}
}
