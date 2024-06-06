package usecase

import (
	"Tweethere-Auth/pkg/config"
	"Tweethere-Auth/pkg/domain"
	"Tweethere-Auth/pkg/helper"
	"Tweethere-Auth/pkg/randomnumbergenerator"
	interfaces "Tweethere-Auth/pkg/repository/interface"
	services "Tweethere-Auth/pkg/usecase/interface"
	"Tweethere-Auth/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	cfg            config.Config
	authRepository interfaces.AuthRepository
}

func NewAuthUseCase(cfg config.Config, repository interfaces.AuthRepository) services.AdminUseCase {
	return &authUseCase{
		cfg:            cfg,
		authRepository: repository,
	}
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
	emaill := helper.IsValidEmail(user.Email)
	if !emaill {
		return &domain.TokenUser{}, errors.New("email not in correct format")
	}

	phonee := helper.PhoneValidation(user.Phone)
	if !phonee {
		return &domain.TokenUser{}, errors.New("phone number are not in correct format")
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
	emaill := helper.IsValidEmail(user.Email)
	if !emaill {
		return &domain.TokenUser{}, errors.New("email not in correct format")
	}
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
	return &domain.TokenUser{
		User:         userdetailsresponse,
		AccesToken:   accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (ad *authUseCase) UserUpdateProfile(user models.UserProfile, id int) (models.UserProfileResponse, error) {
	fmt.Println("userrnnnnnssss", id)

	// phnn := helper.PhoneValidation(user.Phone)
	// if !phnn {
	// 	return models.UserProfileResponse{}, errors.New("phone not in correct format")
	// }

	// userdetails, err := ad.authRepository.UserUpdateProfile(user, id)

	// if err != nil {
	// 	return models.UserProfileResponse{}, errors.New("error happened while profileupdate")
	// }

	if user.Firstname != "" {
		ad.authRepository.UpdateFirstName(user.Firstname, id)
	}
	if user.Lastname != "" {
		ad.authRepository.UpdateLastName(user.Lastname, id)
	}
	if user.Username != "" {
		ad.authRepository.UpdateUserName(user.Username, id)

	}
	fmt.Println("222222222")
	if user.Email != "" {
		fmt.Println("33333333")
		ok := ad.authRepository.CheckEmail(user.Email)
		if ok {
			return models.UserProfileResponse{}, errors.New("email already in use")
		}
		emaill := helper.IsValidEmail(user.Email)
		if !emaill {
			return models.UserProfileResponse{}, errors.New("email not in correct format")
		}
		ad.authRepository.UpdateUserEmail(user.Email, id)
	}
	fmt.Println("4444444444")
	if user.DateOfBirth != "" {
		ad.authRepository.UpdateDOB(user.DateOfBirth, id)
	}
	fmt.Println("555555555")

	if user.Bio != "" {
		ad.authRepository.UpdateBIO(user.Bio, id)
	}
	fmt.Println("666666666")
	return ad.authRepository.UserDetails(id)
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
	fmt.Println("aaaaaa")
	if err != nil {
		return err
	}
	if user.IsBlocked {
		fmt.Println("bbbbb")
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

// func (ad *authUseCase) ChangePassword(id int, passworddetails models.ChangePassword) error {

// 	userpassword, err := ad.authRepository.GetPassword(id)
// 	if err != nil {
// 		return errors.New("internal error")
// 	}
// 	fmt.Println("userpass", userpassword)
// 	fmt.Println("000000000000")

// 	oldhash, errrs := helper.PasswordHash(passworddetails.Oldpassword)
// 	if errrs != nil {
// 		return errrs
// 	}
// 	fmt.Println("oldhashed password", oldhash)

// 	errs := helper.CompareHashAndPassword(userpassword, passworddetails.Oldpassword)
// 	if errs != nil {
// 		return errs
// 	}
// 	fmt.Println("1111111111")
// 	if passworddetails.NewPassword != passworddetails.RePassword {
// 		return errors.New("passwords are not matching")
// 	}
// 	fmt.Println("22222222222")
// 	newpassword, errr := helper.PasswordHash(passworddetails.NewPassword)
// 	if errr != nil {
// 		return errors.New("error in hashig password")
// 	}
// 	fmt.Println("33333333333")

// 	return ad.authRepository.ChangePassword(id, newpassword)
// }

func (ur *authUseCase) ChangePassword(id int, change models.ChangePassword) error {
	userPassword, err := ur.authRepository.GetPassword(id)
	if err != nil {
		return fmt.Errorf("internal error")
	}
	err = helper.CompareHashAndPassword(userPassword, change.Oldpassword)
	if err != nil {
		return fmt.Errorf("password incorrect")
	}
	if change.NewPassword != change.RePassword {
		return fmt.Errorf("password doesnt match")
	}
	newpassword, err := helper.PasswordHash(change.NewPassword)
	if err != nil {
		return fmt.Errorf("error in hashing password")
	}
	return ur.authRepository.ChangePassword(id, string(newpassword))
}

func (ad *authUseCase) GetUserDetails(id int) ([]models.UserDetails4user, error) {
	userdetails, err := ad.authRepository.GetUserDetails(id)
	if err != nil {
		return []models.UserDetails4user{}, err
	}
	return userdetails, nil
}

func (r *authUseCase) UserOTPLogin(email string) (string, error) {
	otp := randomnumbergenerator.RandomNumber()

	otpString := strconv.Itoa(otp)

	errRemv := r.authRepository.DeleteRecentOtpRequestsBefore5min()
	if errRemv != nil {
		return "", errRemv
	}

	expiration := time.Now().Add(5 * time.Minute)
	errTempSave := r.authRepository.TemporarySavingUserOtp(otp, email, expiration)
	if errTempSave != nil {
		fmt.Println("Can't save temporary data for OTP verification in DB")
		return "", errors.New("OTP verification down, please try again later")
	}
	// name, err := r.userRepository.GetUserName(email)
	// if err != nil {
	// 	return "", err
	// }
	// err = helper.SendVerificationEmailWithOtp(otp, email, name)
	// if err != nil {
	// 	return "",err
	// }

	return otpString, nil
}

func (r *authUseCase) OtpVerification(email, otp string) (bool, error) {
	verified, err := r.authRepository.VerifyOTP(email, otp)
	if err != nil {
		return false, err
	}
	return verified, nil
}

func (ad *authUseCase) FollowReq(id int, userid int) error {
	fmt.Println("111111")
	userExist := ad.authRepository.CheckUserAvailability(id)
	if !userExist {
		return errors.New("user doesnt exist")
	}
	fmt.Println("2222222")
	followuserExist := ad.authRepository.CheckUserAvailability(userid)
	if !followuserExist {
		return errors.New("user doesnt exist")
	}
	fmt.Println("3333333")
	err := ad.authRepository.ExistFollowers(id, userid)
	if err {
		fmt.Println("5555")
		return errors.New("request already exist")
	}
	fmt.Println("44444")
	errs := ad.authRepository.FollowReq(id, userid)
	if errs != nil {
		return errs
	}
	return nil

}

func (ad *authUseCase) AcceptFollowReq(id int, userid int) error {
	userExist := ad.authRepository.CheckUserAvailability(id)
	if !userExist {
		return errors.New("user doesnt exist")
	}
	followuserExist := ad.authRepository.CheckUserAvailability(userid)
	if !followuserExist {
		return errors.New("user doesnt exist")
	}
	req := ad.authRepository.CheckRequest(id, userid)
	if !req {
		return errors.New("no request available")
	}
	alreadyfollow := ad.authRepository.AlreadyAccepted(id, userid)
	if alreadyfollow {
		return errors.New("already exist")
	}

	err := ad.authRepository.AcceptFollowREQ(id, userid)
	if err != nil {
		return err
	}
	return nil

}

func (ad *authUseCase) Unfollow(id int, userid int) error {
	userExist := ad.authRepository.CheckUserAvailability(id)
	if !userExist {
		return errors.New("user doesnt exist")
	}
	follouserExist := ad.authRepository.CheckUserAvailability(userid)
	if !follouserExist {
		return errors.New("user doesnt exist")
	}
	err := ad.authRepository.UnFollow(id, userid)
	if err != nil {
		return err
	}
	return nil
}

func (ad *authUseCase) Followers(id int) ([]models.Followersresponse, error) {
	userid := ad.authRepository.CheckUserAvailability(id)
	if !userid {
		return []models.Followersresponse{}, errors.New("user doenst exist")
	}
	ids, err := ad.authRepository.Followers(id)
	if err != nil {
		return []models.Followersresponse{}, err
	}
	var userresp []models.Followersresponse

	for _, ud := range ids {
		details, err := ad.authRepository.Followdetails(int(ud.FollowingUser))
		if err != nil {
			return []models.Followersresponse{}, err
		}
		userresp = append(userresp, models.Followersresponse{
			Username: details.Username,
			Profile:  details.Profile,
		})
	}
	return userresp, nil

}

func (ad *authUseCase) Followings(id int) ([]models.Followersresponse, error) {
	userid := ad.authRepository.CheckUserAvailability(id)
	if !userid {
		return []models.Followersresponse{}, errors.New("user doenst exist")
	}

	ids, err := ad.authRepository.Followings(id)
	if err != nil {
		return []models.Followersresponse{}, err
	}

	var userresp []models.Followersresponse

	for _, ud := range ids {
		details, err := ad.authRepository.Followdetails(int(ud.FollowingUser))
		if err != nil {
			return []models.Followersresponse{}, err
		}
		userresp = append(userresp, models.Followersresponse{
			Username: details.Username,
			Profile:  details.Profile,
		})
	}
	return userresp, nil
}

func (ad *authUseCase) SendOTP(phone string) error {
	phoneNumber := helper.PhoneValidation(phone)
	if !phoneNumber {
		return errors.New("invalid phone number")
	}
	ok := ad.authRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("user with this phone does not exist")
	}
	helper.TwilioSetup(ad.cfg.ACCOUNTSID, ad.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, ad.cfg.SERVICESID)
	if err != nil {
		return errors.New("error occured while generatig otp")
	}
	return nil
}
func (ad *authUseCase) VerifyOTP(details models.VerifyData) (*domain.TokenUser, error) {
	helper.TwilioSetup(ad.cfg.ACCOUNTSID, ad.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(ad.cfg.SERVICESID, details.Code, details.PhoneNumber)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	userdetails, err := ad.authRepository.UserDetailsUsingPhone(details.PhoneNumber)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	access, refresh, err := helper.GenerateTokenUser(userdetails)
	if err != nil {
		return &domain.TokenUser{}, err
	}
	var user models.UserDetailsResponse

	err = copier.Copy(&user, &userdetails)
	if err != nil {
		return &domain.TokenUser{}, err
	}
	return &domain.TokenUser{
		User:         user,
		AccesToken:   access,
		RefreshToken: refresh,
	}, nil
}

func (ad *authUseCase) UploadProfilepic(id int, file []byte) error {

	username, _ := ad.authRepository.GetUserName(id)

	fileUID := uuid.New()
	fileName := fileUID.String()
	s3path := username + fileName

	url, err := helper.AddImageToAwsS3(file, s3path)
	if err != nil {
		return err
	}
	err = ad.authRepository.UploadProfilepic(id, url)
	if err != nil {
		return err
	}
	return nil
}

func (ad *authUseCase) DoesUserExist(id int64) bool {
	ID := int(id)
	userExist := ad.authRepository.CheckUserAvailability(ID)
	if !userExist {
		return false
	}
	return userExist
}

func (ad *authUseCase) FindUserName(id int64) (string, error) {
	ID := int(id)
	username, err := ad.authRepository.GetUserName(ID)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (ad *authUseCase) UserData(id int) (models.UserData, error) {
	Id := id

	details, err := ad.authRepository.UserData(Id)
	fmt.Println("cccccccccc", details.ID)
	if err != nil {
		return models.UserData{}, err
	}
	return details, nil

}
