package interfaces

import "tweet-service/pkg/utils/models"

type TweetRepository interface {
	AddTweet(userid int, description string) (uint, error)
	UploadMedia(tid uint, url string) error
	GetOurTweet(id int) ([]models.PostResponse, error)
	EditTweet(id int, tweetid int, description string) error
	DeletePost(id int, tweetid int) error
	PostExist(postid int) bool
	LikePost(id int, userid int) error
	UnLikePost(id int, userid int) error
	SavePost(id int, postid int) error
	UnSavePost(id int, postid int) error
	CommentPost(id int, postid int, comment string) error
	RplyCommentPost(id int,postid int, comment string,parentid int)error
	// GetComments(postid int)
}
