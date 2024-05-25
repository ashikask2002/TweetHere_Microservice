package di

import (
	server "tweethere-chat/pkg/api"
	"tweethere-chat/pkg/api/service"
	"tweethere-chat/pkg/config"
	"tweethere-chat/pkg/db"
	"tweethere-chat/pkg/repository"
	"tweethere-chat/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	database, err := db.ConnectDatabase(cfg)

	if err != nil {
		return nil, err
	}

	chatRepository := repository.NewChatRepository(database)
	chatUseCase := usecase.NewChatUseCase(chatRepository)
	chatServiceServer := service.NeChatServer(chatUseCase)

	gprcServer, err := server.NewGRPCServer(cfg, chatServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go chatUseCase.MessageConsumer()
	return gprcServer, nil
}
