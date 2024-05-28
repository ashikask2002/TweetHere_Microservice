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

	tweetClient := client.NewTweetClient(cfg)
	tweetHandler := handler.NewTweetHandler(tweetClient)

	chatClient := client.NewChatClient(cfg)
	chatHandler := handler.NewChatHandler(chatClient)

	noticlient := client.NewNotificationClient(cfg)
	notihandler := handler.NewNotificationHandler(noticlient)

	serverHTTP := server.NewServerHTTP(authHandler, tweetHandler, chatHandler, notihandler)
	return serverHTTP, nil
}
