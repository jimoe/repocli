package commands

import (
	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
)

func configCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: `Description and an example config file. Can also initialize a config-file`,
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			initPath, err := cmd.Flags().GetString("init")
			if err != nil {
				exit(err, cmd)
			}

			if initPath == "" {
				tasks.ConfigExample(cfg)
				return
			}

			path := arguments.NewPath(initPath)
			if err := path.Validate(); err != nil {
				exit(err, cmd)
			}

			err = tasks.ConfigInit(cfg, path)
			if err != nil {
				exit(err, nil)
			}
		},
	}

	cmd.Flags().StringP(
		"init",
		"",
		"",
		"Make a config file based on the example",
	)

	return cmd
}
