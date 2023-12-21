package middleware

import (
	"log"
	"net/http"
	"strings"
	"tokbel/auth"

	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		jwtToken := strings.Split(header, " ")
		if len(jwtToken) < 2 {
			log.Println("invalid authorization format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		if jwtToken[0] != "Bearer" {
			log.Println("invalid authorization scheme")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		userId, err := auth.VerifyJWT(jwtToken[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
