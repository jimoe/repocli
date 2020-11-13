package tasks

import (
	"fmt"

	"github.com/jimoe/editor-and-change-dir/config"
)

func GetTabTitle(cfg *config.Config, path string) {
	for _, r := range cfg.Repoes {
		if r.Path == path && r.Terminal != nil {
			fmt.Println(r.Terminal.Title)
			return
		}
		for _, p := range r.MonoRepo {
			packagePath := fmt.Sprintf("%s/%s", r.Path, p.SubPath)
			if packagePath == path && p.Terminal != nil {
				fmt.Println(p.Terminal.Title)
				return
			}
		}
	}
}
