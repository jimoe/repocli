package commands

import (
	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
)

func createGetDirCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "getdir <alias>",
		Short: `Get the root directory of the repo associated with the given <alias>`,
		Args:  cobra.ExactArgs(1),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			alias := arguments.NewAlias(args[0])
			if err := alias.Validate(); err != nil {
				exit(err, nil)
			}

			err := tasks.GetDir(cfg, alias)
			if err != nil {
				exit(err, nil)
			}
		},
	}
}
