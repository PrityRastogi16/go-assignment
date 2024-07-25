package routes

import (
	"net/http"
	"newAssignment/db"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	result := db.DB.Create(&blog)
	if result.Error != nil {
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
