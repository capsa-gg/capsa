package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/infrastructure/config"
)

var production bool
var configFile string
var configuration *entities.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "capsa",
	Short: "Capsa's server-side application for running the web server and managed instances of Capsa",
	Long: `Capsa's server application contains all tools necessary to run and manage Capsa instances.

This CLI has multiple sub commands groups. Run with the --help flag to see a full list of commands available.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		if configuration != nil && configuration.RootLogger != nil {
			configuration.RootLogger.Named("cmd.Execute").Sugar().Fatalf("error executing command: %s", err)
		} else {
			log.Panicf("error executing command: %s", err)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// TODO: Support config.prod.yml
	//rootCmd.PersistentFlags().BoolVarP(&production, "prod", "p", false, "use production configuration file (./config.prod.yml)")
	// TODO: use dynamic config loading f.e. https://github.com/spf13/cobra-cli/blob/main/tpl/main.go#L95
	//rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.yml, or ./config.prod.yml for production mode)")
}

func initConfig() {
	//if production {
	//	configFile = "config.prod.yml"
	//} else {
	configFile = "config.yml"
	//}

	cfg, err := config.LoadConfig(configFile)
	logger := cfg.RootLogger
	if err != nil {
		logger.Named("initConfig").Sugar().Fatalf("error loading config: %s", err)
	}

	configuration = cfg
}
