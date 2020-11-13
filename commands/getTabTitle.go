package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/aliases"
	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
)

func getTabTitleCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "gettabtitle <alias>",
		Short: `Get the terminal tab title that is associated with the given <alias>`,

		Args: cobra.ExactArgs(1),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			alias := aliases.Alias(args[0])
			if err := alias.Validate(); err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.GetTabTitle(cfg, alias)
		},
	}
}
