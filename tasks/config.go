package tasks

import (
	"errors"
	"fmt"
	"os"

	"github.com/jimoe/repocli/config"
)

const configDescription = `Yaml config
Place config in same dir and with the same name as the executable. Just add '.yml' to the filename.

For the 'editors' section:
	The editor will be executed from the repo path (set in the 'repoes' section).
	If 'params' includes the string '<path>' then it will be replaced with the repo path.

EXAMPLE:
`

const configExample = `
editors:
  - name: goland
    params: nosplash <path>
  - name: code
    params: .
repoes:
  - name:    some-repo-name
    path:    /home/username/code/some-repo-name
    editor:  goland
    aliases:
      - some
      - some-repo
    terminal:
      title: SOME
  - name:    another-repo-name
    path:    /home/username/code/another-repo-name
    editor:  code
    aliases:
      - another
    terminal:
      title: ANOTHER
    monorepo:
      - subpath: packages/name
        terminal:
          title: A name
      - subpath: packages/whatever
        terminal:
          title: A whatever
`

func ConfigExample(cfg *config.Config) {
	fmt.Printf("%s	%s\n", configDescription, configExample)
}

func ConfigInit(cfg *config.Config, forceOverwrite bool) error {
	if err := os.MkdirAll(cfg.Yaml.Path, 0775); err != nil {
		return fmt.Errorf("could not make config directory: %w", err)
	}

	if !forceOverwrite {
		if _, err := os.Stat(cfg.Yaml.PathAndFilename); err == nil {
			return errors.New("config file already exists. I refuse to replace it! You may force it by adding '--force'")
		}
	}

	if err := os.WriteFile(cfg.Yaml.PathAndFilename, []byte(configExample), 0644); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	fmt.Printf("Config file is saved in '%s'. Edit it to suit your repoes.\n", cfg.Yaml.Path)
	fmt.Printf("You may always see the the example config later running '%s config example'\n", cfg.CliName)

	return nil
}
