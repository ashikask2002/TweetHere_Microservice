package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
)


type AdminUseCase interface{
	AdminSignUp(admindetails models.AdminSignUp) (*domain.TokenAdmin,error)
	LoginHandler(admindetails models.AdminLogin) (*domain.TokenAdmin,error)
	BlockUser(id string)error
	UnBlockUser(id string)error
	GetUser(page int)([]models.UserDetails,error)

	UserSignup(userdetails models.UserSignup) (*domain.TokenUser,error)
	UserLogin(userdetails models.UserLogin) (*domain.TokenUser,error)
     
	UserUpdateProfile(user models.UserProfile,id int)(models.UserProfileResponse,error)
	ChangePassword(id int,old string,new string,re string)error
   


}