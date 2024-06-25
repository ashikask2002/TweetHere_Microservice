package interfaces

import (
	"TweetHere-API/pkg/utils/models"
	"mime/multipart"
)

type TweetClient interface {
	AddTweet(id int, file *multipart.FileHeader, postdetails string) error
	AddTweet2(id int, postdetails string) error
	GetOurTweet(id int) ([]models.PostResponse, error)
	GetOthersTweet(id int) ([]models.PostResponse, error)
	EditTweet(id int, postid int, discription string) error
	DeletePost(id int, postId int) error
	LikePost(id int, postid int) error
	UnLikePost(id int, Postid int) error
	SavePost(id int, postid int) error
	UnSavePost(id int, postid int) error
	RplyCommentPost(id int, postid int, comment string, parentid int) error
	CommentPost(id int, postid int, comment string) error
	GetComments(postid int) ([]models.CommentsResponse, error)
	EditComments(id int, commentid int, comment string) error
	DeleteComments(id int, commentid int) error
	Home(useid int)([]models.PostResponses,error)
}
