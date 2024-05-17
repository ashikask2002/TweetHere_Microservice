package domain

import "time"

type Post struct {
	ID          uint      `json:"id" gorm:"primary_key;not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	MediaURL    string    `json:"media"`
	CreatedAt   time.Time `json:"created_at"`
}

type Like struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	PostID    uint      `json:"post_id" gorm:"not_null"`
	CreatedAt time.Time `json:"created_at"`
}

type BookMark struct {
	ID        uint      `json:"id" gorm:"primary_key;not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	PostID    uint      `json:"post_id" gorm:"not_null"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID          uint      `json:"id" gorm:"primary_key;not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	PostID      uint      `json:"post_id" gorm:"not null"`
	CommentText string    `json:"comment_text" gorm:"not null"`
	ParentID    uint      `json:"parent_id"`
	CreatedAt   time.Time `json:"created_at"`
}
