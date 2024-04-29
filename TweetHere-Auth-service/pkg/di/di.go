package di

import (
	server "Tweethere-Auth/pkg/api"
	"Tweethere-Auth/pkg/api/service"
	"Tweethere-Auth/pkg/config"
	"Tweethere-Auth/pkg/db"
	"Tweethere-Auth/pkg/repository"
	"Tweethere-Auth/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAuthRepository(gormDB)
	adminUsecase := usecase.NewAuthUseCase(adminRepository)
	adminServiceServer := service.NewAuthServer(adminUsecase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
