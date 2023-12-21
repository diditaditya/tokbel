package user

import (
	"net/http"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary Login
// @Schemes http
// @Description Login
// @Tags User
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {object} LoginResponse
// @Router /users/login [post]
func Login(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Credentials
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}

		res := handler.ValidateStruct(body)
		if res != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   res,
			})
			return
		}

		token, err := userService.Login(body.Email, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid email or password",
			})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{Token: token})
	}
}
