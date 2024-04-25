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

func NewServerHTTP(adminHandler *handler.AdminHandler) *ServerHttp {
	router := gin.New()

	router.Use(gin.Logger())

	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignup)

	router.Use(middleware.AdminAuthMiddleware())

	return &ServerHttp{engine: router}
}

func (s *ServerHttp) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":4000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
