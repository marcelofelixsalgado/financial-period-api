package balance

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/update"
)

type InputBalanceDto struct {
	periodId   string
	categoryId string
}

func ValidateCreateRequestBody(inputCreateBalanceDto create.InputCreateBalanceDto) *responses.ResponseMessage {
	inputBalanceDto := InputBalanceDto{
		periodId:   inputCreateBalanceDto.PeriodId,
		categoryId: inputCreateBalanceDto.CategoryId,
	}
	return validateRequestBody(inputBalanceDto)
}

func ValidateUpdateRequestBody(inputUpdateBalanceDto update.InputUpdateBalanceDto) *responses.ResponseMessage {

	if inputUpdateBalanceDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}

	inputBalanceDto := InputBalanceDto{
		periodId:   inputUpdateBalanceDto.PeriodId,
		categoryId: inputUpdateBalanceDto.CategoryId,
	}
	return validateRequestBody(inputBalanceDto)
}

func validateRequestBody(inputBalanceDto InputBalanceDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputBalanceDto.periodId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "period_id", "")
	}

	if inputBalanceDto.categoryId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "category_id", "")
	}

	return responseMessage
}
