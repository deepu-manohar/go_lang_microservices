package common

type EndPoint struct {
	Method string
	Url    string
}

func (endPoint *EndPoint) GetCompleteURL(baseUrl string) string {
	return baseUrl + endPoint.Url
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
	Data    any    `json:"data,omitempty"`
}
type LogResponse struct {
	Recorded bool   `json:"recorded"`
	Status   int    `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
}

type LogRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type SendMailRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
