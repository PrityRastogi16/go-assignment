package routes

import (
	"net/http"
	"newAssignment/db"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBind(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	result := db.DB.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create author"})
		return
	}
	c.JSON(http.StatusCreated, author)
}

// List all authors
func ListAuthors(c *gin.Context) {
	var authors []models.Author
	result := db.DB.Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch authors"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// Delete an author
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	result := db.DB.Delete(&models.Author{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
