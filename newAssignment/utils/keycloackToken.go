package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
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
