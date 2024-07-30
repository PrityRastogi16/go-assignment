package business

import (
	"fmt"
	"newAssignment/models"

	"newAssignment/api/services"

	"github.com/gin-gonic/gin"
)

type BusinessError struct {
	StatusCode int
	Message    string
}

// CreateUserBusiness processes the user creation logic
func CreateUserBusiness(c *gin.Context) *BusinessError {
	var user models.User

	// Bind JSON input to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		return &BusinessError{StatusCode: 400, Message: "Invalid input"}
	}

	// Call the service layer to create the user
	err := services.CreateUserService(user)
	fmt.Println(err)
	if err != nil {
		return &BusinessError{StatusCode: 500, Message: "Failed to create user"}
	}

	return nil
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func LoginBusiness(user models.User) (services.TokenResponse, error) {
	return services.LoginService(user)
}
func LogoutBusiness(token string) error {
	return services.LogoutService(token)
}
