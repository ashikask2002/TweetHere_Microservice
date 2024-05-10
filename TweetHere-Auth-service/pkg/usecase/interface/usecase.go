package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(admindetails models.AdminLogin) (*domain.TokenAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
	GetUser(page int) ([]models.UserDetails, error)

	UserSignup(userdetails models.UserSignup) (*domain.TokenUser, error)
	UserLogin(userdetails models.UserLogin) (*domain.TokenUser, error)

	UserUpdateProfile(user models.UserProfile, id int) (models.UserProfileResponse, error)
	ChangePassword(id int, pasworddetails models.ChangePassword) error
	GetUserDetails(id int) ([]models.UserDetails4user, error)
	UserOTPLogin(email string) (string, error)
	OtpVerification(email, otp string) (bool, error)
	FollowReq(id, userid int) error
	AcceptFollowReq(id, userid int) error
	Unfollow(id, userid int) error
	Followers(id int) ([]models.Followersresponse, error)
	Followings(id int) ([]models.Followersresponse, error)
	SendOTP(phone string) error
	VerifyOTP(udetails models.VerifyData)(*domain.TokenUser,error)
}
