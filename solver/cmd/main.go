package main

import (
	"flag"

	"solver/internal/app"
	"solver/internal/config"

	log "github.com/sirupsen/logrus"
)

var (
	configPath = flag.String("config", "./config/config.yaml", "path to config file. default: ./config/config.yaml")
)

// main is the entry point of the application.
func main() {
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	// application is the instance of the application.
	application := app.NewApp(&cfg)
	// Run starts the application.
	application.Run()
}
