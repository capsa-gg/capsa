package config

import (
	"fmt"
	"reflect"

	loggerCreator "github.com/lucianonooijen/capsa/server/internal/infrastructure/logger"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/lucianonooijen/capsa/server/internal/entities"
)

// LoadConfig reads in config file and ENV variables if set, creates the root Zap logger.
func LoadConfig(configFile string) (*entities.Config, error) {
	v := viper.New()

	var config entities.Config

	// Loop over the Config type and add all mapstructure field values to v.BindEnv
	// so that Viper will read them, this means you can define application config
	// both in the config.yml file or in the environment variables
	cf := reflect.TypeOf(config)
	for i := 0; i < cf.NumField(); i++ {
		field := cf.Field(i)
		tagValue := field.Tag.Get("mapstructure")

		if err := v.BindEnv(tagValue); err != nil {
			return nil, fmt.Errorf("error binding value in config: %w", err)
		}
	}

	// For now, we only support a config.yml file in the current directory or env variables.
	// Later on, using AddConfigPath and SetConfigName we can allow
	// other methods of application configuration and utilize the
	// full power of Viper for creating a true 12-factor application
	v.SetConfigFile(configFile) // TODO: Support not having a config.yml file present
	v.AutomaticEnv()            // read in environment variables that match
	fmt.Println("Using config file:", v.ConfigFileUsed())

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	// Marshal Viper config to struct and return
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	logger, err := loggerCreator.New(config.IsDevMode)
	if err != nil {
		return nil, fmt.Errorf("error creating logger: %w", err)
	}

	// Attach root logger
	config.RootLogger = logger

	// Validate config
	validate := validator.New()
	err = validate.Struct(config)

	if err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return &config, nil
}
