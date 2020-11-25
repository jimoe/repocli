package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/arguments"
	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
)

const shortDescription = "Get the terminal tab titles for all repoes or the one that is associated with the given full path."

func getTabTitleCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("gettabtitles"),
		Short: shortDescription,
		Long:  shortDescription + " If path is not found then nothing is returned",

		Args: cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			specific, err := cmd.Flags().GetString("specific")
			if err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			if specific == "" {
				tasks.GetTabTitleList(cfg)
				return
			}

			path := arguments.NewPath(specific)
			if err := path.Validate(); err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.GetTabTitle(cfg, path)
		},
	}

	cmd.Flags().StringP(
		"specific",
		"",
		"",
		"get tab-title for one specific full path (without trailing slash)",
	)

	return cmd
}
