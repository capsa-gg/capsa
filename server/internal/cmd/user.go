package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lucianonooijen/capsa/server/internal/domain/user"
)

var (
	userAddEmail     string
	userAddFirstName string
	userAddLastName  string
	userAddPassword  string
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Command group for managing users",
}

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new user to the database",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := s.Config.RootLogger.Named("user").Named("add").Sugar()

		if userAddEmail == "" {
			log.Fatalf("email argument is required")
		}

		if userAddFirstName == "" {
			log.Fatalf("first argument is required")
		}

		if userAddLastName == "" {
			log.Fatalf("last name argument is required")
		}

		if !s.Config.IsDevMode && userAddPassword != "" {
			log.Fatalf("adding a user with a password is only enabled development mode")
		}

		//nolint:ineffassign // Being explicit here is nice
		var err error = nil

		if userAddPassword == "" {
			err = user.AddNewUser(s, userAddEmail, userAddFirstName, userAddLastName)
			if err != nil {
				log.Fatalf("error adding new user: %s", err)
			}
		} else {
			_, err := user.AddNewUserWithPassword(s, userAddEmail, userAddFirstName, userAddLastName, userAddPassword)
			if err != nil {
				log.Fatalf("error adding new user: %s", err)
			}
		}

		log.Info("user successfully added")
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Command flags
	userAddCmd.Flags().StringVarP(&userAddEmail, "email", "e", "", "New user email")
	userAddCmd.Flags().StringVarP(&userAddFirstName, "firstname", "f", "", "New user first name")
	userAddCmd.Flags().StringVarP(&userAddLastName, "lastname", "l", "", "New user last name")
	userAddCmd.Flags().StringVarP(&userAddPassword, "password", "p", "", "New user last name")

	// Add sub commands to command
	userCmd.AddCommand(userAddCmd)

	// Add to root command
	rootCmd.AddCommand(userCmd)
}
