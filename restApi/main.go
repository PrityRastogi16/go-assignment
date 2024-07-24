package main

import (
	"github.com/abcom/restApi/db"
	"github.com/abcom/restApi/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRouter(server)
	server.Run(":8080") // localhost
}
