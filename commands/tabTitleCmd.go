package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
)

func tabTitleCmd(cfg *config.Config) *cobra.Command {
	const description = "Get the terminal tab titles for all repoes or the one that is associated with the given full path."
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("tabtitle"),
		Short: description,
		Long:  description + " If path is not found then nothing is returned",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			pathFlagValue, err := cmd.Flags().GetString("path")
			if err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			if pathFlagValue == "" {
				tasks.GetTabTitleList(cfg)
				return
			}

			path := arguments.NewPath(pathFlagValue)
			if err := path.Validate(); err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.GetTabTitle(cfg, path)
		},
	}

	cmd.Flags().StringP(
		"path",
		"",
		"",
		"get tab-title for one specific full path (without trailing slash)",
	)

	return cmd
}
