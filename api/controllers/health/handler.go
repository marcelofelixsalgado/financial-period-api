package health

import (
	"encoding/json"
	"net/http"
)

type IHealthHandler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type message struct {
	Status string `json:"status"`
}

type HealthHandler struct {
}

func NewHealthHandler() IHealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {

	successMessage := message{
		Status: "Ok",
	}

	messageJSON, _ := json.Marshal(successMessage)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(messageJSON))
}
