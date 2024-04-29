package repository

import (
	"Tweethere-Auth/pkg/domain"
	interfaces "Tweethere-Auth/pkg/repository/interface"
	"Tweethere-Auth/pkg/utils/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) interfaces.AuthRepository {
	return &authRepository{
		DB: DB,
	}
}

func (ad *authRepository) AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var model models.AdminDetailsResponse

	if err := ad.DB.Raw("INSERT INTO admins (firstname,lastname,email,password) VALUES (?,?,?,?) RETURNING id,firstname,lastname,email", adminDetails.Firstname, adminDetails.Lastname, adminDetails.Email, adminDetails.Password).Scan(&model).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}
	fmt.Println("modelsssssssssssss", model)
	return model, nil
}

func (ad *authRepository) CheckAdminExistByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	res := ad.DB.Where(&domain.Admin{Email: email}).First(&admin)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Admin{}, res.Error
	}
	return &admin, nil
}

func (ad *authRepository) FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error) {
	var user models.AdminSignUp
	err := ad.DB.Raw("SELECT * FROM admins WHERE email = ?", admin.Email).Scan(&user).Error
	if err != nil {
		return models.AdminSignUp{}, errors.New("error checking admin details")
	}
	return user, nil
}

func (ad *authRepository) UserSignUp(userDetails models.UserSignup) (models.UserDetailsResponse, error) {
    var model models.UserDetailsResponse

    // Set IsBlocked to false by default
    userDetails.IsBlocked = false

    if err := ad.DB.Raw("INSERT INTO users (firstname, lastname, username, phone, email, date_of_birth, password, is_blocked) VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, firstname, lastname, email", userDetails.Firstname, userDetails.Lastname, userDetails.Username, userDetails.Phone, userDetails.Email, userDetails.DateOfBirth, userDetails.Password, userDetails.IsBlocked).Scan(&model).Error; err != nil {
        return models.UserDetailsResponse{}, err
    }
    return model, nil
}

func (ad *authRepository) ChekUserExistByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := ad.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.User{}, res.Error
	}
	return &user, nil
}

func (ad *authRepository) FindUserByEmail(user models.UserLogin) (models.UserSignup, error) {
	var userr models.UserSignup
	err := ad.DB.Raw("select * from users where email = ?", user.Email).Scan(&userr).Error
	if err != nil {
		return models.UserSignup{}, errors.New("error checking user details")
	}
	return userr, nil
}
