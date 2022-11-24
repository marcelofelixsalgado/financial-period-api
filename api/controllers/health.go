package controllers

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {

	successMessage := message{
		Status: "Ok",
	}

	messageJSON, _ := json.Marshal(successMessage)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(messageJSON))
}
