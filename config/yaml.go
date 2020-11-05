package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Test    string   `yaml:"test"`
	Editors []Editor `yaml:"editors"`
	Repoes  []Repo   `yaml:"repoes"`
}

type Repo struct {
	Name    string   `yaml:"name"`
	Editor  string   `yaml:"editor"`
	Path    string   `yaml:"path"`
	Aliases []string `yaml:"aliases"`
}

type Editor struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func loadYaml() (YamlConfig, error) {
	filename := fmt.Sprintf("./%s.yml", os.Args[0])

	f, err := os.Open(filename)
	if err != nil {
		return YamlConfig{}, fmt.Errorf("could not read yaml file: %w", err)
	}
	defer f.Close()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return YamlConfig{}, fmt.Errorf("coould new decode yaml: %w", err)
	}

	// fmt.Printf("-------------------------------- yaml: %#v\n\n", cfg)

	return cfg, nil
}
