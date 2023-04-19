package period

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/marcelofelixsalgado/financial-period-api/api/requests"
	"github.com/marcelofelixsalgado/financial-period-api/commons/logger"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/findbyid"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

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
	createUseCase   create.ICreateUseCase
	deleteUseCase   delete.IDeleteUseCase
	findByIdUseCase findbyid.IFindByIdUseCase
	listUseCase     list.IListUseCase
	updateUseCase   update.IUpdateUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "
const outputConversionErrorMessage = "Error trying to convert the output to response body: "
const unableFindEntityErrorMessage = "Unable to find the entity"

func NewPeriodHandler(createUseCase create.ICreateUseCase, deleteUseCase delete.IDeleteUseCase, findByIdUseCase findbyid.IFindByIdUseCase, listUseCase list.IListUseCase, updateUseCase update.IUpdateUseCase) IPeriodHandler {
	return &PeriodHandler{
		createUseCase:   createUseCase,
		deleteUseCase:   deleteUseCase,
		findByIdUseCase: findByIdUseCase,
		listUseCase:     listUseCase,
		updateUseCase:   updateUseCase,
	}
}

func (periodHandler *PeriodHandler) CreatePeriod(ctx echo.Context) error {
	log := logger.GetLogger()

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Errorf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input create.InputCreatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input.TenantId = tenantId

	output, internalStatus, err := periodHandler.createUseCase.Execute(input)
	if internalStatus == status.OverlappingPeriodDates {
		log.Warnf("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.OverlappingPeriodDates, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Warnf("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "code", input.Code)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus != status.Success {
		log.Errorf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusCreated, output)
}

func (periodHandler *PeriodHandler) ListPeriods(ctx echo.Context) error {
	log := logger.GetLogger()

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Errorf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := list.InputListPeriodDto{
		TenantId: tenantId,
	}

	filterParameters, err := requests.SetupFilters(ctx.Request())
	if err != nil {
		log.Warnf("Error parsing the querystring parameters: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := periodHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error listing the entity: %v", err)
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
	log := logger.GetLogger()

	id := ctx.Param("id")

	input := findbyid.InputFindByIdPeriodDto{
		Id: id,
	}

	output, internalStatus, err := periodHandler.findByIdUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (periodHandler *PeriodHandler) UpdatePeriod(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input update.InputUpdatePeriodDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	input.Id = id

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := periodHandler.updateUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Warnf("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "code", input.Code)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InternalServerError {
		log.Errorf("Error updating the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (periodHandler *PeriodHandler) DeletePeriod(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	var input = delete.InputDeletePeriodDto{
		Id: id,
	}

	_, internalStatus, err := periodHandler.deleteUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error removing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.NoContent(http.StatusNoContent)
}
