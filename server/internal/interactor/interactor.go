package interactor

import (
	"database/sql"

	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/entities"
)

// Services contains all the shared services in the application that can be passed to .
type Services struct {
	// Config is the application configuration.
	Config *entities.Config `validate:"required"`

	// DBConn is the database connection instance.
	DBConn *sql.DB `validate:"required"`

	// Database is the instance of the SQLc generated database queries.
	Database *database.Queries `validate:"required"`
}

// GetDomainLogger generated a *zap.SugaredLogger instance for the domain and function.
func (s Services) GetDomainLogger(domain, function string) *zap.SugaredLogger {
	return s.Config.RootLogger.Named("domain").Named(domain).Named(function).Sugar()
}
