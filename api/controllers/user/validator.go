package user

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"
)

type InputUserDto struct {
	name  string
	phone string
	email string
}

func ValidateCreateRequestBody(inputCreateUserDto create.InputCreateUserDto) *responses.ResponseMessage {
	inputUserDto := InputUserDto{
		name:  inputCreateUserDto.Name,
		phone: inputCreateUserDto.Phone,
		email: inputCreateUserDto.Email,
	}
	return validateRequestBody(inputUserDto)
}

func ValidateUpdateRequestBody(inputUpdateUserDto update.InputUpdateUserDto) *responses.ResponseMessage {
	if inputUpdateUserDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputUserDto := InputUserDto{
		name:  inputUpdateUserDto.Name,
		phone: inputUpdateUserDto.Phone,
		email: inputUpdateUserDto.Email,
	}
	return validateRequestBody(inputUserDto)
}

func validateRequestBody(inputUserDto InputUserDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputUserDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputUserDto.phone == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "phone", "")
	}

	if inputUserDto.email == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "email", "")
	}
	return responseMessage
}
