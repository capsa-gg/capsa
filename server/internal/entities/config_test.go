package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

func TestDatabaseConnectionString(t *testing.T) {
	tests := map[string]struct {
		config   entities.Config
		expected string
	}{
		"Default SSL connection": {
			config:   entities.Config{DatabaseHost: "localhost", DatabasePort: 5432, DatabaseName: "testdb", DatabaseUser: "testuser", DatabasePass: "password"},
			expected: "host=localhost port=5432 user=testuser password=password dbname=testdb sslmode=disable",
		},
		"SSL connection disabled": {
			config:   entities.Config{DatabaseHost: "localhost", DatabasePort: 5432, DatabaseName: "testdb", DatabaseUser: "testuser", DatabasePass: "password", DatabaseSSL: true},
			expected: "host=localhost port=5432 user=testuser password=password dbname=testdb",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := test.config.DatabaseConnectionString()
			assert.Equal(t, test.expected, actual)
		})
	}
}
