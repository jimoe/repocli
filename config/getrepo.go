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

func (cfg *Config) GetRepo(alias *arguments.Alias) (*Repo, error) {
	for _, r := range cfg.Repoes {
		if r.Name == alias.String() {
			return r, nil
		}
		if strings.ReplaceAll(r.Name, "-", "") == alias.String() {
			return r, nil
		}
		for _, a := range r.Aliases {
			if a == alias.String() {
				return r, nil
			}
		}
	}

	return nil, &RepoNotFoundError{Message: fmt.Sprintf("'%s' not in config", alias.String())}
}
