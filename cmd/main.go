package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jimoe/editor-and-change-dir/commands"
	"github.com/jimoe/editor-and-change-dir/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("config: %w", err))
	}

	commands.Execute(cfg)
}
