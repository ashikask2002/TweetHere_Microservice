package repository

import (
	"Tweethere-Auth/pkg/domain"
	interfaces "Tweethere-Auth/pkg/repository/interface"
	"Tweethere-Auth/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"

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

func (ad *authRepository) UserUpdateProfile(user models.UserProfile, id int) (models.UserProfileResponse, error) {
	var userdetails models.UserProfileResponse

	fmt.Println("user detailssssssssssss at repository", user, id)

	// Build the query to update the user profile details
	query := ad.DB.Table("users").Where("id = ?", id).Updates(map[string]interface{}{
		"firstname":     user.Firstname,
		"lastname":      user.Lastname,
		"username":      user.Username,
		"phone":         user.Phone,
		"email":         user.Email,
		"date_of_birth": user.DateOfBirth,
		"profile":       user.Profile,
		"bio":           user.Bio,
	})

	// Check for errors in executing the query
	if query.Error != nil {
		return models.UserProfileResponse{}, query.Error
	}

	// Fetch the updated user profile details
	if err := ad.DB.Table("users").Where("id = ?", id).First(&userdetails).Error; err != nil {
		return models.UserProfileResponse{}, err
	}

	return userdetails, nil
}


func (ad *authRepository) UpdateBlockUserByID(user domain.User) error {
	err := ad.DB.Exec("update users set is_blocked = ? where id = ?", user.IsBlocked, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ad *authRepository) GetUser(page int) ([]models.UserDetails, error) {
    if page == 0 {
        page = 1
    }
    offset := (page - 1) * 20
    var userDetails []models.UserDetails

    query := `
        SELECT id, firstname, lastname, username, phone, email, date_of_birth, is_blocked, profile, bio
        FROM users
        LIMIT ? OFFSET ?
    `
    if err := ad.DB.Raw(query, 20, offset).Scan(&userDetails).Error; err != nil {
        return []models.UserDetails{}, err
    }

    return userDetails, nil
}

func (ad *authRepository) GetUserById(id string) (domain.User, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?", user_id).Scan(&count).Error; err != nil {
		return domain.User{}, err
	}
	if count < 1 {
		return domain.User{}, errors.New("user for the given id does not exist")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.User

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.User{}, err
	}

	return userDetails, nil
}

func (ad *authRepository) GetPassword(id int) (string, error) {
	var userPassword string
	err := ad.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil
}

func (c *authRepository) ChangePassword(id int, password string) error {
	err := c.DB.Exec("UPDATE users SET password = $1 WHERE id = $2", password, id).Error
	if err != nil {
		return err
	}
	return nil
}