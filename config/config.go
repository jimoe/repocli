package config

import (
	"github.com/jimoe/editor-and-change-dir/aliases"
)

type Config struct {
	CliConfig
	YamlConfig
}

func Load() (Config, error) {
	cli := getCliConfig()

	yaml, err := loadYaml()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{cli, yaml}
	return cfg, nil
}

func (cfg *Config) GetRepo(theAlias aliases.Alias) (bool, Repo) {
	alias := theAlias.String()

	for _, r := range cfg.Repoes {
		if r.Name == alias {
			return true, r
		}
		for _, a := range r.Aliases {
			if a == alias {
				return true, r
			}
		}
	}

	return false, Repo{}
}
