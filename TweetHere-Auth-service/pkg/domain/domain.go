package domain

import (
	"Tweethere-Auth/pkg/utils/models"
	"time"
)

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}

type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
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
	IsBlocked   bool   `json:"is_blocked"`
	Profile     string `json:"profile"` // URL of profile picture
	Bio         string `json:"bio"`     // User biography
}

type TokenUser struct {
	User         models.UserDetailsResponse
	AccesToken   string
	RefreshToken string
}

type UserOTPLogin struct {
	Email      string    `json:"email" validate:"email"`
	Expiration time.Time `json:"expiration"`
	Otp        int       `json:"otp"`
}

type FollowingRequest struct {
	UserID        uint      `json:"user_id"`
	FollowingUser uint      `json:"following_user"`
	CreatedAt     time.Time `json:"created_at"`
}
type Followings struct {
	UserID        uint      `json:"user_id"`
	FollowingUser uint      `json:"following_user"`
	CreatedAt     time.Time `json:"created_at"`
}

type Followers struct {
	UserID        uint      `json:"user_id"`
	FollowingUser uint      `json:"following_user"`
	CreatedAt     time.Time `json:"created_at"`
}
