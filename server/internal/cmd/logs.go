package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/capsa-gg/capsa/server/internal/domain/logs"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Command group for managing logs",
}

var logsCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans logs from the database and blob storage that are over the configured retention time. WARNING: deletes logs",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := getCmdLogger(s.Config, "logs").Named("clean")

		err := logs.DeleteLogsOverRetention(context.Background(), s)
		if err != nil {
			log.Fatalf("error cleaning logs: %s", err)
		}

		log.Info("logs successfully cleaned")
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Add sub commands to command
	logsCmd.AddCommand(logsCleanCmd)

	// Add to root command
	rootCmd.AddCommand(logsCmd)
}
