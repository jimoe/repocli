package tasks

import (
	"fmt"

	"github.com/jimoe/editor-and-change-dir/color"
	"github.com/jimoe/editor-and-change-dir/config"
	"github.com/jimoe/editor-and-change-dir/execute"
)

func Build(cfg config.Config) {
	// Howto build cli manually when developing cli: `go build -o ~/bin/editorAndChangeDirTest cmd/main.go`
	color.Println("Building cli...")

	binPath := fmt.Sprintf("%s%s", cfg.BinHome, cfg.CliName)
	runDir := cfg.SourceHome
	execute.Run(runDir, "go", "build", "-o", binPath, "cmd/main.go")
}
