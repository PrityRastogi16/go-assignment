package controller

import (
	"net/http"
	"newAssignment/db"
	"newAssignment/models"
	"time"

	"github.com/gin-gonic/gin"
)

func VerifyEmailController(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Token is required"})
		return
	}

	var verification models.VerificationToken
	result := db.DB.Where("token = ?", token).First(&verification)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid or expired token"})
		return
	}

	if time.Now().After(verification.ExpiresAt) {
		c.JSON(http.StatusGone, gin.H{"message": "Token has expired"})
		return
	}

	var user models.User
	result = db.DB.Where("id = ?", verification.UserID).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Update user status to active
	user.Status = "active" // Assuming you have a `Status` field in the `User` model
	result = db.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user status"})
		return
	}

	// Optionally, delete the verification token
	result = db.DB.Delete(&verification)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete verification token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email successfully verified. You can now log in."})
}
