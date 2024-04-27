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
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}
