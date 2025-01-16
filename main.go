package main

import (
	
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
	"example.com/rest-api/routes"
)


func main() {
	
	server := gin.Default()

	db.InitDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
