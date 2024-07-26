package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"newAssignment/db"
	"newAssignment/models"
	"newAssignment/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateKeycloakUser(token string, user models.User) error {
	url := "http://localhost:8080/admin/realms/master/users"

	userData := map[string]interface{}{
		"username":      user.Email,
		"email":         user.Email,
		"enabled":       true,
		"emailVerified": false,
		"firstName":     "FirstName",
		"lastName":      "LastName",
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     "user-password",
				"temporary": false,
			},
		},
	}

	jsonData, err := json.Marshal(userData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create user in keycloak, status code: %d", resp.StatusCode)
	}

	return nil
}

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	token, _ := utils.GenerateToken(user.Email, user.ID)
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	result := db.DB.Create(&user)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User creation failed"})
		return
	}
	loginLink := fmt.Sprintf("https://your-app.com/login?token=%s", token)
	grantType := "password"
	clientID := "client-credentials-test-client"
	clientSecret := "PtygUYw4wU9zhwaIr60jDJArxH9TjZVA"
	username := "prity"
	password := "prity"
	keyCloakToken, _ := utils.GetToken(grantType, clientID, clientSecret, username, password)
	fmt.Println("*****", keyCloakToken, "******")
	err = CreateKeycloakUser(keyCloakToken, user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user in Keycloak"})
		return
	}

	// Email body with login link
	emailBody := fmt.Sprintf("Thank you for registering with us! , Click the link below to login:\n%s\n", loginLink)
	err = utils.TriggerEmailWorkflow(user.Email, "Welcome to Our Service", emailBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send registration email"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created Succesfully"})
}

func Login(c *gin.Context) {
	var user models.User
	var existingUser models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not paser request body."})
		return
	}

	// Check if email already exist
	result := db.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not registered, please signup!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		}
		return
	}
	// Validate password
	passwordIsValid := utils.CheckPasswordHash(user.Password, existingUser.Password)

	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password."})
		return
	}
	output, _ := utils.GetKeycloakAuthToken(user.Email, user.Password)
	fmt.Println(output)
	// Generate token
	token, err := utils.GenerateToken(existingUser.Email, existingUser.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

}

// func login(context *gin.Context) {
// 	var loginRequest struct {
// 		Email    string `json:"email" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}
// 	if err := context.ShouldBindJSON(&loginRequest); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
// 		return
// 	}

// 	var user models.User
// 	if err := db.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
// 		return
// 	}

// 	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
// 		return
// 	}

// 	token, err := utils.GenerateToken(user.Email, int64(user.ID))
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
// }

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

// func login(context *gin.Context) {
// 	var user models.User
// 	err := context.ShouldBindJSON(&user)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
// 		return
// 	}

// 	err = user.ValidateCredentials()
// 	if err != nil {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not validate credentials"})
// 		return
// 	}
// 	token, err := utils.GenerateToken(user.Email, user.ID)
// 	if err != nil {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not athenticate"})
// 	}
// 	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
// }
