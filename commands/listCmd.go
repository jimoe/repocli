package commands

import (
	"fmt"
	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
	"github.com/spf13/cobra"
)

// listCmd lists all names and aliases
func listCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [search]",
		Aliases: []string{"aliases", "names", "keys", "key"},
		Short:   "list all names and aliases, or those that matches your search, in your config",
		Example: fmt.Sprintf("%[1]s list start\n%[1]s list -i part", cfg.CliName),
		Args:    cobra.RangeArgs(0, 1),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			shouldInclude, err := cmd.Flags().GetBool("include")
			if err != nil {
				exit(err, cmd)
			}

			var search string
			if len(args) > 0 {
				search = args[0]
			}

			tasks.ListNamesAndAliases(cfg, search, shouldInclude)
		},
	}

	cmd.Flags().BoolP(
		"include",
		"i",
		false,
		"Search for part of name/alias in stead of start of name/alias",
	)

	return cmd
}
