package di

import (
	server "TweetHere-API/pkg/api"
	"TweetHere-API/pkg/api/handler"
	"TweetHere-API/pkg/client"
	"TweetHere-API/pkg/config"
)

func InitailizeAPI(cfg config.Config) (*server.ServerHttp, error) {
	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	serverHTTP := server.NewServerHTTP(adminHandler)
	return serverHTTP, nil
}
