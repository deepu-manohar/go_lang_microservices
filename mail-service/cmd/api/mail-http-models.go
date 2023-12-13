package main

type SendMailRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type SendMailResponse struct {
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	MailRefId string `json:"mail_ref_id,omitempty"`
}
