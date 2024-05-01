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
	ChangePassword(id int,old string,new string,re string)error

}
