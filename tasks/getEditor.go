package tasks

import (
	"fmt"
	"strings"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

func GetEditor(cfg *config.Config, alias *arguments.Alias) error {
	repo, err := cfg.GetRepo(alias)
	if err != nil {
		return err
	}

	editor, params := getEditorCmd(cfg.Editors, repo)

	// send the editor execution command (sans path) to bash
	fmt.Printf("%s %s", editor, strings.Join(params, " "))
	return nil
}

// We validate the config on startup, so we know there will be an editor to find
func getEditorCmd(editors []*config.Editor, repo *config.Repo) (editorName string, params []string) {
	for _, e := range editors {
		if e.Name == repo.Editor {
			// ignore path
			paramStr := strings.ReplaceAll(e.Params, "<path>", "")
			return e.Name, strings.Split(paramStr, " ")
		}
	}

	return
}
