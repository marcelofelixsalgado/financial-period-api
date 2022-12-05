package login

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/login"
	"net/http"
)

type ILoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginHandler struct {
	loginUseCase login.ILoginUseCase
}

func NewLoginHandler(loginUseCase login.ILoginUseCase) ILoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
	}
}

func (loginHandler *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input login.InputUserLoginDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	// Validating input parameters
	if responseMessage := validateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, internalStatus, err := loginHandler.loginUseCase.Execute(input)
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
