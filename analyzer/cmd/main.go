package main

import (
	"flag"

	"worker/internal/app"
)

var (
	configPath = flag.String("config", "./config/config.yaml", "path to config file. default: ./config/config.yaml")
)

// main is the entry point of the application.
func main() {
	flag.Parse()

	// application is the instance of the application.
	application := app.NewApp(*configPath)
	// Run starts the application.
	application.Run()
}
