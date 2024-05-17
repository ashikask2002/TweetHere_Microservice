package repository

import (
	"errors"
	"fmt"
	"time"
	"tweet-service/pkg/domain"
	interfaces "tweet-service/pkg/repository/interface"
	"tweet-service/pkg/utils/models"

	"gorm.io/gorm"
)

type tweetRepository struct {
	DB *gorm.DB
}

func NewTweetRespository(DB *gorm.DB) interfaces.TweetRepository {
	return &tweetRepository{
		DB: DB,
	}
}

func (th *tweetRepository) AddTweet(userID int, description string) (uint, error) {
	var postID uint

	err := th.DB.Raw("INSERT INTO posts (user_id, description, created_at) VALUES (?, ?, NOW()) RETURNING id", userID, description).Scan(&postID).Error
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (th *tweetRepository) UploadMedia(tid uint, url string) error {
	err := th.DB.Exec("UPDATE posts SET media_url = ? WHERE id = ?", url, tid).Error
	if err != nil {
		return errors.New("error in updating media")
	}
	return nil
}

func (th *tweetRepository) GetOurTweet(uid int) ([]models.PostResponse, error) {
	var userdetails []models.PostResponse

	query := `SELECT user_id, description, media_url, created_at FROM posts WHERE user_id = ?`

	if err := th.DB.Raw(query, uid).Scan(&userdetails).Error; err != nil {
		return nil, err
	}
	fmt.Println("xxxxxxx", userdetails)

	return userdetails, nil
}

func (th *tweetRepository) EditTweet(id int, tweetID int, description string) error {
	// Execute the SQL query to update the description
	err := th.DB.Exec("UPDATE posts SET description = ? WHERE user_id = ? AND id = ?", description, id, tweetID).Error
	if err != nil {
		return errors.New("this post is does not exist in your account")
	}
	return nil
}

func (TH *tweetRepository) DeletePost(id int, tweetID int) error {
	// Execute the SQL query to delete the post
	err := TH.DB.Exec("DELETE FROM posts WHERE user_id = ? AND id = ?", id, tweetID).Error
	if err != nil {
		return errors.New("check your provided ids again")
	}
	return nil
}

func (th *tweetRepository) PostExist(postid int) bool {
	var count int
	if err := th.DB.Raw("SELECT COUNT(*) FROM posts WHERE id=?", postid).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (th *tweetRepository) LikePost(userID, postID int) error {
	// Check if the user has already liked the post
	var count int64
	if err := th.DB.Model(&domain.Like{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// User has already liked the post
		return errors.New("user has already liked the post")
	}

	// Insert a new like record
	like := domain.Like{
		UserID: uint(userID),
		PostID: uint(postID),
	}
	if err := th.DB.Create(&like).Error; err != nil {
		return err
	}

	return nil
}

func (th *tweetRepository) UnLikePost(userID, postID int) error {
	// Check if the user has liked the post
	var count int64
	if err := th.DB.Model(&domain.Like{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		// User has not liked the post
		return errors.New("user has not liked the post")
	}

	// Delete the like record
	if err := th.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&domain.Like{}).Error; err != nil {
		return err
	}

	return nil
}

func (th *tweetRepository) SavePost(userID, postID int) error {
	// Check if the user has already saved the post
	var count int64
	if err := th.DB.Model(&domain.BookMark{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// User has already saved the post
		return errors.New("user has already saved the post")
	}

	// Insert a new bookmark record
	bookmark := domain.BookMark{
		UserID: uint(userID),
		PostID: uint(postID),
	}
	if err := th.DB.Create(&bookmark).Error; err != nil {
		return err
	}

	return nil
}

func (th *tweetRepository) UnSavePost(userID, postID int) error {
	// Check if the user has saved the post
	var count int64
	if err := th.DB.Model(&domain.BookMark{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		// User has not saved the post
		return errors.New("user has not saved the post")
	}

	// Delete the bookmark record
	if err := th.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&domain.BookMark{}).Error; err != nil {
		return err
	}

	return nil
}

func (th *tweetRepository) CommentPost(userid int, postid int, comment string) error {

	var postCount int64
	if err := th.DB.Model(&domain.Post{}).Where("id = ?", postid).Count(&postCount).Error; err != nil {
		return fmt.Errorf("failed to check post existence: %v", err)
	}
	if postCount == 0 {
		return fmt.Errorf("post not found")
	}

	newComment := domain.Comment{
		UserID:      uint(userid),
		PostID:      uint(postid),
		CommentText: comment,
		CreatedAt:   time.Now(),
	}

	if err := th.DB.Create(&newComment).Error; err != nil {
		return fmt.Errorf("failed to insert comment: %v", err)
	}

	return nil
}

func (th *tweetRepository) RplyCommentPost(userid int, postid int, comment string, parentid int) error {
	// Validate that the post exists
	var postCount int64
	if err := th.DB.Model(&domain.Post{}).Where("id = ?", postid).Count(&postCount).Error; err != nil {
		return fmt.Errorf("failed to check post existence: %v", err)
	}
	if postCount == 0 {
		return fmt.Errorf("post not found")
	}

	// Validate that the parent comment exists
	var parentCommentCount int64
	if err := th.DB.Model(&domain.Comment{}).Where("id = ?", parentid).Count(&parentCommentCount).Error; err != nil {
		return fmt.Errorf("failed to check parent comment existence: %v", err)
	}
	if parentCommentCount == 0 {
		return fmt.Errorf("parent comment not found")
	}

	// Insert the new reply comment
	newComment := domain.Comment{
		UserID:      uint(userid),
		PostID:      uint(postid),
		CommentText: comment,
		ParentID:    uint(parentid),
		CreatedAt:   time.Now(),
	}

	if err := th.DB.Create(&newComment).Error; err != nil {
		return fmt.Errorf("failed to insert reply comment: %v", err)
	}

	return nil
}
