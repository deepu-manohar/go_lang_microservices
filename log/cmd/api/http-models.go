package main

type LogRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type LogResponse struct {
	Recorded bool   `json:"recorded"`
	Message  string `json:"message,omitempty"`
}
