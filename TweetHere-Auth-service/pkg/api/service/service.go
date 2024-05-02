package service

import (
	"Tweethere-Auth/pkg/pb"
	interfaces "Tweethere-Auth/pkg/usecase/interface"
	"Tweethere-Auth/pkg/utils/models"
	"context"
	"errors"
	"fmt"
	"strconv"
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

	userLogin := models.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	data, err := ad.authUseCase.UserLogin(userLogin)
	if err != nil {
		return &pb.UserLoginResponse{}, err
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

func (ad *AuthServer) UserUpdateProfile(ctx context.Context, user *pb.UserUpdateProfileRequest) (*pb.UserUpdateProfileResponse, error) {
	userdetails := models.UserProfile{
		Firstname:   user.Update.Firstname,
		Lastname:    user.Update.Lastname,
		Username:    user.Update.Username,
		Phone:       user.Update.Phone,
		Email:       user.Update.Email,
		DateOfBirth: user.Update.DateOfBirth,
		Profile:     user.Update.Profile,
		Bio:         user.Update.Bio,
	}
	id := user.Id
	fmt.Println("id isssssss", id)

	data, err := ad.authUseCase.UserUpdateProfile(userdetails, int(id))
	if err != nil {
		return &pb.UserUpdateProfileResponse{}, errors.New("error happened in Auth service")
	}

	userdetailsresponse := pb.UserUpdateProfileResponse{
		Firstname:   data.Firstname,
		Lastname:    data.Lastname,
		Username:    data.Username,
		Phone:       data.Phone,
		Email:       data.Email,
		DateOfBirth: data.DateOfBirth,
		Profile:     data.Profile,
		Bio:         data.Bio,
	}

	return &userdetailsresponse, nil
}

func (ad *AuthServer) GetUser(ctx context.Context, page *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	paage := page.Page

	users, err := ad.authUseCase.GetUser(int(paage))
	if err != nil {
		return &pb.GetUserResponse{}, errors.New("problem in gettind userdetails")
	}

	var userdetails []*pb.UserDetailsAtAdmin

	for _, user := range users {
		userdetails = append(userdetails, &pb.UserDetailsAtAdmin{
			Id:          int64(user.ID),
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Username:    user.Username,
			Phone:       user.Phone,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			Isblocked:   user.IsBlocked,
			Profile:     user.Profile,
			Bio:         user.Bio,
		})
	}
	return &pb.GetUserResponse{
		Users: userdetails,
	}, nil

}

func (ad *AuthServer) BlockUser(ctx context.Context, req *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {
	userid := strconv.FormatUint(req.Id, 10)

	err := ad.authUseCase.BlockUser(userid)

	if err != nil {
		return &pb.BlockUserResponse{
			Succes: false,
			Error:  "error happened while blocking",
		}, nil
	}

	return &pb.BlockUserResponse{
		Succes: true,
	}, nil
}

func (ad *AuthServer) UnBlockUser(ctx context.Context, req *pb.UnBlockUserRequest) (*pb.UnBlockUserResponse, error) {
	userid := strconv.FormatUint(req.Id, 10)
	fmt.Println("aaaaaaaaaaaaa", userid)

	err := ad.authUseCase.UnBlockUser(userid)

	if err != nil {
		return &pb.UnBlockUserResponse{
			Succces: false,
			Error:   "error happened while blocking",
		}, nil
	}

	return &pb.UnBlockUserResponse{
		Succces: true,
	}, nil
}

func (ad *AuthServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	err := ad.authUseCase.ChangePassword(int(req.Id), req.Oldpassword, req.Newpassword, req.Repassword)
	if err != nil {
		return &pb.ChangePasswordResponse{
			Error: err.Error(),
		}, nil
	}

	return &pb.ChangePasswordResponse{}, nil

}
