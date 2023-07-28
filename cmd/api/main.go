package main

import (
	"log"

	config "github.com/sreenathsvrm/voix-me/pkg/config"
	di "github.com/sreenathsvrm/voix-me/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}