package tasks

import (
	"fmt"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

func GetDir(cfg *config.Config, alias *arguments.Alias) error {
	repo, subPath, err := cfg.GetRepo(alias)
	if err != nil {
		return err
	}

	fullPath := repo.Path
	if subPath != "" {
		fullPath = fmt.Sprintf("%s/%s", repo.Path, subPath)
	}

	fmt.Println(fullPath)
	return nil
}
