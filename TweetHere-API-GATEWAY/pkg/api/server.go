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


	router.POST("/user/signup",authHandler.UserSignUp)
	router.POST("/user/login",authHandler.UserLogin)

	router.Use(middleware.AdminAuthMiddleware())
	

	return &ServerHttp{engine: router}
}

func (s *ServerHttp) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":5000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
