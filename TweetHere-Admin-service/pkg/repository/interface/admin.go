package interfaces

import (
	"TweetHere-admin/pkg/domain"
	"TweetHere-admin/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistByEmail(email string) (*domain.Admin, error)
}
