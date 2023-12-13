package client

import (
	"broker/common"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type LoggerClient struct {
	client           *http.Client
	writeLogEndPoint common.EndPoint
}

const (
	logBaseURL = "http://log/"
	writeLog   = "log"
)

func NewLoggerClient() LoggerClient {
	return LoggerClient{
		client: &http.Client{},
		writeLogEndPoint: common.EndPoint{
			Method: "POST",
			Url:    writeLog,
		},
	}
}

func (l *LoggerClient) log(request common.LogRequest) common.LogResponse {
	var logResponse = common.LogResponse{
		Recorded: false,
		Message:  "Internal Server Error",
		Status:   http.StatusInternalServerError,
	}
	authRequest, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return logResponse
	}
	authServiceRequest, err := http.NewRequest(l.writeLogEndPoint.Method, l.writeLogEndPoint.GetCompleteURL(logBaseURL),
		bytes.NewBuffer(authRequest))
	if err != nil {
		log.Println(err)
		logResponse.Message = err.Error()
		return logResponse
	}
	response, err := l.client.Do(authServiceRequest)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println(err)
		logResponse.Message = err.Error()
		return logResponse
	}
	defer response.Body.Close()
	logResponse.Recorded = true
	logResponse.Message = ""
	logResponse.Status = http.StatusOK
	return logResponse
}
