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
		repo, _, err = cfg.GetRepo(alias)
		if err != nil {
			return err
		}
	}

	editorCmd, params, err := getEditor(cfg.Editors, repo)
	if err != nil {
		return fmt.Errorf("error getting editor: %w", err)
	}

	cmd := exec.Command(editorCmd, params...)
	cmd.Dir = repo.Path
	cmd.Stdin = os.Stdin
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start editorCmd (%s) for '%s': %w\n", repo.Editor, repo.Name, err)
	}

	if shouldReturnDir {
		fmt.Println(repo.Path)
	}
	return nil
}

// We validate the config on startup, so we know there will be an editor to find
func getEditor(editors []*config.Editor, repo *config.Repo) (
	cmd string, params []string, err error,
) {
	for _, e := range editors {
		if e.Name != repo.Editor {
			continue
		}

		cmd = e.Cmd
		// handle backward compatibility from before 'e.Cmd' was added
		if cmd == "" {
			cmd = e.Name
		}

		paramStr := e.Params

		// normal
		paramStr = strings.ReplaceAll(e.Params, "<path>", repo.Path)

		// wsl
		placeholder := "<wslpath>"
		if strings.Contains(paramStr, placeholder) {
			convertedPath, err := convertPathToWsl(repo.Path)
			if err != nil {
				return "", nil, fmt.Errorf("converting path to wsl: %w", err)
			}
			paramStr = strings.ReplaceAll(paramStr, placeholder, convertedPath)
		}

		return cmd, strings.Split(paramStr, " "), nil
	}

	return
}

func convertPathToWsl(path string) (string, error) {
	cmd := exec.Command("wslpath", "-w", path)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to convert path using wslpath: %w", err)
	}
	return string(output), nil
}
