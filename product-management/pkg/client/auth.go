package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type AuthUserRequestBody struct {
	Username string `json:"email"`
	Password string `json:"password"`
	ApiToken string `json:"authenticationApiKey"`
}

func NewAuthUserRequest(username string, password string, apiToken string) ([]byte, error) {
	request := AuthUserRequestBody{
		Username: username,
		Password: password,
		ApiToken: apiToken,
	}

	return json.Marshal(request)
}

func Authenticate(username string, password string) (*http.Response, error) {

	body, err := NewAuthUserRequest(username, password, os.Getenv("AUTHENTICATION_API_KEY"))
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	authURL := fmt.Sprintf("%s:%s",
		os.Getenv("AUTH_URL"),
		"/user/token",
	)

	httpClient := &http.Client{}
	return httpClient.Post(
		authURL,
		"application/json",
		bytes.NewBuffer(body),
	)
}

func ValidateToken(token string) (*http.Response, error) {

	authURL := fmt.Sprintf("%s:%s",
		os.Getenv("AUTH_URL"),
		"/user/validation",
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", authURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("token", token)
	return client.Do(req)
}
