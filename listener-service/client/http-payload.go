package client

type LogResponse struct {
	Recorded bool   `json:"recorded"`
	Status   int    `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
}

type LogRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type EndPoint struct {
	method string
	url    string
}

func (endPoint *EndPoint) getCompleteURL(baseUrl string) string {
	return baseUrl + endPoint.url
}
