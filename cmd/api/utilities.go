package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	js, err := json.MarshalIndent(wrapper, "", "\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, status int, err error) error {
	type jsonError struct {
		Message string `json:"message"`
	}
	return app.writeJSON(w, status, jsonError{err.Error()}, "error")
}
