package main

import (
	_ "TweetHere-API/cmd/docs"
	"TweetHere-API/pkg/config"
	"TweetHere-API/pkg/di"
	"TweetHere-API/pkg/logging"
	"fmt"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Go + Gin SocialMedia TweetHereeee
// @version 1.0.0
// @description Tweet Here is a social Media Platform
// @contact.name API Support
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
// @host localhost:5000
// @BasePath /
// @query.collection.format multi

func main() {

	logging.Init()

	logEntry := logging.GetLogger().WithField("context", "loading config")

	config, configErr := config.LoadConfig()
	if configErr != nil {
		// log.Fatal("cannot load config: ", configErr)
		logEntry.WithError(configErr).Fatal("cannot log config")
	}

	logEntry = logging.GetLogger().WithField("contex", "Initializing API")
	server, diErr := di.InitailizeAPI(config)

	if diErr != nil {
		// log.Fatal("cannot start server", diErr)
		logEntry.WithError(diErr).Fatal("cannot start server")
	} else {
		logEntry = logging.GetLogger().WithField("context", "starting server")
		fmt.Println("sssssssssssssss", logEntry)
		server.Start()
	}
}
