package interfaces

import "tweet-service/pkg/utils/models"

type TweetUseCase interface {
	AddTweet(id int64, file []byte, discritption string) error
	AddTweet2(id int64, discritption string) error
	GetOurTweet(id int) ([]models.PostResponse, error)
	GetOthersTweet(id int) ([]models.PostResponse, error)
	EditTweet(id int, postid int, description string) error
	DeletePost(id int, postid int) error
	LikePost(id int, postid int) error
	UnLikePost(id int, postid int) error
	SavePost(id int, postid int) error
	UnSavePost(id int, postid int) error
	CommentPost(id int, postid int, comment string) error
	RplyCommentPost(id int, postid int, comment string, parentid int) error
	GetComments(postid int) ([]models.CommentsResponse, error)
	EditComments(id int, commentid int, comment string) error
	DeleteComments(id int, commentid int) error
}
