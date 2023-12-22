package product

import (
	"net/http"
	"tokbel/entity"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @Schemes http
// @Description Create a new product
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param product body ProductRequest true "New product data"
// @Success 201 {object} entity.Product
// @Router /products [post]
func Create(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body ProductRequest
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

		product := entity.Product{
			Title:      body.Title,
			Price:      body.Price,
			Stock:      body.Stock,
			CategoryId: body.CategoryId,
		}
		created, err := service.Create(&product)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, created)
	}
}
