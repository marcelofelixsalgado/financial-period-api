package group

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/update"
)

type InputGroupDto struct {
	code      string
	name      string
	groupType string
}

func ValidateCreateRequestBody(inputCreateGroupDto create.InputCreateGroupDto) *responses.ResponseMessage {
	inputGroupDto := InputGroupDto{
		code:      inputCreateGroupDto.Code,
		name:      inputCreateGroupDto.Name,
		groupType: inputCreateGroupDto.Type.Code,
	}
	return validateRequestBody(inputGroupDto)
}

func ValidateUpdateRequestBody(inputUpdateGroupDto update.InputUpdateGroupDto) *responses.ResponseMessage {
	if inputUpdateGroupDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputGroupDto := InputGroupDto{
		code:      inputUpdateGroupDto.Code,
		name:      inputUpdateGroupDto.Name,
		groupType: inputUpdateGroupDto.Type.Code,
	}
	return validateRequestBody(inputGroupDto)
}

func validateRequestBody(inputGroupDto InputGroupDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputGroupDto.code == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "code", "")
	}

	if inputGroupDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputGroupDto.groupType == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "group.type", "")
	}

	return responseMessage
}
