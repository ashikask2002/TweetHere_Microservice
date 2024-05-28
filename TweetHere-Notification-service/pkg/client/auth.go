package client

import (
	"context"
	"fmt"
	"tweethere-Notification/pkg/config"
	pb "tweethere-Notification/pkg/pb/auth"
	"tweethere-Notification/pkg/utils/models"

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

func (ad *authClient) UserData(userid int) (models.UserData, error) {
	fmt.Println("iddd", userid)
	data, err := ad.Client.UserData(context.Background(), &pb.UserDataRequest{
		Userid: int64(userid),
	})
	fmt.Println("wwwwwwwwwwwwwwww", data.Userid, data.Username, data.Profile)
	if err != nil {
		fmt.Println("ssssss", err)
		return models.UserData{}, err

	}
	return models.UserData{
		UserId:   int(data.Userid),
		Username: data.Username,
		Profile:  data.Profile,
	}, nil
}
