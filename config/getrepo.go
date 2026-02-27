package config

import (
	"fmt"
	"strings"

	"github.com/jimoe/repocli/arguments"
)

type RepoNotFoundError struct {
	Message string
}

func (e *RepoNotFoundError) Error() string {
	return e.Message
}

func (cfg *Config) GetRepo(alias *arguments.Alias) (*Repo, string, error) {
	for _, r := range cfg.Repoes {
		if r.Name == alias.String() {
			return r, "", nil
		}
		if strings.ReplaceAll(r.Name, "-", "") == alias.String() {
			return r, "", nil
		}
		for _, a := range r.Aliases {
			if a == alias.String() {
				return r, "", nil
			}
		}
		if len(r.MonoRepo) > 0 {
			for _, m := range r.MonoRepo {
				if m.SubPath == alias.String() {
					return r, m.SubPath, nil
				}
				if strings.ReplaceAll(m.SubPath, "-", "") == alias.String() {
					return r, m.SubPath, nil
				}
				for _, a := range m.Aliases {
					if a == alias.String() {
						return r, m.SubPath, nil
					}
				}
			}
		}
	}

	return nil, "", &RepoNotFoundError{Message: fmt.Sprintf("'%s' not in config", alias.String())}
}
