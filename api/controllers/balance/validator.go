package balance

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/create"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
)

type InputBalanceDto struct {
	periodId    string
	categoryId  string
	limitAmount float32
}

func ValidateCreateRequestBody(inputCreateBalanceDto create.InputCreateBalanceDto) *responses.ResponseMessage {
	inputBalanceDto := InputBalanceDto{
		periodId:   inputCreateBalanceDto.PeriodId,
		categoryId: inputCreateBalanceDto.CategoryId,
	}

	responseMessage := responses.NewResponseMessage()

	if inputBalanceDto.periodId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "period_id", "")
	}

	if inputBalanceDto.categoryId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "category_id", "")
	}

	return responseMessage
}
