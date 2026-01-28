package tasks

import (
	"fmt"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

func GetDir(cfg *config.Config, alias *arguments.Alias) error {
	repo, err := cfg.GetRepo(alias)
	if err != nil {
		return err
	}

	fmt.Println(repo.Path)
	return nil
}
