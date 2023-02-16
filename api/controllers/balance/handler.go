package balance

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/update"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IBalanceHandler interface {
	CreateBalance(ctx echo.Context) error
	ListBalances(ctx echo.Context) error
	GetBalanceById(ctx echo.Context) error
	UpdateBalance(ctx echo.Context) error
	DeleteBalance(ctx echo.Context) error
}

type BalanceHandler struct {
	createUseCase create.ICreateUseCase
	listUseCase   list.IListUseCase
	findUseCase   find.IFindUseCase
	updateUseCase update.IUpdateUseCase
	deleteUseCase delete.IDeleteUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "
const outputConversionErrorMessage = "Error trying to convert the output to response body: "
const unableFindEntityErrorMessage = "Unable to find the entity"

func NewBalanceHandler(createUseCase create.ICreateUseCase,
	listUseCase list.IListUseCase,
	findUseCase find.IFindUseCase,
	updateUseCase update.IUpdateUseCase,
	deleteUseCase delete.IDeleteUseCase) IBalanceHandler {
	return &BalanceHandler{
		createUseCase: createUseCase,
		listUseCase:   listUseCase,
		findUseCase:   findUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (balanceHandler *BalanceHandler) CreateBalance(ctx echo.Context) error {
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

	var input create.InputCreateBalanceDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input.TenantId = tenantId

	output, internalStatus, err := balanceHandler.createUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Printf("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "period_id", input.PeriodId).AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "category_id", input.CategoryId)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus != status.Success {
		log.Printf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusCreated, output)
}

func (balanceHandler *BalanceHandler) ListBalances(ctx echo.Context) error {
	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Printf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := list.InputListBalanceDto{
		TenantId: tenantId,
	}

	output, internalStatus, err := balanceHandler.listUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Printf("Error listing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.NoRecordsFound {
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output.Balances)
}

func (balanceHandler *BalanceHandler) GetBalanceById(ctx echo.Context) error {
	id := ctx.Param("id")

	input := find.InputFindBalanceDto{
		Id: id,
	}

	output, internalStatus, err := balanceHandler.findUseCase.Execute(input)
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

func (balanceHandler *BalanceHandler) UpdateBalance(ctx echo.Context) error {
	id := ctx.Param("id")

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

	var input update.InputUpdateBalanceDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Printf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	input.Id = id
	input.TenantId = tenantId

	output, internalStatus, err := balanceHandler.updateUseCase.Execute(input)
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

func (balanceHandler *BalanceHandler) DeleteBalance(ctx echo.Context) error {
	id := ctx.Param("id")

	var input = delete.InputDeleteBalanceDto{
		Id: id,
	}

	_, internalStatus, err := balanceHandler.deleteUseCase.Execute(input)
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
