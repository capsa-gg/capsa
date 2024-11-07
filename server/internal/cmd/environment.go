package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lucianonooijen/capsa/server/internal/domain/admin"
)

var (
	environmentAddName  string
	environmentAddTitle string
)

var envCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env"},
	Short:   "Command group for managing environments for a given title",
}

var envAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new environment to a title",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := s.Config.RootLogger.Named("env").Named("add").Sugar()

		if environmentAddName == "" {
			log.Fatalf("name argument is required")
		}

		if environmentAddTitle == "" {
			log.Fatalf("title argument is required")
		}

		err := admin.AddNewEnvironment(s, environmentAddTitle, environmentAddName)
		if err != nil {
			log.Fatalf("error adding new environment to title: %s", err)
		}

		log.Info("environment successfully added")
	},
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all titles and environments with the associated keys",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := s.Config.RootLogger.Named("env").Named("list").Sugar()

		res, err := admin.ListAllTitlesAndEnvironments(s)
		if err != nil {
			log.Fatalf("error listing environment and titles: %s", err)
		}

		log.Infof("")
		log.Infof("| %-20s | %-20s | %-36s |", "Title", "Environment", "Unreal Engine .ini Key")
		log.Infof("| -------------------- | -------------------- | ------------------------------------ |")

		for _, row := range res {
			log.Infof("| %-20s | %-20s | %-36s |", row.Title, row.EnvironmentName, row.EnvironmentKey)
		}

		log.Infof("")

		log.Info("all data printed")
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Command flags
	envAddCmd.Flags().StringVarP(&environmentAddName, "name", "n", "", "New environment name")
	envAddCmd.Flags().StringVarP(&environmentAddTitle, "title", "t", "", "Title name for the new environment")

	// Add sub commands to command
	envCmd.AddCommand(envAddCmd)
	envCmd.AddCommand(envListCmd)

	// Add to root command
	rootCmd.AddCommand(envCmd)
}
