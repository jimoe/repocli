package tasks

import (
	"fmt"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

func GetDir(cfg *config.Config, alias *arguments.Alias) error {
	var found bool
	var repo *config.Repo
	if found, repo = cfg.GetRepo(alias); !found {
		return fmt.Errorf("'%s' is not in config", alias)
	}

	// send the path to bash so it can cd to it
	fmt.Println(repo.Path)
	return nil
}
