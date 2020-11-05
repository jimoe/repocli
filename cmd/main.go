package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/jimoe/editor-and-change-dir/commands"
	"github.com/jimoe/editor-and-change-dir/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config: %w", err)
	}

	commands.Execute(cfg)
}
