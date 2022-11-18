package routes

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-month-api/api/responses"
	"marcelofelixsalgado/financial-month-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-month-api/pkg/usecase/month/create"
	"net/http"
)

func CreateMonth(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input create.InputCreateMonthDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		message := responses.NewResponseMessage()
		message.AddMessageByIssue(responses.MalformedRequest, "body", "", "")
		jsonMessage, err := message.GetJsonMessage()
		if err != nil {
			message = responses.NewResponseMessage()
			message.AddMessageByErrorCode(responses.InternalServerError)
			jsonMessage, _ = message.GetJsonMessage()
		}
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	repository := repository.NewRepository()

	output, err := create.Execute(input, repository)
	if err != nil {
		log.Printf("Error creating the entity: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.Write(jsonMessage)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error converting struct to response body: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.Write(jsonMessage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(outputJSON)
}
