package tasks

import (
	"fmt"
	"strings"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

func GetEditor(cfg *config.Config, alias *arguments.Alias) error {
	repo, _, err := cfg.GetRepo(alias)
	if err != nil {
		return err
	}

	editorCmd, params, err := getEditor(cfg.Editors, repo)
	if err != nil {
		return fmt.Errorf("error getting editor: %w", err)
	}

	fmt.Printf("%s %s", editorCmd, strings.Join(params, " "))
	return nil
}
