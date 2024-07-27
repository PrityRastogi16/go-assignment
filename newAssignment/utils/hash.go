package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET_KEY = "prity"

const KeyCloakBaseUrl string = "http://localhost:8080"

type KeyCloakAdminTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateToken(email string, userId uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(JWT_SECRET_KEY))
}

func GetKeycloakAccessToken(username, password string) (*KeyCloakAdminTokenResponse, error) {
	apiEndpoint := "/realms/master/protocol/openid-connect/token"
	apiURL := KeyCloakBaseUrl + apiEndpoint
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "client-credentials-test-client")
	data.Set("client_secret", "PtygUYw4wU9zhwaIr60jDJArxH9TjZVA")
	data.Set("username", username)
	data.Set("password", password)

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// Set the content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp KeyCloakAdminTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	// Return the response body as a string
	return &tokenResp, nil
}
