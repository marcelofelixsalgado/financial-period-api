package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type message struct {
	Status string `json: "status"`
}

func Health(w http.ResponseWriter, r *http.Request) {

	successMessage := message{
		Status: "Ok",
	}

	messageJSON, err := json.Marshal(successMessage)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(messageJSON))
}
