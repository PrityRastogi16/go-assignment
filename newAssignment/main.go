package main

import (
	"newAssignment/db"
	"newAssignment/inits"
	"newAssignment/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	inits.LoadEnv()
	db.InitDB()
	server := gin.Default()
	routes.RegisterRouter(server)
	server.Run(":2002")

}