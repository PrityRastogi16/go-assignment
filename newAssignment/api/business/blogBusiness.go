package business

import (
	"newAssignment/api/services"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

func CreateBlogBusiness(c *gin.Context) error {
	return services.CreateBlogService(c)
}

func GetBlogsBusiness() ([]models.Blog, error) {
	return services.GetBlogsService()
}

func DeleteBlogBusiness(id string) error {
	return services.DeleteBlogService(id)
}

func UpdateBlogBusiness(id string, updatedBlog models.Blog) (models.Blog, error) {
	return services.UpdateBlogService(id, updatedBlog)
}
