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

	rootCmd.AddCommand(createGetDirCmd(cfg))
	rootCmd.AddCommand(createTabTitleCmd(cfg))
	rootCmd.AddCommand(createEditorCmd(cfg))
	configCmd := createConfigCmd(cfg)
	configCmd.AddCommand(createConfigExampleCmd(cfg))
	configCmd.AddCommand(createConfigInitCmd(cfg))
	rootCmd.AddCommand(configCmd)

	_ = rootCmd.Execute()
}

// exit prints the error then runs os.Exit().
// If cmd is not nil then cmd.usage() is executed before os.Exit().
func exit(err error, cmd *cobra.Command) {
	fmt.Printf("Error: %s\n", err.Error())

	if cmd != nil {
		fmt.Println()
		_ = cmd.Usage()
	}

	os.Exit(1)
}
