package main

import (
	"broker/common"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (app *Config) broker(w http.ResponseWriter, r *http.Request) {
	payload := common.BrokerResponse{Error: false, Message: "Success"}
	var requestBody = common.BrokerRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Got error while reading error %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("got request %s,%s", requestBody.Message, requestBody.Sender)
	writeJson(w, 200, payload)
}

func (app *Config) signin(w http.ResponseWriter, r *http.Request) {
	var signUpRequest = common.SignUpRequest{}
	readJson(w, r, &signUpRequest)
	authResponse := app.AuthClient.SignUp(common.AuthRequest{
		Email:    signUpRequest.Email,
		Password: signUpRequest.Password,
	})
	brokerResponse := common.BrokerResponse{
		Error:   false,
		Message: "Sucess",
	}
	if authResponse.Error {
		errorJson(w, errors.New(authResponse.Message), authResponse.Status)
		return
	}
	brokerResponse.Data = authResponse.Data
	writeJson(w, http.StatusOK, brokerResponse)
}

func (app *Config) signup(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) log(w http.ResponseWriter, r *http.Request) {
	var loggingRequest = common.BrokerLogRequest{}
	readJson(w, r, &loggingRequest)
	log.Println("producing log event to rabbitmq")
	app.produceLogEvent(w, loggingRequest)
	log.Println("done producing log event to rabbitmq")
	log.Println("pushing log event to rpc")
	app.LogRPCClient.SendLog(loggingRequest)
	log.Println("done pushing log event to rpc")
	log.Println("pushing log event to GRPC")
	app.LogGRPCClient.SendLog(loggingRequest)
	log.Println("done pushing log event to GRPC")
	// loggingResponse := app.LoggerClient.log(LogRequest{
	// 	Name: loggingRequest.Name,
	// 	Data: loggingRequest.Data,
	// })
	// if loggingResponse.Status != http.StatusOK {
	// 	errorJson(w, errors.New(loggingResponse.Message), loggingResponse.Status)
	// 	return
	// }
	// brokerResponse := BrokerResponse{
	// 	Error:   false,
	// 	Message: "Success",
	// 	Data: BrokerLogResponse{
	// 		Recorded: true,
	// 	},
	// }
	// writeJson(w, http.StatusOK, brokerResponse)
}

func (app *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	var brokerRequest = common.BrokerSendMailRequest{}
	readJson(w, r, &brokerRequest)
	err := app.MailerClient.SendMail(common.SendMailRequest{
		From:    brokerRequest.From,
		To:      brokerRequest.To,
		Subject: brokerRequest.Subject,
		Message: brokerRequest.Message,
	})
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}
	brokerResponse := common.BrokerResponse{
		Error:   false,
		Message: "Success",
	}
	writeJson(w, http.StatusOK, brokerResponse)
}

func (app *Config) produceLogEvent(w http.ResponseWriter, request common.BrokerLogRequest) {
	data, _ := json.Marshal(&request)
	logEvent := common.LogEvent{
		Name: "log",
		Data: string(data),
	}
	data, _ = json.Marshal(&logEvent)
	err := app.LogProducer.Produce(string(data), "log.WARNING")
	if err != nil {
		log.Println("failed to publish ", err)
		errorJson(w, err, http.StatusInternalServerError)
		return
	}
	brokerResponse := common.BrokerResponse{
		Error:   false,
		Message: "Logged into queue!!",
	}
	writeJson(w, http.StatusOK, brokerResponse)
}
