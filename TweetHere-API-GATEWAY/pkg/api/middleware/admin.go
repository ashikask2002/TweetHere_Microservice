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
