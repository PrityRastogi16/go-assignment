package business

import (
	"newAssignment/api/services"
	"newAssignment/models"
)

func CreateAuthorBusiness(author models.Author) error {
	return services.CreateAuthorService(author)
}

// ListAuthorsBusiness retrieves all authors and handles any business logic
func ListAuthorsBusiness() ([]models.Author, error) {
	return services.ListAuthorsService()
}

// DeleteAuthorBusiness deletes an author and handles any business logic
func DeleteAuthorBusiness(id string) error {
	return services.DeleteAuthorService(id)
}
