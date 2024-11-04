//nolint:gochecknoinits // Cobra needs usage of init functions
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lucianonooijen/capsa/server/internal/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Command group for running the server interactions",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(_ *cobra.Command, _ []string) {
		c := getAndValidateConfig()
		log := c.RootLogger.Named("server").Sugar()

		if err := server.Start(log, c); err != nil {
			c.RootLogger.Sugar().Fatalf("server errored: %s", err)
		}
	},
}

func init() {
	// Add sub commands to command
	serverCmd.AddCommand(serverStartCmd)

	// Add to root command
	rootCmd.AddCommand(serverCmd)
}
