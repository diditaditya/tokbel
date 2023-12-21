package user

import (
	"fmt"
	"net/http"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

type TopUpRequest struct {
	Balance int `json:"balance" validate:"required,gte=0,lte=100000000"`
}

// @Summary Top up
// @Schemes http
// @Description Top up user balance
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param balance body TopUpRequest true "Amount to add to user balance"
// @Success 200 {object} handler.MessageResponse
// @Router /users/topup [patch]
func TopUp(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body TopUpRequest
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

		id := c.GetInt("userId")

		newBalance, err := userService.TopUp(id, body.Balance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		resp := handler.MessageResponse{
			Message: fmt.Sprintf("Your balance has been successfully updated to Rp %d", newBalance),
		}
		c.JSON(http.StatusOK, resp)
	}
}
