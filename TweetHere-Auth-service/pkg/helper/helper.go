package helper

import (
	"Tweethere-Auth/pkg/config"
	"Tweethere-Auth/pkg/utils/models"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"
)

type authCustomClaimsAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

type authCustomClaimsUser struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func PasswordHash(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashpassword)
	return hash, nil
}

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {

	claims := &authCustomClaimsAdmin{
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Email:     admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))

	if err != nil {
		fmt.Println("Error is ", err)
		return "", err
	}
	return tokenString, nil
}

func GenerateTokenUser(user models.UserDetailsResponse) (string, string, error) {

	accessTokenclaims := &authCustomClaimsUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenclaims := &authCustomClaimsUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenclaims)
	accessTokenstring, err := accessToken.SignedString([]byte("accesssecret"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenclaims)
	refreshTokenstring, err := refreshToken.SignedString([]byte("refreshsecret"))
	if err != nil {
		return "", "", err
	}

	return accessTokenstring, refreshTokenstring, nil
}

func CompareHashAndPassword(a string, b string) error {
	fmt.Println("aaaa is ", a)
	fmt.Println("bbb is", b)
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		fmt.Println("ncncn", err)
		return err
	}
	fmt.Println("erorr is ", err)
	return nil
}

func PhoneValidation(phone string) bool {
	phoneNumber := phone
	pattern := `^\d{10}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(phoneNumber)
	return value
}

var client *twilio.RestClient

func TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func TwilioSendOTP(phone string, serviceID string) (string, error) {
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {
		return " ", err
	}
	return *resp.Sid, nil
}

func TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")
}

func AddImageToAwsS3(file []byte, filename string) (string, error) {
	cfg, _ := config.LoadConfig()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.REGION),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWS_ACCESS_KEY_ID,
			cfg.AWS_SECRET_ACCESS_KEY,
			"",
		),
	})

	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	bucketName := "tweethere"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)
	return url, nil
}
