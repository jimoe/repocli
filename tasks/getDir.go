package tasks

import (
	"fmt"

	. "github.com/jimoe/editor-and-change-dir/config"
)

func GetDir(alias string) {
	for _, repo := range Cfg.Repoes {
		if repo.Name == alias || exsists(alias, repo.Aliases) {
			// send the path to bash so it can cd to it
			fmt.Println(repo.Path)
			return
		}
	}

	fmt.Printf("  -- '%s' is not in config", alias)
}

func exsists(a string, slice []string) bool {
	for _, s := range slice {
		if s == a {
			return true
		}
	}
	return false
}
