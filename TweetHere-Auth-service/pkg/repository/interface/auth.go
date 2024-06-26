package interfaces

import (
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/utils/models"
	"time"
)

type AuthRepository interface {
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistByEmail(email string) (*domain.Admin, error)
	UpdateBlockUserByID(user domain.User) error
	GetUser(page int) ([]models.UserDetails, error)
	GetUserById(id string) (domain.User, error)

	UserSignUp(userdetails models.UserSignup) (models.UserDetailsResponse, error)
	ChekUserExistByEmail(email string) (*domain.User, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignup, error)
	UserUpdateProfile(user models.UserProfile, id int) (models.UserProfileResponse, error)
	GetPassword(id int) (string, error)
	ChangePassword(id int, password string) error
	GetUserDetails(id int) ([]models.UserDetails4user, error)
	DeleteRecentOtpRequestsBefore5min() error
	TemporarySavingUserOtp(otp int, userEmail string, expiration time.Time) error
	VerifyOTP(email, otp string) (bool, error)
	CheckUserAvailability(userid int) bool
	ExistFollowreq(id int, userid int) bool
	ExistFollowers(userid, followingUserID int) bool
	FollowReq(userid int, followinguserid int) error
	CheckRequest(userid int, followinguserid int) bool
	AlreadyAccepted(userid int, followinguserid int) bool
	AcceptFollowREQ(userID, FollowingUserID int) error
	UnFollow(userID, UnFollowUserID int) error
	Followers(userID int) ([]models.FollowResp, error)
	Followdetails(userid int) (models.Followersresponse, error)
	Followings(userID int) ([]models.FollowResp, error)
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error)
	UploadProfilepic(id int,url string)error
	GetUserName(id int)(string,error)
	UserData(id int)(models.UserData,error)
	UpdateFirstName(firstname string, userID int) error 
	UpdateLastName(lastname string, userID int) error
	UpdateUserName(username string, userID int) error
	UpdateDOB(dob string, userID int) error
	UpdateUserEmail(email string, userID int) error
	UpdateBIO(bio string, userID int) error
	CheckEmail(email string) bool
	UserDetails(userID int) (models.UserProfileResponse, error)
	GetFollowingUsers(userID int) ([]models.FollowUsers, error)
}
