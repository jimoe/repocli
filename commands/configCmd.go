package commands

import (
	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/config"
	"github.com/jimoe/repocli/tasks"
)

func configCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: `Show an example config file`,
		Args:  cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			tasks.ConfigExample(cfg)
		},
	}
}
