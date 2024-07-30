package services

import (
	"newAssignment/db"
	"newAssignment/models"
)

func CreateAuthorService(author models.Author) error {
	result := db.DB.Create(&author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ListAuthorsService() ([]models.Author, error) {
	var authors []models.Author
	result := db.DB.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func DeleteAuthorService(id string) error {
	result := db.DB.Delete(&models.Author{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func CreateAuthor(c *gin.Context) {
// 	var author models.Author
// 	if err := c.ShouldBind(&author); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
// 		return
// 	}
// 	result := db.DB.Create(&author)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create author"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, author)
// }

// type ErrorResponse struct {
// 	Error string `json:"error"`
// }
// type AuthorResponse struct {
// 	ID        uint       `json:"id"`
// 	CreatedAt time.Time  `json:"createdAt"`
// 	UpdatedAt time.Time  `json:"updatedAt"`
// 	DeletedAt *time.Time `json:"deletedAt,omitempty"`
// 	Name      string     `json:"name"`
// 	Bio       string     `json:"bio"`
// }

// ListAuthors returns a list of all authors from the database
// @Summary List all authors
// @Description Returns a list of all authors from the database
// @Tags Author
// @Produce json
// @Success 200
// @Failure 500
// @Router /author [get]
// func ListAuthors(c *gin.Context) {
// 	var authors []models.Author
// 	result := db.DB.Find(&authors)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to fetch authors"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, authors)
// }
// func DeleteAuthor(c *gin.Context) {
// 	id := c.Param("id")
// 	result := db.DB.Delete(&models.Author{}, id)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete author"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
// }
