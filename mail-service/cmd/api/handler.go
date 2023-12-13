package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	var request = SendMailRequest{}
	err := readJson(w, r, &request)
	if err != nil {
		log.Println("error in reading sendmail request ", err)
		errorJson(w, err, http.StatusInternalServerError)
		return
	}
	msg := Message{
		From:    request.From,
		To:      request.To,
		Subject: request.Subject,
		Data:    request.Message,
	}
	err = app.Mailer.SendMail(msg)
	if err != nil {
		log.Println("error in reading sendmail request ", err)
		errorJson(w, err, http.StatusInternalServerError)
		return
	}
	response := SendMailResponse{
		Error:     false,
		Message:   "Success!!",
		MailRefId: "",
	}
	writeJson(w, http.StatusAccepted, response)
}
