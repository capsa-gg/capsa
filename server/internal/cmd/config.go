package cmd

import (
	"fmt"
	"os"

	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/infrastructure/config"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

func initConfig() {
	configFile = "config.yml"

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		fmt.Printf("error loading config: %s\n", err)
		os.Exit(1)
	}

	if err != nil {
		getCmdLogger(cfg, "initConfig").Fatalf("error loading config: %s", err)
	}

	configuration = cfg
}

func getAndValidateConfig() *entities.Config {
	if configuration == nil {
		fmt.Println("config nil")
		os.Exit(1)
	}

	if configuration.RootLogger == nil {
		fmt.Println("rootlogger nil")
		os.Exit(1)
	}

	return configuration
}

func getAndValidateServicesInteractor() *interactor.Services {
	c := getAndValidateConfig()

	s, err := interactor.NewServices(c)
	if err != nil {
		fmt.Printf("error creating services interactor: %s\n", err)
		os.Exit(1)
	}

	return s
}
