package tasks

import (
	"fmt"
	"github.com/jimoe/repocli/config"
	"sort"
	"strings"
)

func ListNamesAndAliases(cfg *config.Config, search string, shouldInclude bool) {
	var names []string
	for _, r := range cfg.YamlConfig.Repoes {
		names = append(names, r.Name)
		names = append(names, r.Aliases...)
	}

	if search != "" {
		names = filterOnSearchString(names, search, shouldInclude)
	}

	sort.Strings(names)

	var leadTextAddon string
	if search != "" {
		switch shouldInclude {
		case true:
			leadTextAddon = " containing"
		case false:
			leadTextAddon = " starting with"
		}
	}

	fmt.Printf("All names and aliases%s:\n\n%s\n\n", leadTextAddon, strings.Join(names, "\n"))
}

func filterOnSearchString(names []string, search string, shouldInclude bool) (filtered []string) {
	for _, a := range names {
		if (shouldInclude && strings.Contains(a, search)) || (!shouldInclude && strings.HasPrefix(a, search)) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}
