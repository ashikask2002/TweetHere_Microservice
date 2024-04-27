package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistByEmail(email string) (*domain.Admin, error)
}
