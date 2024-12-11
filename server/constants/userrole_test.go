package constants_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/capsa-gg/capsa/server/constants"
)

func TestUserRoleFromString(t *testing.T) {
	tests := map[string]struct {
		input     string
		wantValue constants.UserRole
		wantErr   bool
	}{
		"Valid Admin": {
			input:     "Admin",
			wantValue: constants.UserRoleAdmin,
			wantErr:   false,
		},
		"Valid User": {
			input:     "User",
			wantValue: constants.UserRoleUser,
			wantErr:   false,
		},
		"Admin invalid casing": {
			input:     "admin",
			wantValue: "",
			wantErr:   true,
		},
		"User invalid casing": {
			input:     "user",
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
			got, err := constants.UserRoleFromString(tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.wantValue, got)
		})
	}
}
