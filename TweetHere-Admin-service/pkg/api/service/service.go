package service

import (
	"Tweethere-Auth/pkg/pb"
	interfaces "Tweethere-Auth/pkg/usecase/interface"
	"Tweethere-Auth/pkg/utils/models"
	"context"
	"fmt"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AuthServiceServer {
	return &AdminServer{
		adminUseCase: useCase,
	}
}

func (ad *AdminServer) AdminSignUp(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	adminSignup := models.AdminSignUp{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	fmt.Println("service", adminSignup)

	res, err := ad.adminUseCase.AdminSignUp(adminSignup)
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

func (ad *AdminServer) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	fmt.Println("ssssssssssss", req)
	adminLogin := models.AdminLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Println("adminnnnnnnnnnnn",adminLogin)

	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
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
