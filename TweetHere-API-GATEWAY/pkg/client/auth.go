package client

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/config"
	pb "TweetHere-API/pkg/pb/auth"
	"TweetHere-API/pkg/utils/models"
	"context"
	"errors"
	"fmt"
	"strconv"

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

	user, err := ad.Client.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    userdetails.Email,
		Password: userdetails.Password,
	})
	if err != nil {
		fmt.Println("zoooooooooooooooooooo", err)
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

func (ad *authClient) UserUpdateProfile(userdetails models.UserProfile, id int) (models.UserProfileResponse, error) {
	fmt.Println("isssss", id)
	user, err := ad.Client.UserUpdateProfile(context.Background(), &pb.UserUpdateProfileRequest{
		Update: &pb.UserUpdateProfile{
			Firstname:   userdetails.Firstname,
			Lastname:    userdetails.Lastname,
			Username:    userdetails.Username,
			Phone:       userdetails.Phone,
			Email:       userdetails.Email,
			DateOfBirth: userdetails.DateOfBirth,
			Profile:     userdetails.Profile,
			Bio:         userdetails.Bio,
		},
		Id: int64(id),
	})
	if err != nil {
		return models.UserProfileResponse{}, errors.New("error in adding user profile")
	}

	return models.UserProfileResponse{
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Username:    user.Username,
		Phone:       user.Phone,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Profile:     user.Profile,
		Bio:         userdetails.Bio,
	}, nil
}

func (ad *authClient) BlockUser(id string) error {
	user_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	_, berr := ad.Client.BlockUser(context.Background(), &pb.BlockUserRequest{
		Id: user_id,
	})
	if berr != nil {
		return berr
	}
	return nil
}

func (ad *authClient) UnBlockUser(id string) error {
	user_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	_, berr := ad.Client.UnBlockUser(context.Background(), &pb.UnBlockUserRequest{
		Id: user_id,
	})
	if berr != nil {

		return berr
	}
	return nil
}

func (ad *authClient) GetUser(page int) ([]models.UserDetails, error) {
	res, err := ad.Client.GetUser(context.Background(), &pb.GetUserRequest{
		Page: int64(page),
	})
	if err != nil {
		return nil, errors.New("error in getting user details")
	}

	var userdetails []models.UserDetails
	for _, ud := range res.Users {
		userdetails = append(userdetails, models.UserDetails{
			ID:          uint(ud.Id),
			Firstname:   ud.Firstname,
			Lastname:    ud.Lastname,
			Username:    ud.Username,
			Phone:       ud.Phone,
			Email:       ud.Email,
			DateOfBirth: ud.DateOfBirth,
			IsBlocked:   ud.Isblocked,
			Profile:     ud.Profile,
			Bio:         ud.Bio,
		})
	}

	return userdetails, nil
}

func (ad *authClient) ChangePassword(id int, old string, new string, re string) error {
	_, err := ad.Client.ChangePassword(context.Background(), &pb.ChangePasswordRequest{
		Id:          int64(id),
		Oldpassword: old,
		Newpassword: new,
		Repassword:  re,
	})
	if err != nil {
		return err
	}
	return nil
}
