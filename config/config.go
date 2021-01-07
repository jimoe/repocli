package config

import (
	"fmt"
	"os"
)

type Config struct {
	CliName string
	Version string
	Yaml    *Yaml
	*YamlConfig
}

type Yaml struct {
	Filename        string
	Path            string
	PathAndFilename string
}

func Load() (cfg *Config, err error) {
	cfg = &Config{}
	cfg.CliName = "repocli"
	cfg.Version = "v1.0.0"

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return cfg, fmt.Errorf("could not get os user-specific config directory. Message: %w", err)
	}

	cfg.Yaml = &Yaml{}
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

func (cfg *Config) String() string {
	return fmt.Sprintf("cliname: %s\nversion: %s\nyaml: %v\n%v\n",
		cfg.CliName, cfg.Version, cfg.Yaml, cfg.YamlConfig)
}

func (y *Yaml) String() string {
	return fmt.Sprintf("\n  Filename: %s\n  Path: %s\n  PathAndFilename: %+v",
		y.Filename, y.Path, y.PathAndFilename)
}
