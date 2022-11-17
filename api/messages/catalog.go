package messages

import "errors"

type ErrorCode string

const (
	InvalidRequestSyntax ErrorCode = "INVALID_REQUEST_SYNTAX" // HTTP 400 - Request is not well-formed, syntactically incorrect, or violates schema
	UnprocessableEntity  ErrorCode = "UNPROCESSABLE_ENTITY"   // HTTP 422 - The request is semantically incorrect or fails business validation
	InvalidClient        ErrorCode = "INVALID_CLIENT"         // HTTP 401 - Client authentication failed
	NotAuthorized        ErrorCode = "NOT_AUTHORIZED"         // HTTP 403 - Authorization failed due to insufficient permissions
	ResourceNotFound     ErrorCode = "RESOURCE_NOT_FOUND"     // HTTP 404 - The specified resource does not found
	MethodNotAllowed     ErrorCode = "METHOD_NOT_ALLOWED"     // HTTP 405 - Invalid path and HTTP method combination
	Conflict             ErrorCode = "CONFLICT"               // HTTP 409 - The request could not be completed due to a conflict with the current state of the target resource
	UnsupportedMediaType ErrorCode = "UNSUPPORTED_MEDIA_TYPE" // HTTP 415 - The server does not support the request body media type
	InternalServerError  ErrorCode = "INTERNAL_SERVER_ERROR"  // HTTP 500 - A system or application error occurred
	BadGateway           ErrorCode = "BAD_GATEWAY"            // HTTP 502 - The server returned an invalid response
	ServiceUnavailable   ErrorCode = "SERVICE_UNAVAILABLE"    // HTTP 503 - The server cannot handle the request for a service due to temporary maintenance
	GatewayTimeout       ErrorCode = "GATEWAY_TIMEOUT"        // HTTP 504 - The server did not send the response in the expected time
)

type Issue string

const (
	DecimalsNotSupported           Issue = "DECIMALS_NOT_SUPPORTED"
	InvalidBooleanValue            Issue = "INVALID_BOOLEAN_VALUE"
	InvalidParameter               Issue = "INVALID_PARAMETER"
	InvalidStringValue             Issue = "INVALID_STRING_VALUE"
	MalformedRequest               Issue = "MALFORMED_REQUEST"
	MissingRequiredField           Issue = "MISSING_REQUIRED_FIELD"
	CannotBeNegative               Issue = "CANNOT_BE_NEGATIVE"
	CannotBeZeroOrNegative         Issue = "CANNOT_BE_ZERO_OR_NEGATIVE"
	ConditionalFieldNotAllowed     Issue = "CONDITIONAL_FIELD_NOT_ALLOWED"
	ConditionalInvalidValue        Issue = "CONDITIONAL_INVALID_VALUE"
	ConditionalMissingField        Issue = "CONDITIONAL_MISSING_FIELD"
	ConditionalValueTooHigh        Issue = "CONDITIONAL_VALUE_TOO_HIGH"
	ConditionalValueTooLow         Issue = "CONDITIONAL_VALUE_TOO_LOW"
	FieldValueTooHigh              Issue = "FIELD_VALUE_TOO_HIGH"
	FieldValueTooLow               Issue = "FIELD_VALUE_TOO_LOW"
	InvalidArrayMaxItems           Issue = "INVALID_ARRAY_MAX_ITEMS"
	InvalidArrayMinItems           Issue = "INVALID_ARRAY_MIN_ITEMS"
	InvalidDateValue               Issue = "INVALID_DATE_VALUE"
	InvalidDateTimeValue           Issue = "INVALID_DATE_TIME_VALUE"
	InvalidDecimalValue            Issue = "INVALID_DECIMAL_VALUE"
	InvalidIntegerValue            Issue = "INVALID_INTEGER_VALUE"
	InvalidParameterFormat         Issue = "INVALID_PARAMETER_FORMAT"
	InvalidParameterValue          Issue = "INVALID_PARAMETER_VALUE"
	InvalidParameterValueBlank     Issue = "INVALID_PARAMETER_VALUE_BLANK"
	InvalidStringLength            Issue = "INVALID_STRING_LENGTH"
	InvalidStringMaxLength         Issue = "INVALID_STRING_MAX_LENGTH"
	InvalidStringMinLength         Issue = "INVALID_STRING_MIN_LENGTH"
	InvalidURLValue                Issue = "INVALID_URL_VALUE"
	InvalidUUIDValue               Issue = "INVALID_UUID_STRING"
	AuthenticationFailure          Issue = "AUTHENTICATION_FAILURE"
	PermissionDenied               Issue = "PERMISSION_DENIED"
	RequiredScopeMissing           Issue = "REQUIRED_SCOPE_MISSING"
	InvalidResourceId              Issue = "INVALID_RESOURCE_ID"
	InvalidURI                     Issue = "INVALID_URI"
	NoRecordsFound                 Issue = "NO_RECORDS_FOUND"
	MethodNotSupported             Issue = "METHOD_NOT_SUPPORTED"
	EntityWithSameKeyAlreadyExists Issue = "ENTITY_WITH_SAME_KEY_ALREADY_EXISTS"
	MissingContentType             Issue = "MISSING_CONTENT_TYPE"
	InvalidContentType             Issue = "INVALID_CONTENT_TYPE"
)

