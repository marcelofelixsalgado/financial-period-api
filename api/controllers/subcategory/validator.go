package subcategory

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/create"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/update"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
)

type InputSubCategoryDto struct {
	code     string
	name     string
	category Category
}

type Category struct {
	id string
}

func ValidateCreateRequestBody(inputCreateSubCategoryDto create.InputCreateSubCategoryDto) *responses.ResponseMessage {
	inputSubCategoryDto := InputSubCategoryDto{
		code: inputCreateSubCategoryDto.Code,
		name: inputCreateSubCategoryDto.Name,
		category: Category{
			id: inputCreateSubCategoryDto.Category.Id,
		},
	}
	return validateRequestBody(inputSubCategoryDto)
}

func ValidateUpdateRequestBody(inputUpdateSubCategoryDto update.InputUpdateSubCategoryDto) *responses.ResponseMessage {
	if inputUpdateSubCategoryDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputSubCategoryDto := InputSubCategoryDto{
		code: inputUpdateSubCategoryDto.Code,
		name: inputUpdateSubCategoryDto.Name,
		category: Category{
			id: inputUpdateSubCategoryDto.Category.Id,
		},
	}
	return validateRequestBody(inputSubCategoryDto)
}

func validateRequestBody(inputSubCategoryDto InputSubCategoryDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputSubCategoryDto.code == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "code", "")
	}

	if inputSubCategoryDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputSubCategoryDto.category.id == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "category.id", "")
	}

	return responseMessage
}
