package responses_test

import (
	"errors"
	"fmt"
	"marcelofelixsalgado/financial-month-api/api/responses"
	"reflect"
	"testing"
)

type testCase struct {
	issueParameter    responses.Issue
	locationParameter responses.Location
	fieldParameter    string
	valueParameter    string
	errorReturned     string
}

func TestAddMessageSuccess(t *testing.T) {
	testCases := []testCase{
		{
			issueParameter:    responses.DecimalsNotSupported,
			locationParameter: responses.Body,
			fieldParameter:    "field1",
			valueParameter:    "value1",
			errorReturned:     "",
		},
		{
			issueParameter:    responses.DecimalsNotSupported,
			locationParameter: "",
			fieldParameter:    "field1",
			valueParameter:    "value1",
			errorReturned:     "location is required",
		},
		{
			issueParameter:    responses.DecimalsNotSupported,
			locationParameter: responses.Body,
			fieldParameter:    "",
			valueParameter:    "value1",
			errorReturned:     "field is required",
		},
		{
			issueParameter:    responses.DecimalsNotSupported,
			locationParameter: responses.Body,
			fieldParameter:    "field",
			valueParameter:    "",
			errorReturned:     "value is required",
		},
		{
			issueParameter:    responses.InvalidParameter,
			locationParameter: responses.Body,
			fieldParameter:    "field",
			valueParameter:    "",
			errorReturned:     "",
		},
		{
			issueParameter:    responses.MalformedRequest,
			locationParameter: responses.Body,
			fieldParameter:    "",
			valueParameter:    "",
			errorReturned:     "",
		},
	}

	for index, testCase := range testCases {
		// err := messages.AddMessageByIssue(testCase.issueParameter, testCase.locationParameter, testCase.fieldParameter, testCase.valueParameter)
		responseMessage := responses.NewResponseMessage()
		err := responseMessage.AddMessageByIssue(testCase.issueParameter, testCase.locationParameter, testCase.fieldParameter, testCase.valueParameter)
		if err == nil {
			err = errors.New("")
		}
		if fmt.Sprint(err) != testCase.errorReturned {
			t.Errorf("Test Case [%d/%d] - The error [%s] is different from what it was expetected [%s]", index+1, len(testCases), err, testCase.errorReturned)
		}
	}
}

func TestGetMessages(t *testing.T) {

	expectedMessage := responses.ResponseMessage{
		ErrorCode:      "INVALID_REQUEST_SYNTAX",
		Message:        "Request is not well-formed, syntactically incorrect, or violates schema",
		HttpStatusCode: 400,
		Details: []responses.ResponseMessageDetail{
			{
				Issue:       "DECIMALS_NOT_SUPPORTED",
				Description: "Field value does not support decimals",
				Location:    "body",
				Field:       "field3",
				Value:       "value3",
			},
		},
	}

	actualMessage := responses.NewResponseMessage()
	actualMessage.AddMessageByIssue(responses.DecimalsNotSupported, responses.Body, "field3", "value3")

	fmt.Println(actualMessage.GetMessage())

	if !reflect.DeepEqual(actualMessage.GetMessage(), expectedMessage) {
		t.Errorf("Expected message: [%+v]  is not equal Returned Message: [%+v]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesMultipleDetais(t *testing.T) {

	expectedMessage := responses.ResponseMessage{
		ErrorCode:      "INVALID_REQUEST_SYNTAX",
		Message:        "Request is not well-formed, syntactically incorrect, or violates schema",
		HttpStatusCode: 400,
		Details: []responses.ResponseMessageDetail{
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
	}

	actualMessage := responses.NewResponseMessage()
	actualMessage.AddMessageByIssue(responses.DecimalsNotSupported, responses.Body, "field1", "value1")
	actualMessage.AddMessageByIssue(responses.DecimalsNotSupported, responses.Body, "field2", "value2")

	if !reflect.DeepEqual(actualMessage.GetMessage(), expectedMessage) {
		t.Errorf("Expected message: [%+v]  is not equal Returned Message: [%+v]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesReplacementSuccess(t *testing.T) {

	expectedMessage := responses.ResponseMessage{
		ErrorCode:      "UNPROCESSABLE_ENTITY",
		Message:        "The request is semantically incorrect or fails business validation",
		HttpStatusCode: 422,
		Details: []responses.ResponseMessageDetail{
			{
				Issue:       "CONDITIONAL_FIELD_NOT_ALLOWED",
				Description: "field1 is not allowed when field field2 is set to value2",
				Location:    "body",
				Field:       "field1",
				Value:       "value1",
			},
		},
	}

	actualMessage := responses.NewResponseMessage()
	actualMessage.AddMessageByIssue(responses.ConditionalFieldNotAllowed, responses.Body, "field1", "value1", "field1", "field2", "value2")

	if !reflect.DeepEqual(actualMessage.GetMessage(), expectedMessage) {
		t.Errorf("Expected message: [%+v]  is not equal Returned Message: [%+v]", expectedMessage, actualMessage)
	}
}

func TestGetMessagesReplacementFail(t *testing.T) {

	expectedError := fmt.Errorf("wrong number of argumentos passed. expected: [3] - received: [1]")

	actualMessage := responses.NewResponseMessage()
	err := actualMessage.AddMessageByIssue(responses.ConditionalFieldNotAllowed, responses.Body, "field1", "value1", "field2")

	if fmt.Sprint(err) != fmt.Sprint(expectedError) {
		t.Errorf("The error [%s] is different from what it was expetected [%s]", err, expectedError)
	}
}
