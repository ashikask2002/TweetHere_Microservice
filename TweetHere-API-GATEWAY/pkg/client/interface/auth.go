package interfaces

import "TweetHere-API/pkg/utils/models"

type AdminClient interface {
	AdminSignUp(admindetails models.AdminSignup) (models.TokenAdmin, error)
	AdminLogin(admindetails models.AdminLogin) (models.TokenAdmin, error)

	UserSignup(userdetails models.UserSignup) (models.TokenUser,error)
	UserLogin(userdetails models.UserLogin) (models.TokenUser,error)
	UserUpdateProfile(userdetails models.UserProfile) (models.UserProfileResponse)
}
