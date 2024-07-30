package main

import (
	"newAssignment/api/routes"
	"newAssignment/db"
	"newAssignment/inits"

	_ "newAssignment/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 	Tag Service API
// @version	2.0
// @description A Tag service API in Go using Gin framework
// @host 	localhost:2002
// @BasePath /
// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

func InitSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
func main() {

	inits.LoadEnv()
	db.InitDB()
	server := gin.Default()
	server.Static("/uploads", "./uploads")
	routes.RegisterRouter(server)
	InitSwagger(server)
	server.Run(":2002")
}
