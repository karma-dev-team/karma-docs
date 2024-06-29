package main

import (
	"log"

	"github.com/karma-dev-team/karma-docs/internal/config"
	"github.com/karma-dev-team/karma-docs/internal/server"
)

func main() {
	config, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp(config)

	if err := app.Run(config.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
