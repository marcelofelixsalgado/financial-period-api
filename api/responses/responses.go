package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"net/http"
)

type IResponseMessage interface {
	AddMessageByErrorCode(faults.ErrorCode) ResponseMessage
	AddMessageByIssue(faults.Issue, Location, string, string, ...string) ResponseMessage
	GetMessage() ResponseMessage
	Write(http.ResponseWriter)
}

type ResponseMessage struct {
	HttpStatusCode int                     `json:"-"`
	ErrorCode      string                  `json:"error_code"`
	Message        string                  `json:"message"`
	Details        []ResponseMessageDetail `json:"details,omitempty"`
}

type ResponseMessageDetail struct {
	Issue       string   `json:"issue"`
	Description string   `json:"description"`
	Location    Location `json:"location,omitempty"`
	Field       string   `json:"field,omitempty"`
	Value       string   `json:"value,omitempty"`
}

type Location string

const (
	Body           Location = "body"
	Header         Location = "header"
	QueryParameter Location = "query_parameter"
	PathParameter  Location = "path_parameter"
)

func NewResponseMessage() *ResponseMessage {
	return &ResponseMessage{}
}

func (responseMessage *ResponseMessage) GetMessage() ResponseMessage {
	return ResponseMessage{
		HttpStatusCode: responseMessage.HttpStatusCode,
		ErrorCode:      responseMessage.ErrorCode,
		Message:        responseMessage.Message,
		Details:        responseMessage.Details,
	}
}

// func (responseMessage *ResponseMessage) getJsonMessage() ([]byte, error) {
// 	message := responseMessage.getMessage()
// 	messageJSON, err := json.Marshal(message)
// 	if err != nil {
// 		return nil, fmt.Errorf("error converting struct to response body: %s", err)
// 	}
// 	return messageJSON, nil
// }

func (responseMessage *ResponseMessage) AddMessageByErrorCode(errorCode faults.ErrorCode) *ResponseMessage {
	referenceMessage, err := faults.FindByErrorCode(errorCode)
	if err != nil {
		log.Printf("Error trying to find the error by code: [%v]: - %v", errorCode, err)
		return NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
	}

	responseMessage.ErrorCode = string(referenceMessage.ErrorCode)
	responseMessage.Message = referenceMessage.Message
	responseMessage.HttpStatusCode = referenceMessage.HttpStatusCode

	return responseMessage
}

func (responseMessage *ResponseMessage) AddMessageByIssue(issue faults.Issue, location Location, field string, value string, descriptionArgs ...string) *ResponseMessage {

	referenceResponse, referenceResponseDetail, err := faults.FindByIssue(issue)
	if err != nil {
		log.Printf("Error trying to find the error by issue: [%v] - %v", issue, err)
		return NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
	}

	if referenceResponseDetail.LocationRequired && location == "" {
		log.Printf("Error trying to define a response message - location is required")
		return NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
	}
	if referenceResponseDetail.FieldRequired && field == "" {
		log.Printf("Error trying to define a response message - field is required")
		return NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
	}
	if referenceResponseDetail.ValueRequired && value == "" {
		log.Printf("Error trying to define a response message - value is required")
		return NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
	}
	if referenceResponseDetail.DescriptionArgs != len(descriptionArgs) {
		log.Printf("Error trying to define a response message - wrong number of argumentos passed. expected: [%d] - received: [%d]", referenceResponseDetail.DescriptionArgs, len(descriptionArgs))
		return NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
	}

	// If responseMessage doesn't exists yet
	if responseMessage.ErrorCode == "" {
		responseMessage.AddMessageByErrorCode(referenceResponse.ErrorCode)
	}

	messageDetail := buildMessageDetail(issue, referenceResponseDetail.Description, location, field, value, descriptionArgs)
	responseMessage.Details = append(responseMessage.Details, messageDetail)

	return responseMessage
}

func (responseMessage *ResponseMessage) AddMessageByInternalStatus(internalStatus status.InternalStatus, location Location, field string, value string) *ResponseMessage {

	switch internalStatus {
	case status.InternalServerError:
		responseMessage.AddMessageByErrorCode(faults.InternalServerError)
	case status.InvalidResourceId:
		responseMessage.AddMessageByIssue(faults.InvalidResourceId, location, field, value)
	case status.NoRecordsFound:
		responseMessage.AddMessageByIssue(faults.NoRecordsFound, location, field, value)
	}

	return responseMessage
}

func (responseMessage *ResponseMessage) Write(w http.ResponseWriter) {
	w.WriteHeader(responseMessage.GetMessage().HttpStatusCode)
	if err := json.NewEncoder(w).Encode(responseMessage.GetMessage()); err != nil {
		log.Printf("Error trying to encode response body message: %v", err)
	}
}

func buildMessageDetail(issue faults.Issue, description string, location Location, field string, value string, descriptionArgs []string) ResponseMessageDetail {

	switch len(descriptionArgs) {
	case 1:
		description = fmt.Sprintf(description, descriptionArgs[0])
	case 2:
		description = fmt.Sprintf(description, descriptionArgs[0], descriptionArgs[1])
	case 3:
		description = fmt.Sprintf(description, descriptionArgs[0], descriptionArgs[1], descriptionArgs[2])
	case 4:
		description = fmt.Sprintf(description, descriptionArgs[0], descriptionArgs[1], descriptionArgs[2], descriptionArgs[3])
	}

	return ResponseMessageDetail{
		Issue:       string(issue),
		Description: description,
		Location:    location,
		Field:       field,
		Value:       value,
	}
}
