package routes

import (
	"net/http"

	"github.com/abcom/restApi/models"
	"github.com/abcom/restApi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{"message": "User creation failed"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created Succesfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not validate credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not athenticate"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
