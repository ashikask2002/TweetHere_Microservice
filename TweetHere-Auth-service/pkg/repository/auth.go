package repository

import (
	"Tweethere-Auth/pkg/domain"
	interfaces "Tweethere-Auth/pkg/repository/interface"
	"Tweethere-Auth/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

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


func (ur *authRepository) UserDetails(userID int) (models.UserProfileResponse, error) {
	var userDetails models.UserProfileResponse
	err := ur.DB.Raw("SELECT firstname, lastname, username,phone,email,date_of_birth,profile,bio  FROM users WHERE id = ?", userID).Row().Scan(&userDetails.Firstname, &userDetails.Lastname, &userDetails.Username, &userDetails.Phone, &userDetails.Email, &userDetails.DateOfBirth, &userDetails.Profile, &userDetails.Bio)
	if err != nil {
		fmt.Println("Error retrieving user details:", err)
		return models.UserProfileResponse{}, err
	}
	return userDetails, nil
}


func (ur *authRepository) UpdateFirstName(firstname string, userID int) error {
	err := ur.DB.Exec("UPDATE users SET firstname= ? WHERE id = ?", firstname, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) UpdateLastName(lastname string, userID int) error {
	err := ur.DB.Exec("UPDATE users SET lastname= ? WHERE id = ?", lastname, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) UpdateUserName(username string, userID int) error {
	err := ur.DB.Exec("UPDATE users SET username= ? WHERE id = ?", username, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) UpdateDOB(dob string, userID int) error {
	err := ur.DB.Exec("UPDATE users SET date_of_birth = ? WHERE id = ?", dob, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) UpdateUserEmail(email string, userID int) error {
	err := ur.DB.Exec("UPDATE users SET email= ? WHERE id = ?", email, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) UpdateBIO(bio string, userID int) error {
	err := ur.DB.Exec("UPDATE users SET bio= ? WHERE id = ?", bio, userID).Error
	if err != nil {
		return err
	}
	return nil
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
	fmt.Println("am in repo change get password")
	var userPassword string
	err := ad.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	fmt.Println("njaan thanne", id, userPassword)
	return userPassword, nil
}

func (c *authRepository) ChangePassword(id int, password string) error {
	fmt.Println("i am in repository")
	err := c.DB.Exec("UPDATE users SET password = $1 WHERE id = $2", password, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *authRepository) GetUserDetails(id int) ([]models.UserDetails4user, error) {
	var userDetails []models.UserDetails4user

	// Execute query to fetch user details
	if err := c.DB.Table("users").Select("id, firstname, lastname, username, phone, email, date_of_birth, profile, bio").Where("id = ?", id).Find(&userDetails).Error; err != nil {
		return nil, err
	}

	return userDetails, nil
}

func (ur *authRepository) DeleteRecentOtpRequestsBefore5min() error {
	query := "DELETE FROM user_otp_logins WHERE expiration < CURRENT_TIMESTAMP - INTERVAL '5 minutes';"
	err := ur.DB.Exec(query).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) TemporarySavingUserOtp(otp int, userEmail string, expiration time.Time) error {

	query := `INSERT INTO user_otp_logins (email, otp, expiration) VALUES ($1, $2, $3)`
	err := ur.DB.Exec(query, userEmail, otp, expiration).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *authRepository) VerifyOTP(email, otp string) (bool, error) {
	query := "SELECT COUNT(*) FROM user_otp_logins WHERE email = ? AND otp = ? AND expiration > CURRENT_TIMESTAMP;"
	var count int64
	if err := ur.DB.Raw(query, email, otp).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ad *authRepository) CheckUserAvailability(userid int) bool {
	var count int
	if err := ad.DB.Raw("SELECT COUNT(*) FROM users WHERE id=?", userid).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}
func (ad *authRepository) CheckEmail(email string) bool {
	var count int
	if err := ad.DB.Raw("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (ad *authRepository) ExistFollowreq(userid, followingUserID int) bool {
	var count int
	err := ad.DB.Raw("SELECT COUNT(*) FROM following_requests WHERE following_user= ? AND user_id = ? ", userid, followingUserID).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (ad *authRepository) ExistFollowers(userid, followingUserID int) bool {
	var count int
	err := ad.DB.Raw("SELECT COUNT(*) FROM followings WHERE following_user= ? AND user_id = ? ", followingUserID, userid).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (ad *authRepository) FollowReq(userid int, followinguserid int) error {
	err := ad.DB.Exec("INSERT INTO following_requests (user_id,following_user,created_at)VALUES(?,?,NOW())", userid, followinguserid).Error
	if err != nil {
		return err
	}
	err = ad.DB.Exec("INSERT INTO followings (user_id,following_user,created_at)VALUES(?,?,NOW())", userid, followinguserid).Error
	if err != nil {
		return err
	}

	return nil
}

func (ad *authRepository) CheckRequest(userid int, followinguserid int) bool {
	var request models.FollowingRequest
	err := ad.DB.Raw("SELECT following_user, user_id FROM following_requests WHERE following_user = ? AND user_id = ?", userid, followinguserid).Scan(&request).Error
	if err != nil {
		return false
	}
	return request.UserID != 0
}

func (ad *authRepository) AlreadyAccepted(userid int, followinguserid int) bool {
	var count int
	err := ad.DB.Raw("SELECT COUNT(*) FROM followers WHERE user_id = ? AND following_user = ?", userid, followinguserid).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (ur *authRepository) AcceptFollowREQ(userID, FollowingUserID int) error {
	err := ur.DB.Exec("INSERT INTO followers (user_id,following_user,created_at) VALUES(?,?,NOW())", userID, FollowingUserID).Error
	if err != nil {
		return err
	}

	err = ur.DB.Exec("DELETE FROM following_requests WHERE user_id = ? AND following_user = ?", FollowingUserID, userID).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *authRepository) UnFollow(userID, UnFollowUserID int) error {
	err := ur.DB.Exec("DELETE FROM followings WHERE user_id = ? AND following_user = ?", userID, UnFollowUserID).Error
	if err != nil {
		return err
	}
	errs := ur.DB.Exec("DELETE FROM followers WHERE user_id = ? AND following_user = ?", UnFollowUserID, userID).Error
	if errs != nil {
		return err
	}
	return nil
}

func (ur *authRepository) Followers(userID int) ([]models.FollowResp, error) {
	var response []models.FollowResp
	err := ur.DB.Raw("SELECT following_user FROM followers WHERE user_id = ?", userID).Scan(&response).Error
	if err != nil {
		return []models.FollowResp{}, err
	}
	return response, nil
}

func (ur *authRepository) Followdetails(userid int) (models.Followersresponse, error) {
	var response models.Followersresponse
	err := ur.DB.Raw("SELECT username,profile FROM users WHERE id = ?", userid).Scan(&response).Error
	if err != nil {
		return models.Followersresponse{}, err
	}
	return response, nil
}

func (ur *authRepository) Followings(userID int) ([]models.FollowResp, error) {
	var response []models.FollowResp
	err := ur.DB.Raw("SELECT following_user FROM followings WHERE user_id = ?", userID).Scan(&response).Error
	if err != nil {
		return []models.FollowResp{}, err
	}
	return response, nil
}

func (ot *authRepository) FindUserByMobileNumber(phone string) bool {

	var count int
	if err := ot.DB.Raw("select count(*) from users where phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (ot *authRepository) UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	if err := ot.DB.Raw("select * from users where phone = ?", phone).Scan(&userDetails).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userDetails, nil

}

func (ot *authRepository) UploadProfilepic(id int, url string) error {
	// Execute the SQL query to update the profile URL
	err := ot.DB.Exec("UPDATE users SET profile = ? WHERE id = ?", url, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (ot *authRepository) GetUserName(id int) (string, error) {
	var name string
	err := ot.DB.Raw("SELECT username FROM users WHERE id = ?", id).Scan(&name).Error
	if err != nil {
		return "", err
	}
	return name, nil
}

func (at *authRepository) UserData(id int) (models.UserData, error) {
	var details models.UserData
	err := at.DB.Raw("SELECT id,username,profile FROM users WHERE id = ?", id).Scan(&details).Error
	if err != nil {
		return models.UserData{}, errors.New("error in getting user data")
	}
	return details, nil
}
