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
	adminClient := client.NewAdminClient(cfg)

	adminHandler := handler.NewAdminHandler(adminClient)

	serverHTTP := server.NewServerHTTP(adminHandler)
	return serverHTTP, nil
}
