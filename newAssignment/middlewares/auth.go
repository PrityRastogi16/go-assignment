package middlewares

import (
	"fmt"
	"net/http"
	"newAssignment/db"
	"newAssignment/models"
	"newAssignment/utils"

	"github.com/gin-gonic/gin"
)

// func Authenticate(context *gin.Context) {
// 	token := context.Request.Header.Get("Authorization")
// 	if token == "" {
// 		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token not found"})
// 		return
// 	}
// 	userId, err := utils.VerifyToken(token)
// 	if err != nil {
// 		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
// 		return
// 	}
// 	context.Set("userId", userId)
// 	context.Next()
// }

func Authenticate(c *gin.Context) {
	var user models.User

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Validate token
	// userId, err := utils.VerifyToken(token)

	// Make api call to keycloak to validate token
	resp, err := utils.GetKeyclaokUserInfo(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token, login again!"})
		return
	}

	email, ok := resp["email"].(string)
	if !ok {
		fmt.Println("ERROR: email field not found or is not a string")
	}

	userDetails := db.DB.Where("email = ?", email).Find(&user)
	if userDetails.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User with this email does not exists in user table.",
		})
		return
	}
	// Set in request object
	c.Set("userId", user.ID)
	c.Next()
}
