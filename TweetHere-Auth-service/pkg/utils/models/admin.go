package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type AdminSignUp struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"  validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type UserSignup struct {
	ID          uint   `josn:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
	IsBlocked   bool   `json:"is_blocked"`
}

type UserDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `jsons:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
