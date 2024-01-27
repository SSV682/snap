package main

import (
	"flag"

	"worker/internal/app"
)

var (
	configPath = flag.String("config", "./config/config.yaml", "path to config file. default: ./config/config.yaml")
)

func main() {
	flag.Parse()

	application := app.NewApp(*configPath)
	application.Run()
}
