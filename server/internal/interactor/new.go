package interactor

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/data/emails"
	"github.com/capsa-gg/capsa/server/internal/data/logchunks"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/blobstorage"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/mailer"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/passhash"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
)

// NewServices initializes and validates a new instance of Services.
func NewServices(c *entities.Config) (*Services, error) { //nolint:gocyclo // This is expected in this init logic
	ctx := context.Background()

	// Database connection
	dbConn, err := pgxpool.New(ctx, c.DatabaseConnectionString())
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Ping database
	err = dbConn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	// Database instance
	db := database.New(dbConn)

	// Validate JWK config options
	if (c.JwkPrivateKeyPath != "" && c.JwkPrivateKeyBase64 != "") || (c.JwkPrivateKeyPath == "" && c.JwkPrivateKeyBase64 == "") {
		return nil, errors.New("either jwk private key path, or base64 should be defined, but not both")
	}

	// Load JWK key
	var jwkKeys *rsa.PrivateKey
	if c.JwkPrivateKeyPath != "" { // Loading from disk
		jwkKeys, err = token.LoadPrivateKeyFromPath(c.JwkPrivateKeyPath)
	} else {
		// Remove spaces
		base64StringNoSpaces := strings.ReplaceAll(c.JwkPrivateKeyBase64, " ", "")

		jwkKeys, err = token.LoadPrivateKeyFromBase64String(base64StringNoSpaces)
	}

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

	// Blob storage
	blobStorage, err := blobstorage.New(c)
	if err != nil {
		return nil, fmt.Errorf("error generating blobstorage instance: %w", err)
	}

	// Logblobs
	logChunks := logchunks.New(c, blobStorage)

	s := Services{
		Config:    c,
		DBPool:    dbConn,
		Database:  db,
		Token:     tokenInstance,
		Passhash:  passHash,
		Emails:    email,
		LogChunks: logChunks,
	}

	// Validate config
	validate := validator.New()
	err = validate.Struct(s)

	if err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return &s, nil
}
