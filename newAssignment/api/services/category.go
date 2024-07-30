package services

import (
	"fmt"
	"newAssignment/db"
	"newAssignment/models"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// CreateCategoryService handles category creation with image upload
func CreateCategoryService(c *gin.Context) error {
	var category models.Category

	userId, ok := c.Get("userId")
	if !ok {
		return fmt.Errorf("user not authenticated")
	}

	file, err := c.FormFile("image")
	if err != nil {
		return fmt.Errorf("image upload failed: %w", err)
	}

	if err := c.ShouldBind(&category); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	category.CreatedBy = userId.(uint)

	// Check if category already exists
	var existingCategory models.Category
	if err := db.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return fmt.Errorf("category with this name already exists")
	}

	// Save uploaded file
	uploadDir := "uploads/category"
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create upload directory: %w", err)
	}
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}
	category.Image = filePath

	// Create category in the database
	if result := db.DB.Create(&category); result.Error != nil {
		return fmt.Errorf("failed to create category: %w", result.Error)
	}

	return nil
}

// ListCategoriesService retrieves all categories from the database
func ListCategoriesService() ([]models.Category, error) {
	var categories []models.Category
	if result := db.DB.Find(&categories); result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// DeleteCategoryService deletes a category by ID
func DeleteCategoryService(id string) error {
	if result := db.DB.Delete(&models.Category{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
