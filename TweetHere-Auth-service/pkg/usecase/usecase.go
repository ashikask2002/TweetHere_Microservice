package usecase

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/helper"
	interfaces "Tweethere-Auth/pkg/repository/interface"
	services "Tweethere-Auth/pkg/usecase/interface"
	"Tweethere-Auth/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	authRepository interfaces.AuthRepository
}

func NewAuthUseCase(repository interfaces.AuthRepository) services.AdminUseCase {
	return &authUseCase{
		authRepository: repository,
	}
}

func (ad *authUseCase) AdminSignUp(admin models.AdminSignUp) (*domain.TokenAdmin, error) {
	email, err := ad.authRepository.CheckAdminExistByEmail(admin.Email)
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
	admindata, err := ad.authRepository.AdminSignUp(admin)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("could not added the user")

	}
	tokenString, err := helper.GenerateTokenAdmin(admindata)
	fmt.Println("token errrr is", err)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: admindata,
		Token: tokenString,
	}, nil
}

func (ad *authUseCase) LoginHandler(admin models.AdminLogin) (*domain.TokenAdmin, error) {
	email, err := ad.authRepository.CheckAdminExistByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenAdmin{}, errors.New("this email doesnt exist")
	}

	admindetails, err := ad.authRepository.FindAdminByEmail(admin)
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

func (ad *authUseCase) UserSignup(user models.UserSignup) (*domain.TokenUser, error) {
	email, err := ad.authRepository.ChekUserExistByEmail(user.Email)
	if err != nil {
		return &domain.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenUser{}, errors.New("user with this email already exist")
	}
	hashedpassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return &domain.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashedpassword
	userdata, err := ad.authRepository.UserSignUp(user)
	if err != nil {
		return &domain.TokenUser{}, errors.New("could not added the uesr data")
	}
	accessToken, refreshToken, err := helper.GenerateTokenUser(userdata)
	if err != nil {
		return &domain.TokenUser{}, err
	}
	return &domain.TokenUser{
		User:         userdata,
		AccesToken:   accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (ad *authUseCase) UserLogin(user models.UserLogin) (*domain.TokenUser, error) {
	fmt.Println("ssssssssssssss", user)
	email, err := ad.authRepository.ChekUserExistByEmail(user.Email)
	if err != nil {
		return &domain.TokenUser{}, errors.New("error in email checking part")
	}
	if email == nil {
		return &domain.TokenUser{}, errors.New("this user doesnt exist")
	}

	userdetails, err := ad.authRepository.FindUserByEmail(user)

	if err != nil {
		return &domain.TokenUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userdetails.Password), []byte(user.Password))
	if err != nil {
		return &domain.TokenUser{}, errors.New("password is not matching")
	}
	var userdetailsresponse models.UserDetailsResponse
	err = copier.Copy(&userdetailsresponse, userdetails)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	accessToken, refreshToken, err := helper.GenerateTokenUser(userdetailsresponse)

	if err != nil {
		return &domain.TokenUser{}, errors.New("error in genereate toeknuser")
	}
	fmt.Println("osssssssssssssss", userdetails)
	fmt.Println("userrrrrrrr", userdetailsresponse)
	return &domain.TokenUser{
		User:         userdetailsresponse,
		AccesToken:   accessToken,
		RefreshToken: refreshToken,
	}, nil
}
