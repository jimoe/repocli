package tasks

import (
	"fmt"

	"github.com/jimoe/editor-and-change-dir/aliases"
	"github.com/jimoe/editor-and-change-dir/config"
)

func Editor(cfg config.Config, alias aliases.Alias) {
	var found bool
	var repo config.Repo
	if found, repo = cfg.GetRepo(alias); !found {
		fmt.Printf("  -- '%s' is not in config", alias)
		return
	}

	fmt.Printf("do something fancy with %v\n\n", repo)
}
