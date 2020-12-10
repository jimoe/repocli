package commands

import (
	"fmt"

	"github.com/jimoe/repocli/config"
	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/tasks"
)

func createBuildCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("build"),
		Short: "Rebuild this cli from source",
		Args:  cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			err := tasks.Build(cfg)
			if err != nil {
				exit(err, nil)
			}
		},
	}
}
