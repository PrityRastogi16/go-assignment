package routes

import (
	"fmt"
	"net/http"
	"newAssignment/db"
	"newAssignment/models"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse form"})
		return
	}

	// Get file from form
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get file"})
		return
	}
	defer file.Close()

	// Create a directory to save the image if it doesn't exist
	imageDir := "./uploads/"
	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create image directory"})
		return
	}

	// Save the file
	filePath := filepath.Join(imageDir, filepath.Base(fileHeader.Filename))

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save file"})
		return
	}
	defer out.Close()

	if _, err := out.ReadFrom(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save file"})
		return
	}

	// Construct image URL or path
	imageURL := fmt.Sprintf("image%s", filepath.Base(filePath))
	var blog models.Blog
	if err := c.ShouldBind(&blog); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	blog.Image = imageURL
	result := db.DB.Create(&blog)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create blog"})
		return
	}
	c.JSON(http.StatusCreated, blog)
}

func GetBlog(c *gin.Context) {
	var blog []models.Blog
	result := db.DB.Find(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create blog"})
		return
	}
	c.JSON(http.StatusOK, blog)

}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	result := db.DB.Delete(&models.Blog{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	id := c.Param("id")

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	var existingBlog models.Blog
	if err := db.DB.First(&existingBlog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		return
	}
	existingBlog.Title = blog.Title
	existingBlog.Description = blog.Description
	existingBlog.AuthorID = blog.AuthorID
	existingBlog.CategoryID = blog.CategoryID
	existingBlog.Image = blog.Image
	if err := db.DB.Save(&existingBlog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, existingBlog)
}
