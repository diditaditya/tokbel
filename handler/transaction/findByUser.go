package transaction

import (
	"net/http"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Get all transactions of current user
// @Schemes http
// @Description Get all transactions of current user
// @Tags Transaction
// @Produce json
// @Security BearerAuth
// @Success 200 {array} []entity.TransactionHistory
// @Router /transactions/my-transactions [get]
func FindByUser(service *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("userId")
		data := service.FindByUserId(userId)
		c.JSON(http.StatusOK, data)
	}
}
