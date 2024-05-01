package middleware

import (
	"TweetHere-API/pkg/helper"
	"TweetHere-API/pkg/utils/response"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenheader := c.GetHeader("authorization")
		fmt.Println("tokennnnnnn", tokenheader)
		if tokenheader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "no auth head provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splited := strings.Split(tokenheader, " ")
		if len(splited) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "invalid token format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenpart := splited[1]
		tokenclaims, err := helper.ValidateToken(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Set("tokenClaims", tokenclaims)
		
		c.Next()
	}
}

func UserAuthMiddleware(c *gin.Context) {

	// accessToken := c.Request.Header.Get("Authorization")
	// fmt.Println("access_token", accessToken)
	// _, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("accesssecret"), nil
	// })

	// if err != nil {
	// 	fmt.Println("it happpens here")
	// 	// The access token is invalid.
	// 	c.AbortWithStatusJSON(401, gin.H{
	// 		"message": "token error",
	// 		"err":     err.Error(),
	// 	})
	// 	return
	// }

	// c.Next()

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		err := errors.New("token invalid")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unautherized",
			"err": err.Error()})
		c.Abort()
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("accesssecret"), nil
	})

	if err != nil || !token.Valid {
		log.Println("Token error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token error",
			"err": err.Error()})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := errors.New("token invalid")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token error happened",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}

	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		err := errors.New("token invalidd")
		c.JSON(http.StatusForbidden, gin.H{"message": "token error happened",
			"err": err.Error()})
		c.Abort()
		return
	}
	fmt.Println("here is the idddddd", id)
	c.Set("id", int(id))

	c.Next()
}
