package controller

import (
	"net/http"
	"newAssignment/api/business"
	"newAssignment/models"

	"github.com/gin-gonic/gin"
)

// SignupController handles user registration
// @Summary User Signup
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User Details"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /signup [post]
func CreateUserController(c *gin.Context) {
	// Call the business logic layer
	err := business.CreateUserBusiness(c)
	if err != nil {
		c.JSON(err.StatusCode, ErrorResponse{err.Message})
		return
	}

	c.JSON(200, SuccessResponse{"User created successfully"})
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// LoginController handles user login
// @Summary User Login
// @Description Authenticate a user and return an access token
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User Credentials"
// @Success 200 {object} SuccessTokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func LoginController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid request"})
		return
	}

	tokenResponse, err := business.LoginBusiness(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

// LogoutController handles user logout
// @Summary Logout a user
// @Description Invalidate the user's session token
// @Tags User
// @Security bearerToken
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /logout [post]
func LogoutController(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Missing token"})
		return
	}

	err := business.LogoutBusiness(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Logout Successfully")
}
