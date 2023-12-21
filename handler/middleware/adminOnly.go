package middleware

import (
	"net/http"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

func AdminOnly(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("userId")

		user, err := userService.FindById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "forbidden",
			})
			return
		}
		c.Next()
	}
}
