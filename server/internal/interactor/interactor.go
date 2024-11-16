package interactor

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/data/emails"
	"github.com/capsa-gg/capsa/server/internal/data/logchunks"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/passhash"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
)

// Services contains all the shared services in the application that can be passed to .
type Services struct {
	// Config is the application configuration.
	Config *entities.Config `validate:"required"`

	// DBPool is the database connection instance using pgx.
	DBPool *pgxpool.Pool `validate:"required"`

	// Database is the instance of the SQLc generated database queries.
	Database *database.Queries `validate:"required"`

	// Token is used to generate and validate JWTs, JWKs.
	Token *token.Token `validate:"required"`

	// Passhash is used to generate and compare password hashes.
	Passhash *passhash.PassHash `validate:"required"`

	// Emails is used to send transactional emails.
	Emails *emails.Emails `validate:"required"`

	// LogChunks is used to manage blobs for log chunks.
	LogChunks *logchunks.LogChunks `validate:"required"`
}

// GetDomainLogger generated a *zap.SugaredLogger instance for the domain and function.
func (s Services) GetDomainLogger(domain, function string) *zap.SugaredLogger {
	return s.Config.RootLogger.Named("Domain").Named(domain).Named(function).Sugar()
}
