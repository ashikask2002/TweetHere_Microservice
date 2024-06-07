package main

import (
	"fmt"
	"tweethere-Notification/pkg/config"
	"tweethere-Notification/pkg/di"
	"tweethere-Notification/pkg/logging"
)

func main() {
	logging.Init()
	logEntry := logging.GetLogger().WithField("context", "loading config")

	config, configErr := config.LoadConfig()
	if configErr != nil {
		// log.Fatal("cannot load config: ", configErr)
		logEntry.WithError(configErr).Fatal("cannot log config")
	}
	logEntry = logging.GetLogger().WithField("contex", "Initializing API")
	server, diErr := di.InitializeApi(config)
	if diErr != nil {
		// log.Fatal("cannot start server: ", diErr)
		logEntry.WithError(diErr).Fatal("cannot start server")
	} else {
		logEntry = logging.GetLogger().WithField("context", "starting server")
		fmt.Println("sssssssssssssss", logEntry)
		server.Start()
	}

}
