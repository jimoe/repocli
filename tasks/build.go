package tasks

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jimoe/repocli/config"
)

func Build(cfg *config.Config) error {
	// Howto build cli manually when developing cli: `go build -o ~/bin/repocli cmd/main.go`
	fmt.Println("Building cli...")

	outputFile := fmt.Sprintf("%s/%s", cfg.Cli.BinPath, cfg.CliName)

	cmd := exec.Command("go", "build", "-o", outputFile, "cmd/main.go")
	cmd.Dir = cfg.Cli.SourcePath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build cli: %w", err)
	}

	return nil
}
