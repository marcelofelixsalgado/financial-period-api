package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"marcelofelixsalgado/financial-month-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-month-api/pkg/usecase/month/create"
	"net/http"
)

func CreateMonth(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input create.InputCreateMonthDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// if err := messages.AddMessageByIssue(messages.MalformedRequest, "body", "", ""); err != nil {
		// 	messages.AddMessageByErrorCode(messages.InternalServerError)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// }
		// outputJSON, err := json.Marshal(messages.GetMessages())
		// if err != nil {
		// 	messages.AddMessageByErrorCode(messages.InternalServerError)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// }
		// w.Write(outputJSON)
		// return
	}

	repository := repository.NewRepository()

	output, err := create.Execute(input, repository)
	if err != nil {
		log.Printf("Error creating the entity: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error converting struct to response body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(outputJSON)
}
