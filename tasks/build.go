package tasks

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jimoe/repocli/config"
)

func Build(cfg *config.Config) error {
	fmt.Println("Building cli...")
	return build(cfg.Cli.BinPath, cfg.Cli.SourcePath, cfg.CliName)
}

func build(binPath, sourcePath, cliName string) error {
	outputFile := fmt.Sprintf("%s/%s", binPath, cliName)

	cmd := exec.Command("go", "build", "-o", outputFile, "cmd/main.go")
	cmd.Dir = sourcePath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build cli: %w", err)
	}

	return nil
}
