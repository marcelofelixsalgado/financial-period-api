package faults

import "errors"

type ErrorCode string

type Catalog struct {
	List []ReferenceResponse
}

type ReferenceResponse struct {
	ErrorCode      ErrorCode
	Message        string
	HttpStatusCode int
	Details        []ReferenceResponseDetail
}

type ReferenceResponseDetail struct {
	Issue            Issue
	Description      string
	DescriptionArgs  int // Number of argument on description
	LocationRequired bool
	FieldRequired    bool
	ValueRequired    bool
}

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
	ConditionalGreaterThan         Issue = "CONDITIONAL_MUST_BE_GREATER_THAN"
	ConditionalLowerThan           Issue = "CONDITIONAL_MUST_BE_LOWER_THAN"
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

var catalog = Catalog{
	List: []ReferenceResponse{
		{
			ErrorCode:      InvalidRequestSyntax,
			Message:        "Request is not well-formed, syntactically incorrect, or violates schema",
			HttpStatusCode: 400,
			Details: []ReferenceResponseDetail{
				{
					Issue:            DecimalsNotSupported,
					Description:      "Field value does not support decimals",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidBooleanValue,
					Description:      "Field value is invalid. Expected values: true or false",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidParameter,
					Description:      "Cannot be specified as part of the request",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    false,
				},
				{
					Issue:            InvalidStringValue,
					Description:      "Field value is invalid. It should be of type string",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            MalformedRequest,
					Description:      "The request payload is not well formed",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    false,
					ValueRequired:    false,
				},
				{
					Issue:            MissingRequiredField,
					Description:      "A required field is missing",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    false,
				},
			},
		},
		{
			ErrorCode:      UnprocessableEntity,
			Message:        "The request is semantically incorrect or fails business validation",
			HttpStatusCode: 422,
			Details: []ReferenceResponseDetail{
				{
					Issue:            CannotBeNegative,
					Description:      "Must be greater than or equal to zero",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            CannotBeZeroOrNegative,
					Description:      "Must be greater than zero",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            ConditionalFieldNotAllowed,
					Description:      "%s is not allowed when field %s is set to %s",
					DescriptionArgs:  3,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},

				{
					Issue:            ConditionalFieldNotAllowed,
					Description:      "%s is not allowed when field %s is set to %s",
					DescriptionArgs:  3,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            ConditionalGreaterThan,
					Description:      "The field %s must be greater than field %s",
					DescriptionArgs:  2,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            ConditionalLowerThan,
					Description:      "The field %s must be lower than field %s",
					DescriptionArgs:  2,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            ConditionalInvalidValue,
					Description:      "{field1} cannot be set to {value1}, if {field2} is set to {value2}",
					DescriptionArgs:  3,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            ConditionalMissingField,
					Description:      "{field1} is required if {field2} is set to {value2}",
					DescriptionArgs:  3,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    false,
				},
				{
					Issue:            ConditionalValueTooHigh,
					Description:      "{field1} cannot be set to {value1}; it cannot be higher than {field2} ({value2})",
					DescriptionArgs:  4,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            ConditionalValueTooLow,
					Description:      "{field1} cannot be set to {value1}; it cannot be lower than {field2} ({value2})",
					DescriptionArgs:  4,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            FieldValueTooHigh,
					Description:      "Field value cannot be higher than {max}",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            FieldValueTooLow,
					Description:      "Field value cannot be lower than {min}",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidArrayMaxItems,
					Description:      "The number of array items cannot be higher than {max}",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidArrayMinItems,
					Description:      "The number of array items cannot be lower than {min}",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidDateValue,
					Description:      "Field value is invalid. Expected format: YYYY-MM-DD",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidDateTimeValue,
					Description:      "Field value is invalid. Expected format: YYYY-MM-DDThh:mm:ssZ",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidDecimalValue,
					Description:      "Field value is invalid. It should have a value with maximum [N] digits and [N] fractions",
					DescriptionArgs:  2,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidIntegerValue,
					Description:      "Field value should have a value with maximum {N} digits",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidParameterFormat,
					Description:      "Field value does not conform to the expected format: {format}",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidParameterValue,
					Description:      "Field value is invalid",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidParameterValueBlank,
					Description:      "Field value cannot be blank",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    false,
				},
				{
					Issue:            InvalidStringLength,
					Description:      "Field length should be {N} characters",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidStringMaxLength,
					Description:      "Field value exceeded the maximum allowed number of {N} characters",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidStringMinLength,
					Description:      "Field value should be at least {N} characters",
					DescriptionArgs:  1,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidURLValue,
					Description:      "Field value is invalid. It should be in the format of a valid URL",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidUUIDValue,
					Description:      "Invalid UUID string format",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
			},
		},
		{
			ErrorCode:      InvalidClient,
			Message:        "Client authentication failed",
			HttpStatusCode: 401,
			Details: []ReferenceResponseDetail{
				{
					Issue:            AuthenticationFailure,
					Description:      "Authentication failed due to missing authorization header, or invalid authentication credentials",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
			},
		},
		{
			ErrorCode:      NotAuthorized,
			Message:        "Authorization failed due to insufficient permissions",
			HttpStatusCode: 403,
			Details: []ReferenceResponseDetail{
				{
					Issue:            PermissionDenied,
					Description:      "You do not have permission to access or perform operations on this resource",
					DescriptionArgs:  0,
					LocationRequired: false,
					FieldRequired:    false,
					ValueRequired:    false,
				},
				{
					Issue:            RequiredScopeMissing,
					Description:      "Access token does not have required scope",
					DescriptionArgs:  0,
					LocationRequired: false,
					FieldRequired:    false,
					ValueRequired:    false,
				},
			},
		},
		{
			ErrorCode:      ResourceNotFound,
			Message:        "The specified resource does not found",
			HttpStatusCode: 404,
			Details: []ReferenceResponseDetail{
				{
					Issue:            InvalidResourceId,
					Description:      "Specified resource ID does not exist",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
				{
					Issue:            InvalidURI,
					Description:      "The URI requested is invalid",
					DescriptionArgs:  0,
					LocationRequired: false,
					FieldRequired:    false,
					ValueRequired:    false,
				},
				{
					Issue:            NoRecordsFound,
					Description:      "Records not found. Please check the input parameters and try again",
					DescriptionArgs:  0,
					LocationRequired: false,
					FieldRequired:    false,
					ValueRequired:    false,
				},
			},
		},
		{
			ErrorCode:      MethodNotAllowed,
			Message:        "Invalid path and HTTP method combination",
			HttpStatusCode: 405,
			Details: []ReferenceResponseDetail{
				{
					Issue:            MethodNotSupported,
					Description:      "The server does not implement the requested path and HTTP method",
					DescriptionArgs:  0,
					LocationRequired: false,
					FieldRequired:    false,
					ValueRequired:    false,
				},
			},
		},
		{
			ErrorCode:      Conflict,
			Message:        "The request could not be completed due to a conflict with the current state of the target resource",
			HttpStatusCode: 409,
			Details: []ReferenceResponseDetail{
				{
					Issue:            EntityWithSameKeyAlreadyExists,
					Description:      "An entity with the same already exists",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    true,
				},
			},
		},
		{
			ErrorCode:      UnsupportedMediaType,
			Message:        "The server does not support the request body media type",
			HttpStatusCode: 415,
			Details: []ReferenceResponseDetail{
				{
					Issue:            MissingContentType,
					Description:      "A required Content Type header is missing",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    false,
				},
				{
					Issue:            InvalidContentType,
					Description:      "The specified Content Type header is invalid",
					DescriptionArgs:  0,
					LocationRequired: true,
					FieldRequired:    true,
					ValueRequired:    false,
				},
			},
		},
		{
			ErrorCode:      InternalServerError,
			Message:        "A system or application error occurred",
			HttpStatusCode: 500,
		},
		{
			ErrorCode:      BadGateway,
			Message:        "The server returned an invalid response",
			HttpStatusCode: 502,
		},
		{
			ErrorCode:      ServiceUnavailable,
			Message:        "The server cannot handle the request for a service due to temporary maintenance",
			HttpStatusCode: 503,
		},
		{
			ErrorCode:      GatewayTimeout,
			Message:        "The server did not send the response in the expected time",
			HttpStatusCode: 504,
		},
	},
}

func FindByErrorCode(errorCode ErrorCode) (ReferenceResponse, error) {
	for _, value := range catalog.List {
		if value.ErrorCode == errorCode {
			return value, nil
		}
	}
	return ReferenceResponse{}, errors.New("error code not found: " + string(errorCode))
}

func FindByIssue(issue Issue) (ReferenceResponse, ReferenceResponseDetail, error) {
	for _, message := range catalog.List {
		for _, detail := range message.Details {
			if detail.Issue == issue {
				return message, detail, nil
			}
		}
	}
	return ReferenceResponse{}, ReferenceResponseDetail{}, errors.New("issue not found" + string(issue))
}
