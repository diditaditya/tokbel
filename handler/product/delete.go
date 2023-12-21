package product

import (
	"net/http"
	"strconv"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a product
// @Schemes http
// @Description Delete a product
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product id"
// @Success 200 {object} entity.Product
// @Router /products/{id} [put]
func Delete(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		productId, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}

		err = service.Delete(productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
			return
		}

		resp := handler.MessageResponse{
			Message: "success",
		}
		c.JSON(http.StatusOK, resp)
	}
}
