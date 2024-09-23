package routes

import (
	"fmt"
	"go_practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"could not save message": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	user, err := models.GetUserByID(userId)
	fmt.Println(user)
	fmt.Println(err)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, user)
}

func updateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	user, err := models.GetUserByID(userId)

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"could not GetUserByID": err})
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"could not GetUserByID ShouldBindJSON": err})
		return
	}

	updatedUser.ID = user.ID

	err = updatedUser.Update()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"could not update user": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})

}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	user, err := models.GetUserByID(userId)
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"could not GetUserByID": err})
		return
	}

	err = user.Delete()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"could not delete": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user delete successfully"})
}
