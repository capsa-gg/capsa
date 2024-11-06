package interactor

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/data/emails"
	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/infrastructure/mailer"
	"github.com/lucianonooijen/capsa/server/internal/infrastructure/passhash"
	"github.com/lucianonooijen/capsa/server/internal/infrastructure/token"
)

// NewServices initializes and validates a new instance of Services.
func NewServices(c *entities.Config) (*Services, error) {
	// Database connection
	dbConn, err := sql.Open("postgres", c.DatabaseConnectionString())
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Ping database
	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	// Database instance
	db := database.New(dbConn)

	// Load JWK key
	jwkKeys, err := token.LoadPrivateKeyFromPath(c.JwkPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error loading private key from path: %w", err)
	}

	// Token instance
	tokenInstance, err := token.New(c, jwkKeys)
	if err != nil {
		return nil, fmt.Errorf("error generating jwk instance: %w", err)
	}

	// Passhash instance
	passHash := passhash.New()

	// Mailer instance (used for Emails instance)
	mail := mailer.New(c)
	email := emails.New(c, mail)

	s := Services{
		Config:   c,
		DBConn:   dbConn,
		Database: db,
		Token:    tokenInstance,
		Passhash: passHash,
		Emails:   email,
	}

	// Validate config
	validate := validator.New()
	err = validate.Struct(s)

	if err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return &s, nil
}
