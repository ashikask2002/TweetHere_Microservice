package models

import "time"

type PostResponse struct {
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Url         string    `json:"url" gorm:"column:media_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type CommentsResponse struct {
	UserId    int       `json:"user_id"`
	Username  string    `json:"username"`
	Profile   string    `json:"profile"`
	Comment   string    `json:"comment" gorm:"column:comment_text"`
	CreatedAt time.Time `json:"created_at"`
}

type UserData struct {
	UserID   int    `json:"user_id" `
	Username string `json:"username"`
	Profile  string `json:"profile"`
}
