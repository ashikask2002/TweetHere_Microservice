package models

import "time"

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type AdminSignup struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}

type TokenAdmin struct {
	Admin AdminDetailsResponse
	Token string
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"  validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type UserSignup struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
}

type User struct {
	ID          uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname   string `json:"firstname" gorm:"validate:required"`
	Lastname    string `json:"lastname" gorm:"validate:required"`
	Username    string `json:"username" gorm:"validate:required"`
	Phone       string `json:"phone" gorm:"validate:required"`
	Email       string `json:"email" gorm:"validate:required"`
	DateOfBirth string `json:"date_of_birth" gorm:"validate:required"`
	Password    string `json:"password" gorm:"validate:required"`
}

type UserDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `jsons:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
type TokenUser struct {
	User         UserDetailsResponse
	AccesToken   string
	RefreshToken string
}

type UserProfile struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Profile     string `json:"profile"`
	Bio         string `json:"bio"`
}

type UserProfileResponse struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Profile     string `json:"profile"`
	Bio         string `json:"bio"`
}

type UserDetails struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	IsBlocked   bool   `json:"is_blocked"`
	Profile     string `json:"profile"`
	Bio         string `json:"bio"`
}

type ChangePassword struct {
	Oldpassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	RePassword  string `json:"re_password"`
}
type UserDetails4user struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Profile     string `json:"profile"`
	Bio         string `json:"bio"`
}
type UserOTPLogin struct {
	Email string `json:"email" validate:"email"`
}

type OtpVerification struct {
	Email string `json:"email" validate:"email"`
	Otp   string `json:"otp" validate:"required,len=4,number"`
}

type Followersresponse struct {
	Username string `json:"username"`
	Profile  string `json:"profile"`
}

type OTPData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
}
type VerifyData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}

type PostDetails struct {
	Description string `json:"description"`
}

type PostResponse struct {
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Url         string    `json:"url" gorm:"column:media_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type CommentsResponse struct {
	Username  string    `json:"username"`
	Profile   string    `json:"profile"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
