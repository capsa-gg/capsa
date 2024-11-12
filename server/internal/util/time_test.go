package util_test

import (
	"testing"
	"time"

	"github.com/capsa-gg/capsa/server/internal/util"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	input    any
	expected time.Time
	hasError bool
}

func TestExtractTimeFromAny(t *testing.T) {
	testTime := time.Date(2023, 11, 7, 10, 15, 30, 0, time.UTC)

	// Define test cases
	tests := map[string]testCase{
		"Valid time.Time": {
			input:    time.Date(2023, 11, 7, 10, 15, 30, 0, time.UTC),
			expected: time.Date(2023, 11, 7, 10, 15, 30, 0, time.UTC),
			hasError: false,
		},
		"Valid *time.Time": {
			input:    &testTime,
			expected: time.Date(2023, 11, 7, 10, 15, 30, 0, time.UTC),
			hasError: false,
		},
		"Valid string RFC3339": {
			input:    "2023-11-07T10:15:30Z",
			expected: time.Date(2023, 11, 7, 10, 15, 30, 0, time.UTC),
			hasError: false,
		},
		"Valid int64 (Unix timestamp)": {
			input:    int64(1672531200),
			expected: time.Unix(1672531200, 0),
			hasError: false,
		},
		"Valid float64 (Unix timestamp)": {
			input:    float64(1672531200),
			expected: time.Unix(1672531200, 0),
			hasError: false,
		},
		"Invalid string": {
			input:    "invalid",
			expected: time.Time{},
			hasError: true,
		},
		"Invalid type (int)": {
			input:    123,
			expected: time.Time{},
			hasError: true,
		},
		"Nil input": {
			input:    nil,
			expected: time.Time{},
			hasError: true,
		},
	}

	// Run test cases
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := util.ExtractTimeFromAny(test.input)
			if test.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
