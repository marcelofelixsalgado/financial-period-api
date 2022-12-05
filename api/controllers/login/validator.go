package login

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/login"
)

type InputLoginDto struct {
	email    string
	password string
}

func validateRequestBody(inputUserLoginDto login.InputUserLoginDto) *responses.ResponseMessage {
	inputLoginDto := InputLoginDto{
		email:    inputUserLoginDto.Email,
		password: inputUserLoginDto.Password,
	}

	responseMessage := responses.NewResponseMessage()

	if inputLoginDto.email == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "email", "")
	}

	if inputLoginDto.password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}
