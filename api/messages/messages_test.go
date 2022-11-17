package messages_test

import (
	"errors"
	"fmt"
	"marcelofelixsalgado/financial-month-api/api/messages"
	"reflect"
	"testing"
)

type testCase struct {
	issueParameter    messages.Issue
	locationParameter messages.Location
	fieldParameter    string
	valueParameter    string
	errorReturned     string
}

func TestAddMessageSuccess(t *testing.T) {
	testCases := []testCase{
		{
			issueParameter:    messages.DecimalsNotSupported,
			locationParameter: messages.Body,
			fieldParameter:    "field1",
			valueParameter:    "value1",
			errorReturned:     "",
		},
		{
			issueParameter:    messages.DecimalsNotSupported,
			locationParameter: "",
			fieldParameter:    "field1",
			valueParameter:    "value1",
			errorReturned:     "location is required",
		},
		{
			issueParameter:    messages.DecimalsNotSupported,
			locationParameter: messages.Body,
			fieldParameter:    "",
			valueParameter:    "value1",
			errorReturned:     "field is required",
		},
		{
			issueParameter:    messages.DecimalsNotSupported,
			locationParameter: messages.Body,
			fieldParameter:    "field",
			valueParameter:    "",
			errorReturned:     "value is required",
		},
		{
			issueParameter:    messages.InvalidParameter,
			locationParameter: messages.Body,
			fieldParameter:    "field",
			valueParameter:    "",
			errorReturned:     "",
		},
		{
			issueParameter:    messages.MalformedRequest,
			locationParameter: messages.Body,
			fieldParameter:    "",
			valueParameter:    "",
			errorReturned:     "",
		},
	}

	for index, testCase := range testCases {
		err := messages.AddMessageByIssue(testCase.issueParameter, testCase.locationParameter, testCase.fieldParameter, testCase.valueParameter)
		if err == nil {
			err = errors.New("")
		}
		if fmt.Sprint(err) != testCase.errorReturned {
			t.Errorf("Test Case [%d/%d] - The error [%s] is different from what it was expetected [%s]", index+1, len(testCases), err, testCase.errorReturned)
		}
	}
}

