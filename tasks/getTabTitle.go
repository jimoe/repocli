package tasks

import (
	"fmt"

	"github.com/jimoe/editor-and-change-dir/aliases"
	"github.com/jimoe/editor-and-change-dir/config"
)

func GetTabTitle(cfg *config.Config, alias aliases.Alias) {
	var found bool
	var repo *config.Repo
	if found, repo = cfg.GetRepo(alias); !found {
		fmt.Printf("  -- '%s' is not in config", alias)
		return
	}

	// if repo.Terminal

	// send the path to bash so it can cd to it
	fmt.Println(repo)
}
