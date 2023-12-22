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
// @Success 200 {array} []UserTransaction
// @Router /transactions/my-transactions [get]
func FindByUser(service *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("userId")
		data := service.FindByUserId(userId)
		resp := []UserTransaction{}
		for _, trx := range data {
			userTrx := UserTransaction{
				Id:         trx.Id,
				ProductId:  trx.ProductId,
				UserId:     trx.UserId,
				Quantity:   trx.Quantity,
				TotalPrice: trx.TotalPrice,
				Product:    trx.Product,
			}
			resp = append(resp, userTrx)
		}
		c.JSON(http.StatusOK, resp)
	}
}
