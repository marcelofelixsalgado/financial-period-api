package credentials

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	loginUsecase "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"
)

type InputUserCredentialsDto struct {
	userId   string
	password string
}

func ValidateCreateRequestBody(inputCreateUserCredentialsDto create.InputCreateUserCredentialsDto) *responses.ResponseMessage {
	inputUserCredentialsDto := InputUserCredentialsDto{
		userId:   inputCreateUserCredentialsDto.UserId,
		password: inputCreateUserCredentialsDto.Password,
	}
	return validateRequestBody(inputUserCredentialsDto)
}

func ValidateUpdateRequestBody(inputUpdateUserCredentialsDto update.InputUpdateUserCredentialsDto) *responses.ResponseMessage {
	if inputUpdateUserCredentialsDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputUserCredentialsDto := InputUserCredentialsDto{
		userId:   inputUpdateUserCredentialsDto.UserId,
		password: inputUpdateUserCredentialsDto.Password,
	}
	return validateRequestBody(inputUserCredentialsDto)
}

func validateRequestBody(inputUserCredentialsDto InputUserCredentialsDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputUserCredentialsDto.userId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputUserCredentialsDto.password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}

func ValidateLoginRequestBody(inputUserLoginDto loginUsecase.InputUserLoginDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()

	if inputUserLoginDto.Email == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "email", "")
	}

	if inputUserLoginDto.Password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}
