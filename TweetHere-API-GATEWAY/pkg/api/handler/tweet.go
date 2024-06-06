package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/logging"
	"TweetHere-API/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TweetHandler struct {
	GRPC_Client interfaces.TweetClient
}

func NewTweetHandler(tweetClient interfaces.TweetClient) *TweetHandler {
	return &TweetHandler{
		GRPC_Client: tweetClient,
	}
}

func (ad *TweetHandler) AddTweet(c *gin.Context) {
	//  var postDetails models.PostDetails
	logEntry := logging.GetLogger().WithField("context", "AddTweetHandler")
	logEntry.Info("Processing Add Tweet request")
	id_string, _ := c.Get("id")
	id, _ := id_string.(int)
	discription := c.PostForm("discription")

	form, _ := c.MultipartForm()

	files := form.File["files"]
	fmt.Println("fileeeeeee", files)

	if len(files) > 0 {
		for _, file := range files {
			err := ad.GRPC_Client.AddTweet(id, file, discription)
			if err != nil {
				logEntry.WithError(err).Error("Error opening uploaded file")
				errs := response.ClientResponse(http.StatusBadRequest, "could not change one or more images", nil, err.Error())
				c.JSON(http.StatusBadRequest, errs)
				return
			}
		}
	} else {
		err := ad.GRPC_Client.AddTweet2(id, discription)
		if err != nil {
			logEntry.WithError(err).Error("Error adding tweet")
			errs := response.ClientResponse(http.StatusBadRequest, "error happened while posting", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
	}
	logEntry.Info("Successfully added tweet for user")
	successres := response.ClientResponse(http.StatusOK, "successfully added the tweet", nil, nil)
	c.JSON(http.StatusOK, successres)

}

func (ad *TweetHandler) GetOurTweet(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetOurTweetHandler")
	logEntry.Info("Processing GetOurTweet request")
	id_string, _ := c.Get("id")
	id, _ := id_string.(int)

	detailsm, err := ad.GRPC_Client.GetOurTweet(id)
	if err != nil {
		logEntry.WithError(err).Error("error in getting tweet")
		errs := response.ClientResponse(http.StatusBadRequest, "error in getting your posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("successfully got all request")
	succesres := response.ClientResponse(http.StatusOK, "successfully got all your posts", detailsm, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) GetOthersTweet(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetOthersTweetHandler")
	logEntry.Info("Processing GetOthersTweet request")
	userId := c.Query("id")
	ID, err := strconv.Atoi(userId)

	if err != nil {
		logEntry.WithError(err).Error("error id conversion")
		errs := response.ClientResponse(http.StatusBadRequest, "error in coversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	detailsm, err := ad.GRPC_Client.GetOthersTweet(ID)
	if err != nil {
		logEntry.WithError(err).Error("Error in getting others tweet call")
		errs := response.ClientResponse(http.StatusBadRequest, "error in getting their  posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("successfully got their posts")
	succesres := response.ClientResponse(http.StatusOK, "successfully got all their posts", detailsm, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) EditTweet(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "EditTweetHandler")
	logEntry.Info("Processing EditTweet request")
	id_string, _ := c.Get("id")
	id := id_string.(int)
	discription := c.Query("discription")

	postid := c.Query("id")
	postID, err := strconv.Atoi(postid)
	if err != nil {
		logEntry.WithError(err).Error("conversion error of id")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errr := ad.GRPC_Client.EditTweet(id, postID, discription)
	if errr != nil {
		logEntry.WithError(err).Error("error while Edit")
		err := response.ClientResponse(http.StatusBadRequest, "errro in editing your post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	logEntry.Info("successfully edited")
	succesres := response.ClientResponse(http.StatusOK, "successfully edited your post", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) DeletePost(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "DeleteHandler")
	logEntry.Info("Processing DeletePost request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("conversion error of id")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errr := ad.GRPC_Client.DeletePost(id, postID)
	if errr != nil {
		logEntry.WithError(err).Error("error while deleting")
		errs := response.ClientResponse(http.StatusBadRequest, "error in deleting video", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("Deletion successful")
	succesres := response.ClientResponse(http.StatusOK, "successfully deleted the post", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) LikePost(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "LikePostHandler")
	logEntry.Info("Processing LikePost request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errr := ad.GRPC_Client.LikePost(id, postID)
	if errr != nil {
		logEntry.WithError(errr).Error("error in like post")
		errs := response.ClientResponse(http.StatusBadRequest, "error while like the post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("like post successful")
	succesres := response.ClientResponse(http.StatusOK, "successfully liked the post", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) UnLikePost(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UnLikePostHandler")
	logEntry.Info("Processing UnLikePost request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errr := ad.GRPC_Client.UnLikePost(id, postID)
	if errr != nil {
		logEntry.WithError(errr).Error("Error Unliking call")
		errs := response.ClientResponse(http.StatusBadRequest, "errror while unliking the post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("unlike is Success")
	succesres := response.ClientResponse(http.StatusOK, "successfully unliked the post", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) SavePost(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "SavePostHandler")
	logEntry.Info("Processing SavePost request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errr := ad.GRPC_Client.SavePost(id, postID)
	if errr != nil {
		logEntry.WithError(errr).Error("error in saving")
		errs := response.ClientResponse(http.StatusBadRequest, "error in saving the post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("SavePost is success")
	succesre := response.ClientResponse(http.StatusOK, "successfully added  into bookmarks", nil, nil)
	c.JSON(http.StatusOK, succesre)
}

func (ad *TweetHandler) UnSavePost(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UnSavePostHandler")
	logEntry.Info("Processing UnSavePost request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errr := ad.GRPC_Client.UnSavePost(id, postID)
	if errr != nil {
		logEntry.WithError(errr).Error("error in unsaving")
		errs := response.ClientResponse(http.StatusBadRequest, "error in unsaving post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("unsave successfull")
	sucsres := response.ClientResponse(http.StatusOK, "successfully unsaved the post", nil, nil)
	c.JSON(http.StatusOK, sucsres)
}

func (ad *TweetHandler) CommentPost(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "CommentPostHandler")
	logEntry.Info("Processing CommentPost request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	comment := c.Query("comment")
	parent_id := c.Query("parentid")
	if parent_id != "" {
		parentid, errr := strconv.Atoi(parent_id)
		if errr != nil {
			logEntry.WithError(errr).Error("parent id conversion error")
			errs := response.ClientResponse(http.StatusBadRequest, "error in conversion parent id", nil, errr.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		errrs := ad.GRPC_Client.RplyCommentPost(id, postID, comment, parentid)
		if errrs != nil {
			logEntry.WithError(errrs).Error("error in replying")
			errs := response.ClientResponse(http.StatusBadRequest, "error while replying comment", nil, errrs.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		logEntry.Info("replying success")
		succesres := response.ClientResponse(http.StatusOK, "successfully replyed comment", nil, nil)
		c.JSON(http.StatusOK, succesres)
	} else {
		errs := ad.GRPC_Client.CommentPost(id, postID, comment)
		if errs != nil {
			logEntry.WithError(errs).Error("error in commenting")
			errs := response.ClientResponse(http.StatusBadRequest, "error while commenting ", nil, errs.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		logEntry.Info("commentig successfull")
		succesres := response.ClientResponse(http.StatusOK, "successfully commented the post", nil, nil)
		c.JSON(http.StatusOK, succesres)
	}

}

func (ad *TweetHandler) GetComments(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetCommentsHandler")
	logEntry.Info("Processing GetComments request")
	post_id := c.Query("id")
	postid, err := strconv.Atoi(post_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	details, errr := ad.GRPC_Client.GetComments(postid)
	if errr != nil {
		logEntry.WithError(errr).Error("error in Getcomments")
		errs := response.ClientResponse(http.StatusBadRequest, "error in getting comments", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("getComment is success")
	succesres := response.ClientResponse(http.StatusOK, "successfully got all comments", details, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) EditComments(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "EditCommentHandler")
	logEntry.Info("Processing EditComment request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	comment_id := c.Query("commentid")
	commentid, err := strconv.Atoi(comment_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	comment := c.Query("comment")
	errs := ad.GRPC_Client.EditComments(id, commentid, comment)
	if errs != nil {
		logEntry.WithError(errs).Error("update comment cause error")
		errr := response.ClientResponse(http.StatusBadRequest, "error in updating comment", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errr)
		return
	}
	logEntry.Info("edit comment successfull")
	succesres := response.ClientResponse(http.StatusOK, "successfully edited the comment", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) DeleteComments(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "DeleteCommentHandler")
	logEntry.Info("Processing DeleteComment request")
	id_string, _ := c.Get("id")
	id := id_string.(int)
	comment_id := c.Query("commentid")
	commentid, err := strconv.Atoi(comment_id)
	if err != nil {
		logEntry.WithError(err).Error("id conversion error")
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errs := ad.GRPC_Client.DeleteComments(id, commentid)
	if errs != nil {
		logEntry.WithError(err).Error("error in commentdelete")
		errres := response.ClientResponse(http.StatusBadRequest, "error in delete the comments", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}
	logEntry.Info("deletecomment Success")
	succesres := response.ClientResponse(http.StatusOK, "successfully deleted the comment", nil, nil)
	c.JSON(http.StatusOK, succesres)

}
