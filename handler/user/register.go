package user

import (
	"net/http"
	"tokbel/entity"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

// @Summary Register a new user
// @Schemes http
// @Description Register a new user
// @Tags User
// @Produce json
// @Param user body RegisterRequest true "Registration Request JSON"
// @Success 201 {object} entity.User
// @Router /users/register [post]
func Register(service *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body RegisterRequest
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

		found, _ := service.FindByEmail(body.Email)
		if found != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "email already exists",
			})
			return
		}

		newUser := entity.User{}
		newUser.FullName = body.FullName
		newUser.Email = body.Email
		newUser.Password = body.Password
		newUser.Role = "customer"

		created, err := service.Register(&newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		created.Password = ""
		c.JSON(http.StatusCreated, created)
	}
}
