package cmd

import (
	"github.com/spf13/cobra"

	"github.com/capsa-gg/capsa/server/internal/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Command group for running the server interactions",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := s.Config.RootLogger.Named("server").Sugar()

		log.Info("initializing server instance")

		svr, err := server.New(s)
		if err != nil {
			log.Fatalf("error initializing server instance: %s", err)
		}

		if err := svr.Start(); err != nil {
			log.Fatalf("server errored: %s", err)
		}
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Add sub commands to command
	serverCmd.AddCommand(serverStartCmd)

	// Add to root command
	rootCmd.AddCommand(serverCmd)
}
