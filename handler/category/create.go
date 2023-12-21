package category

import (
	"net/http"
	"tokbel/entity"
	"tokbel/handler"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new category
// @Schemes http
// @Description Create a new category
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Param category body CategoryRequest true "New category data"
// @Success 201 {object} CategoryResponse
// @Router /categories [post]
func Create(service *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body CategoryRequest
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

		category := entity.Category{
			Type: body.Type,
		}
		created, err := service.Create(&category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating new category",
			})
		}

		resp := CategoryResponse{
			Id:                created.Id,
			Type:              created.Type,
			SoldProductAmount: created.SoldProductAmount,
			CreatedAt:         category.CreatedAt,
		}

		c.JSON(http.StatusCreated, resp)
	}
}
