package status

type InternalStatus string

const (
	Success             InternalStatus = "OK"
	InvalidResourceId   InternalStatus = "INVALID_RESOURCE_ID"
	NoRecordsFound      InternalStatus = "NO_RECORDS_FOUND"
	InternalServerError InternalStatus = "INVALID_CONTENT_TYPE"
)
