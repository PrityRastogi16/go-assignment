package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Credentials struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}

type KeycloakUser struct {
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Enabled     bool          `json:"enabled"`
	Credentials []Credentials `json:"credentials"`
}


func GetToken(grantType, clientID, clientSecret, username, password string) (string, error) {
	// Keycloak token endpoint
	tokenURL := "http://localhost:8080/realms/master/protocol/openid-connect/token"
	data := map[string]string{
		"grant_type":    grantType,
		"client_id":     clientID,
		"client_secret": clientSecret,
		"username":      username,
		"password":      password,
	}
	form := url.Values{}
	for key, value := range data {
		form.Set(key, value)
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func GetKeycloakAuthToken(email, password string) (string, error) {
	baseURL := "http://localhost:8080"
	authURL := baseURL + "/realms/master/protocol/openid-connect/auth?client_id=client-credentials-test-client&redirect_uri=http://localhost:2002/&response_type=code&scope=openid"
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// Step 1: Get the login form
	resp, err := client.Get(authURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Step 2: Extract action URL using regex
	re := regexp.MustCompile(`action="([^"]+)"`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("action URL not found")
	}
	actionURL := matches[1]
	actionURL = strings.Replace(actionURL, "&amp;", "&", -1)

	// Step 3: Submit the login form
	formData := url.Values{
		"username": {email},
		"password": {password},
	}
	loginResp, err := client.PostForm(actionURL, formData)
	if err != nil {
		return "", err
	}
	defer loginResp.Body.Close()

	loginBody, err := io.ReadAll(loginResp.Body)
	if err != nil {
		return "", err
	}

	// Return the response body as a string for simplicity
	return string(loginBody), nil

}
func CreateKeycloakUser(adminToken, email, password string) bool {
	apiEndpoint := "/admin/realms/master/users"
	apiURL := KeyCloakBaseUrl + apiEndpoint

	// Generate random password to store in keycloak
	// randomPassword := utils.GenerateRandomPassword()

	user := KeycloakUser{
		Username: email,
		Email:    email,
		Enabled:  true,
		Credentials: []Credentials{
			{
				Type:      "password",
				Value:     password,
				Temporary: false,
			},
		},
	}
	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println("ERROR: failed to marshal user data", err)
		return false
	}
	// Make POST api call to keycloak to create new user in master realm
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("ERROR: failed to create request", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: failed to send request", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("ERROR: failed to create user", resp)
		return false
	}

	fmt.Println("SUCCESS: User created successfully in keycloak", resp)

	return true
}

func GetKeycloakUser(adminToken, email string) (string, error) {
	apiEndpoint := "/admin/realms/master/users"
	apiURL := KeyCloakBaseUrl + apiEndpoint

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("ERROR: failed to create request", err)
		return "", err
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: failed to make request", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR: failed to get user", resp)
		return "", errors.New("failed to get user")
	}

	var users []struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Println("ERROR: failed to decode response", err)
		return "", err
	}

	if len(users) == 0 {
		return "", fmt.Errorf("user not found")
	}
	return users[0].ID, nil

}

func GetKeyclaokUserInfo(token string) (map[string]interface{}, error) {

	apiEndpoint := "/realms/master/protocol/openid-connect/userinfo"
	apiURL := KeyCloakBaseUrl + apiEndpoint

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("ERROR: failed to create request", err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: failed to make request", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR: failed to get user info", resp)
		return nil, errors.New("failed to get user")
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("ERROR: failed to decode response body:", err)
		return nil, err
	}

	return result, nil
}
