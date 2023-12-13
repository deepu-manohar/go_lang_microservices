package main

import (
	"encoding/json"
	"net/http"
)

const maxBytes = 1048576

func readJson(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func writeJson(w http.ResponseWriter, status int,
	data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, error := w.Write(out)
	return error
}

func errorJson(w http.ResponseWriter, err error, status int) error {
	defaultStatus := http.StatusBadRequest
	if status != 0 {
		defaultStatus = status
	}
	responseData := SendMailResponse{Error: true, Message: err.Error()}
	return writeJson(w, defaultStatus, responseData)
}
