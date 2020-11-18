package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/jimoe/editor-and-change-dir/arguments"
)

type YamlConfig struct {
	Cli     *Cli      `yaml:"cli"`
	Editors []*Editor `yaml:"editors"`
	Repoes  []*Repo   `yaml:"repoes"`
}

type Cli struct {
	SourcePath string `yaml:"sourcePath"`
	BinPath    string `yaml:"binPath"`
}

type Editor struct {
	Name   string `yaml:"name"`
	Params string `yaml:"params"`
}

type Repo struct {
	Name     string      `yaml:"name"`
	Path     string      `yaml:"path"`
	Editor   string      `yaml:"editor"`
	Aliases  []string    `yaml:"aliases"`
	Terminal *Terminal   `yaml:"terminal"`
	MonoRepo []*MonoRepo `yaml:"monorepo"`
}

type MonoRepo struct {
	SubPath  string    `yaml:"subpath"`
	Terminal *Terminal `yaml:"terminal"`
}

type Terminal struct {
	Title string `yaml:"title"`
}

func loadYaml() (*YamlConfig, error) {
	ex, err := os.Executable()
	if err != nil {
		return &YamlConfig{}, fmt.Errorf("could not find executable: %w", err)
	}

	filename := fmt.Sprintf("%s.yml", ex)

	f, err := os.Open(filename)
	if err != nil {
		return &YamlConfig{}, fmt.Errorf("could not read yaml file: %w", err)
	}
	defer f.Close()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return &YamlConfig{}, fmt.Errorf("could new decode yaml: %w", err)
	}

	err = cfg.Validate()
	if err != nil {
		return &YamlConfig{}, fmt.Errorf("could not validate yaml: %w", err)
	}

	return &cfg, nil
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

		if err := arguments.NewPath(r.Path).Validate(); err != nil {
			return fmt.Errorf("path not valid: %w (%v)", err, r)
		}
		for _, m := range r.MonoRepo {
			if err := arguments.NewSubPath(m.SubPath).Validate(); err != nil {
				return fmt.Errorf("supath not valid: %w (%v)", err, m)
			}
		}

		// validate that the given editor exists in the editor-list
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

func (ycfg *YamlConfig) String() string {
	return fmt.Sprintf("cli: %v\n\neditors: %v\n\nrepoes:%v", ycfg.Cli, ycfg.Editors, ycfg.Repoes)
}

func (c *Cli) String() string {
	return fmt.Sprintf("{ Source: %s, Bin: %s }", c.SourcePath, c.BinPath)
}

func (e *Editor) String() string {
	return fmt.Sprintf("\n{ Name: %s, Paramas:%s }", e.Name, e.Params)
}

func (r *Repo) String() string {
	return fmt.Sprintf("\n- Name: %s\n  Editor: %s\n  Path: %s\n  Aliases: %v\n  Terminal: %v\n  MonoRepo: %v",
		r.Name, r.Editor, r.Path, r.Aliases, r.Terminal, r.MonoRepo)
}

func (m *MonoRepo) String() string {
	return fmt.Sprintf("\n  - SubPath: %s\n    Terminal: %+v", m.SubPath, m.Terminal)
}

func (t *Terminal) String() string {
	return fmt.Sprintf("{ Title: %s }", t.Title)
}
