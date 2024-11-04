package entities

import "go.uber.org/zap"

// Config is the application configuration.
type Config struct {
	// Application configuration
	IsDevMode      bool   `mapstructure:"development"` // Not marked required as it would fail when `false` is given as value
	ServerPort     int    `mapstructure:"api_port" validate:"required"`
	ServerHostname string `mapstructure:"api_hostname" validate:"required"`
	WebappHostname string `mapstructure:"webapp_hostname" validate:"required"`
	JWTSecret      string `mapstructure:"jwt_secret" validate:"required"` // TODO: JWK

	// DBConn configuration
	DatabaseHost string `mapstructure:"db_host" validate:"required"`
	DatabasePort int    `mapstructure:"db_port" validate:"required"`
	DatabaseName string `mapstructure:"db_name" validate:"required"`
	DatabaseUser string `mapstructure:"db_user" validate:"required"`
	DatabasePass string `mapstructure:"db_pass" validate:"required"`
	DatabaseSSL  bool   `mapstructure:"db_ssl"` // Not marked required as it would fail when `false` is given as value

	// Email configuration
	EmailSenderName  string `mapstructure:"email_sender_name" validate:"required"`
	EmailSenderEmail string `mapstructure:"email_sender_email" validate:"required"`
	SendinblueAPIKey string `mapstructure:"sendinblue_api_key" validate:"required"`

	// Application logic configuration
	LogRetentionDays    int `mapstructure:"log_retention_days" validate:"required"`
	LogMaxDurationHours int `mapstructure:"log_max_duration_hours" validate:"required"`

	// Application-wide config
	RootLogger *zap.Logger `validate:"required"`
}
