package service

import (
	"Tweethere-Auth/pkg/logging"
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
	logEntry := logging.GetLogger().WithField("method", "AdminLogin")
	logEntry.Info("Processing AdminLogin request with email", req.GetEmail())

	fmt.Println("ssssssssssss", req)
	adminLogin := models.AdminLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Println("adminnnnnnnnnnnn", adminLogin)

	admin, err := ad.authUseCase.LoginHandler(adminLogin)
	if err != nil {
		logEntry.WithError(err).Error("Error logging in admin")
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

	logEntry := logging.GetLogger().WithField("method", "UserSignup")
	logEntry.Info("Processing UserSignup request with email", req.GetEmail())

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
		logEntry.WithError(err).Error("Error signing up user")
		return &pb.UserSignupResponse{}, err
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
	logEntry := logging.GetLogger().WithField("method", "UserLogin")
	logEntry.Info("Processing UserLogin request", req.GetEmail())

	userLogin := models.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	data, err := ad.authUseCase.UserLogin(userLogin)
	if err != nil {
		logEntry.WithError(err).Error("Error Login user")
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
	logEntry := logging.GetLogger().WithField("method", "UserUpdateProfile")
	logEntry.Info("Processing UserUpdateProfile request for user ID", user.GetId())

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
		logEntry.WithError(err).Error("Error updating user profile")
		return &pb.UserUpdateProfileResponse{}, err
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
	logEntry := logging.GetLogger().WithField("method", "GetUser")
	logEntry.Info("Processing GetUser request for page", page.GetPage())
	paage := page.Page

	users, err := ad.authUseCase.GetUser(int(paage))
	if err != nil {
		logEntry.WithError(err).Error("Error getting users")
		return &pb.GetUserResponse{}, err
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
	logEntry := logging.GetLogger().WithField("method", "BlockUser")

	// Log request details (including user ID)
	logEntry.Info("Processing BlockUser request for user ID", req.GetId())
	userid := strconv.FormatUint(req.Id, 10)

	err := ad.authUseCase.BlockUser(userid)

	if err != nil {
		logEntry.WithError(err).Error("Error blocking user")
		return &pb.BlockUserResponse{
			Succes: false,
			Error:  "error happened while blocking",
		}, err
	}
	logEntry.Info("blocked successfully")
	return &pb.BlockUserResponse{
		Succes: true,
	}, nil
}

func (ad *AuthServer) UnBlockUser(ctx context.Context, req *pb.UnBlockUserRequest) (*pb.UnBlockUserResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "UnBlockUser")
	logEntry.Info("Processing UnBlockUser request for user ID", req.GetId())

	userid := strconv.FormatUint(req.Id, 10)
	fmt.Println("aaaaaaaaaaaaa", userid)

	err := ad.authUseCase.UnBlockUser(userid)

	if err != nil {
		logEntry.WithError(err).Error("Error unblocking user")
		return &pb.UnBlockUserResponse{
			Succces: false,
			Error:   "error happened while blocking",
		}, err
	}
	logEntry.Info("unblocked successfully")
	return &pb.UnBlockUserResponse{
		Succces: true,
	}, nil
}

func (ad *AuthServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {

	logEntry := logging.GetLogger().WithField("method", "ChangePassword")
	logEntry.Info("Processing ChangePassword request for user ID", req.GetId())

	UserID := req.Id
	passworddetails := models.ChangePassword{
		Oldpassword: req.Oldpassword,
		NewPassword: req.Newpassword,
		RePassword:  req.Repassword,
	}
	err := ad.authUseCase.ChangePassword(int(UserID), passworddetails)
	if err != nil {
		logEntry.WithError(err).Error("Error changing password")
		return &pb.ChangePasswordResponse{
			Error: err.Error(),
		}, err
	}
	logEntry.Info("Password changed successfully for user ID", UserID)
	return &pb.ChangePasswordResponse{}, nil

}

func (ad *AuthServer) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "GetUserDetails")
	logEntry.Info("Processing GetUserDetails request for user ID", req.GetId())

	id := req.Id

	users, err := ad.authUseCase.GetUserDetails(int(id))
	if err != nil {
		logEntry.WithError(err).Error("Error getting user details")
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
	logEntry := logging.GetLogger().WithField("method", "UserOTPLogin")
	logEntry.Info("Processing UserOTPLogin request for email (masked)")
	otp, err := us.authUseCase.UserOTPLogin(req.Email)
	if err != nil {
		logEntry.WithError(err).Info("Error generating OTP for email (masked)")
		return &pb.UserOTPLoginResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	logEntry.Info("Generated OTP for email (masked)")
	return &pb.UserOTPLoginResponse{
		Status: 200,
		Otp:    otp,
	}, nil
}

func (us *AuthServer) OtpVerification(ctx context.Context, req *pb.OtpVerificationRequest) (*pb.OtpVerificationResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "OtpVerification")
	logEntry.Info("Processing OtpVerification request for email (masked), OTP", req.GetOtp())
	verified, err := us.authUseCase.OtpVerification(req.Email, req.Otp)
	if err != nil {
		logEntry.WithError(err).Info("Error verifying OTP for email (masked)")
		return &pb.OtpVerificationResponse{
			Status: 400,
			Error:  err.Error(),
		}, err
	}
	logEntry.Info("OTP verified successfully for email (masked)")
	return &pb.OtpVerificationResponse{
		Status:   200,
		Verified: verified,
	}, nil
}

func (ad *AuthServer) FollowReq(ctx context.Context, req *pb.FollowReqRequest) (*pb.FollowReqResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "FollowReq")
	logEntry.Info("Processing FollowReq request. User ID:", req.GetUserID(), " following user ID:", req.GetFollowingUser())

	id, userId := req.UserID, req.FollowingUser

	err := ad.authUseCase.FollowReq(int(id), int(userId))
	if err != nil {
		logEntry.WithError(err).Error("Error following user")
		return &pb.FollowReqResponse{}, err
	}
	logEntry.Info("successfully sended follow req")
	return &pb.FollowReqResponse{}, nil

}

func (ad *AuthServer) AcceptFollowReq(ctx context.Context, req *pb.AcceptFollowReqRequest) (*pb.AcceptFollowReqResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "AcceptFollowReq")
	logEntry.Info("Processing AcceptFollowReq request. User ID:", req.GetUserID(), " accepting follow request from user ID:", req.GetFollowingUser())

	id, userID := req.UserID, req.FollowingUser

	err := ad.authUseCase.AcceptFollowReq(int(id), int(userID))
	if err != nil {
		logEntry.WithError(err).Error("Error accepting follow request")
		return &pb.AcceptFollowReqResponse{}, err
	}
	logEntry.Info("User ID", userID, " successfully accepted follow request from user ID")
	return &pb.AcceptFollowReqResponse{}, nil
}

