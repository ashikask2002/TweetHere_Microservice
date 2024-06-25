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

type FollowingRequest struct {
	FollowingUser int `json:"following_user"`
	UserID        int `json:"user_id"`
}
type Followersresponse struct {
	Username string `json:"username"`
	Profile  string `json:"profile"`
}

type FollowResp struct {
	FollowingUser int `json:"following_user"`
}

type VerifyData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}

type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Profile  string `json:"profile"`
}
type FollowUsers struct {
	FollowingUser int `json:"following_user"`
}
