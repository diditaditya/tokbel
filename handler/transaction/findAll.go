package transaction

import (
	"net/http"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Get all transactions
// @Schemes http
// @Description Get all transactions (Admin only)
// @Tags Transaction
// @Produce json
// @Security BearerAuth
// @Success 200 {array} []entity.TransactionHistory
// @Router /transactions/user-transactions [get]
func FindAll(service *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := service.FindAll()
		c.JSON(http.StatusOK, data)
	}
}
