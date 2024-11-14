package constants_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/capsa-gg/capsa/server/constants"
)

func TestLogTypeFromString(t *testing.T) {
	tests := map[string]struct {
		input     string
		wantValue constants.LogType
		wantErr   bool
	}{
		"Valid Client": {
			input:     "Client",
			wantValue: constants.LogTypeClient,
			wantErr:   false,
		},
		"Valid Server": {
			input:     "Server",
			wantValue: constants.LogTypeServer,
			wantErr:   false,
		},
		"Client invalid casing": {
			input:     "client",
			wantValue: "",
			wantErr:   true,
		},
		"Server invalid casing": {
			input:     "server",
			wantValue: "",
			wantErr:   true,
		},
		"Invalid Input": {
			input:     "Unknown",
			wantValue: "",
			wantErr:   true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := constants.LogTypeFromString(tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, got, tt.wantValue)
		})
	}
}
