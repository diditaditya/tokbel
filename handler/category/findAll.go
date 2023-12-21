package category

import (
	"net/http"
	"tokbel/services"

	"github.com/gin-gonic/gin"
)

// @Summary Get all categories
// @Schemes http
// @Description Get all categories
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Success 200 {array} []entity.Category
// @Router /categories [get]
func FindAll(service *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := service.FindAll()
		c.JSON(http.StatusOK, categories)
	}
}
