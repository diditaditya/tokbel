package middleware

import (
	"net/http"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

func ForbidRoles(userService *services.UserService, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("userId")

		user, err := userService.FindById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		for _, role := range roles {
			if user.Role != role {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": "forbidden",
				})
				return
			}
		}

		c.Next()
	}
}
