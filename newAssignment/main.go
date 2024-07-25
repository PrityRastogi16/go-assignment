package main

import (
	"newAssignment/db"
	"newAssignment/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRouter(server)
	server.Run(":2002") // localhost
}
