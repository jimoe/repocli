package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
)

func createEditorCmd(cfg *config.Config) *cobra.Command {
	const description = "Open repo associated with <alias>, or '.' for current directory, in the editor defined in the config."
	cmd := &cobra.Command{
		Use:     fmt.Sprintf(`editor <alias>`),
		Aliases: []string{"e"},
		Short:   description,
		Long:    description + "\nThe name of the repo (exact or with all hyphens removed) can also be used.",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			shouldReturnDir, err := cmd.Flags().GetBool("returndir")
			if err != nil {
				exit(err, cmd)
			}

			alias := arguments.NewAlias(args[0])
			if err := alias.Validate(); err != nil {
				exit(err, nil)
			}

			err = tasks.Editor(cfg, alias, shouldReturnDir)
			if err != nil {
				exit(err, nil)
			}
		},
	}

	cmd.Flags().BoolP(
		"returndir",
		"d",
		false,
		"If you want to return the homedir of the repo for use in your shell",
	)

	return cmd
}
