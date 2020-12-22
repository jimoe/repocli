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
	cmd := &cobra.Command{
		Use:   "init",
		Short: `Write config file based on the example in you os user specific config dir`,
		Args:  cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			forceOverwrite, err := cmd.Flags().GetBool("force")
			if err != nil {
				exit(err, cmd)
			}

			err = tasks.ConfigInit(cfg, forceOverwrite)
			if err != nil {
				exit(err, nil)
			}
		},
	}

	cmd.Flags().BoolP(
		"force",
		"",
		false,
		"If config file already exists, this will force the example config to overwrite the existing config.",
	)

	return cmd
}

func configWhereCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "where",
		Short: "write path to where the config is/will be stored",
		Args:  cobra.ExactArgs(0),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cfg.Yaml.PathAndFilename)
		},
	}
}
