package execute

import (
	"os"
	"os/exec"

	"github.com/jimoe/editor-and-change-dir/color"
)

func Run(runDir string, command string, args ...string) {
	cmd := exec.Command(command, args...)

	cmd.Dir = runDir
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red.Printf("Error: Failed to run command (%s) (%s)\n", command, runDir)
		// fmt.Println("----------------------", err)
		os.Exit(1)
	}
}
