package subcategory

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/marcelofelixsalgado/financial-period-api/api/requests"
	"github.com/marcelofelixsalgado/financial-period-api/commons/logger"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/create"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/delete"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/find"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/list"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/update"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/labstack/echo/v4"
)

type ISubCategoryHandler interface {
	CreateSubCategory(ctx echo.Context) error
	ListCategories(ctx echo.Context) error
	GetSubCategoryById(ctx echo.Context) error
	UpdateSubCategory(ctx echo.Context) error
	DeleteSubCategory(ctx echo.Context) error
}

type SubCategoryHandler struct {
	createUseCase create.ICreateUseCase
	deleteUseCase delete.IDeleteUseCase
	findUseCase   find.IFindUseCase
	listUseCase   list.IListUseCase
	updateUseCase update.IUpdateUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "
const unableFindEntityErrorMessage = "Unable to find the entity"

func NewSubCategoryHandler(createUseCase create.ICreateUseCase, deleteUseCase delete.IDeleteUseCase, findUseCase find.IFindUseCase, listUseCase list.IListUseCase, updateUseCase update.IUpdateUseCase) ISubCategoryHandler {
	return &SubCategoryHandler{
		createUseCase: createUseCase,
		updateUseCase: updateUseCase,
		findUseCase:   findUseCase,
		listUseCase:   listUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (categoryHandler SubCategoryHandler) CreateSubCategory(ctx echo.Context) error {
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

	var input create.InputCreateSubCategoryDto

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

	output, internalStatus, err := categoryHandler.createUseCase.Execute(input)
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

func (categoryHandler SubCategoryHandler) ListCategories(ctx echo.Context) error {
	log := logger.GetLogger()

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Errorf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := list.InputListSubCategoryDto{
		TenantId: tenantId,
	}

	filterParameters, err := requests.SetupFilters(ctx.Request())
	if err != nil {
		log.Warnf("Error parsing the querystring parameters: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := categoryHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error listing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.NoRecordsFound {
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output.SubCategories)
}

func (categoryHandler SubCategoryHandler) GetSubCategoryById(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	input := find.InputFindSubCategoryDto{
		Id: id,
	}

	output, internalStatus, err := categoryHandler.findUseCase.Execute(input)
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

func (categoryHandler SubCategoryHandler) UpdateSubCategory(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input update.InputUpdateSubCategoryDto

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

	output, internalStatus, err := categoryHandler.updateUseCase.Execute(input)
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

func (categoryHandler SubCategoryHandler) DeleteSubCategory(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	var input = delete.InputDeleteSubCategoryDto{
		Id: id,
	}

	_, internalStatus, err := categoryHandler.deleteUseCase.Execute(input)
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
