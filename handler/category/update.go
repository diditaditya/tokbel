package category

import (
	"net/http"
	"strconv"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Update a category
// @Schemes http
// @Description Update a category
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Param category body CategoryRequest true "New category data"
// @Success 200 {object} CategoryResponse
// @Router /categories/:id [patch]
func Update(service *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		categoryId, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}

		var body CategoryRequest
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

		updated, err := service.UpdateType(categoryId, body.Type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
		}

		resp := CategoryResponse{
			Id:                updated.Id,
			Type:              updated.Type,
			SoldProductAmount: updated.SoldProductAmount,
			CreatedAt:         updated.CreatedAt,
		}

		c.JSON(http.StatusOK, resp)
	}
}
