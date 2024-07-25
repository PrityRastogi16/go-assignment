package routes

import (
	"net/http"

	"newAssignment/models"
	"newAssignment/utils"

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
		return
	}
	err = utils.TriggerEmailWorkflow(user.Email, "Welcome to Our Service", "Thank you for registering with us!")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send registration email"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created Succesfully"})
}

// func signup(context *gin.Context) {
//     var user models.User
//     err := context.ShouldBindJSON(&user)
//     if err != nil {
//         context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
//         return
//     }
//     err = user.Save()
//     if err != nil {
//         context.JSON(http.StatusInternalServerError, gin.H{"message": "User creation failed"})
//         return
//     }

//     // Trigger Temporal workflow to send a registration email
//     err = utils.TriggerEmailWorkflow(user.Email, "Welcome to Our Service", "Thank you for registering with us!")
//     if err != nil {
//         context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send registration email"})
//         return
//     }

//     context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
// }

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
