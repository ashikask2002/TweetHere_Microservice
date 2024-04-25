package usecase

import (
	"TweetHere-admin/pkg/domain"
	"TweetHere-admin/pkg/helper"
	interfaces "TweetHere-admin/pkg/repository/interface"
	services "TweetHere-admin/pkg/usecase/interface"
	"TweetHere-admin/pkg/utils/models"
	"errors"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func newAdminUseCase(repository interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repository,
	}
}

func (ad *adminUseCase) AdminSignUp(admin models.AdminSignUp) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenAdmin{}, errors.New("user with this email already exist")
	}
	hashedpassword, err := helper.PasswordHash(admin.Password)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error in hashing password")
	}
	admin.Password = hashedpassword
	admindata, err := ad.adminRepository.AdminSignUp(admin)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("could not added the user")

	}
	tokenString, err := helper.GenerateTokenAdmin(admindata)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: admindata,
		Token: tokenString,
	}, nil
}

func (ad *adminUseCase) LoginHandler(admin models.AdminLogin) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenAdmin{}, errors.New("this email doesnt exist")
	}

	admindetails, err := ad.adminRepository.FindAdminByEmail(admin)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admindetails.Password), []byte(admin.Password))
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("password not matching")
	}

	var AdminDetailsResponse models.AdminDetailsResponse

	err = copier.Copy(&AdminDetailsResponse, &admindetails)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	tokenstring, err := helper.GenerateTokenAdmin(AdminDetailsResponse)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}
	return &domain.TokenAdmin{
		Admin: AdminDetailsResponse,
		Token: tokenstring,
	}, nil
}
