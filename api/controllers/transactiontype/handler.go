package transactiontype

import (
	"marcelofelixsalgado/financial-period-api/api/requests"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/commons/logger"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/list"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITransactionTypeHandler interface {
	GetTransactionTypeByCode(ctx echo.Context) error
	ListTransactionTypes(ctx echo.Context) error
}

type TransactionTypeHandler struct {
	findUseCase find.IFindUseCase
	listUseCase list.IListUseCase
}

const unableFindEntityErrorMessage = "Unable to find the entity"

func NewTransactionTypeHandler(findUseCase find.IFindUseCase, listUseCase list.IListUseCase) ITransactionTypeHandler {
	return &TransactionTypeHandler{
		findUseCase: findUseCase,
		listUseCase: listUseCase,
	}
}

func (transactionTypeHandler TransactionTypeHandler) GetTransactionTypeByCode(ctx echo.Context) error {
	log := logger.GetLogger()

	code := ctx.Param("code")

	input := find.InputFindTransactionTypeDto{
		Code: code,
	}

	output, internalStatus, err := transactionTypeHandler.findUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "code", code)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (transactionTypeHandler TransactionTypeHandler) ListTransactionTypes(ctx echo.Context) error {
	log := logger.GetLogger()

	input := list.InputListTransactionTypeDto{}

	filterParameters, err := requests.SetupFilters(ctx.Request())
	if err != nil {
		log.Warnf("Error parsing the querystring parameters: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := transactionTypeHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error listing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.NoRecordsFound {
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output.TransactionTypes)
}
