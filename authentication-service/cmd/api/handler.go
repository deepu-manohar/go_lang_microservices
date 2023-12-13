package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Config) signin(w http.ResponseWriter, r *http.Request) {
	var request = AuthRequest{}
	err := readJson(w, r, &request)
	if err != nil {
		errorJson(w, err, http.StatusBadRequest)
		return
	}
	app.logRequest(request)
	log.Printf("Got request %s , %s", request.Email, request.Password)
	userDto, err := app.authService.GetUser(request.Email, request.Password)
	if err != nil {
		errorJson(w, err, http.StatusUnauthorized)
		return
	}
	writeJson(w, 200,
		AuthResponse{
			Error:   false,
			Message: "",
			Data: UserResponse{
				Name:     userDto.Name,
				Email:    userDto.Email,
				IsActive: userDto.IsActive,
			},
		})

}

func (app *Config) signup(w http.ResponseWriter, r *http.Request) {
}

func (app *Config) logRequest(data any) {
	logData := fmt.Sprintf("%#v", data)
	log.Printf("Going to log %s", logData)
	app.loggerService.Log(logData)
}
