package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"marcelofelixsalgado/financial-month-api/pkg/infrastructure/month/repository/mysql"
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
		log.Printf("Error converting request body to struct: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository := mysql.NewRepository()

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
