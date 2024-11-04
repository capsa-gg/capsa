package cmd

import (
	"github.com/lucianonooijen/capsa/server/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Command group for running the server interactions",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
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
