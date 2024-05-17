package client

import (
	"context"
	"fmt"
	"tweet-service/pkg/config"
	pb "tweet-service/pkg/pb/authh"

	"google.golang.org/grpc"
)

type authClient struct {
	Client pb.AuthServiceClient
}

func NewAuthClient(cfg *config.Config) *authClient {

	grpcConnection, err := grpc.Dial(cfg.AUTH_SVC_URL, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewAuthServiceClient(grpcConnection)

	return &authClient{
		Client: grpcClient,
	}
}

func (au *authClient) DoesUserExist(id int64) (bool, error) {
	res, err := au.Client.DoesUserExist(context.Background(), &pb.DoesUserExistRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("kkkkkkkkkk", err)
		return false, err
	}
	return res.Data, nil
}

func (au *authClient) FindUserName(id int64) (string, error) {
	res, err := au.Client.FindUserName(context.Background(), &pb.FindUserNameRequest{
		Id: id,
	})
	if err != nil {
		return "", err
	}
	return res.Name, nil
}
