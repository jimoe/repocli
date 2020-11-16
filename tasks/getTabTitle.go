package tasks

import (
	"fmt"

	"github.com/jimoe/editor-and-change-dir/arguments"
	"github.com/jimoe/editor-and-change-dir/config"
)

func GetTabTitle(cfg *config.Config, path *arguments.Path) {
	for _, r := range cfg.Repoes {
		if r.Path == path.String() && r.Terminal != nil {
			fmt.Println(r.Terminal.Title)
			return
		}
		for _, p := range r.MonoRepo {
			packagePath := fmt.Sprintf("%s/%s", r.Path, p.SubPath)
			if packagePath == path.String() && p.Terminal != nil {
				fmt.Println(p.Terminal.Title)
				return
			}
		}
	}
}
