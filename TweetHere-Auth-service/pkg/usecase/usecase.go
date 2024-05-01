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

func (ad *authUseCase) UserUpdateProfile(user models.UserProfile, id int) (models.UserProfileResponse, error) {
	fmt.Println("userdetaillllllll at usecase ", user, id)

	userdetails, err := ad.authRepository.UserUpdateProfile(user, id)

	if err != nil {
		return models.UserProfileResponse{}, errors.New("error happened while profileupdate")
	}
	return userdetails, nil
}

func (ad *authUseCase) GetUser(page int) ([]models.UserDetails, error) {
	userdetails, err := ad.authRepository.GetUser(page)
	if err != nil {
		return []models.UserDetails{}, err
	}
	return userdetails, nil
}

func (ad *authUseCase) BlockUser(id string) error {
	user, err := ad.authRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if user.IsBlocked {
		return errors.New("user already blocked")
	} else {
		user.IsBlocked = true
	}
	err = ad.authRepository.UpdateBlockUserByID(user)
	if err != nil {
		return errors.New("failed to block")
	}
	return nil

}

func (ad *authUseCase) UnBlockUser(id string) error {
	user, err := ad.authRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if !user.IsBlocked {
		return errors.New("user already unblocked")
	} else {
		user.IsBlocked = false
	}
	err = ad.authRepository.UpdateBlockUserByID(user)
	if err != nil {
		return errors.New("failed to unblock")
	}
	return nil

}

func (ad *authUseCase) ChangePassword(id int, old string, new string, re string) error {
	userpassword, err := ad.authRepository.GetPassword(id)
	if err != nil {
		errors.New("internal error")
	}

	err = helper.CompareHashAndPassword(userpassword, old)
	if err != nil {
		return errors.New("old password incorrect")
	}
	if new != re {
		return errors.New("passwords are not matching")
	}
	newpassword, errr := helper.PasswordHash(new)
	if errr != nil {
		return errors.New("error in hashig password")
	}

	return ad.authRepository.ChangePassword(id, newpassword)
}
