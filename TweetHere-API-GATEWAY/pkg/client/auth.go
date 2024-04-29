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

type authClient struct {
	Client pb.AuthServiceClient
}

func NewAuthClient(cfg config.Config) interfaces.AdminClient {
	fmt.Println("clent", cfg.AdminSvcUrl)
	grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect ", err)
	}
	grpcClient := pb.NewAuthServiceClient(grpcConnection)

	return &authClient{
		Client: grpcClient,
	}
}

func (ad *authClient) AdminSignUp(admindetails models.AdminSignup) (models.TokenAdmin, error) {
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

func (ad *authClient) AdminLogin(admindetails models.AdminLogin) (models.TokenAdmin, error) {
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

//userrr ///////////////

func (ad *authClient) UserSignup(userdetails models.UserSignup) (models.TokenUser, error) {
	user, err := ad.Client.UserSignup(context.Background(), &pb.UserSignupRequest{
		Firstname:   userdetails.Firstname,
		Lastname:    userdetails.Lastname,
		Username:    userdetails.Username,
		Phone:       userdetails.Phone,
		Email:       userdetails.Email,
		DateOfBirth: userdetails.DateOfBirth,
		Password:    userdetails.Password,
	})
	if err != nil {
		return models.TokenUser{}, err
	}

	return models.TokenUser{
		User: models.UserDetailsResponse{
			ID:        uint(user.Response.Info.Id),
			Firstname: user.Response.Info.Firstname,
			Lastname:  user.Response.Info.Lastname,
			Username:  user.Response.Info.Username,
			Email:     user.Response.Info.Email,
		},
		AccesToken:   user.Response.AccessToken,
		RefreshToken: user.Response.RefreshToken,
	}, nil
}

func (ad *authClient) UserLogin(userdetails models.UserLogin) (models.TokenUser, error) {
	fmt.Println("loginnnnnnnnnnnnnnndetailsss", userdetails)
	user, err := ad.Client.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    userdetails.Email,
		Password: userdetails.Password,
	})

	fmt.Println("login222222222222222", user)
	if err != nil {
		return models.TokenUser{}, err
	}

	return models.TokenUser{
		User: models.UserDetailsResponse{
			ID:        uint(user.Respone.Info.Id),
			Firstname: user.Respone.Info.Firstname,
			Lastname:  user.Respone.Info.Lastname,
			Username:  user.Respone.Info.Username,
			Email:     user.Respone.Info.Email,
		},
		AccesToken:   user.Respone.AccessToken,
		RefreshToken: user.Respone.RefreshToken,
	}, nil
}


func (ad *authClient) UserUpdateProfile(userdetails models.UserProfile) (models.UserProfileResponse,error){
	user, err := ad.Client.UserUpdateProfile(context.Background(),&pb.UserUpdateProfileRequest{
		Firstname: userdetails.Firstname,
		Lastname: userdetails.Lastname,
		Username: userdetails.Username,
		Phone: userdetails.Phone,
		Email: userdetails.Email,
		Profile: userdetails.Profile,
		Bio: userdetails.Bio,
	})
}
