package di

import (
	server "tweet-service/pkg/api"
	"tweet-service/pkg/api/service"
	"tweet-service/pkg/client"
	"tweet-service/pkg/config"
	"tweet-service/pkg/db"
	"tweet-service/pkg/repository"
	"tweet-service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	tweetRepository := repository.NewTweetRespository(gormDB)
	tweetClient := client.NewAuthClient(&cfg)
	tweetUsecase := usecase.NewTweetUseCase(tweetRepository, tweetClient)
	tweetServiceServer := service.NewTweetServer(tweetUsecase)
	grpcServer, err := server.NewGRPCServer(cfg, tweetServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
