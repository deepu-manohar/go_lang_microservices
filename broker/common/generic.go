package common

type BrokerResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type BrokerRequest struct {
	Message string `json:"message"`
	Sender  string `json:"sender"`
}

type BrokerLogRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type BrokerSendMailRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type BrokerLogResponse struct {
	Recorded bool `json:"recorded"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive int    `json:"isActive"`
}

type LogEvent struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
