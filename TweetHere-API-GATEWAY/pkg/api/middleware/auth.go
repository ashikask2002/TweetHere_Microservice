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

func AdminAuthMiddleware(c *gin.Context) {
	fmt.Println("arunnnnnnnnnn")

	tokenheader := c.GetHeader("Authorization")
	fmt.Println("tokennnnnnn", tokenheader)
	if tokenheader == "" {
		response := response.ClientResponse(http.StatusUnauthorized, "no auth head provided", nil, nil)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	splited := strings.Split(tokenheader, " ")
	if len(splited) != 2 {
		response := response.ClientResponse(http.StatusUnauthorized, "invalid token formattt", nil, nil)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	tokenpart := splited[1]
	tokenclaims, err := helper.ValidateTokenAdmin(tokenpart)
	if err != nil {
		response := response.ClientResponse(http.StatusUnauthorized, "Invalid Tokennnnn", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	c.Set("tokenClaims", tokenclaims)

	c.Next()
}

func UserAuthMiddleware(c *gin.Context) {

	fmt.Println("ashikkkkkkkkkkkk")
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
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token errorrr",
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

func AuthMiddleware(c *gin.Context) {
	tokenheader := c.GetHeader("Authorization")
	if tokenheader == "" {
		response := response.ClientResponse(http.StatusUnauthorized, "no auth head provided", nil, nil)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	splited := strings.Split(tokenheader, " ")
	if len(splited) != 2 {
		response := response.ClientResponse(http.StatusUnauthorized, "error in splitting", nil, nil)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	fmt.Println("spllitedtoken", splited[1])
	fmt.Println("splited", splited[0])
	tokenpart := splited[0]
	tokenpart1 := splited[1]
	if tokenpart == "admin" {
		tokenclaims, err := helper.ValidateTokenAdmin(tokenpart1)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token admin", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Set("tokenClaims", tokenclaims)

		c.Next()
	} else if tokenpart == "user" {

		token, err := jwt.Parse(tokenpart1, func(token *jwt.Token) (interface{}, error) {
			return []byte("accesssecret"), nil
		})

		if err != nil || !token.Valid {
			log.Println("Token error:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token errorrr",
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

		// tokenclaims, err := helper.ValidateTokenUser(tokenpart1)
		// if err != nil {
		// 	response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token user", nil, err.Error())
		// 	c.JSON(http.StatusUnauthorized, response)
		// 	c.Abort()
		// 	return
		// }
		// 	claims, ok := tokenclaims..(jwt.MapClaims)
		// if !ok {
		// 	err := errors.New("token invalid")
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "token error happened",
		// 		"err":     err.Error(),
		// 	})
		// 	c.Abort()
		// 	return
		// }
		// id, ok := tokenclaims["id"].(float64)
		// c.Set("tokenClaims", tokenclaims)
		// fmt.Println("tokenclaimssss", tokenclaims)
		// fmt.Println("nnnnnnnnn", tokenclaims.ID)
		// c.Set("id", tokenclaims.ID)

		// fmt.Println("qqqqqqq", tokenclaims.Id)

		// c.Next()
	} else {
		err := errors.New("privilage not met")
		response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token format", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

}
