package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJson from http body
// Go 1.18.x you can any as type of data if go version lower than 1.18 you need to use interface{} as data type in case of you need to accept any type of data
func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	// 1mb = 1048576
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decodeData := json.NewDecoder(r.Body)
	err := decodeData.Decode(data)
	if err != nil {
		return err
	}

	err = decodeData.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have a single of JSON value")
	}
	return nil
}

// writeJson
func (app *Config) writeJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	jsonData, err := json.Marshal(data)
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
	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// errorJson
func (app *Config) errorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJson(w, statusCode, payload)
}
