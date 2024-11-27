package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/domain/user"
)

var (
	userAddEmail     string
	userAddFirstName string
	userAddLastName  string
	userAddPassword  string
	userAddRole      string

	userDeactivateEmail string

	userReactivateEmail string
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
		log := getCmdLogger(s.Config, "user").Named("add")

		if userAddEmail == "" {
			log.Fatal("email argument is required")
		}

		if userAddFirstName == "" {
			log.Fatal("first argument is required")
		}

		if userAddLastName == "" {
			log.Fatal("last name argument is required")
		}

		role, err := constants.UserRoleFromString(userAddRole)
		if err != nil {
			log.Fatalf("defined role is not valid: %s", err)
		}

		if !s.Config.IsDevMode && userAddPassword != "" {
			log.Fatalf("adding a user with a password is only enabled development mode")
		}

		if userAddPassword == "" {
			_, err := user.AddNewUser(context.Background(), s, userAddEmail, userAddFirstName, userAddLastName, role)
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

//nolint:dupl // Similar logic, but this is cleaner than an attempt at making it generic, this is fine.
var userDeactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivates a user",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := getCmdLogger(s.Config, "fetchedUser").Named("deactivate")

		if userDeactivateEmail == "" {
			log.Fatal("email argument is required")
		}

		fetchedUser, err := s.Database.GetUserByEmail(context.Background(), userDeactivateEmail)
		if err != nil {
			log.Fatalf("cannot retrieve fetchedUser with email %s: %s", userDeactivateEmail, err)
		}

		err = user.DeactivateUser(context.Background(), s, fetchedUser.UserUuid)
		if err != nil {
			log.Fatalf("error marking user as deactivated: %s", err)
		}

		log.Info("user successfully deactivated")
	},
}

//nolint:dupl // Similar logic, but this is cleaner than an attempt at making it generic, this is fine.
var userReactivateCmd = &cobra.Command{
	Use:   "reactivate",
	Short: "Reactivates a deactivated user",
	Run: func(_ *cobra.Command, _ []string) {
		s := getAndValidateServicesInteractor()
		log := getCmdLogger(s.Config, "fetchedUser").Named("reactivate")

		if userReactivateEmail == "" {
			log.Fatal("email argument is required")
		}

		fetchedUser, err := s.Database.GetUserByEmail(context.Background(), userReactivateEmail)
		if err != nil {
			log.Fatalf("cannot retrieve fetchedUser with email %s: %s", userReactivateEmail, err)
		}

		err = user.ReactivateUser(context.Background(), s, fetchedUser.UserUuid)
		if err != nil {
			log.Fatalf("error reactivating user: %s", err)
		}

		log.Info("user successfully reactivated")
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	// Command flags
	userAddCmd.Flags().StringVarP(&userAddEmail, "email", "e", "", "New user email")
	userAddCmd.Flags().StringVarP(&userAddFirstName, "firstname", "f", "", "New user first name")
	userAddCmd.Flags().StringVarP(&userAddLastName, "lastname", "l", "", "New user last name")
	userAddCmd.Flags().StringVarP(&userAddPassword, "password", "p", "", "New user last name")
	userAddCmd.Flags().StringVarP(&userAddRole, "role", "r", "User", "New user role, 'Admin' or 'User', defaults to 'User'")

	userDeactivateCmd.Flags().StringVarP(&userDeactivateEmail, "email", "e", "", "Email of the user to be deactivated")

	userReactivateCmd.Flags().StringVarP(&userReactivateEmail, "email", "e", "", "Email of the user to be reactivated")

	// Add sub commands to command
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userDeactivateCmd)
	userCmd.AddCommand(userReactivateCmd)

	// Add to root command
	rootCmd.AddCommand(userCmd)
}
