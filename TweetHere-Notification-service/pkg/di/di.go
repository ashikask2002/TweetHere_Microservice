package di

import (
	server "tweethere-Notification/pkg/api"
	"tweethere-Notification/pkg/api/service"
	"tweethere-Notification/pkg/client"
	"tweethere-Notification/pkg/config"
	"tweethere-Notification/pkg/db"
	"tweethere-Notification/pkg/repository"
	"tweethere-Notification/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	notiRepository := repository.NewnotiRepository(gormDB)
	noticlient := client.NewAuthClient(&cfg)
	noriUseCase := usecase.NewnotiUsecase(notiRepository, noticlient)
	notiServiceServer := service.NewnotiServer(noriUseCase)
	grpcserver, err := server.NewGRPCServer(cfg, notiServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go noriUseCase.ConsumeNotification()
	return grpcserver, nil
}
