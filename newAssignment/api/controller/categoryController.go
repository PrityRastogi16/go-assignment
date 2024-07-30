package controller

import (
	"net/http"
	"newAssignment/api/business"

	"github.com/gin-gonic/gin"
)

// CreateCategoryController handles the creation of a category
// @Summary Create a category
// @Description Creates a new category with an image
// @Tags Category
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Category Image"
// @Param category body models.Category true "Category Details"
// @Success 201 {object} models.CategoryCreationResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /category [post]
// @Security bearerToken
func CreateCategoryController(c *gin.Context) {
	if err := business.CreateCategoryBusiness(c); err != nil {
		if err.Error() == "user not authenticated" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{"User not authenticated"})
		} else if err.Error() == "image upload failed" {
			c.JSON(http.StatusBadRequest, ErrorResponse{"Image upload failed"})
		} else if err.Error() == "invalid request" {
			c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid request"})
		} else if err.Error() == "category with this name already exists" {
			c.JSON(http.StatusBadRequest, ErrorResponse{"Category with this name already exists"})
		} else if err.Error() == "failed to save image" {
			c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to save image"})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to create category"})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// ListCategoriesController retrieves and returns all categories
// @Summary List all categories
// @Description Fetches a list of all categories from the database
// @Tags Category
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500 {object} ErrorResponse
// @Router /categories [get]
// @Security bearerToken
func ListCategoriesController(c *gin.Context) {
	categories, err := business.ListCategoriesBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// DeleteCategoryController handles the deletion of a category
// @Summary Delete a category
// @Description Deletes a category by its ID
// @Tags Category
// @Param id path string true "Category ID"
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /category/{id} [delete]
// @Security bearerToken
func DeleteCategoryController(c *gin.Context) {
	id := c.Param("id")
	if err := business.DeleteCategoryBusiness(id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{"Category deleted successfully"})
}
