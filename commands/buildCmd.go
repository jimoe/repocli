package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/arguments"
	"github.com/jimoe/editor-and-change-dir/color"
	. "github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
)

var buildServiceArgs = arguments.Arglist{"cli"}

var buildCmd = &cobra.Command{
	Use: fmt.Sprintf(`build cli
  %s build cli`, Cfg.CliName),
	Short: "Rebuild the cli",

	Args:      cobra.RangeArgs(1, 1),
	ValidArgs: buildServiceArgs,

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.OnlyValidArgs(cmd, args); err != nil {
			color.Red.Printf("Error: %s\n\n", err.Error())
			_ = cmd.Usage()
			os.Exit(1)
		}

		service := args[0]
		tasks.Build(service)
	},
}
