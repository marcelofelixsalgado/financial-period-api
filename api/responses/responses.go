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

// Return a JSON response to request
// func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		log.Fatal(err)
// 	}
// }

// // Rerturn an error in JSON format
// func JSONErrorByCode(w http.ResponseWriter, errorCode faults.ErrorCode) {
// 	message := NewResponseMessage()
// 	message.AddMessageByErrorCode(errorCode)
// 	//jsonMessage, _ := message.GetJsonMessage()
// 	JSON(w, message.GetMessage().HttpStatusCode, message.GetMessage())
// 	// w.WriteHeader(message.GetMessage().HttpStatusCode)
// 	// w.Write(jsonMessage)
// }

// ----------------------------------------------------------------------------------------

// var messageList map[ErrorCode]ResponseBodyMessage

// func _AddMessageByErrorCode(errorCode ErrorCode) (ResponseBodyMessage, error) {
// 	referenceMessage, err := findByErrorCode(errorCode)
// 	if err != nil {
// 		return ResponseBodyMessage{}, err
// 	}

// 	if messageList == nil {
// 		messageList = make(map[ErrorCode]ResponseBodyMessage)
// 	}

// 	var message ResponseBodyMessage
// 	message, exists := messageList[referenceMessage.errorCode]
// 	if !exists {
// 		message = buildMessage(referenceMessage.errorCode, referenceMessage.message)
// 	}

// 	// If the incoming message is exclusive
// 	if referenceMessage.exclusive {
// 		for currentKey := range messageList {
// 			currentMessage, _ := findByErrorCode(currentKey)
// 			if referenceMessage.priority > currentMessage.priority {
// 				// if the incoming message has higher priority over the current message,
// 				// then clean the current message list
// 				ResetMessageList()
// 				break
// 			} else {
// 				// if the incoming message has lower priority over the current message,
// 				// then ignore the current message
// 				return ResponseBodyMessage{}, errors.New("message ignored")
// 			}
// 		}
// 	}

// 	messageList[referenceMessage.errorCode] = message

// 	return message, nil
// }

// func AddMessageByIssue(issue Issue, location Location, field string, value string, descriptionArgs ...string) error {
// 	referenceMessage, referenceMessageDetail, err := findByIssue(issue)
// 	if err != nil {
// 		return err
// 	}

// 	if referenceMessageDetail.locationRequired && location == "" {
// 		return errors.New("location is required")
// 	}
// 	if referenceMessageDetail.fieldRequired && field == "" {
// 		return errors.New("field is required")
// 	}
// 	if referenceMessageDetail.valueRequired && value == "" {
// 		return errors.New("value is required")
// 	}
// 	if referenceMessageDetail.description_args != len(descriptionArgs) {
// 		return fmt.Errorf("wrong number of argumentos passed. expected: [%d] - received: [%d]", referenceMessageDetail.description_args, len(descriptionArgs))
// 	}

// 	message, err := AddMessageByErrorCode(referenceMessage.errorCode)
// 	if err != nil {
// 		return err
// 	}

// 	messageDetail := buildMessageDetail(issue, referenceMessageDetail.description, location, field, value, descriptionArgs)
// 	message.Details = append(message.Details, messageDetail)

// 	messageList[referenceMessage.errorCode] = message

// 	return nil
// }

// func (responseMessage ResponseMessage) GetMessage() ResponseMessage {
// 	return responseMessage
// }

// func ResetMessageList() {
// 	messageList = make(map[ErrorCode]ResponseBodyMessage)
// }

// func buildMessage(errorCode ErrorCode, message string, httpStatusCode int) *ResponseMessage {
// 	return &ResponseMessage{
// 		ErrorCode:      string(errorCode),
// 		Message:        message,
// 		HttpStatusCode: httpStatusCode,
// 	}
// }

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
