package tasks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jimoe/repocli/arguments"
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

func ConfigInit(cfg *config.Config, path *arguments.Path) error {
	if _, err := os.Stat(path.String()); os.IsNotExist(err) {
		return errors.New("given <path> does not exist")
	}

	filename := fmt.Sprintf("%s/%s.yml", path, cfg.CliName)
	if err := ioutil.WriteFile(filename, []byte(configExample), 0644); err != nil {
		return fmt.Errorf("could not write example config to file: %w", err)
	}

	fmt.Printf("Config file is saved at '%s'. Edit it to suit your repoes.\n", path)
	fmt.Printf("You may always see the the example config later running '%s config example'\n", cfg.CliName)

	return nil
}