func (ad *AuthServer) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*pb.UnfollowResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "Unfollow")
	logEntry.Info("Processing Unfollow request. User ID:", req.GetUserID(), " unfollowing user ID:", req.GetFollowingUser())
	id, UserID := req.UserID, req.FollowingUser

	err := ad.authUseCase.Unfollow(int(id), int(UserID))
	if err != nil {
		logEntry.WithError(err).Error("Error unfollowing user")
		return &pb.UnfollowResponse{}, err
	}
	logEntry.Info(" successfully unfollowed user ")
	return &pb.UnfollowResponse{}, nil
}

func (ad *AuthServer) Followers(ctx context.Context, req *pb.FollowersRequest) (*pb.FollowersResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "Followers")
	logEntry.Info("Processing Followers request for user ID:", req.GetUserID())
	id := req.UserID

	details, err := ad.authUseCase.Followers(int(id))
	if err != nil {
		logEntry.WithError(err).Error("Error getting followers")
		return &pb.FollowersResponse{}, err
	}
	var userdetails []*pb.FollowResponse

	for _, user := range details {
		userdetails = append(userdetails, &pb.FollowResponse{
			Username:    user.Username,
			UserProfile: user.Profile,
		})
	}
	logEntry.Info("Successfully retrieved followers for user ")
	return &pb.FollowersResponse{
		Users: userdetails,
	}, nil

}

