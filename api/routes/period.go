package routes

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePeriod(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
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

	var input create.InputCreatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error converting input data: %s", err)
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
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error converting struct to response body: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(outputJSON)
}

func ListPeriods(w http.ResponseWriter, r *http.Request) {
	var input list.InputListPeriodDto

	repository := repository.NewRepository()

	output, err := list.Execute(input, repository)
	if err != nil {
		log.Printf("Error listing the entity: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	outputJSON, err := json.Marshal(output.Periods)
	if err != nil {
		log.Printf("Error converting struct to response body: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func GetPeriodById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	input := find.InputFindPeriodDto{
		Id: Id,
	}

	repository := repository.NewRepository()

	output, err := find.Execute(input, repository)
	if err != nil {
		log.Printf("Error finding the entity: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error converting struct to response body: %s", err)
		message := responses.NewResponseMessage()
		message.AddMessageByErrorCode(responses.InternalServerError)
		jsonMessage, _ := message.GetJsonMessage()
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func UpdatePeriod(w http.ResponseWriter, r *http.Request) {
}

func DeletePeriod(w http.ResponseWriter, r *http.Request) {
}
