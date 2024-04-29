package middleware

import (
	"TweetHere-API/pkg/helper"
	"TweetHere-API/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenheader := c.GetHeader("authorization")
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



// func userAuthMiddleware(c *gin.Context) {
// 	accessToken := c.Request.Header.Get("Authorization")
// 	fmt.Println("access_token", accessToken)

// 	// Parse and validate the access token
// 	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("accesssecret"), nil
// 	})

// 	if err != nil {
// 		fmt.Println("error parsing token:", err)
// 		// The access token is invalid.
// 		c.AbortWithStatusJSON(401, gin.H{
// 			"message": "token error",
// 			"err":     err.Error(),
// 		})
// 		return
// 	}

// 	// Check if the token is valid
// 	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
// 		fmt.Println("invalid token")
// 		c.AbortWithStatusJSON(401, gin.H{
// 			"message": "invalid token",
// 		})
// 		return
// 	}

// 	// Token is valid, proceed with the request
// 	c.Next()
// }

