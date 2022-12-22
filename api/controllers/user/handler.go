package user

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/requests"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type IUserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	ListUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	createUseCase create.ICreateUseCase
	deleteUseCase delete.IDeleteUseCase
	findUseCase   find.IFindUseCase
	listUseCase   list.IListUseCase
	updateUseCase update.IUpdateUseCase
}

func NewUserHandler(createUseCase create.ICreateUseCase, deleteUseCase delete.IDeleteUseCase, findUseCase find.IFindUseCase, listUseCase list.IListUseCase, updateUseCase update.IUpdateUseCase) IUserHandler {
	return &UserHandler{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
		findUseCase:   findUseCase,
		listUseCase:   listUseCase,
		updateUseCase: updateUseCase,
	}
}

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input create.InputCreateUserDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, internalStatus, err := userHandler.createUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Printf("Error trying to create the entity - duplicate key: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "phone", input.Phone).AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "email", input.Email).Write(w)
		return
	}
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

func (userHandler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	var input list.InputListUserDto

	filterParameters, err := requests.SetupFilters(r)
	if err != nil {
		log.Printf("Error parsing the querystring parameters: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "").Write(w)
		return
	}

	output, internalStatus, err := userHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Printf("Error listing the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.NoRecordsFound {
		responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "").Write(w)
		return
	}

	outputJSON, err := json.Marshal(output.Users)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func (userHandler *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	sameUser, err := checkSameUser(id, r)
	if err != nil {
		log.Printf("Error extracting the user id from token: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if !sameUser {
		log.Printf("A user cannot update another user info")
		responses.NewResponseMessage().AddMessageByIssue(faults.PermissionDenied, "", "", "").Write(w)
		return
	}

	input := find.InputFindUserDto{
		Id: id,
	}

	output, internalStatus, err := userHandler.findUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity")
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id).Write(w)
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

func (userHandler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	sameUser, err := checkSameUser(id, r)
	if err != nil {
		log.Printf("Error extracting the user id from token: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if !sameUser {
		log.Printf("A user cannot update another user info")
		responses.NewResponseMessage().AddMessageByIssue(faults.PermissionDenied, "", "", "").Write(w)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input update.InputUpdateUserDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}
	input.Id = id

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, internalStatus, err := userHandler.updateUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error updating the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id).Write(w)
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

func (userHandler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	sameUser, err := checkSameUser(id, r)
	if err != nil {
		log.Printf("Error extracting the user id from token: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if !sameUser {
		log.Printf("A user cannot update another user info")
		responses.NewResponseMessage().AddMessageByIssue(faults.PermissionDenied, "", "", "").Write(w)
		return
	}

	var input = delete.InputDeleteUserDto{
		Id: id,
	}

	_, internalStatus, err := userHandler.deleteUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error removing the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id).Write(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func checkSameUser(requestUserId string, r *http.Request) (bool, error) {
	tokenUserId, err := auth.ExtractUserId(r)
	if err != nil {
		return false, err
	}
	if !strings.EqualFold(requestUserId, tokenUserId) {
		return false, nil
	}
	return true, nil
}
