package routes

import (
	"newAssignment/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	// server.GET("/events", getEvents)
	server.GET("category", listCategories)
	server.DELETE("/category/:id", deleteCategory)
	server.POST("/category", createCategory)
	server.DELETE("/events/:id/register", middlewares.Authenticate)
	server.POST("/signup", signup)
	server.POST("/login", Login)

	server.POST("/author", CreateAuthor)
	server.GET("/author", ListAuthors)
	server.DELETE("/author/:id", DeleteAuthor)

	server.POST("/blog", CreateBlog)
	server.GET("/blog", GetBlog)
	server.DELETE("/blog/:id", DeleteBlog)
	server.PUT("/blog/:id", UpdateBlog)
}
