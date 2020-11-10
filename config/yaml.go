package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Editors []*Editor `yaml:"editors"`
	Repoes  []*Repo   `yaml:"repoes"`
}

type Repo struct {
	Name    string   `yaml:"name"`
	Editor  string   `yaml:"editor"`
	Path    string   `yaml:"path"`
	Aliases []string `yaml:"aliases"`
}

type Editor struct {
	Name   string `yaml:"name"`
	Params string `yaml:"params"`
}

func loadYaml() (YamlConfig, error) {
	ex, err := os.Executable()
	if err != nil {
		return YamlConfig{}, fmt.Errorf("could not find executable: %w", err)
	}

	filename := fmt.Sprintf("%s.yml", ex)

	f, err := os.Open(filename)
	if err != nil {
		return YamlConfig{}, fmt.Errorf("could not read yaml file: %w", err)
	}
	defer f.Close()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return YamlConfig{}, fmt.Errorf("could new decode yaml: %w", err)
	}

	err = cfg.Validate()
	if err != nil {
		return YamlConfig{}, fmt.Errorf("could not validate yaml: %w", err)
	}

	return cfg, nil
}

func (ycfg *YamlConfig) Validate() error {
	for _, e := range ycfg.Editors {
		if e.Name == "" {
			return fmt.Errorf("missing name in editor: %v", e)
		}
	}

	for _, r := range ycfg.Repoes {
		if r.Name == "" {
			return fmt.Errorf("missing 'name' in repo: %v", r)
		}
		if r.Editor == "" {
			return fmt.Errorf("missing 'editor' in repo: %v", r)
		}
		if r.Path == "" {
			return fmt.Errorf("missing 'path' in repo: %v", r)
		}

		var found bool
		for _, e := range ycfg.Editors {
			if r.Editor == e.Name {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("editor in repo not found in editor-list: %s", r.Editor)
		}
	}

	return nil
}
