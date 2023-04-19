package category

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/create"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/update"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
)

type InputCategoryDto struct {
	code            string
	name            string
	transactiontype TransactionType
}

type TransactionType struct {
	code string
}

func ValidateCreateRequestBody(inputCreateCategoryDto create.InputCreateCategoryDto) *responses.ResponseMessage {
	inputCategoryDto := InputCategoryDto{
		code: inputCreateCategoryDto.Code,
		name: inputCreateCategoryDto.Name,
		transactiontype: TransactionType{
			code: inputCreateCategoryDto.TransactionType.Code,
		},
	}
	return validateRequestBody(inputCategoryDto)
}

func ValidateUpdateRequestBody(inputUpdateCategoryDto update.InputUpdateCategoryDto) *responses.ResponseMessage {
	if inputUpdateCategoryDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputCategoryDto := InputCategoryDto{
		code: inputUpdateCategoryDto.Code,
		name: inputUpdateCategoryDto.Name,
		transactiontype: TransactionType{
			code: inputUpdateCategoryDto.TransactionType.Code,
		},
	}
	return validateRequestBody(inputCategoryDto)
}

func validateRequestBody(inputCategoryDto InputCategoryDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputCategoryDto.code == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "code", "")
	}

	if inputCategoryDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputCategoryDto.transactiontype.code == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "transaction_type.code", "")
	}

	return responseMessage
}
