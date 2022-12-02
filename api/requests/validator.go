package requests

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"
	"time"
)

type InputPeriodDto struct {
	code      string
	name      string
	year      int
	startDate string
	endDate   string
}

func ValidateCreateRequestBody(inputCreatePeriodDto create.InputCreatePeriodDto) *responses.ResponseMessage {
	inputPeriodDto := InputPeriodDto{
		code:      inputCreatePeriodDto.Code,
		name:      inputCreatePeriodDto.Name,
		year:      inputCreatePeriodDto.Year,
		startDate: inputCreatePeriodDto.StartDate,
		endDate:   inputCreatePeriodDto.EndDate,
	}
	return validateRequestBody(inputPeriodDto)
}
func ValidateUpdateRequestBody(inputUpdatePeriodDto update.InputUpdatePeriodDto) *responses.ResponseMessage {
	if inputUpdatePeriodDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputPeriodDto := InputPeriodDto{
		code:      inputUpdatePeriodDto.Code,
		name:      inputUpdatePeriodDto.Name,
		year:      inputUpdatePeriodDto.Year,
		startDate: inputUpdatePeriodDto.StartDate,
		endDate:   inputUpdatePeriodDto.EndDate,
	}
	return validateRequestBody(inputPeriodDto)
}

func validateRequestBody(inputPeriodDto InputPeriodDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputPeriodDto.code == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "code", "")
	}

	if inputPeriodDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputPeriodDto.year == 0 {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "year", "")
	}

	if inputPeriodDto.startDate == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "start_date", "")
	}

	if inputPeriodDto.endDate == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "end_date", "")
	}

	validDates := true
	startDate, err := time.Parse(time.RFC3339, inputPeriodDto.startDate)
	if err != nil {
		validDates = false
		responseMessage.AddMessageByIssue(faults.InvalidDateTimeValue, responses.Body, "start_date", inputPeriodDto.startDate)
	}

	endDate, err := time.Parse(time.RFC3339, inputPeriodDto.endDate)
	if err != nil {
		validDates = false
		responseMessage.AddMessageByIssue(faults.InvalidDateTimeValue, responses.Body, "end_date", inputPeriodDto.endDate)
	}

	if validDates && (startDate.Equal(endDate) || startDate.After(endDate)) {
		responseMessage.AddMessageByIssue(faults.ConditionalLowerThan, responses.Body, "start_date", inputPeriodDto.startDate, "start_date", "end_date")
	}

	return responseMessage
}
