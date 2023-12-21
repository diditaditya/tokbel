package category

import (
	"net/http"
	"strconv"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a category
// @Schemes http
// @Description Delete a category
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category id"
// @Success 200 {object} handler.MessageResponse
// @Router /categories/:id [delete]
func Delete(service *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		categoryId, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}

		err = service.Delete(categoryId)
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
