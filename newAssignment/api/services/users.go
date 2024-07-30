package services

import (
	"fmt"
	"newAssignment/db"
	"newAssignment/models"
	"newAssignment/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUserService(user models.User) error {
	// Check if email already exists
	var existingUser models.User
	result := db.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return fmt.Errorf("user with this email already exists")
	}
	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword
	// Get Keycloak admin token
	keyCloakAdminTokenResp, err := utils.GetKeycloakAccessToken("prity", "prity")
	if err != nil {
		return fmt.Errorf("failed to get Keycloak admin token: %w", err)
	}
	accessToken := keyCloakAdminTokenResp.AccessToken

	// Create Keycloak user
	status := utils.CreateKeycloakUser(accessToken, user.Email, user.Password)
	if !status {
		return fmt.Errorf("failed to create Keycloak user")
	}
	result = db.DB.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to store user in database: %w", result.Error)
	}

	// Send registration email
	verificationToken := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)
	verification := models.VerificationToken{
		UserID:    user.ID,
		Token:     verificationToken,
		ExpiresAt: expiresAt,
	}
	fmt.Println("11111111cxcvbn,")
	result = db.DB.Create(&verification)
	if result.Error != nil {
		return fmt.Errorf("failed to store verification token: %w", result.Error)
	}
	verificationLink := fmt.Sprintf("http://localhost:2002/verify?token=%s", verificationToken)
	emailBody := fmt.Sprintf("Thank you for registering with us! Click the link below to verify your email address:\n\n%s", verificationLink)
	err = utils.TriggerEmailWorkflow(user.Email, "Welcome to Our Service", emailBody)
	if err != nil {
		return fmt.Errorf("failed to send registration email: %w", err)
	}

	return nil
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func LoginService(user models.User) (TokenResponse, error) {
	var existingUser models.User

	// Check if the user exists
	result := db.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return TokenResponse{}, fmt.Errorf("user not registered, please signup")
		}
		return TokenResponse{}, fmt.Errorf("failed to check user existence: %w", result.Error)
	}

	// Validate password
	passwordIsValid := utils.CheckPasswordHash(user.Password, existingUser.Password)
	if !passwordIsValid {
		return TokenResponse{}, fmt.Errorf("invalid password")
	}

	// Generate token
	keyCloakAdminTokenResp, err := utils.GetKeycloakAccessToken(user.Email, user.Password)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("failed to get Keycloak admin token: %w", err)
	}

	accessToken := keyCloakAdminTokenResp.AccessToken
	refreshToken := keyCloakAdminTokenResp.RefreshToken

	return TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func LogoutService(token string) error {
	// Call Keycloak to invalidate the token
	err := utils.InvalidateKeycloakToken(token)
	if err != nil {
		return fmt.Errorf("failed to logout user: %w", err)
	}
	return nil
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
