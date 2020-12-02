package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jimoe/repocli/commands"
	"github.com/jimoe/repocli/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("config: %w", err))
	}

	commands.Execute(cfg)
}
