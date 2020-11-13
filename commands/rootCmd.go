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
for editors: if params is <path> then use repo path as param when starting editor
example:

editors:
	-	name: goland
		params: nosplash <path>
	- name: code
		params: .
repoes:
  - name:    some-repo-name
    path:    /home/username/code/some-repo-name
    editor:  goland
    aliases:
      - some
			- some-repo
    terminal:
      title: SOME
  - name:    another-repo-name
    path:    /home/username/code/another-repo-name
    editor:  code
    monorepo:
			- subpath: packages/name
				terminal:
          title: A name
			- subpath: packages/whatever
				terminal:
          title: A whatever
`

func Execute(cfg *config.Config) {
	rootCmd := &cobra.Command{
		Use:     cfg.CliName,
		Version: cfg.Version,
		Short:   "Cli to handle repoes",
		Long:    fmt.Sprintf(longDescription, aliases.ValidAliasChars),
	}

	rootCmd.AddCommand(buildCmd(cfg))
	rootCmd.AddCommand(getDirCmd(cfg))
	rootCmd.AddCommand(getTabTitleCmd(cfg))
	rootCmd.AddCommand(editorCmd(cfg))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
