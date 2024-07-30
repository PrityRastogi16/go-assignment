package routes

import (
	"newAssignment/api/controller"
	"newAssignment/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	// server.GET("/events", getEvents)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/category", controller.CreateCategoryController)
	authenticated.GET("/categories", controller.ListCategoriesController)
	authenticated.DELETE("/category/:id", controller.DeleteCategoryController)
	server.POST("/signup", controller.CreateUserController)
	server.POST("/login", controller.LoginController)
	server.POST("/logout", controller.LogoutController)
	server.GET("/verify", controller.VerifyEmailController)
	authenticated.POST("/author", controller.CreateAuthorController)
	authenticated.GET("/author", controller.ListAuthorsController)
	authenticated.DELETE("/author/:id", controller.DeleteAuthorController)
	authenticated.POST("/blog", controller.CreateBlogController)
	authenticated.GET("/blog", controller.GetBlogsController)
	authenticated.DELETE("/blog/:id", controller.DeleteBlogController)
	authenticated.PUT("/blog/:id", controller.UpdateBlogController)
}
