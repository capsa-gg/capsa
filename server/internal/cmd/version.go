package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/capsa-gg/capsa/server/constants"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of Capsa",
	Run: func(_ *cobra.Command, _ []string) {
		log.Printf("Capsa version: %s", constants.Version)
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Add to root command
	rootCmd.AddCommand(versionCmd)
}
