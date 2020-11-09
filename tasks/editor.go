package tasks

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jimoe/editor-and-change-dir/aliases"
	"github.com/jimoe/editor-and-change-dir/color"
	"github.com/jimoe/editor-and-change-dir/config"
)

func Editor(cfg config.Config, alias aliases.Alias) {
	var found bool
	var repo config.Repo
	if found, repo = cfg.GetRepo(alias); !found {
		fmt.Printf("  -- '%s' is not in config", alias)
		return
	}

	paramDir := getParamDir(repo.Editor, repo.Path)
	if paramDir == "" {
		fmt.Printf("  -- Unknown editor: '%s'", repo.Editor)
		return
	}

	cmd := exec.Command(repo.Editor, paramDir)
	cmd.Dir = repo.Path
	cmd.Stdin = os.Stdin
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		color.Red.Printf("Error: Failed to start editor (%s) for '%s': %w\n", repo.Editor, repo.Name, err)
		os.Exit(1)
	}

	// send the path to bash so it can cd to it
	fmt.Println(repo.Path)
}

func getParamDir(editor string, path string) string {
	switch editor {
	case "webstorm":
		return path
	case "goland":
		return path
	case "code":
		return "."
	default:
		return ""
	}
}
