package main

import (
	"gokul-mylsami/rest-api/db"
	"gokul-mylsami/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
