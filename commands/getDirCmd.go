package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/color"
	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
	"github.com/jimoe/editor-and-change-dir/validate"
)

func getDirCmd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "getdir <alias>",
		Short: `Get the homedir of the repo associated with the given <alias>`,

		Args: cobra.ExactArgs(1),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			alias := args[0]
			if err := validate.Alias(alias); err != nil {
				color.Red.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.GetDir(cfg, alias)
		},
	}
}
