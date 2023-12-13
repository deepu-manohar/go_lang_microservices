package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type LogRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

const logURL = "http://log/log"
const authLoggerName = "AuthService"

type LoggerService struct {
	client *http.Client
}

func NewLoggerService() LoggerService {
	return LoggerService{
		client: &http.Client{},
	}
}

func (l *LoggerService) Log(data string) {
	var request = LogRequest{
		Name: authLoggerName,
		Data: data,
	}
	logRequest, err := json.Marshal(request)
	if err != nil {
		log.Println("failed to log request ", err)
		return
	}
	serviceRequest, err := http.NewRequest("POST", logURL,
		bytes.NewBuffer(logRequest))
	if err != nil {
		log.Println("failed to log request ", err)
		return
	}
	response, err := l.client.Do(serviceRequest)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println("failed while logging request ", err)
	}
}
