package commands

import (
	"fmt"

	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/tasks"
)

func buildCmd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("build"),
		Short: "Rebuild this cli",

		Args: cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			tasks.Build()
		},
	}
}
