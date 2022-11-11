package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marcelofelixsalgado/financial-month-api/pkg/usecase/month/create"
	"net/http"
)

func CreateMonth(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var input create.InputCreateMonthDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		w.Write([]byte("Error converting request body to struct"))
		return
	}

	output := create.Execute(input)

	fmt.Println(output)
}
