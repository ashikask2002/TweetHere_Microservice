package server

import (
	"TweetHere-API/pkg/api/handler"
	"TweetHere-API/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler *handler.AuthHandler) *ServerHttp {
	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/admin/login", authHandler.LoginHandler)
	router.POST("/admin/signup", authHandler.AdminSignUp)

	router.POST("/user/signup", authHandler.UserSignUp)
	router.POST("/user/login", authHandler.UserLogin)
	router.POST("/user/otplogin", authHandler.UserOTPLogin)
	router.POST("/user/otpverify", authHandler.VerifyOTP)

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

		}

		adminmanagement := router.Group("/admins")
		{
			adminmanagement.GET("/userdetails", authHandler.GetUser)
			adminmanagement.PATCH("/block", authHandler.BlockUser)
			adminmanagement.PATCH("/unblock", authHandler.UnBlockUser)

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
