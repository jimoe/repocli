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

	editorCmd, params := getEditor(cfg.Editors, repo, true)

	fmt.Printf("%s %s", editorCmd, strings.Join(params, " "))
	return nil
}
