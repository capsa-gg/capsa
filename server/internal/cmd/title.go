package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lucianonooijen/capsa/server/internal/domain/admin"
)

var (
	titleAddName string
)

var titleCmd = &cobra.Command{
	Use:   "title",
	Short: "Command group for managing titles",
}

var titleAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new title to the database",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := s.Config.RootLogger.Named("title").Named("add").Sugar()

		if titleAddName == "" {
			log.Fatalf("name argument is required")
		}

		err := admin.AddNewTitle(s, titleAddName)
		if err != nil {
			log.Fatalf("error adding new title: %s", err)
		}

		log.Info("title successfully added")
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Command flags
	titleAddCmd.Flags().StringVarP(&titleAddName, "name", "n", "", "New title name")

	// Add sub commands to command
	titleCmd.AddCommand(titleAddCmd)

	// Add to root command
	rootCmd.AddCommand(titleCmd)
}
