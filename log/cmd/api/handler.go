package main

import (
	"errors"
	"log"
	"log/data"
	"net/http"
)

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var request LogRequest
	err := readJson(w, r, &request)
	if err != nil {
		log.Panicln("Error reading request ", err)
		errorJson(w, errors.New("Invalid request"), http.StatusBadRequest)
		return
	}
	log.Printf("recording log for %s, %s", request.Name, request.Data)
	event := data.LogEntry{
		Name: request.Name,
		Data: request.Data,
	}
	err = app.LogRepo.Insert(event)
	if err != nil {
		log.Panicln("Error reading request ", err)
		errorJson(w, errors.New("Failed to record"), http.StatusInternalServerError)
		return
	}
	response := LogResponse{
		Recorded: true,
	}
	writeJson(w, 200, response)
}
