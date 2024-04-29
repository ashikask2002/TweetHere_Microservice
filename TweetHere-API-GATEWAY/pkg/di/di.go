package di

import (
	server "TweetHere-API/pkg/api"
	"TweetHere-API/pkg/api/handler"
	"TweetHere-API/pkg/client"
	"TweetHere-API/pkg/config"
)

func InitailizeAPI(cfg config.Config) (*server.ServerHttp, error) {
	adminClient := client.NewAuthClient(cfg)
	authHandler := handler.NewAuthHandler(adminClient)

	serverHTTP := server.NewServerHTTP(authHandler)
	return serverHTTP, nil
}
