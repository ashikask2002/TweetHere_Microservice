package main

import (
	"TweetHere-API/pkg/config"
	"TweetHere-API/pkg/di"
	"fmt"
	"log"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	fmt.Println("heloooooooooooooooooooooooooo")
	server, diErr := di.InitailizeAPI(config)
	fmt.Println("serverrrrrrrrrrrrrrrrr")

	if diErr != nil {
		log.Fatal("cannot start server", diErr)
	} else {
		server.Start()
	}
}
