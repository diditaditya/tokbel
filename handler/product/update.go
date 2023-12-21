package product

import (
	"fmt"
	"net/http"
	"strconv"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Update a product
// @Schemes http
// @Description Update a product
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product id"
// @Param product body ProductRequest true "Product data"
// @Success 200 {object} entity.Product
// @Router /products/{id} [put]
func Update(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		fmt.Println("==>> id str: ", idStr)
		productId, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
				"error":   err.Error(),
			})
			return
		}

		var body ProductRequest
		err = c.ShouldBindJSON(&body)
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

		found, err := service.FindById(productId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
				"error":   res,
			})
			return
		}

		found.Title = body.Title
		found.Price = body.Price
		found.Stock = body.Stock
		found.CategoryId = body.CategoryId

		updated, err := service.Update(found)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
		}

		c.JSON(http.StatusOK, updated)
	}
}
