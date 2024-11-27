package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/capsa-gg/capsa/server/internal/util"
)

func TestIsValidTitleOrEnvironmentName(t *testing.T) {
	tests := map[string]struct {
		input string
		valid bool
	}{
		"Valid alphanumeric": {
			input: "ExampleTitle",
			valid: true,
		},
		"Invalid with Greek alphabet": {
			input: "ΤεστΤεστΤεστ",
			valid: false,
		},
		"Invalid with space": {
			input: "User 123",
			valid: false,
		},
		"Invalid with special character": {
			input: "user@123",
			valid: false,
		},
		"Valid numeric": {
			input: "123",
			valid: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := util.IsValidTitleOrEnvironmentName(tt.input)
			require.Equal(t, tt.valid, result)
		})
	}
}
