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
	// server.POST("/login", login)

	server.POST("/author", CreateAuthor)
	server.GET("/author", ListAuthors)
	server.DELETE("/author/:id", DeleteAuthor)
}
