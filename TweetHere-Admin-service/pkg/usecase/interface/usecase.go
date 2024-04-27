package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
)


type AdminUseCase interface{
	AdminSignUp(admindetails models.AdminSignUp) (*domain.TokenAdmin,error)
	LoginHandler(admindetails models.AdminLogin) (*domain.TokenAdmin,error)
}