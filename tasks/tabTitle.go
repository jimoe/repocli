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
			packagePath := packagePath(r, p)
			if packagePath == path.String() && p.Terminal != nil {
				fmt.Println(p.Terminal.Title)
				return
			}
		}
	}
}

func GetTabTitleList(cfg *config.Config) {
	for _, r := range cfg.Repoes {
		if r.Terminal != nil {
			fmt.Print(format(r.Path, r.Terminal.Title))
		}
		for _, p := range r.MonoRepo {
			if p.Terminal != nil {
				packagePath := packagePath(r, p)
				fmt.Print(format(packagePath, p.Terminal.Title))
			}
		}
	}
}

func packagePath(r *config.Repo, p *config.MonoRepo) string {
	return fmt.Sprintf("%s/%s", r.Path, p.SubPath)
}

func format(path, title string) string {
	return fmt.Sprintf("%s;%s\n", path, title)
}
