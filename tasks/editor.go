package tasks

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

func Editor(cfg *config.Config, alias *arguments.Alias, shouldReturnDir bool) error {
	var repo *config.Repo
	var err error

	if alias.String() == "." {
		if len(cfg.Editors) == 0 {
			return errors.New("no editors listed in config")
		}
		repo = &config.Repo{
			Name:   "CurrentDirectory",
			Path:   ".",
			Editor: cfg.Editors[0].Name, // assume the first one is the preferred one
		}
	} else {
		repo, err = cfg.GetRepo(alias)
		if err != nil {
			return err
		}
	}

	editor, params := getEditor(cfg.Editors, repo)

	cmd := exec.Command(editor, params...)
	cmd.Dir = repo.Path
	cmd.Stdin = os.Stdin
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start editor (%s) for '%s': %w\n", repo.Editor, repo.Name, err)
	}

	if shouldReturnDir {
		fmt.Println(repo.Path)
	}
	return nil
}

// We validate the config on startup, so we know there will be an editor to find
func getEditor(editors []*config.Editor, repo *config.Repo) (editorName string, params []string) {
	for _, e := range editors {
		if e.Name == repo.Editor {
			paramStr := strings.ReplaceAll(e.Params, "<path>", repo.Path)
			return e.Name, strings.Split(paramStr, " ")
		}
	}

	return
}