type Catalog struct {
	List []ReferenceMessage
}

type ReferenceMessage struct {
	errorCode ErrorCode
	message   string
	details   []ReferenceMessageDetail
}

type ReferenceMessageDetail struct {
	issue            Issue
	description      string
	locationRequired bool
	fieldRequired    bool
	valueRequired    bool
}

var catalog = Catalog{
	List: []ReferenceMessage{
		{
			errorCode: InvalidRequestSyntax,
			message:   "Request is not well-formed, syntactically incorrect, or violates schema",
			details: []ReferenceMessageDetail{
				{
					issue:            DecimalsNotSupported,
					description:      "Field value does not support decimals",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidBooleanValue,
					description:      "Field value is invalid. Expected values: true or false",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidParameter,
					description:      "Request is not well-formed, syntactically incorrect, or violates schema",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    false,
				},
				{
					issue:            InvalidStringValue,
					description:      "Field value is invalid. It should be of type string",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            MalformedRequest,
					description:      "The request payload is not well formed",
					locationRequired: true,
					fieldRequired:    false,
					valueRequired:    false,
				},
				{
					issue:            MissingRequiredField,
					description:      "A required field is missing",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    false,
				},
			},
		},
		{
			errorCode: UnprocessableEntity,
			message:   "The request is semantically incorrect or fails business validation",
			details: []ReferenceMessageDetail{
				{
					issue:            CannotBeNegative,
					description:      "Must be greater than or equal to zero",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            CannotBeZeroOrNegative,
					description:      "Must be greater than zero",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            ConditionalFieldNotAllowed,
					description:      "{field1} is not allowed when field {field2} is set to {value2}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            ConditionalInvalidValue,
					description:      "{field1} cannot be set to {value1}, if {field2} is set to {value2}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            ConditionalMissingField,
					description:      "{field1} is required if {field2} is set to {value2}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    false,
				},
				{
					issue:            ConditionalValueTooHigh,
					description:      "{field1} cannot be set to {value1}; it cannot be higher than {field2} ({value2})",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            ConditionalValueTooLow,
					description:      "{field1} cannot be set to {value1}; it cannot be lower than {field2} ({value2})",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            FieldValueTooHigh,
					description:      "Field value cannot be higher than {max}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            FieldValueTooLow,
					description:      "Field value cannot be lower than {min}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidArrayMaxItems,
					description:      "The number of array items cannot be higher than {max}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidArrayMinItems,
					description:      "The number of array items cannot be lower than {min}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidDateValue,
					description:      "Field value is invalid. Expected format: YYYY-MM-DD",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidDateTimeValue,
					description:      "Field value is invalid. Expected format: YYYY-MM-DDThh:mm:ssZ",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidDecimalValue,
					description:      "Field value is invalid. It should have a value with maximum [N] digits and [N] fractions",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidIntegerValue,
					description:      "Field value should have a value with maximum {N} digits",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidParameterFormat,
					description:      "Field value does not conform to the expected format: {format}",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidParameterValue,
					description:      "Field value is invalid",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidParameterValueBlank,
					description:      "Field value cannot be blank",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    false,
				},
				{
					issue:            InvalidStringLength,
					description:      "Field length should be {N} characters",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidStringMaxLength,
					description:      "Field value exceeded the maximum allowed number of {N} characters",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidStringMinLength,
					description:      "Field value should be at least {N} characters",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidURLValue,
					description:      "Field value is invalid. It should be in the format of a valid URL",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidUUIDValue,
					description:      "Invalid UUID string format",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
			},
		},
		{
			errorCode: InvalidClient,
			message:   "Client authentication failed",
			details: []ReferenceMessageDetail{
				{
					issue:            AuthenticationFailure,
					description:      "Authentication failed due to missing authorization header, or invalid authentication credentials",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
			},
		},
		{
			errorCode: NotAuthorized,
			message:   "Authorization failed due to insufficient permissions",
			details: []ReferenceMessageDetail{
				{
					issue:            PermissionDenied,
					description:      "You do not have permission to access or perform operations on this resource",
					locationRequired: false,
					fieldRequired:    false,
					valueRequired:    false,
				},
				{
					issue:            RequiredScopeMissing,
					description:      "Access token does not have required scope",
					locationRequired: false,
					fieldRequired:    false,
					valueRequired:    false,
				},
			},
		},
		{
			errorCode: ResourceNotFound,
			message:   "The specified resource does not found",
			details: []ReferenceMessageDetail{
				{
					issue:            InvalidResourceId,
					description:      "Specified resource ID does not exist",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
				{
					issue:            InvalidURI,
					description:      "The URI requested is invalid",
					locationRequired: false,
					fieldRequired:    false,
					valueRequired:    false,
				},
				{
					issue:            NoRecordsFound,
					description:      "Records not found. Please check the input parameters and try again",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
			},
		},
		{
			errorCode: MethodNotAllowed,
			message:   "Invalid path and HTTP method combination",
			details: []ReferenceMessageDetail{
				{
					issue:            MethodNotSupported,
					description:      "The server does not implement the requested path and HTTP method",
					locationRequired: false,
					fieldRequired:    false,
					valueRequired:    false,
				},
			},
		},
		{
			errorCode: Conflict,
			message:   "The request could not be completed due to a conflict with the current state of the target resource",
			details: []ReferenceMessageDetail{
				{
					issue:            EntityWithSameKeyAlreadyExists,
					description:      "An entity with the same already exists",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    true,
				},
			},
		},
		{
			errorCode: UnsupportedMediaType,
			message:   "The server does not support the request body media type",
			details: []ReferenceMessageDetail{
				{
					issue:            MissingContentType,
					description:      "A required Content Type header is missing",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    false,
				},
				{
					issue:            InvalidContentType,
					description:      "The specified Content Type header is invalid",
					locationRequired: true,
					fieldRequired:    true,
					valueRequired:    false,
				},
			},
		},
		{
			errorCode: InternalServerError,
			message:   "A system or application error occurred",
		},
		{
			errorCode: BadGateway,
			message:   "The server returned an invalid response",
		},
		{
			errorCode: ServiceUnavailable,
			message:   "The server cannot handle the request for a service due to temporary maintenance",
		},
		{
			errorCode: GatewayTimeout,
			message:   "The server did not send the response in the expected time",
		},
	},
}

func findByErrorCode(errorCode ErrorCode) (ReferenceMessage, error) {
	for _, value := range catalog.List {
		if value.errorCode == errorCode {
			return value, nil
		}
	}
	return ReferenceMessage{}, errors.New("error code not found: " + string(errorCode))
}

func findByIssue(issue Issue) (ReferenceMessage, ReferenceMessageDetail, error) {
	for _, message := range catalog.List {
		for _, detail := range message.details {
			if detail.issue == issue {
				return message, detail, nil
			}
		}
	}
	return ReferenceMessage{}, ReferenceMessageDetail{}, errors.New("issue not found" + string(issue))
}
