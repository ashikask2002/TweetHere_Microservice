package client

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/config"
	pb "TweetHere-API/pkg/pb/auth"
	"TweetHere-API/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AuthServiceClient
}

// func NewAdminClient(cfg config.Config) (interfaces.AdminClient, error) {
//     fmt.Println("clent", cfg.AdminSvcUrl)
//     grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
//     if err != nil {
//         fmt.Println("could not connect ", err)
//         return nil, err // Return error
//     }
//     grpcClient := pb.NewAuthServiceClient(grpcConnection)

//     return &adminClient{
//         Client: grpcClient,
//     }, nil // No error
// }

func NewAdminClient(cfg config.Config) interfaces.AdminClient {
	fmt.Println("cfgggggggggg", cfg)
	fmt.Println("clent", cfg.AdminSvcUrl)
	grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect ", err)
	}
	grpcClient := pb.NewAuthServiceClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}
}

func (ad *adminClient) AdminSignUp(admindetails models.AdminSignup) (models.TokenAdmin, error) {
	fmt.Println("ddd", admindetails)
	admin, err := ad.Client.AdminSignUp(context.Background(), &pb.AdminSignupRequest{
		Firstname: admindetails.Firstname,
		Lastname:  admindetails.Lastname,
		Email:     admindetails.Email,
		Password:  admindetails.Password,
	})
	fmt.Println("error is", err)
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}

func (ad *adminClient) AdminLogin(admindetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email:    admindetails.Email,
		Password: admindetails.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}
