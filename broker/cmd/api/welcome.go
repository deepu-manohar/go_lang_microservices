package main

import (
	"broker/common"
	"encoding/json"
	"log"
	"net/http"
)

func (app *Config) welcome(w http.ResponseWriter, r *http.Request) {
	message := common.BrokerResponse{
		Error:   false,
		Message: "welcome",
	}
	out, _ := json.Marshal(message)
	log.Printf("Got request to welcome resource")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
