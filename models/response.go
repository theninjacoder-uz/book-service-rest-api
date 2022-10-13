package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	IsOk    bool        `json:"isOk"`
	Data    interface{} `json:"data"`
}

func setPayload(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func (res *Response) SuccessReponse(w http.ResponseWriter, statusCode int) {
	setPayload(w, statusCode, res)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {

	if err != nil {
		setPayload(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	setPayload(w, http.StatusBadRequest, nil)
}
