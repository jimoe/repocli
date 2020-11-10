package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/aliases"
	"github.com/jimoe/editor-and-change-dir/config"
)

const longDescription = `Cli to simplify jobs regarding repoes.

An <alias> may only contain the characters %[1]s

Yaml config
Place config in same dir and same name as executable. Just add '.yml' to the filename
example:

repoes:
  - name:    some-repo-name
    path:    /home/username/code/some-repo-name
    editor:  goland
    aliases:
      - some
			- some-repo
  - name:    another-repo-name
    path:    /home/username/code/another-repo-name
    editor:  code
`

func Execute(cfg config.Config) {
	rootCmd := &cobra.Command{
		Use:     cfg.CliName,
		Version: cfg.Version,
		Short:   "Cli to handle repoes",
		Long:    fmt.Sprintf(longDescription, aliases.ValidAliasChars),
	}

	rootCmd.AddCommand(buildCmd(cfg))
	rootCmd.AddCommand(getDirCmd(cfg))
	rootCmd.AddCommand(editorCmd(cfg))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
