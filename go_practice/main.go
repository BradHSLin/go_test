package main

import (
	"fmt"
	"go_practice/db"
	"go_practice/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	// fmt.Println(server)
	fmt.Println("Running!")

	server.Run(":8080")

}
