package routes

import (
	"newAssignment/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	// server.GET("/events", getEvents)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	server.GET("category", listCategories)
	authenticated.DELETE("/category/:id", deleteCategory)
	authenticated.POST("/category", createCategory)
	server.POST("/signup", signup)
	server.POST("/login", Login)

	authenticated.POST("/author", CreateAuthor)
	server.GET("/author", ListAuthors)
	authenticated.DELETE("/author/:id", DeleteAuthor)

	authenticated.POST("/blog", CreateBlog)
	server.GET("/blog", GetBlog)
	authenticated.DELETE("/blog/:id", DeleteBlog)
	authenticated.PUT("/blog/:id", UpdateBlog)
}