func (ad *AuthServer) Followings(ctx context.Context, req *pb.FollowingRequest) (*pb.FollowingResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "Followings")
	logEntry.Info("Processing Followings request for user ID:", req.GetUserID())
	id := req.UserID

	details, err := ad.authUseCase.Followings(int(id))
	if err != nil {
		logEntry.WithError(err).Error("Error getting followings")
		return &pb.FollowingResponse{}, err
	}
	var userdetails []*pb.FollowResponse

	for _, user := range details {
		userdetails = append(userdetails, &pb.FollowResponse{
			Username:    user.Username,
			UserProfile: user.Profile,
		})
	}
	logEntry.Info("Successfully retrieved followings for user ")
	return &pb.FollowingResponse{
		Users: userdetails,
	}, nil
}

func (ad *AuthServer) SendOTP(ctx context.Context, req *pb.SendOTPRequest) (*pb.SendOTPResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "SendOTP")
	logEntry.Info("Processing SendOTP request for phone (masked)")
	phone := req.Phone

	err := ad.authUseCase.SendOTP(phone)
	if err != nil {
		logEntry.WithError(err).Error("Error sending OTP")
		return &pb.SendOTPResponse{}, err
	}
	logEntry.Info("OTP sent successfully for phone (masked)")
	return &pb.SendOTPResponse{}, nil
}

func (ad *AuthServer) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "VerifyOTP")
	logEntry.Info("Processing VerifyOTP request for phone (masked), code", req.GetCode())
	udetails := models.VerifyData{
		PhoneNumber: req.Phone,
		Code:        req.Code,
	}
	details, err := ad.authUseCase.VerifyOTP(udetails)
	if err != nil {
		logEntry.WithError(err).Error("Error verifying OTP")
		return &pb.VerifyOTPResponse{}, err
	}
	logEntry.Info("OTP verified successfully for phone (masked)")
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

func (ad *AuthServer) UploadProfilepic(ctx context.Context, req *pb.UploadProfilepicRequest) (*pb.UploadProfilepicResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "UploadProfilepic")
	logEntry.Info("Processing UploadProfilepic request for user ID:", req.GetUserID())
	id := req.UserID
	file := req.File

	err := ad.authUseCase.UploadProfilepic(int(id), file)
	if err != nil {
		logEntry.WithError(err).Error("Error uploading profile picture")
		return &pb.UploadProfilepicResponse{}, err
	}
	logEntry.Info("Profile picture uploaded successfully for user")
	return &pb.UploadProfilepicResponse{}, nil
}

func (ad *AuthServer) DoesUserExist(ctx context.Context, req *pb.DoesUserExistRequest) (*pb.DoesUserExistResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "DoesUserExist")
	logEntry.Info("Processing DoesUserExist request for ID:", req.GetId())
	id := req.Id
	fmt.Println("iddddd is ", id)

	res := ad.authUseCase.DoesUserExist(id)
	if !res {
		logEntry.Info("User with ID (masked) does not exist")
		return &pb.DoesUserExistResponse{}, errors.New("user doesnt exist")
	}
	logEntry.Info("User with IDexists")
	return &pb.DoesUserExistResponse{
		Data: true,
	}, nil
}

func (ad *AuthServer) FindUserName(ctx context.Context, req *pb.FindUserNameRequest) (*pb.FindUserNameResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "FindUserName")
	logEntry.Info("Processing FindUserName request for ID:", req.GetId())

	id := req.Id

	res, err := ad.authUseCase.FindUserName(id)
	if err != nil {
		logEntry.WithError(err).Error("Error finding username")
		return &pb.FindUserNameResponse{}, err
	}
	logEntry.Info("Successfully found username for ID", req.Id)
	return &pb.FindUserNameResponse{
		Name: res,
	}, nil
}

func (ad *AuthServer) UserData(ctx context.Context, req *pb.UserDataRequest) (*pb.UserDataResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "UserData")
	logEntry.Info("Processing UserData request for user ID:", req.GetUserid())
	id := req.Userid
	res, err := ad.authUseCase.UserData(int(id))

	if err != nil {
		logEntry.WithError(err).Error("Error getting user data")
		return &pb.UserDataResponse{}, err
	}
	logEntry.Info("Successfully retrieved user data for ID", req.Userid)
	return &pb.UserDataResponse{
		Userid:   int64(res.ID),
		Username: res.Username,
		Profile:  res.Profile,
	}, nil
}
