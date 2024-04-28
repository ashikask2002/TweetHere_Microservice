package di

import (
	server "TweetHere-API/pkg/api"
	"TweetHere-API/pkg/api/handler"
	"TweetHere-API/pkg/client"
	"TweetHere-API/pkg/config"
	"fmt"
)

func InitailizeAPI(cfg config.Config) (*server.ServerHttp, error) {
	fmt.Println("vbbbbbbbbbbbbbbbbbbbbbbb")
	adminClient := client.NewAuthClient(cfg)
	authHandler := handler.NewAuthHandler(adminClient)

	serverHTTP := server.NewServerHTTP(authHandler)
	return serverHTTP, nil
}
