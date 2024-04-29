package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
)

type AuthRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistByEmail(email string) (*domain.Admin, error)
    
	UserSignUp(userdetails models.UserSignup)(models.UserDetailsResponse,error)
	ChekUserExistByEmail(email string)(*domain.User,error)
	FindUserByEmail(user models.UserLogin)(models.UserSignup,error)
    	
}
