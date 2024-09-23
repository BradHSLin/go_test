package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/users", createUser)
	server.GET("/users/", getUsers)
	server.GET("/users/:id", getUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id", deleteUser)
}
