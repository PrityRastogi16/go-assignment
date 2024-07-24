package routes

import (
	"github.com/abcom/restApi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authenticate)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
