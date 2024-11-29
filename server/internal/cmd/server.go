package cmd

import (
	"github.com/spf13/cobra"

	"github.com/capsa-gg/capsa/server/internal/infrastructure/recurringjobs"
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
		log := getCmdLogger(s.Config, "server")

		log.Info("initializing recurring job instance")

		jobrunner, err := recurringjobs.New(s.Config, s)
		if err != nil {
			log.Fatalf("error initializing recurring jobs: %s", err)
		}

		jobrunner.Start()

		log.Info("initializing server instance")

		svr, err := server.New(s)
		if err != nil {
			log.Fatalf("error initializing server instance: %s", err)
		}

		if err := svr.Start(); err != nil {
			log.Fatalf("server errored: %s", err)
		}

		jobrunner.Stop()
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Add sub commands to command
	serverCmd.AddCommand(serverStartCmd)

	// Add to root command
	rootCmd.AddCommand(serverCmd)
}
