package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
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
	id_string, _ := c.Get("id")
	id, _ := id_string.(int)
	discription := c.PostForm("discription")

	form, _ := c.MultipartForm()
	// if err != nil {
	// 	errorres := response.ClientResponse(http.StatusBadRequest, "error in retreiving post", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorres)
	// 	return
	// }
	files := form.File["files"]
	fmt.Println("fileeeeeee", files)

	// if len(files) == 0 {
	// 	errres := response.ClientResponse(http.StatusBadRequest, "no files are provided", nil, nil)
	// 	c.JSON(http.StatusBadRequest, errres)
	// 	return
	// }
	if len(files) > 0 {
		for _, file := range files {
			err := ad.GRPC_Client.AddTweet(id, file, discription)
			if err != nil {
				errs := response.ClientResponse(http.StatusBadRequest, "could not change one or more images", nil, err.Error())
				c.JSON(http.StatusBadRequest, errs)
				return
			}
		}
	} else {
		err := ad.GRPC_Client.AddTweet2(id, discription)
		if err != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "error happened while posting", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
	}

	successres := response.ClientResponse(http.StatusOK, "successfully added the tweet", nil, nil)
	c.JSON(http.StatusOK, successres)

}

func (ad *TweetHandler) GetOurTweet(c *gin.Context) {
	id_string, _ := c.Get("id")
	id, _ := id_string.(int)

	detailsm, err := ad.GRPC_Client.GetOurTweet(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in getting your posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully got all your posts", detailsm, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) GetOthersTweet(c *gin.Context) {
	userId := c.Query("id")
	ID, err := strconv.Atoi(userId)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in coversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	detailsm, err := ad.GRPC_Client.GetOthersTweet(ID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in getting their  posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully got all their posts", detailsm, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) EditTweet(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)
	discription := c.Query("discription")

	postid := c.Query("id")
	postID, err := strconv.Atoi(postid)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errr := ad.GRPC_Client.EditTweet(id, postID, discription)
	if errr != nil {
		err := response.ClientResponse(http.StatusBadRequest, "errro in editing your post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully edited your post", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) DeletePost(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errr := ad.GRPC_Client.DeletePost(id, postID)
	if errr != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in deleting video", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully deleted the post", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) LikePost(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errr := ad.GRPC_Client.LikePost(id, postID)
	if errr != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error while like the post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully liked the post", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) UnLikePost(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errr := ad.GRPC_Client.UnLikePost(id, postID)
	if errr != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "errror while unliking the post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully unliked the post", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *TweetHandler) SavePost(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errr := ad.GRPC_Client.SavePost(id, postID)
	if errr != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in saving the post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	succesre := response.ClientResponse(http.StatusOK, "successfully added  into bookmarks", nil, nil)
	c.JSON(http.StatusOK, succesre)
}

func (ad *TweetHandler) UnSavePost(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errr := ad.GRPC_Client.UnSavePost(id, postID)
	if errr != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in unsaving post", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	sucsres := response.ClientResponse(http.StatusOK, "successfully unsaved the post", nil, nil)
	c.JSON(http.StatusOK, sucsres)
}

func (ad *TweetHandler) CommentPost(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	post_id := c.Query("id")
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	comment := c.Query("comment")
	parent_id := c.Query("parentid")
	if parent_id != "" {
		parentid, errr := strconv.Atoi(parent_id)
		if errr != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "error in conversion parent id", nil, errr.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		errrs := ad.GRPC_Client.RplyCommentPost(id, postID, comment, parentid)
		if errrs != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "error while replying comment", nil, errrs.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		succesres := response.ClientResponse(http.StatusOK, "successfully replyed comment", nil, nil)
		c.JSON(http.StatusOK, succesres)
	} else {
		errs := ad.GRPC_Client.CommentPost(id, postID, comment)
		if errs != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "error while commenting ", nil, errs.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		succesres := response.ClientResponse(http.StatusOK, "successfully commented the post", nil, nil)
		c.JSON(http.StatusOK, succesres)
	}

}

func (ad *TweetHandler) GetComments(c *gin.Context) {
	post_id := c.Query("id")
	postid, err := strconv.Atoi(post_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	details, errr := ad.GRPC_Client.GetComments(postid)
	if errr != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in getting comments", nil, errr.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully got all comments", details, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) EditComments(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	comment_id := c.Query("commentid")
	commentid, err := strconv.Atoi(comment_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	comment := c.Query("comment")
	errs := ad.GRPC_Client.EditComments(id, commentid, comment)
	if errs != nil {
		errr := response.ClientResponse(http.StatusBadRequest, "error in updating comment", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errr)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully edited the comment", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *TweetHandler) DeleteComments(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)
	comment_id := c.Query("commentid")
	commentid, err := strconv.Atoi(comment_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	errs := ad.GRPC_Client.DeleteComments(id, commentid)
	if errs != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "error in delete the comments", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully deleted the comment", nil, nil)
	c.JSON(http.StatusOK, succesres)

}
