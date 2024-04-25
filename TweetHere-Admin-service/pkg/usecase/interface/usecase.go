package interfaces

import (
	"TweetHere-admin/pkg/domain"
	"TweetHere-admin/pkg/utils/models"
)


type AdminUseCase interface{
	AdminSignUp(admindetails models.AdminSignUp) (*domain.TokenAdmin,error)
	LoginHandler(admindetails models.AdminLogin) (*domain.TokenAdmin,error)
}