package messages

import (
	"errors"
	"fmt"
	"sort"
)

type ResponseBodyMessage struct {
	ErrorCode string                      `json:"error_code"`
	Message   string                      `json:"message"`
	Details   []ResponseBodyMessageDetail `json:"details"`
}

type ResponseBodyMessageDetail struct {
	Issue       string   `json:"issue"`
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Field       string   `json:"field"`
	Value       string   `json:"value"`
}

type Location string

const (
	Body           Location = "body"
	Header         Location = "header"
	QueryParameter Location = "query_parameter"
)

var messageList map[ErrorCode]ResponseBodyMessage

func AddMessageByErrorCode(errorCode ErrorCode) (ResponseBodyMessage, error) {
	referenceMessage, err := findByErrorCode(errorCode)
	if err != nil {
		return ResponseBodyMessage{}, err
	}

	if messageList == nil {
		messageList = make(map[ErrorCode]ResponseBodyMessage)
	}

	var message ResponseBodyMessage
	message, exists := messageList[referenceMessage.errorCode]
	if !exists {
		message = buildMessage(referenceMessage.errorCode, referenceMessage.message)
	}

	// If the incoming message is exclusive
	if referenceMessage.exclusive {
		for currentKey, _ := range messageList {
			currentMessage, _ := findByErrorCode(currentKey)
			if referenceMessage.priority > currentMessage.priority {
				// if the incoming message has higher priority over the current message,
				// then clean the current message list
				ResetMessageList()
				break
			} else {
				// if the incoming message has lower priority over the current message,
				// then ignore the current message
				return ResponseBodyMessage{}, errors.New("message ignored")
			}
		}
	}

	messageList[referenceMessage.errorCode] = message

	return message, nil
}

func AddMessageByIssue(issue Issue, location Location, field string, value string, descriptionArgs ...string) error {
	referenceMessage, referenceMessageDetail, err := findByIssue(issue)
	if err != nil {
		return err
	}

	if referenceMessageDetail.locationRequired && location == "" {
		return errors.New("location is required")
	}
	if referenceMessageDetail.fieldRequired && field == "" {
		return errors.New("field is required")
	}
	if referenceMessageDetail.valueRequired && value == "" {
		return errors.New("value is required")
	}
	if referenceMessageDetail.description_args != len(descriptionArgs) {
		return fmt.Errorf("wrong number of argumentos passed. expected: [%d] - received: [%d]", referenceMessageDetail.description_args, len(descriptionArgs))
	}

	message, err := AddMessageByErrorCode(referenceMessage.errorCode)
	if err != nil {
		return err
	}

	messageDetail := buildMessageDetail(issue, referenceMessageDetail.description, location, field, value, descriptionArgs)
	message.Details = append(message.Details, messageDetail)

	messageList[referenceMessage.errorCode] = message

	return nil
}

func GetMessages() []ResponseBodyMessage {

	returnList := []ResponseBodyMessage{}

	// Sorting the list
	keys := make([]string, 0)
	for key, _ := range messageList {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)

	// Building a message list from the sorted ErrorCode sequence
	for _, key := range keys {
		returnList = append(returnList, messageList[ErrorCode(key)])
	}
	return returnList
}

func ResetMessageList() {
	messageList = make(map[ErrorCode]ResponseBodyMessage)
}

func buildMessage(errorCode ErrorCode, message string) ResponseBodyMessage {
	return ResponseBodyMessage{
		ErrorCode: string(errorCode),
		Message:   message,
	}
}

func buildMessageDetail(issue Issue, description string, location Location, field string, value string, descriptionArgs []string) ResponseBodyMessageDetail {

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
	return ResponseBodyMessageDetail{
		Issue:       string(issue),
		Description: description,
		Location:    location,
		Field:       field,
		Value:       value,
	}
}
