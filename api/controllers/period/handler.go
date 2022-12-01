package period

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/requests"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"net/http"

	"github.com/gorilla/mux"
)

type IPeriodHandler interface {
	CreatePeriod(w http.ResponseWriter, r *http.Request)
	ListPeriods(w http.ResponseWriter, r *http.Request)
	GetPeriodById(w http.ResponseWriter, r *http.Request)
	UpdatePeriod(w http.ResponseWriter, r *http.Request)
	DeletePeriod(w http.ResponseWriter, r *http.Request)
}

type PeriodHandler struct {
	createUseCase create.ICreateUseCase
	deleteUseCase delete.IDeleteUseCase
	findUseCase   find.IFindUseCase
	listUseCase   list.IListUseCase
	updateUseCase update.IUpdateUseCase
}

func NewPeriodHandler(createUseCase create.ICreateUseCase, deleteUseCase delete.IDeleteUseCase, findUseCase find.IFindUseCase, listUseCase list.IListUseCase, updateUseCase update.IUpdateUseCase) IPeriodHandler {
	return &PeriodHandler{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
		findUseCase:   findUseCase,
		listUseCase:   listUseCase,
		updateUseCase: updateUseCase,
	}
}

func (periodHandler *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input create.InputCreatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	output, internalStatus, err := periodHandler.createUseCase.Execute(input)
	if internalStatus != status.Success {
		log.Printf("Error trying to create the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(outputJSON)
}

func (periodHandler *PeriodHandler) ListPeriods(w http.ResponseWriter, r *http.Request) {
	var input list.InputListPeriodDto

	filterParameters, err := requests.SetupFilters(r)
	if err != nil {
		log.Printf("Error parsing the querystring parameters: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "").Write(w)
		return
	}

	output, internalStatus, err := periodHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Printf("Error listing the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.NoRecordsFound {
		responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "").Write(w)
		return
	}

	outputJSON, err := json.Marshal(output.Periods)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func (periodHandler *PeriodHandler) GetPeriodById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	input := find.InputFindPeriodDto{
		Id: Id,
	}

	output, internalStatus, err := periodHandler.findUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", Id).Write(w)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func (periodHandler *PeriodHandler) UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input update.InputUpdatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}
	input.Id = Id

	output, internalStatus, err := periodHandler.updateUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error updating the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", Id).Write(w)
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func (periodHandler *PeriodHandler) DeletePeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	var input = delete.InputDeletePeriodDto{
		Id: Id,
	}

	_, internalStatus, err := periodHandler.deleteUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error removing the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", Id).Write(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}