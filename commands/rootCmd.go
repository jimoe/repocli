package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/validate"
)

func Execute(cfg config.Config) {
	var longDescription = fmt.Sprintf(`Cli to simplify jobs regarding repoes.

An alias my only contain the characters ½[2]s

Help:
	%[1]s help [command]
`,
		cfg.CliName,
		validate.ValidAliasChars)

	rootCmd := &cobra.Command{
		Use:     cfg.CliName,
		Version: cfg.Version,
		Short:   "Cli to handle repoes",
		Long:    longDescription,
	}

	rootCmd.AddCommand(buildCmd(cfg))
	rootCmd.AddCommand(getDirCmd(cfg))
	rootCmd.AddCommand(editorCmd(cfg))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
