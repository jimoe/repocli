package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/arguments"
	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
)

const shortDescription = "Get the terminal tab title that is associated with the given <full path> (pwd)."

func getTabTitleCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gettabtitle <full path>", // TODO: oppdatere denne
		Short: shortDescription,
		Long:  shortDescription + " If path is not found then nothing is returned",

		Args: cobra.RangeArgs(0, 1),

		Run: func(cmd *cobra.Command, args []string) {
			returnList, err := cmd.Flags().GetBool("list")
			if err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			if returnList {
				if len(args) > 0 {
					fmt.Printf("Error: Do not enter path when the list flag is set\n\n")
					_ = cmd.Usage()
					os.Exit(1)
				}
				tasks.GetTabTitleList(cfg)
				return
			}

			if len(args) < 1 {
				fmt.Printf("Error: You need to add the path as a parameter\n\n")
				_ = cmd.Usage()
				os.Exit(1)
			}

			path := arguments.NewPath(args[0])
			if err := path.Validate(); err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.GetTabTitle(cfg, path)
		},
	}

	cmd.Flags().BoolP(
		"list",
		"l",
		false,
		"If you want to return a list of all paths and corresponding tab-titles for use in your shell",
	)

	return cmd
}
