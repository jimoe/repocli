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

	editorCmd, params := getEditorCmd(cfg.Editors, repo)

	fmt.Printf("%s %s", editorCmd, strings.Join(params, " "))
	return nil
}

// We validate the config on startup, so we know there will be an editor to find
func getEditorCmd(editors []*config.Editor, repo *config.Repo) (cmd string, params []string) {
	for _, e := range editors {
		if e.Name == repo.Editor {
			// ignore path
			paramStr := strings.ReplaceAll(e.Params, "<path>", "")

			cmd = e.Cmd
			// handle backward compatibility from before 'Cmd' was added
			if cmd == "" {
				cmd = e.Name
			}

			return cmd, strings.Split(paramStr, " ")
		}
	}

	return
}
