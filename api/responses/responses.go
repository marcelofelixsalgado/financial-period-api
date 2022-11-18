package responses

import (
	"encoding/json"
	"errors"
	"fmt"
)

type IResponseMessage interface {
	NewResponseMessage() *ResponseMessage
	AddMessageByErrorCode(ErrorCode) error
	AddMessageByIssue(Issue, Location, string, string, ...string) error
	GetMessage() ResponseMessage
}

type ResponseMessage struct {
	HttpStatusCode int                     `json:"-"`
	ErrorCode      string                  `json:"error_code"`
	Message        string                  `json:"message"`
	Details        []ResponseMessageDetail `json:"details"`
}

type ResponseMessageDetail struct {
	Issue       string   `json:"issue"`
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Field       string   `json:"field,omitempty"`
	Value       string   `json:"value,omitempty"`
}

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

func (responseMessage *ResponseMessage) GetJsonMessage() ([]byte, error) {
	message := responseMessage.GetMessage()
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to response body: %s", err)
	}
	return messageJSON, nil
}

func (responseMessage *ResponseMessage) AddMessageByErrorCode(errorCode ErrorCode) error {
	referenceMessage, err := findByErrorCode(errorCode)
	if err != nil {
		return err
	}

	responseMessage.ErrorCode = string(referenceMessage.errorCode)
	responseMessage.Message = referenceMessage.message
	responseMessage.HttpStatusCode = referenceMessage.httpStatusCode

	return nil
}

func (responseMessage *ResponseMessage) AddMessageByIssue(issue Issue, location Location, field string, value string, descriptionArgs ...string) error {

	referenceResponse, referenceResponseDetail, err := findByIssue(issue)
	if err != nil {
		return err
	}

	if referenceResponseDetail.locationRequired && location == "" {
		return errors.New("location is required")
	}
	if referenceResponseDetail.fieldRequired && field == "" {
		return errors.New("field is required")
	}
	if referenceResponseDetail.valueRequired && value == "" {
		return errors.New("value is required")
	}
	if referenceResponseDetail.description_args != len(descriptionArgs) {
		return fmt.Errorf("wrong number of argumentos passed. expected: [%d] - received: [%d]", referenceResponseDetail.description_args, len(descriptionArgs))
	}

	if responseMessage.ErrorCode == "" {
		err = responseMessage.AddMessageByErrorCode(referenceResponse.errorCode)
		if err != nil {
			return err
		}
	}

	messageDetail := buildMessageDetail(issue, referenceResponseDetail.description, location, field, value, descriptionArgs)

	responseMessage.Details = append(responseMessage.Details, messageDetail)

	return nil
}

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

func buildMessageDetail(issue Issue, description string, location Location, field string, value string, descriptionArgs []string) ResponseMessageDetail {

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
