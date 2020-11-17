package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/arguments"
	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
)

func editorCmd(cfg *config.Config) *cobra.Command {
	eCmd := &cobra.Command{
		Use:   fmt.Sprintf(`editor <alias>`),
		Short: "Open repo associated with <alias> in the editor defined in the repo-config",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shouldReturnDir, err := cmd.Flags().GetBool("returndir")
			if err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			alias := arguments.NewAlias(args[0])
			if err := alias.Validate(); err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.Editor(cfg, alias, shouldReturnDir)
		},
	}

	eCmd.Flags().BoolP(
		"returndir",
		"d",
		false,
		"If you want to return the homedir of the repo for use in your shell",
	)

	return eCmd
}
