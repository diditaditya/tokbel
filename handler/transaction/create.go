package transaction

import (
	"net/http"
	"tokbel/entity"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create a transaction
// @Schemes http
// @Description Create a transaction
// @Tags Transaction
// @Produce json
// @Security BearerAuth
// @Param transaction body TransactionRequest true "New transaction data"
// @Success 201 {object} CreateTransactionResponse
// @Router /transactions [post]
func Create(service *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body TransactionRequest
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

		trx := entity.TransactionHistory{
			ProductId: body.ProductId,
			Quantity:  body.Quantity,
			UserId:    c.GetInt("userId"),
		}
		created, err := service.Create(&trx)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		resp := CreateTransactionResponse{
			Message: "You have successfully purchase the product",
			TransactionBill: TransactionBill{
				TotalPrice:   created.TotalPrice,
				Quantity:     created.Quantity,
				ProductTitle: created.Product.Title,
			},
		}

		c.JSON(http.StatusCreated, resp)
	}
}
