package client

import (
	"broker/common"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type MailerClient struct {
	client           *http.Client
	sendMailEndPoint common.EndPoint
}

const mailerBaseURL = "http://mail-service/"

func NewMailerClient() MailerClient {
	return MailerClient{
		client: &http.Client{},
		sendMailEndPoint: common.EndPoint{
			Method: "POST",
			Url:    "mail/send",
		},
	}
}

func (mailerClient *MailerClient) SendMail(request common.SendMailRequest) error {
	sendMailRequest, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return err
	}
	authServiceRequest, err := http.NewRequest(mailerClient.sendMailEndPoint.Method, mailerClient.sendMailEndPoint.GetCompleteURL(mailerBaseURL),
		bytes.NewBuffer(sendMailRequest))
	if err != nil {
		log.Println(err)
		return err
	}
	response, err := mailerClient.client.Do(authServiceRequest)
	if err != nil {
		log.Println(err)
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusAccepted {
		data := make([]byte, 99999)
		response.Body.Read(data)
		log.Println(string(data))
		return errors.New("failed to send mail")
	}
	return nil
}