func TestGetMessages(t *testing.T) {

	expectedMessage := []messages.ResponseBodyMessage{
		{
			ErrorCode: "INVALID_REQUEST_SYNTAX",
			Message:   "Request is not well-formed, syntactically incorrect, or violates schema",
			Details: []messages.ResponseBodyMessageDetail{
				{
					Issue:       "DECIMALS_NOT_SUPPORTED",
					Description: "Field value does not support decimals",
					Location:    "body",
					Field:       "field3",
					Value:       "value3",
				},
			},
		},
	}

	messages.ResetMessageList()
	messages.AddMessageByIssue(messages.DecimalsNotSupported, messages.Body, "field3", "value3")

	actualMessage := messages.GetMessages()

	if !reflect.DeepEqual(actualMessage, expectedMessage) {
		t.Errorf("Expected message: [%s]  is not equal Returned Message: [%s]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesMultipleDetais(t *testing.T) {

	expectedMessage := []messages.ResponseBodyMessage{
		{
			ErrorCode: "INVALID_REQUEST_SYNTAX",
			Message:   "Request is not well-formed, syntactically incorrect, or violates schema",
			Details: []messages.ResponseBodyMessageDetail{
				{
					Issue:       "DECIMALS_NOT_SUPPORTED",
					Description: "Field value does not support decimals",
					Location:    "body",
					Field:       "field1",
					Value:       "value1",
				},
				{
					Issue:       "DECIMALS_NOT_SUPPORTED",
					Description: "Field value does not support decimals",
					Location:    "body",
					Field:       "field2",
					Value:       "value2",
				},
			},
		},
	}

	messages.ResetMessageList()
	messages.AddMessageByIssue(messages.DecimalsNotSupported, messages.Body, "field1", "value1")
	messages.AddMessageByIssue(messages.DecimalsNotSupported, messages.Body, "field2", "value2")

	actualMessage := messages.GetMessages()

	if !reflect.DeepEqual(actualMessage, expectedMessage) {
		t.Errorf("Expected message: [%s]  is not equal Returned Message: [%s]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesMultipleErrorCodes(t *testing.T) {

	expectedMessage := []messages.ResponseBodyMessage{
		{
			ErrorCode: "INVALID_REQUEST_SYNTAX",
			Message:   "Request is not well-formed, syntactically incorrect, or violates schema",
			Details: []messages.ResponseBodyMessageDetail{
				{
					Issue:       "DECIMALS_NOT_SUPPORTED",
					Description: "Field value does not support decimals",
					Location:    "body",
					Field:       "field1",
					Value:       "value1",
				},
				{
					Issue:       "DECIMALS_NOT_SUPPORTED",
					Description: "Field value does not support decimals",
					Location:    "header",
					Field:       "field2",
					Value:       "value2",
				},
			},
		},
		{
			ErrorCode: "UNPROCESSABLE_ENTITY",
			Message:   "The request is semantically incorrect or fails business validation",
			Details: []messages.ResponseBodyMessageDetail{
				{
					Issue:       "CANNOT_BE_NEGATIVE",
					Description: "Must be greater than or equal to zero",
					Location:    "query_parameter",
					Field:       "field3",
					Value:       "value3",
				},
			},
		},
	}

	messages.ResetMessageList()
	messages.AddMessageByIssue(messages.DecimalsNotSupported, messages.Body, "field1", "value1")
	messages.AddMessageByIssue(messages.DecimalsNotSupported, messages.Header, "field2", "value2")
	messages.AddMessageByIssue(messages.CannotBeNegative, messages.QueryParameter, "field3", "value3")

	actualMessage := messages.GetMessages()

	if !reflect.DeepEqual(actualMessage, expectedMessage) {
		t.Errorf("Expected message: [%s]  is not equal Returned Message: [%s]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesWithExclusiveTrue(t *testing.T) {

	expectedMessage := []messages.ResponseBodyMessage{
		{
			ErrorCode: "UNSUPPORTED_MEDIA_TYPE",
			Message:   "The server does not support the request body media type",
			Details: []messages.ResponseBodyMessageDetail{
				{
					Issue:       "MISSING_CONTENT_TYPE",
					Description: "A required Content Type header is missing",
					Location:    "header",
					Field:       "Content-Type",
				},
			},
		},
	}

	messages.ResetMessageList()
	messages.AddMessageByIssue(messages.DecimalsNotSupported, messages.Body, "field1", "value1")
	messages.AddMessageByIssue(messages.MissingContentType, messages.Header, "Content-Type", "")

	actualMessage := messages.GetMessages()

	if !reflect.DeepEqual(actualMessage, expectedMessage) {
		t.Errorf("Expected message: [%s]  is not equal Returned Message: [%s]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesReplacementSuccess(t *testing.T) {

	expectedMessage := []messages.ResponseBodyMessage{
		{
			ErrorCode: "UNPROCESSABLE_ENTITY",
			Message:   "The request is semantically incorrect or fails business validation",
			Details: []messages.ResponseBodyMessageDetail{
				{
					Issue:       "CONDITIONAL_FIELD_NOT_ALLOWED",
					Description: "field1 is not allowed when field field2 is set to value2",
					Location:    "body",
					Field:       "field1",
					Value:       "value1",
				},
			},
		},
	}

	messages.ResetMessageList()
	messages.AddMessageByIssue(messages.ConditionalFieldNotAllowed, messages.Body, "field1", "value1", "field1", "field2", "value2")

	actualMessage := messages.GetMessages()

	if !reflect.DeepEqual(actualMessage, expectedMessage) {
		t.Errorf("Expected message: [%s]  is not equal Returned Message: [%s]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesReplacementFail(t *testing.T) {

	expectedError := fmt.Errorf("wrong number of argumentos passed. expected: [3] - received: [1]")

	messages.ResetMessageList()
	err := messages.AddMessageByIssue(messages.ConditionalFieldNotAllowed, messages.Body, "field1", "value1", "field2")

	if fmt.Sprint(err) != fmt.Sprint(expectedError) {
		t.Errorf("The error [%s] is different from what it was expetected [%s]", err, expectedError)
	}
}
