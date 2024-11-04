//nolint:gochecknoinits // Cobra needs usage of init functions
package cmd

import (
	"database/sql"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/internal/infrastructure/migrator"
)

var (
	dbMigrateDirection string
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Command group for performing database interactions",
}

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs server migrations",
	Run: func(_ *cobra.Command, _ []string) {
		c := getAndValidateConfig()
		log := c.RootLogger.Named("database").Named("migrations").Sugar()

		db, err := sql.Open("postgres", c.DatabaseConnectionString())
		if err != nil {
			log.Fatalf("error opening database connection: %s", err)
		}

		direction := getMigrationDirection(log)
		log.Infof("migration direction: %s", direction)

		migrate := migrator.New(db, c.DatabaseName)
		err = migrate(direction)

		if err != nil {
			log.Fatalf("error migrating database: %s", err)
		}

		log.Info("database migrated successfully")
	},
}

func getMigrationDirection(log *zap.SugaredLogger) migrator.Direction {
	if dbMigrateDirection == "" {
		log.Fatal("direction argument is required, use '-d up' or '-d down' when running the command")
	}

	if dbMigrateDirection == "up" {
		return migrator.UpAll
	}

	if dbMigrateDirection == "down" {
		return migrator.DownAll
	}

	log.Fatalf("direction '%s' is not a valid value", dbMigrateDirection)

	return "" // Keep the compiler happy, this should never be reached due to fatal log
}

func init() {
	// Command flags
	dbMigrateCmd.Flags().StringVarP(&dbMigrateDirection, "direction", "d", "", "The direction for the database migrations")

	// Add sub commands to command
	dbCmd.AddCommand(dbMigrateCmd)

	// Add to root command
	rootCmd.AddCommand(dbCmd)
}
