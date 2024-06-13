package server

import (
	"TweetHere-API/pkg/api/handler"
	"TweetHere-API/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler *handler.AuthHandler, tweetHandler *handler.TweetHandler, chatHandler *handler.ChatHandler, notihandler *handler.NotificationHandler, videocallHandler *handler.VideoCallHandler) *ServerHttp {
	router := gin.New()
	router.Use(gin.Logger())

	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/exit", videocallHandler.ExitPage)
	router.GET("/error", videocallHandler.ErrorPage)
	router.GET("/index", videocallHandler.IndexedPage)

	router.POST("/admin/login", authHandler.LoginHandler)

	router.POST("/user/signup", authHandler.UserSignUp)
	router.POST("/user/login", authHandler.UserLogin)
	router.POST("/user/otplogin", authHandler.UserOTPLogin)
	router.POST("/user/otpverify", authHandler.VerifyOTP)
	router.POST("/user/sendOTP", authHandler.SendOTP)
	router.POST("/user/verifyOTP", authHandler.VerifyOTP)

	router.Use(middleware.AuthMiddleware)
	{

		usermanagement := router.Group("/users")
		{
			usermanagement.POST("/profile", authHandler.UserUpdateProfile)
			usermanagement.PATCH("/changepassword", authHandler.ChangePassword)
			usermanagement.GET("/getyoudetails", authHandler.GetUserDetails)
			usermanagement.POST("/followreq", authHandler.FollowReq)
			usermanagement.POST("/acceptfollowreq", authHandler.AcceptFollowreq)
			usermanagement.POST("unfollow", authHandler.Unfollow)
			usermanagement.GET("/followers", authHandler.Followers)
			usermanagement.GET("/followings", authHandler.Followings)
			usermanagement.POST("/profilepic", authHandler.UploadProfilepic)

			usermanagement.POST("/addpost", tweetHandler.AddTweet)
			usermanagement.GET("/getourpost", tweetHandler.GetOurTweet)
			usermanagement.GET("/getotherspost", tweetHandler.GetOthersTweet)
			usermanagement.PATCH("/editpost", tweetHandler.EditTweet)
			usermanagement.DELETE("/deletpost", tweetHandler.DeletePost)
			usermanagement.POST("/likepost", tweetHandler.LikePost)
			usermanagement.POST("/unlikepost", tweetHandler.UnLikePost)
			usermanagement.POST("/savepost", tweetHandler.SavePost)
			usermanagement.POST("/unsavepost", tweetHandler.UnSavePost)
			usermanagement.POST("/commentpost", tweetHandler.CommentPost)
			usermanagement.PATCH("/editcomment", tweetHandler.EditComments)
			usermanagement.DELETE("/deletecomment", tweetHandler.DeleteComments)
			usermanagement.GET("/getcomments", tweetHandler.GetComments)
		}

		adminmanagement := router.Group("/admins")
		{
			adminmanagement.GET("/userdetails", authHandler.GetUser)
			adminmanagement.PATCH("/block", authHandler.BlockUser)
			adminmanagement.PATCH("/unblock", authHandler.UnBlockUser)

		}

		chatmanagement := router.Group("/chat")
		{
			chatmanagement.GET("", chatHandler.FriendMessage)
			chatmanagement.GET("message", chatHandler.GetChat)
		}

		notificationmanagement := router.Group("/noti")
		{
			notificationmanagement.GET("", notihandler.GetNotification)
		}
	}

	return &ServerHttp{engine: router}
}

func (s *ServerHttp) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":5000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
