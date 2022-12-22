package credentials

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	loginUsecase "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"
)

func ValidateCreateRequestBody(inputCreateUserCredentialsDto create.InputCreateUserCredentialsDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()

	if inputCreateUserCredentialsDto.UserId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputCreateUserCredentialsDto.Password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}

func ValidateUpdateRequestBody(inputUpdateUserCredentialsDto update.InputUpdateUserCredentialsDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()
	if inputUpdateUserCredentialsDto.UserId == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}

	if inputUpdateUserCredentialsDto.UserId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputUpdateUserCredentialsDto.NewPassword == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "new_password", "")
	}

	if inputUpdateUserCredentialsDto.CurrentPassword == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "current_password", "")
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
