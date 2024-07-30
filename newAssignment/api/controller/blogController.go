package controller

import (
	"net/http"
	"newAssignment/api/business"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

// CreateBlogController handles the craetion of Blog with image
// @Summary Create Blog Post
// @Description Create a new blog post
// @Tags Blogs
// @Accept multipart/form-data
// @Produce json
// @param image formData file true "Blog Image"
// @Param blog body models.Blog true "Blog Details"
// @Success 201 {object} models.Blog
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /blog [post]
// @Security bearerToken
func CreateBlogController(c *gin.Context) {
	if err := business.CreateBlogBusiness(c); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to create blog"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Blog created successfully"})
}

// GetBlogsController retrieves all the blogs
// @Summary Get all blogs
// @Description Get Blogs
// @Tags Blogs
// @Produce json
// @Success 200 {array} models.Blog
// @Failure 500 {object} ErrorResponse
// @Router /blog [get]
// @Security bearerToken
func GetBlogsController(c *gin.Context) {
	blogs, err := business.GetBlogsBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to fetch blogs"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// DeleteBlogController handles the deletion of a blog post
// @Summary Delete a blog post
// @Description Deletes a blog post by its ID
// @Tags Blogs
// @Param id path string true "Blog ID"
// @Success 200
// @Failure 500 {object} ErrorResponse
// @Router /blog/{id} [delete]
// @Security bearerToken
func DeleteBlogController(c *gin.Context) {
	id := c.Param("id")
	if err := business.DeleteBlogBusiness(id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to delete blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// UpdateBlogController handles the updation of a blog post
// @Summary Update a blog post
// @Description Update the details of existing blog
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Param blog body models.Blog true "Updated Blog Details"
// @Success 200 {object} models.Blog
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /blog/{id} [put]
// @Security bearerToken
func UpdateBlogController(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid data"})
		return
	}
	updatedBlog, err := business.UpdateBlogBusiness(id, blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to update blog"})
		return
	}
	c.JSON(http.StatusOK, updatedBlog)
}
