package login

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/marcelofelixsalgado/financial-period-api/commons/logger"
	loginUsecase "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/labstack/echo/v4"
)

type ILoginHandler interface {
	Login(ctx echo.Context) error
}

type LoginHandler struct {
	loginUseCase loginUsecase.ILoginUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "

func NewLoginHandler(loginUseCase loginUsecase.ILoginUseCase) ILoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
	}
}

func (userCredentialsHandler *LoginHandler) Login(ctx echo.Context) error {
	log := logger.GetLogger()

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input loginUsecase.InputUserLoginDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Validating input parameters
	if responseMessage := ValidateLoginRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := userCredentialsHandler.loginUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("Unable finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.Body, "email", input.Email)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.LoginFailed {
		log.Infof("Login failed: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.LoginFailed, responses.Body, "password", input.Password)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}
