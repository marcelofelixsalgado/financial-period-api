package messages

import "errors"

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

func AddMessageByErrorCode(errorCode ErrorCode) error {
	referenceMessage, err := findByErrorCode(errorCode)
	if err != nil {
		return err
	}

	if messageList == nil {
		messageList = make(map[ErrorCode]ResponseBodyMessage)
	}

	var message ResponseBodyMessage
	message, exists := messageList[referenceMessage.errorCode]
	if !exists {
		message = buildMessage(referenceMessage.errorCode, referenceMessage.message)
	}

	messageList[referenceMessage.errorCode] = message

	return nil
}

func AddMessageByIssue(issue Issue, location Location, field string, value string) error {
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

	if messageList == nil {
		messageList = make(map[ErrorCode]ResponseBodyMessage)
	}

	messageDetail := buildMessageDetail(issue, referenceMessageDetail.description, location, field, value)

	var message ResponseBodyMessage
	message, exists := messageList[referenceMessage.errorCode]
	if !exists {
		message = buildMessage(referenceMessage.errorCode, referenceMessage.message)
	}

	message.Details = append(message.Details, messageDetail)

	messageList[referenceMessage.errorCode] = message

	return nil
}

func GetMessages() []ResponseBodyMessage {
	returnList := []ResponseBodyMessage{}
	for _, value := range messageList {
		returnList = append(returnList, value)
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

func buildMessageDetail(issue Issue, description string, location Location, field string, value string) ResponseBodyMessageDetail {
	return ResponseBodyMessageDetail{
		Issue:       string(issue),
		Description: description,
		Location:    location,
		Field:       field,
		Value:       value,
	}
}
