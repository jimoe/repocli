package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

const longDescription = `Cli to simplify jobs regarding repoes.

An <alias> may only contain the characters %[1]s
`

func Execute(cfg *config.Config) {
	rootCmd := &cobra.Command{
		Use:     cfg.CliName,
		Version: cfg.Version,
		Short:   "Cli to handle repoes",
		Long:    fmt.Sprintf(longDescription, arguments.ValidAliasChars),
	}

	rootCmd.AddCommand(buildCmd(cfg))
	rootCmd.AddCommand(getDirCmd(cfg))
	rootCmd.AddCommand(tabTitleCmd(cfg))
	rootCmd.AddCommand(editorCmd(cfg))
	rootCmd.AddCommand(configCmd(cfg))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
