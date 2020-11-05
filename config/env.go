package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	CliName string `envconfig:"CLI_NAME" required:"true"`

	SourceHome string `envconfig:"SOURCE_HOME" required:"true"`
	BinHome    string `envconfig:"BIN_HOME" required:"true"`
}

var pathError *os.PathError

func LoadEnv() (EnvConfig, error) {
	err := godotenv.Load()
	if err != nil && !errors.As(err, &pathError) {
		log.Printf("Error while loading config-file. %s", err)
	}

	var config EnvConfig
	err = envconfig.Process("", &config)
	if err != nil {
		return EnvConfig{}, err
	}

	// fmt.Printf("-------------------------------- env: %#v\n\n", config)

	return config, nil
}
