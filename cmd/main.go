package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jimoe/repocli/commands"
	"github.com/jimoe/repocli/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil && isFatalError(err) {
		log.Fatal(fmt.Errorf("config: %w", err))
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
