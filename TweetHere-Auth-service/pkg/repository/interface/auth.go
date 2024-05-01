package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
)

type AuthRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistByEmail(email string) (*domain.Admin, error)
	UpdateBlockUserByID(user domain.User)error
	GetUser(page int)([]models.UserDetails,error)
	GetUserById(id string)(domain.User,error)
	
    
	UserSignUp(userdetails models.UserSignup)(models.UserDetailsResponse,error)
	ChekUserExistByEmail(email string)(*domain.User,error)
	FindUserByEmail(user models.UserLogin)(models.UserSignup,error)
	UserUpdateProfile(user models.UserProfile,id int)(models.UserProfileResponse,error)
	GetPassword(id int) (string, error)
	ChangePassword(id int, password string) error
	
    	
}
