package migrator

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5" //nolint:gocritic,stylecheck // Needs import for using pgx.WithInstance
	"github.com/golang-migrate/migrate/v4/source/httpfs"

	// Needs side effect from pgx/v5.
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5" //nolint:gocritic,stylecheck // Needs import for side effect

	"github.com/capsa-gg/capsa/server/migrations"
)

// Direction indicates the migration direction.
type Direction string

// UpAll and DownAll indicate the migration direction.
const (
	UpAll   = Direction("UpAll")
	DownAll = Direction("DownAll")
)

// New returns a migrator closure that will accept UpAll or DownAll to up or down migrate the database.
func New(dbConn *sql.DB, dbName string) func(Direction) error {
	return func(direction Direction) error {
		// Direction check
		if direction != UpAll && direction != DownAll {
			return fmt.Errorf("migration direction should be 'UpAll' or 'DownAll'")
		}

		driver, err := pgx.WithInstance(dbConn, &pgx.Config{})
		if err != nil {
			return fmt.Errorf("invalid target postgres instance, %w", err)
		}
		// Source instance for the migrations embedded in server binary
		sourceInstance, err := httpfs.New(http.FS(migrations.Migrations), ".")
		if err != nil {
			return fmt.Errorf("invalid source instance, %w", err)
		}

		// Create migrator instance
		migrator, err := migrate.NewWithInstance("httpfs", sourceInstance, dbName, driver)
		if err != nil {
			return fmt.Errorf("failed to initialize migrate instance, %w", err)
		}

		// Do the actual migration
		if direction == UpAll {
			return handleMigratorErrors(migrator.Up())
		}

		if direction == DownAll {
			return handleMigratorErrors(migrator.Down())
		}

		return errors.New("you should not see this")
	}
}

func handleMigratorErrors(err error) error {
	if errors.Is(err, migrate.ErrNoChange) { // Do not report error when no database change
		return nil
	}

	return err
}
