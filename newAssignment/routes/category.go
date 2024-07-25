package routes

import (
	"net/http"
	"newAssignment/db"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

func createCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if result := db.DB.Create(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": category})
}

func listCategories(c *gin.Context) {
	var categories []models.Category
	if result := db.DB.Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// Delete Category
func deleteCategory(c *gin.Context) {
	id := c.Param("id")
	if result := db.DB.Delete(&models.Category{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
