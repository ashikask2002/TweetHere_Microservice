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
	UserID := req.Id
	passworddetails := models.ChangePassword{
		Oldpassword: req.Oldpassword,
		NewPassword: req.Newpassword,
		RePassword:  req.Repassword,
	}
	err := ad.authUseCase.ChangePassword(int(UserID), passworddetails)
	if err != nil {
		return &pb.ChangePasswordResponse{
			Error: err.Error(),
		}, nil
	}

	return &pb.ChangePasswordResponse{}, nil

}

func (ad *AuthServer) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	id := req.Id

	users, err := ad.authUseCase.GetUserDetails(int(id))
	if err != nil {
		return &pb.GetUserDetailsResponse{}, err
	}
	var userdetails []*pb.GetUserDetailsforUser

	for _, user := range users {
		userdetails = append(userdetails, &pb.GetUserDetailsforUser{
			Id:          int64(user.ID),
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Username:    user.Username,
			Phone:       user.Phone,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			Profile:     user.Profile,
			Bio:         user.Bio,
		})
	}
	return &pb.GetUserDetailsResponse{
		Userdetails: userdetails,
		Error:       "nil",
	}, nil

}

func (us *AuthServer) UserOTPLogin(ctx context.Context, req *pb.UserOTPLoginRequest) (*pb.UserOTPLoginResponse, error) {
	otp, err := us.authUseCase.UserOTPLogin(req.Email)
	if err != nil {
		return &pb.UserOTPLoginResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.UserOTPLoginResponse{
		Status: 200,
		Otp:    otp,
	}, nil
}

func (us *AuthServer) OtpVerification(ctx context.Context, req *pb.OtpVerificationRequest) (*pb.OtpVerificationResponse, error) {
	verified, err := us.authUseCase.OtpVerification(req.Email, req.Otp)
	if err != nil {
		return &pb.OtpVerificationResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.OtpVerificationResponse{
		Status:   200,
		Verified: verified,
	}, nil
}

func (ad *AuthServer) FollowReq(ctx context.Context, req *pb.FollowReqRequest) (*pb.FollowReqResponse, error) {
	id, userId := req.UserID, req.FollowingUser

	err := ad.authUseCase.FollowReq(int(id), int(userId))
	if err != nil {
		return &pb.FollowReqResponse{}, err
	}
	return &pb.FollowReqResponse{}, nil

}

func (ad *AuthServer) AcceptFollowReq(ctx context.Context, req *pb.AcceptFollowReqRequest) (*pb.AcceptFollowReqResponse, error) {
	id, userID := req.UserID, req.FollowingUser

	err := ad.authUseCase.AcceptFollowReq(int(id), int(userID))
	if err != nil {
		return &pb.AcceptFollowReqResponse{}, err
	}
	return &pb.AcceptFollowReqResponse{}, nil
}

func (ad *AuthServer) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*pb.UnfollowResponse, error) {
	id, UserID := req.UserID, req.FollowingUser

	err := ad.authUseCase.Unfollow(int(id), int(UserID))
	if err != nil {
		return &pb.UnfollowResponse{}, err
	}
	return &pb.UnfollowResponse{}, nil
}

func (ad *AuthServer) Followers(ctx context.Context, req *pb.FollowersRequest) (*pb.FollowersResponse, error) {
	id := req.UserID

	details, err := ad.authUseCase.Followers(int(id))
	if err != nil {
		return &pb.FollowersResponse{}, err
	}
	var userdetails []*pb.FollowResponse

	for _, user := range details {
		userdetails = append(userdetails, &pb.FollowResponse{
			Username:    user.Username,
			UserProfile: user.Profile,
		})
	}
	return &pb.FollowersResponse{
		Users: userdetails,
	}, nil

}

func (ad *AuthServer) Followings(ctx context.Context, req *pb.FollowingRequest) (*pb.FollowingResponse, error) {
	id := req.UserID

	details, err := ad.authUseCase.Followings(int(id))
	if err != nil {
		return &pb.FollowingResponse{}, err
	}
	var userdetails []*pb.FollowResponse

	for _, user := range details {
		userdetails = append(userdetails, &pb.FollowResponse{
			Username:    user.Username,
			UserProfile: user.Profile,
		})
	}
	return &pb.FollowingResponse{
		Users: userdetails,
	}, nil
}

func (ad *AuthServer) SendOTP(ctx context.Context, req *pb.SendOTPRequest) (*pb.SendOTPResponse, error) {
	phone := req.Phone

	err := ad.authUseCase.SendOTP(phone)
	if err != nil {
		return &pb.SendOTPResponse{}, err
	}
	return &pb.SendOTPResponse{}, nil
}

func (ad *AuthServer) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	udetails := models.VerifyData{
		PhoneNumber: req.Phone,
		Code:        req.Code,
	}
	details, err := ad.authUseCase.VerifyOTP(udetails)
	if err != nil {
		return &pb.VerifyOTPResponse{}, err
	}
	userdetails := pb.UserInfo{
		Id:        int64(details.User.ID),
		Firstname: details.User.Firstname,
		Lastname:  details.User.Lastname,
		Username:  details.User.Username,
		Email:     details.User.Email,
	}
	return &pb.VerifyOTPResponse{
		Info:         &userdetails,
		AccessToken:  details.AccesToken,
		RefreshToken: details.RefreshToken,
	}, nil
}
