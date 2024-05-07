package interfaces

import "TweetHere-API/pkg/utils/models"

type AdminClient interface {
	AdminSignUp(admindetails models.AdminSignup) (models.TokenAdmin, error)
	AdminLogin(admindetails models.AdminLogin) (models.TokenAdmin, error)
    BlockUser(id string)error
	UnBlockUser(id string)error
	GetUser(page int) ([]models.UserDetails,error)



	UserSignup(userdetails models.UserSignup) (models.TokenUser,error)
	UserLogin(userdetails models.UserLogin) (models.TokenUser,error)
	UserUpdateProfile(userdetails models.UserProfile,id int) (models.UserProfileResponse,error)
	ChangePassword(id int,passworddetails models.ChangePassword)error
	GetUserDetails(id int)( []models.UserDetails4user,error)
	UserOTPLogin(email string) (string, error)
	OtpVerification(email, otp string) (bool, error)
	FollowReq(id int,userid int)error
	AcceptFollowreq(id int,userid int)error
	Unfollow(id int,userid int)error
	Followers(id int)([]models.Followersresponse,error)
	Followings(id int)([]models.Followersresponse,error)

}
