package client

import (
	"broker/common"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type AuthClient struct {
	client         *http.Client
	singupEndpoint common.EndPoint
	singinEndpoint common.EndPoint
}

const authBaseURL = "http://authentication-service/"

func NewAuthClient() AuthClient {
	return AuthClient{
		client: &http.Client{},
		singupEndpoint: common.EndPoint{
			Method: "POST",
			Url:    "signup",
		},
		singinEndpoint: common.EndPoint{
			Method: "POST",
			Url:    "signin",
		},
	}
}

func (authClient *AuthClient) SignUp(request common.AuthRequest) common.AuthResponse {
	authResponse := common.AuthResponse{
		Error:   true,
		Message: "Internal Server Error",
		Status:  http.StatusInternalServerError,
	}
	authRequest, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return authResponse
	}
	authServiceRequest, err := http.NewRequest(authClient.singupEndpoint.Method, authClient.singinEndpoint.GetCompleteURL(authBaseURL),
		bytes.NewBuffer(authRequest))
	if err != nil {
		log.Println(err)
		authResponse.Message = err.Error()
		return authResponse
	}
	response, err := authClient.client.Do(authServiceRequest)
	if err != nil {
		log.Println(err)
		authResponse.Message = err.Error()
		return authResponse
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&authResponse)
	log.Printf("Got response from auth service %+v", authResponse)
	if err != nil {
		authResponse.Error = true
		authResponse.Message = "unexpected error"
		authResponse.Status = http.StatusInternalServerError
		return authResponse
	} else if authResponse.Error || response.StatusCode == http.StatusUnauthorized {
		return authResponse
	} else if response.StatusCode != http.StatusOK {
		log.Println(authResponse)
		authResponse.Message = "failed calling auth service"
		return authResponse
	}

	return authResponse
}
