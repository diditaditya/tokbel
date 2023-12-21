package product

import (
	"net/http"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Get all products
// @Schemes http
// @Description Get all products
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Success 200 {array} []entity.Product
// @Router /products [get]
func FindAll(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		products := service.FindAll()
		c.JSON(http.StatusOK, products)
	}
}
