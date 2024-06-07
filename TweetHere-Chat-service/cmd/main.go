package main

import (
	"fmt"
	"tweethere-chat/pkg/config"
	"tweethere-chat/pkg/di"
	"tweethere-chat/pkg/logging"
)

func main() {
	logging.Init()
	logEntry := logging.GetLogger().WithField("context", "loading config")
	config, configErr := config.LoadConfig()
	if configErr != nil {
		// log.Fatal("cannot load config: ", configErr)
		logEntry.WithError(configErr).Fatal("cannot log config")
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		// log.Fatal("cannot start server: ", diErr)
		logEntry.WithError(diErr).Fatal("cannot start server")
	} else {
		logEntry = logging.GetLogger().WithField("context", "starting server")
		fmt.Println("sssssssssssssss", logEntry)
		server.Start()
	}

}
