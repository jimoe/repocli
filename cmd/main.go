package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jimoe/repocli/commands"
	"github.com/jimoe/repocli/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil && isFatalError(err) {
		fmt.Printf("config: %s", err.Error())
		os.Exit(1)
	}

	commands.Execute(cfg)
}

func isFatalError(err error) bool {
	var e *os.PathError
	if errors.As(err, &e) {
		if os.IsNotExist(e.Unwrap()) {
			return false
		}
	}

	return true
}
