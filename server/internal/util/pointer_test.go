package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/capsa-gg/capsa/server/internal/util"
)

func TestPtrTo(t *testing.T) {
	// Define a map of test cases
	testCases := map[string]any{
		"Int":    42,
		"String": "hello",
		"Float":  3.14,
		"Bool":   true,
		"Struct": struct {
			Field        int
			anotherField int
		}{Field: 10, anotherField: 42},
	}

	// Iterate over the test cases
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			got := util.PtrTo(tt)
			require.Equal(t, &tt, got)
		})
	}
}
