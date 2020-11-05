package commands

import (
	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/tasks"
)

var getDirCmd = &cobra.Command{
	Use:   "getdir <alias>",
	Short: `Get the homedir of repo with <alias>`,

	Args: cobra.ExactArgs(1),

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		tasks.GetDir(alias)
	},
}
