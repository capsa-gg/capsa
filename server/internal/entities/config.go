package entities

import (
	"fmt"

	"go.uber.org/zap"
)

// Config is the application configuration.
type Config struct {
	// Application configuration
	IsDevMode      bool   `mapstructure:"development"` // Not marked required as it would fail when `false` is given as value
	ServerPort     int    `mapstructure:"api_port" validate:"required"`
	ServerProtocol string `mapstructure:"api_protocol" validate:"oneof=http https"`
	ServerHostname string `mapstructure:"api_hostname" validate:"required"`
	WebappHostname string `mapstructure:"webapp_hostname" validate:"required"`

	// JWK data
	JwkPrivateKeyPath   string `mapstructure:"jwk_private_key_path"`   // Not marked as required, this is checked elsewhere
	JwkPrivateKeyBase64 string `mapstructure:"jwk_private_key_base64"` // Not marked as required, this is checked elsewhere

	// DBPool configuration
	DatabaseHost string `mapstructure:"db_host" validate:"required"`
	DatabasePort int    `mapstructure:"db_port" validate:"required"`
	DatabaseName string `mapstructure:"db_name" validate:"required"`
	DatabaseUser string `mapstructure:"db_user" validate:"required"`
	DatabasePass string `mapstructure:"db_pass" validate:"required"`
	DatabaseSSL  bool   `mapstructure:"db_ssl"` // Not marked required as it would fail when `false` is given as value

	// Blob storage configuration
	BlobStorageEndpoint string `mapstructure:"blobstorage_endpoint" validate:"required"`
	BlobStorageRegion   string `mapstructure:"blobstorage_region" validate:"required"`
	BlobStorageKey      string `mapstructure:"blobstorage_key" validate:"required"`
	BlobStorageSecret   string `mapstructure:"blobstorage_secret" validate:"required"`
	BlobStorageBucket   string `mapstructure:"blobstorage_bucket" validate:"required"`

	// Email configuration
	EmailSenderName  string `mapstructure:"email_sender_name" validate:"required"`
	EmailSenderEmail string `mapstructure:"email_sender_email" validate:"required"`
	BrevoAPIKey      string `mapstructure:"brevo_api_key" validate:"required"`

	// Application logic configuration
	LogRetentionDays    int `mapstructure:"log_retention_days" validate:"required"`
	LogMaxDurationHours int `mapstructure:"log_max_duration_hours" validate:"required"`

	// Application-wide config
	RootLogger *zap.Logger `validate:"required"`
}

// DatabaseConnectionString returns database connection string for connection to PostgreSQL.
func (c *Config) DatabaseConnectionString() string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.DatabaseHost, c.DatabasePort, c.DatabaseUser, c.DatabasePass, c.DatabaseName)

	if !c.DatabaseSSL {
		connStr += " sslmode=disable"
	}

	return connStr
}
