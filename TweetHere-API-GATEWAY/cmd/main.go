package main

import (
	"TweetHere-API/pkg/config"
	"TweetHere-API/pkg/di"
	"log"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	server, diErr := di.InitailizeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server", diErr)
	} else {
		server.Start()
	}
}
