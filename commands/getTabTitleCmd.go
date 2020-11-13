package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/tasks"
)

const shortDescription = "Get the terminal tab title that is associated with the given <full path> (pwd)."

func getTabTitleCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "gettabtitle <full path>",
		Short: shortDescription,
		Long:  shortDescription + " If path is not found then nothing is returned",

		Args: cobra.ExactArgs(1),

		DisableFlagsInUseLine: true,

		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			if err := validatePath(path); err != nil {
				fmt.Printf("Error: %s\n\n", err.Error())
				_ = cmd.Usage()
				os.Exit(1)
			}

			tasks.GetTabTitle(cfg, path)
		},
	}
}

func validatePath(path string) error {
	validChars := "a-zA-Z0-9-_/."
	regStr := fmt.Sprintf(`^[%s]+$`, validChars)
	if regexp.MustCompile(regStr).MatchString(path) {
		return nil
	}
	return fmt.Errorf("illegal character in <path> (%s)", validChars)
}
