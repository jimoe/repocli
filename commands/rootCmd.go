package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/config"
)

func Execute(cfg config.Config) {
	var longDescription = fmt.Sprintf(`Cli to simplify jobs regarding repoes.

Help:
	%[1]s help [command]
`,
		cfg.CliName)

	rootCmd := &cobra.Command{
		Use:     cfg.CliName,
		Version: "dummy",
		Short:   "Cli to handle repoes",
		Long:    longDescription,
	}

	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(getDirCmd)
	rootCmd.AddCommand(editorCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
