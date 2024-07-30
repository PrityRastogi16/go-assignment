package business

import (
	"newAssignment/api/services"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

// CreateCategoryBusiness handles the business logic for creating a category
func CreateCategoryBusiness(c *gin.Context) error {
	return services.CreateCategoryService(c)
}

// ListCategoriesBusiness retrieves all categories
func ListCategoriesBusiness() ([]models.Category, error) {
	return services.ListCategoriesService()
}

// DeleteCategoryBusiness handles the deletion of a category
func DeleteCategoryBusiness(id string) error {
	return services.DeleteCategoryService(id)
}
