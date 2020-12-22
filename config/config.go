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

type RepoNotFoundError struct {
	Message string
}

func (e *RepoNotFoundError) Error() string {
	return e.Message
}

func (cfg *Config) GetRepo(alias *arguments.Alias) (*Repo, error) {
	for _, r := range cfg.Repoes {
		if r.Name == alias.String() {
			return r, nil
		}
		if strings.ReplaceAll(r.Name, "-", "") == alias.String() {
			return r, nil
		}
		for _, a := range r.Aliases {
			if a == alias.String() {
				return r, nil
			}
		}
	}

	return &Repo{}, &RepoNotFoundError{Message: fmt.Sprintf("'%s' not in config", alias.String())}
}
