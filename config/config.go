package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/jimoe/repocli/arguments"
)

type Config struct {
	CliName string
	Version string
	Yaml    struct {
		Path            string
		Filename        string
		PathAndFilename string
	}
	*YamlConfig
}

func Load() (cfg *Config, err error) {
	cfg = &Config{}
	cfg.CliName = "repocli"
	cfg.Version = "v1.0.0"

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return cfg, fmt.Errorf("could not get os user-specific config directory. Message: %w", err)
	}

	cfg.Yaml.Filename = "config.yml"
	cfg.Yaml.Path = fmt.Sprintf("%s/%s", userConfigDir, cfg.CliName)
	cfg.Yaml.PathAndFilename = fmt.Sprintf("%s/%s", cfg.Yaml.Path, cfg.Yaml.Filename)

	yaml, err := loadYaml(cfg.Yaml.PathAndFilename)
	if err != nil {
		cfg.YamlConfig = &YamlConfig{} // to avoid nil pointer error without testing all places it is used
		return cfg, fmt.Errorf("yaml: %w", err)
	}

	cfg.YamlConfig = yaml

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
