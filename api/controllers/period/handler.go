package period

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/requests"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IPeriodHandler interface {
	CreatePeriod(ctx echo.Context) error
	ListPeriods(ctx echo.Context) error
	GetPeriodById(ctx echo.Context) error
	UpdatePeriod(ctx echo.Context) error
	DeletePeriod(ctx echo.Context) error
}

type PeriodHandler struct {
	createUseCase create.ICreateUseCase
	deleteUseCase delete.IDeleteUseCase
	findUseCase   find.IFindUseCase
	listUseCase   list.IListUseCase
	updateUseCase update.IUpdateUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "
const outputConversionErrorMessage = "Error trying to convert the output to response body: "
const unableFindEntityErrorMessage = "Unable to find the entity"

func NewPeriodHandler(createUseCase create.ICreateUseCase, deleteUseCase delete.IDeleteUseCase, findUseCase find.IFindUseCase, listUseCase list.IListUseCase, updateUseCase update.IUpdateUseCase) IPeriodHandler {
	return &PeriodHandler{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
		findUseCase:   findUseCase,
		listUseCase:   listUseCase,
		updateUseCase: updateUseCase,
	}
}

func (periodHandler *PeriodHandler) CreatePeriod(ctx echo.Context) error {

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Printf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Printf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input create.InputCreatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(ctx.Response().Writer)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input.TenantId = tenantId

	output, internalStatus, err := periodHandler.createUseCase.Execute(input)
	if internalStatus != status.Success {
		log.Printf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusCreated, output)
}

func (periodHandler *PeriodHandler) ListPeriods(ctx echo.Context) error {

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Printf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := list.InputListPeriodDto{
		TenantId: tenantId,
	}

	filterParameters, err := requests.SetupFilters(ctx.Request())
	if err != nil {
		log.Printf("Error parsing the querystring parameters: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := periodHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Printf("Error listing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.NoRecordsFound {
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output.Periods)
}

func (periodHandler *PeriodHandler) GetPeriodById(ctx echo.Context) error {
	id := ctx.Param("id")

	input := find.InputFindPeriodDto{
		Id: id,
	}

	output, internalStatus, err := periodHandler.findUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (periodHandler *PeriodHandler) UpdatePeriod(ctx echo.Context) error {
	id := ctx.Param("id")

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Printf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input update.InputUpdatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	input.Id = id

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(ctx.Response().Writer)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := periodHandler.updateUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error updating the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (periodHandler *PeriodHandler) DeletePeriod(ctx echo.Context) error {
	id := ctx.Param("id")

	var input = delete.InputDeletePeriodDto{
		Id: id,
	}

	_, internalStatus, err := periodHandler.deleteUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error removing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Printf("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.NoContent(http.StatusNoContent)
}
