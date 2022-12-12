package credentials

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	createUseCase "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	loginUsecase "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"
	updateUseCase "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"net/http"

	"github.com/gorilla/mux"
)

type IUserCredentialsHandler interface {
	CreateUserCredentials(w http.ResponseWriter, r *http.Request)
	UpdateUserCredentials(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserCredentialsHandler struct {
	createUseCase createUseCase.ICreateUseCase
	updateUseCase updateUseCase.IUpdateUseCase
	loginUseCase  loginUsecase.ILoginUseCase
}

func NewUserCredentialsHandler(createUseCase createUseCase.ICreateUseCase, updateUseCase updateUseCase.IUpdateUseCase, loginUseCase loginUsecase.ILoginUseCase) IUserCredentialsHandler {
	return &UserCredentialsHandler{
		createUseCase: createUseCase,
		updateUseCase: updateUseCase,
		loginUseCase:  loginUseCase,
	}
}

func (userCredentialsHandler *UserCredentialsHandler) CreateUserCredentials(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input create.InputCreateUserCredentialsDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}
	input.UserId = Id

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, internalStatus, err := userCredentialsHandler.createUseCase.Execute(input)
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

func (userCredentialsHandler *UserCredentialsHandler) UpdateUserCredentials(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input update.InputUpdateUserCredentialsDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}
	input.Id = Id

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, internalStatus, err := userCredentialsHandler.updateUseCase.Execute(input)
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

func (userCredentialsHandler *UserCredentialsHandler) Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input loginUsecase.InputUserLoginDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	// Validating input parameters
	if responseMessage := ValidateLoginRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, internalStatus, err := userCredentialsHandler.loginUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("Unable finding the entity: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.Body, "email", input.Email).Write(w)
		return
	}
	if internalStatus == status.LoginFailed {
		log.Printf("Login failed: %v", err)
		responses.NewResponseMessage().AddMessageByInternalStatus(status.LoginFailed, responses.Body, "password", input.Password).Write(w)
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
