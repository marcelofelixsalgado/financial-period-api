package user

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"

	userCredentialsCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	userCredentialsUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
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

func ValidateUserCredentialsCreateRequestBody(inputCreateUserCredentialsDto userCredentialsCreate.InputCreateUserCredentialsDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()

	if inputCreateUserCredentialsDto.UserId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputCreateUserCredentialsDto.Password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}

func ValidateUserCredentialsUpdateRequestBody(inputUpdateUserCredentialsDto userCredentialsUpdate.InputUpdateUserCredentialsDto) *responses.ResponseMessage {
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
