package service

import (
	"Tweethere-Auth/pkg/pb"
	interfaces "Tweethere-Auth/pkg/usecase/interface"
	"Tweethere-Auth/pkg/utils/models"
	"context"
	"errors"
	"fmt"
)

type AuthServer struct {
	authUseCase interfaces.AdminUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(useCase interfaces.AdminUseCase) pb.AuthServiceServer {
	return &AuthServer{
		authUseCase: useCase,
	}
}

func (ad *AuthServer) AdminSignUp(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	adminSignup := models.AdminSignUp{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	fmt.Println("service", adminSignup)

	res, err := ad.authUseCase.AdminSignUp(adminSignup)
	if err != nil {
		return &pb.AdminSignupResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(res.Admin.ID),
		Firstname: res.Admin.Firstname,
		Lastname:  res.Admin.Lastname,
		Email:     res.Admin.Email,
	}
	return &pb.AdminSignupResponse{
		Status:       201,
		AdminDetails: adminDetails,
		Token:        res.Token,
	}, nil
}

func (ad *AuthServer) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	fmt.Println("ssssssssssss", req)
	adminLogin := models.AdminLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Println("adminnnnnnnnnnnn", adminLogin)

	admin, err := ad.authUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}

	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}

func (ad *AuthServer) UserSignup(ctx context.Context, req *pb.UserSignupRequest) (*pb.UserSignupResponse, error) {
	userSignUp := models.UserSignup{
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Username:    req.Username,
		Phone:       req.Phone,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Password:    req.Password,
	}

	data, err := ad.authUseCase.UserSignup(userSignUp)
	if err != nil {
		return &pb.UserSignupResponse{}, errors.New("error in signing up")
	}
	userDetials := pb.UserResponse{
		Info: &pb.UserInfo{
			Id:        int64(data.User.ID),
			Firstname: data.User.Firstname,
			Lastname:  data.User.Lastname,
			Username:  data.User.Username,
			Email:     data.User.Email,
		},
		AccessToken:  data.AccesToken,
		RefreshToken: data.RefreshToken,
	}
	return &pb.UserSignupResponse{
		Response: &userDetials,
		Error:    " ",
	}, err
}

func (ad *AuthServer) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	fmt.Println("userloginnnnnnnnn", req)
	userLogin := models.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	data, err := ad.authUseCase.UserLogin(userLogin)
	if err != nil {
		return &pb.UserLoginResponse{}, errors.New("error in userlogin")
	}
	userdeteils := pb.UserResponse{
		Info: &pb.UserInfo{
			Id:        int64(data.User.ID),
			Firstname: data.User.Firstname,
			Lastname:  data.User.Lastname,
			Username:  data.User.Username,
			Email:     data.User.Email,
		},
		AccessToken:  data.AccesToken,
		RefreshToken: data.RefreshToken,
	}
	return &pb.UserLoginResponse{
		Respone: &userdeteils,
		Error:   " ",
	}, err

}
