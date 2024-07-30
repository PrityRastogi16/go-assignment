package services

import (
	"fmt"
	"newAssignment/db"
	"newAssignment/models"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// CreateBlogService handles blog creation with image upload
func CreateBlogService(c *gin.Context) error {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		return fmt.Errorf("unable to parse form: %w", err)
	}

	// Get file from form
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		return fmt.Errorf("unable to get file: %w", err)
	}
	defer file.Close()

	// Create a directory to save the image if it doesn't exist
	imageDir := "./uploads/"
	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create image directory: %w", err)
	}

	// Save the file
	filePath := filepath.Join(imageDir, filepath.Base(fileHeader.Filename))

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to save file: %w", err)
	}
	defer out.Close()

	if _, err := out.ReadFrom(file); err != nil {
		return fmt.Errorf("unable to save file: %w", err)
	}

	// Construct image URL or path
	imageURL := fmt.Sprintf("image%s", filepath.Base(filePath))

	var blog models.Blog
	if err := c.ShouldBind(&blog); err != nil {
		return fmt.Errorf("invalid data: %w", err)
	}
	blog.Image = imageURL
	result := db.DB.Create(&blog)
	if result.Error != nil {
		return fmt.Errorf("failed to create blog: %w", result.Error)
	}
	return nil
}

// GetBlogsService retrieves all blogs from the database
func GetBlogsService() ([]models.Blog, error) {
	var blogs []models.Blog
	result := db.DB.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

// DeleteBlogService deletes a blog by ID
func DeleteBlogService(id string) error {
	result := db.DB.Delete(&models.Blog{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateBlogService updates a blog by ID
func UpdateBlogService(id string, updatedBlog models.Blog) (models.Blog, error) {
	var existingBlog models.Blog
	if err := db.DB.First(&existingBlog, id).Error; err != nil {
		return models.Blog{}, err
	}
	existingBlog.Title = updatedBlog.Title
	existingBlog.Description = updatedBlog.Description
	existingBlog.AuthorID = updatedBlog.AuthorID
	existingBlog.CategoryID = updatedBlog.CategoryID
	existingBlog.Image = updatedBlog.Image
	if err := db.DB.Save(&existingBlog).Error; err != nil {
		return models.Blog{}, err
	}
	return existingBlog, nil
}

// func CreateBlog(c *gin.Context) {
// 	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse form"})
// 		return
// 	}

// 	// Get file from form
// 	file, fileHeader, err := c.Request.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get file"})
// 		return
// 	}
// 	defer file.Close()

// 	// Create a directory to save the image if it doesn't exist
// 	imageDir := "./uploads/"
// 	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create image directory"})
// 		return
// 	}

// 	// Save the file
// 	filePath := filepath.Join(imageDir, filepath.Base(fileHeader.Filename))

// 	out, err := os.Create(filePath)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save file"})
// 		return
// 	}
// 	defer out.Close()

// 	if _, err := out.ReadFrom(file); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save file"})
// 		return
// 	}

// 	// Construct image URL or path
// 	imageURL := fmt.Sprintf("image%s", filepath.Base(filePath))
// 	var blog models.Blog
// 	if err := c.ShouldBind(&blog); err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
// 		return
// 	}
// 	blog.Image = imageURL
// 	result := db.DB.Create(&blog)
// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create blog"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, blog)
// }

// func GetBlog(c *gin.Context) {
// 	var blog []models.Blog
// 	result := db.DB.Find(&blog)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create blog"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, blog)

// }

// func DeleteBlog(c *gin.Context) {
// 	id := c.Param("id")
// 	result := db.DB.Delete(&models.Blog{}, id)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete blog"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
// }

// func UpdateBlog(c *gin.Context) {
// 	var blog models.Blog
// 	id := c.Param("id")

// 	if err := c.ShouldBindJSON(&blog); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
// 		return
// 	}
// 	var existingBlog models.Blog
// 	if err := db.DB.First(&existingBlog, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
// 		return
// 	}
// 	existingBlog.Title = blog.Title
// 	existingBlog.Description = blog.Description
// 	existingBlog.AuthorID = blog.AuthorID
// 	existingBlog.CategoryID = blog.CategoryID
// 	existingBlog.Image = blog.Image
// 	if err := db.DB.Save(&existingBlog).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update blog"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, existingBlog)
// }
