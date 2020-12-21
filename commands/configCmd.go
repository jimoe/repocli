package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
)

func createConfigCmd(cfg *config.Config) *cobra.Command {
	const description = "config-file example and initialization"

	return &cobra.Command{
		Use:   "config",
		Short: fmt.Sprintf("%s. Run '%s help config' for details", description, cfg.CliName),
		Long:  description,
		Args:  cobra.ExactArgs(0),
	}
}

func createConfigExampleCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "example",
		Short: `Description and example config file`,
		Args:  cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			tasks.ConfigExample(cfg)
		},
	}
}

func createConfigInitCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: `Make a config file based on the example in you os user specific config dir`,
		Args:  cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			err := tasks.ConfigInit(cfg)
			if err != nil {
				exit(err, nil)
			}
		},
	}
}
