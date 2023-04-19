package user

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/marcelofelixsalgado/financial-period-api/api/requests"
	"github.com/marcelofelixsalgado/financial-period-api/commons/logger"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/delete"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/find"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/list"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"

	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	userCredentialsCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	userCredentialsUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	CreateUser(ctx echo.Context) error
	ListUsers(ctx echo.Context) error
	GetUserById(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	CreateUserCredentials(ctx echo.Context) error
	UpdateUserCredentials(ctx echo.Context) error
}

type UserHandler struct {
	userCreateUseCase            create.ICreateUseCase
	userDeleteUseCase            delete.IDeleteUseCase
	userFindUseCase              find.IFindUseCase
	userListUseCase              list.IListUseCase
	userUpdateUseCase            update.IUpdateUseCase
	userCredentialsCreateUseCase userCredentialsCreate.ICreateUseCase
	userCredentialsUpdateUseCase userCredentialsUpdate.IUpdateUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "
const userAccessAnotherUserErrorMessage = "The user is not allowed to access another user info"
const extractingUserIdErrorMessage = "Error extracting the user id from token: "

func NewUserHandler(userCreateUseCase create.ICreateUseCase,
	userDeleteUseCase delete.IDeleteUseCase,
	userFindUseCase find.IFindUseCase,
	userListUseCase list.IListUseCase,
	userUpdateUseCase update.IUpdateUseCase,
	userCredentialsCreateUseCase userCredentialsCreate.ICreateUseCase,
	userCredentialsUpdateUseCase userCredentialsUpdate.IUpdateUseCase) IUserHandler {
	return &UserHandler{
		userCreateUseCase:            userCreateUseCase,
		userDeleteUseCase:            userDeleteUseCase,
		userFindUseCase:              userFindUseCase,
		userListUseCase:              userListUseCase,
		userUpdateUseCase:            userUpdateUseCase,
		userCredentialsCreateUseCase: userCredentialsCreateUseCase,
		userCredentialsUpdateUseCase: userCredentialsUpdateUseCase,
	}
}

func (userHandler *UserHandler) CreateUser(ctx echo.Context) error {
	log := logger.GetLogger()

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input create.InputCreateUserDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := userHandler.userCreateUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Infof("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "phone", input.Phone).AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "email", input.Email)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus != status.Success {
		log.Errorf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusCreated, output)
}

func (userHandler *UserHandler) ListUsers(ctx echo.Context) error {
	log := logger.GetLogger()

	var input list.InputListUserDto

	filterParameters, err := requests.SetupFilters(ctx.Request())
	if err != nil {
		log.Warnf("Error parsing the querystring parameters: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := userHandler.userListUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error listing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.NoRecordsFound {
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output.Users)
}

func (userHandler *UserHandler) GetUserById(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	sameUser, err := checkSameUser(id, ctx.Request())
	if err != nil {
		log.Errorf("%s%v", extractingUserIdErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if !sameUser {
		log.Infof("%s", userAccessAnotherUserErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.PermissionDenied, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := find.InputFindUserDto{
		Id: id,
	}

	output, internalStatus, err := userHandler.userFindUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("Unable finding the entity")
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (userHandler *UserHandler) UpdateUser(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	sameUser, err := checkSameUser(id, ctx.Request())
	if err != nil {
		log.Errorf("%s%v", extractingUserIdErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if !sameUser {
		log.Infof("%s", userAccessAnotherUserErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.PermissionDenied, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input update.InputUpdateUserDto

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

	output, internalStatus, err := userHandler.userUpdateUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error updating the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("Unable finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (userHandler *UserHandler) DeleteUser(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	sameUser, err := checkSameUser(id, ctx.Request())
	if err != nil {
		log.Errorf("%s%v", extractingUserIdErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if !sameUser {
		log.Infof("%s", userAccessAnotherUserErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.PermissionDenied, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input = delete.InputDeleteUserDto{
		Id: id,
	}

	_, internalStatus, err := userHandler.userDeleteUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error removing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("Unable finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func checkSameUser(requestUserId string, r *http.Request) (bool, error) {
	tokenUserId, err := auth.ExtractUserId(r)
	if err != nil {
		return false, err
	}
	if !strings.EqualFold(requestUserId, tokenUserId) {
		return false, nil
	}
	return true, nil
}

func (userHandler *UserHandler) CreateUserCredentials(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input userCredentialsCreate.InputCreateUserCredentialsDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	input.UserId = id

	// Validating input parameters
	if responseMessage := ValidateUserCredentialsCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := userHandler.userCredentialsCreateUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Infof("Error trying to create the entity - The user already has a password")
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus != status.Success {
		log.Errorf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusCreated, output)
}

func (userHandler *UserHandler) UpdateUserCredentials(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input userCredentialsUpdate.InputUpdateUserCredentialsDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	input.UserId = id

	// Validating input parameters
	if responseMessage := ValidateUserCredentialsUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := userHandler.userCredentialsUpdateUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error updating the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("Error updating the entity - Unable finding the entity")
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.PasswordsDontMatch {
		log.Infof("Error updating the entity - passwords don't match")
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}
