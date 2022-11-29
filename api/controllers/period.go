package controllers

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/requests"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/database"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePeriod(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		message := responses.NewResponseMessage()
		message.AddMessageByIssue(responses.MalformedRequest, "body", "", "")
		jsonMessage, err := message.GetJsonMessage()
		if err != nil {
			responses.JSONErrorByCode(w, responses.InternalServerError)
			return
		}
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	var input create.InputCreatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		message := responses.NewResponseMessage()
		message.AddMessageByIssue(responses.MalformedRequest, "body", "", "")
		jsonMessage, err := message.GetJsonMessage()
		if err != nil {
			responses.JSONErrorByCode(w, responses.InternalServerError)
			return
		}
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	repository := repository.NewRepository(database.ConnectionPool)

	output, err := create.Execute(input, repository)
	if err != nil {
		log.Printf("Error trying to create the entity: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(outputJSON)
}

func ListPeriods(w http.ResponseWriter, r *http.Request) {
	var input list.InputListPeriodDto

	filterParameters, err := requests.SetupFilters(r)
	if err != nil {
		log.Printf("Error parsing the querystring parameters: %v", err)
		message := responses.NewResponseMessage()
		message.AddMessageByIssue(responses.MalformedRequest, "query_parameter", "", "")
		jsonMessage, err := message.GetJsonMessage()
		if err != nil {
			responses.JSONErrorByCode(w, responses.InternalServerError)
			return
		}
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	repository := repository.NewRepository(database.ConnectionPool)

	output, err := list.Execute(input, filterParameters, repository)
	if err != nil {
		log.Printf("Error listing the entity: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	outputJSON, err := json.Marshal(output.Periods)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
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

	repository := repository.NewRepository(database.ConnectionPool)

	output, err := find.Execute(input, repository)
	if err != nil {
		log.Printf("Error finding the entity: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		message := responses.NewResponseMessage()
		message.AddMessageByIssue(responses.MalformedRequest, "body", "", "")
		jsonMessage, err := message.GetJsonMessage()
		if err != nil {
			responses.JSONErrorByCode(w, responses.InternalServerError)
			return
		}
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}

	var input update.InputUpdatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		message := responses.NewResponseMessage()
		message.AddMessageByIssue(responses.MalformedRequest, "body", "", "")
		jsonMessage, err := message.GetJsonMessage()
		if err != nil {
			responses.JSONErrorByCode(w, responses.InternalServerError)
		}
		w.WriteHeader(message.GetMessage().HttpStatusCode)
		w.Write(jsonMessage)
		return
	}
	input.Id = Id

	repository := repository.NewRepository(database.ConnectionPool)

	output, err := update.Execute(input, repository)
	if err != nil {
		log.Printf("Error updating the entity: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func DeletePeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	var input = delete.InputDeletePeriodDto{
		Id: Id,
	}

	repository := repository.NewRepository(database.ConnectionPool)

	_, err := delete.Execute(input, repository)
	if err != nil {
		log.Printf("Error removing the entity: %v", err)
		responses.JSONErrorByCode(w, responses.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
