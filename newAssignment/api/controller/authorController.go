package controller

import (
	"net/http"
	"newAssignment/api/business"
	"newAssignment/models"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorResponse defines the structure of error responses
type ErrorResponse struct {
	Error string `json:"error"`
}
type SuccessResponse struct {
	Success string `json:"error"`
}
type SuccessTokenResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

// AuthorResponse defines the structure of author responses
type AuthorResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Name      string     `json:"name"`
	Bio       string     `json:"bio"`
}

// CreateAuthorController handles the creation of an author
// @Summary Create an author
// @Description Creates a new author
// @Tags Author
// @Accept json
// @Produce json
// @Param author body models.Author true "Author Details"
// @Success 201 {object} models.Author
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /author [post]
// @Security bearerToken
func CreateAuthorController(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid data"})
		return
	}

	err := business.CreateAuthorBusiness(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to create author"})
		return
	}

	c.JSON(http.StatusCreated, author)
}

// ListAuthorsController retrieves and returns all authors
// @Summary List all authors
// @Description Fetches a list of all authors from the database
// @Tags Author
// @Produce json
// @Success 200 {array} models.Author
// @Failure 500 {object} ErrorResponse
// @Router /author [get]
// @Security bearerToken
func ListAuthorsController(c *gin.Context) {
	authors, err := business.ListAuthorsBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to fetch authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// DeleteAuthorController handles the deletion of an author
// @Summary Delete an author
// @Description Deletes an author by its ID
// @Tags Author
// @Param id path string true "Author ID"
// @Success 200
// @Failure 500 {object} ErrorResponse
// @Router /author/{id} [delete]
// @Security bearerToken
func DeleteAuthorController(c *gin.Context) {
	id := c.Param("id")
	err := business.DeleteAuthorBusiness(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to delete author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
