package config

import (
	"fmt"
	"strings"

	"github.com/jimoe/repocli/arguments"
)

type Config struct {
	*CliConfig
	*YamlConfig
}

func Load() (*Config, error) {
	cli := getCliConfig()

	yaml, err := loadYaml()
	if err != nil {
		return &Config{}, fmt.Errorf("yaml: %w", err)
	}

	cfg := &Config{cli, yaml}
	return cfg, nil
}

func (cfg *Config) GetRepo(alias *arguments.Alias) (bool, *Repo) {
	for _, r := range cfg.Repoes {
		if r.Name == alias.String() {
			return true, r
		}
		if strings.ReplaceAll(r.Name, "-", "") == alias.String() {
			return true, r
		}
		for _, a := range r.Aliases {
			if a == alias.String() {
				return true, r
			}
		}
	}

	return false, &Repo{}
}
